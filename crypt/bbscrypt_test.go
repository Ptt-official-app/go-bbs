package crypt

import (
	"reflect"
	"testing"
)

func TestFcrypt(t *testing.T) {
	type args struct {
		key  []byte
		salt []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				key:  []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '1', 0},
				salt: []byte{65, 65, 51, 81, 66, 104, 76, 87, 107, 49, 66, 87, 65, 0},
			},
			want: []byte{65, 65, 51, 81, 66, 104, 76, 87, 107, 49, 66, 87, 65, 0},
		},
		{
			args: args{
				key:  []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '1', 0},
				salt: []byte{65, 63, 51, 81, 66, 104, 76, 87, 107, 49, 66, 87, 65, 0},
			},
			want: []byte{65, 63, 65, 52, 56, 79, 53, 115, 114, 113, 80, 83, 85, 0},
		},
		{
			args: args{
				key:  []byte{'A', 'B', 'C', 'D', '4', '5', '6', '7', '8', '9', '0', '1', 0},
				salt: []byte{65, 65, 57, 86, 98, 117, 101, 90, 88, 111, 106, 65, 65, 0},
			},
			want: []byte{65, 65, 57, 86, 98, 117, 101, 90, 88, 111, 106, 65, 65, 0},
		},
		{
			args: args{
				key:  []byte{'0', '0', '0', '0', '0', '0', '0', '0', 0},
				salt: []byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', 0},
			},
			want: []byte{'0', '0', 'C', 'f', 'V', '1', '4', '6', 'Z', 'J', 'd', 'L', 'c', 0},
		},
		{
			args: args{
				key:  []byte{'0', '0', '0', '0', '0', '0', '0', '0', '1', '1', '2', '3', '1', '2', '3', '1', '2', '3', 0},
				salt: []byte{'0', '0', '0', '0', '0', '0', '0', '0', '0', 0},
			},
			want: []byte{'0', '0', 'C', 'f', 'V', '1', '4', '6', 'Z', 'J', 'd', 'L', 'c', 0},
		},
		{
			args: args{
				key:  []byte{'0', '0', '0', '0', '0', '0', '0', '0', '1', '1', '2', '3', '1', '2', 's', 'd', 'f', 'a', 's', 'd', 'f', '3', '1', '2', '3', 0},
				salt: []byte{'0', '0', '0', '0', '1', '0', '0', '1', '1', 'a', 's', 'f', 'a', 's', 'd', 'f', 's', 'a', 'f', 0},
			},
			want: []byte{'0', '0', 'C', 'f', 'V', '1', '4', '6', 'Z', 'J', 'd', 'L', 'c', 0},
		},
		{
			args: args{
				key:  []byte{'0', '0', '0', '0', '0', '0', '0', '0', '1', '1', '2', '3', '1', '2', 's', 'd', 'f', 'a', 's', 'd', 'f', '3', '1', '2', '3', 0},
				salt: []byte{'0', '2', '1', '0', '1', '0', '0', '1', '1', 'a', 's', 'f', 'a', 's', 'd', 'f', 's', 'a', 'f', 0},
			},
			want: []byte{'0', '2', 'v', '6', 'A', 'D', 'q', 'e', 'C', 's', 'b', '1', '2', 0},
		},
	}

	gots := make([][]byte, len(tests))
	for idx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Fcrypt(tt.args.key, tt.args.salt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fcrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fcrypt() = %v, want %v", got, tt.want)
			}
			gots[idx] = got
		})
	}

	if reflect.DeepEqual(gots[0], gots[1]) {
		t.Errorf("Fcrypt: affected with multiple calls gots[0]: %v gots[1]: %v", gots[0], gots[1])
	}
}
