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
	FILE_LOCAL     = 0x01 /* local saved,  non-mail */
	FILE_READ      = 0x01 /* already read, mail only */
	FILE_MARKED    = 0x02 /* non-mail + mail */
	FILE_DIGEST    = 0x04 /* digest,       non-mail */
	FILE_REPLIED   = 0x04 /* replied,      mail only */
	FILE_BOTTOM    = 0x08 /* push_bottom,  non-mail */
	FILE_MULTI     = 0x08 /* multi send,   mail only */
	FILE_SOLVED    = 0x10 /* problem solved, sysop/BM non-mail only */
	FILE_HIDE      = 0x20 /* hide,	in announce */
	FILE_BID       = 0x20 /* bid,		in non-announce */
	FILE_BM        = 0x40 /* BM only,	in announce */
	FILE_VOTE      = 0x40 /* for vote,	in non-announce */
	FILE_ANONYMOUS = 0x80 /* anonymous file */

)
const (
	PTT_IDLEN = 12
	PTT_TTLEN = 64
	PTT_FNLEN = 28
)

const (
	PosOfFilename  = 0
	PosOfModified  = PosOfFilename + PTT_FNLEN
	PosOfRecommend = 1 + PosOfModified + 4
	PosOfOwner     = PosOfRecommend + 1
	PosOfDate      = PosOfOwner + PTT_IDLEN + 2
	PosOfTitle     = PosOfDate + 6

	PosOfUnionMulti = 1 + PosOfTitle + PTT_TTLEN + 1
	PosOfFilemode   = PosOfUnionMulti + 4
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
	ret.Filename = string(bytes.Trim(data[PosOfFilename:+PosOfFilename+PTT_FNLEN], "\x00"))

	modifiedInt := binary.LittleEndian.Uint32(data[PosOfModified : PosOfModified+4])
	ret.Modified = time.Unix(int64(modifiedInt), 0)

	ret.Recommend = int8(data[PosOfRecommend])
	ret.Owner = string(bytes.Trim(data[PosOfOwner:PosOfOwner+PTT_IDLEN+2], "\x00"))
	ret.Date = string(bytes.Trim(data[PosOfDate:PosOfDate+6], "\x00"))
	ret.Title = Big5ToUtf8(string(bytes.Trim(data[PosOfTitle:PosOfTitle+PTT_TTLEN+1], "\x00")))
	// log.Println("PosOfUnionMulti:", PosOfUnionMulti, data[PosOfUnionMulti])

	ret.Money = int(binary.LittleEndian.Uint32(data[PosOfUnionMulti : PosOfUnionMulti+4]))
	ret.AnnoUid = int(binary.LittleEndian.Uint32(data[PosOfUnionMulti : PosOfUnionMulti+4]))

	ret.Filemode = uint8(data[PosOfFilemode])

	if ret.IsVotePost() {
		ret.VoteLimits.Posts = data[PosOfUnionMulti+0]
		ret.VoteLimits.Logins = data[PosOfUnionMulti+1]
		ret.VoteLimits.Regtime = data[PosOfUnionMulti+2]
		ret.VoteLimits.Badpost = data[PosOfUnionMulti+3]
	}

	// ret.Title = binary.LittleEndian.Uint8(data[PTT_FNLEN+5+PTT_IDLEN+2+6 : PTT_FNLEN+5+PTT_IDLEN+2+6+PTT_TTLEN+1])

	return &ret, nil
}

func (h *FileHeader) MarshalToByte() ([]byte, error) {
	ret := make([]byte, 128)

	copy(ret[PosOfFilename:PosOfFilename+PTT_FNLEN], h.Filename)
	binary.LittleEndian.PutUint32(ret[PosOfModified:PosOfModified+4], uint32(h.Modified.Unix()))

	ret[PosOfRecommend] = byte(h.Recommend)
	copy(ret[PosOfOwner:PosOfOwner+PTT_IDLEN+2], h.Owner)
	copy(ret[PosOfDate:PosOfDate+6], h.Date)
	copy(ret[PosOfTitle:PosOfTitle+PTT_TTLEN+1], Utf8ToBig5(h.Title))

	// TODO: Check file mode for set Money or AnnoUid ... etc

	return ret, nil
}

func (h *FileHeader) IsVotePost() bool {
	return h.Filemode&FILE_VOTE != 0
}
