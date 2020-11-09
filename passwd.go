package bbs

import (
	"io"
	"log"
	"os"

	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

// https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h
type Userec struct {
	Version  uint32
	Userid   string
	Realname string
	Nickname string
	Passwd   string
	Pad1     uint8

	Uflag        uint32
	_unused1     uint32
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
	userecRaw, err := ptttype.NewUserecRawWithFile(file)
	if err != nil {
		return nil, err
	}

	user := NewUserecFromRaw(userecRaw)

	return user, nil
}

func NewUserecFromRaw(userecRaw *ptttype.UserecRaw) *Userec {
	user := &Userec{}
	user.Version = userecRaw.Version
	user.Userid = types.CstrToString(userecRaw.UserID[:])
	user.Realname = Big5ToUtf8(types.CstrToBytes(userecRaw.RealName[:]))
	user.Nickname = Big5ToUtf8(types.CstrToBytes(userecRaw.Nickname[:]))
	user.Passwd = types.CstrToString(userecRaw.PasswdHash[:])
	user.Pad1 = userecRaw.Pad1

	user.Uflag = userecRaw.UFlag
	user._unused1 = userecRaw.Unused1
	user.Userlevel = userecRaw.UserLevel
	user.Numlogindays = userecRaw.NumLoginDays
	user.Numposts = userecRaw.NumPosts
	user.Firstlogin = uint32(userecRaw.FirstLogin)
	user.Lastlogin = uint32(userecRaw.LastLogin)
	user.Lasthost = types.CstrToString(userecRaw.LastHost[:])

	return user
}
