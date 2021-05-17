package bbs

import (
	"bytes"
	"regexp"
	"unsafe"
)

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var re = regexp.MustCompile(ansi)

func FilterANSI(dst, src []byte) []byte {
	out := re.ReplaceAllString(ToString(src), "")
	copy(dst[:], out)
	firstNull := bytes.IndexAny(dst, "\x00")
	return dst[:firstNull]
}

func FilterStringANSI(src string) string {
	return re.ReplaceAllString(src, "")
}

func ToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
