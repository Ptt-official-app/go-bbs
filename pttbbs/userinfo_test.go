package pttbbs

import (
	"encoding/hex"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestOpenUserInfoCache(t *testing.T) {
	actualUserInfoCache, err := OpenUserInfoCache("testcase/userinfo/01.USERINFO")
	if err != nil {
		t.Errorf("OpenUserInfoCache() error =%v", err)
		return
	}

	expected := []UserInfoCache{
		{
			uid:       1,
			pid:       29798,
			SockAddr:  0,
			UserLevel: 65535,
			userID:    "SYSOP",
			nickname:  "神",
			from:      "1.164.105.65",
			fromIP:    1097442305,
			DarkChess: GameScore{
				Win:  0,
				Lose: 0,
				Tie:  0,
			},
			AngelPause:  0,
			FriendTotal: 1,
			NumFriends:  0,
			// MyFriend not set bc of zero value
			// FriendOnline not set bc most value is zero
			// Reject not set bc of zero value
			MsgCount:  0,
			Birth:     false,
			Active:    false,
			Invisible: false,
			Mode:      22,
			Pager:     true,
			Conn6: GameScore{
				Win:  0,
				Lose: 0,
				Tie:  0,
			},
			LastAct:    time.Date(2021, 2, 12, 14, 36, 22, 0, time.UTC),
			Alerts:     0,
			Signature:  0,
			DestUID:    0,
			DestUIP:    0,
			SockActive: 0,
			InChat:     0,
			ChatID:     "",
			LockMode:   0,
			Turn:       false,
			MateID:     "",
			Color:      0,
			Five: GameScore{
				Win:  0,
				Lose: 0,
				Tie:  0,
			},
			ChineseChess: GameScore{
				Win:  0,
				Lose: 0,
				Tie:  0,
			},
			ChessEloRating: 0,

			GoChess: GameScore{
				Win:  0,
				Lose: 0,
				Tie:  0,
			},
			WithMe: 0,
			brcID:  0,
			wbTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	// setting expected FriendOnline value
	expected[0].FriendOnline[0] = 268435462
	// setting expected Msgs value
	// TODO: update this part after getting the actual struct for MessageQueueCache
	for idx := range expected[0].Msgs {
		expected[0].Msgs[idx] = MessageQueueCache(make([]byte, 100))
	}

	for idx, actual := range actualUserInfoCache {
		ok := idx < len(expected)
		if !ok {
			t.Errorf("expected case %d not exist", idx)
			break
		}

		exp := expected[idx]

		if actual.uid != exp.uid {
			t.Errorf("uid not match with index %d, expected: %v, got: %v", idx, exp.uid, actual.uid)
		}

		if actual.pid != exp.pid {
			t.Errorf("pid not match with index %d, expected: %v, got: %v", idx, exp.pid, actual.pid)
		}

		if actual.SockAddr != exp.SockAddr {
			t.Errorf("SockAddr not match with index %d, expected: %v, got: %v", idx, exp.SockAddr, actual.SockAddr)
		}

		if actual.UserLevel != exp.UserLevel {
			t.Errorf("UserLevel not match with index %d, expected: %v, got: %v", idx, exp.UserLevel, actual.UserLevel)
		}

		if actual.userID != exp.userID {
			t.Errorf("userID not match with index %d, expected: %v, got: %v", idx, exp.userID, actual.userID)
		}

		if actual.nickname != exp.nickname {
			t.Errorf("nickname not match with index %d, expected: %v, got: %v", idx, exp.nickname, actual.nickname)
		}

		if actual.from != exp.from {
			t.Errorf("from not match with index %d, expected: %v, got: %v", idx, exp.from, actual.from)
		}

		if actual.fromIP != exp.fromIP {
			t.Errorf("fromIP not match with index %d, expected: %v, got: %v", idx, exp.fromIP, actual.fromIP)
		}

		if !reflect.DeepEqual(actual.DarkChess, exp.DarkChess) {
			t.Errorf("DarkChess not match with index %d, expected: %v, got: %v", idx, exp.DarkChess, actual.DarkChess)
		}

		if actual.AngelPause != exp.AngelPause {
			t.Errorf("AngelPause not match with index %d, expected: %v, got: %v", idx, exp.AngelPause, actual.AngelPause)
		}

		if actual.FriendTotal != exp.FriendTotal {
			t.Errorf("FriendTotal not match with index %d, expected: %v, got: %v", idx, exp.FriendTotal, actual.FriendTotal)
		}

		if actual.NumFriends != exp.NumFriends {
			t.Errorf("Numfriends not match with index %d, expected: %v, got: %v", idx, exp.NumFriends, actual.NumFriends)
		}

		if !reflect.DeepEqual(actual.MyFriend, exp.MyFriend) {
			t.Errorf("MyFriend not match with index %d, expected: %v, got: %v", idx, exp.MyFriend, actual.MyFriend)
		}

		if !reflect.DeepEqual(actual.FriendOnline, exp.FriendOnline) {
			t.Errorf("FriendOnline not match with index %d, expected: %v, got: %v", idx, exp.FriendOnline, actual.FriendOnline)
		}

		if !reflect.DeepEqual(actual.Reject, exp.Reject) {
			t.Errorf("Reject not match with index %d, expected: %v, got: %v", idx, exp.Reject, actual.Reject)
		}

		if actual.MsgCount != exp.MsgCount {
			t.Errorf("MsgCount not match with index %d, expected: %v, got: %v", idx, exp.MsgCount, actual.MsgCount)
		}

		if !reflect.DeepEqual(actual.Msgs, exp.Msgs) {
			t.Errorf("Msgs not match with index %d, expected: %v, got: %v", idx, exp.Msgs, actual.Msgs)
		}

		if actual.Birth != exp.Birth {
			t.Errorf("Birth not match with index %d, expected: %v, got: %v", idx, exp.Birth, actual.Birth)
		}

		if actual.Active != exp.Active {
			t.Errorf("Active not match with index %d, expected: %v, got: %v", idx, exp.Active, actual.Active)
		}

		if actual.Invisible != exp.Invisible {
			t.Errorf("Invisible not match with index %d, expected: %v, got: %v", idx, exp.Invisible, actual.Invisible)
		}

		if actual.Mode != exp.Mode {
			t.Errorf("Mode not match with index %d, expected: %v, got: %v", idx, exp.Mode, actual.Mode)
		}

		if actual.Pager != exp.Pager {
			t.Errorf("Pager not match with index %d, expected: %v, got: %v", idx, exp.Pager, actual.Pager)
		}

		if !reflect.DeepEqual(actual.Conn6, exp.Conn6) {
			t.Errorf("Conn6 not match with index %d, expected: %v, got: %v", idx, exp.Conn6, actual.Conn6)
		}

		if actual.LastAct.Sub(exp.LastAct) != 0 {
			t.Errorf("LastAct not match with index %d, expected: %v, got: %v", idx, exp.LastAct, actual.LastAct)
		}

		if actual.Alerts != exp.Alerts {
			t.Errorf("Alerts not match with index %d, expected: %v, got: %v", idx, exp.Alerts, actual.Alerts)
		}

		if actual.Signature != exp.Signature {
			t.Errorf("Mode not match with index %d, expected: %v, got: %v", idx, exp.Mode, actual.Mode)
		}

		if actual.DestUID != exp.DestUID {
			t.Errorf("DestUID not match with index %d, expected: %v, got: %v", idx, exp.DestUID, actual.DestUID)
		}

		if actual.DestUIP != exp.DestUIP {
			t.Errorf("DestUIP not match with index %d, expected: %v, got: %v", idx, exp.DestUIP, actual.DestUIP)
		}

		if actual.SockActive != exp.SockActive {
			t.Errorf("SockActive not match with index %d, expected: %v, got: %v", idx, exp.SockActive, actual.SockActive)
		}

		if actual.InChat != exp.InChat {
			t.Errorf("InChat not match with index %d, expected: %v, got: %v", idx, exp.InChat, actual.InChat)
		}

		if actual.ChatID != exp.ChatID {
			t.Errorf("ChatID not match with index %d, expected: %v, got: %v", idx, exp.ChatID, actual.ChatID)
		}

		if actual.LockMode != exp.LockMode {
			t.Errorf("LockMode not match with index %d, expected: %v, got: %v", idx, exp.LockMode, actual.LockMode)
		}

		if actual.Turn != exp.Turn {
			t.Errorf("Turn not match with index %d, expected: %v, got: %v", idx, exp.Turn, actual.Turn)
		}

		if actual.MateID != exp.MateID {
			t.Errorf("MateID not match with index %d, expected: %v, got: %v", idx, exp.MateID, actual.MateID)
		}

		if actual.Color != exp.Color {
			t.Errorf("Color not match with index %d, expected: %v, got: %v", idx, exp.Color, actual.Color)
		}

		if !reflect.DeepEqual(actual.Five, exp.Five) {
			t.Errorf("Five not match with index %d, expected: %v, got: %v", idx, exp.Five, actual.Five)
		}

		if !reflect.DeepEqual(actual.ChineseChess, exp.ChineseChess) {
			t.Errorf("ChineseChess not match with index %d, expected: %v, got: %v", idx, exp.ChineseChess, actual.ChineseChess)
		}

		if actual.ChessEloRating != exp.ChessEloRating {
			t.Errorf("ChessEloRating not match with index %d, expected: %v, got: %v", idx, exp.ChessEloRating, actual.ChessEloRating)
		}

		if !reflect.DeepEqual(actual.GoChess, exp.GoChess) {
			t.Errorf("GoChess not match with index %d, expected: %v, got: %v", idx, exp.GoChess, actual.GoChess)
		}

		if actual.WithMe != exp.WithMe {
			t.Errorf("Color not match with index %d, expected: %v, got: %v", idx, exp.Color, actual.Color)
		}

		if actual.brcID != exp.brcID {
			t.Errorf("Color not match with index %d, expected: %v, got: %v", idx, exp.Color, actual.Color)
		}

		if actual.wbTime.Sub(exp.wbTime) != 0 {
			t.Errorf("wbTime not match with index %d, expected: %v, got: %v", idx, exp.wbTime, actual.wbTime)
		}

	}
}

func TestUnmarshalUserInfo(t *testing.T) {
	file, err := os.Open("testcase/userinfo/01.USERINFO")
	if err != nil {
		t.Errorf("Error when opening test file.")
	}

	buf := make([]byte, 3484)

	_, err = file.Read(buf)
	if err != nil && err != io.EOF {
		t.Errorf("Error when reading test file.")
	}

	type TestCase struct {
		Input    UserInfoCache
		Expected []byte
	}

	testcase := []TestCase{
		{
			Input: UserInfoCache{
				uid:       1,
				pid:       29798,
				SockAddr:  0,
				UserLevel: 65535,
				userID:    "SYSOP",
				nickname:  "神",
				from:      "1.164.105.65",
				fromIP:    1097442305,
				DarkChess: GameScore{
					Win:  0,
					Lose: 0,
					Tie:  0,
				},
				AngelPause:  0,
				FriendTotal: 1,
				NumFriends:  0,
				// MyFriend not set bc of zero value
				// FriendOnline not set bc most value is zero
				// Reject not set bc of zero value
				MsgCount:  0,
				Birth:     false,
				Active:    false,
				Invisible: false,
				Mode:      22,
				Pager:     true,
				Conn6: GameScore{
					Win:  0,
					Lose: 0,
					Tie:  0,
				},
				LastAct:    time.Date(2021, 2, 12, 14, 36, 22, 0, time.UTC),
				Alerts:     0,
				Signature:  0,
				DestUID:    0,
				DestUIP:    0,
				SockActive: 0,
				InChat:     0,
				ChatID:     "",
				LockMode:   0,
				Turn:       false,
				MateID:     "",
				Color:      0,
				Five: GameScore{
					Win:  0,
					Lose: 0,
					Tie:  0,
				},
				ChineseChess: GameScore{
					Win:  0,
					Lose: 0,
					Tie:  0,
				},
				ChessEloRating: 0,

				GoChess: GameScore{
					Win:  0,
					Lose: 0,
					Tie:  0,
				},
				WithMe: 0,
				brcID:  0,
				wbTime: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			Expected: buf,
		},
	}

	// setting input FriendOnline value
	testcase[0].Input.FriendOnline[0] = 268435462
	// setting input Msgs value
	// TODO: update this part after getting the actual struct for MessageQueueCache
	for idx := range testcase[0].Input.Msgs {
		testcase[0].Input.Msgs[idx] = MessageQueueCache(make([]byte, 100))
	}

	for idx, c := range testcase {
		actualBinary := c.Input.MarshalBinary()
		t.Logf("Binary Output:%q", actualBinary)
		actualString := hex.Dump(actualBinary)
		expectedString := hex.Dump(c.Expected)
		if len(expectedString) != len(actualString) {
			t.Errorf("Length not match in test case %d, expected: \n%d\n, got: \n%d", idx, len(expectedString), len(actualString))
		}
		for strIdx, val := range actualString {
			if byte(val) != expectedString[strIdx] {
				t.Errorf("Byte index %d not match in test case %d, expected: \n%v\n, got: \n%v", strIdx, idx, expectedString[strIdx], byte(val))
			}
		}

	}
}
