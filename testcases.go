package bbs

var (
	testUserecBig51 = &UserecRaw{
		Version:    PTT_PASSWD_VERSION,
		UserID:     [PTT_IDLEN + 1]byte{83, 89, 83, 79, 80},
		RealName:   [PTT_REALNAMESZ]byte{67, 111, 100, 105, 110, 103, 77, 97, 110},
		Nickname:   [PTT_NICKNAMESZ]byte{175, 171},
		PasswdHash: [PTT_PASSLEN]byte{98, 104, 119, 118, 79, 74, 116, 102, 84, 49, 84, 65, 73, 0},

		UFlag:        33557088,
		UserLevel:    536871943,
		NumLoginDays: 2,
		NumPosts:     0,
		FirstLogin:   1600681288,
		LastLogin:    1600756094,
		LastHost:     [PTT_IPV4LEN + 1]byte{53, 57, 46, 49, 50, 52, 46, 49, 54, 55, 46, 50, 50, 54},
		Address:      [PTT_ADDRESSSZ]byte{183, 115, 166, 203, 191, 164, 164, 108, 181, 234, 182, 109, 175, 81, 166, 179, 167, 248, 53, 52, 51, 184, 185},
		Over18:       1,
		Pager:        1,
		Career:       [PTT_CAREERSZ]byte{165, 254, 180, 186, 179, 110, 197, 233},
		LastSeen:     1600681288,
	}

	testUserec1 = &Userec{
		Version:      PTT_PASSWD_VERSION,
		Userid:       "SYSOP",
		Realname:     "CodingMan",
		Nickname:     "神",
		Passwd:       "bhwvOJtfT1TAI",
		Uflag:        33557088,
		Userlevel:    536871943,
		Numlogindays: 2,
		Numposts:     0,
		Firstlogin:   1600681288,
		Lastlogin:    1600756094,
		Lasthost:     "59.124.167.226",
	}

	testUserec2 = &Userec{
		Version:      PTT_PASSWD_VERSION,
		Userid:       "CodingMan",
		Realname:     "朱元璋",
		Nickname:     "程式俠",
		Passwd:       "u8mLG.ktfOk3w",
		Uflag:        33557216,
		Userlevel:    31,
		Numlogindays: 1,
		Numposts:     0,
		Firstlogin:   1600737659,
		Lastlogin:    1600737960,
		Lasthost:     "59.124.167.226",
	}

	testUserec3 = &Userec{
		Version:      PTT_PASSWD_VERSION,
		Userid:       "pichu",
		Realname:     "Pichu",
		Nickname:     "Pichu",
		Passwd:       "KO27TyME.3/tw",
		Uflag:        33557216,
		Userlevel:    7,
		Numlogindays: 1,
		Numposts:     0,
		Firstlogin:   1600755675,
		Lastlogin:    1600766204,
		Lasthost:     "103.246.218.43",
	}

	testUserec4 = &Userec{
		Version:      PTT_PASSWD_VERSION,
		Userid:       "Kahou",
		Realname:     "林嘉豪",
		Nickname:     "Kahou",
		Passwd:       "V3nkaYTLnDPUA",
		Uflag:        33557216,
		Userlevel:    7,
		Numlogindays: 1,
		Numposts:     0,
		Firstlogin:   1600758266,
		Lastlogin:    1600758266,
		Lasthost:     "180.217.174.18",
	}

	testUserec5 = &Userec{
		Version:      PTT_PASSWD_VERSION,
		Userid:       "Kahou2",
		Realname:     "Kahou",
		Nickname:     "kahou",
		Passwd:       "R7shIAOZgQCKs",
		Uflag:        33557216,
		Userlevel:    31,
		Numlogindays: 1,
		Numposts:     0,
		Firstlogin:   1600758939,
		Lastlogin:    1600760401,
		Lasthost:     "180.217.174.18",
	}
	testUserecEmpty = &Userec{}

	testOpenUserecFile1 []*Userec = nil
)
