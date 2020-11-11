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
	shm.IsTest = true
	shm.LoadUHash()
	shm.AttachSHM()
}

func teardownTest() {
	shm.IsTest = false
	ptttype.SetBBSHOME(origBBSHOME)
	isTest = false
}
