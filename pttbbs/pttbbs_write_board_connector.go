package pttbbs

import (
	"fmt"

	"github.com/Ptt-official-app/go-bbs"
)

func (c *Connector) NewBoardRecord(args map[string]interface{}) (bbs.BoardRecord, error) {
	// return fmt.Errorf("not implement")
	record := NewBoardHeader()

	boardID, ok := args["board_id"].(string)
	if !ok {
		return nil, fmt.Errorf("NewBoardRecord: board_id must not be empty")
	}
	record.SetBoardID(boardID)

	title, ok := args["title"].(string)
	if !ok {
		return nil, fmt.Errorf("NewBoardRecord: title must not be empty")
	}
	record.SetTitle(title)

	return record, nil
}

// AddBoardRecordFileRecord given record file name and new record, should append
// file record in that file.
func (c *Connector) AddBoardRecordFileRecord(name string, brd bbs.BoardRecord) error {
	b, ok := brd.(*BoardHeader)
	if !ok {
		return fmt.Errorf("brd should be create with NewBoardRecord")
	}
	return AppendBoardHeaderFileRecord(name, b)
}

// UpdateBoardRecordFileRecord update boardRecord brd on index in record file,
// index is start with 0
func (c *Connector) UpdateBoardRecordFileRecord(name string, index uint, brd bbs.BoardRecord) error {
	return fmt.Errorf("not implement")
}

// ReadBoardRecordFileRecord return boardRecord brd on index in record file.
func (c *Connector) ReadBoardRecordFileRecord(name string, index uint) (bbs.BoardRecord, error) {
	return nil, fmt.Errorf("not implement")
}

// RemoveBoardRecordFileRecord remove boardRecord brd on index in record file.
func (c *Connector) RemoveBoardRecordFileRecord(name string, index uint) error {
	return fmt.Errorf("not implement")
}

var _ bbs.WriteBoardConnector = &Connector{}
