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
				t.Errorf("NewUserecFromBig5() = %v, expected %v", got, tt.expected)
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
		{
			args:     args{"testcase/passwd/01.PASSWDS.corrupt"},
			expected: testOpenUserecFile1[:TEST_N_OPEN_USER_FILE_1-1],
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OpenUserecFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenUserecFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for idx, eachGot := range got {
				each := eachGot.Version
				expected := tt.expected[idx].Version
				if !reflect.DeepEqual(each, expected) {
					t.Errorf("(%v/%v) Version: OpenUserecFile() = %v, expected %v", idx, len(got), each, expected)
				}
			}

			for idx, eachGot := range got {
				each := eachGot.Userid
				expected := tt.expected[idx].Userid
				if !reflect.DeepEqual(each, expected) {
					t.Errorf("(%v/%v) Userid: OpenUserecFile() = %v, expected %v", idx, len(got), each, expected)
				}
			}

			for idx, eachGot := range got {
				each := eachGot.Realname
				expected := tt.expected[idx].Realname
				if !reflect.DeepEqual(each, expected) {
					t.Errorf("(%v/%v) Realname: OpenUserecFile() = %v, expected %v", idx, len(got), each, expected)
				}
			}

			for idx, eachGot := range got {
				each := eachGot.Nickname
				expected := tt.expected[idx].Nickname
				if !reflect.DeepEqual(each, expected) {
					t.Errorf("(%v/%v) Nickname: OpenUserecFile() = %v, expected %v", idx, len(got), each, expected)
				}
			}

			for idx, eachGot := range got {
				each := eachGot.Passwd
				expected := tt.expected[idx].Passwd
				if !reflect.DeepEqual(each, expected) {
					t.Errorf("(%v/%v) Passwd: OpenUserecFile() = %v, expected %v", idx, len(got), each, expected)
				}
			}

			for idx, eachGot := range got {
				each := eachGot.Pad1
				expected := tt.expected[idx].Pad1
				if !reflect.DeepEqual(each, expected) {
					t.Errorf("(%v/%v) Pad1: OpenUserecFile() = %v, expected %v", idx, len(got), each, expected)
				}
			}

			for idx, eachGot := range got {
				each := eachGot.Uflag
				expected := tt.expected[idx].Uflag
				if !reflect.DeepEqual(each, expected) {
					t.Errorf("(%v/%v) Uflag: OpenUserecFile() = %v, expected %v", idx, len(got), each, expected)
				}
			}

			for idx, eachGot := range got {
				each := eachGot._unused1
				expected := tt.expected[idx]._unused1
				if !reflect.DeepEqual(each, expected) {
					t.Errorf("(%v/%v) _unused1: OpenUserecFile() = %v, expected %v", idx, len(got), each, expected)
				}
			}

			for idx, eachGot := range got {
				each := eachGot.Userlevel
				expected := tt.expected[idx].Userlevel
				if !reflect.DeepEqual(each, expected) {
					t.Errorf("(%v/%v) Userlevel: OpenUserecFile() = %v, expected %v", idx, len(got), each, expected)
				}
			}

			for idx, eachGot := range got {
				each := eachGot.Numlogindays
				expected := tt.expected[idx].Numlogindays
				if !reflect.DeepEqual(each, expected) {
					t.Errorf("(%v/%v) Numlogindays: OpenUserecFile() = %v, expected %v", idx, len(got), each, expected)
				}
			}

			for idx, eachGot := range got {
				each := eachGot.Numposts
				expected := tt.expected[idx].Numposts
				if !reflect.DeepEqual(each, expected) {
					t.Errorf("(%v/%v) Numposts: OpenUserecFile() = %v, expected %v", idx, len(got), each, expected)
				}
			}

			for idx, eachGot := range got {
				each := eachGot.Firstlogin
				expected := tt.expected[idx].Firstlogin
				if !reflect.DeepEqual(each, expected) {
					t.Errorf("(%v/%v) Firstlogin: OpenUserecFile() = %v, expected %v", idx, len(got), each, expected)
				}
			}

			for idx, eachGot := range got {
				each := eachGot.Lastlogin
				expected := tt.expected[idx].Lastlogin
				if !reflect.DeepEqual(each, expected) {
					t.Errorf("(%v/%v) Lastlogin: OpenUserecFile() = %v, expected %v", idx, len(got), each, expected)
				}
			}

			for idx, eachGot := range got {
				each := eachGot.Lasthost
				expected := tt.expected[idx].Lasthost
				if !reflect.DeepEqual(each, expected) {
					t.Errorf("(%v/%v) Lasthost: OpenUserecFile() = %v, expected %v", idx, len(got), each, expected)
				}
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("OpenUserecFile() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
