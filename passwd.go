package bbs

import (
	"encoding/binary"
	"io"
	"log"
	"os"
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
	PosOfPttPasswdOver18       = PosOfPttPasswdJustify + PosOfPttPasswdJustify + PTT_REGLEN + 1 + 3
	PosOfPttPasswdPagerUiType  = PosOfPttPasswdOver18 + 1
	PosOfPttPasswdPager        = PosOfPttPasswdPagerUiType + 1
	PosOfPttPasswdInvisible    = PosOfPttPasswdPager + 1
	PosOfPttPasswdExMailBox    = PosOfPttPasswdInvisible + 2

	PosOfPttPasswdCareer        = PosOfPttPasswdExMailBox + 4
	PosOfPttPasswdRole          = PosOfPttPasswdCareer + PTT_PHONESZ + 1 + 44
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

	PosOfPttPasswdSignature = PosOfPttPasswdUaVersion + 2
	PosOfPttPasswdBadPost   = PosOfPttPasswdSignature + 2
	PosOfPttPasswdDarkTie   = PosOfPttPasswdBadPost + 2
	PosOfPttPasswdMyAngel   = PosOfPttPasswdDarkTie + 2

	PosOfPttPasswdChessEloRating    = PosOfPttPasswdMyAngel + PTT_IDLEN + 1 + 1
	PosOfPttPasswdWithMe            = PosOfPttPasswdChessEloRating + 2
	PosOfPttPasswdTimeRemoveBadPost = PosOfPttPasswdWithMe + 4
	PosOfPttPasswdTimeViolateLaw    = PosOfPttPasswdTimeRemoveBadPost + 4
)

// type UserecRaw struct {
// 	Version uint32

// 	//Require const-bytes to have correct Unsafe.Sizeof
// 	UserID     [PTT_IDLEN + 1]byte  /* 使用者ID (alpha-number only) */
// 	RealName   [PTT_REALNAMESZ]byte /* 真實姓名 */
// 	Nickname   [PTT_NICKNAMESZ]byte /* 暱稱 */
// 	PasswdHash [PTT_PASSLEN]byte    /* 密碼 (hashed bytes) */
// 	Pad1       byte

// 	UFlag        uint32                /* 習慣, see uflags.h */
// 	Unused1      uint32                /* 從前放習慣2, 使用前請先清0 */
// 	UserLevel    uint32                /* 權限 */
// 	NumLoginDays uint32                /* 上線資歷 (每日最多+1的登入次數) */
// 	NumPosts     uint32                /* 文章篇數 */
// 	FirstLogin   time4                 /* 註冊時間 */
// 	LastLogin    time4                 /* 最近上站時間(包含隱身) */
// 	LastHost     [PTT_IPV4LEN + 1]byte /* 上次上站來源 */
// 	Money        int32                 /* Ptt幣 */
// 	Unused2      [4]byte

// 	Email       [PTT_EMAILSZ]byte    /* Email */
// 	Address     [PTT_ADDRESSSZ]byte  /* 住址 */
// 	Justify     [PTT_REGLEN + 1]byte /* 審核資料 */
// 	UnusedBirth [3]uint8             /* 生日 月日年 */
// 	Over18      uint8                /* 是否已滿18歲 */
// 	PagerUIType uint8                /* 呼叫器界面類別 (was: WATER_*) */
// 	Pager       uint8                /* 呼叫器狀態 */
// 	Invisible   uint8                /* 隱形狀態 */
// 	Unused4     [2]byte
// 	Exmailbox   uint32 /* 購買信箱數 */

// 	// r3968 移出 sizeof(chicken_t)=128 bytes
// 	Unused5       [4]byte
// 	Career        [PTT_CAREERSZ]byte /* 學歷職業 */
// 	UnusedPhone   [PTT_PHONESZ]byte  /* 電話 */
// 	Unused6       uint32              從前放轉換前的 numlogins, 使用前請先清0
// 	Chkpad1       [44]byte
// 	Role          uint32 /* Role-specific permissions */
// 	LastSeen      time4  /* 最近上站時間(隱身不計) */
// 	TimeSetAngel  time4  /* 上次得到天使時間 */
// 	TimePlayAngel time4  /* 上次與天使互動時間 (by day) */
// 	// 以上應為 sizeof(chicken_t) 同等大小

// 	LastSong  time4  /* 上次點歌時間 */
// 	LoginView uint32 /* 進站畫面 */
// 	Unused8   uint8  // was: channel
// 	Pad2      uint8

