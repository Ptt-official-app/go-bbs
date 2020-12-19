package bbs

import (
	"encoding/binary"
	"io"
	"log"
	"os"
)

// https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h
type Userec struct {
	Version  uint32
	Userid   string
	Realname string
	Nickname string
	Passwd   string

	Uflag        uint32
	Userlevel    uint32
	Numlogindays uint32
	Numposts     uint32
	Firstlogin   uint32
	Lastlogin    uint32
	Lasthost     string
	// TODO
}

func OpenUserecFile(filename string) ([]*Userec, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ret := []*Userec{}

	for {
		user, eachErr := NewUserecWithFile(file)
		if eachErr != nil {
			// io.EOF is reading correctly to the end the file.
			if eachErr == io.EOF {
				break
			}

			err = eachErr
			break
		}
		ret = append(ret, user)
	}

	return ret, err

}

func NewUserecWithFile(file *os.File) (*Userec, error) {
	userecRaw := &UserecRaw{}

	err := binary.Read(file, binary.LittleEndian, userecRaw)
	if err != nil {
		return nil, err
	}

	user := NewUserecFromUserecRaw(userecRaw)

	return user, nil
}

func NewUserecFromUserecRaw(userecRaw *UserecRaw) *Userec {
	user := &Userec{}
	user.Version = userecRaw.Version
	user.Userid = CstrToString(userecRaw.UserID[:])
	user.Realname = Big5ToUtf8(CstrToBytes(userecRaw.RealName[:]))
	user.Nickname = Big5ToUtf8(CstrToBytes(userecRaw.Nickname[:]))
	user.Passwd = CstrToString(userecRaw.PasswdHash[:])

	user.Uflag = userecRaw.UFlag

	user.Userlevel = userecRaw.UserLevel
	user.Numlogindays = userecRaw.NumLoginDays
	user.Numposts = userecRaw.NumPosts
	user.Firstlogin = uint32(userecRaw.FirstLogin)
	user.Lastlogin = uint32(userecRaw.LastLogin)
	user.Lasthost = CstrToString(userecRaw.LastHost[:])

	return user
}
