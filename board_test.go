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
		{
			BrdName:            "1...........",
			Title:              ".... Σ中央政府  《高壓危險,非人可敵》",
			BM:                 "",
			Brdattr:            PTT_BRD_GROUPBOARD,
			VoteLimitPosts:     0,
			VoteLimitLogins:    0,
			ChessCountry:       "",
			BUpdate:            time.Unix(int64(0), 0),
			PostLimitPosts:     0,
			PostLimitLogins:    0,
			BVote:              0,
			VTime:              time.Unix(int64(0), 0),
			Level:              PTT_PERM_SYSOP,
			PermReload:         time.Unix(int64(0), 0),
			Gid:                1,
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
		{
			BrdName:            "junk",
			Title:              "發電 ◎雜七雜八的垃圾",
			BM:                 "",
			Brdattr:            0,
			VoteLimitPosts:     0,
			VoteLimitLogins:    0,
			ChessCountry:       "",
			BUpdate:            time.Unix(int64(0), 0),
			PostLimitPosts:     0,
			PostLimitLogins:    0,
			BVote:              0,
			VTime:              time.Unix(int64(0), 0),
			Level:              PTT_PERM_SYSOP,
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
		{
			BrdName:            "Security",
			Title:              "發電 ◎站內系統安全",
			BM:                 "",
			Brdattr:            0,
			VoteLimitPosts:     0,
			VoteLimitLogins:    0,
			ChessCountry:       "",
			BUpdate:            time.Unix(int64(0), 0),
			PostLimitPosts:     0,
			PostLimitLogins:    0,
			BVote:              0,
			VTime:              time.Unix(int64(0), 0),
			Level:              PTT_PERM_SYSOP,
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
		{
			BrdName:            "2...........",
			Title:              ".... Σ市民廣場     報告  站長  ㄜ！",
			BM:                 "",
			Brdattr:            PTT_BRD_GROUPBOARD,
			VoteLimitPosts:     0,
			VoteLimitLogins:    0,
			ChessCountry:       "",
			BUpdate:            time.Unix(int64(0), 0),
			PostLimitPosts:     0,
			PostLimitLogins:    0,
			BVote:              0,
			VTime:              time.Unix(int64(0), 0),
			Level:              PTT_PERM_SYSOP,
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

	for index, header := range headers[0:2] {

		if header.BrdName != expected[index].BrdName {
			t.Errorf("BoardName not match in index %d, expected: %s, got: %s", index, expected[index].BrdName, header.BrdName)
		}
		if header.Title != expected[index].Title {
			t.Errorf("Title not match in index %d, expected: %s, got: %s", index, expected[index].Title, header.Title)
		}
		if header.BM != expected[index].BM {
			t.Errorf("BM not match in index %d, expected: %s, got: %s", index, expected[index].BM, header.BM)
		}
		if header.VoteLimitPosts != expected[index].VoteLimitPosts {
			t.Errorf("VoteLimitPosts not match in index %d, expected: %d, got: %d", index, expected[index].VoteLimitPosts, header.VoteLimitPosts)
		}
		if header.VoteLimitLogins != expected[index].VoteLimitLogins {
			t.Errorf("VoteLimitLogins not match in index %d, expected: %d, got: %d", index, expected[index].VoteLimitLogins, header.VoteLimitLogins)
		}
		if header.ChessCountry != expected[index].ChessCountry {
			t.Errorf("ChessCountry not match in index %d, expected: %s, got: %s", index, expected[index].ChessCountry, header.ChessCountry)
		}
		if header.BUpdate != expected[index].BUpdate {
			t.Errorf("BUpdate not match in index %d, expected: %s, got: %s", index, expected[index].BUpdate, header.BUpdate)
		}
		if header.PostLimitPosts != expected[index].PostLimitPosts {
			t.Errorf("PostLimitPosts not match in index %d, expected: %d, got: %d", index, expected[index].PostLimitPosts, header.PostLimitPosts)
		}
		if header.BVote != expected[index].BVote {
			t.Errorf("BVote not match in index %d, expected: %d, got: %d", index, expected[index].BVote, header.BVote)
		}
		if header.VTime != expected[index].VTime {
			t.Errorf("VTime not match in index %d, expected: %s, got: %s", index, expected[index].VTime, header.VTime)
		}
		if header.Level != expected[index].Level {
			t.Errorf("Level not match in index %d, expected: %d, got: %d", index, expected[index].Level, header.Level)
		}
		if header.PermReload != expected[index].PermReload {
			t.Errorf("PermReload not match in index %d, expected: %s, got: %s", index, expected[index].PermReload, header.PermReload)
		}
		if header.Gid != expected[index].Gid {
			t.Errorf("Gid not match in index %d, expected: %d, got: %d", index, expected[index].Gid, header.Gid)
		}
		for i := 0; i < 2; i++ {
			if header.Next[i] != expected[index].Next[i] {
				t.Errorf("Nextnot match in index %d, expected: %d, got: %d", index, expected[index].Next[i], header.Next[i])
			}
		}
		for i := 0; i < 2; i++ {
			if header.FirstChild[i] != expected[index].FirstChild[i] {
				t.Errorf("FirstChild not match in index %d, expected: %d, got: %d", index, expected[index].FirstChild[i], header.FirstChild[i])
			}
		}
		if header.Parent != expected[index].Parent {
			t.Errorf("Parent not match in index %d, expected: %d, got: %d", index, expected[index].Parent, header.Parent)
		}
		if header.ChildCount != expected[index].ChildCount {
			t.Errorf("ChildCount not match in index %d, expected: %d, got: %d", index, expected[index].ChildCount, header.ChildCount)
		}
		if header.Nuser != expected[index].Nuser {
			t.Errorf("Nuser not match in index %d, expected: %d, got: %d", index, expected[index].Nuser, header.Nuser)
		}
		if header.PostExpire != expected[index].PostExpire {
			t.Errorf("PostExpire not match in index %d, expected: %d, got: %d", index, expected[index].PostExpire, header.PostExpire)
		}
		if header.EndGamble != expected[index].EndGamble {
			t.Errorf("EndGamble not match in index %d, expected: %s, got: %s", index, expected[index].EndGamble, header.EndGamble)
		}
		if header.PostType != expected[index].PostType {
			t.Errorf("PostType not match in index %d, expected: %s, got: %s", index, expected[index].PostType, header.PostType)
		}
		if header.PostTypeF != expected[index].PostTypeF {
			t.Errorf("PostTypeF not match in index %d, expected: %s, got: %s", index, expected[index].PostTypeF, header.PostTypeF)
		}
		if header.FastRecommendPause != expected[index].FastRecommendPause {
			t.Errorf("FastRecommendPause not match in index %d, expected: %d, got: %d", index, expected[index].FastRecommendPause, header.FastRecommendPause)
		}
		if header.VoteLimitBadPost != expected[index].VoteLimitBadPost {
			t.Errorf("VoteLimitBadPost not match in index %d, expected: %d, got: %d", index, expected[index].VoteLimitBadPost, header.VoteLimitBadPost)
		}
		if header.PostLimitBadPost != expected[index].PostLimitBadPost {
			t.Errorf("PostLimitBadPost not match in index %d, expected: %d, got: %d", index, expected[index].PostLimitBadPost, header.PostLimitBadPost)
		}
		if header.SRexpire != expected[index].SRexpire {
			t.Errorf("SRexpire not match in index %d, expected: %s, got: %s", index, expected[index].SRexpire, header.SRexpire)
		}

	}

}
