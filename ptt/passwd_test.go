package ptt

import (
	"reflect"
	"testing"

	"github.com/PichuChen/go-bbs/ptttype"
)

func Test_initCurrentUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	userid1 := [ptttype.IDLEN + 1]byte{}
	copy(userid1[:], []byte("SYSOP"))

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
			args:  args{&userid1},
			want:  1,
			want1: testUserecBig51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := initCurrentUser(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("initCurrentUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("initCurrentUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("initCurrentUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
