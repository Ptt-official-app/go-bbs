package bbs

import (
	"testing"
)

func TestSomething(t *testing.T) {
	_, err := OpenBoardHeaderFile("testcase/board/00.BRD")
	if err != nil {
		t.Error(err)
	}

}
