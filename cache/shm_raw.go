package cache

import (
	"unsafe"

	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

type SHMRaw struct {
	Version int32 // SHM_VERSION   for verification
	Size    int32 // sizeof(SHM_t) for verification

	/* uhash */
	/* uhash is a userid->uid hash table -- jochang */
	Userid       [ptttype.MAX_USERS][ptttype.IDLEN + 1]byte
	Gap1         [ptttype.IDLEN + 1]byte
	NextInHash   [ptttype.MAX_USERS]int32
	Gap2         [types.INT32_SZ]byte
	Money        [ptttype.MAX_USERS]int32
	Gap3         [types.INT32_SZ]byte
	CooldownTime [ptttype.MAX_USERS]types.Time4
	Gap4         [types.INT32_SZ]byte
	HashHead     [1 << ptttype.HASH_BITS]int32
	Gap5         [types.INT32_SZ]byte
	Number       int32 /* # of users total */
	Loaded       int32 /* .PASSWD has been loaded? */

	/* utmpshm */
	UInfo  [ptttype.USHM_SIZE]ptttype.UserInfoRaw
	Gap6   [ptttype.USER_INFO_RAW_SZ]byte
	Sorted [2][9][ptttype.USHM_SIZE]int32
	/* 第一維double buffer 由currsorted指向目前使用的
	   第二維sort type */
	Gap7          [types.INT32_SZ]byte
	CurrSorted    int32
	UTMPUptime    types.Time4
	UTMPNumber    int
	UTMPNeedSort  byte
	UTMPBusyState byte

	/* brdshm */
	Gap8          [types.INT32_SZ]byte
	BMCache       [ptttype.MAX_BOARD][ptttype.MAX_BMs]int32
	Gap9          [types.INT32_SZ]byte
	BCache        [ptttype.MAX_BOARD]ptttype.BoardHeaderRaw
	Gap10         [types.INT32_SZ]byte
	BSorted       [2][ptttype.MAX_BOARD]int32 /* 0: by name 1: by class */ /* 裡頭存的是 bid-1 */
	Gap11         [types.INT32_SZ]byte
	NHOTs         uint8
	HBcache       [ptttype.HOTBOARDCACHE]int32
	Gap12         [types.INT32_SZ]byte
	BusyStateB    [ptttype.MAX_BOARD]types.Time4
	Gap13         [types.INT32_SZ]byte
	Total         [ptttype.MAX_BOARD]int32
	Gap14         [types.INT32_SZ]byte
	NBottom       [ptttype.MAX_BOARD]uint8 /* number of bottom */
	Gap15         [types.INT32_SZ]byte
	Hbfl          [ptttype.MAX_BOARD][ptttype.MAX_FRIEND + 1]int /* hidden board friend list, 0: load time, 1-MAX_FRIEND: uid */
	Gap16         [types.INT32_SZ]byte
	LastPostTime  [ptttype.MAX_BOARD]types.Time4
	Gap17         [types.INT32_SZ]byte
	BUptime       types.Time4
	BTouchTime    types.Time4
	BNumber       int32
	BBusyState    int32
	CloseVoteTime types.Time4

	/* pttcache */
	Notes   [ptttype.MAX_ADBANNER][256 * ptttype.MAX_ADBANNER_HEIGHT]byte
	Gap18   [types.INT32_SZ]byte
	TodayIs [20]byte
	// FIXME remove it
	NeverUsedNNotes_ [ptttype.MAX_ADBANNER_SECTION]int32 /* 一節中有幾個 看板 */
	Gap19            [types.INT32_SZ]byte
	// FIXME remove it
	NeverUsedNextRefresh_ [ptttype.MAX_ADBANNER_SECTION]int32 /* 下一次要refresh的 看板 */
	Gap20                 [types.INT32_SZ]byte
	LoginMsg              ptttype.MsgQueueRaw /* 進站水球 */
	LastFilm              int32
	LastUsong             int32
	PUptime               types.Time4
	PTouchTime            types.Time4
	PBusyState            int

	/* SHM 中的全域變數, 可用 shmctl 設定或顯示. 供動態調整或測試使用 */
	GV2 shmGV2

	/* statistic */
	Statistic [ptttype.STAT_MAX]uint32

	// 從前作為故鄉使用 (fromcache). 現已被 daemon/fromd 取代。
	DeprecatedHomeIp_   [ptttype.MAX_FROM]uint32
	DeprecatedHomeMask_ [ptttype.MAX_FROM]uint32
	DeprecatedHomeDesc_ [ptttype.MAX_FROM][32]byte
	DeprecatedHomeNum_  int32

	MaxUser    int32
	MaxTime    types.Time4
	FUptime    types.Time4
	FTouchTime types.Time4
	FBusyState int32
}

type shmGV2 struct {
	DyMaxMctive  int32       /* 動態設定最大人數上限     */
	TooManyUsers int32       /* 超過人數上限不給進的個數 */
	NoonLineUser int32       /* 站上使用者不高亮度顯示   */
	Now          types.Time4 // __attribute__ ((deprecated));
	NWelcomes    int32
	Shutdown     int32 /* shutdown flag */
}

const SHM_RAW_SZ = unsafe.Sizeof(SHMRaw{})
