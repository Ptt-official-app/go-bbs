package ptt

import (
	"github.com/PichuChen/go-bbs/cmbbs"
	"github.com/PichuChen/go-bbs/ptttype"
)

func initCurrentUser(userID *[ptttype.IDLEN + 1]byte) (int, *ptttype.UserecBig5, error) {
	return cmbbs.PasswdLoadUser(userID)
}
