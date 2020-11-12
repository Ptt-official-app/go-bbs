package bbs

//FixedBytesLen
//
//Same effect as strlen (length until \0)
//See tests for more examples.
//
//Params
//	bytes: bytes
//
//Return
//  int: length of the fixed-bytes
func FixedBytesLen(fixedBytes []byte) int {
	for idx, c := range fixedBytes {
		if c == 0 {
			return idx
		}
	}
	return len(fixedBytes)
}

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
	len := FixedBytesLen(fixedBytes)
	return string(fixedBytes[:len])
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
	len := FixedBytesLen(fixedBytes)
	return fixedBytes[:len]
}
