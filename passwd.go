package bbs

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"os"

	"github.com/PichuChen/go-bbs/ptttype"
)

var (
	PosOfPTTUserecVersionPos = 0
	PosOfPTTUserecUseridPos  = 4
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
			err = eachErr
			break
		}
		ret = append(ret, user)
	}
	if err == io.EOF {
		err = nil
	}

	return ret, err

}

func NewUserecWithFile(file *os.File) (*Userec, error) {
	userBig5 := &ptttype.UserecBig5{}

	err := binary.Read(file, binary.LittleEndian, userBig5)
	if err != nil {
		return nil, err
	}

	user := NewUserecFromBig5(userBig5)

	return user, nil
}

func NewUserecFromBig5(userBig5 *ptttype.UserecBig5) *Userec {
	user := &Userec{}
	user.Version = userBig5.Version
	user.Userid = ptttype.FixedBytesToString(userBig5.UserID[:])
	user.Realname = Big5ToUtf8(ptttype.FixedBytesToBytes(userBig5.RealName[:]))
	user.Nickname = Big5ToUtf8(ptttype.FixedBytesToBytes(userBig5.Nickname[:]))
	user.Passwd = ptttype.FixedBytesToString(userBig5.PasswdHash[:])
	user.Pad1 = userBig5.Pad1

	user.Uflag = userBig5.UFlag
	user._unused1 = userBig5.Unused1
	user.Userlevel = userBig5.UserLevel
	user.Numlogindays = userBig5.NumLoginDays
	user.Numposts = userBig5.NumPosts
	user.Firstlogin = uint32(userBig5.FirstLogin)
	user.Lastlogin = uint32(userBig5.LastLogin)
	user.Lasthost = ptttype.FixedBytesToString(userBig5.LastHost[:])

	return user
}

func NewUserecWithByte(data []byte) (*Userec, error) {

	ret := Userec{}

	ret.Version = binary.LittleEndian.Uint32(data[PosOfPTTUserecVersionPos : PosOfPTTUserecVersionPos+4])
	ret.Userid = string(bytes.Trim(data[PosOfPTTUserecUseridPos:PosOfPTTUserecUseridPos+PTT_IDLEN+1], "\x00"))

	// modifiedInt := binary.LittleEndian.Uint32(data[PosOfPTTModified : PosOfPTTModified+4])
	// ret.Modified = time.Unix(int64(modifiedInt), 0)

	// ret.Recommend = int8(data[PosOfPTTRecommend])
	// ret.Owner = string(bytes.Trim(data[PosOfPTTOwner:PosOfPTTOwner+PTT_IDLEN+2], "\x00"))
	// ret.Date = string(bytes.Trim(data[PosOfPTTDate:PosOfPTTDate+6], "\x00"))
	// ret.Title = Big5ToUtf8(string(bytes.Trim(data[PosOfPTTTitle:PosOfPTTTitle+PTT_TTLEN+1], "\x00")))
	// // log.Println("PosOfUnionMulti:", PosOfUnionMulti, data[PosOfUnionMulti])

	// ret.Money = int(binary.LittleEndian.Uint32(data[PosOfPTTUnionMulti : PosOfPTTUnionMulti+4]))
	// ret.AnnoUid = int(binary.LittleEndian.Uint32(data[PosOfPTTUnionMulti : PosOfPTTUnionMulti+4]))

	// ret.Filemode = uint8(data[PosOfPTTFilemode])

	// if ret.IsVotePost() {
	// 	ret.VoteLimits.Posts = data[PosOfPTTUnionMulti+0]
	// 	ret.VoteLimits.Logins = data[PosOfPTTUnionMulti+1]
	// 	ret.VoteLimits.Regtime = data[PosOfPTTUnionMulti+2]
	// 	ret.VoteLimits.Badpost = data[PosOfPTTUnionMulti+3]
	// }

	// ret.Title = binary.LittleEndian.Uint8(data[PTT_FNLEN+5+PTT_IDLEN+2+6 : PTT_FNLEN+5+PTT_IDLEN+2+6+PTT_TTLEN+1])

	return &ret, nil
}
