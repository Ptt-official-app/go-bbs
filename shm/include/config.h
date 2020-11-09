/* $Id$ */
#ifndef INCLUDE_CONFIG_H
#define INCLUDE_CONFIG_H

#include <syslog.h>
#include "../pttbbs.conf"

#define BBSPROG         BBSHOME "/bin/mbbsd"         /* �D�{�� */
#define BAN_FILE        "BAN"                        /* �����q�i�� */
#define LOAD_FILE       "/proc/loadavg"              /* for Linux */

/////////////////////////////////////////////////////////////////////////////
// System Name Configuration �t�ΦW�ٳ]�w

/* �t�ΦW(�l���)�A��ĳ�O�W�L 3 �Ӧr���C �Ԩ� sample/pttbbs.conf */
#ifndef BBSMNAME
#define BBSMNAME        "Ptt"
#endif

/* �t�ΦW(����)�A��ĳ��n 4 �Ӧr���C �Ԩ� sample/pttbbs.conf */
#ifndef BBSMNAME2
#define BBSMNAME2       "��tt"
#endif

/* �����W�A��ĳ��n 3 �Ӧr���C �Ԩ� sample/pttbbs.conf */
#ifndef MONEYNAME
#define MONEYNAME       BBSMNAME "��"
#endif

/* AID ��ܪ����x�W�١C �Y IP �Ӫ��Хt��w�q�C */
#ifndef AID_HOSTNAME
#define AID_HOSTNAME    MYHOSTNAME
#endif

/////////////////////////////////////////////////////////////////////////////
// Themes �D�D�t��

#ifndef TITLE_COLOR 
#define TITLE_COLOR             ANSI_COLOR(0;1;37;46)   /* �D�e���W����D�C */
#endif

#ifndef HLP_CATEGORY_COLOR
#define HLP_CATEGORY_COLOR      ANSI_COLOR(0;1;32)  /* ������椺������ */
#endif

#ifndef HLP_DESCRIPTION_COLOR
#define HLP_DESCRIPTION_COLOR   ANSI_COLOR(0)       /* ������椺������ */
#endif

#ifndef HLP_KEYLIST_COLOR
#define HLP_KEYLIST_COLOR       ANSI_COLOR(0;1;36)  /* ������椺���䶵 */
#endif

/////////////////////////////////////////////////////////////////////////////
// OS Settings �@�~�t�ά����]�w

#ifndef BBSUSER
#define BBSUSER "bbs"
#endif

#ifndef BBSUID
#define BBSUID (9999)
#endif

#ifndef BBSGID
#define BBSGID (99)
#endif

#ifndef TAR_PATH
#define TAR_PATH "tar"
#endif

#ifndef MUTT_PATH
#define MUTT_PATH "mutt"
#endif

#ifndef MAXPATHLEN
#define MAXPATHLEN (256)
#endif

#ifndef PATHLEN
#define PATHLEN (256)
#endif

#ifndef DEFAULT_FOLDER_CREATE_PERM
#define DEFAULT_FOLDER_CREATE_PERM (0755)
#endif

#ifndef DEFAULT_FILE_CREATE_PERM
#define DEFAULT_FILE_CREATE_PERM (0644)
#endif

#ifndef SHM_KEY
#define SHM_KEY         1228
#endif

#ifndef PASSWDSEM_KEY
#define PASSWDSEM_KEY   2010    /* semaphore key */
#endif

#ifndef SYSLOG_FACILITY
#define SYSLOG_FACILITY   LOG_LOCAL0
#endif

#ifndef MEM_CHECK
#define MEM_CHECK 0x98761234
#endif

#ifndef RELAY_SERVER_IP                     /* �H���~�H�� mail server */
#define RELAY_SERVER_IP    "127.0.0.1"
#endif

#ifndef XCHATD_ADDR
#define XCHATD_ADDR     ":3838"
#endif

/////////////////////////////////////////////////////////////////////////////
// Default Board Names �w�]�ݪO�W��

#ifndef BN_BUGREPORT
#define BN_BUGREPORT "SYSOP"
#endif

#ifndef BN_SYSOP
#define BN_SYSOP "SYSOP"
#endif

#ifndef BN_ID_PROBLEM
#define BN_ID_PROBLEM "SYSOP"
#endif

#ifndef BN_LAW
#define BN_LAW  BBSMNAME "Law"
#endif

#ifndef BN_NEWBIE
#define BN_NEWBIE BBSMNAME "Newhand"
#endif

#ifndef BN_TEST
#define BN_TEST "Test"
#endif

#ifndef BN_NOTE
#define BN_NOTE "Note"
#endif

