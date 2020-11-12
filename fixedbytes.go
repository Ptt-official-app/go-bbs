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
func FixedBytesLen(bytes []byte) int {
	for idx, c := range bytes {
		if c == 0 {
			return idx
		}
	}
	return len(bytes)
}

//FixedBytesToString
//
//Only the bytes until \0 when converting to string.
//See tests for more examples.
//
//Params
//	bytes: bytes
//
//Return
//	string: string
func FixedBytesToString(bytes []byte) string {
	len := FixedBytesLen(bytes)
	return string(bytes[:len])
}

//FixedBytesToBytes
//
//Only the bytes until \0.
//See tests for more examples.
//
//Params
//	bytes: fixed-bytes
//
//Return
//	[]byte: bytes
func FixedBytesToBytes(bytes []byte) []byte {
	len := FixedBytesLen(bytes)
	return bytes[:len]
}
