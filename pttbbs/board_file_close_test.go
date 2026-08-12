package pttbbs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenBoardHeaderFileClosesFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), ".BRD")
	board := NewBoardHeader()
	board.SetBoardID("SYSOP")
	board.SetTitle("system board")
	if err := AppendBoardHeaderFileRecord(path, board); err != nil {
		t.Fatal(err)
	}

	records, err := OpenBoardHeaderFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(records) != 1 || records[0].BoardID() != "SYSOP" {
		t.Fatalf("unexpected board records: %#v", records)
	}

	// Windows rejects removing an open file, so this catches the leaked handle
	// that Unix-like systems otherwise hide by allowing unlink-on-open.
	if err := os.Remove(path); err != nil {
		t.Fatalf("board file is still open: %v", err)
	}
}
