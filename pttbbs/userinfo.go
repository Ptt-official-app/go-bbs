package pttbbs

import (
	"encoding/binary"
	"io"
	"log"
	"os"
	"time"
)

const (
	PosOfUserInfoUID      = 0
	PosOfUserInfoPID      = PosOfUserInfoUID + 4
	PosOfUserInfoSockAddr = PosOfUserInfoPID + 4

	PosOfUserInfoUserLevel  = PosOfUserInfoSockAddr + 4
	PosOfUserInfoUserID     = PosOfUserInfoUserLevel + 4
	PosOfUserInfoNickname   = PosOfUserInfoUserID + IDLength + 1
	PosOfUserInfoFrom       = PosOfUserInfoNickname + NicknameSize
	PosOfUserInfoFromIP     = PosOfUserInfoFrom + MachineNameLength
	PosOfUserInfoDarkWin    = PosOfUserInfoFromIP + 4
	PosOfUserInfoDarkLose   = PosOfUserInfoDarkWin + 2
	PosOfUserInfoAngelPause = PosOfUserInfoDarkLose + 2 + 1 // gap_0
	PosOfUserInfoDarkTie    = PosOfUserInfoAngelPause + 1

	PosOfUserInfoFriendTotal = PosOfUserInfoDarkTie + 2
	PosOfUserInfoNumFriends  = PosOfUserInfoFriendTotal + 4

	PosOfUserInfoMyFriend     = PosOfUserInfoNumFriends + 2 + 2             // _unused3
	PosOfUserInfoFriendOnline = PosOfUserInfoMyFriend + MaxFriend*4 + 4     // gap_1
	PosOfUserInfoReject       = PosOfUserInfoFriendOnline + MaxFriend*4 + 4 // gap_2

	PosOfUserInfoMsgCount = PosOfUserInfoReject + MaxReject*4 + 4 // gap_3
	PosOfUserInfoMsgs     = PosOfUserInfoMsgCount + 1 + 3         // _unused4

	PosOfUserInfoBirth     = PosOfUserInfoMsgs + MaxMsgs*100 + 100 // gap_4 and 100 is the size of msgque_t when run on my machine
	PosOfUserInfoActive    = PosOfUserInfoBirth + 1
	PosOfUserInfoInvisible = PosOfUserInfoActive + 1
	PosOfUserInfoMode      = PosOfUserInfoInvisible + 1
	PosOfUserInfoPager     = PosOfUserInfoMode + 1

	PosOfUserInfoConn6Win  = PosOfUserInfoPager + 1 + 1 // unused5
	PosOfUserInfoLastAct   = PosOfUserInfoConn6Win + 2
	PosOfUserInfoAlerts    = PosOfUserInfoLastAct + 4
	PosOfUserInfoConn6Lose = PosOfUserInfoAlerts + 1 + 1 // unused_mind

	PosOfUserInfoSig        = PosOfUserInfoConn6Lose + 2 + 1 // unused_mind2
	PosOfUserInfoConn6Tie   = PosOfUserInfoSig + 1
	PosOfUserInfoDestUID    = PosOfUserInfoConn6Tie + 2
	PosOfUserInfoDestUIP    = PosOfUserInfoDestUID + 4
	PosOfUserInfoSockActive = PosOfUserInfoDestUIP + 4

	PosOfUserInfoInChat = PosOfUserInfoSockActive + 1
	PosOfUserInfoChatID = PosOfUserInfoInChat + 1

	PosOfUserInfoLockMode = PosOfUserInfoChatID + ChatIDLength
	PosOfUserInfoTurn     = PosOfUserInfoLockMode + 1
	PosOfUserInfoMateID   = PosOfUserInfoTurn + 1
	PosOfUserInfoColor    = PosOfUserInfoMateID + IDLength + 1

	PosOfUserInfoFiveWin        = PosOfUserInfoColor + 1
	PosOfUserInfoFiveLose       = PosOfUserInfoFiveWin + 2
	PosOfUserInfoFiveTie        = PosOfUserInfoFiveLose + 2
	PosOfUserInfoChcWin         = PosOfUserInfoFiveTie + 2
	PosOfUserInfoChcLose        = PosOfUserInfoChcWin + 2
	PosOfUserInfoChcTie         = PosOfUserInfoFiveLose + 2
	PosOfUserInfoChessEloRating = PosOfUserInfoChcTie + 2
	PosOfUserInfoGoWin          = PosOfUserInfoChessEloRating + 2
	PosOfUserInfoGoLose         = PosOfUserInfoGoWin + 2
	PosOfUserInfoGoTie          = PosOfUserInfoGoLose + 2

	PosOfUserInfoWithMe = PosOfUserInfoGoTie + 2
	PosOfUserInfoBrcID  = PosOfUserInfoWithMe + 4

	PosOfUserInfoWBTime = PosOfUserInfoBrcID + 4
)

