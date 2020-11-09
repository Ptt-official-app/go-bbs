package ptttype

var (
	//////////
	// make.conf
	//////////
	BBSHOME = "/home/bbs" /* BBS home-dir. MUST-NOT set this directly. use SetBBSHOME() */

	//////////
	// pttbbs.conf
	//////////	/* 使用 HUGETLB shared memory . 目前只在 Linux 上有效 */
	USE_HUGETLB = false

	/* 讓過於熱門或被鬧的版冷靜, SHM 會變大一些些 */
	USE_COOLDOWN = true

	/* 是否可以開新的 shm */
	IS_CREATE_SHM = false

	/* 在某些平台之下, shared-memory規定需要為一定的 aligned size,
	   如在 linux x86_64 下使用 HUGETLB 時需為 4MB aligned,
	   而在 linux ia64 下使用 HUGETLB時需為 256MB aligned.
	   單位為 bytes */
	SHMALIGNEDSIZE = 1048576 * 1 // 4MB for x86_64, 1MB for development

	SHM_KEY = 1228

	/////////////////////////////////////////////////////////////////////////////
	// Site settings 站台功能設定

	MAX_POST_MONEY = 100 /* 發表文章稿費的上限 */

	MAX_CHICKEN_MONEY = 100 /* 養雞場獲利上限 */

	MAX_GUEST_LIFE = 3 * 24 * 60 * 60 /* 最長未認證使用者保留時間(秒) */

	MAX_EDIT_LINE = 2048 /* 文章最長編輯長度 */

	MAX_EDIT_LINE_LARGE = 32000 // 大檔最長編輯長度

	MAX_LIFE = 120 * 24 * 60 * 60 /* 最長使用者保留時間(秒) */

	KEEP_DAYS_REGGED = 120 /* 已註冊使用者保留多久 */

	KEEP_DAYS_UNREGGED = 15 /* 未註冊使用者保留多久 */

	THREAD_SEARCH_RANGE = 500 /* 系列文章搜尋上限 */

	/* 定義是否使用外籍使用者註冊
	   及外國人最長居留時間，之後需向站方申請永久居留權 */
	FOREIGN_REG = false

	FOREIGN_REG_DAY = 30 /* 外籍使用者試用日期上限 */

	FORCE_PROCESS_REGISTER_FORM = 0

	/* 定義看板好友名單將會在幾秒鐘後失效強迫重載 */
	HBFLexpire = 432000

	MAX_EXKEEPMAIL = 1000 /* 最多信箱加大多少封 */

	/* 在轉信時附上的時區. 若在台灣, 中國大陸等地, 用預設的即可.          */
	INNTIMEZONE = "+0800 (CST)"

	ADD_EXMAILBOX = 0 /* 贈送信箱 */

	BADPOST_CLEAR_DURATION = 180 // 消劣文時間限制

	BADPOST_MIN_CLEAR_DURATION = 3 // 劣文首消時間限制

	MAX_CROSSNUM = 9 /* 最多crosspost次數 */

	/* (deprecated) more.c 中文章頁數上限(lines/22), +4 for safe */
	MAX_PAGES = MAX_EDIT_LINE/22 + 4

	MAX_QUERYLINES = 16 /* 顯示 Query/Plan 訊息最大行數 */

	MAX_LOGIN_INFO = 128 /* 最多上線通知人數 */

	MAX_POST_INFO = 32 /* 最多新文章通知人數 */

	MAX_NAMELIST = 128 /* 最多其他特別名單人數 */

	MAX_NOTE = 20 /* 最多保留幾篇留言？ */

	MAX_SIGLINES = 6 /* 簽名檔引入最大行數 */

	LOGINATTEMPTS = 3 /* 最大進站失誤次數 */

	MAX_KEEPMAIL = 200 /* 一般 user 最多保留幾封 MAIL？ */

	MAX_KEEPMAIL_SOFTLIMIT = 2500 /* 除 admin 外，無法寄給此人 */

	MAX_KEEPMAIL_HARDLIMIT = 20000 /* 信箱數量的上限，超過就不給寄信 */

	BADCIDCHARS = " *" /* Chat Room 中禁用於 nick 的字元 */

	MAX_ROOM = 16 /* 聊天室最多有幾間包廂？ */

	MAXTAGS = 255 /* t(tag) 的最大數量 */

	WRAPMARGIN = 511 /* 編輯器 wrap 長度 */

	// XXX we may need to take care of FROMHOST
	/*
	   #ifdef USE_MASKED_FROMHOST
	   #define FROMHOST    fromhost_masked
	   #else
	   #define FROMHOST    fromhost
	   #endif
	*/

	/////////////////////////////////////////////////////////////////////////////
	// Logging 記錄設定

	LOG_CONF_KEYWORD = false // 記錄搜尋的關鍵字

	LOG_CONF_INTERNETMAIL = false // 記錄 internet outgoing mail

	LOG_CONF_PUSH = false // 記錄推文

	LOG_CONF_EDIT_CALENDAR = false // 記錄編輯行事曆

	LOG_CONF_POST = false // 記錄發文

	LOG_CONF_CRAWLER = false // 記錄 crawlers

	LOG_CONF_CROSSPOST = false // 記錄轉錄

	LOG_CONF_BAD_REG_CODE = false // 記錄打錯的註冊碼

	LOG_CONF_VALIDATE_REG = false // 記錄審核註冊單

	LOG_CONF_MASS_DELETE = false // 記錄大量刪除檔案

	LOG_CONF_OSONG_VERBOSE = false // 詳細點播記錄

	LOG_CONF_EDIT_TITLE = false // 編輯標題記錄

	/////////////////////////////////////////////////////////////////////////////
	// Default Configurations 預設參數

	// 若想停用下列參數請在 pttbbs.conf 定義 NO_XXX (ex, NO_LOGINASNEW)
	LOGINASNEW = true /* 採用上站申請帳號制度 */

	DOTIMEOUT = true /* 處理閒置時間 */

	INTERNET_EMAIL = true /* 支援 InterNet Email 功能(含 Forward) */

	SHOWUID = true /* 站長可看見使用者 UID */

	SHOWBOARD = true /* 站長可看見使用者看板 */

	SHOWPID = true /* 站長可看見使用者 PID */

	HAVE_ANONYMOUS = true /* 提供 Anonymous 板 */

	HAVE_ORIGIN = true /* 顯示 author 來自何處 */

	USE_BSMTP = true /* 使用opus的BSMTP 寄收信? */

	REJECT_FLOOD_POST = true /* 防止BlahBlah式灌水 */

	// #define  HAVE_INFO               /* 顯示程式版本說明 */
	// #define  HAVE_LICENSE            /* 顯示 GNU 版權畫面 */
	// #define  HAVE_REPORT             /* (轉信)系統追蹤報告 */

	IDLE_TIMEOUT   = 43200 /* 一般情況之 timeout (12hr) */
	SHOW_IDLE_TIME = true  /* 顯示閒置時間 */

	//////////
	// common.h
	//////////
	SZ_RECENTLOGIN = 16000 // size of max recent log before rotation
	SZ_RECENTPAY   = 16000

	// XXX FN_SAFEDEL
	// https://github.com/ptt/pttbbs/blob/master/include/common.h#L68
	// pttbbs commit: 6bdd36898bde207683a441cdffe2981e95de5b20
	USE_EDIT_HISTORY      = false
	FN_SAFEDEL            = ".deleted"
	FN_SAFEDEL_PREFIX_LEN = 8 // must match FN_SAFEDEL

	STR_SAFEDEL_TITLE = "(本文已被刪除)"

	SAFE_ARTICLE_DELETE_NUSER = 2
)
