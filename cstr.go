package bbs

import "bytes"

//Cstr
//
//[]byte with C String property in that \0 is considered as the end of the bytes/string.
//It is used to convert from fixed-length bytes to string or []byte with no \0.
//
//Naming Cstr instead of CString is to avoid confusion with C.CString
//(C.CString is from string, and should be compatible with string, not with []byte)
//(We also have str(len/cpy/cmp) functions in C)
//
//See tests for more examples of how to use fixed-bytes with Cstr to get no-\0 string / []byte
type Cstr []byte

//CstrToString
//
//Only the bytes until \0 when converting to string.
//See tests for more examples.
//
//Params
//	cstr
//
//Return
//	string: string
func CstrToString(cstr Cstr) string {
	theBytes := CstrToBytes(cstr)
	return string(theBytes)
}

//CstrToBytes
//
//Only the bytes until \0.
//See tests for more examples.
//
//Params
//	cstr
//
//Return
//	[]byte: bytes
func CstrToBytes(cstr Cstr) []byte {
	len := bytes.IndexByte(cstr, 0x00)
	if len == -1 {
		return cstr
	}

	return cstr[:len]
}
