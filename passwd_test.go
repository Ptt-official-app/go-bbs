package bbs

import (
	"testing"
)

func TestParseUserRecordHeader(t *testing.T) {

	headers, err := OpenUserecFile("testcase/passwd/01.PASSWDS")
	if err != nil {
		t.Error(err)
	}

	expected := []Userec{
		{
			Version: 4194,
		},
		{},
	}

	for index, header := range headers {
		if len(expected) <= index {
			return
		}
		t.Logf("version: %d", header.Version)
		if header.Version != expected[index].Version {
			t.Logf("lena :%d %d", (header.Version), (expected[index].Version))
			t.Errorf("Version not match in index %d, expected: %d, got: %d", index, expected[index].Version, header.Version)
		}

		// if header.Modified.Sub(expected[index].Modified) != 0 {
		// 	t.Errorf("Modified not match in index %d, expected: %s, got: %s", index, expected[index].Modified, header.Modified)
		// }
		// if header.Recommend != expected[index].Recommend {
		// 	t.Errorf("Recommend not match in index %d, expected: %q, got: %q", index, expected[index].Recommend, header.Recommend)
		// }
		// if header.Owner != expected[index].Owner {
		// 	t.Errorf("Owner not match in index %d, expected: %s, got: %s", index, expected[index].Owner, header.Owner)
		// }
		// if header.Date != expected[index].Date {
		// 	t.Logf("Date :%d %d", len(header.Date), len(expected[index].Date))
		// 	t.Errorf("Date not match in index %d, expected: %q, got: %q", index, expected[index].Date, header.Date)
		// }
		// if header.Title != expected[index].Title {
		// 	t.Errorf("Title not match in index %d, expected: %q, got: %q", index, expected[index].Title, header.Title)
		// }
		// if header.Money != expected[index].Money {
		// 	t.Errorf("Money not match in index %d, expected: %q, got: %q", index, expected[index].Money, header.Money)
		// }
		// if header.AnnoUid != expected[index].AnnoUid {
		// 	t.Errorf("AnnoUid not match in index %d, expected: %q, got: %q", index, expected[index].AnnoUid, header.AnnoUid)
		// }
		// if header.VoteLimits != expected[index].VoteLimits {
		// 	t.Errorf("VoteLimits not match in index %d, expected: %q, got: %q", index, expected[index].VoteLimits, header.VoteLimits)
		// }
		// if header.ReferRef != expected[index].ReferRef {
		// 	t.Errorf("ReferRef not match in index %d, expected: %q, got: %q", index, expected[index].ReferRef, header.ReferRef)
		// }
		// if header.Filemode != expected[index].Filemode {
		// 	t.Errorf("Filemode not match in index %d, expected: %q, got: %q", index, expected[index].Filemode, header.Filemode)
		// }

	}

}
