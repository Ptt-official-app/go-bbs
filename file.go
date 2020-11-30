package bbs

//
// For Current PTT
// Please see https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h
// fileheader_t
//
// Refer data:
// [Re: [問題] FileHeader 的結構???](https://www.ptt.cc/bbs/PttCurrent/M.1219675989.A.F85.html)

import (
	"bytes"
	"encoding/binary"

	// "iconv"
	"io"
	"log"
	"os"
	"time"
)

const (
	PosOfPTTFilename  = 0
	PosOfPTTModified  = PosOfPTTFilename + PTT_FNLEN
	PosOfPTTRecommend = 1 + PosOfPTTModified + 4
	PosOfPTTOwner     = PosOfPTTRecommend + 1
	PosOfPTTDate      = PosOfPTTOwner + PTT_IDLEN + 2
	PosOfPTTTitle     = PosOfPTTDate + 6

	PosOfPTTUnionMulti = 1 + PosOfPTTTitle + PTT_TTLEN + 1
	PosOfPTTFilemode   = PosOfPTTUnionMulti + 4
)

type VoteLimits struct {
	Posts   uint8
	Logins  uint8
	Regtime uint8
	Badpost uint8
}

type FileHeader struct {
	Filename  string
	Modified  time.Time
	Recommend int8   // Important Level
	Owner     string // uid[.]
	Date      string
	Title     string

	Money   int
	AnnoUid int
	VoteLimits
	ReferRef  uint // 至底公告？
	ReferFlag bool // 至底公告？

	Filemode uint8

	Postno int32 // FormosaBBS only
}

func OpenFileHeaderFile(filename string) ([]*FileHeader, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ret := []*FileHeader{}

	for {
		hdr := make([]byte, 128)
		_, err := file.Read(hdr)
		// log.Println(len, err)
		if err == io.EOF {
			break
		}

		f, err := NewFileHeaderWithByte(hdr)
		if err != nil {
			return nil, err
		}
		ret = append(ret, f)
		// log.Println(f.Filename)

	}

	return ret, nil

}

func NewFileHeaderWithByte(data []byte) (*FileHeader, error) {

	ret := FileHeader{}
	ret.Filename = string(bytes.Trim(data[PosOfPTTFilename:+PosOfPTTFilename+PTT_FNLEN], "\x00"))

	modifiedInt := binary.LittleEndian.Uint32(data[PosOfPTTModified : PosOfPTTModified+4])
	ret.Modified = time.Unix(int64(modifiedInt), 0)

	ret.Recommend = int8(data[PosOfPTTRecommend])
	ret.Owner = string(bytes.Trim(data[PosOfPTTOwner:PosOfPTTOwner+PTT_IDLEN+2], "\x00"))
	ret.Date = string(bytes.Trim(data[PosOfPTTDate:PosOfPTTDate+6], "\x00"))
	ret.Title = Big5ToUtf8(bytes.Trim(data[PosOfPTTTitle:PosOfPTTTitle+PTT_TTLEN+1], "\x00"))
	// log.Println("PosOfUnionMulti:", PosOfUnionMulti, data[PosOfUnionMulti])

	ret.Money = int(binary.LittleEndian.Uint32(data[PosOfPTTUnionMulti : PosOfPTTUnionMulti+4]))
	ret.AnnoUid = int(binary.LittleEndian.Uint32(data[PosOfPTTUnionMulti : PosOfPTTUnionMulti+4]))

	ret.Filemode = uint8(data[PosOfPTTFilemode])

	if ret.IsVotePost() {
		ret.VoteLimits.Posts = data[PosOfPTTUnionMulti+0]
		ret.VoteLimits.Logins = data[PosOfPTTUnionMulti+1]
		ret.VoteLimits.Regtime = data[PosOfPTTUnionMulti+2]
		ret.VoteLimits.Badpost = data[PosOfPTTUnionMulti+3]
	}

	// ret.Title = binary.LittleEndian.Uint8(data[PTT_FNLEN+5+PTT_IDLEN+2+6 : PTT_FNLEN+5+PTT_IDLEN+2+6+PTT_TTLEN+1])

	return &ret, nil
}

func (h *FileHeader) MarshalToByte() ([]byte, error) {
	ret := make([]byte, 128)

	copy(ret[PosOfPTTFilename:PosOfPTTFilename+PTT_FNLEN], h.Filename)
	binary.LittleEndian.PutUint32(ret[PosOfPTTModified:PosOfPTTModified+4], uint32(h.Modified.Unix()))

	ret[PosOfPTTRecommend] = byte(h.Recommend)
	copy(ret[PosOfPTTOwner:PosOfPTTOwner+PTT_IDLEN+2], h.Owner)
	copy(ret[PosOfPTTDate:PosOfPTTDate+6], h.Date)
	copy(ret[PosOfPTTTitle:PosOfPTTTitle+PTT_TTLEN+1], Utf8ToBig5(h.Title))

	// TODO: Check file mode for set Money or AnnoUid ... etc

	return ret, nil
}

func (h *FileHeader) IsVotePost() bool {
	return h.Filemode&PTT_FILE_VOTE != 0
}
