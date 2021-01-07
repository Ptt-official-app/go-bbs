package bbs

import (
	"encoding/hex"
	"strings"
	"testing"
	"time"
)

func TestParseFileHeader(t *testing.T) {
	headers, err := OpenFileHeaderFile("testcase/file/01.DIR")
	if err != nil {
		t.Error(err)
	}

	expected := []FileHeader{
		{
			Filename:  "M.1599059246.A.CF6",
			Modified:  time.Date(2020, time.September, 02, 15, 31, 28, 0, time.UTC),
			Recommend: 0,
			Owner:     "SYSOP",
			Date:      " 9/02",
			Title:     "[閒聊] 自己的文章自己寫",

			Money:   0,
			AnnoUid: 0,
			VoteLimits: VoteLimits{
				Posts:   0,
				Logins:  0,
				Regtime: 0,
				Badpost: 0,
			},
			ReferRef:  0,
			ReferFlag: true,
			Filemode:  0,
		},
		{
			Filename:  "M.1599059415.A.FBA",
			Modified:  time.Date(2020, time.September, 02, 15, 31, 37, 0, time.UTC),
			Recommend: 0,
			Owner:     "SYSOP",
			Date:      " 9/02",
			Title:     "[討論] 賞大稻埕煙火遠離人潮！",

			Money:   0,
			AnnoUid: 0,
			VoteLimits: VoteLimits{
				Posts:   0,
				Logins:  0,
				Regtime: 0,
				Badpost: 0,
			},
			ReferRef:  0,
			ReferFlag: true,
			Filemode:  0,
		},
		{
			Filename:  "M.1599059496.A.2BE",
			Modified:  time.Date(2020, time.September, 02, 15, 31, 46, 0, time.UTC),
			Recommend: 0,
			Owner:     "SYSOP",
			Date:      " 9/02",
			Title:     "[公告] 何不？ 五大寬容",

			Money:   0,
			AnnoUid: 0,
			VoteLimits: VoteLimits{
				Posts:   0,
				Logins:  0,
				Regtime: 0,
				Badpost: 0,
			},
			ReferRef:  0,
			ReferFlag: true,
			Filemode:  0,
		},
	}

	for index, header := range headers {
		CheckPttFileHeader(t, index, header, &expected[index])
	}
}

func TestParseFileHeader02(t *testing.T) {
	headers, err := OpenFileHeaderFile("testcase/file/02.DIR")
	if err != nil {
		t.Error(err)
	}

	expected := []FileHeader{
		{
			Filename:  "M.1604489415.A.C31",
			Modified:  time.Date(2020, 11, 04, 11, 30, 14, 0, time.UTC),
			Recommend: 0,
			Owner:     "pichu",
			Date:      "11/04",
			Title:     "[問題] test",

			Money:   0,
			AnnoUid: 0,
			VoteLimits: VoteLimits{
				Posts:   0,
				Logins:  0,
				Regtime: 0,
				Badpost: 0,
			},
			ReferRef:  0,
			ReferFlag: true,
			Filemode:  0,
		},
	}

	for index, header := range headers {
		CheckPttFileHeader(t, index, header, &expected[index])
	}
}
func TestParsePttFileHeaderTreasures(t *testing.T) {
	headers, err := OpenFileHeaderFile("testcase/file/03.DIR")
	if err != nil {
		t.Error(err)
	}

	expected := []FileHeader{
		{
			Filename:  "D6D8",
			Modified:  time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			Recommend: 0,
			Owner:     "SYSOP",
			Date:      "12/20",
			Title:     "◆ Folder 1.1.1.1",

			Money:   0,
			AnnoUid: 0,
			VoteLimits: VoteLimits{
				Posts:   0,
				Logins:  0,
				Regtime: 0,
				Badpost: 0,
			},
			ReferRef:  0,
			ReferFlag: true,
			Filemode:  0,
		},
	}

	for index, header := range headers {
		CheckPttFileHeader(t, index, header, &expected[index])
	}
}

