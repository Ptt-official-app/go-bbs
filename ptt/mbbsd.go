package ptt

import (
	"github.com/PichuChen/go-bbs/cmbbs"
	"github.com/PichuChen/go-bbs/ptttype"
)

//LoginQuery
//Params
//	userID: userID
//	passwd: passwd
//	ip: ip
//	isHashed: whether the input-passwd is already hashed.
//
//Return
//	*UserecBig5: current user
//  error: err
func LoginQuery(userID *[ptttype.IDLEN + 1]byte, passwd []byte, ip [ptttype.IPV4LEN + 1]byte, isHashed bool) (*ptttype.UserecBig5, error) {
	if !cmbbs.IsValidUserID(userID) {
		return nil, ptttype.ErrInvalidUserID
	}

	_, cuser, err := initCurrentUser(userID)
	if err != nil {
		return nil, err
	}

	isValid, err := cmbbs.CheckPasswd(cuser.PasswdHash[:], passwd, isHashed)
	if err != nil || !isValid {
		cmbbs.LogAttempt(userID, ip, true)
		return nil, err
	}

	return cuser, nil
}
