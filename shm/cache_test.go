package shm

import (
	"testing"
)

func TestAttachSHM(t *testing.T) {
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
