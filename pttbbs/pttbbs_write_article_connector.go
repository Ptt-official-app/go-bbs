package pttbbs

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Ptt-official-app/go-bbs"
	"github.com/Ptt-official-app/go-bbs/filelock"
)

func (c *Connector) NewArticleRecord(args map[string]interface{}) (bbs.ArticleRecord, error) {

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

	data := fmt.Sprintf("作者: %s 看板: %s\n標題: %s \n時間: %s\n",
		owner, boardID, title, date)

	if _, err := f.Write([]byte(data)); err != nil {
		return nil, err
	}

	filelock.Unlock(f)

	return record, nil
}

func (c *Connector) AddArticleRecordFileRecord(name string, article bbs.ArticleRecord) error {
	a, ok := article.(*FileHeader)
	if !ok {
		return fmt.Errorf("article should be create with NewArticleRecord")
	}
	return AppendFileHeaderFileRecord(name, a)
}
