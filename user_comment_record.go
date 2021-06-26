package bbs

import (
	"fmt"
	"regexp"
	"time"
)

var (
	ErrNotUserComment         = fmt.Errorf("data is not a user comment")
	ErrUserCommentEmptyUserID = fmt.Errorf("user comment has empty name")
)

var (
	userCommentPattern = regexp.MustCompile(`([a-zA-Z][a-zA-Z0-9]+):.*([0-9][0-9]\/[0-9][0-9]\s[0-9][0-9]:[0-9][0-9])`)
)

type UserCommentRecord interface {
	CommentOrder() uint32
	CommentTime() time.Time
	Owner() string
	IP() string
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
}

// NewUserCommentRecord parses the data and returns the user comment record.
//  Return error when input data is not matched the user comment pattern.
func NewUserCommentRecord(order uint32, data string, boardID string, ar ArticleRecord) (UserCommentRecord, error) {
	owner, ctime, err := parseUserComment(data)
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
func parseUserComment(data string) (owner string, ctime time.Time, err error) {
	matches := userCommentPattern.FindStringSubmatch(data)
	// The 1st record is entire matched result of row.
	// The 2nd record is group owner result, EX: "pichu".
	const ownerIdx = 1
	// The 3rd record is group time result,  EX: "05/15 01:06".
	const timeIdx = 2
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

	return owner, ctime, nil
}
