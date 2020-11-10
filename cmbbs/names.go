package cmbbs

import "github.com/PichuChen/go-bbs/ptttype"

func IsValidUserID(userID *[ptttype.IDLEN + 1]byte) bool {
	if userID == nil {
		return false
	}

	len := ptttype.FixedBytesLen(userID[:])
	if len < 2 || len > ptttype.IDLEN {
		return false
	}

	if !isalpha(userID[0]) {
		return false
	}

	for idx, c := range userID {
		if idx == len {
			break
		}

		if !isalnum(c) {
			return false
		}
	}

	return true
}

func isalpha(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}

	if c >= 'a' && c <= 'z' {
		return true
	}

	return false
}

func isnumber(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isalnum(c byte) bool {
	return isalpha(c) || isnumber(c)
}
