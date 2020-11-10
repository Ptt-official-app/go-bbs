package cmbbs

import (
	"encoding/binary"
	"os"
	"reflect"

	"github.com/PichuChen/go-bbs/crypt"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/shm"
	log "github.com/sirupsen/logrus"
)

func CheckPasswd(expected []byte, input []byte, isHashed bool) (bool, error) {
	if isHashed {
		return reflect.DeepEqual(expected, input), nil
	}
	pw, err := crypt.Fcrypt(input, expected)
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(pw, expected), nil
}

func LogAttempt(userID *[ptttype.IDLEN + 1]byte, ip [ptttype.IPV4LEN + 1]byte, isWithUserHome bool) {
}

func PasswdLoadUser(userID *[ptttype.IDLEN + 1]byte) (int, *ptttype.UserecBig5, error) {
	if userID == nil || userID[0] == 0 {
		return 0, nil, ptttype.ErrInvalidUserID
	}

	log.Debugf("PasswdLoadUser: to SearchUserBig5: userID: %v", userID)
	usernum, _, err := shm.SearchUserBig5(userID[:], false)
	log.Debugf("PasswdLoadUser: after SearchUserBig5: userID: %v usernum: %v err :%v", userID, usernum, err)
	if err != nil {
		return 0, nil, err
	}

	if usernum > ptttype.MAX_USERS {
		return 0, nil, ptttype.ErrInvalidUserID
	}

	user, err := PasswdQuery(usernum)
	if err != nil {
		return 0, nil, err
	}

	return usernum, user, nil
}

func PasswdQuery(num int) (*ptttype.UserecBig5, error) {
	if num < 1 || num > ptttype.MAX_USERS {
		return nil, ptttype.ErrInvalidUserID
	}

	file, err := os.Open(ptttype.FN_PASSWD)
	if err != nil {
		return nil, err
	}

	user := &ptttype.UserecBig5{}
	offset := ptttype.USERREBIG5SZ * (int64(num) - 1)
	_, err = file.Seek(offset, 0)
	if err != nil {
		return nil, err
	}
	err = binary.Read(file, binary.LittleEndian, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
