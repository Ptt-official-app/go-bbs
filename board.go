package bbs

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"os"
	"time"
)

//
// For Current PTT
// Please see https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h
// boardheader_t
//

type BoardHeader struct {
	BrdName            string
	Title              string
	BM                 string
	Brdattr            uint32 // uid[.]
	ChessCountry       string
	VoteLimitPosts     uint8
	VoteLimitLogins    uint8
	BUpdate            time.Time
	PostLimitPosts     uint8
	PostLimitLogins    uint8
	BVote              uint8
	VTime              time.Time
	Level              uint32
	PermReload         time.Time
	Gid                int32
	Next               []int32
	FirstChild         []int32
	Parent             int32
	ChildCount         int32
	Nuser              int32
	PostExpire         int32
	EndGamble          time.Time
	PostType           string
	PostTypeF          string
	FastRecommendPause uint8
	VoteLimitBadPost   uint8
	PostLimitBadPost   uint8
	SRexpire           time.Time
}

const (
	PTT_BTLEN = 48
)

const (
	PosOfPTTBoardName       = 0
	PosOfPTTBoardTitle      = PosOfPTTBoardName + PTT_IDLEN + 1
	PosOfPTTBM              = PosOfPTTBoardTitle + PTT_BTLEN + 1
	PosOfBrdAttr            = 3 + PTT_IDLEN*3 + 3 + PosOfPTTBM
	PosOfChessCountry       = PosOfBrdAttr + 4
	PosOfVoteLimitPosts     = PosOfChessCountry + 1
	PosOfVoteLimitLogins    = PosOfVoteLimitPosts + 1
	PosOfBUpdate            = 1 + PosOfVoteLimitLogins + 1
	PosOfPostLimitPosts     = PosOfBUpdate + 4
	PosOfPostLimitLogins    = PosOfPostLimitPosts + 1
	PosOfBVote              = 1 + PosOfPostLimitLogins + 1
	PosOfVTime              = PosOfBVote + 1
	PosOfLevel              = PosOfVTime + 4
	PosOfPermReload         = PosOfLevel + 4
	PosOfGid                = PosOfPermReload + 4
	PosOfNext               = PosOfGid + 4
	PosOfFirstChild         = PosOfNext + 8
	PosOfParent             = PosOfFirstChild + 8
	PosOfChildCount         = PosOfParent + 4
	PosOfNuser              = PosOfChildCount + 4
	PosOfPostExpire         = PosOfNuser + 4
	PosOfEndGamble          = PosOfPostExpire + 4
	PosOfPostType           = PosOfEndGamble + 4
	PosOfPostTypeF          = PosOfPostType + 33
	PosOfFastRecommendPause = PosOfPostTypeF + 1
	PosOfVoteLimitBadPost   = PosOfFastRecommendPause + 1
	PosOfPostLimitBadPost   = PosOfVoteLimitBadPost + 1
	PosOfSRExpire           = 3 + PosOfPostLimitBadPost + 1
)

const (
	PTT_BRD_POSTMASK   = 0x00000020
	PTT_BRD_GROUPBOARD = 0x00000008
	PTT_PERM_SYSOP     = 000000040000
	PTT_PERM_BM        = 000000002000
	PTT_BRD_HIDE       = 0x00000010
)

func OpenBoardHeaderFile(filename string) ([]*BoardHeader, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ret := []*BoardHeader{}

	for {
		hdr := make([]byte, 256)
		_, err := file.Read(hdr)
		// log.Println(len, err)
		if err == io.EOF {
			break
		}

		f, err := NewBoardHeaderWithByte(hdr)
		if err != nil {
			return nil, err
		}
		ret = append(ret, f)
		// log.Println(f.Filename)

	}

	return ret, nil
}