#ifndef BN_SECURITY
#define BN_SECURITY "Security"
#endif

#ifndef BN_RECORD
#define BN_RECORD "Record"
#endif

#ifndef BN_FOREIGN
#define BN_FOREIGN BBSMNAME "Foreign"
#endif

#ifndef BN_DELETED
#define BN_DELETED "deleted"
#endif

#ifndef BN_JUNK
#define BN_JUNK "junk"
#endif 

#ifndef BN_POLICELOG
#define BN_POLICELOG    "PoliceLog"
#endif

#ifndef BN_UNANONYMOUS
#define BN_UNANONYMOUS "UnAnonymous"
#endif

#ifndef BN_NEWIDPOST
#define BN_NEWIDPOST "NEWIDPOST"
#endif

#ifndef BN_ALLPOST
#define BN_ALLPOST "ALLPOST"
#endif

#ifndef BN_ALLHIDPOST
#define BN_ALLHIDPOST "ALLHIDPOST"
#endif

/////////////////////////////////////////////////////////////////////////////
// Performance Parameters �į�Ѽ�

#ifndef MAX_USERS                        /* �̰����U�H�� */
#define MAX_USERS         (150000)
#endif

#ifndef MAX_ACTIVE
#define MAX_ACTIVE        (1024)         /* �̦h�P�ɤW���H�� */
#endif

#ifndef MAX_GUEST
#define MAX_GUEST         (100)          /* �̦h guest �W���H�� */
#endif

#ifndef MAX_CPULOAD
#define MAX_CPULOAD       (70)           /* CPU �̰�load */
#endif

#ifndef DEBUGSLEEP_SECONDS
#define DEBUGSLEEP_SECONDS (3600)        /* debug ���ݮɶ� */
#endif

#ifndef MAX_BOARD
#define MAX_BOARD         (8192)         /* �̤j�}�O�Ӽ� */
#endif

#ifndef HASH_BITS
#define HASH_BITS         (16)           /* userid->uid hashing bits */
#endif

#ifndef OVERLOADBLOCKFDS
#define OVERLOADBLOCKFDS  (0)            /* �W����|�O�d�o��h�� fd */
#endif

#ifndef MAX_FRIEND
#define MAX_FRIEND        (256)          /* ���J cache ���̦h�B�ͼƥ� */
#endif

#ifndef MAX_REJECT
#define MAX_REJECT        (32)           /* ���J cache ���̦h�a�H�ƥ� */
#endif

#ifndef MAX_MSGS
#define MAX_MSGS          (10)           /* ���y(���T)�ԭ@�W�� */
#endif

#ifndef MAX_ADBANNER
#define MAX_ADBANNER      (500)          /* �̦h�ʺA�ݪO�� */
#endif

#ifndef MAX_SWAPUSED
#define MAX_SWAPUSED      (0.7)          /* SWAP�̰��ϥβv */
#endif

#ifndef HOTBOARDCACHE
#define HOTBOARDCACHE     (0)            /* �����ݪO�֨� */
#endif

#ifndef TARQUEUE_TIME_STR
#define TARQUEUE_TIME_STR   "�`�]"       // �ݪO�ƥ��ɶ��T�� (���P contab �@�P)
#endif

/////////////////////////////////////////////////////////////////////////////
// More system messages �t�ΰT��

#ifndef RECYCLE_BIN_NAME
#define RECYCLE_BIN_NAME "�귽�^����" // "�U����"
#endif

#ifndef RECYCLE_BIN_OWNER
#define RECYCLE_BIN_OWNER "[" RECYCLE_BIN_NAME "]"
#endif

#ifndef TIME_CAPSULE_NAME
#define TIME_CAPSULE_NAME "Magical Index" // "Time Capsule"
#endif

/////////////////////////////////////////////////////////////////////////////
// Site settings ���x�\��]�w

#ifndef MAX_POST_MONEY                      /* �o��峹�Z�O���W�� */
#define MAX_POST_MONEY          (100)
#endif

#ifndef MAX_CHICKEN_MONEY                   /* �i������Q�W�� */
#define MAX_CHICKEN_MONEY       (100)
#endif

#ifndef MAX_GUEST_LIFE                      /* �̪����{�ҨϥΪ̫O�d�ɶ�(��) */
#define MAX_GUEST_LIFE          (3 * 24 * 60 * 60)
#endif

#ifndef MAX_EDIT_LINE
#define MAX_EDIT_LINE           (2048)      /* �峹�̪��s����� */
#endif 

#ifndef MAX_EDIT_LINE_LARGE                 // �j�ɳ̪��s�����
#define MAX_EDIT_LINE_LARGE     (32000)
#endif

