package bbs

import (
	// "log"
	"testing"
	"time"
)

func TestParseLoginRecent(t *testing.T) {

	actualRecords, err := OpenLoginRecentFile("testcase/logins_recent/01.logins.recent")
	if err != nil {
		t.Error(err)
	}
	expectedRecords := []LoginRecentRecord{
		{
			LoginStartTime: time.Date(2020, 10, 24, 05, 57, 18, 0, time.UTC),
			FromHost:       "103.246.218.43",
		},
		{
			LoginStartTime: time.Date(2020, 11, 04, 10, 50, 49, 0, time.UTC),
			FromHost:       "103.246.218.43",
		},
		{
			LoginStartTime: time.Date(2020, 11, 04, 11, 29, 35, 0, time.UTC),
			FromHost:       "103.246.218.43",
		},
	}
	for index, actual := range actualRecords {

		if len(expectedRecords) <= index {
			t.Errorf("len(expectedRecords) <= len(actualRecords)")
			return
		}
		expected := expectedRecords[index]
		if actual.LoginStartTime != expected.LoginStartTime {
			t.Errorf("Login start time not match in index %d, expected: %v, got: %v", index, expected.LoginStartTime, actual.LoginStartTime)

		}

		if actual.FromHost != expected.FromHost {
			t.Errorf("from host not match in index %d, expected: %v, got: %v", index, expected.FromHost, actual.FromHost)

		}

	}

}
