package pttbbs

import (
	"fmt"
	"os"

	"github.com/Ptt-official-app/go-bbs"
)

func (c *Connector) AppendBoardArticleFile(filename string, content []byte) error {
	// TODO: Lockfile

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("AppendNewLine: Cannot open file %s, err: %s", filename, err)
	}
	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		return fmt.Errorf("AppendNewLine: Write to file (%s) failed, err: %w", filename, err)
	}
	return nil
}

func (c *Connector) UpdateArticleRecordFileRecord(filename string, index uint, article bbs.ArticleRecord) error {
	a, ok := article.(*FileHeader)
	if !ok {
		return fmt.Errorf("article should be create with NewArticleRecord or get by ReadArticleRecordFileRecord")
	}
	return UpdateFileHeaderFileRecord(filename, index, a)
}
