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

package pttbbs

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/Ptt-official-app/go-bbs"
)

var testBoardHeaders = []BoardHeader{
	{
		BrdName:            "SYSOP",
		title:              "嘰哩 ◎站長好!",
		bm:                 "",
		Brdattr:            BoardPostMask,
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
		title:              ".... Σ中央政府  《高壓危險,非人可敵》",
		bm:                 "",
		Brdattr:            BoardGroupBoard,
		VoteLimitPosts:     0,
		VoteLimitLogins:    0,
		ChessCountry:       "",
		BUpdate:            time.Unix(int64(0), 0),
		PostLimitPosts:     0,
		PostLimitLogins:    0,
		BVote:              0,
		VTime:              time.Unix(int64(0), 0),
		Level:              PermSYSOP,
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
		title:              "發電 ◎雜七雜八的垃圾",
		bm:                 "",
		Brdattr:            0,
		VoteLimitPosts:     0,
		VoteLimitLogins:    0,
		ChessCountry:       "",
		BUpdate:            time.Unix(int64(0), 0),
		PostLimitPosts:     0,
		PostLimitLogins:    0,
		BVote:              0,
		VTime:              time.Unix(int64(0), 0),
		Level:              PermSYSOP,
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
		title:              "發電 ◎站內系統安全",
		bm:                 "",
		Brdattr:            0,
		VoteLimitPosts:     0,
		VoteLimitLogins:    0,
		ChessCountry:       "",
		BUpdate:            time.Unix(int64(0), 0),
		PostLimitPosts:     0,
		PostLimitLogins:    0,
		BVote:              0,
		VTime:              time.Unix(int64(0), 0),
		Level:              PermSYSOP,
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
		title:              ".... Σ市民廣場     報告  站長  ㄜ！",
		bm:                 "",
		Brdattr:            BoardGroupBoard,
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
		BrdName:            "ALLPOST",
		title:              "嘰哩 ◎跨板式LOCAL新文章",
		bm:                 "",
		Brdattr:            BoardPostMask,
		VoteLimitPosts:     0,
		VoteLimitLogins:    0,
		ChessCountry:       "",
		BUpdate:            time.Unix(int64(0), 0),
		PostLimitPosts:     0,
		PostLimitLogins:    0,
		BVote:              0,
		VTime:              time.Unix(int64(0), 0),
		Level:              PermSYSOP,
		PermReload:         time.Unix(int64(0), 0),
		Gid:                5,
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
		BrdName:            "deleted",
		title:              "嘰哩 ◎資源回收筒",
		bm:                 "",
		Brdattr:            0,
		VoteLimitPosts:     0,
		VoteLimitLogins:    0,
		ChessCountry:       "",
		BUpdate:            time.Unix(int64(0), 0),
		PostLimitPosts:     0,
		PostLimitLogins:    0,
		BVote:              0,
		VTime:              time.Unix(int64(0), 0),
		Level:              PermBM,
		PermReload:         time.Unix(int64(0), 0),
		Gid:                5,
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
		BrdName:            "Note",
		title:              "嘰哩 ◎動態看板及歌曲投稿",
		bm:                 "",
		Brdattr:            0,
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
		Gid:                5,
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
		BrdName:            "Record",
		title:              "嘰哩 ◎我們的成果",
		bm:                 "",
		Brdattr:            0 | BoardPostMask,
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
		Gid:                5,
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
		BrdName:            "WhoAmI",
		title:              "嘰哩 ◎呵呵，猜猜我是誰！",
		bm:                 "",
		Brdattr:            0,
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
		Gid:                5,
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
		BrdName:            "EditExp",
		title:              "嘰哩 ◎範本精靈投稿區",
		bm:                 "",
		Brdattr:            0,
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
		Gid:                5,
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
		BrdName:            "ALLHIDPOST",
		title:              "嘰哩 ◎跨板式LOCAL新文章(隱板)",
		bm:                 "",
		Brdattr:            BoardPostMask | BoardHide,
		VoteLimitPosts:     0,
		VoteLimitLogins:    0,
		ChessCountry:       "",
		BUpdate:            time.Unix(int64(0), 0),
		PostLimitPosts:     0,
		PostLimitLogins:    0,
		BVote:              0,
		VTime:              time.Unix(int64(0), 0),
		Level:              PermSYSOP,
		PermReload:         time.Unix(int64(0), 0),
		Gid:                5,
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

var testBoardBinary = []string{
	`	
			5359 534f 5000 0000 0000 0000 00bc 54ad
f920 a1b7 afb8 aaf8 a66e 2100 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 2000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0200 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
`,
	`
312e 2e2e 2e2e 2e2e 2e2e 2e2e 002e 2e2e
2e20 a355 a4a4 a5a1 ac46 a9b2 2020 a16d
b0aa c0a3 a64d c049 2cab 44a4 48a5 69bc
c4a1 6e00 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0800 0000 0000 0000
0000 0000 0000 0000 0000 0000 0040 0000
0000 0000 0100 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
`,
}

func TestBoardHeader(t *testing.T) {
	headers, err := OpenBoardHeaderFile("testcase/board/01.BRD")
	if err != nil {
		t.Error(err)
	}
	expected := testBoardHeaders
	for index, header := range headers[0:11] {

		if header.BrdName != expected[index].BrdName {
			t.Errorf("BoardName not match in index %d, expected: %s, got: %s", index, expected[index].BrdName, header.BrdName)
		}
		if header.title != expected[index].title {
			t.Errorf("Title not match in index %d, expected: %s, got: %s", index, expected[index].title, header.title)
		}
		if header.bm != expected[index].bm {
			t.Errorf("BM not match in index %d, expected: %s, got: %s", index, expected[index].bm, header.bm)
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
		if header.Brdattr != expected[index].Brdattr {
			t.Errorf("Raw Brdattr not match in index %d, expected: %08X, got: %08X", index, expected[index].Brdattr, header.Brdattr)
		}

	}
}

func TestBoardHeaderInfo(t *testing.T) {
	headers, err := OpenBoardHeaderFile("testcase/board/01.BRD")
	if err != nil {
		t.Error(err)
	}

	expected := testBoardHeaders
	for index, header := range headers[0:11] {
		info := bbs.BoardRecordInfo(header)
		if info.GetPostLimitPosts() != expected[index].PostLimitPosts {
			t.Errorf("posts not match in index %d, expected: %d, got: %d", index, expected[index].PostLimitPosts, info.GetPostLimitPosts())
		}
		if info.GetPostLimitLogins() != expected[index].PostLimitLogins {
			t.Errorf("logins not match in index %d, expected: %d, got: %d", index, expected[index].PostLimitLogins, info.GetPostLimitLogins())
		}
		if info.GetPostLimitBadPost() != expected[index].PostLimitBadPost {
			t.Errorf("badpost not match in index %d, expected: %d, got: %d", index, expected[index].PostLimitBadPost, info.GetPostLimitBadPost())
		}
	}
}

func TestBoardHeaderSettings(t *testing.T) {
	headers, err := OpenBoardHeaderFile("testcase/board/01.BRD")
	if err != nil {
		t.Error(err)
	}

	expected := testBoardHeaders
	for index, header := range headers[0:11] {
		settings := bbs.BoardRecordSettings(header)
		if settings.IsHide() != expected[index].IsHide() {
			t.Errorf("IsHide not match in index %d, expected: %v, got: %v", index, expected[index].IsHide(), settings.IsHide())
		}
		if settings.IsPostMask() != expected[index].IsPostMask() {
			t.Errorf("IsPostMask not match in index %d, expected: %v, got: %v", index, expected[index].IsPostMask(), settings.IsPostMask())
		}
		if settings.IsAnonymous() != expected[index].IsAnonymous() {
			t.Errorf("IsAnonymous not match in index %d, expected: %v, got: %v", index, expected[index].IsAnonymous(), settings.IsAnonymous())
		}
		if settings.IsDefaultAnonymous() != expected[index].IsDefaultAnonymous() {
			t.Errorf("IsDefaultAnonymous not match in index %d, expected: %v, got: %v", index, expected[index].IsDefaultAnonymous(), settings.IsDefaultAnonymous())
		}
		if settings.IsNoCredit() != expected[index].IsNoCredit() {
			t.Errorf("IsNoCredit not match in index %d, expected: %v, got: %v", index, expected[index].IsNoCredit(), settings.IsNoCredit())
		}
		if settings.IsVoteBoard() != expected[index].IsVoteBoard() {
			t.Errorf("IsVoteBoard not match in index %d, expected: %v, got: %v", index, expected[index].IsVoteBoard(), settings.IsVoteBoard())
		}
		if settings.IsWarnEL() != expected[index].IsWarnEL() {
			t.Errorf("IsWarnEL not match in index %d, expected: %v, got: %v", index, expected[index].IsWarnEL(), settings.IsWarnEL())
		}
		if settings.IsTop() != expected[index].IsTop() {
			t.Errorf("IsTop not match in index %d, expected: %v, got: %v", index, expected[index].IsTop(), settings.IsTop())
		}
		if settings.IsNoRecommend() != expected[index].IsNoRecommend() {
			t.Errorf("IsNoRecommend not match in index %d, expected: %v, got: %v", index, expected[index].IsNoRecommend(), settings.IsNoRecommend())
		}
		if settings.IsAngelAnonymous() != expected[index].IsAngelAnonymous() {
			t.Errorf("IsAngelAnonymous not match in index %d, expected: %v, got: %v", index, expected[index].IsAngelAnonymous(), settings.IsAngelAnonymous())
		}
		if settings.IsBMCount() != expected[index].IsBMCount() {
			t.Errorf("IsBMCount not match in index %d, expected: %v, got: %v", index, expected[index].IsBMCount(), settings.IsBMCount())
		}
		if settings.IsNoBoo() != expected[index].IsNoBoo() {
			t.Errorf("IsNoBoo not match in index %d, expected: %v, got: %v", index, expected[index].IsNoBoo(), settings.IsNoBoo())
		}
		if settings.IsRestrictedPost() != expected[index].IsRestrictedPost() {
			t.Errorf("IsRestrictedPost not match in index %d, expected: %v, got: %v", index, expected[index].IsRestrictedPost(), settings.IsRestrictedPost())
		}
		if settings.IsGuestPost() != expected[index].IsGuestPost() {
			t.Errorf("IsGuestPost not match in index %d, expected: %v, got: %v", index, expected[index].IsGuestPost(), settings.IsGuestPost())
		}
		if settings.IsCooldown() != expected[index].IsCooldown() {
			t.Errorf("IsCooldown not match in index %d, expected: %v, got: %v", index, expected[index].IsCooldown(), settings.IsCooldown())
		}
		if settings.IsCPLog() != expected[index].IsCPLog() {
			t.Errorf("IsCPLog not match in index %d, expected: %v, got: %v", index, expected[index].IsCPLog(), settings.IsCPLog())
		}
		if settings.IsNoFastRecommend() != expected[index].IsNoFastRecommend() {
			t.Errorf("IsNoFastRecommend not match in index %d, expected: %v, got: %v", index, expected[index].IsNoFastRecommend(), settings.IsNoFastRecommend())
		}
		if settings.IsIPLogRecommend() != expected[index].IsIPLogRecommend() {
			t.Errorf("IsIPLogRecommend not match in index %d, expected: %v, got: %v", index, expected[index].IsIPLogRecommend(), settings.IsIPLogRecommend())
		}
		if settings.IsOver18() != expected[index].IsOver18() {
			t.Errorf("IsOver18 not match in index %d, expected: %v, got: %v", index, expected[index].IsOver18(), settings.IsOver18())
		}
		if settings.IsNoReply() != expected[index].IsNoReply() {
			t.Errorf("IsNoReply not match in index %d, expected: %v, got: %v", index, expected[index].IsNoReply(), settings.IsNoReply())
		}
		if settings.IsAlignedComment() != expected[index].IsAlignedComment() {
			t.Errorf("IsAlignedComment not match in index %d, expected: %v, got: %v", index, expected[index].IsAlignedComment(), settings.IsAlignedComment())
		}
		if settings.IsNoSelfDeletePost() != expected[index].IsNoSelfDeletePost() {
			t.Errorf("IsNoSelfDeletePost not match in index %d, expected: %v, got: %v", index, expected[index].IsNoSelfDeletePost(), settings.IsNoSelfDeletePost())
		}
		if settings.IsBMMaskContent() != expected[index].IsBMMaskContent() {
			t.Errorf("IsBMMaskContent not match in index %d, expected: %v, got: %v", index, expected[index].IsBMMaskContent(), settings.IsBMMaskContent())
		}
	}
}

func TestAppendAndRemoveBoardRecord(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "board_test_*")
	if err != nil {
		t.Errorf("create tmp file error: %v", err)
	}
	t.Logf("tmpfile: %v", tmpfile.Name())
	// t.Errorf("create tmp file error: %v", err)
	filename := tmpfile.Name()
	expectedBrdName := "XXXSSS"
	brd := BoardHeader{
		BrdName: expectedBrdName,
	}
	err = AppendBoardHeaderFileRecord(filename, &brd)
	if err != nil {
		t.Errorf("AppendBoardHeaderFileRecord error: %v", err)
	}

	headers, err := OpenBoardHeaderFile(filename)
	if err != nil {
		t.Error(err)
	}

	if len(headers) != 1 {
		t.Errorf("AppendBoardHeaderFileRecord failed, len(headers) expected: 1, got %v", len(headers))
	}

	if headers[0].BrdName != expectedBrdName {
		t.Errorf("AppendBoardHeaderFileRecord BrdName not match, expected: %v, got %v", expectedBrdName, headers[0].BrdName)
	}

	err = RemoveBoardHeaderFileRecord(filename, 0)
	if err != nil {
		t.Errorf("RemoveBoardHeaderFileRecord error: %v", err)
	}

	headers, err = OpenBoardHeaderFile(filename)
	if err != nil {
		t.Error(err)
	}

	if len(headers) != 0 {
		t.Errorf("RemoveBoardHeaderFileRecord failed, len(headers) expected: 0, got %v", len(headers))
	}

	defer os.Remove(tmpfile.Name()) // clean up
}

func TestMarshalBinary(t *testing.T) {
	type TestCase struct {
		Input    BoardHeader
		Expected []byte
	}
	// Only test first 2 cases since the byte split issue
	// in BrdName and title
	testcase := make([]TestCase, len(testBoardBinary))
	for i := 0; i < len(testcase); i++ {
		testcase[i] = TestCase{
			Input:    testBoardHeaders[i],
			Expected: hexToByte(testBoardBinary[i]),
		}
	}

	for index, c := range testcase {
		b, err := c.Input.MarshalBinary()
		t.Logf("log: %q, %q", b, err)
		if hex.Dump(b) != hex.Dump(c.Expected) {
			t.Errorf("Expected byte not match in index %d, expected: \n%s\n, got: \n%s", index, hex.Dump(c.Expected), hex.Dump(b))
		}

	}
}
