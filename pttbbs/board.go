// Copyright 2020 Pichu Chen, The PTT APP Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pttbbs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Ptt-official-app/go-bbs/filelock"
)

//
// For Current PTT
// Please see https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h
// boardheader_t
//

type BoardHeader struct {
	BrdName         string
	title           string
	bm              string
	Brdattr         uint32 // uid[.]
	ChessCountry    string
	VoteLimitPosts  uint8
	VoteLimitLogins uint8
	BUpdate         time.Time
	PostLimitPosts  uint8
	PostLimitLogins uint8
	BVote           uint8
	VTime           time.Time
	Level           uint32
	PermReload      time.Time

	// Parent class id, first item is start from 1.
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

func (b *BoardHeader) BoardID() string            { return b.BrdName }
func (b *BoardHeader) SetBoardID(newValue string) { b.BrdName = newValue }

func (b *BoardHeader) Title() string            { return b.title }
func (b *BoardHeader) SetTitle(newValue string) { b.title = newValue }

func (b *BoardHeader) IsClass() bool   { return b.IsGroupBoard() }
func (b *BoardHeader) ClassID() string { return fmt.Sprintf("%v", b.Gid) }

func (b *BoardHeader) IsNoCount() bool          { return b.Brdattr&0x00000002 != 0 }
func (b *BoardHeader) IsGroupBoard() bool       { return b.Brdattr&0x00000008 != 0 } // Class
func (b *BoardHeader) IsHide() bool             { return b.Brdattr&0x00000010 != 0 } // BoardHide board or friend only
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

func (b *BoardHeader) GetPostLimitPosts() uint8   { return b.PostLimitPosts }
func (b *BoardHeader) GetPostLimitLogins() uint8  { return b.PostLimitLogins }
func (b *BoardHeader) GetPostLimitBadPost() uint8 { return b.PostLimitBadPost }

func (b *BoardHeader) BM() []string {
	if b.bm == "" {
		return []string{}
	}
	return strings.Split(b.bm, "/")
}

const (
	// BoardTitleLength https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L165
	BoardTitleLength = 48
)

const (
	PosOfBoardName          = 0
	PosOfBoardTitle         = PosOfBoardName + IDLength + 1
	PosOfBM                 = PosOfBoardTitle + BoardTitleLength + 1
	PosOfBrdAttr            = 3 + PosOfBM + IDLength*3 + 3
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
	PosOfFirstChild         = PosOfNext + 4*2
	PosOfParent             = PosOfFirstChild + 4*2
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
	// BoardPostMask https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L211
	BoardPostMask = 0x00000020
	// BoardGroupBoard https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L209
	BoardGroupBoard = 0x00000008
	// PermSYSOP https://github.com/ptt/pttbbs/blob/master/include/perm.h#L22
	PermSYSOP = 000000040000
	// PermBM https://github.com/ptt/pttbbs/blob/master/include/perm.h#L18
	PermBM = 000000002000
	// BoardHide https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h#L210
	BoardHide = 0x00000010

	BoardHeaderRecordLength = 256
)

func NewBoardHeader() *BoardHeader {
	return &BoardHeader{}
}

func OpenBoardHeaderFile(filename string) ([]*BoardHeader, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ret := []*BoardHeader{}

	for {
		hdr := make([]byte, BoardHeaderRecordLength)
		_, err := file.Read(hdr)
		// log.Println(len, err)
		if err == io.EOF {
			break
		}

		f, err := UnmarshalBoardHeader(hdr)
		if err != nil {
			return nil, err
		}
		ret = append(ret, f)
		// log.Println(f.Filename)

	}

	return ret, nil
}

func AppendBoardHeaderFileRecord(filename string, newBoardHeader *BoardHeader) error {
	// If the file doesn't exist, create it, or append to the file

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	err = filelock.Lock(f)
	if err != nil {
		// File is lock
		return err
	}

	data, err := newBoardHeader.MarshalBinary()
	if err != nil {
		return err
	}

	if _, err := f.Write(data); err != nil {
		return err
	}

	filelock.Unlock(f)
	return nil
}

func RemoveBoardHeaderFileRecord(filename string, index int) error {

	fi, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("fi.OpenFile error: %v", err)
	}
	defer fi.Close()

	err = filelock.Lock(fi)
	if err != nil {
		// File is lock
		return err
	}

	fo, err := os.OpenFile(filename, os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("fo.OpenFile error: %v", err)
	}
	defer fo.Close()

	_, err = fi.Seek(int64((index+1)*BoardHeaderRecordLength), os.SEEK_SET)
	if err != nil {
		return fmt.Errorf("fi.Seek error: %v", err)
	}
	_, err = fo.Seek(int64((index)*BoardHeaderRecordLength), os.SEEK_SET)
	if err != nil {
		return fmt.Errorf("fo.Seek error: %v", err)
	}
	_, err = io.CopyBuffer(fo, fi, make([]byte, 256))
	if err != nil {
		return fmt.Errorf("copy error: %v", err)
	}
	log.Println("copy finish")

	size, err := fo.Seek(0, os.SEEK_CUR)
	if err != nil {
		return fmt.Errorf("fo.Seek for SEEK_CUR error: %v", err)
	}

	fo.Truncate(size)
	filelock.Unlock(fi)
	return nil

}

