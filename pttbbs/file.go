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

// File header is in the board directory, it indicates the article's metainfo
// such as filename, author, title... usually without content.
//
// For Current PTT
// Please see https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h
// fileheader_t
//
// Refer data:
// [Re: [問題] FileHeader 的結構???](https://www.ptt.cc/bbs/PttCurrent/M.1219675989.A.F85.html)
//
// The `PosOf...` variables will be unexported soon.

package pttbbs

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"os"
	"time"

	"github.com/Ptt-official-app/go-bbs/filelock"
)

const (
	PosOfFileHeaderFilename  = 0
	PosOfFileHeaderModified  = PosOfFileHeaderFilename + FileNameLength
	PosOfFileHeaderRecommend = 1 + PosOfFileHeaderModified + 4
	PosOfFileHeaderOwner     = PosOfFileHeaderRecommend + 1
	PosOfFileHeaderDate      = PosOfFileHeaderOwner + IDLength + 2
	PosOfFileHeaderTitle     = PosOfFileHeaderDate + 6

	PosOfFileHeaderUnionMulti = 1 + PosOfFileHeaderTitle + TitleLength + 1
	PosOfFileHeaderFilemode   = PosOfFileHeaderUnionMulti + 4
)

// VoteLimits shows the limitation of a vote post.
type VoteLimits struct {
	Posts   uint8
	Logins  uint8
	Regtime uint8
	Badpost uint8
}

// FileHeader records article's metainfo
type FileHeader struct {
	filename  string
	modified  time.Time
	recommend int8   // Important Level
	owner     string // uid[.]
	date      string
	title     string

	money   int
	AnnoUID int
	VoteLimits
	ReferRef  uint // 至底公告？
	ReferFlag bool // 至底公告？

	Filemode uint8
}

func (f *FileHeader) Filename() string            { return f.filename }
func (f *FileHeader) SetFilename(newValue string) { f.filename = newValue }

func (f *FileHeader) Modified() time.Time               { return f.modified }
func (f *FileHeader) SetModified(newModified time.Time) { f.modified = newModified }

func (f *FileHeader) Recommend() int           { return int(f.recommend) }
func (f *FileHeader) AddRecommend(update int8) { f.recommend = f.recommend + update }

func (f *FileHeader) Owner() string            { return f.owner }
func (f *FileHeader) SetOwner(newValue string) { f.owner = newValue }

func (f *FileHeader) Date() string            { return f.date }
func (f *FileHeader) SetDate(newValue string) { f.date = newValue }

func (f *FileHeader) Title() string            { return f.title }
func (f *FileHeader) SetTitle(newValue string) { f.title = newValue }

func (f *FileHeader) Money() int { return f.money }

func NewFileHeader() *FileHeader {
	return &FileHeader{}
}

// OpenFileHeaderFile function open a .DIR file in board directory.
// It returns slice of FileHeader.
func OpenFileHeaderFile(filename string) ([]*FileHeader, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer file.Close()

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
		// log.Println(f.filename)

	}

	return ret, nil

}

func AppendFileHeaderFileRecord(filename string, newFileHeader *FileHeader) error {

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	err = filelock.Lock(f)
	if err != nil {
		// File is locked
		return err
	}
	defer filelock.Unlock(f)

	data, err := newFileHeader.MarshalToByte()
	if err != nil {
		return err
	}

	if _, err := f.Write(data); err != nil {
		return err
	}

	// TODO: update BoardHeader ?
	// https://github.com/ptt/pttbbs/blob/4d56e77f264960e43e060b77e442e166e5706417/mbbsd/syspost.c#L35

	return nil
}

func NewFileHeaderWithByte(data []byte) (*FileHeader, error) {

	ret := FileHeader{}
	ret.filename = string(bytes.Trim(data[PosOfFileHeaderFilename:+PosOfFileHeaderFilename+FileNameLength], "\x00"))

	modifiedInt := binary.LittleEndian.Uint32(data[PosOfFileHeaderModified : PosOfFileHeaderModified+4])
	ret.modified = time.Unix(int64(modifiedInt), 0)

	ret.recommend = int8(data[PosOfFileHeaderRecommend])
	ret.owner = string(bytes.Trim(data[PosOfFileHeaderOwner:PosOfFileHeaderOwner+IDLength+2], "\x00"))
	ret.date = string(bytes.Trim(data[PosOfFileHeaderDate:PosOfFileHeaderDate+6], "\x00"))
	ret.title = big5uaoToUTF8String(bytes.Trim(data[PosOfFileHeaderTitle:PosOfFileHeaderTitle+TitleLength+1], "\x00"))
	// log.Println("PosOfUnionMulti:", PosOfUnionMulti, data[PosOfUnionMulti])

	ret.money = int(binary.LittleEndian.Uint32(data[PosOfFileHeaderUnionMulti : PosOfFileHeaderUnionMulti+4]))
	ret.AnnoUID = int(binary.LittleEndian.Uint32(data[PosOfFileHeaderUnionMulti : PosOfFileHeaderUnionMulti+4]))

	ret.Filemode = uint8(data[PosOfFileHeaderFilemode])

	if ret.IsVotePost() {
		ret.VoteLimits.Posts = data[PosOfFileHeaderUnionMulti+0]
		ret.VoteLimits.Logins = data[PosOfFileHeaderUnionMulti+1]
		ret.VoteLimits.Regtime = data[PosOfFileHeaderUnionMulti+2]
		ret.VoteLimits.Badpost = data[PosOfFileHeaderUnionMulti+3]
	}

	// ret.title = binary.LittleEndian.Uint8(data[FileNameLength+5+IDLength+2+6 : FileNameLength+5+IDLength+2+6+TitleLength+1])

	return &ret, nil
}

func (f *FileHeader) MarshalToByte() ([]byte, error) {
	ret := make([]byte, 128)

	copy(ret[PosOfFileHeaderFilename:PosOfFileHeaderFilename+FileNameLength], f.filename)
	binary.LittleEndian.PutUint32(ret[PosOfFileHeaderModified:PosOfFileHeaderModified+4], uint32(f.modified.Unix()))

	ret[PosOfFileHeaderRecommend] = byte(f.recommend)
	copy(ret[PosOfFileHeaderOwner:PosOfFileHeaderOwner+IDLength+2], f.owner)
	copy(ret[PosOfFileHeaderDate:PosOfFileHeaderDate+6], f.date)
	copy(ret[PosOfFileHeaderTitle:PosOfFileHeaderTitle+TitleLength+1], utf8ToBig5UAOString(f.title))

	// TODO: Check file mode for set Money or AnnoUID ... etc

	return ret, nil
}

func (f *FileHeader) IsVotePost() bool {
	return f.Filemode&FileVote != 0
}
