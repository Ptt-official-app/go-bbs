package bbs

import (
	"testing"
)

func TestOpenBadLoginFile(t *testing.T) {
	type testCase struct {
		filename string
		expected []*LoginAttempt
	}
	testCases := []*testCase{
		{
			filename: "testcase/bad_logins/logins.bad",
			expected: nil,
		},
		{
			filename: "testcase/bad_logins/test01/logins.bad",
			expected: nil,
		},
	}

	for _, c := range testCases {
		attemps, err := OpenBadLoginFile(c.filename)
		if err != nil {
			t.Errorf("Failed to open logins.bad. Err %v", err)
		}
		for _, l := range attemps {
			if l.FromHost == "" {
				t.Error("FromHost should never be empty")
			}
			if l.LoginStartTime.IsZero() {
				t.Error("LoginStartTime should not be zero")
			}
			if l.UserId == "" && l.Success {
				t.Error("If UserId is empty, Success must be false")
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
