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
	"github.com/Ptt-official-app/go-bbs/crypt"

	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const (
	PosOfPasswdVersion      = 0
	PosOfPasswdUserID       = PosOfPasswdVersion + 4
	PosOfPasswdRealName     = PosOfPasswdUserID + IDLength + 1
	PosOfPasswdNickname     = PosOfPasswdRealName + RealNameSize
	PosOfPasswdPassword     = PosOfPasswdNickname + NicknameSize
	PosOfPasswdUserFlag     = 1 + PosOfPasswdPassword + PasswordLength
	PosOfPasswdUserLevel    = 4 + PosOfPasswdUserFlag + 4
	PosOfPasswdNumLoginDays = PosOfPasswdUserLevel + 4
	PosOfPasswdNumPosts     = PosOfPasswdNumLoginDays + 4
	PosOfPasswdFirstLogin   = PosOfPasswdNumPosts + 4
	PosOfPasswdLastLogin    = PosOfPasswdFirstLogin + 4
	PosOfPasswdLastHost     = PosOfPasswdLastLogin + 4
	PosOfPasswdMoney        = PosOfPasswdLastHost + IPV4Length + 1
	PosOfPasswdEmail        = 4 + PosOfPasswdMoney + 4
	PosOfPasswdAddress      = PosOfPasswdEmail + EmailSize
	PosOfPasswdJustify      = PosOfPasswdAddress + AddressSize
	PosOfPasswdOver18       = 3 + PosOfPasswdJustify + RegistrationLength + 1
	PosOfPasswdPagerUIType  = PosOfPasswdOver18 + 1
	PosOfPasswdPager        = PosOfPasswdPagerUIType + 1
	PosOfPasswdInvisible    = PosOfPasswdPager + 1
	PosOfPasswdExMailBox    = 2 + PosOfPasswdInvisible + 1

	PosOfPasswdCareer        = 4 + PosOfPasswdExMailBox + 4
	PosOfPasswdRole          = 20 + 4 + 44 + PosOfPasswdCareer + CareerSize
	PosOfPasswdLastSeen      = PosOfPasswdRole + 4
	PosOfPasswdTimeSetAngel  = PosOfPasswdLastSeen + 4
	PosOfPasswdTimePlayAngel = PosOfPasswdTimeSetAngel + 4

	PosOfPasswdLastSong  = PosOfPasswdTimePlayAngel + 4
	PosOfPasswdLoginView = PosOfPasswdLastSong + 4

	PosOfPasswdLawCounter = 1 + 1 + PosOfPasswdLoginView + 4
	PosOfPasswdFiveWin    = PosOfPasswdLawCounter + 2
	PosOfPasswdFiveLose   = PosOfPasswdFiveWin + 2
	PosOfPasswdFiveTie    = PosOfPasswdFiveLose + 2
	PosOfPasswdChcWin     = PosOfPasswdFiveTie + 2
	PosOfPasswdChcLose    = PosOfPasswdChcWin + 2
	PosOfPasswdChcTie     = PosOfPasswdChcLose + 2
	PosOfPasswdConn6Win   = PosOfPasswdChcTie + 2
	PosOfPasswdConn6Lose  = PosOfPasswdConn6Win + 2
	PosOfPasswdConn6Tie   = PosOfPasswdConn6Lose + 2
	PosOfPasswdGoWin      = 2 + PosOfPasswdConn6Tie + 2
	PosOfPasswdGoLose     = PosOfPasswdGoWin + 2
	PosOfPasswdGoTie      = PosOfPasswdGoLose + 2
	PosOfPasswdDarkWin    = PosOfPasswdGoTie + 2
	PosOfPasswdDarkLose   = PosOfPasswdDarkWin + 2
	PosOfPasswdUaVersion  = PosOfPasswdDarkLose + 2

	PosOfPasswdSignature = PosOfPasswdUaVersion + 1
	PosOfPasswdBadPost   = 1 + PosOfPasswdSignature + 1
	PosOfPasswdDarkTie   = PosOfPasswdBadPost + 1
	PosOfPasswdMyAngel   = PosOfPasswdDarkTie + 2

	PosOfPasswdChessEloRating    = 1 + PosOfPasswdMyAngel + IDLength + 1
	PosOfPasswdWithMe            = PosOfPasswdChessEloRating + 2
	PosOfPasswdTimeRemoveBadPost = PosOfPasswdWithMe + 4
	PosOfPasswdTimeViolateLaw    = PosOfPasswdTimeRemoveBadPost + 4
)

// https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h

type UserecGameScore struct {
	Win  uint16
	Lose uint16
	Tie  uint16
}

