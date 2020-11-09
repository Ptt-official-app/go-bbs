package cache

import "github.com/PichuChen/go-bbs/ptttype"

const (
	testShmKey = 2000000
)

var (
	IsTest = false

	origBBSHome = ""
)

func setupTest() {
	IsTest = true
	origBBSHome = ptttype.SetBBSHOME("./testcase")
}

func teardownTest() {
	ptttype.SetBBSHOME(origBBSHome)
	IsTest = false
}
