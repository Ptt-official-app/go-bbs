package pttbbs

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/Ptt-official-app/go-bbs"
)

func TestDoAddRecommend(t *testing.T) {
	c := Connector{"./testcase"}

	originalModified, _ := time.Parse(time.RFC3339, "2021-05-23T15:04:05Z")
	fileName := "M.1606672292.A.B23"
	file := &FileHeader{
		filename:  fileName,
		owner:     "SYSOP",
		date:      "Sun Nov 29 17:51:30 2020",
		title:     "[討論] 偶爾要發個廢文",
		recommend: 0,
		modified:  originalModified,
	}

	expectedRecommend := 0
	expectedNewLine := "\u001B[1;31m→ \u001B[33mSYSOP\u001B[m\u001B[33m:test                                                   \u001B[m推 01/17 11:36"
	boardPath := "./testcase/boards/t/test/"
	filePath := boardPath + "/" + fileName
	stat, _ := os.Stat(filePath)
	oriFileSize := stat.Size()

	err := c.DoAddRecommend(&boardPath, file, 0, &expectedNewLine, bbs.RecommendTypeArrow)

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

	// input is arrow type
	if file.Recommend() != expectedRecommend {
		t.Errorf("recommend not match, expected: %d, \ngot: %d", expectedRecommend, file.Recommend())
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

func TestDoAddRecommendUpdateRecommend(t *testing.T) {
	c := Connector{"./testcase"}

	originalModified, _ := time.Parse(time.RFC3339, "2021-05-23T15:04:05Z")
	fileName := "M.1606672292.A.B23"
	file := &FileHeader{
		filename:  fileName,
		owner:     "SYSOP",
		date:      "Sun Nov 29 17:51:30 2020",
		title:     "[討論] 偶爾要發個廢文",
		recommend: 0,
		modified:  originalModified,
	}

	expectedNewLine := "\u001B[1;31m→ \u001B[33mSYSOP\u001B[m\u001B[33m:test                                                   \u001B[m推 01/17 11:36"
	boardPath := "./testcase/boards/t/test/"

	// 0 + 0 = 0
	_ = c.DoAddRecommend(&boardPath, file, 0, &expectedNewLine, bbs.RecommendTypeArrow)
	if file.Recommend() != 0 {
		t.Errorf("RecommendTypeArrow should plus 0")
	}

	// 0 + 1 = 1
	_ = c.DoAddRecommend(&boardPath, file, 0, &expectedNewLine, bbs.RecommendTypeGood)
	if file.Recommend() != 1 {
		t.Errorf("RecommendTypeGood should plus 1")
	}

	// 1 - 1 = 0
	_ = c.DoAddRecommend(&boardPath, file, 0, &expectedNewLine, bbs.RecommendTypeBad)
	if file.Recommend() != 0 {
		t.Errorf("RecommendTypeBad should minus 1")
	}

	// MaxRecommend == 100
	file.AddRecommend(100)
	_ = c.DoAddRecommend(&boardPath, file, 0, &expectedNewLine, bbs.RecommendTypeGood)
	if file.Recommend() != 100 {
		t.Errorf("MaxRecommend count is 100")
	}

	// -MaxRecommend == -100
	file.AddRecommend(-100)
	file.AddRecommend(-100)
	_ = c.DoAddRecommend(&boardPath, file, 0, &expectedNewLine, bbs.RecommendTypeBad)
	if file.Recommend() != -100 {
		t.Errorf("-MaxRecommend count is -100")
	}

}