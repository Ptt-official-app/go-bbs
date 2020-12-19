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
			Userid:       "SYSOP",
			Realname:     "CodingMan",
			Nickname:     "神",
			Passwd:       "bhwvOJtfT1TAI",
			Uflag:        0x02000A60,
			Userlevel:    0x20000407,
			Numlogindays: 2,
			Numposts:     0,
			Firstlogin:   1600681288,
			Lastlogin:    1600756094,
			Lasthost:     "59.124.167.226",
		},
		Userec{
			Version:      4194,
			Userid:       "CodingMan",
			Realname:     "朱元璋",
			Nickname:     "程式俠",
			Passwd:       "u8mLG.ktfOk3w",
			Uflag:        0x02000AE0,
			Userlevel:    0x0000001F,
			Numlogindays: 1,
			Numposts:     0,
			Firstlogin:   1600737659,
			Lastlogin:    1600737960,
			Lasthost:     "59.124.167.226",
		},
		Userec{
			Version:      4194,
			Userid:       "pichu",
			Realname:     "Pichu",
			Nickname:     "Pichu",
			Passwd:       "KO27TyME.3/tw",
			Uflag:        0x02000AE0,
			Userlevel:    0x00000007,
			Numlogindays: 1,
			Numposts:     0,
			Firstlogin:   1600755675,
			Lastlogin:    1600766204,
			Lasthost:     "103.246.218.43",
		},
		Userec{
			Version:      4194,
			Userid:       "Kahou",
			Realname:     "林嘉豪",
			Nickname:     "Kahou",
			Passwd:       "V3nkaYTLnDPUA",
			Uflag:        0x02000AE0,
			Userlevel:    0x00000007,
			Numlogindays: 1,
			Numposts:     0,
			Firstlogin:   1600758266,
			Lastlogin:    1600758266,
			Lasthost:     "180.217.174.18",
		},
		Userec{
			Version:      4194,
			Userid:       "Kahou2",
			Realname:     "Kahou",
			Nickname:     "kahou",
			Passwd:       "R7shIAOZgQCKs",
			Uflag:        0x02000AE0,
			Userlevel:    0x0000001F,
			Numlogindays: 1,
			Numposts:     0,
			Firstlogin:   1600758939,
			Lastlogin:    1600760401,
			Lasthost:     "180.217.174.18",
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

		if actual.Userid != expected.Userid {
			t.Errorf("Userid not match with index %d, expected: %v, got: %v", index, expected.Userid, actual.Userid)
		}

		if actual.Realname != expected.Realname {
			t.Errorf("Realname not match with index %d, expected: %v, got: %v", index, expected.Realname, actual.Realname)
		}

		if actual.Nickname != expected.Nickname {
			t.Errorf("Nickname not match with index %d, expected: %v, got: %v", index, expected.Nickname, actual.Nickname)
		}

		if actual.Passwd != expected.Passwd {
			t.Errorf("Passwd not match with index %d, expected: %v, got: %v", index, expected.Passwd, actual.Passwd)
		}

		if actual.Uflag != expected.Uflag {
			t.Errorf("Uflag not match with index %d, expected: 0x%08X, got: 0x%08X", index, expected.Uflag, actual.Uflag)
		}

		if actual.Userlevel != expected.Userlevel {
			t.Errorf("Userlevel not match with index %d, expected: 0x%08X, got: 0x%08X", index, expected.Userlevel, actual.Userlevel)
		}

		if actual.Numlogindays != expected.Numlogindays {
			t.Errorf("Numlogindays not match with index %d, expected: %v, got: %v", index, expected.Numlogindays, actual.Numlogindays)
		}

		if actual.Numposts != expected.Numposts {
			t.Errorf("Numposts not match with index %d, expected: %v, got: %v", index, expected.Numposts, actual.Numposts)
		}

		if actual.Firstlogin != expected.Firstlogin {
			t.Errorf("Firstlogin not match with index %d, expected: %v, got: %v", index, expected.Firstlogin, actual.Firstlogin)
		}

		if actual.Lastlogin != expected.Lastlogin {
			t.Errorf("Lastlogin not match with index %d, expected: %v, got: %v", index, expected.Lastlogin, actual.Lastlogin)
		}

		if actual.Lasthost != expected.Lasthost {
			t.Errorf("Lasthost not match with index %d, expected: %v, got: %v", index, expected.Lasthost, actual.Lasthost)
		}
	}

}
