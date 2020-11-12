package bbs

import (
	"reflect"
	"testing"
)

func TestCstrToBytes(t *testing.T) {
	setupTest()
	defer teardownTest()

	str1 := [13]byte{}
	str2 := [13]byte{}
	copy(str2[:], []byte("123"))
	str3 := [10]byte{}
	copy(str3[:], []byte("0123456789"))
	str4 := [10]byte{}
	copy(str4[:], []byte("01234\x006789"))

	type args struct {
		cstr Cstr
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "init",
			args: args{str1[:]},
			want: []byte{},
		},
		{
			name: "with only 3 letters",
			args: args{str2[:]},
			want: []byte("123"),
		},
		{
			name: "with no 0",
			args: args{str3[:]},
			want: []byte("0123456789"),
		},
		{
			name: "cutoff at str4[5]",
			args: args{str4[:]},
			want: []byte("01234"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CstrToBytes(tt.args.cstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CstrToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCstrToString(t *testing.T) {
	setupTest()
	defer teardownTest()

	str1 := [13]byte{}
	str2 := [13]byte{}
	copy(str2[:], []byte("123"))
	str3 := [10]byte{}
	copy(str3[:], []byte("0123456789"))
	str4 := [10]byte{}
	copy(str4[:], []byte("01234\x006789"))

	type args struct {
		cstr Cstr
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "init",
			args: args{str1[:]},
			want: "",
		},
		{
			name: "with only 3 letters",
			args: args{str2[:]},
			want: "123",
		},
		{
			name: "with no 0",
			args: args{str3[:]},
			want: "0123456789",
		},
		{
			name: "cutoff at str4[5]",
			args: args{str4[:]},
			want: "01234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CstrToString(tt.args.cstr); got != tt.want {
				t.Errorf("CstrToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
