#define INCLUDE_VAR_H
#include "bbs.h"

const char * const str_permid[] = {
    "���v�O",			/* PERM_BASIC */
    "�i�J��ѫ�",		/* PERM_CHAT */
    "��H���",			/* PERM_PAGE */
    "�o��峹",			/* PERM_POST */
    "���U�{�ǻ{��",		/* PERM_LOGINOK */
    "�H��L�W��",		/* PERM_MAILLIMIT */
    "�����N",			/* PERM_CLOAK */
    "�ݨ��Ԫ�",			/* PERM_SEECLOAK */
    "�ä[�O�d�b��",		/* PERM_XEMPT */
    "���������N",		/* PERM_DENYPOST */
    "�O�D",			/* PERM_BM */
    "�b���`��",			/* PERM_ACCOUNTS */
    "��ѫ��`��",		/* PERM_CHATCLOAK */
    "�ݪO�`��",			/* PERM_BOARD */
    "����",			/* PERM_SYSOP */
    "BBSADM",			/* PERM_POSTMARK */
    "���C�J�Ʀ�]",		/* PERM_NOTOP */
    "�H�k�q�r��",		/* PERM_VIOLATELAW */
#ifdef PLAY_ANGEL
    "�i����p�Ѩ�",		/* PERM_ANGEL */
#else
    "�p�Ѩ�(�����L��)",
#endif
    "�����\\�{�ҽX���U",	/* PERM_NOREGCODE */
    "��ı����",			/* PERM_VIEWSYSOP */
    "�[��ϥΪ̦���",		/* PERM_LOGUSER */
    "�͹ܤ��v",		        /* PERM_NOCITIZEN */
    "�s�ժ�",			/* PERM_SYSSUPERSUBOP */
    "�b���f�ֲ�",		/* PERM_ACCTREG */
    "�{����",			/* PERM_PRG */
    "���ʲ�",			/* PERM_ACTION */
    "���u��",			/* PERM_PAINT */
    "ĵ���`��",			/* PERM_POLICE_MAN */
    "�p�ժ�",			/* PERM_SYSSUBOP */
    "�h�𯸪�",			/* PERM_OLDSYSOP */
    "ĵ��"			/* PERM_POLICE */
};

const char * const str_roleid[] = {
    "(�Ѩ�)�~�ޱb��(CIA)",	/* 0x00000001 ROLE_ANGEL_CIA */
    "(�Ѩ�)���ʥαb��",         /* 0x00000002 ROLE_ANGEL_ACTIVITY */
    "",		                /* 0x00000004 */
    "",		                /* 0x00000008 */
    "",		                /* 0x00000010 */
    "",		                /* 0x00000020 */
    "",		                /* 0x00000040 */
    "(�Ѩ�)�j�Ѩ�",	        /* 0x00000080 ROLE_ANGEL_ARCHANGEL */
    "(ĵ��)�ΦWĵ��",           /* 0x00000100 ROLE_POLICE_ANONYMOUS */
    "",		                /* 0x00000200 */
    "",		                /* 0x00000400 */
    "",		                /* 0x00000800 */
    "",		                /* 0x00001000 */
    "",		                /* 0x00002000 */
    "",		                /* 0x00004000 */
    "",	                        /* 0x00008000 */
    "",		                /* 0x00010000 */
    "",		                /* 0x00020000 */
    "",		                /* 0x00040000 */
    "",		                /* 0x00080000 */
    "",		                /* 0x00100000 */
    "",		                /* 0x00200000 */
    "",		                /* 0x00400000 */
    "",		                /* 0x00800000 */
    "",		                /* 0x01000000 */
    "",		                /* 0x02000000 */
    "",		                /* 0x04000000 */
    "",		                /* 0x08000000 */
    "",		                /* 0x10000000 */
    "",		                /* 0x20000000 */
    "",		                /* 0x40000000 */
    "",		                /* 0x80000000 */
};

