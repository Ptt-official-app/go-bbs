package pttbbs

import (
	"bytes"
	"github.com/Ptt-official-app/go-bbs"
	"os"
	"testing"
	"time"
)

func TestAppendNewLine(t *testing.T) {
	c := Connector{"./testcase"}

	originalModified, _ := time.Parse(time.RFC3339, "2021-05-23T15:04:05Z")
	fileName := "M.1606672292.A.B23"
	file := bbs.ArticleRecord(&FileHeader{
		filename:  fileName,
		owner:     "SYSOP",
		date:      "Sun Nov 29 17:51:30 2020",
		title:     "[討論] 偶爾要發個廢文",
		recommend: 0,
		modified:  originalModified,
	})

	expectedNewLine := "\u001B[1;31m→ \u001B[33mSYSOP\u001B[m\u001B[33m:test                                                   \u001B[m推 01/17 11:36"
	boardPath := "./testcase/boards/t/test/"
	filePath := boardPath + "/" + fileName
	stat, _ := os.Stat(filePath)
	oriFileSize := stat.Size()

	err := c.AppendNewLine(boardPath, file, expectedNewLine)

	if err != nil {
		t.Errorf("Unexpected Error happened: %s", err)
	}

	if file.Modified() == originalModified {
		t.Errorf("Modified does not change, original: %s, \nchanged: %s", originalModified, file.Modified())
	}

	if file.Modified().UnixNano() < originalModified.UnixNano() {
		t.Errorf(
			"Modified is not less tha`n original, original: %d, changed: %d",
			originalModified.UnixNano(), originalModified.UnixNano(),
		)
	}

	stat, _ = os.Stat(filePath)
	newFileSize := stat.Size()
	newLineSize := newFileSize - oriFileSize
	buf := make([]byte, newLineSize)
	fileHandle, _ := os.Open(filePath)
	_, err = fileHandle.ReadAt(buf, oriFileSize)
	actualNewLine := bytes.NewBuffer(buf)

	if expectedNewLine != actualNewLine.String() {
		t.Errorf("newline not matched, expected: %s, \ngot: %s", expectedNewLine, actualNewLine)
	}
}
