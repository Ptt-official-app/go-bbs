package bbs

import (
	"testing"
)

func TestGetPasswdsPath(t *testing.T) {

	type TestCase struct {
		input    string
		expected string
	}
	cases := []TestCase{

		{
			input:    "/root",
			expected: "/root/.PASSWDS",
		},
	}

	for i, c := range cases {
		actual, err := GetPasswdsPath(c.input)
		if err != nil {
			t.Errorf("GetPasswdsPath err != nil on index %d", i)
		}
		if actual != c.expected {
			t.Errorf("GetPasswdsPath result not match on index %d with input:%v , expected: %v, got: %v",
				i, c.input, c.expected, actual)
		}
	}

}
