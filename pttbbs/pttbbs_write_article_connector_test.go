package pttbbs

import (
	"os"
	"testing"

	"github.com/Ptt-official-app/go-bbs"
)

func TestNewArticleRecord(t *testing.T) {

	c := Connector{"../home/bbs"}

	input := map[string]interface{}{
		"board_id": "SYSOP",
		"owner":    "nickyanggg",
		"date":     "5/15",
		"title":    "Test",
	}

	expected := bbs.ArticleRecord(&FileHeader{
		owner: "nickyanggg",
		date:  "5/15",
		title: "Test",
	})

	actual, err := c.NewArticleRecord(input)
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

	err = os.Remove("../home/bbs/boards/S/SYSOP/" + actual.Filename())
	if err != nil {
		t.Error(err)
	}
}
