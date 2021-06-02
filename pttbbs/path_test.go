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
		userID string
	}
	type TestCase struct {
		input    Input
		expected string
	}
	cases := []TestCase{

		{
			input: Input{
				wd:     "/root",
				userID: "SYSOP",
			},
			expected: "/root/home/S/SYSOP/.fav",
		},
		{
			input: Input{
				wd:     "/root/",
				userID: "SYSOP",
			},
			expected: "/root//home/S/SYSOP/.fav",
		},
		{
			input: Input{
				wd:     "/root//",
				userID: "SYSOP",
			},
			expected: "/root///home/S/SYSOP/.fav",
		},
		{
			input: Input{
				wd:     "/root",
				userID: "sysop",
			},
			expected: "/root/home/s/sysop/.fav",
		},
	}

	for i, c := range cases {
		actual, err := GetUserFavoritePath(c.input.wd, c.input.userID)
		if err != nil {
			t.Errorf("GetUserFavoritePath err != nil on index %d", i)
		}
		if actual != c.expected {
			t.Errorf("GetUserFavoritePath result not match on index %d with input:%v , expected: %v, got: %v",
				i, c.input, c.expected, actual)
		}
	}

}

func TestGetUserDraftPath(t *testing.T) {

	type Input struct {
		wd      string
		userID  string
		draftID string
	}
	type TestCase struct {
		input    Input
		expected string
	}
	cases := []TestCase{

		{
			input: Input{
				wd:      "/root",
				userID:  "SYSOP",
				draftID: "1",
			},
			expected: "/root/home/S/SYSOP/buf.1",
		},
		{
			input: Input{
				wd:      "/root/",
				userID:  "SYSOP",
				draftID: "2",
			},
			expected: "/root//home/S/SYSOP/buf.2",
		},
		{
			input: Input{
				wd:      "/root//",
				userID:  "SYSOP",
				draftID: "3",
			},
			expected: "/root///home/S/SYSOP/buf.3",
		},
		{
			input: Input{
				wd:      "/root",
				userID:  "sysop",
				draftID: "4",
			},
			expected: "/root/home/s/sysop/buf.4",
		},
	}

	for i, c := range cases {
		actual, err := GetUserDraftPath(c.input.wd, c.input.userID, c.input.draftID)
		if err != nil {
			t.Errorf("GetUserDraftPath err != nil on index %d", i)
		}
		if actual != c.expected {
			t.Errorf("GetUserDraftPath result not match on index %d with input:%v , expected: %v, got: %v",
				i, c.input, c.expected, actual)
		}
	}

}

func TestGetUserMailPath(t *testing.T) {

	type Input struct {
		wd       string
		userID   string
		filename string
	}
	type TestCase struct {
		input    Input
		expected string
	}
	cases := []TestCase{

		{
			input: Input{
				wd:       "/root",
				userID:   "SYSOP",
				filename: "M.1600751073.A.BC9",
			},
			expected: "/root/home/S/SYSOP/M.1600751073.A.BC9",
		},
		{
			input: Input{
				wd:       "/root/",
				userID:   "SYSOP",
				filename: "M.16007514073.A.BC9",
			},
			expected: "/root//home/S/SYSOP/M.16007514073.A.BC9",
		},
		{
			input: Input{
				wd:       "/root//",
				userID:   "SYSOP",
				filename: "M.2600751073.A.BC9",
			},
			expected: "/root///home/S/SYSOP/M.2600751073.A.BC9",
		},
		{
			input: Input{
				wd:       "/root",
				userID:   "sysop",
				filename: "M.1600751073.B.BC9",
			},
			expected: "/root/home/s/sysop/M.1600751073.B.BC9",
		},
	}

	for i, c := range cases {
		actual, err := GetUserMailPath(c.input.wd, c.input.userID, c.input.filename)
		if err != nil {
			t.Errorf("GetUserMailPath err != nil on index %d", i)
		}
		if actual != c.expected {
			t.Errorf("GetUserMailPath result not match on index %d with input:%v , expected: %v, got: %v",
				i, c.input, c.expected, actual)
		}
	}

}

