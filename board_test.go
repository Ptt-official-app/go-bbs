package bbs

import (
	"strings"
	"testing"
)

func TestParseBoardHeader(t *testing.T) {

	headers, err := OpenBoardHeaderFile("testcase/board/01.BRD")
	if err != nil {
		t.Error(err)
	}

	expected := []BoardHeader{
		{
			Boardname:       "SYSOP",
			Title:           "嘰哩 ◎站長好!",
			BoardModerators: []string{""},
		},
		{
			Boardname:       "1...........",
			Title:           ".... Σ中央政府  《高壓危險,非人可敵》",
			BoardModerators: []string{""},
		},
		{
			Boardname:       "junk",
			Title:           "發電 ◎雜七雜八的垃圾",
			BoardModerators: []string{""},
		},
		{
			Boardname:       "Security",
			Title:           "發電 ◎站內系統安全",
			BoardModerators: []string{""},
		},
		{
			Boardname:       "2...........",
			Title:           ".... Σ市民廣場     報告  站長  ㄜ！",
			BoardModerators: []string{""},
		},
		{
			Boardname:       "ALLPOST",
			Title:           "嘰哩 ◎跨板式LOCAL新文章",
			BoardModerators: []string{""},
		},
		{
			Boardname:       "deleted",
			Title:           "嘰哩 ◎資源回收筒",
			BoardModerators: []string{""},
		},
		{
			Boardname:       "Note",
			Title:           "嘰哩 ◎動態看板及歌曲投稿",
			BoardModerators: []string{""},
		},
		{
			Boardname:       "Record",
			Title:           "嘰哩 ◎我們的成果",
			BoardModerators: []string{""},
		},
		{
			Boardname:       "WhoAmI",
			Title:           "嘰哩 ◎呵呵，猜猜我是誰！",
			BoardModerators: []string{""},
		},
		{
			Boardname:       "EditExp",
			Title:           "嘰哩 ◎範本精靈投稿區",
			BoardModerators: []string{""},
		},
		{
			Boardname:       "ALLHIDPOST",
			Title:           "嘰哩 ◎跨板式LOCAL新文章(隱板)",
			BoardModerators: []string{""},
		},
		{
			Boardname:       "ptt_app",
			Title:           "測試 ◎大家來玩吧",
			BoardModerators: []string{"SYSOP"},
		},
		{
			Boardname:       "test",
			Title:           "測試 ◎測試測試測試",
			BoardModerators: []string{"SYSOP"},
		},
	}

	for index, header := range headers {
		if header.Boardname != expected[index].Boardname {
			t.Logf("len :%d %d", len(header.Boardname), len(expected[index].Boardname))
			t.Errorf("Boardname not match in index %d, expected: %s, got: %s", index, expected[index].Boardname, header.Boardname)
		}
		if header.Title != expected[index].Title {
			t.Logf("len :%d %d", len(header.Title), len(expected[index].Title))
			t.Errorf("Title not match in index %d, expected: %s, got: %s", index, expected[index].Title, header.Title)
		}
		if strings.Join(header.BoardModerators, "}{") != strings.Join(expected[index].BoardModerators, "}{") {
			t.Logf("len :%d %d", len(header.BoardModerators), len(expected[index].BoardModerators))
			t.Errorf("BoardModerators not match in index %d, expected: %s, got: %s", index, expected[index].BoardModerators, header.BoardModerators)
		}

	}

}
