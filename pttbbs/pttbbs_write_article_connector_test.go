package pttbbs

import (
	"os"
	"testing"

	"github.com/Ptt-official-app/go-bbs"
)

func TestNewArticleRecord(t *testing.T) {

	c := Connector{"./testcase"}
	filename, _ := c.CreateBoardArticleFilename("SYSOP")

	input := map[string]string{
		"filename": filename,
		"owner":    "nickyanggg",
		"date":     " 5/15",
		"title":    "Test",
	}

	expected := bbs.ArticleRecord(&FileHeader{
		filename: filename,
		owner:    "nickyanggg",
		date:     " 5/15",
		title:    "Test",
	})
	actual, err := c.NewArticleRecord(
		input["filename"],
		input["owner"],
		input["date"],
		input["title"],
	)
	if err != nil {
		t.Error(err)
	}

	if expected.Owner() != actual.Owner() {
		t.Errorf("owner not match, expected: %s, got: %s", expected.Owner(), actual.Owner())
	}

	if expected.Date() != actual.Date() {
		t.Errorf("date not match, expected: %s, got: %s", expected.Date(), actual.Date())
	}

	if expected.Title() != actual.Title() {
		t.Errorf("title not match, expected: %s, got: %s", expected.Title(), actual.Title())
	}

	err = os.Remove("./testcase/boards/S/SYSOP/" + actual.Filename())
	if err != nil {
		t.Error(err)
	}
}

// Testing implement write article connector
var _ bbs.WriteArticleConnector = &Connector{}

// func TestImplementWriteArticleConnector(t *testing.T) {
// 	_ :=
// }
