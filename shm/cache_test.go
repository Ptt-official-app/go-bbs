package shm

import (
	"testing"

	"reflect"

	"github.com/PichuChen/go-bbs/ptttype"
)

func TestAttachSHM(t *testing.T) {
	setupTest()
	defer teardownTest()

	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AttachSHM(); (err != nil) != tt.wantErr {
				t.Errorf("AttachSHM() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_doSearchUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		userID   string
		isReturn bool
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args:  args{userID: "SYSOP"},
			want:  1,
			want1: "",
		},
		{
			args:  args{userID: "SYSOP", isReturn: true},
			want:  1,
			want1: "SYSOP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := doSearchUser(tt.args.userID, tt.args.isReturn)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoSearchUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DoSearchUser() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DoSearchUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}

}

func Test_doSearchUserBig5(t *testing.T) {
	setupTest()
	defer teardownTest()

	userID1 := [ptttype.IDLEN + 1]byte{}
	copy(userID1[:5], []byte("SYSOP"))

	type args struct {
		userID   []byte
		isReturn bool
	}
	tests := []struct {
		name    string
		args    args
		want    int
		want1   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args:  args{userID: userID1[:]},
			want:  1,
			want1: nil,
		},
		{
			args:  args{userID: userID1[:], isReturn: true},
			want:  1,
			want1: userID1[:],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := doSearchUserBig5(tt.args.userID, tt.args.isReturn)
			if (err != nil) != tt.wantErr {
				t.Errorf("doSearchUserBig5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("doSearchUserBig5() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("doSearchUserBig5() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
