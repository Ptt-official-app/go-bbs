package bbs

import "bytes"

//FixedBytesToString
//
//Only the bytes until \0 when converting to string.
//See tests for more examples.
//
//Params
//	fixedBytes: bytes
//
//Return
//	string: string
func FixedBytesToString(fixedBytes []byte) string {
	theBytes := FixedBytesToBytes(fixedBytes)
	return string(theBytes)
}

//FixedBytesToBytes
//
//Only the bytes until \0.
//See tests for more examples.
//
//Params
//	fixedBytes: fixed-bytes
//
//Return
//	[]byte: bytes
func FixedBytesToBytes(fixedBytes []byte) []byte {
	len := bytes.IndexByte(fixedBytes, 0)
	if len == -1 {
		return fixedBytes
	}

	return fixedBytes[:len]
}
