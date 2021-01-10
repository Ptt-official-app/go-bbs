package bbs

import (
	"testing"
)

func TestParseFormosaBBSFileHeader(t *testing.T) {

	headers, err := OpenFormosaBBSFileHeaderFile("testcase/file/01.DIR")
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
		if header.ReferRef != expected[index].ReferRef {
			t.Errorf("ReferRef not match in index %d, expected: %q, got: %q", index, expected[index].ReferRef, header.ReferRef)
		}
		if header.Filemode != expected[index].Filemode {
			t.Errorf("Filemode not match in index %d, expected: %q, got: %q", index, expected[index].Filemode, header.Filemode)
		}

	}

}
