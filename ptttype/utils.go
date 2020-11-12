package ptttype

import log "github.com/sirupsen/logrus"

//SetBBSHOME
//
//This is to safely set BBSHOME
//
//Params
//	bbshome: new bbshome
//
//Return
//	string: original bbshome
func SetBBSHOME(bbshome string) string {
	origBBSHome := BBSHOME
	log.Debugf("SetBBSHOME: %v", bbshome)

	// config.go
	BBSHOME = bbshome
	BBSPROG = BBSHOME + BBSPROGPOSTFIX

	//common.go
	FN_CONF_BANIP = BBSHOME + FN_CONF_BANIP_POSTFIX // 禁止連線的 IP 列表
	FN_PASSWD = BBSHOME + FN_PASSWD_POSTFIX         /* User records */

	return origBBSHome
}

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