const char * const str_permboard[] = {
    "(�L�@��)",			/* deprecated: BRD_NOZAP */
    "���C�J�έp",		/* BRD_NOCOUNT */
    "(�L�@��)",			/* deprecated: BRD_NOTRAN */
    "�s�ժO",			/* BRD_GROUPBOARD */
    "���êO",			/* BRD_HIDE */
    "����(���ݳ]�w)",		/* BRD_POSTMASK */
    "�ΦW�O",			/* BRD_ANONYMOUS */
    "�w�]�ΦW�O",		/* BRD_DEFAULTANONYMOUS */
    "�o��L���y",		/* BRD_NOCREDIT, was: BRD_BAD */
    "�s�p�M�άݪO",		/* BRD_VOTEBOARD */
    "�wĵ�i�n�o��",		/* BRD_WARNEL */
    "�����ݪO�s��",		/* BRD_TOP */
    "���i����",                 /* BRD_NORECOMMEND */
    "�p�Ѩϥi�ΦW",		/* BRD_ANGELANONYMOUS */
    "�O�D�]�w�C�J�O��",		/* BRD_BMCOUNT */
    "�s���ݪO",                 /* BRD_SYMBOLIC */
    "���i�N",                   /* BRD_NOBOO */
    "(�L�@��)",                 /* deprecated: BRD_LOCALSAVE */
    "���ݪO�|���o��",           /* BRD_RESTRICTEDPOST */
    "Guest�i�H�o��",            /* BRD_GUESTPOST */
#ifdef USE_COOLDOWN
    "�N�R",			/* BRD_COOLDOWN */
#else
    "�N�R(�����L��)",		/* BRD_COOLDOWN */
#endif
#ifdef USE_AUTOCPLOG
    "�۰ʯd����O��",		/* BRD_CPLOG */
#else
    "����O��(�����L��)",	/* BRD_CPLOG */
#endif
    "�T��ֳt����",		/* BRD_NOFASTRECMD */
    "����O�� IP",		/* BRD_IPLOGRECMD */
    "�Q�K�T",			/* BRD_OVER18 */
    "���������",		/* BRD_ALIGNEDCMT */
    "���i�ۧR",                 /* BRD_NOSELFDELPOST */
    "�O�D�i�R�S�w��r",         /* BRD_BM_MASK_CONTENT */
    "�S�Q��",
    "�S�Q��",
    "�S�Q��",
    "�S�Q��",
};

/* modes.h */
const char * const str_pager_modes[PAGER_MODES] =
{
    "����",	// PAGER_OFF
    "���}",	// PAGER_ON
    "�ޱ�",	// PAGER_DISABLE
    "����",	// PAGER_ANTIWB
    "�n��",	// PAGER_FRIENDONLY
};

int             usernum;
int             currmode = 0;
int             currsrmode = 0;
int             currbid;
char            quote_file[80] = "\0";
char            quote_user[80] = "\0";
char            currtitle[TTLEN + 1] = "\0";
const char     *currboard = "\0";
char            currBM[IDLEN * 3 + 10];
char            margs[64] = "\0";	/* main argv list */
pid_t           currpid;	/* current process ID */
time4_t         login_start_time, last_login_time;
time4_t         start_time;
userec_t        pwcuser;	/* current user structure */
unsigned int    currbrdattr;
unsigned int    currstat;

/* global string variables */
/* filename */

char * const fn_passwd = FN_PASSWD;
char * const fn_board = FN_BOARD;
const char * const fn_plans = "plans";
const char * const fn_writelog = "writelog";
const char * const fn_talklog = "talklog";
const char * const fn_overrides = FN_OVERRIDES;
const char * const fn_reject = FN_REJECT;
const char * const fn_notes = "notes";
const char * const fn_water = FN_WATER;
const char * const fn_visable = FN_VISABLE;
const char * const fn_mandex = "/.Names";
const char * const fn_boardlisthelp = FN_BRDLISTHELP;
const char * const fn_boardhelp = FN_BOARDHELP;

/* are descript in userec.loginview */

char           * const loginview_file[NUMVIEWFILE][2] = {
    {"", "(�O�d)"},
    {FN_TOPSONG, "�߱��I���Ʀ�]"},
    {"etc/topusr", "�Q�j�Ʀ�]"},
    {"etc/topusr100", "�ʤj�Ʀ�]"},
    {"", "(�O�d)"},
    {"", "(�O�d)"},
    {"etc/day", "����Q�j���D"},
    {"etc/week", "�@�g���Q�j���D"},
    {"etc/today", "���ѤW���H��"},
    {"etc/yesterday", "�Q��W���H��"},
    {"etc/history", "���v�W������"},
    {"etc/topboardman", "��ذϱƦ�]"},
    {"etc/topboard.tmp", "�ݪO�H��Ʀ�]"},
    {NULL, NULL}
};

/* message */
char           * const msg_separator = MSG_SEPARATOR;

char           * const msg_cancel = MSG_CANCEL;
char           * const msg_usr_left = MSG_USR_LEFT;

char           * const msg_sure_ny = MSG_SURE_NY;
char           * const msg_sure_yn = MSG_SURE_YN;

char           * const msg_bid = MSG_BID;
char           * const msg_uid = MSG_UID;

char           * const msg_del_ok = MSG_DEL_OK;
char           * const msg_del_ny = MSG_DEL_NY;

char           * const msg_fwd_ok = MSG_FWD_OK;
char           * const msg_fwd_err1 = MSG_FWD_ERR1;
char           * const msg_fwd_err2 = MSG_FWD_ERR2;

