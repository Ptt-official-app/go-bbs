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
	}

}
