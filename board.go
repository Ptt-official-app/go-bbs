package bbs

//
// For Current PTT
// Please see https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h
// boardheader_t
//

import (
	"io"
	"log"
	"os"
	"strings"
)

type BoardHeader struct {
	Boardname       string
	Title           string
	BoardModerators []string
}

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
	// var posOfPTTBrdname = posOfPTTBrdname + PTT_IDLEN

	ret := BoardHeader{}
	ret.Boardname = string(CstrToBytes(data[posOfPTTBrdname : +posOfPTTBrdname+PTT_IDLEN+1]))
	ret.Title = Big5ToUtf8(CstrToBytes(data[posOfTitle : posOfTitle+PTT_BTLEN+1]))
	ret.BoardModerators = strings.Split(Big5ToUtf8(CstrToBytes(data[posOfBM:posOfBM+3*PTT_IDLEN+3])), "/")

	return &ret, nil
}