char           * const err_board_update = ERR_BOARD_UPDATE;
char           * const err_bid = ERR_BID;
char           * const err_uid = ERR_UID;
char           * const err_filename = ERR_FILENAME;

char           * const str_mail_address = "." BBSUSER "@" MYHOSTNAME;
char           * const str_reply = "Re:";
char           * const str_forward = "Fw:";
char           * const str_legacy_forward = "[���]";
char           * const str_space = " \t\n\r";
char           * const str_sysop = "SYSOP";
char           * const str_author1 = STR_AUTHOR1;
char           * const str_author2 = STR_AUTHOR2;
char           * const str_post1 = STR_POST1;
char           * const str_post2 = STR_POST2;
char           * const BBSName = BBSNAME;

/* MAX_MODES is defined in common.h */

char           * const ModeTypeTable[] = {
    "�o�b",			/* IDLE */
    "�D���",			/* MMENU */
    "�t�κ��@",			/* ADMIN */
    "�l����",			/* MAIL */
    "��Ϳ��",			/* TMENU */
    "�ϥΪ̿��",		/* UMENU */
    "XYZ ���",			/* XMENU */
    "�����ݪO",			/* CLASS */
    "Play���",			/* PMENU */
    "�s�S�O�W��",		/* NMENU */
    BBSMNAME2 "�q�c��",		/* PSALE */
    "�o��峹",			/* POSTING */
    "�ݪO�C��",			/* READBRD */
    "�\\Ū�峹",		/* READING */
    "�s�峹�C��",		/* READNEW */
    "��ܬݪO",			/* SELECT */
    "Ū�H",			/* RMAIL */
    "�g�H",			/* SMAIL */
    "��ѫ�",			/* CHATING */
    "��L",			/* XMODE */
    "�M��n��",			/* FRIEND */
    "�W�u�ϥΪ�",		/* LAUSERS */
    "�ϥΪ̦W��",		/* LUSERS */
    "�l�ܯ���",			/* MONITOR */
    "�I�s",			/* PAGE */
    "�d��",			/* TQUERY */
    "���",			/* TALK  */
    "�s�W����",			/* EDITPLAN */
    "�sñ�W��",			/* EDITSIG */
    "�벼��",			/* VOTING */
    "�]�w���",			/* XINFO */
    "�H������",			/* MSYSOP */
    "�L�L�L",			/* WWW */
    "���j�ѤG",			/* BIG2 */
    "�^��",			/* REPLY */
    "�Q���y����",		/* HIT */
    "���y�ǳƤ�",		/* DBACK */
    "���O��",			/* NOTE */
    "�s��峹",			/* EDITING */
    "�o�t�γq�i",		/* MAILALL */
    "�N���",			/* MJ */
    "�q���ܤ�",			/* P_FRIEND */
    "�W���~��",			/* LOGIN */
    "�d�r��",			/* DICT */
    "�����P",			/* BRIDGE */
    "���ɮ�",			/* ARCHIE */
    "���a��",			/* GOPHER */
    "��News",			/* NEWS */
    "���Ѳ��;�",		/* LOVE */
    "�s�軲�U��",		/* EDITEXP */
    "�ӽ�IP��}",		/* IPREG */
    "���޿줽��",		/* NetAdm */
    "������~�{",		/* DRINK */
    "�p���",			/* CAL */
    "�s��y�k��",		/* PROVERB */
    "���G��",			/* ANNOUNCE */
    "��y���O",			/* EDNOTE */
    "�^�~½Ķ��",		/* CDICT */
    "�˵��ۤv���~",		/* LOBJ */
    "�߱��I��",			/* OSONG */
    "�P�d���P��",		/* CHICKEN */
    "���m��",			/* TICKET */
    "�q�Ʀr",			/* GUESSNUM */
    "�C�ֳ�",			/* AMUSE */
    "��H�¥մ�",		/* OTHELLO */
    "����l",			/* DICE */
    "�o�����",			/* VICE */
    "�G�G��ing",		/* BBCALL */
    "ú�@��",			/* CROSSPOST */
    "���l��",			/* M_FIVE */
    "21�Iing",			/* JACK_CARD */
    "10�I�bing",		/* TENHALF */
    "�W�ŤE�Q�E",		/* CARD_99 */
    "�����d��",			/* RAIL_WAY */
    "�j�M���",			/* SREG */
    "�U�H��",			/* CHC */
    "�U�t��",			/* DARK */
    "NBA�j�q��",		/* TMPJACK */
    BBSMNAME2 "�d�]�t��",		/* JCEE */
    "���s�峹",			/* REEDIT */
    "������",                   /* BLOGGING */
    "�ݴ�",			/* CHESSWATCHING */
    "�U���",			/* UMODE_GO */
    "[�t�ο��~]",		/* DEBUGSLEEPING */
    "�s����",			/* UMODE_CONN6 */
    "�¥մ�",			/* REVERSI */
    "BBS-Lua",			/* UMODE_BBSLUA */
    "����ʵe",			/* UMODE_ASCIIMOVIE */
    "",
    "",
    "", // 90
    "",
    "",
    "",
    "",
    "",
    "",
    "",
    "",
    "",
    "", // 100
    "",
    "",
    "",
    "",
    "",
    "",
    "",
    "",
    "",
    "", // 110
    "",
    "",
    "",
    "",
    "",
    "",
    "",
    "",
    "",
    "", // 120
    "",
    "",
    "",
    "",
    "",
    "",
    "",
    ""
};

