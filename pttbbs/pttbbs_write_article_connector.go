package pttbbs

import (
	"github.com/Ptt-official-app/go-bbs"

	"fmt"
	"math/rand"
	"os"
	"time"
)

func (c *Connector) NewArticleRecord(args map[string]interface{}) (bbs.ArticleRecord, error) {

	record := NewFileHeader()

	boardID, ok := args["board_id"].(string)
	if !ok {
		return nil, fmt.Errorf("NewArticleRecord: board_id must not be empty")
	}

	filename := ""
	dtime := time.Now().Unix()
	rand.Seed(dtime)

	for {
		filename = fmt.Sprintf("M.%d.A.%3.3X", dtime, rand.Intn(0x1000))
		path, err := c.GetBoardArticleFilePath(boardID, filename)
		if err != nil {
			return nil, err
		}
		// check if filename exists already
		if _, err := os.Stat(path); err != nil {
			break
		}
	}

	record.SetFilename(filename)

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

	return record, nil
}

func (c *Connector) AddArticleRecordFileRecord(name string, article bbs.ArticleRecord) error {
	a, ok := article.(*FileHeader)
	if !ok {
		return fmt.Errorf("article should be create with NewArticleRecord")
	}
	return AppendFileHeaderFileRecord(name, a)
}
