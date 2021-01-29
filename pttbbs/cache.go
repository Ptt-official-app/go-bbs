package pttbbs

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
	"syscall"
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

	// 	int   version;  // SHM_VERSION   for verification // 0, 4
	//     int   size;	    // sizeof(SHM_t) for verification // 4, 4

	//     /* uhash */
	//     /* uhash is a userid->uid hash table -- jochang */
	//     char    userid[MAX_USERS][IDLEN + 1]; // 8, 650
	//     char    gap_1[IDLEN+1]; // 658, 13
	//     int     next_in_hash[MAX_USERS]; // 672, 200
	//     char    gap_2[sizeof(int)]; //
	//     int     money[MAX_USERS]; //    ,200
	//     char    gap_3[sizeof(int)];
	//     // TODO(piaip) Always have this var - no more #ifdefs in structure.
	// #ifdef USE_COOLDOWN
	//     time4_t cooldowntime[MAX_USERS];
	// #endif
	//     char    gap_4[sizeof(int)];
	//     int     hash_head[1 << HASH_BITS];
	//     char    gap_5[sizeof(int)];
	//     int     number;				/* # of users total */
	//     int     loaded;				/* .PASSWD has been loaded? */

	//     /* utmpshm */
	//     userinfo_t      uinfo[USHM_SIZE]; // 有 NOKILLWATERBALL: 3484,
	//     char    gap_6[sizeof(userinfo_t)];
	//     int             sorted[2][9][USHM_SIZE]; // 朋友列表
	//                     /* 第一維double buffer 由currsorted指向目前使用的
	// 		       第二維sort type */
	//     char    gap_7[sizeof(int)];
	//     int     currsorted;
	//     time4_t UTMPuptime;
	//     int     UTMPnumber;
	//     char    UTMPneedsort;
	//     char    UTMPbusystate;

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

	//     /* brdshm */
	//     char    gap_8[sizeof(int)];
	//     int     BMcache[MAX_BOARD][MAX_BMs];
	//     char    gap_9[sizeof(int)];
	//     boardheader_t   bcache[MAX_BOARD];
	//     char    gap_10[sizeof(int)];
	//     int     bsorted[2][MAX_BOARD]; /* 0: by name 1: by class */ /* 裡頭存的是 bid-1 */
	//     char    gap_11[sizeof(int)];
	//     // TODO(piaip) Always have this var - no more #ifdefs in structure.
	// #if HOTBOARDCACHE
	//     unsigned char    nHOTs;
	//     int              HBcache[HOTBOARDCACHE];
	// #endif
	//     char    gap_12[sizeof(int)];
	//     time4_t busystate_b[MAX_BOARD];
	//     char    gap_13[sizeof(int)];
	//     int     total[MAX_BOARD];
	//     char    gap_14[sizeof(int)];
	//     unsigned char  n_bottom[MAX_BOARD]; /* number of bottom */
	//     char    gap_15[sizeof(int)];
	//     int     hbfl[MAX_BOARD][MAX_FRIEND + 1]; /* hidden board friend list, 0: load time, 1-MAX_FRIEND: uid */
	//     char    gap_16[sizeof(int)];
	//     time4_t lastposttime[MAX_BOARD];
	//     char    gap_17[sizeof(int)];
	//     time4_t Buptime;
	//     time4_t Btouchtime;
	//     int     Bnumber;
	//     int     Bbusystate;
	//     time4_t close_vote_time;

	//     /* pttcache */
	//     char    notes[MAX_ADBANNER][256*MAX_ADBANNER_HEIGHT];
	//     char    gap_18[sizeof(int)];
	//     char    today_is[20];
	// FIXME remove it
	// int     __never_used__n_notes[MAX_ADBANNER_SECTION];      /* 一節中有幾個 看板 */
	// char    gap_19[sizeof(int)];
	// // FIXME remove it
	// int     __never_used__next_refresh[MAX_ADBANNER_SECTION]; /* 下一次要refresh的 看板 */
	// char    gap_20[sizeof(int)];

	posOfLoginMsg   int
	posOfLastFilm   int
	posOflastUSong  int
	posOfPUptime    int
	posOfPTouchTime int
	posOfPBusyState int

	//    msgque_t loginmsg;  /* 進站水球 */
	//    int     last_film;
	//    int     last_usong;
	//    time4_t Puptime;
	//    time4_t Ptouchtime;
	//    int     Pbusystate;

	//    /* SHM 中的全域變數, 可用 shmctl 設定或顯示. 供動態調整或測試使用 */
	//    union {
	// int     v[512];
	// struct {
	//     int     dymaxactive;  /* 動態設定最大人數上限     */
	//     int     toomanyusers; /* 超過人數上限不給進的個數 */
	//     int     noonlineuser; /* 站上使用者不高亮度顯示   */
	//     time4_t now __attribute__ ((deprecated));
	//     int     nWelcomes;
	//     int     shutdown;     /* shutdown flag */

	//     /* 注意, 應保持 align sizeof(int) */
	// } e;
	//    } GV2;
	/* statistic */
	// unsigned int    statistic[STAT_MAX];

	posOfStatistic int

	// 從前作為故鄉使用 (fromcache). 現已被 daemon/fromd 取代。
	// unsigned int    _deprecated_home_ip[MAX_FROM];
	// unsigned int    _deprecated_home_mask[MAX_FROM];
	// char            _deprecated_home_desc[MAX_FROM][32];
	// int        	    _deprecated_home_num;

	posOfMaxUser   int
	posOfMaxTime   int
	posOfFUptime   int
	posOfTouchTime int
	posOfBusyState int

	// int     max_user
	// time4_t max_time
	// time4_t Fuptime
	// time4_t Ftouchtime
	// int     Fbusystate
}