// UserInfoCache is `userinfo_t` in pttstruct.h, all of those object will be store in shared memory
type UserInfoCache struct {
	uid      int32
	pid      int32
	SockAddr int32

	UserLevel uint32
	userID    string
	nickname  string
	from      string // machine name the user called in from
	fromIP    uint32

	DarkChess  GameScore
	AngelPause uint8

	FriendTotal int32
	NumFriends  int16

	MyFriend     [MaxFriend]int32
	FriendOnline [MaxFriend]uint32

	Reject [MaxReject]int32

	MsgCount uint8
	Msgs     [MaxMsgs]MessageQueueCache

	Birth     bool // whether it's birthday today
	Active    bool
	Invisible bool
	Mode      uint8
	Pager     bool

	Conn6   GameScore
	LastAct time.Time
	Alerts  uint8 // mail alert, passwd update ...

	Signature  uint8
	DestUID    int32
	DestUIP    int32
	SockActive uint8

	InChat uint8
	ChatID string

	LockMode uint8
	Turn     bool
	MateID   string
	Color    uint8

	// game record
	Five           GameScore
	ChineseChess   GameScore
	ChessEloRating uint16
	GoChess        GameScore

	WithMe uint32
	brcID  uint32

	wbTime time.Time // not sure if this is present
}

// function used for testing unmarshal and marshal function
func OpenUserInfoCache(filename string) ([]*UserInfoCache, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	ret := []*UserInfoCache{}

	for {
		buf := make([]byte, 3484)

		_, err := file.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			return nil, err
		}

		cache := UnmarshalUserInfo(buf)

		ret = append(ret, cache)
	}

	return ret, nil
}

