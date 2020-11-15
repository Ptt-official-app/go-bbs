package bbs

import (
	"testing"
	"time"
)

func TestSomething(t *testing.T) {
	headers, err := OpenBoardHeaderFile("testcase/board/01.BRD")
	if err != nil {
		t.Error(err)
	}

	expected := []BoardHeader{
		{
			BrdName:            "SYSOP",
			Title:              "嘰哩 ◎站長好!",
			BM:                 "",
			Brdattr:            PTT_BRD_POSTMASK,
			VoteLimitPosts:     0,
			VoteLimitLogins:    0,
			ChessCountry:       "",
			BUpdate:            time.Unix(int64(0), 0),
			PostLimitPosts:     0,
			PostLimitLogins:    0,
			BVote:              0,
			VTime:              time.Unix(int64(0), 0),
			Level:              0,
			PermReload:         time.Unix(int64(0), 0),
			Gid:                2,
			Next:               []int32{0, 0},
			FirstChild:         []int32{0, 0},
			Parent:             0,
			ChildCount:         0,
			Nuser:              0,
			PostExpire:         0,
			EndGamble:          time.Unix(int64(0), 0),
			PostType:           "",
			PostTypeF:          "",
			FastRecommendPause: 0,
			VoteLimitBadPost:   0,
			PostLimitBadPost:   0,
			SRexpire:           time.Unix(int64(0), 0),
		},
	}

	for index, header := range headers {
		if header.BrdName != expected[index].BrdName {
			t.Errorf("Board name not match in index %d, expected: %s, got: %s", index, expected[index].BrdName, header.BrdName)
		}
	}

}