#ifndef MAX_LIFE                            /* �̪��ϥΪ̫O�d�ɶ�(��) */
#define MAX_LIFE                (120 * 24 * 60 * 60)
#endif

#ifndef KEEP_DAYS_REGGED
#define KEEP_DAYS_REGGED        (120)       /* �w���U�ϥΪ̫O�d�h�[ */
#endif

#ifndef KEEP_DAYS_UNREGGED
#define KEEP_DAYS_UNREGGED      (15)        /* �����U�ϥΪ̫O�d�h�[ */
#endif

#ifndef MAX_FROM
#define MAX_FROM                (300)       /* �̦h�G�m�� */
#endif

#ifndef THREAD_SEARCH_RANGE
#define THREAD_SEARCH_RANGE     (500)       /* �t�C�峹�j�M�W�� */
#endif

#ifndef FOREIGN_REG_DAY
#define FOREIGN_REG_DAY         (30)        /* �~�y�ϥΪ̸եΤ���W�� */
#endif

#ifndef FORCE_PROCESS_REGISTER_FORM
#define FORCE_PROCESS_REGISTER_FORM     (0)
#endif

#ifndef HBFLexpire
#define HBFLexpire        (432000)      /* 5 days */
#endif

#ifndef MAX_EXKEEPMAIL
#define MAX_EXKEEPMAIL    (1000)        /* �̦h�H�c�[�j�h�֫� */
#endif

#ifndef INNTIMEZONE
#define INNTIMEZONE       "+0000 (UTC)" /* ��H�� timestamp ���ɰ� */
#endif

#ifndef ADD_EXMAILBOX
#define ADD_EXMAILBOX     0             /* �ذe�H�c */
#endif


#ifndef BADPOST_CLEAR_DURATION
#define BADPOST_CLEAR_DURATION  (180)   // ���H��ɶ�����
#endif

#ifndef BADPOST_MIN_CLEAR_DURATION
#define BADPOST_MIN_CLEAR_DURATION (3)  // �H�孺���ɶ�����
#endif

#ifndef MAX_CROSSNUM
#define MAX_CROSSNUM      (9)           /* �̦hcrosspost���� */
#endif

/* (deprecated) more.c ���峹���ƤW��(lines/22), +4 for safe */
#ifndef MAX_PAGES
#define MAX_PAGES         (MAX_EDIT_LINE / 22 + 4)
#endif

#ifndef MAX_ADBANNER_SECTION
#define MAX_ADBANNER_SECTION (10)       /* �̦h�ʺA�ݪO���O */
#endif

#ifndef MAX_ADBANNER_HEIGHT
#define MAX_ADBANNER_HEIGHT  (11)       /* �̤j�ʺA�ݪO���e���� */
#endif

#ifndef MAX_QUERYLINES
#define MAX_QUERYLINES    (16)          /* ��� Query/Plan �T���̤j��� */
#endif

#ifndef MAX_LOGIN_INFO
#define MAX_LOGIN_INFO    (128)         /* �̦h�W�u�q���H�� */
#endif

#ifndef MAX_POST_INFO
#define MAX_POST_INFO     (32)          /* �̦h�s�峹�q���H�� */
#endif

#ifndef MAX_NAMELIST
#define MAX_NAMELIST      (128)         /* �̦h��L�S�O�W��H�� */
#endif

#ifndef MAX_NOTE
#define MAX_NOTE          (20)          /* �̦h�O�d�X�g�d���H */
#endif

#ifndef MAX_SIGLINES
#define MAX_SIGLINES      (6)           /* ñ�W�ɤޤJ�̤j��� */
#endif

#ifndef MAX_REVIEW
#define MAX_REVIEW        (7)           /* �̦h���y�^�U */
#endif

#ifndef NUMVIEWFILE
#define NUMVIEWFILE       (14)          /* �i���e���̦h�� */
#endif

#ifndef LOGINATTEMPTS
#define LOGINATTEMPTS     (3)           /* �̤j�i�����~���� */
#endif

#ifndef MAX_KEEPMAIL
#define MAX_KEEPMAIL            (200)   /* �@�� user �̦h�O�d�X�� MAIL�H */
#endif

#ifndef MAX_KEEPMAIL_SOFTLIMIT
#define MAX_KEEPMAIL_SOFTLIMIT  (2500)  /* �� admin �~�A�L�k�H�����H */
#endif

#ifndef MAX_KEEPMAIL_HARDLIMIT
#define MAX_KEEPMAIL_HARDLIMIT  (20000) /* �H�c�ƶq���W���A�W�L�N�����H�H */
#endif

