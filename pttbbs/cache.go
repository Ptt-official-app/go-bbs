package pttbbs

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/PichuChen/go-bbs/cache"
)

type cachePos struct {
	posOfVersion    int
	posOfSize       int
	posOfUserId     int
	posOfNextInHash int
	posOfMoney      int

	posOfCooldownTime int

	posOfHashHeader int
	posOfUserNumber int
	posOfUserLoaded int

	// utmpshm
	posOfUInfo  int
	posOfSorted int

	posOfCurrentSorted int
	posOfUTMPUptime    int
	posOfUTMPNumber    int
	posOfUTMPNeedSort  int
	posOfUTMPBusyState int

	posOfBMCache     int
	posOfBoardCache  int
	posOfSortedBoard int

	posOfHotBoardCache int

	posOfBoardBusyState        int
	posOfBoardTotal            int
	posOfBoardBottomNumber     int
	posOfHiddenBoardFriendList int
	posOfBoardLastPostTime     int
	posOfBUptime               int
	posOfBNumber               int
	posOfBBusystat             int
	posOfCloseVoteTime         int

	posOfNotes   int
	posOfTodayIs int

	posOfLoginMsg   int
	posOfLastFilm   int
	posOflastUSong  int
	posOfPUptime    int
	posOfPTouchTime int
	posOfPBusyState int

	posOfStatistic int

	posOfMaxUser   int
	posOfMaxTime   int
	posOfFUptime   int
	posOfTouchTime int
	posOfBusyState int
}

// MemoryMappingSetting provides parameters for calculating the memory position of
//  the relevant fields.
type MemoryMappingSetting struct {
	AlignmentBytes int // 1, 2, 4 or 8, 1 for no aligment

	MaxUsers int
	IDLen    int
}

// Cache provides an IPC(inter-process communication) bridge with process-based
// pttbbs process, and shars the cache with board info, user info ...
type Cache struct {
	cache.Cache
	*MemoryMappingSetting
	cachePos
}

