package cmbbs

import (
	"testing"

	"github.com/PichuChen/go-bbs/ptttype"
)

func TestIsValidUserID(t *testing.T) {
	setupTest()
	defer teardownTest()

	userID0 := [ptttype.IDLEN + 1]byte{}

	userID1 := [ptttype.IDLEN + 1]byte{}
	copy(userID1[:], []byte("S"))

	userID2 := [ptttype.IDLEN + 1]byte{}
	copy(userID2[:], []byte("SYSOP"))

	userID3 := [ptttype.IDLEN + 1]byte{}
	copy(userID3[:], []byte("S1234567891234"))

	userID4 := [ptttype.IDLEN + 1]byte{}
	copy(userID4[:], []byte("SYSOP,-"))

	userID5 := [ptttype.IDLEN + 1]byte{}
	copy(userID5[:], []byte("SYSOP1"))

	userID6 := [ptttype.IDLEN + 1]byte{}
	copy(userID6[:], []byte("1SYSOP"))

	userID7 := [ptttype.IDLEN + 1]byte{}
	copy(userID7[:], []byte("S1"))

	userID8 := [ptttype.IDLEN + 1]byte{}
	copy(userID8[:], []byte("Ss"))

	type args struct {
		userID *[ptttype.IDLEN + 1]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "nil",
			args: args{nil},
			want: false,
		},
		{
			name: "",
			args: args{&userID0},
			want: false,
		},
		{
			name: "S",
			args: args{&userID1},
			want: false,
		},
		{
			name: "SYSOP",
			args: args{&userID2},
			want: true,
		},
		{
			name: "too long",
			args: args{&userID3},
			want: false,
		},
		{
			name: "not alnum",
			args: args{&userID4},
			want: false,
		},
		{
			name: "SYSOP1",
			args: args{&userID5},
			want: true,
		},
		{
			name: "1SYSOP (not alpha in 0st char)",
			args: args{&userID6},
			want: false,
		},
		{
			name: "S1 (alnum)",
			args: args{&userID7},
			want: true,
		},
		{
			name: "Ss (all alpha)",
			args: args{&userID7},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidUserID(tt.args.userID); got != tt.want {
				t.Errorf("IsValidUserID() = %v, want %v", got, tt.want)
			}
		})
	}
}
