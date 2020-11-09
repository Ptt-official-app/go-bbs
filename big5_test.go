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
		name string
		args args
		want Big5String
	}{
		// TODO: Add test cases.
		{
			name: "test0",
			args: args{input: "新的目錄"},
			want: Big5String{183, 115, 170, 186, 165, 216, 191, 253},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Utf8ToBig5(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Utf8ToBig5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBig5ToUtf8(t *testing.T) {
	type args struct {
		input Big5String
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test0",
			args: args{input: Big5String{183, 115, 170, 186, 165, 216, 191, 253}},
			want: "新的目錄",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Big5ToUtf8(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Big5ToUtf8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBig5BytesToUtf8(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test0",
			args: args{input: []byte{183, 115, 170, 186, 165, 216, 191, 253}},
			want: "新的目錄",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Big5ToUtf8(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Big5ToUtf8() = %v, want %v", got, tt.want)
			}
		})
	}
}
