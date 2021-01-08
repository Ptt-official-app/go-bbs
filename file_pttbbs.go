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

package bbs

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
	PosOfPttFileHeaderFilename  = 0
	PosOfPttFileHeaderModified  = PosOfPttFileHeaderFilename + PTT_FNLEN
	PosOfPttFileHeaderRecommend = 1 + PosOfPttFileHeaderModified + 4
	PosOfPttFileHeaderOwner     = PosOfPttFileHeaderRecommend + 1
	PosOfPttFileHeaderDate      = PosOfPttFileHeaderOwner + PTT_IDLEN + 2
	PosOfPttFileHeaderTitle     = PosOfPttFileHeaderDate + 6

	PosOfPttFileHeaderUnionMulti = 1 + PosOfPttFileHeaderTitle + PTT_TTLEN + 1
	PosOfPttFileHeaderFilemode   = PosOfPttFileHeaderUnionMulti + 4
)

type VoteLimits struct {
	Posts   uint8
	Logins  uint8
	Regtime uint8
	Badpost uint8
}

// FileHeader records article's metainfo
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

// OpenFileHeaderFile function open a .DIR file in board directory.
// It returns slice of FileHeader.
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
	ret.Filename = string(bytes.Trim(data[PosOfPttFileHeaderFilename:+PosOfPttFileHeaderFilename+PTT_FNLEN], "\x00"))

	modifiedInt := binary.LittleEndian.Uint32(data[PosOfPttFileHeaderModified : PosOfPttFileHeaderModified+4])
	ret.Modified = time.Unix(int64(modifiedInt), 0)

	ret.Recommend = int8(data[PosOfPttFileHeaderRecommend])
	ret.Owner = string(bytes.Trim(data[PosOfPttFileHeaderOwner:PosOfPttFileHeaderOwner+PTT_IDLEN+2], "\x00"))
	ret.Date = string(bytes.Trim(data[PosOfPttFileHeaderDate:PosOfPttFileHeaderDate+6], "\x00"))
	ret.Title = Big5ToUtf8(bytes.Trim(data[PosOfPttFileHeaderTitle:PosOfPttFileHeaderTitle+PTT_TTLEN+1], "\x00"))
	// log.Println("PosOfUnionMulti:", PosOfUnionMulti, data[PosOfUnionMulti])

	ret.Money = int(binary.LittleEndian.Uint32(data[PosOfPttFileHeaderUnionMulti : PosOfPttFileHeaderUnionMulti+4]))
	ret.AnnoUid = int(binary.LittleEndian.Uint32(data[PosOfPttFileHeaderUnionMulti : PosOfPttFileHeaderUnionMulti+4]))

	ret.Filemode = uint8(data[PosOfPttFileHeaderFilemode])

	if ret.IsVotePost() {
		ret.VoteLimits.Posts = data[PosOfPttFileHeaderUnionMulti+0]
		ret.VoteLimits.Logins = data[PosOfPttFileHeaderUnionMulti+1]
		ret.VoteLimits.Regtime = data[PosOfPttFileHeaderUnionMulti+2]
		ret.VoteLimits.Badpost = data[PosOfPttFileHeaderUnionMulti+3]
	}

	// ret.Title = binary.LittleEndian.Uint8(data[PTT_FNLEN+5+PTT_IDLEN+2+6 : PTT_FNLEN+5+PTT_IDLEN+2+6+PTT_TTLEN+1])

	return &ret, nil
}

func (h *FileHeader) MarshalToByte() ([]byte, error) {
	ret := make([]byte, 128)

	copy(ret[PosOfPttFileHeaderFilename:PosOfPttFileHeaderFilename+PTT_FNLEN], h.Filename)
	binary.LittleEndian.PutUint32(ret[PosOfPttFileHeaderModified:PosOfPttFileHeaderModified+4], uint32(h.Modified.Unix()))

	ret[PosOfPttFileHeaderRecommend] = byte(h.Recommend)
	copy(ret[PosOfPttFileHeaderOwner:PosOfPttFileHeaderOwner+PTT_IDLEN+2], h.Owner)
	copy(ret[PosOfPttFileHeaderDate:PosOfPttFileHeaderDate+6], h.Date)
	copy(ret[PosOfPttFileHeaderTitle:PosOfPttFileHeaderTitle+PTT_TTLEN+1], Utf8ToBig5(h.Title))

	// TODO: Check file mode for set Money or AnnoUid ... etc

	return ret, nil
}

func (h *FileHeader) IsVotePost() bool {
	return h.Filemode&PTT_FILE_VOTE != 0
}
