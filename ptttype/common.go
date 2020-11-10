package ptttype

//////////
// common.h
//////////
const (
	STR_GUEST  = "guest" // guest account
	STR_REGNEW = "new"   // 用來建新帳號的名稱

	STR_RECOVER = "/recover" // recover function

	// XXX DEFAULT_BOARD
	//#define DEFAULT_BOARD   str_sysop

	// BBS Configuration Files
	FN_CONF_EDITABLE    = "etc/editable"       // 站務可編輯的系統檔案列表
	FN_CONF_RESERVED_ID = "etc/reserved.id"    // 保留系統用無法註冊的 ID
	FN_CONF_BINDPORTS   = "etc/bindports.conf" // 預設要提供連線服務的 port 列表

	// BBS Data File Names
	FN_CHICKEN                 = "chicken"
	FN_USSONG                  = "ussong"    /* 點歌統計 */
	FN_POST_NOTE               = "post.note" /* po文章備忘錄 */
	FN_POST_BID                = "post.bid"
	FN_MONEY                   = "etc/money"
	FN_OVERRIDES               = "overrides"
	FN_REJECT                  = "reject"
	FN_WATER                   = "water"           // 舊水桶
	FN_BANNED                  = "banned"          // 新水桶
	FN_BANNED_HISTORY          = "banned.history"  // 新水桶之歷史記錄
	FN_BADPOST_HISTORY         = "badpost.history" // 劣文歷史記錄
	FN_CANVOTE                 = "can_vote"
	FN_VISABLE                 = "visable" // 不知道是誰拼錯的，將錯就錯吧...
	FN_ALOHAED                 = "alohaed" // 上站要通知我的名單 (編輯用)
	FN_ALOHA                   = "aloha"   // 我上站要通知的名單 (自動產生)
	FN_USIES                   = "usies"   /* BBS log */
	FN_DIR                     = ".DIR"
	FN_DIR_BOTTOM              = ".DIR.bottom"
	FN_BOARD                   = ".BRD"    /* board list */
	FN_USEBOARD                = "usboard" /* 看板統計 */
	FN_TOPSONG                 = "etc/topsong"
	FN_TICKET                  = "ticket"
	FN_TICKET_END              = "ticket.end"
	FN_TICKET_LOCK             = "ticket.end.lock"
	FN_TICKET_ITEMS            = "ticket.items"
	FN_TICKET_RECORD           = "ticket.data"
	FN_TICKET_USER             = "ticket.user"
	FN_TICKET_OUTCOME          = "ticket.outcome"
	FN_TICKET_BRDLIST          = "boardlist"
	FN_BRDLISTHELP             = "etc/boardlist.help"
	FN_BOARDHELP               = "etc/board.help"
	FN_MAIL_ACCOUNT_SYSOP      = "etc/mail_account_sysop"      // 帳號站長信箱列表
	FN_MAIL_ACCOUNT_SYSOP_DESC = "etc/mail_account_sysop_desc" // 帳號站長信箱說明
	FN_USERMEMO                = "memo.txt"                    // 使用者個人記事本
	FN_BADLOGIN                = "logins.bad"                  // in BBSHOME & user directory
	FN_RECENTLOGIN             = "logins.recent"               // in user directory
	FN_FORWARD                 = ".forward"                    /* auto forward */

	FN_RECENTPAY = "money.recent"

	// 自訂刪除文章時出現的標題與檔案
	FN_EDITHISTORY = ".history"

	MSG_CLOAKED = "已進入隱形模式(不列於使用者名單上)"
	MSG_UNCLOAK = "已離開隱形模式(公開於使用者名單上)"

	MSG_WORKING = "處理中，請稍候..."

	MSG_CANCEL   = "取消。"
	MSG_USR_LEFT = "使用者已經離開了"
	MSG_NOBODY   = "目前無人上線"

	MSG_DEL_OK     = "刪除完畢"
	MSG_DEL_CANCEL = "取消刪除"
	MSG_DEL_ERROR  = "刪除錯誤"
	MSG_DEL_NY     = "請確定刪除(Y/N)?[N] "

	MSG_FWD_OK   = "文章轉寄完成!"
	MSG_FWD_ERR1 = "轉寄錯誤: 系統錯誤 system error"
	MSG_FWD_ERR2 = "轉寄錯誤: 位址錯誤 address error"

	MSG_SURE_NY = "請您確定(Y/N)？[N] "
	MSG_SURE_YN = "請您確定(Y/N)？[Y] "

	MSG_BID    = "請輸入看板名稱: "
	MSG_UID    = "請輸入使用者代號: "
	MSG_PASSWD = "請輸入您的密碼: "

	MSG_BIG_BOY     = "我是大帥哥! ^o^Y"
	MSG_BIG_GIRL    = "世紀大美女 *^-^*"
	MSG_LITTLE_BOY  = "我是底迪啦... =)"
	MSG_LITTLE_GIRL = "最可愛的美眉! :>"
	MSG_MAN         = "麥當勞叔叔 (^O^)"
	MSG_WOMAN       = "叫我小阿姨!! /:>"
	MSG_PLANT       = "植物也有性別喔.."
	MSG_MIME        = "礦物總沒性別了吧"

	ERR_BOARD_OPEN   = ".BOARD 開啟錯誤"
	ERR_BOARD_UPDATE = ".BOARD 更新有誤"
	ERR_PASSWD_OPEN  = ".PASSWDS 開啟錯誤"

	ERR_BID      = "你搞錯了啦！沒有這個板喔！"
	ERR_UID      = "這裡沒有這個人啦！"
	ERR_PASSWD   = "密碼不對喔！請檢查帳號及密碼大小寫有無輸入錯誤。"
	ERR_FILENAME = "檔名不正確！"

	TN_ANNOUNCE = "[公告]"

	STR_AUTHOR1 = "作者:"
	STR_AUTHOR2 = "發信人:"
	STR_POST1   = "看板:"
	STR_POST2   = "站內:"

	STR_LOGINDAYS     = "登入次數"
	STR_LOGINDAYS_QTY = "次"

	/* AIDS */
	AID_DISPLAYNAME = "文章代碼(AID)"
	/* end of AIDS */

	/* QUERY_ARTICLE_URL */
	URL_DISPLAYNAME = "文章網址"
	/* end of QUERY_ARTICLE_URL */

	/* LONG MESSAGES */

	MSG_SEPARATOR = "───────────────────────────────────────"

	/* Flags to getdata input function */
	NOECHO   = 0
	DOECHO   = 1
	LCECHO   = 2
	NUMECHO  = 4
	GCARRY   = 8 // (from M3) do not empty input buffer.
	PASSECHO = 0x10

	YEA = true /* Booleans  (Yep, for true and false) */
	NA  = false

	EQUSTR = 0 /* for strcmp */

	/* 好友關係 */
	IRH       = 1  /* I reject him.		*/
	HRM       = 2  /* He reject me.		*/
	IBH       = 4  /* I am board friend of him.	*/
	IFH       = 8  /* I friend him (He is one of my friends). */
	HFM       = 16 /* He friends me (I am one of his friends). */
	ST_FRIEND = IBH | IFH | HFM
	ST_REJECT = IRH | HRM

	//XXX not sure what it is. #define QCAST           int (*)(const void *, const void *)
	//XXX replaced by ToUpper #define chartoupper(c)  ((c >= 'a' && c <= 'z') ? c+'A'-'a' : c)

	LEN_AUTHOR1 = 5
	LEN_AUTHOR2 = 7

	/* the title of article will be truncate to PROPER_TITLE_LEN */
	PROPER_TITLE_LEN = 42

	/* ----------------------------------------------------- */
	/* 水球模式 邊界定義                                     */
	/* ----------------------------------------------------- */
	WB_OFO_USER_TOP    = 7
	WB_OFO_USER_BOTTOM = 11
	WB_OFO_USER_NUM    = WB_OFO_USER_BOTTOM - WB_OFO_USER_TOP + 1
	WB_OFO_USER_LEFT   = 28
	WB_OFO_MSG_TOP     = 15
	WB_OFO_MSG_BOTTOM  = 22
	WB_OFO_MSG_LEFT    = 4

	/* ----------------------------------------------------- */
	/* 標題類形                                              */
	/* ----------------------------------------------------- */
	SUBJECT_NORMAL  = 0
	SUBJECT_REPLY   = 1
	SUBJECT_FORWARD = 2
	SUBJECT_LOCKED  = 3

	/* ----------------------------------------------------- */
	/* 群組名單模式   Ptt                                    */
	/* ----------------------------------------------------- */
	FRIEND_OVERRIDE = 0
	FRIEND_REJECT   = 1
	FRIEND_ALOHA    = 2
	// #define FRIEND_POST     3	    // deprecated
	FRIEND_SPECIAL = 4
	FRIEND_CANVOTE = 5
	BOARD_WATER    = 6
	BOARD_VISABLE  = 7

	LOCK_THIS  = 1 // lock這線不能重複玩
	LOCK_MULTI = 2 // lock所有線不能重複玩

	MAX_MODES      = 127
	MAX_RECOMMENDS = 100

	STR_CURSOR  = ">"
	STR_CURSOR2 = "●"
	STR_UNCUR   = " "
	STR_UNCUR2  = "  "

	NOTREPLYING    = -1
	REPLYING       = 0
	RECVINREPLYING = 1

	/* ----------------------------------------------------- */
	/* 編輯器選項                                            */
	/* ----------------------------------------------------- */
	EDITFLAG_TEXTONLY   = 0x00000001
	EDITFLAG_UPLOAD     = 0x00000002
	EDITFLAG_ALLOWLARGE = 0x00000004
	EDITFLAG_ALLOWTITLE = 0x00000008
	// #define EDITFLAG_ANONYMOUS  (0x00000010)
	EDITFLAG_KIND_NEWPOST   = 0x00000010
	EDITFLAG_KIND_REPLYPOST = 0x00000020
	EDITFLAG_KIND_SENDMAIL  = 0x00000040
	EDITFLAG_KIND_MAILLIST  = 0x00000080
	EDITFLAG_WARN_NOSELFDEL = 0x00000100
	// #define EDITFLAG_ALLOW_LOCAL    (0x00000200)
	EDIT_ABORTED = -1

	/* ----------------------------------------------------- */
	/* 聊天室常數 (xchatd)                                   */
	/* ----------------------------------------------------- */
	EXIT_LOGOUT   = 0
	EXIT_LOSTCONN = -1
	EXIT_CLIERROR = -2
	EXIT_TIMEDOUT = -3
	EXIT_KICK     = -4

	CHAT_LOGIN_OK      = "OK"
	CHAT_LOGIN_EXISTS  = "EX"
	CHAT_LOGIN_INVALID = "IN"
	CHAT_LOGIN_BOGUS   = "BG"

	/* ----------------------------------------------------- */
	/* Grayout Levels                                        */
	/* ----------------------------------------------------- */
	GRAYOUT_COLORBOLD = -2
	GRAYOUT_BOLD      = -1
	GRAYOUT_DARK      = 0
	GRAYOUT_NORM      = 1
	GRAYOUT_COLORNORM = 2

	/* Typeahead */
	TYPEAHEAD_NONE  = -1
	TYPEAHEAD_STDIN = 0

	/* ----------------------------------------------------- */
	/* Constants                                             */
	/* ----------------------------------------------------- */
	DAY_SECONDS   = 86400
	MONTH_SECONDS = DAY_SECONDS * 30
	MILLISECONDS  = 1000 // milliseconds of one second
)

var (
	FN_CONF_BANIP_POSTFIX = "/etc/banip.conf"               // 禁止連線的 IP 列表
	FN_CONF_BANIP         = BBSHOME + FN_CONF_BANIP_POSTFIX // 禁止連線的 IP 列表
	FN_PASSWD_POSTFIX     = "/.PASSWDS"                     /* User records */
	FN_PASSWD             = BBSHOME + FN_PASSWD_POSTFIX     /* User records */

	SHM_HUGETLB = 04000 /* segment is mapped via hugetlb */

	MSG_SELECT_BOARD = ANSIColor("7") + "【 選擇看板 】" + ANSIReset() + "\n" + "請輸入看板名稱(按空白鍵自動搜尋): "
)