// 	VlCount    uint16  /* 違法記錄 ViolateLaw counter */
// 	FiveWin    uint16  /* 五子棋戰績 勝 */
// 	FiveLose   uint16  /* 五子棋戰績 敗 */
// 	FiveTie    uint16  /* 五子棋戰績 和 */
// 	ChcWin     uint16  /* 象棋戰績 勝 */
// 	ChcLose    uint16  /* 象棋戰績 敗 */
// 	ChcTie     uint16  /* 象棋戰績 和 */
// 	Conn6Win   uint16  /* 六子棋戰績 勝 */
// 	Conn6Lose  uint16  /* 六子棋戰績 敗 */
// 	Conn6Tie   uint16  /* 六子棋戰績 和 */
// 	UnusedMind [2]byte /* 舊心情 */
// 	GoWin      uint16  /* 圍棋戰績 勝 */
// 	GoLose     uint16  /* 圍棋戰績 敗 */
// 	GoTie      uint16  /* 圍棋戰績 和 */
// 	DarkWin    uint16  /* 暗棋戰績 勝 */
// 	DarkLose   uint16  /* 暗棋戰績 敗 */
// 	UaVersion  uint8   /* Applicable user agreement version */

// 	Signature uint8               /* 慣用簽名檔 */
// 	Unused10  uint8               /* 從前放好文章數, 使用前請先清0 */
// 	BadPost   uint8               /* 評價為壞文章數 */
// 	DarkTie   uint16              /* 暗棋戰績 和 */
// 	MyAngel   [PTT_IDLEN + 1]byte /* 我的小天使 */
// 	Pad3      byte

// 	ChessEloRating    uint16 /* 象棋等級分 */
// 	WithMe            uint32 /* 我想找人下棋，聊天.... */
// 	TimeRemoveBadPost time4  /* 上次刪除劣文時間 */
// 	TimeViolateLaw    time4  /* 被開罰單時間 */

// 	PadTail [28]byte
// }

// const USERECRAWSZ = int64(unsafe.Sizeof(UserecRaw{}))

// func NewUserecRaw() *UserecRaw {
// 	return &UserecRaw{}
// }

// https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h
type Userec struct {
	Version  uint32 // Magic Number
	UserId   string // 使用者帳號，或稱使用者 ID
	RealName string // 真實姓名
	Nickname string // 暱稱
	Password string // 密碼，預設為 crypt, 不同版本實作可能不同

	UserFlag     uint32
	UserLevel    uint32 // 權限
	NumLoginDays uint32
	NumPosts     uint32
	FirstLogin   uint32
	LastLogin    uint32
	LastHost     string
	Money        int32

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

	ViolateLaw uint16
	FiveWin    uint16
	FiveLose   uint16
	FiveTie    uint16
	ChcWin     uint16
	ChcLose    uint16
	ChcTie     uint16
	Conn6Win   uint16
	Conn6Lose  uint16
	Conn6Tie   uint16
	GoWin      uint16
	GoLose     uint16
	GoTie      uint16
	DarkWin    uint16
	DarkLose   uint16
	UaVersion  uint8

	Signature         uint8
	BadPost           uint8
	DarkTie           uint16
	MyAngel           string
	ChessEloRating    uint16
	WithMe            uint32
	TimeRemoveBadPost time4
	TimeViolateLaw    time4
	// TODO
}

func OpenUserecFile(filename string) ([]*Userec, error) {
	log.Println(filename)
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
	user.UserId = CstrToString(data[PosOfPttPasswdUserId : PosOfPttPasswdUserId+PTT_IDLEN+1])
	user.RealName = Big5ToUtf8(CstrToBytes(data[PosOfPttPasswdRealName : PosOfPttPasswdRealName+20]))
	user.Nickname = Big5ToUtf8(CstrToBytes(data[PosOfPttPasswdNickname : PosOfPttPasswdNickname+24]))
	user.Password = CstrToString(data[PosOfPttPasswdPassword : PosOfPttPasswdPassword+PTT_PASSLEN])

	user.UserFlag = binary.LittleEndian.Uint32(data[PosOfPttPasswdUserFlag : PosOfPttPasswdUserFlag+4])
	user.UserLevel = binary.LittleEndian.Uint32(data[PosOfPttPasswdUserLevel : PosOfPttPasswdUserLevel+4])
	user.NumLoginDays = binary.LittleEndian.Uint32(data[PosOfPttPasswdNumLoginDays : PosOfPttPasswdNumLoginDays+4])
	user.NumPosts = binary.LittleEndian.Uint32(data[PosOfPttPasswdNumPosts : PosOfPttPasswdNumPosts+4])
	user.FirstLogin = binary.LittleEndian.Uint32(data[PosOfPttPasswdFirstLogin : PosOfPttPasswdFirstLogin+4])
	user.LastLogin = binary.LittleEndian.Uint32(data[PosOfPttPasswdLastLogin : PosOfPttPasswdLastLogin+4])
	user.LastHost = CstrToString(data[PosOfPttPasswdLastHost : PosOfPttPasswdLastHost+PTT_IPV4LEN+1])

	return user, nil
}
