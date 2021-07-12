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
			LoginStartTime: time.Date(2020, 10, 24, 0o5, 57, 18, 0, time.UTC),
			FromHost:       "103.246.218.43",
		},
		{
			LoginStartTime: time.Date(2020, 11, 0o4, 10, 50, 49, 0, time.UTC),
			FromHost:       "103.246.218.43",
		},
		{
			LoginStartTime: time.Date(2020, 11, 0o4, 11, 29, 35, 0, time.UTC),
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
