package ptttype

var (
	//////////
	// make.conf
	//////////

	BBSHOME = "/home/bbs" /* BBS home-dir. MUST-NOT set this directly. use SetBBSHOME() */

	//////////
	// pttbbs.conf
	//////////
	/* 定義 BBS 站名位址 */
	BBSNAME    = "新批踢踢"           /* 中文站名 */
	BBSENAME   = "PTT2"           /* 英文站名 */
	MYHOSTNAME = "ptt2.cc"        /* 網路位址 */
	MYIP       = "140.112.30.143" /* IP位址 */

	/* 定義是否查詢文章的 web 版 URL，及 URL 用的 hostname/prefix */
	QUERY_ARTICLE_URL = true                    /* 是否提供查詢文章 URL */
	URL_PREFIX        = "http://www.ptt.cc/bbs" /* URL prefix */
	/*
	   http://www.ptt.cc/bbs/SYSOP/M.1197864962.A.476.html
	   ^^^^^^^^^^^^^^^^^^^^^
	   這個部分
	*/

	/* *** 以下為定義時會多出功能的板名 *** */

	/* 若定義, 提供美工特別用板 */
	BN_ARTDSN = "Artdsn"

	/* 若定義，該板發文不受行限或是可上傳 */
	BN_BBSMOVIE = "BBSmovie"

	// /* 若定義，則.... */
	BN_WHOAMI = "WhoAmI"

	/* 若定義, 則全站所有五子棋/象棋棋譜都會紀錄在此板 */
	BN_FIVECHESS_LOG = BBSMNAME + "Five"
	BN_CCHESS_LOG    = BBSMNAME + "CChess"

	/* 若定義，則動態看板會動態檢查爭議性字眼 */
	BN_NOTE_AGGCHKDIR = "<點歌> 動態看板"

	/* 若定義則啟用修文自動合併系統 */
	EDITPOST_SMARTMERGE = false

	/* 可以設定多重進站畫面 */
	MULTI_WELCOME_LOGIN = false

	/* 若定義, 則所有編輯文章最下方都會加入編輯來源.
	   否則只有 SYSOP板會加入來源                    */
	ALL_REEDIT_LOG = false

	/* 板主可以按大寫 H切換隱形與否 */
	BMCHS = false

	/* 水球整理, 看板備份等等外部程式 */
	OUTJOBSPOOL = true

	/* 若定義, 則不能舉辦賭盤 */
	NO_GAMBLE = true

	/* 可動態透過 GLOBALVAR[9]調整使用者上限 */
	DYMAX_ACTIVE = true

	/* 程式每天最多可以跑多久 (in seconds) 因為有的時候會出現跑不停的 process */
	CPULIMIT_PER_DAY = 30

	/* 若定義, 若程式失敗, 會等待 86400 秒以讓 gdb來 attach */
	DEBUGSLEEP = true

	/* 若定義, 在轉寄位址輸入錯誤時會有讓使用者回報訊息的提示 */
	/* 這個選項存在的原因是因為有部份使用者信誓旦旦說他們沒打錯但看不出程式錯誤 */
	DEBUG_FWDADDRERR = false

	/* 若定義, 則可在外部 (shmctl cmsignal) 要求將 mbbsd將 zapbuf 釋放掉.
	   會使用非正規的記憶體要求函式. (目前只在 FreeBSD上測試過)
	   !!請注意!!
	   除非您確切知道這個能能在做什麼並且有須要,
	   否則請不要打開這個功能!!                                           */
	CRITICAL_MEMORY = false

	/* 對於 port 23的, 會預先 fork 幾隻出來. 如此在系統負荷高的時候,
	   仍可有好的上站率 */
	PRE_FORK = 10

	/* 若定義, 則開啟 Big5 轉 UTF-8 的功能 */
	CONVERT = true

	/* 若定義, 則在文章列表的時候不同日期會標上不同顏色 */
	COLORDATE = false

	/* 若定義, 在使用者註冊之前, 會先顯示出該檔案, 經使用者確認後才能註冊 */
	HAVE_USERAGREEMENT            = "etc/UserAgreement"
	HAVE_USERAGREEMENT_VERSION    = "etc/UserAgreementVersion"
	HAVE_USERAGREEMENT_ACCEPTABLE = "etc/UserAgreementAcceptable"

	/* DBCS 相關設定 */
	/* DBCS Aware: 讓游標不會跑到 DBCS trailing bytes 上 */
	DBCSAWARE = true

	/* 若定義，guest 帳號預設不顯示一字雙色 */
	GUEST_DEFAULT_DBCS_NOINTRESC = false

	/* 使用新式的 pmore (piaip's more) 代替舊式 bug 抓不完的 more 或是簡易的 minimore */
	USE_PMORE = false

	/* 使用 rfork()取代 fork() . 目前只在 FreeBSD上有效 */
	USE_RFORK = false

	/* 使用 HUGETLB shared memory . 目前只在 Linux 上有效 */
	USE_HUGETLB = false

	/* 在某些平台之下, shared-memory規定需要為一定的 aligned size,
	   如在 linux x86_64 下使用 HUGETLB 時需為 4MB aligned,
	   而在 linux ia64 下使用 HUGETLB時需為 256MB aligned.
	   單位為 bytes */
	SHMALIGNEDSIZE = 1048576 * 4 // 4MB for x86_64

	/* 讓過於熱門或被鬧的版冷靜, SHM 會變大一些些 */
	USE_COOLDOWN = true

	/* 若定義, 則在刪除看板文章的時候, 僅會在 .DIR 中標明, 並不會將該資料
	   從 .DIR 中拿掉. 可以避免多項問題 (尤其是熱門看板一堆推薦及編輯時)
	   須配合使用 (尚未完成)                                              */
	SAFE_ARTICLE_DELETE = false

	/* 若定義, 則在傳送水球的時候, 不會直接 kill 該程序. 理論上可以減少大
	   量的系統負和                                                       */
	NOKILLWATERBALL = false

	/* 若定義, 則 SYSOP帳號並不會自動加上站長權限.
	   在第一次啟動時, 您並不能定義 (否則就拿不到站長權了) .
	   而在設定完成後, 若您站長帳號並不叫做 SYSOP,
	   則可透過 NO_SYSOP_ACCOUNT 關閉該帳號, 以避免安全問題發生.          */
	NO_SYSOP_ACCOUNT = false

	/* 開啟小天使小主人功能 */
	PLAY_ANGEL = false

	/* 若定義, 則使用舊式推文 */
	OLDRECOMMEND = false

	/* 若定義, 則 guest 可推文，格式變為 IP+日期 */
	GUESTRECOMMEND = false

	/* 定義幾秒內算快速推文 */
	FASTRECMD_LIMIT = 90

	/* 若定義, 可設定轉錄自動在原文留下記錄 */
	USE_AUTOCPLOG = true

	/* 若定義, 新板設定自動開記錄，不過 USE_AUTOCPLOG 還是要開才有用 */
	DEFAULT_AUTOCPLOG = true

	/* 如果 time_t 是 8 bytes的話 (如 X86_64) */
	TIMET64 = true

	/* 使用 utmpd, 在外部運算好友資料, 如果您確定這個在做什麼才開啟 */
	UTMPD      = false
	UTMPD_ADDR = "192.168.0.1:5120"
	/* 在 cacheserver 上面擋掉狂上下站的使用者 */
	NOFLOODING = false

	/* 使用 daemon/fromd, 使用外部daemon紀錄上站故鄉名稱 */
	FROMD = false

	/* 若定義, 則不允許註冊 guest */
	NO_GUEST_ACCOUNT_REG = false

	/* 限制一個email能註冊的帳號數量 (要使用請在make的時候加 WITH_EMAILDB) */
	EMAILDB_LIMIT = 5

	USE_REG_CAPTCHA            = false
	USE_POST_CAPTCHA_FOR_NOREG = false
	USE_VOTE_CAPTCHA           = false
	USE_REMOTE_CAPTCHA         = false
	CAPTCHA_INSERT_SERVER_ADDR = "127.0.0.1:80"
	CAPTCHA_INSERT_HOST        = CAPTCHA_INSERT_SERVER_ADDR
	CAPTCHA_INSERT_URI         = "/captcha/insert"
	CAPTCHA_INSERT_SECRET      = ""
	CAPTCHA_URL_PREFIX         = "http://localhost/captcha"
	CAPTCHA_CODE_LENGTH        = 32

	REQUIRE_SECURE_CONN_TO_REGISTER  = false
	REQUIRE_VERIFY_EMAIL_AT_REGISTER = false

	/* 前進站畫面 */
	//INSCREEN = "前進站畫面"
	INSCREEN = ""

	//////////
	// config.h
	//////////
	BBSPROGPOSTFIX = "/bin/mbbsd"             /* 主程式 */
	BBSPROG        = BBSHOME + BBSPROGPOSTFIX /* 主程式 */
	BAN_FILE       = "BAN"                    /* 關站通告檔 */
	LOAD_FILE      = "/proc/loadavg"          /* for Linux */

	/////////////////////////////////////////////////////////////////////////////
	// System Name Configuration 系統名稱設定

	/* 系統名(郵件用)，建議別超過 3 個字元。 詳見 sample/pttbbs.conf */
	BBSMNAME = "Ptt"

	/* 系統名(選單用)，建議剛好 4 個字元。 詳見 sample/pttbbs.conf */
	BBSMNAME2 = "Ｐtt"

	/* 錢幣名，建議剛好 3 個字元。 詳見 sample/pttbbs.conf */
	MONEYNAME = BBSMNAME + "幣"

	/* AID 顯示的站台名稱。 若 IP 太長請另行定義。 */
	AID_HOSTNAME = MYHOSTNAME

	/////////////////////////////////////////////////////////////////////////////
	// Themes 主題配色

	TITLE_COLOR = ANSIColor("0;1;37;46") /* 主畫面上方標題列 */

	HLP_CATEGORY_COLOR = ANSIColor("0;1;32") /* 說明表格內分類項 */

	HLP_DESCRIPTION_COLOR = ANSIColor("0") /* 說明表格內說明項 */

	HLP_KEYLIST_COLOR = ANSIColor("0;1;36") /* 說明表格內按鍵項 */

	/////////////////////////////////////////////////////////////////////////////
	// OS Settings 作業系統相關設定

	BBSUSER = "bbs"

	BBSUID = 9999

	BBSGID = 99

	TAR_PATH = "tar"

	MUTT_PATH = "mutt"

	MAXPATHLEN = 256

	PATHLEN = 256

	DEFAULT_FOLDER_CREATE_PERM = 0755

	DEFAULT_FILE_CREATE_PERM = (0644)

	SHM_KEY = 1228

	PASSWDSEM_KEY = 2010 /* semaphore key */

	// SYSLOG_FACILITY = LOG_LOCAL0 (not sure the corresponding LOG_LOCAL0 in golang)

	/* 若定義, 用一個奇怪的數字來檢查我的最愛和看板列表是否錯誤 */
	MEM_CHECK = 0x98761234

	RELAY_SERVER_IP = "127.0.0.1"

	XCHATD_ADDR = ":3838"

	/////////////////////////////////////////////////////////////////////////////
	// Default Board Names 預設看板名稱

	/* *** 以下為預設板名 (見 include/config.h) *** */

	/* 安全紀錄 */
	BN_SECURITY = "Security"
	/* 動態看板的家 */
	BN_NOTE = "Note"
	/* 紀錄 */
	BN_RECORD = "Record"

	/* SYSOP 板 */
	BN_SYSOP = "SYSOP"
	/* 測試板 */
	BN_TEST = "Test"
	/* 發生錯誤時建議的回報板名為此板 */
	BN_BUGREPORT = BBSMNAME + "Bug"
	/* 法律訴訟的板 */
	BN_LAW = BBSMNAME + "Law"
	/* 新手板(會自動進我的最愛) */
	BN_NEWBIE = BBSMNAME + "NewHand"
	/* 找看板(會自動進我的最愛) */
	BN_ASKBOARD = "AskBoard"
	/* 外國板 */
	BN_FOREIGN = BBSMNAME + "Foreign"

	BN_ID_PROBLEM = "SYSOP"

	BN_DELETED = "deleted"

	BN_JUNK = "junk"

	BN_POLICELOG = "PoliceLog"

	BN_UNANONYMOUS = "UnAnonymous"

	BN_NEWIDPOST = "NEWIDPOST"

	BN_ALLPOST = "ALLPOST"

	BN_ALLHIDPOST = "ALLHIDPOST"

	/////////////////////////////////////////////////////////////////////////////
	// Performance Parameters 效能參數

	MAX_USERS = 150000 /* 最高註冊人數 */

	MAX_ACTIVE = 1024 /* 最多同時上站人數 */

	MAX_GUEST = 100 /* 最多 guest 上站人數 */

	MAX_CPULOAD = 70 /* CPU 最高load */

	DEBUGSLEEP_SECONDS = 3600 /* debug 等待時間 */

	MAX_BOARD = 8192 /* 最大開板個數 */

	HASH_BITS = 16 /* userid->uid hashing bits */

	OVERLOADBLOCKFDS = 0 /* 超載後會保留這麼多個 fd */

	MAX_FRIEND = 256 /* 載入 cache 之最多朋友數目 */

	MAX_REJECT = 32 /* 載入 cache 之最多壞人數目 */

	MAX_MSGS = 10 /* 水球(熱訊)忍耐上限 */

	MAX_ADBANNER = 500 /* 最多動態看板數 */

	MAX_SWAPUSED = 0.7 /* SWAP最高使用率 */

	HOTBOARDCACHE = 0 /* 熱門看板快取 */

	TARQUEUE_TIME_STR = "深夜" // 看板備份時間訊息 (應與 contab 一致)

	/////////////////////////////////////////////////////////////////////////////
	// More system messages 系統訊息

	RECYCLE_BIN_NAME = "資源回收筒" // "垃圾桶"

	RECYCLE_BIN_OWNER = "[" + RECYCLE_BIN_NAME + "]"

	TIME_CAPSULE_NAME = "Magical Index" // "Time Capsule"

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

	MAX_FROM = 300 /* 最多故鄉數 */

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

	MAX_ADBANNER_SECTION = 10 /* 最多動態看板類別 */

	MAX_ADBANNER_HEIGHT = 11 /* 最大動態看板內容高度 */

	MAX_QUERYLINES = 16 /* 顯示 Query/Plan 訊息最大行數 */

	MAX_LOGIN_INFO = 128 /* 最多上線通知人數 */

	MAX_POST_INFO = 32 /* 最多新文章通知人數 */

	MAX_NAMELIST = 128 /* 最多其他特別名單人數 */

	MAX_NOTE = 20 /* 最多保留幾篇留言？ */

	MAX_SIGLINES = 6 /* 簽名檔引入最大行數 */

	MAX_REVIEW = 7 /* 最多水球回顧 */

	NUMVIEWFILE = 14 /* 進站畫面最多數 */

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
