package ptt

import (
	"reflect"
	"testing"

	"github.com/PichuChen/go-bbs/ptttype"
)

func TestLoginQuery(t *testing.T) {
	setupTest()
	defer teardownTest()

	userid1 := [ptttype.IDLEN + 1]byte{}
	copy(userid1[:], []byte("SYSOP"))

	type args struct {
		userID   *[ptttype.IDLEN + 1]byte
		passwd   []byte
		ip       [ptttype.IPV4LEN + 1]byte
		isHashed bool
	}
	tests := []struct {
		name    string
		args    args
		want    *ptttype.UserecBig5
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{userID: &userid1, passwd: []byte("123123")},
			want: testUserecBig51,
		},
		{
			args:    args{userID: &userid1, passwd: []byte("124")},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoginQuery(tt.args.userID, tt.args.passwd, tt.args.ip, tt.args.isHashed)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
