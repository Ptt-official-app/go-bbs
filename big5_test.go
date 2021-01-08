// Copyright 2020 Pichu Chen, The PTT APP Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bbs

import (
	"reflect"
	"testing"
)

func TestUtf8ToBig5(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		// TODO: Add test cases.
		{
			name:     "test0",
			args:     args{input: "新的目錄"},
			expected: "\xb7\x73\xaa\xba\xa5\xd8\xbf\xfd",
		},
		{
			name:     "test1",
			args:     args{input: "ピリカピリララ"},
			expected: "\xc7\xd0\xc7\xe6\xc7\xa7\xc7\xd0\xc7\xe6\xc7\xe5\xc7\xe5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Utf8ToBig5(tt.args.input); !reflect.DeepEqual(got, []byte(tt.expected)) {
				t.Errorf("Utf8ToBig5() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestBig5ToUtf8(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		// TODO: Add test cases.
		{
			name:     "test0",
			args:     args{input: "\xb7\x73\xaa\xba\xa5\xd8\xbf\xfd"},
			expected: "新的目錄",
		},
		{
			name:     "test1",
			args:     args{input: "\xc7\xd0\xc7\xe6\xc7\xa7\xc7\xd0\xc7\xe6\xc7\xe5\xc7\xe5"},
			expected: "ピリカピリララ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Big5ToUtf8([]byte(tt.args.input)); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Big5ToUtf8() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