// cacluatePos find out posOf values on runtime with MemoryMappingSetting
// notice that, different compiler option will result in different
// align padding, it may cause bugs.
// the align usually will be 2 with pttbbs by examine result.
//
// See: https://en.wikipedia.org/wiki/Data_structure_alignment
func (c *Cache) caculatePos() {

	c.posOfVersion = 0
	c.posOfSize = c.posOfVersion + 4
	c.posOfUserId = c.posOfSize + 4

	c.posOfNextInHash = c.posOfUserId + c.MaxUsers*(c.IDLen+1) + (c.IDLen + 1)
	// Align
	if c.posOfNextInHash%c.AlignmentBytes != 0 {
		padding := c.AlignmentBytes - c.posOfNextInHash%c.AlignmentBytes
		c.posOfNextInHash += padding
	}
	fmt.Println("c.posOfNextInHash", c.posOfNextInHash)
	c.posOfMoney = c.posOfNextInHash + c.MaxUsers*4 + 4

	// TODO: other pos value
	// 	c.posOfCooldownTime int
	// c.
	// 	c.posOfHashHeader int
	// 	c.posOfUserNumber int
	// 	c.posOfUserLoaded int
	// c.
	// 	c.// utmpshm
	// 	c.posOfUInfo  int
	// 	c.posOfSorted int
	// c.
	// 	c.posOfCurrentSorted int
	// 	c.posOfUTMPUptime    int
	// 	c.posOfUTMPNumber    int
	// 	c.posOfUTMPNeedSort  int
	// 	c.posOfUTMPBusyState int
	// c.
	// 	c.// 	int   version;  // SHM_VERSION   for verification // 0, 4
	// 	c.//     int   size;	    // sizeof(SHM_t) for verification // 4, 4
	// c.
	// 	c.//     /* uhash */
	// 	c.//     /* uhash is a userid->uid hash table -- jochang */
	// 	c.//     char    userid[MAX_USERS][IDLEN + 1]; // 8, 650
	// 	c.//     char    gap_1[IDLEN+1]; // 658, 13
	// 	c.//     int     next_in_hash[MAX_USERS]; // 672, 200
	// 	c.//     char    gap_2[sizeof(int)]; //
	// 	c.//     int     money[MAX_USERS]; //    ,200
	// 	c.//     char    gap_3[sizeof(int)];
	// 	c.//     // TODO(piaip) Always have this var - no more #ifdefs in structure.
	// 	c.// #ifdef USE_COOLDOWN
	// 	c.//     time4_t cooldowntime[MAX_USERS];
	// 	c.// #endif
	// 	c.//     char    gap_4[sizeof(int)];
	// 	c.//     int     hash_head[1 << HASH_BITS];
	// 	c.//     char    gap_5[sizeof(int)];
	// 	c.//     int     number;				/* # of users total */
	// 	c.//     int     loaded;				/* .PASSWD has been loaded? */
	// c.
	// 	c.//     /* utmpshm */
	// 	c.//     userinfo_t      uinfo[USHM_SIZE]; // 有 NOKILLWATERBALL: 3484,
	// 	c.//     char    gap_6[sizeof(userinfo_t)];
	// 	c.//     int             sorted[2][9][USHM_SIZE]; // 朋友列表
	// 	c.//                     /* 第一維double buffer 由currsorted指向目前使用的
	// 	c.// 		       第二維sort type */
	// 	c.//     char    gap_7[sizeof(int)];
	// 	c.//     int     currsorted;
	// 	c.//     time4_t UTMPuptime;
	// 	c.//     int     UTMPnumber;
	// 	c.//     char    UTMPneedsort;
	// 	c.//     char    UTMPbusystate;
	// c.
	// 	c.posOfBMCache     int
	// 	c.posOfBoardCache  int
	// 	c.posOfSortedBoard int
	// c.
	// 	c.posOfHotBoardCache int
	// c.
	// 	c.posOfBoardBusyState        int
	// 	c.posOfBoardTotal            int
	// 	c.posOfBoardBottomNumber     int
	// 	c.posOfHiddenBoardFriendList int
	// 	c.posOfBoardLastPostTime     int
	// 	c.posOfBUptime               int
	// 	c.posOfBNumber               int
	// 	c.posOfBBusystat             int
	// 	c.posOfCloseVoteTime         int
	// c.
	// 	c.posOfNotes   int
	// 	c.posOfTodayIs int
	// c.
	// 	c.//     /* brdshm */
	// 	c.//     char    gap_8[sizeof(int)];
	// 	c.//     int     BMcache[MAX_BOARD][MAX_BMs];
	// 	c.//     char    gap_9[sizeof(int)];
	// 	c.//     boardheader_t   bcache[MAX_BOARD];
	// 	c.//     char    gap_10[sizeof(int)];
	// 	c.//     int     bsorted[2][MAX_BOARD]; /* 0: by name 1: by class */ /* 裡頭存的是 bid-1 */
	// 	c.//     char    gap_11[sizeof(int)];
	// 	c.//     // TODO(piaip) Always have this var - no more #ifdefs in structure.
	// 	c.// #if HOTBOARDCACHE
	// 	c.//     unsigned char    nHOTs;
	// 	c.//     int              HBcache[HOTBOARDCACHE];
	// 	c.// #endif
	// 	c.//     char    gap_12[sizeof(int)];
	// 	c.//     time4_t busystate_b[MAX_BOARD];
	// 	c.//     char    gap_13[sizeof(int)];
	// 	c.//     int     total[MAX_BOARD];
	// 	c.//     char    gap_14[sizeof(int)];
	// 	c.//     unsigned char  n_bottom[MAX_BOARD]; /* number of bottom */
	// 	c.//     char    gap_15[sizeof(int)];
	// 	c.//     int     hbfl[MAX_BOARD][MAX_FRIEND + 1]; /* hidden board friend list, 0: load time, 1-MAX_FRIEND: uid */
	// 	c.//     char    gap_16[sizeof(int)];
	// 	c.//     time4_t lastposttime[MAX_BOARD];
	// 	c.//     char    gap_17[sizeof(int)];
	// 	c.//     time4_t Buptime;
	// 	c.//     time4_t Btouchtime;
	// 	c.//     int     Bnumber;
	// 	c.//     int     Bbusystate;
	// 	c.//     time4_t close_vote_time;
	// c.
	// 	c.//     /* pttcache */
	// 	c.//     char    notes[MAX_ADBANNER][256*MAX_ADBANNER_HEIGHT];
	// 	c.//     char    gap_18[sizeof(int)];
	// 	c.//     char    today_is[20];
	// 	c.// FIXME remove it
	// 	c.// int     __never_used__n_notes[MAX_ADBANNER_SECTION];      /* 一節中有幾個 看板 */
	// 	c.// char    gap_19[sizeof(int)];
	// 	c.// // FIXME remove it
	// 	c.// int     __never_used__next_refresh[MAX_ADBANNER_SECTION]; /* 下一次要refresh的 看板 */
	// 	c.// char    gap_20[sizeof(int)];
	// c.
	// 	c.posOfLoginMsg   int
	// 	c.posOfLastFilm   int
	// 	c.posOflastUSong  int
	// 	c.posOfPUptime    int
	// 	c.posOfPTouchTime int
	// 	c.posOfPBusyState int
	// c.
	// 	c.//    msgque_t loginmsg;  /* 進站水球 */
	// 	c.//    int     last_film;
	// 	c.//    int     last_usong;
	// 	c.//    time4_t Puptime;
	// 	c.//    time4_t Ptouchtime;
	// 	c.//    int     Pbusystate;
	// c.
	// 	c.//    /* SHM 中的全域變數, 可用 shmctl 設定或顯示. 供動態調整或測試使用 */
	// 	c.//    union {
	// 	c.// int     v[512];
	// 	c.// struct {
	// 	c.//     int     dymaxactive;  /* 動態設定最大人數上限     */
	// 	c.//     int     toomanyusers; /* 超過人數上限不給進的個數 */
	// 	c.//     int     noonlineuser; /* 站上使用者不高亮度顯示   */
	// 	c.//     time4_t now __attribute__ ((deprecated));
	// 	c.//     int     nWelcomes;
	// 	c.//     int     shutdown;     /* shutdown flag */
	// c.
	// 	c.//     /* 注意, 應保持 align sizeof(int) */
	// 	c.// } e;
	// 	c.//    } GV2;
	// 	c./* statistic */
	// 	c.// unsigned int    statistic[STAT_MAX];
	// c.
	// 	c.posOfStatistic int
	// c.
	// 	c.// 從前作為故鄉使用 (fromcache). 現已被 daemon/fromd 取代。
	// 	c.// unsigned int    _deprecated_home_ip[MAX_FROM];
	// 	c.// unsigned int    _deprecated_home_mask[MAX_FROM];
	// 	c.// char            _deprecated_home_desc[MAX_FROM][32];
	// 	c.// int        	    _deprecated_home_num;
	// c.
	// 	c.posOfMaxUser   int
	// 	c.posOfMaxTime   int
	// 	c.posOfFUptime   int
	// 	c.posOfTouchTime int
	// 	c.posOfBusyState int

	// 	// int     max_user
	// 	// time4_t max_time
	// 	// time4_t Fuptime
	// 	// time4_t Ftouchtime
	// 	// int     Fbusystate
}

