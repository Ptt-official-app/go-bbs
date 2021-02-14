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
	"github.com/PichuChen/go-bbs/crypt"

	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const (
	PosOfPttPasswdVersion      = 0
	PosOfPttPasswdUserId       = PosOfPttPasswdVersion + 4
	PosOfPttPasswdRealName     = PosOfPttPasswdUserId + PTT_IDLEN + 1
	PosOfPttPasswdNickname     = PosOfPttPasswdRealName + 20
	PosOfPttPasswdPassword     = PosOfPttPasswdNickname + 24
	PosOfPttPasswdUserFlag     = PosOfPttPasswdPassword + 15
	PosOfPttPasswdUserLevel    = PosOfPttPasswdUserFlag + 4 + 4
	PosOfPttPasswdNumLoginDays = PosOfPttPasswdUserLevel + 4
	PosOfPttPasswdNumPosts     = PosOfPttPasswdNumLoginDays + 4
	PosOfPttPasswdFirstLogin   = PosOfPttPasswdNumPosts + 4
	PosOfPttPasswdLastLogin    = PosOfPttPasswdFirstLogin + 4
	PosOfPttPasswdLastHost     = PosOfPttPasswdLastLogin + 4
	PosOfPttPasswdMoney        = PosOfPttPasswdLastHost + PTT_IPV4LEN + 1
	PosOfPttPasswdEmail        = PosOfPttPasswdMoney + 4 + 4
	PosOfPttPasswdAddress      = PosOfPttPasswdEmail + PTT_EMAILSZ
	PosOfPttPasswdJustify      = PosOfPttPasswdAddress + PTT_ADDRESSSZ
	PosOfPttPasswdOver18       = PosOfPttPasswdJustify + PTT_REGLEN + 1 + 3
	PosOfPttPasswdPagerUiType  = PosOfPttPasswdOver18 + 1
	PosOfPttPasswdPager        = PosOfPttPasswdPagerUiType + 1
	PosOfPttPasswdInvisible    = PosOfPttPasswdPager + 1
	PosOfPttPasswdExMailBox    = PosOfPttPasswdInvisible + 1 + 2

	PosOfPttPasswdCareer        = PosOfPttPasswdExMailBox + 4 + 4
	PosOfPttPasswdRole          = PosOfPttPasswdCareer + 40 + 20 + 4 + 44
	PosOfPttPasswdLastSeen      = PosOfPttPasswdRole + 4
	PosOfPttPasswdTimeSetAngel  = PosOfPttPasswdLastSeen + 4
	PosOfPttPasswdTimePlayAngel = PosOfPttPasswdTimeSetAngel + 4

	PosOfPttPasswdLastSong  = PosOfPttPasswdTimePlayAngel + 4
	PosOfPttPasswdLoginView = PosOfPttPasswdLastSong + 4

	PosOfPttPasswdLawCounter = PosOfPttPasswdLoginView + 2
	PosOfPttPasswdFiveWin    = PosOfPttPasswdLawCounter + 2
	PosOfPttPasswdFiveLose   = PosOfPttPasswdFiveWin + 2
	PosOfPttPasswdFiveTie    = PosOfPttPasswdFiveLose + 2
	PosOfPttPasswdChcWin     = PosOfPttPasswdFiveTie + 2
	PosOfPttPasswdChcLose    = PosOfPttPasswdChcWin + 2
	PosOfPttPasswdChcTie     = PosOfPttPasswdChcLose + 2
	PosOfPttPasswdConn6Win   = PosOfPttPasswdChcTie + 2
	PosOfPttPasswdConn6Lose  = PosOfPttPasswdConn6Win + 2
	PosOfPttPasswdConn6Tie   = PosOfPttPasswdConn6Lose + 2
	PosOfPttPasswdGoWin      = PosOfPttPasswdConn6Tie + 4
	PosOfPttPasswdGoLose     = PosOfPttPasswdGoWin + 2
	PosOfPttPasswdGoTie      = PosOfPttPasswdGoLose + 2
	PosOfPttPasswdDarkWin    = PosOfPttPasswdGoTie + 2
	PosOfPttPasswdDarkLose   = PosOfPttPasswdDarkWin + 2
	PosOfPttPasswdUaVersion  = PosOfPttPasswdDarkLose + 2

	PosOfPttPasswdSignature = PosOfPttPasswdUaVersion + 1
	PosOfPttPasswdBadPost   = PosOfPttPasswdSignature + 2
	PosOfPttPasswdDarkTie   = PosOfPttPasswdBadPost + 1
	PosOfPttPasswdMyAngel   = PosOfPttPasswdDarkTie + 2

	PosOfPttPasswdChessEloRating    = PosOfPttPasswdMyAngel + PTT_IDLEN + 1 + 1
	PosOfPttPasswdWithMe            = PosOfPttPasswdChessEloRating + 2
	PosOfPttPasswdTimeRemoveBadPost = PosOfPttPasswdWithMe + 4
	PosOfPttPasswdTimeViolateLaw    = PosOfPttPasswdTimeRemoveBadPost + 4
)