func UnmarshalUserInfo(data []byte) *UserInfoCache {
	cache := &UserInfoCache{}

	cache.uid = int32(binary.LittleEndian.Uint32(data[PosOfUserInfoUID : PosOfUserInfoUID+4]))
	cache.pid = int32(binary.LittleEndian.Uint32(data[PosOfUserInfoPID : PosOfUserInfoPID+4]))
	cache.SockAddr = int32(binary.LittleEndian.Uint32(data[PosOfUserInfoSockAddr : PosOfUserInfoSockAddr+4]))

	cache.UserLevel = binary.LittleEndian.Uint32(data[PosOfUserInfoUserLevel : PosOfUserInfoUserLevel+4])
	cache.userID = newStringFormCString(data[PosOfUserInfoUserID : PosOfUserInfoUserID+IDLength+1])
	cache.nickname = newStringFormBig5UAOCString(data[PosOfUserInfoNickname : PosOfUserInfoNickname+NicknameSize])
	cache.from = newStringFormCString(data[PosOfUserInfoFrom : PosOfUserInfoFrom+MachineNameLength])
	cache.fromIP = binary.LittleEndian.Uint32(data[PosOfUserInfoFromIP : PosOfUserInfoFromIP+4])
	cache.DarkChess.Win = binary.LittleEndian.Uint16(data[PosOfUserInfoDarkWin : PosOfUserInfoDarkWin+2])
	cache.DarkChess.Lose = binary.LittleEndian.Uint16(data[PosOfUserInfoDarkLose : PosOfUserInfoDarkLose+2])
	cache.AngelPause = data[PosOfUserInfoAngelPause]
	cache.DarkChess.Tie = binary.LittleEndian.Uint16(data[PosOfUserInfoDarkTie : PosOfUserInfoDarkTie+2])

	cache.FriendTotal = int32(binary.LittleEndian.Uint32(data[PosOfUserInfoFriendTotal : PosOfUserInfoFriendTotal+4]))
	cache.NumFriends = int16(binary.LittleEndian.Uint16(data[PosOfUserInfoNumFriends : PosOfUserInfoNumFriends+2]))

	for idx := range cache.MyFriend {
		cache.MyFriend[idx] = int32(binary.LittleEndian.Uint32(data[PosOfUserInfoMyFriend+4*idx : PosOfUserInfoMyFriend+4*idx+4]))
	}

	for idx := range cache.FriendOnline {
		cache.FriendOnline[idx] = binary.LittleEndian.Uint32(data[PosOfUserInfoFriendOnline+4*idx : PosOfUserInfoFriendOnline+4*idx+4])
	}

	for idx := range cache.Reject {
		cache.Reject[idx] = int32(binary.LittleEndian.Uint32(data[PosOfUserInfoReject+4*idx : PosOfUserInfoReject+4*idx+4]))
	}

	cache.MsgCount = data[PosOfUserInfoMsgCount]

	for idx := range cache.Msgs {
		// Todo: replace it with the Unmarshal for MessageQueueCache
		cache.Msgs[idx] = data[PosOfUserInfoMsgs+100*idx : PosOfUserInfoMsgs+100*idx+100]
	}

	cache.Birth = (data[PosOfUserInfoBirth] != 0)
	cache.Active = (data[PosOfUserInfoActive] != 0)
	cache.Invisible = (data[PosOfUserInfoInvisible] != 0)
	cache.Mode = data[PosOfUserInfoMode]
	cache.Pager = (data[PosOfUserInfoPager] != 0)

	cache.Conn6.Win = binary.LittleEndian.Uint16(data[PosOfUserInfoConn6Win : PosOfUserInfoConn6Win+2])
	cache.LastAct = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfUserInfoLastAct:PosOfUserInfoLastAct+4])), 0)
	cache.Alerts = data[PosOfUserInfoAlerts]
	cache.Conn6.Lose = binary.LittleEndian.Uint16(data[PosOfUserInfoConn6Lose : PosOfUserInfoConn6Lose+2])

	cache.Signature = data[PosOfUserInfoSig]
	cache.Conn6.Tie = binary.LittleEndian.Uint16(data[PosOfUserInfoConn6Tie : PosOfUserInfoConn6Tie+2])

	cache.DestUID = int32(binary.LittleEndian.Uint32(data[PosOfUserInfoDestUID : PosOfUserInfoDestUID+4]))
	cache.DestUIP = int32(binary.LittleEndian.Uint32(data[PosOfUserInfoDestUIP : PosOfUserInfoDestUIP+4]))
	cache.SockActive = data[PosOfUserInfoSockActive]

	cache.InChat = data[PosOfUserInfoInChat]
	cache.ChatID = newStringFormCString(data[PosOfUserInfoChatID : PosOfUserInfoChatID+ChatIDLength])

	cache.LockMode = data[PosOfUserInfoLockMode]
	cache.Turn = (data[PosOfUserInfoTurn] != 0)

	cache.MateID = newStringFormBig5UAOCString(data[PosOfUserInfoMateID : PosOfUserInfoMateID+IDLength+1])
	cache.Color = data[PosOfUserInfoColor]

	cache.Five.Win = binary.LittleEndian.Uint16(data[PosOfUserInfoFiveWin : PosOfUserInfoFiveWin+2])
	cache.Five.Lose = binary.LittleEndian.Uint16(data[PosOfUserInfoFiveLose : PosOfUserInfoFiveLose+2])
	cache.Five.Tie = binary.LittleEndian.Uint16(data[PosOfUserInfoFiveTie : PosOfUserInfoFiveTie+2])

	cache.ChineseChess.Win = binary.LittleEndian.Uint16(data[PosOfUserInfoChcWin : PosOfUserInfoChcWin+2])
	cache.ChineseChess.Lose = binary.LittleEndian.Uint16(data[PosOfUserInfoChcLose : PosOfUserInfoChcLose+2])
	cache.ChineseChess.Tie = binary.LittleEndian.Uint16(data[PosOfUserInfoChcTie : PosOfUserInfoChcTie+2])

	cache.ChessEloRating = binary.LittleEndian.Uint16(data[PosOfUserInfoChessEloRating : PosOfUserInfoChessEloRating+2])

	cache.GoChess.Win = binary.LittleEndian.Uint16(data[PosOfUserInfoGoWin : PosOfUserInfoGoWin+2])
	cache.GoChess.Lose = binary.LittleEndian.Uint16(data[PosOfUserInfoGoLose : PosOfUserInfoGoLose+2])
	cache.GoChess.Tie = binary.LittleEndian.Uint16(data[PosOfUserInfoGoTie : PosOfUserInfoGoTie+2])

	cache.WithMe = binary.LittleEndian.Uint32(data[PosOfUserInfoWithMe : PosOfUserInfoWithMe+4])
	cache.brcID = binary.LittleEndian.Uint32(data[PosOfUserInfoBrcID : PosOfUserInfoBrcID+4])

	cache.wbTime = time.Unix(int64(binary.LittleEndian.Uint32(data[PosOfUserInfoWBTime:PosOfUserInfoWBTime+4])), 0)
	return cache
}

