package bbs

import (
	"github.com/PichuChen/go-bbs"

	"bytes"
	"encoding/binary"
	"io"
	"log"
	"os"
	"time"
)

const (
	FORMOSABBS_STRLEN = 80

	PosOfFormosaBBSFileHeaderFilename = 0
	PosOfFormosaBBSFileHeaderOwner    = FORMOSABBS_STRLEN
	PosOfFormosaBBSFileHeaderPostno   = FORMOSABBS_STRLEN - 8
	PosOfFormosaBBSFileHeaderModified = PosOfFormosaBBSFileHeaderOwner + FORMOSABBS_STRLEN - 8
	PosOfFormosaBBSFileHeaderTitle    = PosOfFormosaBBSFileHeaderOwner + FORMOSABBS_STRLEN
)

type FileHeader struct {
	Filename  string
	Modified  time.Time
	Recommend int8   // Important Level
	Owner     string // uid[.]
	Date      string
	Title     string

	Money   int
	AnnoUid int
	// VoteLimits
	ReferRef  uint // 至底公告？
	ReferFlag bool // 至底公告？

	Filemode uint8

	Postno int32 // FormosaBBS only
}

func OpenFormosaBBSFileHeaderFile(filename string) ([]*FileHeader, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ret := []*FileHeader{}

	for {
		hdr := make([]byte, 248)
		_, err := file.Read(hdr)
		// log.Println(len, err)
		if err == io.EOF {
			break
		}

		f, err := NewFomosaBBSFileHeaderWithByte(hdr)
		if err != nil {
			return nil, err
		}
		ret = append(ret, f)
		// log.Println(f.Filename)

	}

	return ret, nil

}

func NewFomosaBBSFileHeaderWithByte(data []byte) (*FileHeader, error) {

	ret := FileHeader{}
	ret.Filename = string(bytes.Trim(data[PosOfFormosaBBSFileHeaderFilename:+PosOfFormosaBBSFileHeaderFilename+44], "\x00"))

	modifiedInt := binary.LittleEndian.Uint32(data[PosOfFormosaBBSFileHeaderModified : PosOfFormosaBBSFileHeaderModified+4])
	// log.Println("modifiedInt:", modifiedInt, PosOfFormosaBBSModified)
	ret.Modified = time.Unix(int64(modifiedInt), 0)

	// ret.Recommend = int8(data[PosOfRecommend])
	ret.Owner = string(bytes.Trim(data[PosOfFormosaBBSFileHeaderOwner:PosOfFormosaBBSFileHeaderOwner+72], "\x00"))
	// ret.Date = string(bytes.Trim(data[PosOfDate:PosOfDate+6], "\x00"))
	ret.Title = bbs.Big5ToUtf8(bytes.Trim(data[PosOfFormosaBBSFileHeaderTitle:PosOfFormosaBBSFileHeaderTitle+67], "\x00"))
	// // log.Println("PosOfUnionMulti:", PosOfUnionMulti, data[PosOfUnionMulti])

	// ret.Money = int(binary.LittleEndian.Uint32(data[PosOfUnionMulti : PosOfUnionMulti+4]))
	// ret.AnnoUid = int(binary.LittleEndian.Uint32(data[PosOfUnionMulti : PosOfUnionMulti+4]))

	// ret.Filemode = uint8(data[PosOfFilemode])

	// if ret.IsVotePost() {
	// 	ret.VoteLimits.Posts = data[PosOfUnionMulti+0]
	// 	ret.VoteLimits.Logins = data[PosOfUnionMulti+1]
	// 	ret.VoteLimits.Regtime = data[PosOfUnionMulti+2]
	// 	ret.VoteLimits.Badpost = data[PosOfUnionMulti+3]
	// }

	ret.Postno = int32(binary.LittleEndian.Uint32(data[PosOfFormosaBBSFileHeaderPostno : PosOfFormosaBBSFileHeaderPostno+4]))
	// ret.Title = binary.LittleEndian.Uint8(data[PTT_FNLEN+5+PTT_IDLEN+2+6 : PTT_FNLEN+5+PTT_IDLEN+2+6+PTT_TTLEN+1])

	return &ret, nil
}
