package bbs

//
// For Current PTT
// Please see https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h
// fileheader_t
//

import (
	"bytes"
	"io"
	"log"
	"os"
)

const (
	PTT_TTLEN = 64
	PTT_FNLEN = 28
)

type FileHeader struct {
	Filename string
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
		len, err := file.Read(hdr)
		log.Println(len, err)
		if err == io.EOF {
			break
		}

		f, err := NewFileHeaderWithByte(hdr)
		if err != nil {
			return nil, err
		}
		ret = append(ret, f)
		log.Println(f.Filename)

	}

	return ret, nil

}

func NewFileHeaderWithByte(data []byte) (*FileHeader, error) {
	ret := FileHeader{}
	ret.Filename = string(bytes.Trim(data[0:PTT_FNLEN], "\x00"))

	return &ret, nil
}
