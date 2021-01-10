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
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// LoginRecentRecord records Login record, please see https://github.com/ptt/pttbbs/blob/446c2bc34867286a2a093615ea69501f32c127e4/mbbsd/mbbsd.c#L1123
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
