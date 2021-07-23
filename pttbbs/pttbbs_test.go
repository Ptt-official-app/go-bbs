package pttbbs

import (
	"testing"
)

func TestReadUserFavoriteRecordsFileNoFile(t *testing.T) {
	c := Connector{}
	recs, err := c.ReadUserFavoriteRecordsFile("test/.fav")

	// https://github.com/Ptt-official-app/Ptt-backend/issues/235
	if err != nil || len(recs) != 0 {
		t.Errorf("not return empty favorite records")
	}
}
