package bbs

import (
	"time"
)

// VoteLimits shows the limitation of a vote post.
type VoteLimits struct {
	Posts   uint8
	Logins  uint8
	Regtime uint8
	Badpost uint8
}

// FileHeader records article's metainfo
type FileHeader struct {
	Filename  string
	Modified  time.Time
	Recommend int8   // Important Level
	Owner     string // uid[.]
	Date      string
	Title     string

	Money   int
	AnnoUid int
	VoteLimits
	ReferRef  uint // 至底公告？
	ReferFlag bool // 至底公告？

	Filemode uint8

	Postno int32 // FormosaBBS only
}
