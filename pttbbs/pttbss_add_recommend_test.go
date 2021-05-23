package pttbbs

import (
	"bytes"
	"os"
	"testing"
	"time"


	"github.com/Ptt-official-app/go-bbs"
)

func removeNewLines(filePath string, start int64) {
	fileHandle, _ := os.Open(filePath)
	empty := make([]byte, 0)
	fileHandle.WriteAt(empty, start)
}

func TestDoAddRecommend(t *testing.T) {

	newLineCount := 0
	c := Connector{"./testcase"}

	originalModified, _ := time.Parse(time.RFC3339, "2021-05-23T15:04:05Z")
	fileName := "M.1606672292.A.B23"
	file := &FileHeader{
		filename: fileName,
		owner: "SYSOP",
		date:  "Sun Nov 29 17:51:30 2020",
		title: "[討論] 偶爾要發個廢文",
		recommend: 0,
		modified: originalModified,
	}

	expectedRecommend := 0
	expectedNewLine := "\u001B[1;31m→ \u001B[33mSYSOP\u001B[m\u001B[33m:test                                                   \u001B[m推 01/17 11:36"
	boardPath := "./testcase/boards/t/test/"
	filePath := boardPath + "/" + fileName
	stat, _ := os.Stat(filePath)
	oriFileSize := stat.Size()

	err := c.DoAddRecommend(&boardPath, file, 0, &expectedNewLine, bbs.RecommendTypeArrow)
	newLineCount++

	if err != nil {
		t.Errorf("Unexpected Error happened: %s", err)
	}

	if file.Modified() != originalModified {
		t.Errorf("Modified does not change, original: %s, changed: %s", originalModified, file.Modified())
	}

	if file.Modified().UnixNano() < originalModified.UnixNano() {
		t.Errorf(
			"Modified is not less than original, original: %d, changed: %d",
			originalModified.UnixNano(), originalModified.UnixNano(),
			)
	}

	// input is arrow type
	if file.Recommend() != expectedRecommend {
		t.Errorf("recommend not match, expected: %d, got: %d", expectedRecommend, file.Recommend())
	}

	stat, _ = os.Stat(filePath)
	newFileSize := stat.Size()
	newLineSize := newFileSize - oriFileSize
	buf := make([]byte, newLineSize)
	fileHandle, _ := os.Open(filePath)
	_, err = fileHandle.ReadAt(buf, oriFileSize)
	actualNewLine := bytes.NewBuffer(buf)

	if expectedNewLine == actualNewLine.String() {
		t.Errorf("newline not matched, expected: %s, got: %s", expectedNewLine, actualNewLine)
	}

	removeNewLines(filePath, oriFileSize)
}
