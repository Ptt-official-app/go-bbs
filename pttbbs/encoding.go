package pttbbs

import (
	"github.com/PichuChen/go-bbs"
)

func big5uaoToUTF8String(b []byte) string {
	return bbs.Big5ToUtf8(b)
}

func utf8ToBig5UAOString(s string) []byte {
	return bbs.Utf8ToBig5(s)
}

// newStringFormCString return a string from Null-Terminated String
func newStringFormCString(cs []byte) string {
	return bbs.CstrToString(cs)
}

func newStringFormBig5UAOCString(cs []byte) string {
	return bbs.Big5ToUtf8(bbs.CstrToBytes(cs))
}