// https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h

type UserecGameScore struct {
	Win  uint16
	Lose uint16
	Tie  uint16
}

type Userec struct {
	Version  uint32 // Magic Number
	userId   string // 使用者帳號，或稱使用者 ID
	realName string // 真實姓名
	nickname string // 暱稱
	password string // 密碼，預設為 crypt, 不同版本實作可能不同

	UserFlag     uint32
	UserLevel    uint32 // 權限
	numLoginDays uint32
	numPosts     uint32
	firstLogin   time.Time
	lastLogin    time.Time
	lastHost     string
	money        int32

	Email   string
	Address string
	Justify string

	Over18      bool
	PagerUIType uint8
	Pager       uint8
	Invisible   bool

	ExMailBox uint32

	Career        string
	Role          uint32
	LastSeen      time.Time
	TimeSetAngel  time.Time
	TimePlayAngel time.Time

	LastSong  time.Time
	LoginView uint32

	ViolateLaw   uint16
	Five         UserecGameScore
	ChineseChess UserecGameScore
	Conn6        UserecGameScore
	GoChess      UserecGameScore
	DarkChess    UserecGameScore
	UaVersion    uint8 // User Agreement Version

	Signature         uint8
	BadPost           uint8
	MyAngel           string
	ChessEloRating    uint16
	WithMe            uint32
	TimeRemoveBadPost time.Time
	TimeViolateLaw    time.Time
}

func (u *Userec) HashedPassword() string {
	return u.password
}

// VerifyPassword will check user's password is OK. it will return null
// when OK and error when there are something wrong
func (u *Userec) VerifyPassword(password string) error {
	res, err := crypt.Fcrypt([]byte(password), []byte(u.password[:2]))
	if err != nil {
		return err
	}
	str := strings.Trim(string(res), "\x00")
	// log.Println("res", str, err, []byte(str), []byte(u.Password))

	if str != u.password {
		return fmt.Errorf("password incorrect")
	}
	return nil
}

func (u *Userec) UserId() string { return u.userId }

// Nickname return a string for user's nickname, this string may change
// depend on user's mood, return empty string if this bbs system do not support
func (u *Userec) Nickname() string { return u.nickname }

// RealName return a string for user's real name, this string may not be changed
// return empty string if this bbs system do not support
func (u *Userec) RealName() string { return u.realName }

// NumLoginDays return how many days this have been login since account created.
func (u *Userec) NumLoginDays() int { return int(u.numLoginDays) }

// NumPosts return how many posts this user has posted.
func (u *Userec) NumPosts() int { return int(u.numPosts) }

// Money return the money this user have.
func (u *Userec) Money() int { return int(u.money) }

func (u *Userec) LastLogin() time.Time {
	return u.lastLogin
}

func (u *Userec) LastHost() string {
	return u.lastHost
}

func OpenUserecFile(filename string) ([]*Userec, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ret := []*Userec{}

	for {
		buf := make([]byte, 512)
		_, err := file.Read(buf)
		// log.Println(len, buf, err)
		if err == io.EOF {
			break
		}

		f, err := NewUserecWithByte(buf)
		if err != nil {
			return nil, err
		}
		ret = append(ret, f)
		// log.Println(f.Filename)

	}

	return ret, nil

}

