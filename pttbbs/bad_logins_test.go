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
	"testing"
	"time"
)

// The format of bad login equal recent login, you can also see testfile of recent login
func TestOpenBadLoginFile(t *testing.T) {
	type testCase struct {
		filename string
		expected []*LoginAttempt
	}
	testCases := []*testCase{
		{
			filename: "testcase/bad_logins/logins.bad",
			expected: []*LoginAttempt{
				{
					Success:        true,
					UserID:         "SYSOP",
					LoginStartTime: time.Date(2021, 0o1, 0o1, 10, 8, 56, 0, time.UTC),
					FromHost:       "172.22.0.1",
				},
				{
					Success:        true,
					UserID:         "SYSOP",
					LoginStartTime: time.Date(2021, 0o1, 0o1, 10, 10, 50, 0, time.UTC),
					FromHost:       "172.22.0.1",
				},
				{
					Success:        true,
					UserID:         "abc123456789",
					LoginStartTime: time.Date(2021, 0o1, 0o1, 10, 11, 9, 0, time.UTC),
					FromHost:       "172.22.0.1",
				},
				{
					Success:        true,
					UserID:         "test01",
					LoginStartTime: time.Date(2021, 0o1, 0o1, 10, 11, 23, 0, time.UTC),
					FromHost:       "172.22.0.1",
				},
				{
					Success:        true,
					UserID:         "test02",
					LoginStartTime: time.Date(2021, 0o1, 0o1, 10, 11, 35, 0, time.UTC),
					FromHost:       "172.22.0.1",
				},

				{
					Success:        true,
					UserID:         "test03",
					LoginStartTime: time.Date(2021, 0o1, 0o1, 10, 11, 45, 0, time.UTC),
					FromHost:       "172.22.0.1",
				},
				{
					Success:        true,
					UserID:         "test04",
					LoginStartTime: time.Date(2021, 0o1, 0o1, 10, 13, 35, 0, time.UTC),
					FromHost:       "172.22.0.1",
				},
				{
					Success:        true,
					UserID:         "test05",
					LoginStartTime: time.Date(2021, 0o1, 0o1, 10, 13, 45, 0, time.UTC),
					FromHost:       "172.22.0.1",
				},
				{
					Success:        true,
					UserID:         "SYSOP",
					LoginStartTime: time.Date(2021, 0o1, 0o1, 10, 13, 53, 0, time.UTC),
					FromHost:       "172.22.0.1",
				},
				{
					Success:        true,
					UserID:         "test06",
					LoginStartTime: time.Date(2021, 0o1, 0o1, 10, 14, 38, 0, time.UTC),
					FromHost:       "172.22.0.1",
				},

				{
					Success:        true,
					UserID:         "SYSOP",
					LoginStartTime: time.Date(2021, 0o1, 0o1, 10, 14, 46, 0, time.UTC),
					FromHost:       "172.22.0.1",
				},
				{
					Success:        false,
					UserID:         "test01",
					LoginStartTime: time.Date(2021, 0o1, 0o1, 10, 15, 16, 0, time.UTC),
					FromHost:       "172.22.0.1",
				},
				{
					Success:        false,
					UserID:         "test02",
					LoginStartTime: time.Date(2021, 0o1, 0o1, 10, 15, 19, 0, time.UTC),
					FromHost:       "172.22.0.1",
				},
				{
					Success:        false,
					UserID:         "test03",
					LoginStartTime: time.Date(2021, 0o1, 0o1, 10, 15, 22, 0, time.UTC),
					FromHost:       "172.22.0.1",
				},
				{
					Success:        true,
					UserID:         "test04",
					LoginStartTime: time.Date(2021, 0o1, 0o1, 10, 15, 38, 0, time.UTC),
					FromHost:       "172.22.0.1",
				},
			},
		},
		{
			filename: "testcase/bad_logins/test01/logins.bad",
			expected: []*LoginAttempt{
				{
					Success:        false,
					UserID:         "",
					LoginStartTime: time.Date(2021, 0o1, 0o1, 10, 15, 16, 0, time.UTC),
					FromHost:       "172.22.0.1",
				},
			},
		},
	}

	for i, c := range testCases {
		actual, err := OpenBadLoginFile(c.filename)
		if err != nil {
			t.Errorf("Failed to open logins.bad. Err %v", err)
		}
		if len(c.expected) != len(actual) {
			t.Errorf("expceted result length (=%d) not match actual (=%d) on case %d", len(c.expected), len(actual), i)
		}

		for itemIndex, l := range actual {
			actualItem := l
			expectedItem := c.expected[itemIndex]
			if actualItem.FromHost == "" {
				t.Error("FromHost should never be empty")
			}
			if actualItem.LoginStartTime.IsZero() {
				t.Error("LoginStartTime should not be zero")
			}
			if actualItem.UserID == "" && actualItem.Success {
				t.Error("If UserID is empty, Success must be false")
			}

			// Testing for matching field
			if actualItem.Success != expectedItem.Success {
				t.Errorf("Success not match with index %d:%d, expected: %v, got: %v",
					i, itemIndex, expectedItem.Success, actualItem.Success)
			}
			if actualItem.UserID != expectedItem.UserID {
				t.Errorf("UserID not match with index %d:%d, expected: %v, got: %v",
					i, itemIndex, expectedItem.UserID, actualItem.UserID)
			}
			if actualItem.LoginStartTime.Sub(expectedItem.LoginStartTime) != 0 {
				t.Errorf("LoginStartTime not match with index %d:%d, expected: %v, got: %v",
					i, itemIndex, expectedItem.LoginStartTime, actualItem.LoginStartTime)
			}
			if actualItem.FromHost != expectedItem.FromHost {
				t.Errorf("FromHost not match with index %d:%d, expected: %v, got: %v",
					i, itemIndex, expectedItem.FromHost, actualItem.FromHost)
			}
		}
	}
}

func TestLoginAttempt(t *testing.T) {
	testLines := []string{
		" SYSOP       [01/01/2021 10:08:56 Fri] ?@172.22.0.1",
		"-test03      [01/12/2021 13:14:15 Tue] ?@1.2.3.4",
		" test03      [12/30/2021 21:55:59 Thu] ?@255.255.255.255",
		" abc123456789[01/01/2021 10:11:09 Fri] ?@127.0.0.1",
		"-abc123456789[01/01/2021 10:11:09 Fri] ?@192.168.1.1",
		"[01/01/2021 01:02:03 Fri] 1.2.3.4",
		"[01/12/2021 13:14:15 Tue] 255.255.255.255",
		"[12/30/2021 21:55:59 Thu] 100.100.100.100",
	}

	for _, line := range testLines {
		attempt := &LoginAttempt{}
		err := attempt.UnmarshalText([]byte(line))
		if err != nil {
			t.Errorf("Failed to unmarshal line %s. Err %v", line, err)
		}
		formatted, err := attempt.MarshalText()
		if err != nil {
			t.Errorf("Failed to marshal LoginAttempt. Err %v", err)
		}
		if string(formatted) != line {
			t.Errorf("Marshaled != original. Original: '%s'. Marshaled: '%s'.", line, string(formatted))
		}
	}
}