func (u *UserInfoCache) MarshalBinary() []byte {
	ret := make([]byte, 3484)

	binary.LittleEndian.PutUint32(ret[PosOfUserInfoUID:PosOfUserInfoUID+4], uint32(u.uid))
	binary.LittleEndian.PutUint32(ret[PosOfUserInfoPID:PosOfUserInfoPID+4], uint32(u.pid))
	binary.LittleEndian.PutUint32(ret[PosOfUserInfoSockAddr:PosOfUserInfoSockAddr+4], uint32(u.SockAddr))

	binary.LittleEndian.PutUint32(ret[PosOfUserInfoUserLevel:PosOfUserInfoUserLevel+4], u.UserLevel)

	copy(ret[PosOfUserInfoUserID:PosOfUserInfoUserID+IDLength+1], utf8ToBig5UAOString(u.userID))
	copy(ret[PosOfUserInfoNickname:PosOfUserInfoNickname+NicknameSize], utf8ToBig5UAOString(u.nickname))
	copy(ret[PosOfUserInfoFrom:PosOfUserInfoFrom+MachineNameLength], utf8ToBig5UAOString(u.from))

	binary.LittleEndian.PutUint32(ret[PosOfUserInfoFromIP:PosOfUserInfoFromIP+4], u.fromIP)
	binary.LittleEndian.PutUint16(ret[PosOfUserInfoDarkWin:PosOfUserInfoDarkWin+2], u.DarkChess.Win)
	binary.LittleEndian.PutUint16(ret[PosOfUserInfoDarkLose:PosOfUserInfoDarkLose+2], u.DarkChess.Lose)

	ret[PosOfUserInfoAngelPause] = u.AngelPause

	binary.LittleEndian.PutUint16(ret[PosOfUserInfoDarkTie:PosOfUserInfoDarkTie+2], u.DarkChess.Tie)

	binary.LittleEndian.PutUint32(ret[PosOfUserInfoFriendTotal:PosOfUserInfoFriendTotal+4], uint32(u.FriendTotal))
	binary.LittleEndian.PutUint16(ret[PosOfUserInfoNumFriends:PosOfUserInfoNumFriends+2], uint16(u.NumFriends))

	for idx, val := range u.MyFriend {
		binary.LittleEndian.PutUint32(ret[PosOfUserInfoMyFriend+4*idx:PosOfUserInfoMyFriend+4*idx+4], uint32(val))
	}

	for idx, val := range u.FriendOnline {
		binary.LittleEndian.PutUint32(ret[PosOfUserInfoFriendOnline+4*idx:PosOfUserInfoFriendOnline+4*idx+4], uint32(val))
	}

	for idx, val := range u.Reject {
		binary.LittleEndian.PutUint32(ret[PosOfUserInfoReject+4*idx:PosOfUserInfoReject+4*idx+4], uint32(val))
	}

	ret[PosOfUserInfoMsgCount] = u.MsgCount

	for idx, val := range u.Msgs {
		// Todo: replace it with the MarshalBinary for MessageQueueCache
		copy(ret[PosOfUserInfoMsgs+100*idx:PosOfUserInfoMsgs+100*idx+100], val)
	}

	if u.Birth {
		ret[PosOfUserInfoBirth] = 1
	} else {
		ret[PosOfUserInfoBirth] = 0
	}

	if u.Active {
		ret[PosOfUserInfoActive] = 1
	} else {
		ret[PosOfUserInfoActive] = 0
	}
	if u.Invisible {
		ret[PosOfUserInfoInvisible] = 1
	} else {
		ret[PosOfUserInfoInvisible] = 0
	}

	ret[PosOfUserInfoMode] = u.Mode

	if u.Pager {
		ret[PosOfUserInfoPager] = 1
	} else {
		ret[PosOfUserInfoPager] = 0
	}

	binary.LittleEndian.PutUint16(ret[PosOfUserInfoConn6Win:PosOfUserInfoConn6Win+2], u.Conn6.Win)

	binary.LittleEndian.PutUint32(ret[PosOfUserInfoLastAct:PosOfUserInfoLastAct+4], uint32(u.LastAct.Unix()))
	ret[PosOfUserInfoAlerts] = u.Alerts
	binary.LittleEndian.PutUint16(ret[PosOfUserInfoConn6Lose:PosOfUserInfoConn6Lose+2], u.Conn6.Lose)

	ret[PosOfUserInfoSig] = u.Signature
	binary.LittleEndian.PutUint16(ret[PosOfUserInfoConn6Tie:PosOfUserInfoConn6Tie+2], u.Conn6.Tie)

	binary.LittleEndian.PutUint32(ret[PosOfUserInfoDestUID:PosOfUserInfoDestUID+4], uint32(u.DestUID))
	binary.LittleEndian.PutUint32(ret[PosOfUserInfoDestUIP:PosOfUserInfoDestUIP+4], uint32(u.DestUIP))

	ret[PosOfUserInfoSockActive] = u.SockActive
	ret[PosOfUserInfoInChat] = u.InChat
	copy(ret[PosOfUserInfoChatID:PosOfUserInfoChatID+ChatIDLength], utf8ToBig5UAOString(u.ChatID))

	ret[PosOfUserInfoLockMode] = u.LockMode

	if u.Turn {
		ret[PosOfUserInfoTurn] = 1
	} else {
		ret[PosOfUserInfoTurn] = 0
	}

	copy(ret[PosOfUserInfoMateID:PosOfUserInfoMateID+IDLength+1], utf8ToBig5UAOString(u.MateID))

	ret[PosOfUserInfoColor] = u.Color

	binary.LittleEndian.PutUint16(ret[PosOfUserInfoFiveWin:PosOfUserInfoFiveWin+2], u.Five.Win)
	binary.LittleEndian.PutUint16(ret[PosOfUserInfoFiveLose:PosOfUserInfoFiveLose+2], u.Five.Lose)
	binary.LittleEndian.PutUint16(ret[PosOfUserInfoFiveTie:PosOfUserInfoFiveTie+2], u.Five.Tie)

	binary.LittleEndian.PutUint16(ret[PosOfUserInfoChcWin:PosOfUserInfoChcWin+2], u.ChineseChess.Win)
	binary.LittleEndian.PutUint16(ret[PosOfUserInfoChcLose:PosOfUserInfoChcLose+2], u.ChineseChess.Lose)
	binary.LittleEndian.PutUint16(ret[PosOfUserInfoChcTie:PosOfUserInfoChcTie+2], u.ChineseChess.Tie)

	binary.LittleEndian.PutUint16(ret[PosOfUserInfoChessEloRating:PosOfUserInfoChessEloRating+2], u.ChessEloRating)

	binary.LittleEndian.PutUint16(ret[PosOfUserInfoGoWin:PosOfUserInfoGoWin+2], u.GoChess.Win)
	binary.LittleEndian.PutUint16(ret[PosOfUserInfoGoLose:PosOfUserInfoGoLose+2], u.GoChess.Lose)
	binary.LittleEndian.PutUint16(ret[PosOfUserInfoGoTie:PosOfUserInfoGoTie+2], u.GoChess.Tie)

	binary.LittleEndian.PutUint32(ret[PosOfUserInfoWithMe:PosOfUserInfoWithMe+4], u.WithMe)
	binary.LittleEndian.PutUint32(ret[PosOfUserInfoBrcID:PosOfUserInfoBrcID+4], u.brcID)

	binary.LittleEndian.PutUint32(ret[PosOfUserInfoWBTime:PosOfUserInfoWBTime+4], uint32(u.wbTime.Unix()))

	return ret

}
