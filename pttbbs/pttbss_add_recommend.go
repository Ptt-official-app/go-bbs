package pttbbs

import (
	"fmt"
	"os"
	"time"

	"github.com/Ptt-official-app/go-bbs"
	"github.com/Ptt-official-app/go-bbs/filelock"
)

func (c *Connector) DoAddRecommend(
	direct *string, article *FileHeader, ent int, buf *string, recommendType bbs.RecommendType,
) error {

	var update int8
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

	if _, err = fileHandle.WriteString(*buf); err != nil {
		return fmt.Errorf("DoAddRecommend: Write to file (%s) failed, err: %s", path, err)
	}
	if recommendType == bbs.RecommendTypeGood && article.Recommend() < MaxRecommends {
		update++
	} else if recommendType == bbs.RecommendTypeBad && article.Recommend() > -MaxRecommends {
		update--
	}
	article.AddRecommend(update)

	fileStat, _ := os.Stat(path)
	article.SetModified(fileStat.ModTime())

	if article.Modified().Unix() > 0 {
		// TODO: modify_dir_lite
		// TODO: brc_addlist
	}
	// TODO: ENDSTAT
	return nil
}
