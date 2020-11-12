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
		expected []byte
	}{
		// TODO: Add test cases.
		{
			name:     "test0",
			args:     args{input: "新的目錄"},
			expected: []byte{0xb7, 0x73, 0xaa, 0xba, 0xa5, 0xd8, 0xbf, 0xfd},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Utf8ToBig5(tt.args.input); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Utf8ToBig5() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestBig5ToUtf8(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		// TODO: Add test cases.
		{
			name:     "test0",
			args:     args{input: []byte{0xb7, 0x73, 0xaa, 0xba, 0xa5, 0xd8, 0xbf, 0xfd}},
			expected: "新的目錄",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Big5ToUtf8(tt.args.input); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Big5ToUtf8() = %v, want %v", got, tt.expected)
			}
		})
	}
}