func CheckPttFileHeader(t *testing.T, index int, actual *FileHeader, expected *FileHeader) {

	if actual.Filename != expected.Filename {
		t.Logf("lena :%d %d", len(actual.Filename), len(expected.Filename))
		t.Errorf("Filename not match in index %d, expected: %s, got: %s", index, expected.Filename, actual.Filename)
	}

	if actual.Modified.Sub(expected.Modified) != 0 {
		t.Errorf("Modified not match in index %d, expected: %s, got: %s", index, expected.Modified, actual.Modified)
	}
	if actual.Recommend != expected.Recommend {
		t.Errorf("Recommend not match in index %d, expected: %q, got: %q", index, expected.Recommend, actual.Recommend)
	}
	if actual.Owner != expected.Owner {
		t.Errorf("Owner not match in index %d, expected: %s, got: %s", index, expected.Owner, actual.Owner)
	}
	if actual.Date != expected.Date {
		t.Logf("Date :%d %d", len(actual.Date), len(expected.Date))
		t.Errorf("Date not match in index %d, expected: %q, got: %q", index, expected.Date, actual.Date)
	}
	if actual.Title != expected.Title {
		t.Errorf("Title not match in index %d, expected: %q, got: %q", index, expected.Title, actual.Title)
	}
	if actual.Money != expected.Money {
		t.Errorf("Money not match in index %d, expected: %q, got: %q", index, expected.Money, actual.Money)
	}
	if actual.AnnoUid != expected.AnnoUid {
		t.Errorf("AnnoUid not match in index %d, expected: %q, got: %q", index, expected.AnnoUid, actual.AnnoUid)
	}
	if actual.VoteLimits != expected.VoteLimits {
		t.Errorf("VoteLimits not match in index %d, expected: %q, got: %q", index, expected.VoteLimits, actual.VoteLimits)
	}
	if actual.ReferRef != expected.ReferRef {
		t.Errorf("ReferRef not match in index %d, expected: %q, got: %q", index, expected.ReferRef, actual.ReferRef)
	}
	if actual.Filemode != expected.Filemode {
		t.Errorf("Filemode not match in index %d, expected: %q, got: %q", index, expected.Filemode, actual.Filemode)
	}
}

func TestEncodingFileHeader(t *testing.T) {
	type TestCase struct {
		Input    FileHeader
		Expected []byte
	}

	testcase := []TestCase{
		{
			Input: FileHeader{
				Filename:  "M.1599059246.A.CF6",
				Modified:  time.Date(2020, time.September, 02, 15, 31, 28, 0, time.UTC),
				Recommend: 0,
				Owner:     "SYSOP",
				Date:      " 9/02",
				Title:     "[閒聊] 自己的文章自己寫",

				Money:   0,
				AnnoUid: 0,
				VoteLimits: VoteLimits{
					Posts:   0,
					Logins:  0,
					Regtime: 0,
					Badpost: 0,
				},
				ReferRef:  0,
				ReferFlag: true,
				Filemode:  0,
			},
			Expected: hexToByte(`
			4d2e 3135 3939 3035 3932 3436 2e41 2e43
4636 0000 0000 0000 0000 0000 d0ba 4f5f
0000 5359 534f 5000 0000 0000 0000 0000
2039 2f30 3200 5bb6 a2b2 e15d 20a6 dba4
76aa baa4 e5b3 b9a6 dba4 76bc 6700 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
			`),
		},
		{
			Input: FileHeader{
				Filename:  "M.1599059415.A.FBA",
				Modified:  time.Date(2020, time.September, 02, 15, 31, 37, 0, time.UTC),
				Recommend: 0,
				Owner:     "SYSOP",
				Date:      " 9/02",
				Title:     "[討論] 賞大稻埕煙火遠離人潮！",

				Money:   0,
				AnnoUid: 0,
				VoteLimits: VoteLimits{
					Posts:   0,
					Logins:  0,
					Regtime: 0,
					Badpost: 0,
				},
				ReferRef:  0,
				ReferFlag: true,
				Filemode:  0,
			},
			Expected: hexToByte(`
4d2e 3135 3939 3035 3934 3135 2e41 2e46
4241 0000 0000 0000 0000 0000 d9ba 4f5f
0000 5359 534f 5000 0000 0000 0000 0000
2039 2f30 3200 5bb0 51bd d75d 20bd e0a4
6abd 5fd1 4cb7 cfa4 f5bb b7c2 f7a4 48bc
e9a1 4900 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
			`),
		},
		{
			Input: FileHeader{
				Filename:  "M.1599059496.A.2BE",
				Modified:  time.Date(2020, time.September, 02, 15, 31, 46, 0, time.UTC),
				Recommend: 0,
				Owner:     "SYSOP",
				Date:      " 9/02",
				Title:     "[公告] 何不？ 五大寬容",

				Money:   0,
				AnnoUid: 0,
				VoteLimits: VoteLimits{
					Posts:   0,
					Logins:  0,
					Regtime: 0,
					Badpost: 0,
				},
				ReferRef:  0,
				ReferFlag: true,
				Filemode:  0,
			},
			Expected: hexToByte(`
4d2e 3135 3939 3035 3934 3936 2e41 2e32
4245 0000 0000 0000 0000 0000 e2ba 4f5f
0000 5359 534f 5000 0000 0000 0000 0000
2039 2f30 3200 5ba4 bda7 695d 20a6 f3a4
a3a1 4820 a4ad a46a bc65 ae65 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
			`),
		},
	}

	for index, c := range testcase {
		b, err := c.Input.MarshalToByte()
		t.Logf("log: %q, %q", b, err)
		if hex.Dump(b) != hex.Dump(c.Expected) {
			t.Errorf("Expected byte not match in index %d, expected: \n%s\n, got: \n%s", index, hex.Dump(c.Expected), hex.Dump(b))
		}

	}

}

func hexToByte(input string) []byte {
	s := strings.ReplaceAll(input, " ", "")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\n", "")
	b, _ := hex.DecodeString(s)
	return b
}
