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
				salt: []byte("AA3QBhLWk1BWA"),
			},
			expected: []byte("AA3QBhLWk1BWA\x00"),
		},
		{
			args: args{
				key:  []byte("012345678901"),
				salt: []byte("A?3QBhLWk1BWA"),
			},
			expected: []byte("A?A48O5srqPSU\x00"),
		},
		{
			args: args{
				key:  []byte("ABCD45678901"),
				salt: []byte("AA9VbueZXojAA"),
			},
			expected: []byte("AA9VbueZXojAA\x00"),
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
		{
			args: args{
				key:  []byte("123123"),
				salt: []byte("bhwvOJtfT1TAI"),
			},
			expected: []byte("bhwvOJtfT1TAI\x00"),
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
}
