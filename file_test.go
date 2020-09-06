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
		if header.Filename != expected[index].Filename {
			t.Logf("lena :%d %d", len(header.Filename), len(expected[index].Filename))
			t.Errorf("Filename not match in index %d, expected: %s, got: %s", index, expected[index].Filename, header.Filename)
		}

		if header.Modified.Sub(expected[index].Modified) != 0 {
			t.Errorf("Modified not match in index %d, expected: %s, got: %s", index, expected[index].Modified, header.Modified)
		}
		if header.Recommend != expected[index].Recommend {
			t.Errorf("Recommend not match in index %d, expected: %q, got: %q", index, expected[index].Recommend, header.Recommend)
		}
		if header.Owner != expected[index].Owner {
			t.Errorf("Owner not match in index %d, expected: %s, got: %s", index, expected[index].Owner, header.Owner)
		}
		if header.Date != expected[index].Date {
			t.Logf("Date :%d %d", len(header.Date), len(expected[index].Date))
			t.Errorf("Date not match in index %d, expected: %q, got: %q", index, expected[index].Date, header.Date)
		}
		if header.Title != expected[index].Title {
			t.Errorf("Title not match in index %d, expected: %q, got: %q", index, expected[index].Title, header.Title)
		}
		if header.Money != expected[index].Money {
			t.Errorf("Money not match in index %d, expected: %q, got: %q", index, expected[index].Money, header.Money)
		}
		if header.AnnoUid != expected[index].AnnoUid {
			t.Errorf("AnnoUid not match in index %d, expected: %q, got: %q", index, expected[index].AnnoUid, header.AnnoUid)
		}
		if header.VoteLimits != expected[index].VoteLimits {
			t.Errorf("VoteLimits not match in index %d, expected: %q, got: %q", index, expected[index].VoteLimits, header.VoteLimits)
		}
		if header.ReferRef != expected[index].ReferRef {
			t.Errorf("ReferRef not match in index %d, expected: %q, got: %q", index, expected[index].ReferRef, header.ReferRef)
		}
		if header.Filemode != expected[index].Filemode {
			t.Errorf("Filemode not match in index %d, expected: %q, got: %q", index, expected[index].Filemode, header.Filemode)
		}

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

func TestParseFormosaBBSFileHeader(t *testing.T) {

	headers, err := OpenFormosaBBSFileHeaderFile("testcase/formosabbs_file/01.DIR")
	if err != nil {
		t.Error(err)
	}

	expected := []FileHeader{
		{
			Filename: "M.1444066232.A",
			Owner:    "programchen",
			Date:     "",
			Title:    "program",
			Postno:   3362,
		},
		{
			Filename: "M.1456585692.A",
			Owner:    "download",
			Date:     "",
			Title:    "中山大學學生套房出租",
			Postno:   3363,
		},
		{
			Filename: "M.1469020580.A",
			Owner:    "programchen",
			Date:     "",
			Title:    "are you busy ",
			Postno:   3364,
		},
		{
			Filename: "M.1476541525.A",
			Owner:    "programchen",
			Date:     "",
			Title:    "test",
			Postno:   3365,
		},
		{
			Filename: "M.1483330205.A",
			Owner:    "programchen",
			Date:     "",
			Title:    "TEST new",
			Postno:   3366,
		},
		{
			Filename: "M.1484320784.A",
			Owner:    "programchen",
			Date:     "",
			Title:    "rainy ",
			Postno:   3367,
		},
		{
			Filename: "M.1486650586.A",
			Owner:    "programchen",
			Date:     "",
			Title:    "cold night",
			Postno:   3368,
		},
		{
			Filename: "M.1489418653.A",
			Owner:    "programchen",
			Date:     "",
			Title:    "good ngiht",
			Postno:   3369,
		},
		{
			Filename: "M.1499780490.A",
			Owner:    "programchen",
			Date:     "",
			Title:    "good day",
			Postno:   3370,
		},
		{
			Filename: "M.1502982246.A",
			Owner:    "frugal",
			Date:     "",
			Title:    "test",
			Postno:   3371,
		},
		{
			Filename: "M.1504100036.A",
			Owner:    "programchen",
			Date:     "",
			Title:    "2133",
			Postno:   3372,
		},
		{
			Filename: "M.1509734390.A",
			Owner:    "programchen",
			Date:     "",
			Title:    "0239",
			Postno:   3373,
		},
		{
			Filename: "M.1522244563.A",
			Owner:    "programchen",
			Date:     "",
			Title:    "online",
			Postno:   3374,
		},
		{
			Filename: "M.1584333992.A",
			Owner:    "salman",
			Date:     "",
			Title:    "  許留",
			Postno:   3376,
		},
		{
			Filename: "M.1596381800.A",
			Owner:    "pichubaby",
			Date:     "",
			Title:    "Ｔｅｓｔ",
			Postno:   3377,
		},
	}

	for index, header := range headers {
		if header.Filename != expected[index].Filename {
			t.Logf("lena :%d %d", len(header.Filename), len(expected[index].Filename))
			t.Errorf("Filename not match in index %d, expected: %s, got: %s", index, expected[index].Filename, header.Filename)
		}

		if header.Postno != expected[index].Postno {
			t.Errorf("Postno not match in index %d, expected: %d, got: %d", index, expected[index].Postno, header.Postno)
		}
		if header.Recommend != expected[index].Recommend {
			t.Errorf("Recommend not match in index %d, expected: %q, got: %q", index, expected[index].Recommend, header.Recommend)
		}
		if header.Owner != expected[index].Owner {
			t.Errorf("Owner not match in index %d, expected: %s, got: %s", index, expected[index].Owner, header.Owner)
		}
		if header.Date != expected[index].Date {
			t.Logf("Date :%d %d", len(header.Date), len(expected[index].Date))
			t.Errorf("Date not match in index %d, expected: %q, got: %q", index, expected[index].Date, header.Date)
		}
		if header.Title != expected[index].Title {
			t.Errorf("Title not match in index %d, expected: %q, got: %q", index, expected[index].Title, header.Title)
		}
		if header.Money != expected[index].Money {
			t.Errorf("Money not match in index %d, expected: %q, got: %q", index, expected[index].Money, header.Money)
		}
		if header.AnnoUid != expected[index].AnnoUid {
			t.Errorf("AnnoUid not match in index %d, expected: %q, got: %q", index, expected[index].AnnoUid, header.AnnoUid)
		}
		if header.VoteLimits != expected[index].VoteLimits {
			t.Errorf("VoteLimits not match in index %d, expected: %q, got: %q", index, expected[index].VoteLimits, header.VoteLimits)
		}
		if header.ReferRef != expected[index].ReferRef {
			t.Errorf("ReferRef not match in index %d, expected: %q, got: %q", index, expected[index].ReferRef, header.ReferRef)
		}
		if header.Filemode != expected[index].Filemode {
			t.Errorf("Filemode not match in index %d, expected: %q, got: %q", index, expected[index].Filemode, header.Filemode)
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
