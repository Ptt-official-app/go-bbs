package pttbbs

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Ptt-official-app/go-bbs"
)

func TestAddUserFavorite(t *testing.T) {
	home := t.TempDir()
	connector := &Connector{home: home}
	writeFavoriteTestBoard(t, home, "SYSOP")
	favPath := prepareFavoriteTestUser(t, connector, "pichu")

	line, err := connector.AddUserFavorite("pichu", bbs.FavoriteCreateOptions{Type: bbs.FavoriteTypeLine})
	if err != nil {
		t.Fatal(err)
	}
	if line.Type() != bbs.FavoriteTypeLine {
		t.Fatalf("line type = %v", line.Type())
	}

	folder, err := connector.AddUserFavorite("pichu", bbs.FavoriteCreateOptions{
		Type:  bbs.FavoriteTypeFolder,
		Title: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
	if folder.Title() != "test" || folder.Type() != bbs.FavoriteTypeFolder {
		t.Fatalf("folder = type %v title %q", folder.Type(), folder.Title())
	}

	board, err := connector.AddUserFavorite("pichu", bbs.FavoriteCreateOptions{
		Type:    bbs.FavoriteTypeBoard,
		BoardID: "sysop",
	})
	if err != nil {
		t.Fatal(err)
	}
	if board.BoardID() != "SYSOP" || board.Type() != bbs.FavoriteTypeBoard {
		t.Fatalf("board = type %v id %q", board.Type(), board.BoardID())
	}

	fav, err := OpenFavFile(favPath)
	if err != nil {
		t.Fatal(err)
	}
	if fav.Version != pttFavoriteVersion {
		t.Errorf("version = %d, want %d", fav.Version, pttFavoriteVersion)
	}
	if fav.Folder.NLines != 1 || fav.Folder.NFolders != 1 || fav.Folder.NBoards != 1 {
		t.Errorf("counts = boards:%d folders:%d lines:%d", fav.Folder.NBoards, fav.Folder.NFolders, fav.Folder.NLines)
	}
	if len(fav.Folder.FavItems) != 3 {
		t.Fatalf("item count = %d, want 3", len(fav.Folder.FavItems))
	}
	if fav.Folder.FavItems[0].FavType != FavItemTypeLine || fav.Folder.FavItems[0].FavAttr&uint8(FavhFav) == 0 {
		t.Errorf("first item is not a valid favorite line")
	}
	if got := fav.Folder.FavItems[1].GetFolder(); got == nil || got.Title != "test" || got.FolderID != 1 || got.ThisFolder == nil {
		t.Errorf("persisted folder = %#v", got)
	}
	if got := fav.Folder.FavItems[2].GetBoard(); got == nil || got.BoardID != 1 {
		t.Errorf("persisted board = %#v", got)
	}
}

func TestAddUserFavoriteRejectsDuplicateBoard(t *testing.T) {
	home := t.TempDir()
	connector := &Connector{home: home}
	writeFavoriteTestBoard(t, home, "SYSOP")
	prepareFavoriteTestUser(t, connector, "pichu")

	options := bbs.FavoriteCreateOptions{Type: bbs.FavoriteTypeBoard, BoardID: "SYSOP"}
	if _, err := connector.AddUserFavorite("pichu", options); err != nil {
		t.Fatal(err)
	}
	if _, err := connector.AddUserFavorite("pichu", options); err == nil || !strings.Contains(err.Error(), "already in favorites") {
		t.Fatalf("expected duplicate board error, got %v", err)
	}
}

func TestAddUserFavoriteValidatesFolderTitle(t *testing.T) {
	home := t.TempDir()
	connector := &Connector{home: home}
	prepareFavoriteTestUser(t, connector, "pichu")

	for _, title := range []string{"", strings.Repeat("測", 25)} {
		if _, err := connector.AddUserFavorite("pichu", bbs.FavoriteCreateOptions{
			Type:  bbs.FavoriteTypeFolder,
			Title: title,
		}); err == nil {
			t.Errorf("folder title %q unexpectedly succeeded", title)
		}
	}
}

func TestAppendFavoriteLineLimit(t *testing.T) {
	folder := newEmptyFavoriteFile().Folder
	for i := 0; i < pttFavoriteMaxLines; i++ {
		if _, err := appendFavoriteLine(folder); err != nil {
			t.Fatalf("line %d: %v", i, err)
		}
	}
	if _, err := appendFavoriteLine(folder); err == nil {
		t.Fatal("65th favorite line unexpectedly succeeded")
	}
}

func TestAddUserFavoriteRejectsUserPathTraversal(t *testing.T) {
	connector := &Connector{home: t.TempDir()}
	for _, userID := range []string{"../SYSOP", `..\\SYSOP`, "home/SYSOP"} {
		if _, err := connector.AddUserFavorite(userID, bbs.FavoriteCreateOptions{Type: bbs.FavoriteTypeLine}); err == nil {
			t.Errorf("user id %q unexpectedly succeeded", userID)
		}
	}
}

func writeFavoriteTestBoard(t *testing.T, home, boardID string) {
	t.Helper()
	board := NewBoardHeader()
	board.SetBoardID(boardID)
	board.SetTitle("test board")
	path, err := GetBoardPath(home)
	if err != nil {
		t.Fatal(err)
	}
	if err := AppendBoardHeaderFileRecord(path, board); err != nil {
		t.Fatal(err)
	}
}

func prepareFavoriteTestUser(t *testing.T, connector *Connector, userID string) string {
	t.Helper()
	path, err := connector.GetUserFavoriteRecordsPath(userID)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0750); err != nil {
		t.Fatal(err)
	}
	return path
}

var _ bbs.WriteFavoriteConnector = &Connector{}
