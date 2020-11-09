package ptttype

const (
	////
	// pttstruct.h
	IDLEN   = 12 /* Length of bid/uid */
	IPV4LEN = 15 /* a.b.c.d form */

	PASS_INPUT_LEN = 8 /* Length of valid input password length.
	   For DES, set to 8. */
	PASSLEN = 14 /* Length of encrypted passwd field */
	REGLEN  = 38 /* Length of registration data */

	REALNAMESZ = 20 /* Size of real-name field */
	NICKNAMESZ = 24 /* SIze of nick-name field */
	EMAILSZ    = 50 /* Size of email field */
	ADDRESSSZ  = 50 /* Size of address field */
	CAREERSZ   = 40 /* Size of career field */
	PHONESZ    = 20 /* Size of phone field */

	PASSWD_VERSION = 4194

	BTLEN = 48 /* Length of board title */

	MAX_BMs = 4

	USHM_SIZE = MAX_ACTIVE * 41 / 40
)

const (
	// type struct requires const.
	// use "// +build custom" to setup customized config
	MAX_USERS = 60 /* 最高註冊人數 */

	MAX_ACTIVE = 31 /* 最多同時上站人數 */

	MAX_BOARD = 100 /* 最大開板個數 */

	HASH_BITS = 16 /* userid->uid hashing bits */

	MAX_FRIEND = 100 /* 載入 cache 之最多朋友數目 */

	MAX_REJECT = 32 /* 載入 cache 之最多壞人數目 */

	MAX_MSGS = 10 /* 水球(熱訊)忍耐上限 */

	MAX_ADBANNER = 100 /* 最多動態看板數 */

	HOTBOARDCACHE = 0 /* 熱門看板快取 */

	MAX_FROM = 10 /* 最多故鄉數 */

	MAX_REVIEW = 7 /* 最多水球回顧 */

	NUMVIEWFILE = 14 /* 進站畫面最多數 */

	MAX_ADBANNER_SECTION = 10 /* 最多動態看板類別 */

	MAX_ADBANNER_HEIGHT = 11 /* 最大動態看板內容高度 */
)
