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
	"testing"
	"time"
)

func TestOpenUserecFile(t *testing.T) {

	actualUserecs, err := OpenUserecFile("testcase/passwd/01.PASSWDS")
	if err != nil {
		t.Errorf("OpenUserecFile() error = %v", err)
		return
	}
	expected := []Userec{
		{
			Version:       4194,
			userID:        "SYSOP",
			realName:      "CodingMan",
			nickname:      "神",
			password:      "bhwvOJtfT1TAI",
			UserFlag:      0x02000A60,
			UserLevel:     0x20000407,
			numLoginDays:  2,
			numPosts:      0,
			firstLogin:    time.Date(2020, 9, 21, 9, 41, 28, 0, time.UTC),
			lastLogin:     time.Date(2020, 9, 22, 6, 28, 14, 0, time.UTC),
			lastHost:      "59.124.167.226",
			money:         0,
			Address:       "新竹縣子虛鄉烏有村543號",
			Over18:        true,
			Pager:         1,
			Invisible:     false,
			Career:        "全景軟體",
			LastSeen:      time.Date(2020, 9, 21, 9, 41, 28, 0, time.UTC),
			TimeSetAngel:  time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			TimePlayAngel: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			LastSong:      time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),

			TimeRemoveBadPost: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			TimeViolateLaw:    time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			BadPost:           0x01,
			ViolateLaw:        0x01,
			Five: UserecGameScore{
				Win:  0x0001,
				Lose: 0x0002,
				Tie:  0x0003,
			},
			ChineseChess: UserecGameScore{
				Win:  0x0004,
				Lose: 0x0005,
				Tie:  0x0006,
			},
			Conn6: UserecGameScore{
				Win:  0x0007,
				Lose: 0x0008,
				Tie:  0x0009,
			},
			GoChess: UserecGameScore{
				Win:  0x000a,
				Lose: 0x000b,
				Tie:  0x000c,
			},
			DarkChess: UserecGameScore{
				Win:  0x000d,
				Lose: 0x000e,
				Tie:  0x000f,
			},
			MyAngel: "SYSOP",
		},
		{
			Version:       4194,
			userID:        "CodingMan",
			realName:      "朱元璋",
			nickname:      "程式俠",
			password:      "u8mLG.ktfOk3w",
			UserFlag:      0x02000AE0,
			UserLevel:     0x0000001F,
			numLoginDays:  1,
			numPosts:      0,
			firstLogin:    time.Date(2020, 9, 22, 1, 20, 59, 0, time.UTC),
			lastLogin:     time.Date(2020, 9, 22, 1, 26, 00, 0, time.UTC),
			lastHost:      "59.124.167.226",
			Email:         "x",
			Address:       "新竹縣子虛鄉烏有村543號",
			Justify:       "[SYSOP] 09/22/2020 01:25:53 Tue",
			Over18:        true,
			Pager:         1,
			Career:        "全景軟體",
			LastSeen:      time.Date(2020, 9, 22, 1, 26, 00, 0, time.UTC),
			TimeSetAngel:  time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			TimePlayAngel: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			LastSong:      time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),

			TimeRemoveBadPost: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			TimeViolateLaw:    time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			BadPost:           0x02,
			ViolateLaw:        0x02,
			Five: UserecGameScore{
				Win:  0x0101,
				Lose: 0x0102,
				Tie:  0x0103,
			},
			ChineseChess: UserecGameScore{
				Win:  0x0104,
				Lose: 0x0105,
				Tie:  0x0106,
			},
			Conn6: UserecGameScore{
				Win:  0x0107,
				Lose: 0x0108,
				Tie:  0x0109,
			},
			GoChess: UserecGameScore{
				Win:  0x010a,
				Lose: 0x010b,
				Tie:  0x010c,
			},
			DarkChess: UserecGameScore{
				Win:  0x010d,
				Lose: 0x010e,
				Tie:  0x010f,
			},
			MyAngel: "CodingMan",
		},
		{
			Version:       4194,
			userID:        "pichu",
			realName:      "Pichu",
			nickname:      "Pichu",
			password:      "KO27TyME.3/tw",
			UserFlag:      0x02000AE0,
			UserLevel:     0x00000007,
			numLoginDays:  1,
			numPosts:      0,
			firstLogin:    time.Date(2020, 9, 22, 6, 21, 15, 0, time.UTC),
			lastLogin:     time.Date(2020, 9, 22, 9, 16, 44, 0, time.UTC),
			lastHost:      "103.246.218.43",
			Career:        "台灣智慧家庭",
			LastSeen:      time.Date(2020, 9, 22, 9, 16, 44, 0, time.UTC),
			TimeSetAngel:  time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			TimePlayAngel: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			LastSong:      time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			Email:         "pichu@tih.tw",
			Address:       "北市蘆洲區123路3號",
			Justify:       "<Email>",
			Over18:        true,
			Pager:         1,

			TimeRemoveBadPost: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			TimeViolateLaw:    time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			BadPost:           0x03,
			ViolateLaw:        0x03,
			Five: UserecGameScore{
				Win:  0x0011,
				Lose: 0x0012,
				Tie:  0x0013,
			},
			ChineseChess: UserecGameScore{
				Win:  0x0014,
				Lose: 0x0015,
				Tie:  0x0016,
			},
			Conn6: UserecGameScore{
				Win:  0x0017,
				Lose: 0x0018,
				Tie:  0x0019,
			},
			GoChess: UserecGameScore{
				Win:  0x001a,
				Lose: 0x001b,
				Tie:  0x001c,
			},
			DarkChess: UserecGameScore{
				Win:  0x001d,
				Lose: 0x001e,
				Tie:  0x001f,
			},
			MyAngel: "pichu",
		},
		{
			Version:       4194,
			userID:        "Kahou",
			realName:      "林嘉豪",
			nickname:      "Kahou",
			password:      "V3nkaYTLnDPUA",
			UserFlag:      0x02000AE0,
			UserLevel:     0x00000007,
			numLoginDays:  1,
			numPosts:      0,
			firstLogin:    time.Date(2020, 9, 22, 7, 4, 26, 0, time.UTC),
			lastLogin:     time.Date(2020, 9, 22, 7, 4, 26, 0, time.UTC),
			lastHost:      "180.217.174.18",
			Career:        "我的服務單位",
			LastSeen:      time.Date(2020, 9, 22, 7, 4, 26, 0, time.UTC),
			TimeSetAngel:  time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			TimePlayAngel: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			LastSong:      time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			Email:         "creator.kahou@gmail.com",
			Address:       "新北市板橋信義路111號",
			Justify:       "<Email>",
			Over18:        true,
			Pager:         1,

			TimeRemoveBadPost: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			TimeViolateLaw:    time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			BadPost:           0x04,
			ViolateLaw:        0x04,
			Five: UserecGameScore{
				Win:  0x0111,
				Lose: 0x0112,
				Tie:  0x0113,
			},
			ChineseChess: UserecGameScore{
				Win:  0x0114,
				Lose: 0x0115,
				Tie:  0x0116,
			},
			Conn6: UserecGameScore{
				Win:  0x0117,
				Lose: 0x0118,
				Tie:  0x0119,
			},
			GoChess: UserecGameScore{
				Win:  0x011a,
				Lose: 0x011b,
				Tie:  0x011c,
			},
			DarkChess: UserecGameScore{
				Win:  0x011d,
				Lose: 0x011e,
				Tie:  0x011f,
			},
			MyAngel: "Kahou",
		},
		{
			Version:       4194,
			userID:        "Kahou2",
			realName:      "Kahou",
			nickname:      "kahou",
			password:      "R7shIAOZgQCKs",
			UserFlag:      0x02000AE0,
			UserLevel:     0x0000001F,
			numLoginDays:  1,
			numPosts:      0,
			firstLogin:    time.Date(2020, 9, 22, 7, 15, 39, 0, time.UTC),
			lastLogin:     time.Date(2020, 9, 22, 7, 40, 1, 0, time.UTC),
			lastHost:      "180.217.174.18",
			Career:        "我的服務單位",
			LastSeen:      time.Date(2020, 9, 22, 7, 40, 01, 0, time.UTC),
			TimeSetAngel:  time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			TimePlayAngel: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			LastSong:      time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			Email:         "x",
			Address:       "新北市板橋區信義路111號",
			Justify:       "[SYSOP] 09/22/2020 07:51:12 Tue",
			Over18:        true,
			Pager:         1,

			TimeRemoveBadPost: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			TimeViolateLaw:    time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			BadPost:           0x05,
			ViolateLaw:        0x05,
			Five: UserecGameScore{
				Win:  0x0021,
				Lose: 0x0022,
				Tie:  0x0023,
			},
			ChineseChess: UserecGameScore{
				Win:  0x0024,
				Lose: 0x0025,
				Tie:  0x0026,
			},
			Conn6: UserecGameScore{
				Win:  0x0027,
				Lose: 0x0028,
				Tie:  0x0029,
			},
			GoChess: UserecGameScore{
				Win:  0x002a,
				Lose: 0x002b,
				Tie:  0x002c,
			},
			DarkChess: UserecGameScore{
				Win:  0x002d,
				Lose: 0x002e,
				Tie:  0x002f,
			},
			MyAngel: "Kahou2",
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

		if actual.userID != expected.userID {
			t.Errorf("UserID not match with index %d, expected: %v, got: %v", index, expected.userID, actual.userID)
		}

		if actual.realName != expected.realName {
			t.Errorf("RealName not match with index %d, expected: %v, got: %v", index, expected.realName, actual.realName)
		}

		if actual.nickname != expected.nickname {
			t.Errorf("Nickname not match with index %d, expected: %v, got: %v", index, expected.nickname, actual.nickname)
		}

		if actual.password != expected.password {
			t.Errorf("Password not match with index %d, expected: %v, got: %v", index, expected.password, actual.password)
		}

		if actual.UserFlag != expected.UserFlag {
			t.Errorf("UserFlag not match with index %d, expected: 0x%08X, got: 0x%08X", index, expected.UserFlag, actual.UserFlag)
		}

		if actual.UserLevel != expected.UserLevel {
			t.Errorf("UserLevel not match with index %d, expected: 0x%08X, got: 0x%08X", index, expected.UserLevel, actual.UserLevel)
		}

		if actual.numLoginDays != expected.numLoginDays {
			t.Errorf("NumLoginDays not match with index %d, expected: %v, got: %v", index, expected.numLoginDays, actual.numLoginDays)
		}

		if actual.numPosts != expected.numPosts {
			t.Errorf("NumPosts not match with index %d, expected: %v, got: %v", index, expected.numPosts, actual.numPosts)
		}

		if actual.firstLogin.Sub(expected.firstLogin) != 0 {
			t.Errorf("FirstLogin not match with index %d, expected: %v, got: %v", index, expected.firstLogin, actual.firstLogin)
		}

		if actual.lastLogin.Sub(expected.lastLogin) != 0 {
			t.Errorf("LastLogin not match with index %d, expected: %v, got: %v", index, expected.lastLogin, actual.lastLogin)
		}

		if actual.lastHost != expected.lastHost {
			t.Errorf("LastHost not match with index %d, expected: %v, got: %v", index, expected.lastHost, actual.lastHost)
		}

		if actual.money != expected.money {
			t.Errorf("Money not match with index %d, expected: %v, got: %v", index, expected.money, actual.money)
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

		if actual.LastSeen.Sub(expected.LastSeen) != 0 {
			t.Errorf("LastSeen not match with index %d, expected: %v, got: %v", index, expected.LastSeen, actual.LastSeen)
		}

		if actual.TimeSetAngel.Sub(expected.TimeSetAngel) != 0 {
			t.Errorf("TimeSetAngel not match with index %d, expected: %v, got: %v", index, expected.TimeSetAngel, actual.TimeSetAngel)
		}

		if actual.TimePlayAngel.Sub(expected.TimePlayAngel) != 0 {
			t.Errorf("TimePlayAngel not match with index %d, expected: %v, got: %v", index, expected.TimePlayAngel, actual.TimePlayAngel)
		}

		if actual.LastSong.Sub(expected.LastSong) != 0 {
			t.Errorf("LastSong not match with index %d, expected: %v, got: %v", index, expected.LastSong, actual.LastSong)
		}

		if actual.LoginView != expected.LoginView {
			t.Errorf("LoginView not match with index %d, expected: %v, got: %v", index, expected.LoginView, actual.LoginView)
		}

		if actual.ViolateLaw != expected.ViolateLaw {
			t.Errorf("ViolateLaw not match with index %d, expected: %v, got: %v", index, expected.ViolateLaw, actual.ViolateLaw)
		}

		if actual.Five.Win != expected.Five.Win {
			t.Errorf("Five.Win not match with index %d, expected: %v, got: %v", index, expected.Five.Win, actual.Five.Win)
		}

		if actual.Five.Lose != expected.Five.Lose {
			t.Errorf("Five.Lose not match with index %d, expected: %v, got: %v", index, expected.Five.Lose, actual.Five.Lose)
		}

		if actual.Five.Tie != expected.Five.Tie {
			t.Errorf("Five.Tie not match with index %d, expected: %v, got: %v", index, expected.Five.Tie, actual.Five.Tie)
		}

		if actual.ChineseChess.Win != expected.ChineseChess.Win {
			t.Errorf("ChineseChess.Win not match with index %d, expected: %v, got: %v", index, expected.ChineseChess.Win, actual.ChineseChess.Win)
		}

		if actual.ChineseChess.Lose != expected.ChineseChess.Lose {
			t.Errorf("ChineseChess.Lose not match with index %d, expected: %v, got: %v", index, expected.ChineseChess.Lose, actual.ChineseChess.Lose)
		}

		if actual.ChineseChess.Tie != expected.ChineseChess.Tie {
			t.Errorf("ChineseChess.Tie not match with index %d, expected: %v, got: %v", index, expected.ChineseChess.Tie, actual.ChineseChess.Tie)
		}

		if actual.Conn6.Win != expected.Conn6.Win {
			t.Errorf("Conn6.Win not match with index %d, expected: %v, got: %v", index, expected.Conn6.Win, actual.Conn6.Win)
		}

		if actual.Conn6.Lose != expected.Conn6.Lose {
			t.Errorf("Conn6.Lose not match with index %d, expected: %v, got: %v", index, expected.Conn6.Lose, actual.Conn6.Lose)
		}

		if actual.Conn6.Tie != expected.Conn6.Tie {
			t.Errorf("Conn6.Tie not match with index %d, expected: %v, got: %v", index, expected.Conn6.Tie, actual.Conn6.Tie)
		}

		if actual.GoChess.Win != expected.GoChess.Win {
			t.Errorf("GoChess.Win not match with index %d, expected: %v, got: %v", index, expected.GoChess.Win, actual.GoChess.Win)
		}

		if actual.GoChess.Lose != expected.GoChess.Lose {
			t.Errorf("GoChess.Lose not match with index %d, expected: %v, got: %v", index, expected.GoChess.Lose, actual.GoChess.Lose)
		}

		if actual.GoChess.Tie != expected.GoChess.Tie {
			t.Errorf("GoChess.Tie not match with index %d, expected: %v, got: %v", index, expected.GoChess.Tie, actual.GoChess.Tie)
		}

		if actual.DarkChess.Win != expected.DarkChess.Win {
			t.Errorf("DarkChess.Win not match with index %d, expected: %v, got: %v", index, expected.DarkChess.Win, actual.DarkChess.Win)
		}

		if actual.DarkChess.Lose != expected.DarkChess.Lose {
			t.Errorf("DarkChess.Lose not match with index %d, expected: %v, got: %v", index, expected.DarkChess.Lose, actual.DarkChess.Lose)
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

		if actual.DarkChess.Tie != expected.DarkChess.Tie {
			t.Errorf("DarkChess.Tie not match with index %d, expected: %v, got: %v", index, expected.DarkChess.Tie, actual.DarkChess.Tie)
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

		if actual.TimeRemoveBadPost.Sub(expected.TimeRemoveBadPost) != 0 {
			t.Errorf("TimeRemoveBadPost not match with index %d, expected: %v, got: %v", index, expected.TimeRemoveBadPost, actual.TimeRemoveBadPost)
		}

		if actual.TimeViolateLaw.Sub(expected.TimeViolateLaw) != 0 {
			t.Errorf("TimeViolateLaw not match with index %d, expected: %v, got: %v", index, expected.TimeViolateLaw, actual.TimeViolateLaw)
		}

	}

}

func TestEncodingUserec(t *testing.T) {
	type TestCase struct {
		Input    Userec
		Expected []byte
	}

	testcase := []TestCase{
		{
			Input: Userec{
				Version:       4194,
				userID:        "SYSOP",
				realName:      "CodingMan",
				nickname:      "神",
				password:      "bhwvOJtfT1TAI",
				UserFlag:      0x02000A60,
				UserLevel:     0x20000407,
				numLoginDays:  2,
				numPosts:      0,
				firstLogin:    time.Date(2020, 9, 21, 9, 41, 28, 0, time.UTC),
				lastLogin:     time.Date(2020, 9, 22, 6, 28, 14, 0, time.UTC),
				lastHost:      "59.124.167.226",
				money:         0,
				Address:       "新竹縣子虛鄉烏有村543號",
				Over18:        true,
				Pager:         1,
				Invisible:     false,
				Career:        "全景軟體",
				LastSeen:      time.Date(2020, 9, 21, 9, 41, 28, 0, time.UTC),
				TimeSetAngel:  time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
				TimePlayAngel: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
				LastSong:      time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),

				TimeRemoveBadPost: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
				TimeViolateLaw:    time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
				BadPost:           0x01,
				ViolateLaw:        0x01,
				Five: UserecGameScore{
					Win:  0x0001,
					Lose: 0x0002,
					Tie:  0x0003,
				},
				ChineseChess: UserecGameScore{
					Win:  0x0004,
					Lose: 0x0005,
					Tie:  0x0006,
				},
				Conn6: UserecGameScore{
					Win:  0x0007,
					Lose: 0x0008,
					Tie:  0x0009,
				},
				GoChess: UserecGameScore{
					Win:  0x000a,
					Lose: 0x000b,
					Tie:  0x000c,
				},
				DarkChess: UserecGameScore{
					Win:  0x000d,
					Lose: 0x000e,
					Tie:  0x000f,
				},
				MyAngel: "SYSOP",
			},
			Expected: hexToByte(`
			6210 0000 5359 534f 5000 0000 0000 0000
0043 6f64 696e 674d 616e 0000 0000 0000
0000 0000 00af ab00 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0062 6877
764f 4a74 6654 3154 4149 0000 600a 0002
0000 0000 0704 0020 0200 0000 0000 0000
4875 685f 7e99 695f 3539 2e31 3234 2e31
3637 2e32 3236 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 b773 a6cb bfa4 a46c b5ea b66d af51
a6b3 a7f8 3534 33b8 b900 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0100
0100 0000 0000 0000 0000 0000 a5fe b4ba
b36e c5e9 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 4875 685f
0000 0000 0000 0000 0000 0000 0000 0000
0000 0100 0100 0200 0300 0400 0500 0600
0700 0800 0900 0000 0a00 0b00 0c00 0d00
0e00 0000 0001 0f00 5359 534F 5000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
0000 0000 0000 0000 0000 0000 0000 0000
			`),
		},
	}

	for index, c := range testcase {
		b, err := c.Input.MarshalToByte()
		t.Logf("log: %q, %q", b, err)
		if hex.Dump(b) != hex.Dump(c.Expected) {
			t.Errorf("Expected byte not match in index %d, expected: \n%s\n, got: \n%s", index, hex.Dump(c.Expected), hex.Dump(b))
		}

	}

}
