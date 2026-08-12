package pttbbs

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/Ptt-official-app/go-bbs"
)

func TestDeleteBoardArticle(t *testing.T) {
	home := t.TempDir()
	boardID := "SYSOP"
	boardDir := filepath.Join(home, "boards", "S", boardID)
	if err := os.MkdirAll(boardDir, 0755); err != nil {
		t.Fatal(err)
	}

	c := Connector{home: home}
	filename := "M.1723423500.A.123"
	record := NewFileHeader()
	record.SetFilename(filename)
	record.SetModified(time.Unix(1723423500, 0))
	record.AddRecommend(7)
	record.SetOwner("SYSOP")
	record.SetDate(" 8/12")
	record.SetTitle("delete me")

	dirPath := filepath.Join(boardDir, ".DIR")
	if err := AppendFileHeaderFileRecord(dirPath, record); err != nil {
		t.Fatal(err)
	}
	articlePath := filepath.Join(boardDir, filename)
	if err := os.WriteFile(articlePath, []byte("article body"), 0644); err != nil {
		t.Fatal(err)
	}

	if err := c.DeleteBoardArticle(boardID, filename, "SYSOP"); err != nil {
		t.Fatal(err)
	}

	records, err := OpenFileHeaderFile(dirPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(records) != 1 {
		t.Fatalf("expected one tombstone record, got %d", len(records))
	}
	got := records[0]
	if got.Filename() != safeDeleteFilename {
		t.Errorf("filename = %q, want %q", got.Filename(), safeDeleteFilename)
	}
	if got.Owner() != "-" {
		t.Errorf("owner = %q, want -", got.Owner())
	}
	if got.Title() != "(本文已被刪除) [SYSOP]" {
		t.Errorf("title = %q", got.Title())
	}
	if got.Recommend() != 7 {
		t.Errorf("recommend = %d, want 7", got.Recommend())
	}
	if _, err := os.Stat(articlePath); !os.IsNotExist(err) {
		t.Fatalf("original article body still exists: %v", err)
	}
}

func TestDeleteBoardArticleByModerator(t *testing.T) {
	home := t.TempDir()
	boardID := "SYSOP"
	boardDir := filepath.Join(home, "boards", "S", boardID)
	if err := os.MkdirAll(boardDir, 0755); err != nil {
		t.Fatal(err)
	}

	c := Connector{home: home}
	filename := "M.1723423501.A.124"
	record := NewFileHeader()
	record.SetFilename(filename)
	record.SetOwner("author")
	record.SetTitle("delete me")
	if err := AppendFileHeaderFileRecord(filepath.Join(boardDir, ".DIR"), record); err != nil {
		t.Fatal(err)
	}

	if err := c.DeleteBoardArticle(boardID, filename, "SYSOP"); err != nil {
		t.Fatal(err)
	}

	records, err := OpenFileHeaderFile(filepath.Join(boardDir, ".DIR"))
	if err != nil {
		t.Fatal(err)
	}
	if got := records[0].Title(); got != "(本文已被刪除) <author>" {
		t.Errorf("title = %q", got)
	}
}

func TestDeleteBoardArticleRejectsPathTraversal(t *testing.T) {
	c := Connector{home: t.TempDir()}
	for _, filename := range []string{"../.PASSWDS", `..\\.PASSWDS`, "dir/article"} {
		if err := c.DeleteBoardArticle("SYSOP", filename, "SYSOP"); err == nil {
			t.Errorf("DeleteBoardArticle(%q) unexpectedly succeeded", filename)
		}
	}
}

func TestDeleteBoardArticleNotFound(t *testing.T) {
	home := t.TempDir()
	boardDir := filepath.Join(home, "boards", "S", "SYSOP")
	if err := os.MkdirAll(boardDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(boardDir, ".DIR"), nil, 0644); err != nil {
		t.Fatal(err)
	}

	c := Connector{home: home}
	err := c.DeleteBoardArticle("SYSOP", "M.1.A.001", "SYSOP")
	if err == nil || !strings.Contains(err.Error(), "not found") {
		t.Fatalf("expected not found error, got %v", err)
	}
}

var _ bbs.DeleteArticleConnector = &Connector{}
