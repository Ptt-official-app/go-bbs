package bbs

import (
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
