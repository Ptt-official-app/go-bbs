package bbs

import (
	"testing"
)

func TestParseFileHeader(t *testing.T) {

	headers, err := OpenFileHeaderFile("testcase/file/01.DIR")
	if err != nil {
		t.Error(err)
	}

	expected := []FileHeader{
		{
			Filename: "M.1599059246.A.CF6",
		},
		{
			Filename: "M.1599059415.A.FBA",
		},
		{
			Filename: "M.1599059496.A.2BE",
		},
	}

	for index, header := range headers {
		if header.Filename != expected[index].Filename {
			t.Logf("lena :%d %d", len(header.Filename), len(expected[index].Filename))
			t.Errorf("Filename not match in index %d, expected: %s, got: %s", index, expected[index].Filename, header.Filename)
		}

	}

}
