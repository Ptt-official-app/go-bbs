package pttbbs

import (
	"github.com/PichuChen/go-bbs"
)

type Connector struct {
	home string
}

func init() {
	bbs.Register("pttbbs", &Connector{})
}

func (c *Connector) GetUserRecordsPath() (string, error) {
	return GetPasswdsPath(c.home)
}

func (c *Connector) ReadUserRecordsFile(path string) ([]bbs.UserRecord, error) {
	rec, err := OpenUserecFile(path)
	ret := make([]bbs.UserRecord, len(rec))
	for i, v := range rec {
		ret[i] = v
	}
	return ret, err
}

func (c *Connector) GetBoardRecordsPath() (string, error) {
	return GetBoardPath(c.home)
}

func (c *Connector) ReadBoardRecordsFile(path string) ([]bbs.BoardRecord, error) {
	rec, err := OpenBoardHeaderFile(path)
	ret := make([]bbs.BoardRecord, len(rec))
	for i, v := range rec {
		ret[i] = v
	}
	return ret, err
}
