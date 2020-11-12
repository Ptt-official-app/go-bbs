package bbs

import (
	"reflect"
	"testing"
)

func TestNewUserecFromBig5(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		userecraw *UserecRaw
	}
	tests := []struct {
		name     string
		args     args
		expected *Userec
	}{
		// TODO: Add test cases.
		{
			args:     args{testUserecBig51},
			expected: testUserec1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserecFromRaw(tt.args.userecraw); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("NewUserecFromBig5() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestOpenUserecFile(t *testing.T) {
	setupTest()
	defer teardownTest()

	type args struct {
		filename string
	}
	tests := []struct {
		name     string
		args     args
		expected []*Userec
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{"testcase/passwd/01.PASSWDS"},
			expected: testOpenUserecFile1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OpenUserecFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenUserecFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("OpenUserecFile() = %v, want %v", got, tt.expected)
			}
		})
	}
}
