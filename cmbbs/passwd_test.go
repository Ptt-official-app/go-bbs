package cmbbs

import (
	"reflect"
	"testing"

	"github.com/PichuChen/go-bbs/ptttype"
)

func TestPasswdLoadUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	userID1 := [ptttype.IDLEN + 1]byte{}
	copy(userID1[:], []byte("SYSOP"))

	type args struct {
		userID *[ptttype.IDLEN + 1]byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   *ptttype.UserecBig5
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args:  args{&userID1},
			want:  1,
			want1: testUserecBig51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PasswdLoadUser(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("PasswdLoadUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PasswdLoadUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PasswdLoadUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPasswdQuery(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		want    *ptttype.UserecBig5
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{1},
			want: testUserecBig51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PasswdQuery(tt.args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("PasswdQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PasswdQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPasswd(t *testing.T) {
	input1 := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '1', 0}
	input2 := []byte{'0', '1', '2', '4', '4', '5', '6', '7', '8', '9', '0', '1', 0}
	expected1 := [ptttype.PASSLEN]byte{65, 65, 51, 81, 66, 104, 76, 87, 107, 49, 66, 87, 65}

	type args struct {
		expected []byte
		input    []byte
		isHashed bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{expected1[:], input1, false},
			want: true,
		},
		{
			name: "already hashed",
			args: args{expected1[:], expected1[:], true},
			want: true,
		},
		{
			name: "incorrect input",
			args: args{expected1[:], input2, false},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckPasswd(tt.args.expected, tt.args.input, tt.args.isHashed)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckPasswd() = %v, want %v", got, tt.want)
			}
		})
	}
}
