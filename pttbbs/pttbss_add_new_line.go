package pttbbs

import (
	"fmt"
	"os"
	"time"

	"github.com/Ptt-official-app/go-bbs/filelock"
)

func (c *Connector) AddNewLine(
	direct *string, article *FileHeader, buf string,
) error {
	path := *direct + "/" + article.Filename()
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
		return fmt.Errorf("DoAddRecommend: File (%s) is locked", path)
	}

	fileHandle, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0752)
	if err != nil {
		return fmt.Errorf("DoAddRecommend: Cannot open file %s, err: %s", path, err)
	}
	defer fileHandle.Close()

	if _, err = fileHandle.WriteString(buf); err != nil {
		return fmt.Errorf("DoAddRecommend: Write to file (%s) failed, err: %s", path, err)
	}

	fileStat, _ := os.Stat(path)
	article.SetModified(fileStat.ModTime())

	// TODO: ENDSTAT
	return nil
}