func NewBoardHeaderWithByte(data []byte) (*BoardHeader, error) {
	ret := BoardHeader{}

	ret.BrdName = Big5ToUtf8((bytes.Trim(data[PosOfPTTBoardName:PosOfPTTBoardName+PTT_IDLEN+1], "\x00")))
	ret.Title = Big5ToUtf8((bytes.Split(data[PosOfPTTBoardTitle:PosOfPTTBoardTitle+PTT_BTLEN+1], []byte("\x00"))[0])) // Be careful about C-string end char \0
	ret.BM = string(bytes.Trim(data[PosOfPTTBM:PosOfPTTBM+PTT_IDLEN*3+3], "\x00"))
	ret.Brdattr = binary.LittleEndian.Uint32(data[PosOfBrdAttr : PosOfBrdAttr+4])
	ret.VoteLimitPosts = uint8(data[PosOfVoteLimitPosts])
	ret.VoteLimitLogins = uint8(data[PosOfVoteLimitLogins])
	ret.ChessCountry = string(bytes.Trim(data[PosOfChessCountry:PosOfChessCountry+1], "\x00"))
	bUpdateInt := binary.LittleEndian.Uint32(data[PosOfBUpdate : PosOfBUpdate+4])
	ret.BUpdate = time.Unix(int64(bUpdateInt), 0)
	ret.PostLimitPosts = uint8(data[PosOfPostLimitPosts])
	ret.PostLimitLogins = uint8(data[PosOfPostLimitLogins])
	ret.BVote = uint8(data[PosOfBVote])
	vTime := binary.LittleEndian.Uint32(data[PosOfVTime : PosOfVTime+4])
	ret.VTime = time.Unix(int64(vTime), 0)
	ret.Level = binary.LittleEndian.Uint32(data[PosOfLevel : PosOfLevel+4])
	permReload := binary.LittleEndian.Uint32(data[PosOfPermReload : PosOfPermReload+4])
	ret.PermReload = time.Unix(int64(permReload), 0)
	ret.Gid = int32(binary.LittleEndian.Uint32(data[PosOfGid : PosOfGid+4]))

	ret.Next = []int32{int32(binary.LittleEndian.Uint32(data[PosOfNext : PosOfNext+4])), int32(binary.LittleEndian.Uint32(data[PosOfNext+4 : PosOfNext+8]))}
	ret.FirstChild = []int32{int32(binary.LittleEndian.Uint32(data[PosOfFirstChild : PosOfFirstChild+4])), int32(binary.LittleEndian.Uint32(data[PosOfFirstChild+4 : PosOfFirstChild+8]))}

	ret.Parent = int32(binary.LittleEndian.Uint32(data[PosOfParent : PosOfParent+4]))
	ret.ChildCount = int32(binary.LittleEndian.Uint32(data[PosOfChildCount : PosOfChildCount+4]))

	ret.Nuser = int32(binary.LittleEndian.Uint32(data[PosOfNuser : PosOfNuser+4]))
	ret.PostExpire = int32(binary.LittleEndian.Uint32(data[PosOfPostExpire : PosOfPostExpire+4]))
	endGamble := binary.LittleEndian.Uint32(data[PosOfEndGamble : PosOfEndGamble+4])
	ret.EndGamble = time.Unix(int64(endGamble), 0)
	ret.PostType = Big5ToUtf8(bytes.Trim(data[PosOfPostType:PosOfPostType+33], "\x00"))
	ret.PostTypeF = Big5ToUtf8(bytes.Trim(data[PosOfPostTypeF:PosOfPostTypeF+1], "\x00"))

	ret.FastRecommendPause = uint8(data[PosOfFastRecommendPause])
	ret.VoteLimitBadPost = uint8(data[PosOfVoteLimitBadPost])
	ret.PostLimitBadPost = uint8(data[PosOfPostLimitBadPost])
	srExpire := binary.LittleEndian.Uint32(data[PosOfSRExpire : PosOfSRExpire+4])
	ret.SRexpire = time.Unix(int64(srExpire), 0)

	return &ret, nil
}
