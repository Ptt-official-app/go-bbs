package cmbbs

import "github.com/PichuChen/go-bbs/ptttype"

var (
	testUserecBig51 = &ptttype.UserecBig5{
		Version:    4194,
		UserID:     [ptttype.IDLEN + 1]byte{83, 89, 83, 79, 80},
		RealName:   [ptttype.REALNAMESZ]byte{67, 111, 100, 105, 110, 103, 77, 97, 110},
		Nickname:   [ptttype.NICKNAMESZ]byte{175, 171},
		PasswdHash: [ptttype.PASSLEN]byte{98, 104, 119, 118, 79, 74, 116, 102, 84, 49, 84, 65, 73, 0},

		UFlag:        33557088,
		UserLevel:    536871943,
		NumLoginDays: 2,
		NumPosts:     0,
		FirstLogin:   1600681288,
		LastLogin:    1600756094,
		LastHost:     [ptttype.IPV4LEN + 1]byte{53, 57, 46, 49, 50, 52, 46, 49, 54, 55, 46, 50, 50, 54},
		Address:      [ptttype.ADDRESSSZ]byte{183, 115, 166, 203, 191, 164, 164, 108, 181, 234, 182, 109, 175, 81, 166, 179, 167, 248, 53, 52, 51, 184, 185},
		Over18:       1,
		Pager:        1,
		Career:       [ptttype.CAREERSZ]byte{165, 254, 180, 186, 179, 110, 197, 233},
		LastSeen:     1600681288,
	}
)
