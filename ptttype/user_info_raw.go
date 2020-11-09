package ptttype

import (
	"unsafe"

	"github.com/PichuChen/go-bbs/types"
)

type UserInfoRaw struct {
	//Require updating SHM_VERSION if USER_INFO_RAW_SZ is changed.
	Uid      int32       /* Used to find user name in passwd file */
	Pid      types.Pid_t /* kill() to notify user of talk request */
	SockAddr int32       /* ... */

	/* user data */
	UserLevel  uint32
	UserID     [IDLEN + 1]byte
	Nickname   [24]byte
	From       [27]byte       /* machine name the user called in from */
	FromIp     types.InAddr_t // was: int     from_alias;
	DarkWin    uint16
	DarkLose   uint16
	Gap0       byte
	AngelPause uint8 // TODO move to somewhere else in future.
	DarkTie    uint16

	/* friends */
	FriendTotal int32 /* 好友比較的cache 大小 */
	NFriends    int16 /* 下面 friend[] 只用到前幾個,
	   用來 bsearch */
	Unused3_     int16
	MyFriend     [MAX_FRIEND]int32
	Gap1         [4]byte
	FriendOnline [MAX_FRIEND]uint32 /* point到線上好友 utmpshm的位置 */
	/* 好友比較的cache 前兩個bit是狀態 */
	Gap2   [4]byte
	Reject [MAX_REJECT]int32
	Gap3   [4]byte

	/* messages */
	MsgCount byte
	Unused4_ [3]byte
	Msgs     [MAX_MSGS]MsgQueueRaw
	Gap4     [MSG_QUEUE_RAW_SZ]byte /* avoid msgs racing and overflow */

	/* user status */
	Birth        int8  /* 是否是生日 Ptt*/
	Active       uint8 /* When allocated this field is true */
	Invisible    uint8 /* Used by cloaking function in Xyz menu */
	Mode         uint8 /* UL/DL, Talk Mode, Chat Mode, ... */
	Pager        byte  /* pager toggle, YEA, or NA */
	Unused5_     byte
	Conn6Win     uint16
	LastAct      types.Time4 /* 上次使用者動的時間 */
	Alerts       byte        /* mail alert, passwd update... */
	UnusedMind_  byte
	Conn6Lose    uint16
	UnusedMind2_ byte

	/* chatroom/talk/games calling */
	Sig        byte /* signal type */
	Conn6Tie   uint16
	DestUid    int32 /* talk uses this to identify who called */
	DestUip    int32 /* dest index in utmpshm->uinfo[] */
	SockActive uint8 /* Used to coordinate talk requests */

	/* chat */
	InChat uint8    /* for in_chat commands   */
	Chatid [11]byte /* chat id, if in chat mode */

	/* games */
	LockMode uint8           /* 不准 multi_login 玩的東西 */
	Turn     byte            /* 遊戲的先後 */
	Mateid   [IDLEN + 1]byte /* 遊戲對手的 id */
	Color    byte            /* 暗棋 顏色 */

	/* game record */
	FiveWin        uint16
	FiveLose       uint16
	FiveTie        uint16
	ChcWin         uint16
	ChcLose        uint16
	ChcTie         uint16
	ChessEloRating uint16
	GoWin          uint16
	GoLose         uint16
	GoTie          uint16

	/* misc */
	WithMe uint32
	BrcID  uint32

	WBTime types.Time4
}

//Require updating SHM_VERSION if USER_INFO_RAW_SZ is changed.
const USER_INFO_RAW_SZ = unsafe.Sizeof(UserInfoRaw{})
