package cache

import (
	"testing"

	"github.com/PichuChen/go-bbs/ptttype"
	log "github.com/sirupsen/logrus"
)

func TestAttachSHM(t *testing.T) {
	setupTest()
	defer teardownTest()

	_ = NewSHM(testShmKey, ptttype.USE_HUGETLB, true)
	defer CloseSHM()

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

func TestDoSearchUser(t *testing.T) {
	setupTest()
	defer teardownTest()

	err := NewSHM(testShmKey, ptttype.USE_HUGETLB, true)
	if err != nil {
		log.Errorf("TestDoSearchUser: unable to NewSHM: e: %v", err)
		return
	}
	defer CloseSHM()

	_ = LoadUHash()

	type args struct {
		userID   string
		isReturn bool
	}
	tests := []struct {
		name     string
		args     args
		expected int32
		want1    string
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{userID: "SYSOP"},
			expected: 1,
			want1:    "",
		},
		{
			args:     args{userID: "SYSOP", isReturn: true},
			expected: 1,
			want1:    "SYSOP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := doSearchUser(tt.args.userID, tt.args.isReturn)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoSearchUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("DoSearchUser() got = %v, expected%v", got, tt.expected)
			}
			if got1 != tt.want1 {
				t.Errorf("DoSearchUser() got1 = %v, expected%v", got1, tt.want1)
			}
		})
	}
}