func NewUserecWithByte(data []byte) (*Userec, error) {
	user := &Userec{}
	user.Version = binary.LittleEndian.Uint32(data[PosOfPttPasswdVersion : PosOfPttPasswdVersion+4])
	user.userId = newStringFormCString(data[PosOfPttPasswdUserId : PosOfPttPasswdUserId+PTT_IDLEN+1])
	user.realName = newStringFormBig5UAOCString(data[PosOfPttPasswdRealName : PosOfPttPasswdRealName+20])
	user.nickname = newStringFormBig5UAOCString(data[PosOfPttPasswdNickname : PosOfPttPasswdNickname+24])
	user.password = newStringFormCString(data[PosOfPttPasswdPassword : PosOfPttPasswdPassword+PTT_PASSLEN])

	user.UserFlag = binary.LittleEndian.Uint32(data[PosOfPttPasswdUserFlag : PosOfPttPasswdUserFlag+4])
	user.UserLevel = binary.LittleEndian.Uint32(data[PosOfPttPasswdUserLevel : PosOfPttPasswdUserLevel+4])
	user.numLoginDays = binary.LittleEndian.Uint32(data[PosOfPttPasswdNumLoginDays : PosOfPttPasswdNumLoginDays+4])
	user.numPosts = binary.LittleEndian.Uint32(data[PosOfPttPasswdNumPosts : PosOfPttPasswdNumPosts+4])
	user.firstLogin = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfPttPasswdFirstLogin:PosOfPttPasswdFirstLogin+4])), 0)
	user.lastLogin = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfPttPasswdLastLogin:PosOfPttPasswdLastLogin+4])), 0)
	user.lastHost = newStringFormCString(data[PosOfPttPasswdLastHost : PosOfPttPasswdLastHost+PTT_IPV4LEN+1])

	user.money = int32(binary.LittleEndian.Uint32(data[PosOfPttPasswdMoney : PosOfPttPasswdMoney+4]))

	user.Email = newStringFormBig5UAOCString(data[PosOfPttPasswdEmail : PosOfPttPasswdEmail+PTT_EMAILSZ])
	user.Address = newStringFormBig5UAOCString(data[PosOfPttPasswdAddress : PosOfPttPasswdAddress+PTT_ADDRESSSZ])
	user.Justify = newStringFormBig5UAOCString(data[PosOfPttPasswdJustify : PosOfPttPasswdJustify+PTT_REGLEN+1])

	user.Over18 = data[PosOfPttPasswdOver18] != 0
	user.PagerUIType = data[PosOfPttPasswdPagerUiType]
	user.Pager = data[PosOfPttPasswdPager]
	user.Invisible = data[PosOfPttPasswdInvisible] != 0

	user.ExMailBox = binary.LittleEndian.Uint32(data[PosOfPttPasswdExMailBox : PosOfPttPasswdExMailBox+4])

	user.Career = newStringFormBig5UAOCString(data[PosOfPttPasswdCareer : PosOfPttPasswdCareer+40])
	user.Role = binary.LittleEndian.Uint32(data[PosOfPttPasswdRole : PosOfPttPasswdRole+4])
	user.LastSeen = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfPttPasswdLastSeen:PosOfPttPasswdLastSeen+4])), 0)
	user.TimeSetAngel = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfPttPasswdTimeSetAngel:PosOfPttPasswdTimeSetAngel+4])), 0)
	user.TimePlayAngel = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfPttPasswdTimePlayAngel:PosOfPttPasswdTimePlayAngel+4])), 0)

	user.LastSong = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfPttPasswdLastSong:PosOfPttPasswdLastSong+4])), 0)
	user.LoginView = binary.LittleEndian.Uint32(data[PosOfPttPasswdLoginView : PosOfPttPasswdLoginView+4])
	user.ViolateLaw = binary.LittleEndian.Uint16(data[PosOfPttPasswdLawCounter : PosOfPttPasswdLawCounter+2])

	user.Five.Win = binary.LittleEndian.Uint16(data[PosOfPttPasswdFiveWin : PosOfPttPasswdFiveWin+2])
	user.Five.Lose = binary.LittleEndian.Uint16(data[PosOfPttPasswdFiveLose : PosOfPttPasswdFiveLose+2])
	user.Five.Tie = binary.LittleEndian.Uint16(data[PosOfPttPasswdFiveTie : PosOfPttPasswdFiveTie+2])

	user.ChineseChess.Win = binary.LittleEndian.Uint16(data[PosOfPttPasswdChcWin : PosOfPttPasswdChcWin+2])
	user.ChineseChess.Lose = binary.LittleEndian.Uint16(data[PosOfPttPasswdChcLose : PosOfPttPasswdChcLose+2])
	user.ChineseChess.Tie = binary.LittleEndian.Uint16(data[PosOfPttPasswdChcTie : PosOfPttPasswdChcTie+2])

	user.Conn6.Win = binary.LittleEndian.Uint16(data[PosOfPttPasswdConn6Win : PosOfPttPasswdConn6Win+2])
	user.Conn6.Lose = binary.LittleEndian.Uint16(data[PosOfPttPasswdConn6Lose : PosOfPttPasswdConn6Lose+2])
	user.Conn6.Tie = binary.LittleEndian.Uint16(data[PosOfPttPasswdConn6Tie : PosOfPttPasswdConn6Tie+2])

	user.GoChess.Win = binary.LittleEndian.Uint16(data[PosOfPttPasswdGoWin : PosOfPttPasswdGoWin+2])
	user.GoChess.Lose = binary.LittleEndian.Uint16(data[PosOfPttPasswdGoLose : PosOfPttPasswdGoLose+2])
	user.GoChess.Tie = binary.LittleEndian.Uint16(data[PosOfPttPasswdGoTie : PosOfPttPasswdGoTie+2])

	user.DarkChess.Win = binary.LittleEndian.Uint16(data[PosOfPttPasswdDarkWin : PosOfPttPasswdDarkWin+2])
	user.DarkChess.Lose = binary.LittleEndian.Uint16(data[PosOfPttPasswdDarkLose : PosOfPttPasswdDarkLose+2])
	user.UaVersion = data[PosOfPttPasswdUaVersion]

	user.Signature = data[PosOfPttPasswdSignature]
	user.BadPost = data[PosOfPttPasswdBadPost]
	user.DarkChess.Tie = binary.LittleEndian.Uint16(data[PosOfPttPasswdDarkTie : PosOfPttPasswdDarkTie+2])
	user.MyAngel = newStringFormCString(data[PosOfPttPasswdMyAngel : PosOfPttPasswdMyAngel+PTT_IDLEN+1+1])

	user.ChessEloRating = binary.LittleEndian.Uint16(data[PosOfPttPasswdChessEloRating : PosOfPttPasswdChessEloRating+2])
	user.WithMe = binary.LittleEndian.Uint32(data[PosOfPttPasswdWithMe : PosOfPttPasswdWithMe+4])
	user.TimeRemoveBadPost = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfPttPasswdTimeRemoveBadPost:PosOfPttPasswdTimeRemoveBadPost+4])), 0)
	user.TimeViolateLaw = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfPttPasswdTimeViolateLaw:PosOfPttPasswdTimeViolateLaw+4])), 0)

	return user, nil
}

