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
		name     string
		args     args
		expected []byte
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args: args{
				key:  []byte("012345678901"),
				salt: []byte{65, 65, 51, 81, 66, 104, 76, 87, 107, 49, 66, 87, 65, 0},
			},
			expected: []byte{65, 65, 51, 81, 66, 104, 76, 87, 107, 49, 66, 87, 65, 0},
		},
		{
			args: args{
				key:  []byte("012345678901"),
				salt: []byte{65, 63, 51, 81, 66, 104, 76, 87, 107, 49, 66, 87, 65, 0},
			},
			expected: []byte{65, 63, 65, 52, 56, 79, 53, 115, 114, 113, 80, 83, 85, 0},
		},
		{
			args: args{
				key:  []byte("ABCD45678901"),
				salt: []byte{65, 65, 57, 86, 98, 117, 101, 90, 88, 111, 106, 65, 65, 0},
			},
			expected: []byte{65, 65, 57, 86, 98, 117, 101, 90, 88, 111, 106, 65, 65, 0},
		},
		{
			name: "key: 8 0's, salt: 9 0's",
			args: args{
				key:  []byte("00000000"),
				salt: []byte("000000000"),
			},
			expected: []byte("00CfV146ZJdLc\x00"),
		},
		{
			args: args{
				key:  []byte("000000001123123123"),
				salt: []byte("000000000"),
			},
			expected: []byte("00CfV146ZJdLc\x00"),
		},
		{
			args: args{
				key:  []byte("00000000112312sdfasdf3123"),
				salt: []byte("000010011asfasdfsaf"),
			},
			expected: []byte("00CfV146ZJdLc\x00"),
		},
		{
			args: args{
				key:  []byte("00000000112312sdfasdf3123"),
				salt: []byte("021010011asfasdfsaf"),
			},
			expected: []byte("02v6ADqeCsb12\x00"),
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
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Fcrypt() = %v, expected %v", got, tt.expected)
			}
			gots[idx] = got
		})
	}

	if reflect.DeepEqual(gots[0], gots[1]) {
		t.Errorf("Fcrypt: affected with multiple calls gots[0]: %v gots[1]: %v", gots[0], gots[1])
	}
}
