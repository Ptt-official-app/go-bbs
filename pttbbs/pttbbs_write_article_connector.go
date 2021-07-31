package pttbbs

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Ptt-official-app/go-bbs"
	"github.com/Ptt-official-app/go-bbs/filelock"
)

// CreateBoardArticleFilename get available filename for board with boardID, it will test is this filename not exist
// And open a file to occupy this filename
// Please see fhdr_stamp in pttbbs fhdr_stamp.c also
func (c *Connector) CreateBoardArticleFilename(boardID string) (filename string, err error) {
	var f *os.File
	for {
		dtime := time.Now().Unix()
		// TOOD: Check 2038 Problem
		filename = fmt.Sprintf("M.%d.A.%3.3X", dtime, rand.Intn(0x1000))
		path, err := c.GetBoardArticleFilePath(boardID, filename)
		if err != nil {
			return filename, err
		}

		f, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
		if err == nil {
			break
		}

		// TODO: Should log if can not lock file in first time, is system loading too heavy?

	}
	f.Close()
	return

}

// NewArticleRecord returns a new ArticleRecord given a filename, owner, date, title
func (c *Connector) NewArticleRecord(filename, owner, date, title string) (bbs.ArticleRecord, error) {
	ret := &FileHeader{
		filename: filename,
		owner:    owner,
		date:     date,
		title:    title,
	}
	return ret, nil
}

func (c *Connector) WriteBoardArticleFile(path string, content []byte) error {

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("openfile: %w", err)

	}
	err = filelock.Lock(f)
	if err != nil {
		// File is locked
		return fmt.Errorf("filelock lock: %w", err)
	}
	defer filelock.Unlock(f)

	if _, err := f.Write(content); err != nil {
		return err
	}
	return nil

}

func (c *Connector) NewArticleRecordWithMap(args map[string]interface{}) (bbs.ArticleRecord, error) {

	record := NewFileHeader()

	owner, ok := args["owner"].(string)
	if !ok {
		return nil, fmt.Errorf("NewArticleRecord: owner must not be empty")
	}
	record.SetOwner(owner)

	date, ok := args["date"].(string)
	if !ok {
		return nil, fmt.Errorf("NewArticleRecord: date must not be empty")
	}
	record.SetDate(date)

	title, ok := args["title"].(string)
	if !ok {
		return nil, fmt.Errorf("NewArticleRecord: title must not be empty")
	}
	record.SetTitle(title)

	boardID, ok := args["board_id"].(string)
	if !ok {
		return nil, fmt.Errorf("NewArticleRecord: board_id must not be empty")
	}

	filename := ""
	dtime := time.Now().Unix()
	rand.Seed(dtime)

	var f *os.File
	for {
		filename = fmt.Sprintf("M.%d.A.%3.3X", dtime, rand.Intn(0x1000))
		path, err := c.GetBoardArticleFilePath(boardID, filename)
		if err != nil {
			return nil, err
		}

		f, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
		if err == nil {
			break
		}
	}
	defer f.Close()

	record.SetFilename(filename)

	err := filelock.Lock(f)
	if err != nil {
		// File is locked
		return nil, err
	}
	defer filelock.Unlock(f)

	data := fmt.Sprintf("作者: %s 看板: %s\n標題: %s \n時間: %s\n",
		owner, boardID, title, date)

	if _, err := f.Write([]byte(data)); err != nil {
		return nil, err
	}

	return record, nil
}

func (c *Connector) AddArticleRecordFileRecord(name string, article bbs.ArticleRecord) error {
	a, ok := article.(*FileHeader)
	if !ok {
		return fmt.Errorf("article should be create with NewArticleRecord")
	}
	return AppendFileHeaderFileRecord(name, a)
}
