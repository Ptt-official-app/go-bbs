package cmbbs

import (
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/shm"
)

var (
	isTest      = false
	origBBSHOME = ""
)

func setupTest() {
	isTest = true
	origBBSHOME = ptttype.SetBBSHOME("./testcase")
	shm.AttachSHM()
}

func teardownTest() {
	isTest = false
	ptttype.SetBBSHOME(origBBSHOME)
}