#ifndef BADCIDCHARS
#define BADCIDCHARS     " *"            /* Chat Room ���T�Ω� nick ���r�� */
#endif

#ifndef MAX_ROOM
#define MAX_ROOM        (16)            /* ��ѫǳ̦h���X���]�[�H */
#endif

#ifndef MAXTAGS
#define MAXTAGS         (255)           /* t(tag) ���̤j�ƶq */
#endif

#ifndef WRAPMARGIN
#define WRAPMARGIN      (511)           /* �s�边 wrap ���� */
#endif

#ifdef USE_MASKED_FROMHOST
#define FROMHOST    fromhost_masked
#else
#define FROMHOST    fromhost
#endif

/////////////////////////////////////////////////////////////////////////////
// Logging �O���]�w

#ifndef LOG_CONF_KEYWORD        // �O���j�M������r
#define LOG_CONF_KEYWORD        (0)
#endif
#ifndef LOG_CONF_INTERNETMAIL   // �O�� internet outgoing mail
#define LOG_CONF_INTERNETMAIL   (0)
#endif
#ifndef LOG_CONF_PUSH           // �O������
#define LOG_CONF_PUSH           (0)
#endif
#ifndef LOG_CONF_EDIT_CALENDAR  // �O���s���ƾ�
#define LOG_CONF_EDIT_CALENDAR  (0)
#endif
#ifndef LOG_CONF_POST           // �O���o��
#define LOG_CONF_POST           (0)
#endif
#ifndef LOG_CONF_CRAWLER        // �O�� crawlers
#define LOG_CONF_CRAWLER        (0)
#endif
#ifndef LOG_CONF_CROSSPOST      // �O�����
#define LOG_CONF_CROSSPOST      (0)
#endif
#ifndef LOG_CONF_BAD_REG_CODE   // �O�����������U�X
#define LOG_CONF_BAD_REG_CODE   (0)
#endif
#ifndef LOG_CONF_VALIDATE_REG   // �O���f�ֵ��U��
#define LOG_CONF_VALIDATE_REG   (0)
#endif
#ifndef LOG_CONF_MASS_DELETE    // �O���j�q�R���ɮ�
#define LOG_CONF_MASS_DELETE    (0)
#endif
#ifndef LOG_CONF_OSONG_VERBOSE  // �Բ��I���O��
#define LOG_CONF_OSONG_VERBOSE  (0)
#endif
#ifndef LOG_CONF_EDIT_TITLE     // �s����D�O��
#define LOG_CONF_EDIT_TITLE     (0)
#endif

/////////////////////////////////////////////////////////////////////////////
// Default Configurations �w�]�Ѽ�

// �Y�Q���ΤU�C�ѼƽЦb pttbbs.conf �w�q NO_XXX (ex, NO_LOGINASNEW)
#ifndef NO_LOGINASNEW
#define    LOGINASNEW           /* �ĥΤW���ӽбb����� */
#endif

#ifndef NO_DOTIMEOUT
#define    DOTIMEOUT            /* �B�z���m�ɶ� */
#endif

#ifndef NO_INTERNET_EMAIL
#define    INTERNET_EMAIL       /* �䴩 InterNet Email �\��(�t Forward) */
#endif

#ifndef NO_SHOWUID
#define    SHOWUID              /* �����i�ݨ��ϥΪ� UID */
#endif

#ifndef NO_SHOWBOARD
#define    SHOWBOARD            /* �����i�ݨ��ϥΪ̬ݪO */
#endif

#ifndef NO_SHOWPID
#define    SHOWPID              /* �����i�ݨ��ϥΪ� PID */
#endif

#ifndef NO_HAVE_ANONYMOUS
#define    HAVE_ANONYMOUS       /* ���� Anonymous �O */
#endif

#ifndef NO_HAVE_ORIGIN
#define    HAVE_ORIGIN          /* ��� author �Ӧۦ�B */
#endif

#ifndef NO_USE_BSMTP
#define    USE_BSMTP            /* �ϥ�opus��BSMTP �H���H? */
#endif

#ifndef NO_REJECT_FLOOD_POST
#define    REJECT_FLOOD_POST    /* ����BlahBlah����� */
#endif

// #define  HAVE_INFO               /* ��ܵ{���������� */
// #define  HAVE_LICENSE            /* ��� GNU ���v�e�� */
// #define  HAVE_REPORT             /* (��H)�t�ΰl�ܳ��i */

#ifdef  DOTIMEOUT
# define IDLE_TIMEOUT    (43200) /* �@�뱡�p�� timeout (12hr) */
# define SHOW_IDLE_TIME          /* ��ܶ��m�ɶ� */
#endif

#endif
