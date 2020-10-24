package bbs

import (
	"encoding/binary"
	"io"
	"log"
	"os"
)

var (
	PosOfPTTUserecPos = 0
)

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
		hdr := make([]byte, 128)
		_, err := file.Read(hdr)
		// log.Println(len, err)
		if err == io.EOF {
			break
		}

		f, err := NewUserecWithByte(hdr)
		if err != nil {
			return nil, err
		}
		ret = append(ret, f)
		// log.Println(f.Filename)

	}

	return ret, nil

}

func NewUserecWithByte(data []byte) (*Userec, error) {

	ret := Userec{}

	ret.Version = binary.LittleEndian.Uint32(data[PosOfPTTUserecPos : PosOfPTTUserecPos+4])

	// ret.Filename = string(bytes.Trim(data[PosOfPTTFilename:+PosOfPTTFilename+PTT_FNLEN], "\x00"))

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
