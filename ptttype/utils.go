package ptttype

import log "github.com/sirupsen/logrus"

//SetBBSHOME
//
//This is to safely set BBSHOME
//
//Params
//      bbshome: new bbshome
//
//Return
//      string: original bbshome
func SetBBSHOME(bbshome string) string {
	origBBSHome := BBSHOME
	log.Debugf("SetBBSHOME: %v", bbshome)

	// config.go
	BBSHOME = bbshome

	//common.go
	FN_PASSWD = BBSHOME + FN_PASSWD_POSTFIX /* User records */

	return origBBSHome
}
