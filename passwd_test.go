package bbs

import (
	"testing"
)

func TestOpenUserecFile(t *testing.T) {

	actualUserecs, err := OpenUserecFile("testcase/passwd/01.PASSWDS")
	if err != nil {
		t.Errorf("OpenUserecFile() error = %v", err)
		return
	}

	expected := []Userec{
		Userec{
			Version:      4194,
			UserId:       "SYSOP",
			RealName:     "CodingMan",
			Nickname:     "神",
			Password:     "bhwvOJtfT1TAI",
			UserFlag:     0x02000A60,
			UserLevel:    0x20000407,
			NumLoginDays: 2,
			NumPosts:     0,
			FirstLogin:   1600681288,
			LastLogin:    1600756094,
			LastHost:     "59.124.167.226",
			Money:        0,
		},
		Userec{
			Version:      4194,
			UserId:       "CodingMan",
			RealName:     "朱元璋",
			Nickname:     "程式俠",
			Password:     "u8mLG.ktfOk3w",
			UserFlag:     0x02000AE0,
			UserLevel:    0x0000001F,
			NumLoginDays: 1,
			NumPosts:     0,
			FirstLogin:   1600737659,
			LastLogin:    1600737960,
			LastHost:     "59.124.167.226",
		},
		Userec{
			Version:      4194,
			UserId:       "pichu",
			RealName:     "Pichu",
			Nickname:     "Pichu",
			Password:     "KO27TyME.3/tw",
			UserFlag:     0x02000AE0,
			UserLevel:    0x00000007,
			NumLoginDays: 1,
			NumPosts:     0,
			FirstLogin:   1600755675,
			LastLogin:    1600766204,
			LastHost:     "103.246.218.43",
		},
		Userec{
			Version:      4194,
			UserId:       "Kahou",
			RealName:     "林嘉豪",
			Nickname:     "Kahou",
			Password:     "V3nkaYTLnDPUA",
			UserFlag:     0x02000AE0,
			UserLevel:    0x00000007,
			NumLoginDays: 1,
			NumPosts:     0,
			FirstLogin:   1600758266,
			LastLogin:    1600758266,
			LastHost:     "180.217.174.18",
		},
		Userec{
			Version:      4194,
			UserId:       "Kahou2",
			RealName:     "Kahou",
			Nickname:     "kahou",
			Password:     "R7shIAOZgQCKs",
			UserFlag:     0x02000AE0,
			UserLevel:    0x0000001F,
			NumLoginDays: 1,
			NumPosts:     0,
			FirstLogin:   1600758939,
			LastLogin:    1600760401,
			LastHost:     "180.217.174.18",
		},
	}

	for index, actual := range actualUserecs {

		ok := index < len(expected)
		if !ok {
			t.Logf("expected case %d not exist, assume all 0", index)
			break
		}
		expected := expected[index]

		if actual.Version != expected.Version {
			t.Errorf("Version not match with index %d, expected: %v, got: %v", index, expected.Version, actual.Version)
		}

		if actual.UserId != expected.UserId {
			t.Errorf("UserId not match with index %d, expected: %v, got: %v", index, expected.UserId, actual.UserId)
		}

		if actual.RealName != expected.RealName {
			t.Errorf("RealName not match with index %d, expected: %v, got: %v", index, expected.RealName, actual.RealName)
		}

		if actual.Nickname != expected.Nickname {
			t.Errorf("Nickname not match with index %d, expected: %v, got: %v", index, expected.Nickname, actual.Nickname)
		}

		if actual.Password != expected.Password {
			t.Errorf("Password not match with index %d, expected: %v, got: %v", index, expected.Password, actual.Password)
		}

		if actual.UserFlag != expected.UserFlag {
			t.Errorf("UserFlag not match with index %d, expected: 0x%08X, got: 0x%08X", index, expected.UserFlag, actual.UserFlag)
		}

		if actual.UserLevel != expected.UserLevel {
			t.Errorf("UserLevel not match with index %d, expected: 0x%08X, got: 0x%08X", index, expected.UserLevel, actual.UserLevel)
		}

		if actual.NumLoginDays != expected.NumLoginDays {
			t.Errorf("NumLoginDays not match with index %d, expected: %v, got: %v", index, expected.NumLoginDays, actual.NumLoginDays)
		}

		if actual.NumPosts != expected.NumPosts {
			t.Errorf("NumPosts not match with index %d, expected: %v, got: %v", index, expected.NumPosts, actual.NumPosts)
		}

		if actual.FirstLogin != expected.FirstLogin {
			t.Errorf("FirstLogin not match with index %d, expected: %v, got: %v", index, expected.FirstLogin, actual.FirstLogin)
		}

		if actual.LastLogin != expected.LastLogin {
			t.Errorf("LastLogin not match with index %d, expected: %v, got: %v", index, expected.LastLogin, actual.LastLogin)
		}

		if actual.LastHost != expected.LastHost {
			t.Errorf("LastHost not match with index %d, expected: %v, got: %v", index, expected.LastHost, actual.LastHost)
		}

		if actual.Money != expected.Money {
			t.Errorf("Money not match with index %d, expected: %v, got: %v", index, expected.Money, actual.Money)
		}

		if actual.Email != expected.Email {
			t.Errorf("Email not match with index %d, expected: %v, got: %v", index, expected.Email, actual.Email)
		}

		if actual.Address != expected.Address {
			t.Errorf("Address not match with index %d, expected: %v, got: %v", index, expected.Address, actual.Address)
		}

		if actual.Justify != expected.Justify {
			t.Errorf("Justify not match with index %d, expected: %v, got: %v", index, expected.Justify, actual.Justify)
		}

		if actual.Over18 != expected.Over18 {
			t.Errorf("Over18 not match with index %d, expected: %v, got: %v", index, expected.Over18, actual.Over18)
		}

		if actual.PagerUIType != expected.PagerUIType {
			t.Errorf("PagerUIType not match with index %d, expected: %v, got: %v", index, expected.PagerUIType, actual.PagerUIType)
		}

		if actual.Pager != expected.Pager {
			t.Errorf("Pager not match with index %d, expected: %v, got: %v", index, expected.Pager, actual.Pager)
		}

		if actual.Invisible != expected.Invisible {
			t.Errorf("Invisible not match with index %d, expected: %v, got: %v", index, expected.Invisible, actual.Invisible)
		}

		if actual.ExMailBox != expected.ExMailBox {
			t.Errorf("ExMailBox not match with index %d, expected: %v, got: %v", index, expected.ExMailBox, actual.ExMailBox)
		}

		if actual.Career != expected.Career {
			t.Errorf("Career not match with index %d, expected: %v, got: %v", index, expected.Career, actual.Career)
		}

		if actual.Role != expected.Role {
			t.Errorf("Role not match with index %d, expected: %v, got: %v", index, expected.Role, actual.Role)
		}

		if actual.LastSeen != expected.LastSeen {
			t.Errorf("LastSeen not match with index %d, expected: %v, got: %v", index, expected.LastSeen, actual.LastSeen)
		}

		if actual.TimeSetAngel != expected.TimeSetAngel {
			t.Errorf("TimeSetAngel not match with index %d, expected: %v, got: %v", index, expected.TimeSetAngel, actual.TimeSetAngel)
		}

		if actual.TimePlayAngel != expected.TimePlayAngel {
			t.Errorf("TimePlayAngel not match with index %d, expected: %v, got: %v", index, expected.TimePlayAngel, actual.TimePlayAngel)
		}

		if actual.LastSong != expected.LastSong {
			t.Errorf("LastSong not match with index %d, expected: %v, got: %v", index, expected.LastSong, actual.LastSong)
		}

		if actual.LoginView != expected.LoginView {
			t.Errorf("LoginView not match with index %d, expected: %v, got: %v", index, expected.LoginView, actual.LoginView)
		}

		if actual.ViolateLaw != expected.ViolateLaw {
			t.Errorf("ViolateLaw not match with index %d, expected: %v, got: %v", index, expected.ViolateLaw, actual.ViolateLaw)
		}

		if actual.FiveWin != expected.FiveWin {
			t.Errorf("FiveWin not match with index %d, expected: %v, got: %v", index, expected.FiveWin, actual.FiveWin)
		}

		if actual.FiveLose != expected.FiveLose {
			t.Errorf("FiveLose not match with index %d, expected: %v, got: %v", index, expected.FiveLose, actual.FiveLose)
		}

		if actual.FiveTie != expected.FiveTie {
			t.Errorf("FiveTie not match with index %d, expected: %v, got: %v", index, expected.FiveTie, actual.FiveTie)
		}

		if actual.ChcWin != expected.ChcWin {
			t.Errorf("ChcWin not match with index %d, expected: %v, got: %v", index, expected.ChcWin, actual.ChcWin)
		}

		if actual.ChcLose != expected.ChcLose {
			t.Errorf("ChcLose not match with index %d, expected: %v, got: %v", index, expected.ChcLose, actual.ChcLose)
		}

		if actual.ChcTie != expected.ChcTie {
			t.Errorf("ChcTie not match with index %d, expected: %v, got: %v", index, expected.ChcTie, actual.ChcTie)
		}

		if actual.Conn6Win != expected.Conn6Win {
			t.Errorf("Conn6Win not match with index %d, expected: %v, got: %v", index, expected.Conn6Win, actual.Conn6Win)
		}

		if actual.Conn6Lose != expected.Conn6Lose {
			t.Errorf("Conn6Lose not match with index %d, expected: %v, got: %v", index, expected.Conn6Lose, actual.Conn6Lose)
		}

		if actual.Conn6Tie != expected.Conn6Tie {
			t.Errorf("Conn6Tie not match with index %d, expected: %v, got: %v", index, expected.Conn6Tie, actual.Conn6Tie)
		}

		if actual.GoWin != expected.GoWin {
			t.Errorf("GoWin not match with index %d, expected: %v, got: %v", index, expected.GoWin, actual.GoWin)
		}

		if actual.GoLose != expected.GoLose {
			t.Errorf("GoLose not match with index %d, expected: %v, got: %v", index, expected.GoLose, actual.GoLose)
		}

		if actual.GoTie != expected.GoTie {
			t.Errorf("GoTie not match with index %d, expected: %v, got: %v", index, expected.GoTie, actual.GoTie)
		}

		if actual.DarkWin != expected.DarkWin {
			t.Errorf("DarkWin not match with index %d, expected: %v, got: %v", index, expected.DarkWin, actual.DarkWin)
		}

		if actual.DarkLose != expected.DarkLose {
			t.Errorf("DarkLose not match with index %d, expected: %v, got: %v", index, expected.DarkLose, actual.DarkLose)
		}

		if actual.UaVersion != expected.UaVersion {
			t.Errorf("UaVersion not match with index %d, expected: %v, got: %v", index, expected.UaVersion, actual.UaVersion)
		}

		if actual.Signature != expected.Signature {
			t.Errorf("Signature not match with index %d, expected: %v, got: %v", index, expected.Signature, actual.Signature)
		}

		if actual.BadPost != expected.BadPost {
			t.Errorf("BadPost not match with index %d, expected: %v, got: %v", index, expected.BadPost, actual.BadPost)
		}

		if actual.DarkTie != expected.DarkTie {
			t.Errorf("DarkTie not match with index %d, expected: %v, got: %v", index, expected.DarkTie, actual.DarkTie)
		}

		if actual.MyAngel != expected.MyAngel {
			t.Errorf("MyAngel not match with index %d, expected: %v, got: %v", index, expected.MyAngel, actual.MyAngel)
		}

		if actual.ChessEloRating != expected.ChessEloRating {
			t.Errorf("ChessEloRating not match with index %d, expected: %v, got: %v", index, expected.ChessEloRating, actual.ChessEloRating)
		}

		if actual.WithMe != expected.WithMe {
			t.Errorf("WithMe not match with index %d, expected: %v, got: %v", index, expected.WithMe, actual.WithMe)
		}

		if actual.TimeRemoveBadPost != expected.TimeRemoveBadPost {
			t.Errorf("TimeRemoveBadPost not match with index %d, expected: %v, got: %v", index, expected.TimeRemoveBadPost, actual.TimeRemoveBadPost)
		}

		if actual.TimeViolateLaw != expected.TimeViolateLaw {
			t.Errorf("TimeViolateLaw not match with index %d, expected: %v, got: %v", index, expected.TimeViolateLaw, actual.TimeViolateLaw)
		}

	}

}