func (r *Userec) MarshalToByte() ([]byte, error) {
	ret := make([]byte, 512)

	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdVersion:PosOfPttPasswdVersion+4], r.Version)
	copy(ret[PosOfPttPasswdUserId:PosOfPttPasswdUserId+PTT_IDLEN+1], utf8ToBig5UAOString(r.userId))
	copy(ret[PosOfPttPasswdRealName:PosOfPttPasswdRealName+PTT_IDLEN+1], utf8ToBig5UAOString(r.realName))
	copy(ret[PosOfPttPasswdNickname:PosOfPttPasswdNickname+20], utf8ToBig5UAOString(r.nickname))
	copy(ret[PosOfPttPasswdPassword:PosOfPttPasswdPassword+20], utf8ToBig5UAOString(r.password))

	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdUserFlag:PosOfPttPasswdUserFlag+4], r.UserFlag)
	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdUserLevel:PosOfPttPasswdUserLevel+4], r.UserLevel)
	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdNumLoginDays:PosOfPttPasswdNumLoginDays+4], r.numLoginDays)
	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdNumPosts:PosOfPttPasswdNumPosts+4], r.numPosts)
	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdFirstLogin:PosOfPttPasswdFirstLogin+4], uint32(r.firstLogin.Unix()))
	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdLastLogin:PosOfPttPasswdLastLogin+4], uint32(r.lastLogin.Unix()))
	copy(ret[PosOfPttPasswdLastHost:PosOfPttPasswdLastHost+PTT_IPV4LEN+1], utf8ToBig5UAOString(r.lastHost))
	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdMoney:PosOfPttPasswdMoney+4], uint32(r.money))

	copy(ret[PosOfPttPasswdEmail:PosOfPttPasswdEmail+PTT_EMAILSZ], utf8ToBig5UAOString(r.Email))
	copy(ret[PosOfPttPasswdAddress:PosOfPttPasswdAddress+PTT_ADDRESSSZ], utf8ToBig5UAOString(r.Address))
	copy(ret[PosOfPttPasswdJustify:PosOfPttPasswdJustify+PTT_REGLEN], utf8ToBig5UAOString(r.Justify))

	if r.Over18 {
		ret[PosOfPttPasswdOver18] = 1
	} else {
		ret[PosOfPttPasswdOver18] = 0
	}

	ret[PosOfPttPasswdPagerUiType] = r.PagerUIType
	ret[PosOfPttPasswdPager] = r.Pager

	if r.Invisible {
		ret[PosOfPttPasswdInvisible] = 1
	} else {
		ret[PosOfPttPasswdInvisible] = 0
	}

	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdExMailBox:PosOfPttPasswdExMailBox+4], r.ExMailBox)

	copy(ret[PosOfPttPasswdCareer:PosOfPttPasswdCareer+44], utf8ToBig5UAOString(r.Career))

	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdLastSeen:PosOfPttPasswdLastSeen+4], uint32(r.LastSeen.Unix()))
	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdTimeSetAngel:PosOfPttPasswdTimeSetAngel+4], uint32(r.TimeSetAngel.Unix()))
	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdTimePlayAngel:PosOfPttPasswdTimePlayAngel+4], uint32(r.TimePlayAngel.Unix()))

	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdLastSong:PosOfPttPasswdLastSong+4], uint32(r.LastSong.Unix()))
	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdLoginView:PosOfPttPasswdLoginView+4], r.LoginView)
	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdLawCounter:PosOfPttPasswdLawCounter+2], r.ViolateLaw)

	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdFiveWin:PosOfPttPasswdFiveWin+2], r.Five.Win)
	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdFiveLose:PosOfPttPasswdFiveLose+2], r.Five.Lose)
	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdFiveTie:PosOfPttPasswdFiveTie+2], r.Five.Tie)

	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdChcWin:PosOfPttPasswdChcWin+2], r.ChineseChess.Win)
	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdChcLose:PosOfPttPasswdChcLose+2], r.ChineseChess.Lose)
	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdChcTie:PosOfPttPasswdChcTie+2], r.ChineseChess.Tie)

	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdConn6Win:PosOfPttPasswdConn6Win+2], r.Conn6.Win)
	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdConn6Lose:PosOfPttPasswdConn6Lose+2], r.Conn6.Lose)
	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdConn6Tie:PosOfPttPasswdConn6Tie+2], r.Conn6.Tie)

	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdGoWin:PosOfPttPasswdGoWin+2], r.GoChess.Win)
	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdGoLose:PosOfPttPasswdGoLose+2], r.GoChess.Lose)
	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdGoTie:PosOfPttPasswdGoTie+2], r.GoChess.Tie)

	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdDarkWin:PosOfPttPasswdDarkWin+2], r.DarkChess.Win)
	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdDarkLose:PosOfPttPasswdDarkLose+2], r.DarkChess.Lose)
	ret[PosOfPttPasswdUaVersion] = r.UaVersion

	ret[PosOfPttPasswdSignature] = r.Signature
	ret[PosOfPttPasswdBadPost] = r.BadPost
	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdDarkTie:PosOfPttPasswdDarkTie+2], r.DarkChess.Tie)
	copy(ret[PosOfPttPasswdMyAngel:PosOfPttPasswdMyAngel+PTT_IDLEN+1+1], utf8ToBig5UAOString(r.MyAngel))

	binary.LittleEndian.PutUint16(ret[PosOfPttPasswdChessEloRating:PosOfPttPasswdChessEloRating+2], r.ChessEloRating)
	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdWithMe:PosOfPttPasswdWithMe+4], r.WithMe)
	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdTimeRemoveBadPost:PosOfPttPasswdTimeRemoveBadPost+4], uint32(r.TimeRemoveBadPost.Unix()))
	binary.LittleEndian.PutUint32(ret[PosOfPttPasswdTimeViolateLaw:PosOfPttPasswdTimeViolateLaw+4], uint32(r.TimeViolateLaw.Unix()))

	return ret, nil
}
