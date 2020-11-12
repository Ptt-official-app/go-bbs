package bbs

import (
	"reflect"
	"testing"

	"github.com/PichuChen/go-bbs/ptttype"
)

func TestNewUserecFromBig5(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		userBig5 *ptttype.UserecBig5
	}
	tests := []struct {
		name string
		args args
		want *Userec
	}{
		// TODO: Add test cases.
		{
			args: args{testUserecBig51},
			want: testUserec1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserecFromBig5(tt.args.userBig5); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserecFromBig5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOpenUserecFile(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    []*Userec
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{"testcase/passwd/01.PASSWDS"},
			want: testOpenUserecFile1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OpenUserecFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenUserecFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenUserecFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