type Userec struct {
	Version  uint32 // Magic Number
	userID   string // 使用者帳號，或稱使用者 ID
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

func (u *Userec) UserID() string { return u.userID }

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
	user.Version = binary.LittleEndian.Uint32(data[PosOfPasswdVersion : PosOfPasswdVersion+4])
	user.userID = newStringFormCString(data[PosOfPasswdUserID : PosOfPasswdUserID+IDLength+1])
	user.realName = newStringFormBig5UAOCString(data[PosOfPasswdRealName : PosOfPasswdRealName+RealNameSize])
	user.nickname = newStringFormBig5UAOCString(data[PosOfPasswdNickname : PosOfPasswdNickname+NicknameSize])
	user.password = newStringFormCString(data[PosOfPasswdPassword : PosOfPasswdPassword+PasswordLength])

	user.UserFlag = binary.LittleEndian.Uint32(data[PosOfPasswdUserFlag : PosOfPasswdUserFlag+4])
	user.UserLevel = binary.LittleEndian.Uint32(data[PosOfPasswdUserLevel : PosOfPasswdUserLevel+4])
	user.numLoginDays = binary.LittleEndian.Uint32(data[PosOfPasswdNumLoginDays : PosOfPasswdNumLoginDays+4])
	user.numPosts = binary.LittleEndian.Uint32(data[PosOfPasswdNumPosts : PosOfPasswdNumPosts+4])
	user.firstLogin = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfPasswdFirstLogin:PosOfPasswdFirstLogin+4])), 0)
	user.lastLogin = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfPasswdLastLogin:PosOfPasswdLastLogin+4])), 0)
	user.lastHost = newStringFormCString(data[PosOfPasswdLastHost : PosOfPasswdLastHost+IPV4Length+1])

	user.money = int32(binary.LittleEndian.Uint32(data[PosOfPasswdMoney : PosOfPasswdMoney+4]))

	user.Email = newStringFormBig5UAOCString(data[PosOfPasswdEmail : PosOfPasswdEmail+EmailSize])
	user.Address = newStringFormBig5UAOCString(data[PosOfPasswdAddress : PosOfPasswdAddress+AddressSize])
	user.Justify = newStringFormBig5UAOCString(data[PosOfPasswdJustify : PosOfPasswdJustify+RegistrationLength+1])

	user.Over18 = data[PosOfPasswdOver18] != 0
	user.PagerUIType = data[PosOfPasswdPagerUIType]
	user.Pager = data[PosOfPasswdPager]
	user.Invisible = data[PosOfPasswdInvisible] != 0

	user.ExMailBox = binary.LittleEndian.Uint32(data[PosOfPasswdExMailBox : PosOfPasswdExMailBox+4])

	user.Career = newStringFormBig5UAOCString(data[PosOfPasswdCareer : PosOfPasswdCareer+CareerSize])
	user.Role = binary.LittleEndian.Uint32(data[PosOfPasswdRole : PosOfPasswdRole+4])
	user.LastSeen = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfPasswdLastSeen:PosOfPasswdLastSeen+4])), 0)
	user.TimeSetAngel = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfPasswdTimeSetAngel:PosOfPasswdTimeSetAngel+4])), 0)
	user.TimePlayAngel = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfPasswdTimePlayAngel:PosOfPasswdTimePlayAngel+4])), 0)

	user.LastSong = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfPasswdLastSong:PosOfPasswdLastSong+4])), 0)
	user.LoginView = binary.LittleEndian.Uint32(data[PosOfPasswdLoginView : PosOfPasswdLoginView+4])
	user.ViolateLaw = binary.LittleEndian.Uint16(data[PosOfPasswdLawCounter : PosOfPasswdLawCounter+2])

	user.Five.Win = binary.LittleEndian.Uint16(data[PosOfPasswdFiveWin : PosOfPasswdFiveWin+2])
	user.Five.Lose = binary.LittleEndian.Uint16(data[PosOfPasswdFiveLose : PosOfPasswdFiveLose+2])
	user.Five.Tie = binary.LittleEndian.Uint16(data[PosOfPasswdFiveTie : PosOfPasswdFiveTie+2])

	user.ChineseChess.Win = binary.LittleEndian.Uint16(data[PosOfPasswdChcWin : PosOfPasswdChcWin+2])
	user.ChineseChess.Lose = binary.LittleEndian.Uint16(data[PosOfPasswdChcLose : PosOfPasswdChcLose+2])
	user.ChineseChess.Tie = binary.LittleEndian.Uint16(data[PosOfPasswdChcTie : PosOfPasswdChcTie+2])

	user.Conn6.Win = binary.LittleEndian.Uint16(data[PosOfPasswdConn6Win : PosOfPasswdConn6Win+2])
	user.Conn6.Lose = binary.LittleEndian.Uint16(data[PosOfPasswdConn6Lose : PosOfPasswdConn6Lose+2])
	user.Conn6.Tie = binary.LittleEndian.Uint16(data[PosOfPasswdConn6Tie : PosOfPasswdConn6Tie+2])

	user.GoChess.Win = binary.LittleEndian.Uint16(data[PosOfPasswdGoWin : PosOfPasswdGoWin+2])
	user.GoChess.Lose = binary.LittleEndian.Uint16(data[PosOfPasswdGoLose : PosOfPasswdGoLose+2])
	user.GoChess.Tie = binary.LittleEndian.Uint16(data[PosOfPasswdGoTie : PosOfPasswdGoTie+2])

	user.DarkChess.Win = binary.LittleEndian.Uint16(data[PosOfPasswdDarkWin : PosOfPasswdDarkWin+2])
	user.DarkChess.Lose = binary.LittleEndian.Uint16(data[PosOfPasswdDarkLose : PosOfPasswdDarkLose+2])
	user.UaVersion = data[PosOfPasswdUaVersion]

	user.Signature = data[PosOfPasswdSignature]
	user.BadPost = data[PosOfPasswdBadPost]
	user.DarkChess.Tie = binary.LittleEndian.Uint16(data[PosOfPasswdDarkTie : PosOfPasswdDarkTie+2])
	user.MyAngel = newStringFormCString(data[PosOfPasswdMyAngel : PosOfPasswdMyAngel+IDLength+1+1])

	user.ChessEloRating = binary.LittleEndian.Uint16(data[PosOfPasswdChessEloRating : PosOfPasswdChessEloRating+2])
	user.WithMe = binary.LittleEndian.Uint32(data[PosOfPasswdWithMe : PosOfPasswdWithMe+4])
	user.TimeRemoveBadPost = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfPasswdTimeRemoveBadPost:PosOfPasswdTimeRemoveBadPost+4])), 0)
	user.TimeViolateLaw = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfPasswdTimeViolateLaw:PosOfPasswdTimeViolateLaw+4])), 0)

	return user, nil
}