func TestGetLoginRecentPath(t *testing.T) {

	type Input struct {
		wd     string
		userID string
	}
	type TestCase struct {
		input    Input
		expected string
	}
	cases := []TestCase{

		{
			input: Input{
				wd:     "/root",
				userID: "SYSOP",
			},
			expected: "/root/home/S/SYSOP/logins.recent",
		},
		{
			input: Input{
				wd:     "/root/",
				userID: "SYSOP",
			},
			expected: "/root//home/S/SYSOP/logins.recent",
		},
		{
			input: Input{
				wd:     "/root//",
				userID: "SYSOP",
			},
			expected: "/root///home/S/SYSOP/logins.recent",
		},
		{
			input: Input{
				wd:     "/root",
				userID: "sysop",
			},
			expected: "/root/home/s/sysop/logins.recent",
		},
	}

	for i, c := range cases {
		actual, err := GetLoginRecentPath(c.input.wd, c.input.userID)
		if err != nil {
			t.Errorf("GetLoginRecentPath err != nil on index %d", i)
		}
		if actual != c.expected {
			t.Errorf("GetLoginRecentPath result not match on index %d with input:%v , expected: %v, got: %v",
				i, c.input, c.expected, actual)
		}
	}

}

func TestGetBoardArticlesDirectoryPath(t *testing.T) {

	type Input struct {
		wd      string
		boardID string
	}
	type TestCase struct {
		input    Input
		expected string
	}
	cases := []TestCase{

		{
			input: Input{
				wd:      "/root",
				boardID: "SYSOP",
			},
			expected: "/root/boards/S/SYSOP/.DIR",
		},
		{
			input: Input{
				wd:      "/root/",
				boardID: "SYSOP",
			},
			expected: "/root//boards/S/SYSOP/.DIR",
		},
		{
			input: Input{
				wd:      "/root//",
				boardID: "SYSOP",
			},
			expected: "/root///boards/S/SYSOP/.DIR",
		},
		{
			input: Input{
				wd:      "/root",
				boardID: "sysop",
			},
			expected: "/root/boards/s/sysop/.DIR",
		},
	}

	for i, c := range cases {
		actual, err := GetBoardArticlesDirectoryPath(c.input.wd, c.input.boardID)
		if err != nil {
			t.Errorf("GetBoardArticlesDirectoryPath err != nil on index %d", i)
		}
		if actual != c.expected {
			t.Errorf("GetBoardArticlesDirectoryPath result not match on index %d with input:%v , expected: %v, got: %v",
				i, c.input, c.expected, actual)
		}
	}

}

func TestGetBoardArticleFilePath(t *testing.T) {

	type Input struct {
		wd       string
		boardID  string
		filename string
	}
	type TestCase struct {
		input    Input
		expected string
	}
	cases := []TestCase{

		{
			input: Input{
				wd:       "/root",
				boardID:  "SYSOP",
				filename: "M.16007514073.A.BC9",
			},
			expected: "/root/boards/S/SYSOP/M.16007514073.A.BC9",
		},
		{
			input: Input{
				wd:       "/root/",
				boardID:  "SYSOP",
				filename: "M.16007514073.A.BC9",
			},
			expected: "/root//boards/S/SYSOP/M.16007514073.A.BC9",
		},
		{
			input: Input{
				wd:       "/root//",
				boardID:  "SYSOP",
				filename: "M.16007514073.A.BC9",
			},
			expected: "/root///boards/S/SYSOP/M.16007514073.A.BC9",
		},
		{
			input: Input{
				wd:       "/root",
				boardID:  "sysop",
				filename: "M.16007514073.A.BC9",
			},
			expected: "/root/boards/s/sysop/M.16007514073.A.BC9",
		},
	}

	for i, c := range cases {
		actual, err := GetBoardArticleFilePath(c.input.wd, c.input.boardID, c.input.filename)
		if err != nil {
			t.Errorf("GetBoardArticleFilePath err != nil on index %d", i)
		}
		if actual != c.expected {
			t.Errorf("GetBoardArticleFilePath result not match on index %d with input:%v , expected: %v, got: %v",
				i, c.input, c.expected, actual)
		}
	}

}

