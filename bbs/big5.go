package bbs

import (
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

func Utf8ToBig5(input string) []byte {
	utf8ToBig5 := traditionalchinese.Big5.NewEncoder()
	big5, _, _ := transform.Bytes(utf8ToBig5, []byte(input))
	return big5
}

func Big5ToUtf8(input []byte) string {
	big5ToUTF8 := traditionalchinese.Big5.NewDecoder()
	utf8, _, _ := transform.String(big5ToUTF8, string(input))
	return utf8
}