func (u *Userec) MarshalToByte() ([]byte, error) {
	ret := make([]byte, 512)

	binary.LittleEndian.PutUint32(ret[PosOfPasswdVersion:PosOfPasswdVersion+4], u.Version)
	copy(ret[PosOfPasswdUserID:PosOfPasswdUserID+IDLength+1], utf8ToBig5UAOString(u.userID))
	copy(ret[PosOfPasswdRealName:PosOfPasswdRealName+IDLength+1], utf8ToBig5UAOString(u.realName))
	copy(ret[PosOfPasswdNickname:PosOfPasswdNickname+NicknameSize], utf8ToBig5UAOString(u.nickname))
	copy(ret[PosOfPasswdPassword:PosOfPasswdPassword+PasswordLength], utf8ToBig5UAOString(u.password))

	binary.LittleEndian.PutUint32(ret[PosOfPasswdUserFlag:PosOfPasswdUserFlag+4], u.UserFlag)
	binary.LittleEndian.PutUint32(ret[PosOfPasswdUserLevel:PosOfPasswdUserLevel+4], u.UserLevel)
	binary.LittleEndian.PutUint32(ret[PosOfPasswdNumLoginDays:PosOfPasswdNumLoginDays+4], u.numLoginDays)
	binary.LittleEndian.PutUint32(ret[PosOfPasswdNumPosts:PosOfPasswdNumPosts+4], u.numPosts)
	binary.LittleEndian.PutUint32(ret[PosOfPasswdFirstLogin:PosOfPasswdFirstLogin+4], uint32(u.firstLogin.Unix()))
	binary.LittleEndian.PutUint32(ret[PosOfPasswdLastLogin:PosOfPasswdLastLogin+4], uint32(u.lastLogin.Unix()))
	copy(ret[PosOfPasswdLastHost:PosOfPasswdLastHost+IPV4Length+1], utf8ToBig5UAOString(u.lastHost))
	binary.LittleEndian.PutUint32(ret[PosOfPasswdMoney:PosOfPasswdMoney+4], uint32(u.money))

	copy(ret[PosOfPasswdEmail:PosOfPasswdEmail+EmailSize], utf8ToBig5UAOString(u.Email))
	copy(ret[PosOfPasswdAddress:PosOfPasswdAddress+AddressSize], utf8ToBig5UAOString(u.Address))
	copy(ret[PosOfPasswdJustify:PosOfPasswdJustify+RegistrationLength], utf8ToBig5UAOString(u.Justify))

	if u.Over18 {
		ret[PosOfPasswdOver18] = 1
	} else {
		ret[PosOfPasswdOver18] = 0
	}

	ret[PosOfPasswdPagerUIType] = u.PagerUIType
	ret[PosOfPasswdPager] = u.Pager

	if u.Invisible {
		ret[PosOfPasswdInvisible] = 1
	} else {
		ret[PosOfPasswdInvisible] = 0
	}

	binary.LittleEndian.PutUint32(ret[PosOfPasswdExMailBox:PosOfPasswdExMailBox+4], u.ExMailBox)

	copy(ret[PosOfPasswdCareer:PosOfPasswdCareer+CareerSize], utf8ToBig5UAOString(u.Career))

	binary.LittleEndian.PutUint32(ret[PosOfPasswdLastSeen:PosOfPasswdLastSeen+4], uint32(u.LastSeen.Unix()))
	binary.LittleEndian.PutUint32(ret[PosOfPasswdTimeSetAngel:PosOfPasswdTimeSetAngel+4], uint32(u.TimeSetAngel.Unix()))
	binary.LittleEndian.PutUint32(ret[PosOfPasswdTimePlayAngel:PosOfPasswdTimePlayAngel+4], uint32(u.TimePlayAngel.Unix()))

	binary.LittleEndian.PutUint32(ret[PosOfPasswdLastSong:PosOfPasswdLastSong+4], uint32(u.LastSong.Unix()))
	binary.LittleEndian.PutUint32(ret[PosOfPasswdLoginView:PosOfPasswdLoginView+4], u.LoginView)
	binary.LittleEndian.PutUint16(ret[PosOfPasswdLawCounter:PosOfPasswdLawCounter+2], u.ViolateLaw)

	binary.LittleEndian.PutUint16(ret[PosOfPasswdFiveWin:PosOfPasswdFiveWin+2], u.Five.Win)
	binary.LittleEndian.PutUint16(ret[PosOfPasswdFiveLose:PosOfPasswdFiveLose+2], u.Five.Lose)
	binary.LittleEndian.PutUint16(ret[PosOfPasswdFiveTie:PosOfPasswdFiveTie+2], u.Five.Tie)

	binary.LittleEndian.PutUint16(ret[PosOfPasswdChcWin:PosOfPasswdChcWin+2], u.ChineseChess.Win)
	binary.LittleEndian.PutUint16(ret[PosOfPasswdChcLose:PosOfPasswdChcLose+2], u.ChineseChess.Lose)
	binary.LittleEndian.PutUint16(ret[PosOfPasswdChcTie:PosOfPasswdChcTie+2], u.ChineseChess.Tie)

	binary.LittleEndian.PutUint16(ret[PosOfPasswdConn6Win:PosOfPasswdConn6Win+2], u.Conn6.Win)
	binary.LittleEndian.PutUint16(ret[PosOfPasswdConn6Lose:PosOfPasswdConn6Lose+2], u.Conn6.Lose)
	binary.LittleEndian.PutUint16(ret[PosOfPasswdConn6Tie:PosOfPasswdConn6Tie+2], u.Conn6.Tie)

	binary.LittleEndian.PutUint16(ret[PosOfPasswdGoWin:PosOfPasswdGoWin+2], u.GoChess.Win)
	binary.LittleEndian.PutUint16(ret[PosOfPasswdGoLose:PosOfPasswdGoLose+2], u.GoChess.Lose)
	binary.LittleEndian.PutUint16(ret[PosOfPasswdGoTie:PosOfPasswdGoTie+2], u.GoChess.Tie)

	binary.LittleEndian.PutUint16(ret[PosOfPasswdDarkWin:PosOfPasswdDarkWin+2], u.DarkChess.Win)
	binary.LittleEndian.PutUint16(ret[PosOfPasswdDarkLose:PosOfPasswdDarkLose+2], u.DarkChess.Lose)
	ret[PosOfPasswdUaVersion] = u.UaVersion

	ret[PosOfPasswdSignature] = u.Signature
	ret[PosOfPasswdBadPost] = u.BadPost
	binary.LittleEndian.PutUint16(ret[PosOfPasswdDarkTie:PosOfPasswdDarkTie+2], u.DarkChess.Tie)
	copy(ret[PosOfPasswdMyAngel:PosOfPasswdMyAngel+IDLength+1+1], utf8ToBig5UAOString(u.MyAngel))

	binary.LittleEndian.PutUint16(ret[PosOfPasswdChessEloRating:PosOfPasswdChessEloRating+2], u.ChessEloRating)
	binary.LittleEndian.PutUint32(ret[PosOfPasswdWithMe:PosOfPasswdWithMe+4], u.WithMe)
	binary.LittleEndian.PutUint32(ret[PosOfPasswdTimeRemoveBadPost:PosOfPasswdTimeRemoveBadPost+4], uint32(u.TimeRemoveBadPost.Unix()))
	binary.LittleEndian.PutUint32(ret[PosOfPasswdTimeViolateLaw:PosOfPasswdTimeViolateLaw+4], uint32(u.TimeViolateLaw.Unix()))

	return ret, nil
}
