package ptttype

import (
	"encoding/binary"
	"os"
	"unsafe"

	"github.com/PichuChen/go-bbs/types"
)

type UserecRaw struct {
	Version uint32

	//Require const-bytes to have correct Unsafe.Sizeof
	UserID     [IDLEN + 1]byte  /* 使用者ID (alpha-number only) */
	RealName   [REALNAMESZ]byte /* 真實姓名 */
	Nickname   [NICKNAMESZ]byte /* 暱稱 */
	PasswdHash [PASSLEN]byte    /* 密碼 (hashed bytes) */
	Pad1       byte

	UFlag        uint32            /* 習慣, see uflags.h */
	Unused1      uint32            /* 從前放習慣2, 使用前請先清0 */
	UserLevel    uint32            /* 權限 */
	NumLoginDays uint32            /* 上線資歷 (每日最多+1的登入次數) */
	NumPosts     uint32            /* 文章篇數 */
	FirstLogin   types.Time4       /* 註冊時間 */
	LastLogin    types.Time4       /* 最近上站時間(包含隱身) */
	LastHost     [IPV4LEN + 1]byte /* 上次上站來源 */
	Money        int32             /* Ptt幣 */
	Unused2      [4]byte

	Email       [EMAILSZ]byte    /* Email */
	Address     [ADDRESSSZ]byte  /* 住址 */
	Justify     [REGLEN + 1]byte /* 審核資料 */
	UnusedBirth [3]uint8         /* 生日 月日年 */
	Over18      uint8            /* 是否已滿18歲 */
	PagerUIType uint8            /* 呼叫器界面類別 (was: WATER_*) */
	Pager       uint8            /* 呼叫器狀態 */
	Invisible   uint8            /* 隱形狀態 */
	Unused4     [2]byte
	Exmailbox   uint32 /* 購買信箱數 */

	// r3968 移出 sizeof(chicken_t)=128 bytes
	Unused5       [4]byte
	Career        [CAREERSZ]byte /* 學歷職業 */
	UnusedPhone   [PHONESZ]byte  /* 電話 */
	Unused6       uint32         /* 從前放轉換前的 numlogins, 使用前請先清0 */
	Chkpad1       [44]byte
	Role          uint32      /* Role-specific permissions */
	LastSeen      types.Time4 /* 最近上站時間(隱身不計) */
	TimeSetAngel  types.Time4 /* 上次得到天使時間 */
	TimePlayAngel types.Time4 /* 上次與天使互動時間 (by day) */
	// 以上應為 sizeof(chicken_t) 同等大小

	LastSong  types.Time4 /* 上次點歌時間 */
	LoginView uint32      /* 進站畫面 */
	Unused8   uint8       // was: channel
	Pad2      uint8

	VlCount    uint16  /* 違法記錄 ViolateLaw counter */
	FiveWin    uint16  /* 五子棋戰績 勝 */
	FiveLose   uint16  /* 五子棋戰績 敗 */
	FiveTie    uint16  /* 五子棋戰績 和 */
	ChcWin     uint16  /* 象棋戰績 勝 */
	ChcLose    uint16  /* 象棋戰績 敗 */
	ChcTie     uint16  /* 象棋戰績 和 */
	Conn6Win   uint16  /* 六子棋戰績 勝 */
	Conn6Lose  uint16  /* 六子棋戰績 敗 */
	Conn6Tie   uint16  /* 六子棋戰績 和 */
	UnusedMind [2]byte /* 舊心情 */
	GoWin      uint16  /* 圍棋戰績 勝 */
	GoLose     uint16  /* 圍棋戰績 敗 */
	GoTie      uint16  /* 圍棋戰績 和 */
	DarkWin    uint16  /* 暗棋戰績 勝 */
	DarkLose   uint16  /* 暗棋戰績 敗 */
	UaVersion  uint8   /* Applicable user agreement version */

	Signature uint8           /* 慣用簽名檔 */
	Unused10  uint8           /* 從前放好文章數, 使用前請先清0 */
	BadPost   uint8           /* 評價為壞文章數 */
	DarkTie   uint16          /* 暗棋戰績 和 */
	MyAngel   [IDLEN + 1]byte /* 我的小天使 */
	Pad3      byte

	ChessEloRating    uint16      /* 象棋等級分 */
	WithMe            uint32      /* 我想找人下棋，聊天.... */
	TimeRemoveBadPost types.Time4 /* 上次刪除劣文時間 */
	TimeViolateLaw    types.Time4 /* 被開罰單時間 */

	PadTail [28]byte
}

const USEREC_RAW_SZ = int64(unsafe.Sizeof(UserecRaw{}))

func NewUserecRaw() *UserecRaw {
	return &UserecRaw{}
}

func NewUserecRawWithFile(file *os.File) (*UserecRaw, error) {
	userecRaw := &UserecRaw{}

	err := binary.Read(file, binary.LittleEndian, userecRaw)
	if err != nil {
		return nil, err
	}

	return userecRaw, nil
}
