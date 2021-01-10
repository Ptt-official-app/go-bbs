package pttbbs

import (
	"github.com/PichuChen/go-bbs"
)

func big5uaoToUTF8String(b []bytes) string {
	return bbs.Big5ToUtf8(b)
}
