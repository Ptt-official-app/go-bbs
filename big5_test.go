package bbs

import (
	"reflect"
	"testing"
)

func TestUtf8ToBig5(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		// TODO: Add test cases.
		{
			name:     "test0",
			args:     args{input: "新的目錄"},
			expected: "\xb7\x73\xaa\xba\xa5\xd8\xbf\xfd",
		},
		{
			name:     "test1",
			args:     args{input: "ピリカピリララ"},
			expected: "\xc7\xd0\xc7\xe6\xc7\xa7\xc7\xd0\xc7\xe6\xc7\xe5\xc7\xe5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Utf8ToBig5(tt.args.input); !reflect.DeepEqual(got, []byte(tt.expected)) {
				t.Errorf("Utf8ToBig5() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestBig5ToUtf8(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		// TODO: Add test cases.
		{
			name:     "test0",
			args:     args{input: "\xb7\x73\xaa\xba\xa5\xd8\xbf\xfd"},
			expected: "新的目錄",
		},
		{
			name:     "test1",
			args:     args{input: "\xc7\xd0\xc7\xe6\xc7\xa7\xc7\xd0\xc7\xe6\xc7\xe5\xc7\xe5"},
			expected: "ピリカピリララ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Big5ToUtf8([]byte(tt.args.input)); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Big5ToUtf8() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
