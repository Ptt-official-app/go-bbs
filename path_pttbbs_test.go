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
		{
			input:    "/root/",
			expected: "/root//.PASSWDS",
		},
		{
			input:    "/root//",
			expected: "/root///.PASSWDS",
		},
		{
			input:    "home/bbs",
			expected: "home/bbs/.PASSWDS",
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

func TestGetBoardPath(t *testing.T) {

	type TestCase struct {
		input    string
		expected string
	}
	cases := []TestCase{

		{
			input:    "/root",
			expected: "/root/.BRD",
		},
		{
			input:    "/root/",
			expected: "/root//.BRD",
		},
		{
			input:    "/root//",
			expected: "/root///.BRD",
		},
		{
			input:    "home/bbs",
			expected: "home/bbs/.BRD",
		},
	}

	for i, c := range cases {
		actual, err := GetBoardPath(c.input)
		if err != nil {
			t.Errorf("GetBoardPath err != nil on index %d", i)
		}
		if actual != c.expected {
			t.Errorf("GetBoardPath result not match on index %d with input:%v , expected: %v, got: %v",
				i, c.input, c.expected, actual)
		}
	}

}

func TestGetUserFavoritePath(t *testing.T) {

	type Input struct {
		wd     string
		userId string
	}
	type TestCase struct {
		input    Input
		expected string
	}
	cases := []TestCase{

		{
			input: Input{
				wd:     "/root",
				userId: "SYSOP",
			},
			expected: "/root/home/S/SYSOP/.fav",
		},
		{
			input: Input{
				wd:     "/root/",
				userId: "SYSOP",
			},
			expected: "/root//home/S/SYSOP/.fav",
		},
		{
			input: Input{
				wd:     "/root//",
				userId: "SYSOP",
			},
			expected: "/root///home/S/SYSOP/.fav",
		},
		{
			input: Input{
				wd:     "/root",
				userId: "sysop",
			},
			expected: "/root/home/s/sysop/.fav",
		},
	}

	for i, c := range cases {
		actual, err := GetUserFavoritePath(c.input.wd, c.input.userId)
		if err != nil {
			t.Errorf("GetUserFavoritePath err != nil on index %d", i)
		}
		if actual != c.expected {
			t.Errorf("GetUserFavoritePath result not match on index %d with input:%v , expected: %v, got: %v",
				i, c.input, c.expected, actual)
		}
	}

}

func TestGetLoginRecentPath(t *testing.T) {

	type Input struct {
		wd     string
		userId string
	}
	type TestCase struct {
		input    Input
		expected string
	}
	cases := []TestCase{

		{
			input: Input{
				wd:     "/root",
				userId: "SYSOP",
			},
			expected: "/root/home/S/SYSOP/logins.recent",
		},
		{
			input: Input{
				wd:     "/root/",
				userId: "SYSOP",
			},
			expected: "/root//home/S/SYSOP/logins.recent",
		},
		{
			input: Input{
				wd:     "/root//",
				userId: "SYSOP",
			},
			expected: "/root///home/S/SYSOP/logins.recent",
		},
		{
			input: Input{
				wd:     "/root",
				userId: "sysop",
			},
			expected: "/root/home/s/sysop/logins.recent",
		},
	}

	for i, c := range cases {
		actual, err := GetLoginRecentPath(c.input.wd, c.input.userId)
		if err != nil {
			t.Errorf("GetLoginRecentPath err != nil on index %d", i)
		}
		if actual != c.expected {
			t.Errorf("GetLoginRecentPath result not match on index %d with input:%v , expected: %v, got: %v",
				i, c.input, c.expected, actual)
		}
	}

}
