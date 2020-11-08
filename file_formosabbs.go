package bbs

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"os"
	"time"
)

const (
	FORMOSABBS_STRLEN = 80

	PosOfFormosaBBSFilename = 0
	PosOfFormosaBBSOwner    = FORMOSABBS_STRLEN
	PosOfFormosaBBSPostno   = FORMOSABBS_STRLEN - 8
	PosOfFormosaBBSModified = PosOfFormosaBBSOwner + FORMOSABBS_STRLEN - 8
	PosOfFormosaBBSTitle    = PosOfFormosaBBSOwner + FORMOSABBS_STRLEN
)

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
	ret.Filename = string(bytes.Trim(data[PosOfFormosaBBSFilename:+PosOfFormosaBBSFilename+44], "\x00"))

	modifiedInt := binary.LittleEndian.Uint32(data[PosOfFormosaBBSModified : PosOfFormosaBBSModified+4])
	// log.Println("modifiedInt:", modifiedInt, PosOfFormosaBBSModified)
	ret.Modified = time.Unix(int64(modifiedInt), 0)

	// ret.Recommend = int8(data[PosOfRecommend])
	ret.Owner = string(bytes.Trim(data[PosOfFormosaBBSOwner:PosOfFormosaBBSOwner+72], "\x00"))
	// ret.Date = string(bytes.Trim(data[PosOfDate:PosOfDate+6], "\x00"))
	ret.Title = Big5ToUtf8(bytes.Trim(data[PosOfFormosaBBSTitle:PosOfFormosaBBSTitle+67], "\x00"))
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

	ret.Postno = int32(binary.LittleEndian.Uint32(data[PosOfFormosaBBSPostno : PosOfFormosaBBSPostno+4]))
	// ret.Title = binary.LittleEndian.Uint8(data[PTT_FNLEN+5+PTT_IDLEN+2+6 : PTT_FNLEN+5+PTT_IDLEN+2+6+PTT_TTLEN+1])

	return &ret, nil
}
