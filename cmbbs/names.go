package cmbbs

import (
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

func IsValidUserID(userID *[ptttype.IDLEN + 1]byte) bool {

	if userID == nil {
		return false
	}

	theLen := types.CstrLen(userID[:])
	if theLen < 2 || theLen > ptttype.IDLEN {
		return false
	}

	if !isalpha(userID[0]) {
		return false
	}

	for idx, c := range userID {
		if idx == theLen {
			break
		}

		if !isalnum(c) {
			return false
		}
	}

	return true
}

func isalpha(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z')
}

func isalnum(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= '0' && c <= '9') ||
		(c >= 'A' && c <= 'Z')
}