func UnmarshalBoardHeader(data []byte) (*BoardHeader, error) {
	ret := BoardHeader{}

	ret.BrdName = big5uaoToUTF8String(bytes.Split(data[PosOfBoardName:PosOfBoardName+IDLength+1], []byte("\x00"))[0])
	ret.title = big5uaoToUTF8String(bytes.Split(data[PosOfBoardTitle:PosOfBoardTitle+BoardTitleLength+1], []byte("\x00"))[0]) // Be careful about C-string end char \0
	ret.bm = string(bytes.Trim(data[PosOfBM:PosOfBM+IDLength*3+3], "\x00"))
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
	ret.PostType = big5uaoToUTF8String(bytes.Trim(data[PosOfPostType:PosOfPostType+33], "\x00"))
	ret.PostTypeF = big5uaoToUTF8String(bytes.Trim(data[PosOfPostTypeF:PosOfPostTypeF+1], "\x00"))

	ret.FastRecommendPause = uint8(data[PosOfFastRecommendPause])
	ret.VoteLimitBadPost = uint8(data[PosOfVoteLimitBadPost])
	ret.PostLimitBadPost = uint8(data[PosOfPostLimitBadPost])
	srExpire := binary.LittleEndian.Uint32(data[PosOfSRExpire : PosOfSRExpire+4])
	ret.SRexpire = time.Unix(int64(srExpire), 0)

	return &ret, nil
}

func (b *BoardHeader) MarshalBinary() ([]byte, error) {
	ret := make([]byte, BoardHeaderRecordLength)

	copy(ret[PosOfBoardName:PosOfBoardName+IDLength+1], utf8ToBig5UAOString(b.BrdName))
	copy(ret[PosOfBoardTitle:PosOfBoardTitle+BoardTitleLength+1], utf8ToBig5UAOString(b.title))
	copy(ret[PosOfBM:PosOfBM+IDLength*3+3], b.bm)
	binary.LittleEndian.PutUint32(ret[PosOfBrdAttr:PosOfBrdAttr+4], b.Brdattr)
	ret[PosOfVoteLimitPosts] = b.VoteLimitPosts
	ret[PosOfVoteLimitLogins] = b.VoteLimitLogins
	copy(ret[PosOfChessCountry:PosOfChessCountry+1], b.ChessCountry)
	binary.LittleEndian.PutUint32(ret[PosOfBUpdate:PosOfBUpdate+4], uint32(b.BUpdate.Unix()))
	ret[PosOfPostLimitPosts] = b.PostLimitPosts
	ret[PosOfPostLimitLogins] = b.PostLimitLogins
	ret[PosOfBVote] = b.BVote
	binary.LittleEndian.PutUint32(ret[PosOfVTime:PosOfVTime+4], uint32(b.VTime.Unix()))
	binary.LittleEndian.PutUint32(ret[PosOfLevel:PosOfLevel+4], b.Level)
	binary.LittleEndian.PutUint32(ret[PosOfPermReload:PosOfPermReload+4], uint32(b.PermReload.Unix()))
	binary.LittleEndian.PutUint32(ret[PosOfGid:PosOfGid+4], uint32(b.Gid))

	if len(b.Next) == 2 {
		binary.LittleEndian.PutUint32(ret[PosOfNext:PosOfNext+4], uint32(b.Next[0]))
		binary.LittleEndian.PutUint32(ret[PosOfNext+4:PosOfNext+8], uint32(b.Next[1]))
	}

	if len(b.FirstChild) == 2 {
		binary.LittleEndian.PutUint32(ret[PosOfFirstChild:PosOfFirstChild+4], uint32(b.FirstChild[0]))
		binary.LittleEndian.PutUint32(ret[PosOfFirstChild+4:PosOfFirstChild+8], uint32(b.FirstChild[1]))
	}

	binary.LittleEndian.PutUint32(ret[PosOfParent:PosOfParent+4], uint32(b.Parent))
	binary.LittleEndian.PutUint32(ret[PosOfChildCount:PosOfChildCount+4], uint32(b.ChildCount))

	binary.LittleEndian.PutUint32(ret[PosOfNuser:PosOfNuser+4], uint32(b.Nuser))
	binary.LittleEndian.PutUint32(ret[PosOfPostExpire:PosOfPostExpire+4], uint32(b.PostExpire))
	binary.LittleEndian.PutUint32(ret[PosOfEndGamble:PosOfEndGamble+4], uint32(b.EndGamble.Unix()))
	copy(ret[PosOfPostType:PosOfPostType+33], utf8ToBig5UAOString(b.PostType))
	copy(ret[PosOfPostTypeF:PosOfPostTypeF+1], utf8ToBig5UAOString(b.PostTypeF))

	ret[PosOfFastRecommendPause] = b.FastRecommendPause
	ret[PosOfVoteLimitBadPost] = b.VoteLimitBadPost
	ret[PosOfPostLimitBadPost] = b.PostLimitBadPost
	binary.LittleEndian.PutUint32(ret[PosOfSRExpire:PosOfSRExpire+4], uint32(b.SRexpire.Unix()))

	return ret, nil

}
