package bbs

import (
	"strings"
	"testing"
	"time"
)

func TestNewUserCommentRecord(t *testing.T) {
	perfectData := "→ lex: 快一點                        05/15 01:06"
	expectedOrder := uint32(1)
	expectedOwner := "lex"
	expectedTime := time.Date(0, 5, 15, 1, 6, 0, 0, time.UTC)
	expectedBoardID := "SYSOP"
	expectedArticleRecord := &MockArticleRecord{}

	got, err := NewUserCommentRecord(1, perfectData, expectedBoardID, expectedArticleRecord)

	if got.CommentOrder() != expectedOrder {
		t.Errorf("comment order = %v, expected %v\n", got.CommentOrder(), expectedOrder)
	}

	if strings.Compare(got.Owner(), expectedOwner) != 0 {
		t.Errorf("comment owner = %v, expected %v\n", got.Owner(), expectedOwner)
	}

	if !got.CommentTime().Equal(expectedTime) {
		t.Errorf("comment time = %v, expected %v\n", got.CommentTime(), expectedTime)
	}

	if got.BoardID() != expectedBoardID {
		t.Errorf("boardID = %v, expected %v\n", got.BoardID(), expectedBoardID)
	}

	if got.Filename() != expectedArticleRecord.Filename() {
		t.Errorf("boardID = %v, expected %v\n", got.Filename(), expectedArticleRecord.Filename())
	}

	if err != nil {
		t.Errorf("err = %v, should be nil\n", err)
	}
}

func TestParseUserComment(t *testing.T) {

	expectedTime := time.Date(0, 5, 15, 1, 6, 0, 0, time.UTC)
	emptyTime := time.Time{}

	perfectData := "→ lex: 快一點                        05/15 01:06"
	dataWithoutOwner := ": 快一點                             05/15 01:06"
	dataWithInvalidTime := "→ lex: 快一點                        5/15 01:06"
	emptyData := ""

	type args struct {
		data string
	}
	type expected struct {
		owner  string
		ctime  time.Time
		hasErr bool
	}
	tests := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name:     "parse perfect data",
			args:     args{perfectData},
			expected: expected{"lex", expectedTime, false},
		},
		{
			name:     "parse data without owner should return error",
			args:     args{dataWithoutOwner},
			expected: expected{"", emptyTime, true},
		},
		{
			name:     "parse data with invalid time should return error",
			args:     args{dataWithInvalidTime},
			expected: expected{"", emptyTime, true},
		},
		{
			name:     "parse empty data should return error",
			args:     args{emptyData},
			expected: expected{"", emptyTime, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOwner, gotTime, gotErr := parseUserComment(tt.args.data)

			if strings.Compare(gotOwner, tt.expected.owner) != 0 {
				t.Errorf("comment owner = %v, expected %v\n", gotOwner, tt.expected.owner)
			}

			if !gotTime.Equal(tt.expected.ctime) {
				t.Errorf("comment time = %v, expected %v\n", gotTime, tt.expected.ctime)
			}

			if (gotErr != nil) != tt.expected.hasErr {
				t.Errorf("err = %v, should be nil\n", gotErr)
			}
		})
	}
}
