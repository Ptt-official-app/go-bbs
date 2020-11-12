package bbs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type LoginRecentRecord struct {
	LoginStartTime time.Time
	FromHost       string
}

func OpenLoginRecentFile(filename string) ([]*LoginRecentRecord, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ret := []*LoginRecentRecord{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		f, err := NewLoginRecentRecord(line)
		if err != nil {
			return nil, err
		}
		ret = append(ret, f)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return ret, nil

}

func NewLoginRecentRecord(line string) (*LoginRecentRecord, error) {
	seg := strings.Split(line, " ")
	if len(seg) < 3 {
		return nil, fmt.Errorf("format for login recent incorrect")
	}
	t, err := time.Parse("01/02/2006 15:04:05", seg[0]+" "+seg[1])
	if err != nil {
		return nil, err
	}
	return &LoginRecentRecord{
		LoginStartTime: t,
		FromHost:       seg[2],
	}, nil
}