// NewCache returns Cache (SHM) by connectionString, connectionString indicate the shm location
// with uri format  eg. shmkey:1228 or file:/tmp/ramdisk/bbs.shm
func NewCache(connectionString string, settings *MemoryMappingSetting) (*Cache, error) {

	c, err := cache.NewCache(connectionString)
	if err != nil {
		return nil, fmt.Errorf("cache open error: %v", err)
	}
	return nil, fmt.Errorf("unsupport connectionString")

	ret := Cache{
		Cache:                c,
		MemoryMappingSetting: settings,
	}
	ret.caculatePos()
	return &ret, nil
}

// Version returns cache (SHM) version of pttbbs, it will be 4842 on pttbbs version
// 4d56e77 (2009/09 ~ )
func (c *Cache) Version() uint32 {
	// Should be 4842
	return binary.LittleEndian.Uint32(c.Buf()[c.posOfVersion : c.posOfVersion+4])
}

// UserId returns userId string with specific uid, such as "SYSOP",
// uid means the index in PASSWD file, start with 0.
func (c *Cache) UserId(uid int) string {
	// TODO: Check if it is out of range
	s := c.posOfUserId + (c.IDLen+1)*uid
	return string(bytes.Split(c.Buf()[s:s+c.IDLen+1], []byte("\x00"))[0])
}

// Money returns the money user have with specific uid, uid start with 0
func (c *Cache) Money(uid int) int32 {
	// TODO: Check if it is out of range
	s := c.posOfMoney + 4*uid
	return int32(binary.LittleEndian.Uint32(c.Buf()[s : s+4]))
}
