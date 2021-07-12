package pttbbs

import (
	"fmt"
	"os"
	"time"

	"github.com/Ptt-official-app/go-bbs"

	"github.com/Ptt-official-app/go-bbs/filelock"
)

func (c *Connector) AppendNewLine(
	boardPath string, article bbs.ArticleRecord, buf string,
) error {
	path := boardPath + "/" + article.Filename()
	lockRetry, lockWait, lockSuccess := 5, time.Duration(1), false
	// TODO: STARTSTAT

	for lockRetry > 0 {
		lockRetry--
		if filelock.IsLock(path) {
			fmt.Printf("File is locked, please wait. Retry: %d", lockRetry)
			time.Sleep(lockWait)
			continue
		}
		lockSuccess = true
	}
	if !lockSuccess {
		return fmt.Errorf("AppendNewLine: File (%s) is locked", path)
	}

	fileHandle, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o752)
	if err != nil {
		return fmt.Errorf("AppendNewLine: Cannot open file %s, err: %s", path, err)
	}
	defer fileHandle.Close()

	if _, err = fileHandle.WriteString(buf); err != nil {
		return fmt.Errorf("AppendNewLine: Write to file (%s) failed, err: %s", path, err)
	}

	fileStat, _ := os.Stat(path)
	article.SetModified(fileStat.ModTime())

	// TODO: ENDSTAT
	return nil
}
