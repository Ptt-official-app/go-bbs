package pttbbs

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Ptt-official-app/go-bbs/filelock"
)

const (
	safeDeleteFilename = ".d"
	safeDeleteTitle    = "(本文已被刪除)"
)

// DeleteBoardArticle implements PTT SAFE_ARTICLE_DELETE semantics: the
// original .DIR record stays in place as a tombstone while the article body is
// made unavailable from its original filename.
func (c *Connector) DeleteBoardArticle(boardID, filename, deletedBy string) error {
	if filename == "" || filepath.Base(filename) != filename || strings.Contains(filename, `\`) {
		return fmt.Errorf("pttbbs: invalid article filename %q", filename)
	}

	recordsPath, err := c.GetBoardArticleRecordsPath(boardID)
	if err != nil {
		return fmt.Errorf("pttbbs: get article records path: %w", err)
	}
	articlePath, err := c.GetBoardArticleFilePath(boardID, filename)
	if err != nil {
		return fmt.Errorf("pttbbs: get article path: %w", err)
	}

	f, err := os.OpenFile(recordsPath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("pttbbs: open article records: %w", err)
	}
	defer f.Close()

	if err := filelock.Lock(f); err != nil {
		return fmt.Errorf("pttbbs: lock article records: %w", err)
	}
	defer filelock.Unlock(f)

	for index := uint(0); ; index++ {
		raw := make([]byte, 128)
		_, err := io.ReadFull(f, raw)
		switch err {
		case nil:
		case io.EOF:
			return fmt.Errorf("pttbbs: article %q not found", filename)
		case io.ErrUnexpectedEOF:
			return fmt.Errorf("pttbbs: truncated article records file")
		default:
			return fmt.Errorf("pttbbs: read article record: %w", err)
		}

		header, err := NewFileHeaderWithByte(raw)
		if err != nil {
			return fmt.Errorf("pttbbs: decode article record: %w", err)
		}
		if header.Filename() != filename {
			continue
		}

		stagedPath, err := stageArticleForDelete(articlePath)
		if err != nil {
			return fmt.Errorf("pttbbs: stage article body for deletion: %w", err)
		}

		offset := int64(index) * 128
		tombstone := makeSafeDeleteRecord(raw, header, deletedBy, time.Now())
		if err := writeArticleRecordAt(f, offset, tombstone); err != nil {
			restoreStagedArticle(stagedPath, articlePath)
			return err
		}
		if err := f.Sync(); err != nil {
			_ = writeArticleRecordAt(f, offset, raw)
			_ = f.Sync()
			restoreStagedArticle(stagedPath, articlePath)
			return fmt.Errorf("pttbbs: sync article tombstone: %w", err)
		}

		if stagedPath != "" {
			if err := os.Remove(stagedPath); err != nil {
				return fmt.Errorf("pttbbs: article tombstoned but body cleanup failed: %w", err)
			}
		}
		return nil
	}
}

func makeSafeDeleteRecord(raw []byte, header *FileHeader, deletedBy string, now time.Time) []byte {
	out := append([]byte(nil), raw...)

	setFixedField(out[PosOfFileHeaderFilename:PosOfFileHeaderFilename+FileNameLength], []byte(safeDeleteFilename))
	binary.LittleEndian.PutUint32(out[PosOfFileHeaderModified:PosOfFileHeaderModified+4], uint32(now.Unix()))
	setFixedField(out[PosOfFileHeaderOwner:PosOfFileHeaderOwner+IDLength+2], []byte("-"))
	setFixedField(out[PosOfFileHeaderTitle:PosOfFileHeaderTitle+TitleLength+1], utf8ToBig5UAOString(safeDeleteRecordTitle(header, deletedBy)))

	return out
}

func safeDeleteRecordTitle(header *FileHeader, deletedBy string) string {
	owner := header.Owner()
	if header.Filemode&FileAnonymous != 0 || owner == "" || owner == "-" {
		return safeDeleteTitle
	}
	if owner == deletedBy {
		return fmt.Sprintf("%s [%s]", safeDeleteTitle, owner)
	}
	return fmt.Sprintf("%s <%s>", safeDeleteTitle, owner)
}

func setFixedField(dst, value []byte) {
	for i := range dst {
		dst[i] = 0
	}
	copy(dst, value)
}

func writeArticleRecordAt(f *os.File, offset int64, data []byte) error {
	if _, err := f.Seek(offset, io.SeekStart); err != nil {
		return fmt.Errorf("pttbbs: seek article record: %w", err)
	}
	if _, err := f.Write(data); err != nil {
		return fmt.Errorf("pttbbs: write article tombstone: %w", err)
	}
	return nil
}

func stageArticleForDelete(articlePath string) (string, error) {
	if _, err := os.Stat(articlePath); err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}

	tmp, err := os.CreateTemp(filepath.Dir(articlePath), ".go-bbs-delete-*")
	if err != nil {
		return "", err
	}
	stagedPath := tmp.Name()
	if err := tmp.Close(); err != nil {
		_ = os.Remove(stagedPath)
		return "", err
	}
	if err := os.Remove(stagedPath); err != nil {
		return "", err
	}
	if err := os.Rename(articlePath, stagedPath); err != nil {
		return "", err
	}
	return stagedPath, nil
}

func restoreStagedArticle(stagedPath, articlePath string) {
	if stagedPath != "" {
		_ = os.Rename(stagedPath, articlePath)
	}
}
