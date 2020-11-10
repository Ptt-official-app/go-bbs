package ptttype

import log "github.com/sirupsen/logrus"

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

func FixedBytesLen(bytes []byte) int {
	for idx, c := range bytes {
		if c == 0 {
			return idx
		}
	}
	return len(bytes)
}

func FixedBytesToString(bytes []byte) string {
	len := FixedBytesLen(bytes)
	return string(bytes[:len])
}
