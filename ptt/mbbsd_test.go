package ptt

import (
	"reflect"
	"testing"

	"github.com/PichuChen/go-bbs/ptttype"
)

func TestLoginQuery(t *testing.T) {
	setupTest()
	defer teardownTest()

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
		/*
			{
				args: args{userID: &[ptttype.IDLEN + 1]byte{'S', 'Y', 'S', 'O', 'P'}, passwd: []byte{'1', '2', '3'}},
				want: testUserecBig51,
			},
		*/
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