type MemoryMappingSetting struct {
	MaxUsers int
	IDLen    int
}

type Cache struct {
	buf []byte
	*MemoryMappingSetting
	cachePos
}

func (c *Cache) caculatePos() {
	c.posOfVersion = 0
	c.posOfSize = c.posOfVersion + 4
	c.posOfUserId = c.posOfSize + 4

	c.posOfNextInHash = c.posOfUserId + c.MaxUsers*(c.IDLen+1) + (c.IDLen + 1)
	// Align
	if c.posOfNextInHash%2 != 0 {
		c.posOfNextInHash += 1
	}
	fmt.Println("c.posOfNextInHash", c.posOfNextInHash)
	c.posOfMoney = c.posOfNextInHash + c.MaxUsers*4 + 4
	// 	c.posOfMoney      int
	// c.
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

func NewCache(connectionString string, settings *MemoryMappingSetting) (*Cache, error) {
	if strings.HasPrefix(connectionString, "file://") {
		// mmap
		filePath := strings.Replace(connectionString, "file://", "", -1)
		f, err := os.Open(filePath)
		// f, err := os.Open("../../../dump.shm")
		if err != nil {
			fmt.Println("open shm fail", err)
			return nil, fmt.Errorf("open error:", err)
		}
		fd := int(f.Fd())
		fmt.Println("fd:", fd)

		stat, err := f.Stat()
		if err != nil {
			fmt.Println("stat error", err)
			return nil, fmt.Errorf("stat error:", err)
		}

		size := int(stat.Size())
		fmt.Println("size", size)
		b, err := syscall.Mmap(fd, 0, size, syscall.PROT_READ, syscall.MAP_SHARED)
		if err != nil {
			fmt.Println("mmap error", err)
			return nil, fmt.Errorf("mmap error:", err)
		}

		c := Cache{
			buf:                  b,
			MemoryMappingSetting: settings,
		}
		c.caculatePos()
		return &c, nil
	}
	return nil, fmt.Errorf("unsupport connectionString")

}

// func main() {
// 	c, err := NewCache("file://../../dump.shm", &MemoryMappingSetting{
// 		MaxUsers: 50,
// 		IDLen:    12,
// 	})
// 	if err != nil {
// 		fmt.Println("new cache err:", err)
// 		return
// 	}
// 	fmt.Println("version", c.Version())
// 	fmt.Println("userid, 0:", c.UserId(0))
// 	fmt.Println("money, 0:", c.Money(0))
// }

func (c *Cache) Version() uint32 {
	// Should be 4842
	return binary.LittleEndian.Uint32(c.buf[c.posOfVersion : c.posOfVersion+4])
}

func (c *Cache) UserId(uid int) string {
	// TODO: Check for out of range
	s := c.posOfUserId + (c.IDLen+1)*uid
	return string(bytes.Split(c.buf[s:s+c.IDLen+1], []byte("\x00"))[0])
}

func (c *Cache) Money(uid int) int32 {
	// TODO: Check for out of range
	s := c.posOfMoney + 4*uid
	return int32(binary.LittleEndian.Uint32(c.buf[s : s+4]))
}
