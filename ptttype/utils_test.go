package ptttype

import "testing"

func TestBytesLen(t *testing.T) {
	setupTest()
	defer teardownTest()

	str1 := [13]byte{}
	str2 := [13]byte{'a', 'b', 'c'}
	str3 := [10]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	str4 := [10]byte{'0', '1', '2', '3', '4', 0, '6', '7', '8', '9'}

	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "init",
			args: args{str1[:]},
			want: 0,
		},
		{
			name: "with only 3 letters",
			args: args{str2[:]},
			want: 3,
		},
		{
			name: "with no 0",
			args: args{str3[:]},
			want: 10,
		},
		{
			name: "cutoff at str4[5]",
			args: args{str4[:]},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FixedBytesLen(tt.args.bytes); got != tt.want {
				t.Errorf("BytesLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
