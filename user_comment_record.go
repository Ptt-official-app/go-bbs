package bbs

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var (
	ErrNotUserComment          = fmt.Errorf("data is not a user comment")
	ErrUserCommentEmptyUserID  = fmt.Errorf("user comment has empty name")
	ErrUserCommentEmptyComment = fmt.Errorf("user comment detail has empty")
)

var (
	userCommentPattern = regexp.MustCompile(`([a-zA-Z][a-zA-Z0-9]+):.*([0-9][0-9]\/[0-9][0-9]\s[0-9][0-9]:[0-9][0-9])`)
)

type UserCommentRecord interface {
	CommentOrder() uint32
	CommentTime() time.Time
	Owner() string
	IP() string
	Comment() string
	String() string
	BoardID() string
	Filename() string
}

var _ UserCommentRecord = &userCommentRecord{}

type userCommentRecord struct {
	commentOrder uint32
	commentTime  time.Time
	owner        string
	ip           string
	boardID      string
	filename     string
	comment      string
}

// NewUserCommentRecord parses the data and returns the user comment record.
//  Return error when input data is not matched the user comment pattern.
func NewUserCommentRecord(order uint32, data string, boardID string, ar ArticleRecord) (UserCommentRecord, error) {
	owner, ctime, comment, err := parseUserComment(data)
	if err != nil {
		return nil, err
	}
	return &userCommentRecord{
		commentOrder: order,
		commentTime:  ctime,
		owner:        owner,
		ip:           "", // TODO
		boardID:      boardID,
		filename:     ar.Filename(),
		comment:      comment,
	}, nil
}

func (r userCommentRecord) CommentOrder() uint32 {
	return r.commentOrder
}

func (r userCommentRecord) CommentTime() time.Time {
	return r.commentTime
}

func (r userCommentRecord) Owner() string {
	return r.owner
}

func (r userCommentRecord) IP() string {
	return r.ip
}

func (r userCommentRecord) Comment() string {
	return r.comment
}

func (r userCommentRecord) String() string {
	return fmt.Sprintf("order: %d, owner: %s, time: %s", r.commentOrder, r.owner, r.commentTime.Format("01/02 15:04"))
}

func (r userCommentRecord) BoardID() string {
	return r.boardID
}

func (r userCommentRecord) Filename() string {
	return r.filename
}

// parseUserComment returns the owner and time of comment data.
//  Return ErrNotUserComment error when data doesn't match to the pattern.
//  Return other error when data contains the ambiguous value which can't parse.
func parseUserComment(data string) (owner string, ctime time.Time, comment string, err error) {
	matches := userCommentPattern.FindStringSubmatch(data)
	// The 1st record is entire matched result of row.
	// The 2nd record is group owner result, EX: "pichu".
	const ownerIdx = 1
	// The 3rd record is group time result,  EX: "05/15 01:06".
	const timeIdx = 2

	const commentIdx = 0

	if len(matches) < 3 {
		err = ErrNotUserComment
		return
	}

	owner = matches[ownerIdx]
	if len(owner) == 0 {
		err = ErrUserCommentEmptyUserID
		return
	}

	ctimeStr := matches[timeIdx]
	ctime, err = time.Parse("01/02 15:04", ctimeStr)
	if err != nil {
		return
	}

	commentStr := matches[commentIdx]
	if len(commentStr) == 0 {
		err = ErrUserCommentEmptyComment
		return
	}

	//TODO: improve get comment
	ownStr := ":"
	commentTimeRemoveArr := strings.Split(commentStr, ctimeStr)
	commentArr := strings.Split(commentTimeRemoveArr[0], ownStr)

	for key, value := range commentArr {
		//key 0 為使用者名稱，跳過
		if key == 0 {
			continue
		}
		comment += strings.TrimSpace(value)
		//key 2 開始才會有重複冒號的問題
		if key >= 2 {
			comment += ":"
		}
	}

	return owner, ctime, comment, nil
}