func TestGetBoardTreasuresDirectoryPath(t *testing.T) {

	type Input struct {
		wd      string
		boardID string
		path    []string
	}
	type TestCase struct {
		input    Input
		expected string
	}
	cases := []TestCase{

		{
			input: Input{
				wd:      "/root",
				boardID: "SYSOP",
				path:    []string{},
			},
			expected: "/root/man/boards/S/SYSOP/.DIR",
		},
		{
			input: Input{
				wd:      "/root",
				boardID: "SYSOP",
				path: []string{
					"D3D1",
				},
			},
			expected: "/root/man/boards/S/SYSOP/D3D1/.DIR",
		},
		{
			input: Input{
				wd:      "/root",
				boardID: "SYSOP",
				path: []string{
					"D3D1",
					"D3D2",
				},
			},
			expected: "/root/man/boards/S/SYSOP/D3D1/D3D2/.DIR",
		},
	}

	for i, c := range cases {
		actual, err := GetBoardTreasuresDirectoryPath(c.input.wd, c.input.boardID, c.input.path)
		if err != nil {
			t.Errorf("GetBoardTreasuresDirectoryPath err != nil on index %d", i)
		}
		if actual != c.expected {
			t.Errorf("GetBoardTreasuresDirectoryPath result not match on index %d with input:%v , expected: %v, got: %v",
				i, c.input, c.expected, actual)
		}
	}

}

func TestGetBoardTreasureFilePath(t *testing.T) {

	type Input struct {
		wd       string
		boardID  string
		path     []string
		filename string
	}
	type TestCase struct {
		input    Input
		expected string
	}
	cases := []TestCase{

		{
			input: Input{
				wd:       "/root",
				boardID:  "SYSOP",
				path:     []string{},
				filename: "M.16007514073.A.BC9",
			},
			expected: "/root/man/boards/S/SYSOP/M.16007514073.A.BC9",
		},
		{
			input: Input{
				wd:      "/root",
				boardID: "sysop",
				path: []string{
					"D3D1",
				},
				filename: "M.16007514073.A.BC9",
			},
			expected: "/root/man/boards/s/sysop/D3D1/M.16007514073.A.BC9",
		},
		{
			input: Input{
				wd:      "/root",
				boardID: "sysop",
				path: []string{
					"D3D1",
					"D3D2",
				},
				filename: "M.16007514073.A.BC9",
			},
			expected: "/root/man/boards/s/sysop/D3D1/D3D2/M.16007514073.A.BC9",
		},
	}

	for i, c := range cases {
		actual, err := GetBoardTreasureFilePath(c.input.wd, c.input.boardID, c.input.path, c.input.filename)
		if err != nil {
			t.Errorf("GetBoardTreasureFilePath err != nil on index %d", i)
		}
		if actual != c.expected {
			t.Errorf("GetBoardTreasureFilePath result not match on index %d with input:%v , expected: %v, got: %v",
				i, c.input, c.expected, actual)
		}
	}

}

func TestGetBoardNameFilePath(t *testing.T) {

	type Input struct {
		wd      string
		boardID string
	}
	type TestCase struct {
		input    Input
		expected string
	}
	cases := []TestCase{

		{
			input: Input{
				wd:      "/root",
				boardID: "SYSOP",
			},
			expected: "/root/boards/S/SYSOP/.Name",
		},
		{
			input: Input{
				wd:      "/root/",
				boardID: "SYSOP",
			},
			expected: "/root//boards/S/SYSOP/.Name",
		},
		{
			input: Input{
				wd:      "/root//",
				boardID: "SYSOP",
			},
			expected: "/root///boards/S/SYSOP/.Name",
		},
		{
			input: Input{
				wd:      "/root",
				boardID: "sysop",
			},
			expected: "/root/boards/s/sysop/.Name",
		},
	}

	for i, c := range cases {
		actual, err := GetBoardNameFilePath(c.input.wd, c.input.boardID)
		if err != nil {
			t.Errorf("GetBoardNameFilePath err != nil on index %d", i)
		}
		if actual != c.expected {
			t.Errorf("GetBoardNameFilePath result not match on index %d with input:%v , expected: %v, got: %v",
				i, c.input, c.expected, actual)
		}
	}

}
