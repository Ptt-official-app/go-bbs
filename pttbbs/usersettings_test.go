// 實作幾個 method 傳入 userid 回傳該使用者的 BBS 設定
// 目前已經有可以取得使用者userec的method可用了
//
// 請見 userec_t 中的 uflag 和 userlevel
// 每個選項對應的 flag id 請參見 user.c desc1 和 masks1
// user.c - https://github.com/ptt/pttbbs/blob/5715b35f510f48eb5092d32882f1aa09181dc3a1/mbbsd/user.c#L438
// uflags.h - https://github.com/ptt/pttbbs/blob/4d56e77f264960e43e060b77e442e166e5706417/include/uflags.h

package pttbbs

import (
	"reflect"
	"testing"
)

func Test_getUserFlagAllByID(t *testing.T) {
	type args struct {
		userid string
	}
	tests := []struct {
		name    string
		args    args
		want    *uFlags
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getUserFlagAllByID(tt.args.userid)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUserFlagAllByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUserFlagAllByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
