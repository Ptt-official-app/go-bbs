package cmsys

import "testing"

func Test_toupper(t *testing.T) {
	type args struct {
		theByte byte
	}
	tests := []struct {
		name     string
		args     args
		expected byte
	}{
		// TODO: Add test cases.
		{
			args:     args{'a'},
			expected: 'A',
		},
		{
			args:     args{'A'},
			expected: 'A',
		},
		{
			args:     args{'0'},
			expected: '0',
		},
		{
			args:     args{'9'},
			expected: '9',
		},
		{
			args:     args{'?'},
			expected: '?',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toupper(tt.args.theByte); got != tt.expected {
				t.Errorf("toupper() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