/* term.c */
int             b_lines = 23; // bottom line of screen (= t_lines - 1)
int             t_lines = 24; // term lines
int             p_lines = 20; // ���� header(3), footer(1), �e���W�i�H��ܸ�ƪ����
int             t_columns = 80;

/* refer to ansi.h for *len */
char           * const strtstandout = ANSI_REVERSE;
const int       strtstandoutlen = 4;
char           * const endstandout = ANSI_RESET;
const int        endstandoutlen = 3;
char           * const clearbuf = ESC_STR "[H" ESC_STR "[J";
const int        clearbuflen = 6;
char           * const cleolbuf = ESC_STR "[K";
const int        cleolbuflen = 3;
char           * const scrollrev = ESC_STR "M";
const int       scrollrevlen = 2;
int             automargins = 1;

/* io.c */
time4_t         now;
int             KEY_ESC_arg;
int             watermode = -1;
int             wmofo = NOTREPLYING;
/*
 * PAGER_UI_IS(PAGER_UI_ORIG) | PAGER_UI_IS(PAGER_UI_NEW):
 * ????????????????????
 * Ptt ���y�^�U   (FIXME: guessed by scw)
 * watermode = -1 �S�b�^���y
 *           = 0   �b�^�W�@�����y  (Ctrl-R)
 *           > 0   �b�^�e n �����y (Ctrl-R Ctrl-R)
 *
 * PAGER_UI_IS(PAGER_UI_OFO)  by in2
 * wmofo     = NOTREPLYING     �S�b�^���y
 *           = REPLYING        ���b�^���y
 *           = RECVINREPLYING  �^���y���S������y
 *
 * wmofo     >=0  �ɦ�����y�N�u���, ���|��water[]��,
 *                �ݦ^�����y���ɭԤ@���g�J.
 */


/* cache.c */
SHM_t          *SHM;
boardheader_t  *bcache;
userinfo_t     *currutmp;

/* read.c */
int             TagNum = 0;		/* tag's number */
int		TagBoard = -1;		/* TagBoard = 0 : user's mailbox */
                                        /* TagBoard > 0 : bid where last taged */
char            currdirect[64];		/* XXX TODO change this to PATHLEN? */

/* bbs.c */
char            real_name[IDLEN + 1];

/* mbbsd.c */
char            fromhost[STRLEN] = "\0";
char		fromhost_masked[32] = "\0"; // masked 'fromhost'
char            from_cc[STRLEN] = "\0";
char            water_usies = 0;
char            is_first_login_of_today = 0;
char            is_login_ready = 0;
FILE           *fp_writelog = NULL;
water_t         *water, *swater[WB_OFO_USER_NUM], *water_which;

/* chc_play.c */

/* user.c */
#ifdef CHESSCOUNTRY
int user_query_mode;
/*
 * user_query_mode = 0  simple data
 *                 = 1  gomoku chess country data
 *                 = 2  chc chess country data
 *                 = 3  go chess country data
 */
#endif /* defined(CHESSCOUNTRY) */

/* screen.c */
#define SCR_COLS        ANSILINELEN
screenline_t   *big_picture = NULL;
char            roll = 0;
char		msg_occupied = 0;

/* gomo.c */
const char     * const bw_chess[] = {"��", "��", "�C", "�E"};

/* friend.c */
/* Ptt �U�دS�O�W�檺�ɦW */
char           *friend_file[8] = {
    FN_OVERRIDES,
    FN_REJECT,
    FN_ALOHAED,
    "", /* deprecated: post list */
    "", /* may point to other filename */
    FN_CANVOTE,
    FN_WATER,
    FN_VISABLE
};

#ifdef NOKILLWATERBALL
char    reentrant_write_request = 0;
#endif

#ifdef PTTBBS_UTIL
    #define COMMON_TIME (time(0))
#else
    #define COMMON_TIME (now)
#endif
