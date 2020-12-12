package bbs

//
// For Current PTT
// Please see https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h
// boardheader_t
//

import (
	"encoding/binary"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type BoardHeader struct {
	Boardname       string
	Title           string
	BoardModerators []string
	Brdattr         uint32 // raw value
	Chesscountry    uint8
	VoteLimitLogins uint8
	BoardUpdate     time.Time
	PostLimitLogins uint8
	BoardVote       uint8
	VoteTime        time.Time // Vote close time
	Level           uint32
	PermReload      time.Time
	Gid             int32
	// Next [2]int32
	// FirstChild [2]int32
	Parent              int32
	ChildCount          int32
	Nuser               int32 // How many user in this board
	PostExpire          int32 // ?
	EndGamble           time.Time
	posttype            [33]uint8
	postTypeF           uint8
	FastRecommendPuause int // Delay for continuous recommend (Push)
	VoteLimitBadpost    uint8
	PostLimitBadpost    uint8
	SRExpire            time.Time
}

func (b *BoardHeader) IsNoCount() bool          { return b.Brdattr&0x00000002 != 0 }
func (b *BoardHeader) IsGroudBoard() bool       { return b.Brdattr&0x00000008 != 0 } // Class
func (b *BoardHeader) IsHide() bool             { return b.Brdattr&0x00000010 != 0 } // Hide board or friend only
func (b *BoardHeader) IsPostMask() bool         { return b.Brdattr&0x00000020 != 0 } // Has Post or Reading Limition
func (b *BoardHeader) IsAnonymous() bool        { return b.Brdattr&0x00000040 != 0 }
func (b *BoardHeader) IsDefaultAnonymous() bool { return b.Brdattr&0x00000080 != 0 }
func (b *BoardHeader) IsNoCredit() bool         { return b.Brdattr&0x00000100 != 0 }
func (b *BoardHeader) IsVoteBoard() bool        { return b.Brdattr&0x00000200 != 0 }
func (b *BoardHeader) IsWarnEL() bool           { return b.Brdattr&0x00000400 != 0 } // Warning for Remove Board
func (b *BoardHeader) IsTop() bool              { return b.Brdattr&0x00000800 != 0 }
func (b *BoardHeader) IsNoRecommend() bool      { return b.Brdattr&0x00001000 != 0 } // Forbiddent Recommend (Push)
func (b *BoardHeader) IsAngelAnonymous() bool   { return b.Brdattr&0x00002000 != 0 }
func (b *BoardHeader) IsBMCount() bool          { return b.Brdattr&0x00004000 != 0 }
func (b *BoardHeader) IsIsSymbolic() bool       { return b.Brdattr&0x00008000 != 0 } // symbolic link to board
func (b *BoardHeader) IsNoBoo() bool            { return b.Brdattr&0x00010000 != 0 }
func (b *BoardHeader) IsRestrictedPost() bool   { return b.Brdattr&0x00040000 != 0 } // Board Friend only
func (b *BoardHeader) IsGuestPost() bool        { return b.Brdattr&0x00080000 != 0 }
func (b *BoardHeader) IsCooldown() bool         { return b.Brdattr&0x00100000 != 0 }
func (b *BoardHeader) IsCPLog() bool            { return b.Brdattr&0x00200000 != 0 }
func (b *BoardHeader) IsNoFastRecommend() bool  { return b.Brdattr&0x00400000 != 0 }
func (b *BoardHeader) IsIPLogRecommend() bool   { return b.Brdattr&0x00800000 != 0 }
func (b *BoardHeader) IsOver18() bool           { return b.Brdattr&0x01000000 != 0 }
func (b *BoardHeader) IsNoReply() bool          { return b.Brdattr&0x02000000 != 0 }
func (b *BoardHeader) IsAlignedComment() bool   { return b.Brdattr&0x04000000 != 0 }
func (b *BoardHeader) IsNoSelfDeletePost() bool { return b.Brdattr&0x08000000 != 0 }
func (b *BoardHeader) IsBMMaskContent() bool    { return b.Brdattr&0x10000000 != 0 }

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

	const PTT_BTLEN = 48

	var posOfPTTBrdname = 0
	var posOfTitle = posOfPTTBrdname + PTT_IDLEN + 1
	var posOfBM = posOfTitle + PTT_BTLEN + 1
	var posOfBrdattr = posOfBM + 3*PTT_IDLEN + 3 + 3
	// var posOfPTTBrdname = posOfPTTBrdname + PTT_IDLEN

	ret := BoardHeader{}
	ret.Boardname = string(CstrToBytes(data[posOfPTTBrdname : +posOfPTTBrdname+PTT_IDLEN+1]))
	ret.Title = Big5ToUtf8(CstrToBytes(data[posOfTitle : posOfTitle+PTT_BTLEN+1]))
	ret.BoardModerators = strings.Split(Big5ToUtf8(CstrToBytes(data[posOfBM:posOfBM+3*PTT_IDLEN+3])), "/")
	ret.Brdattr = binary.LittleEndian.Uint32(data[posOfBrdattr : posOfBrdattr+4])

	return &ret, nil
}
