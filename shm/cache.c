#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <signal.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include <unistd.h>
#include <errno.h>
#include <fcntl.h>
#include "cmsys.h"
#include "cmbbs.h"
#include "common.h"
#include "var.h"

#include "modes.h" // for DEBUGSLEEPING

//////////////////////////////////////////////////////////////////////////
// This is shared by utility library and core BBS,
// so do not put code using currutmp/cuser here.
//////////////////////////////////////////////////////////////////////////

// these cannot be used!
#define currutmp  YOU_FAILED
#define usernum	  YOU_FAILED
#undef  cuser
#undef  cuser_ref
#define cuser     YOU_FAILED
#define cuser_ref YOU_FAILED
#define abort_bbs YOU_FAILED
#define log_usies YOU_FAILED

/*
 * the reason for "safe_sleep" is that we may call sleep during SIGALRM
 * handler routine, while SIGALRM is blocked. if we use the original sleep,
 * we'll never wake up.
 */
unsigned int
safe_sleep(unsigned int seconds)
{
    /* jochang  sleep�����D�ɥ� */
    sigset_t        set, oldset;

    sigemptyset(&set);
    sigprocmask(SIG_BLOCK, &set, &oldset);
    if (sigismember(&oldset, SIGALRM)) {
	unsigned int    retv;
	// log_usies("SAFE_SLEEP ", "avoid hang");
	sigemptyset(&set);
	sigaddset(&set, SIGALRM);
	sigprocmask(SIG_UNBLOCK, &set, NULL);
	retv = sleep(seconds);
	sigprocmask(SIG_BLOCK, &set, NULL);
	return retv;
    }
    return sleep(seconds);
}

/*
 * section - SHM
 */
static void
attach_err(int shmkey, const char *name)
{
    fprintf(stderr, "[%s error] key = %x\n", name, shmkey);
    fprintf(stderr, "errno = %d: %s\n", errno, strerror(errno));
    exit(1);
}

void           *
attach_shm(int shmkey, int shmsize)
{
    void           *shmptr = (void *)NULL;
    int             shmid;

    shmid = shmget(shmkey, shmsize,
#ifdef USE_HUGETLB
	    SHM_HUGETLB |
#endif
	    0);
    if (shmid < 0) {
	// SHM should be created by uhash_loader, NOT mbbsd or other utils
	attach_err(shmkey, "shmget");
    } else {
	shmptr = (void *)shmat(shmid, NULL, 0);
	if (shmptr == (void *)-1)
	    attach_err(shmkey, "shmat");
    }

    return shmptr;
}

static void 
shm_check_error()
{
    fprintf(stderr, "Please use the source code version corresponding to SHM,\n"
	    "or use ipcrm(1) command to clean share memory.\n");
    exit(1);
}

void
attach_check_SHM(int version, int SHM_t_size)
{
    SHM = attach_shm(SHM_KEY, SHMSIZE);

    // check main program -> common bbs library
    if (version      != SHM_VERSION) {
	fprintf(stderr, "Error: version(%d) != SHM_VERSION(%d)\n",
		version, SHM_VERSION);
	shm_check_error();
    }
    if (SHM_t_size   != sizeof(SHM_t)) {
	fprintf(stderr, "Error: SHM_t_size(%d) != sizeof(SHM_t)(%zd)\n",
		SHM_t_size, sizeof(SHM_t));
	shm_check_error();
    }
    // check common bbs library -> SHM
    if (SHM->version != SHM_VERSION) {
	fprintf(stderr, "Error: SHM->version(%d) != SHM_VERSION(%d)\n", 
		SHM->version, SHM_VERSION);
	shm_check_error();
    }
    if (SHM->size    != sizeof(SHM_t)) {
	fprintf(stderr, "Error: SHM->size(%d) != sizeof(SHM_t)(%zd)\n", 
		SHM->size, sizeof(SHM_t));
	shm_check_error();
    }

    if (!SHM->loaded)		/* (uhash) assume fresh shared memory is
				 * zeroed */
	exit(1);
    if (SHM->Btouchtime == 0)
	SHM->Btouchtime = 1;
    bcache = SHM->bcache;

    if (SHM->Ptouchtime == 0)
	SHM->Ptouchtime = 1;

    if (SHM->Ftouchtime == 0)
	SHM->Ftouchtime = 1;
}

/*
 * section - user cache(including uhash)
 */

/* uhash ****************************************** */
/*
 * the design is this: we use another stand-alone program to create and load
 * data into the hash. (that program could be run in rc-scripts or something
 * like that) after loading completes, the stand-alone program sets loaded to
 * 1 and exits.
 * 
 * the bbs exits if it can't attach to the shared memory or the hash is not
 * loaded yet.
 */

void
add_to_uhash(int n, const char *id)
{
    int            *p, h = StringHash(id)%(1<<HASH_BITS);
    int             times;
    strlcpy(SHM->userid[n], id, sizeof(SHM->userid[n]));

    p = &(SHM->hash_head[h]);

    for (times = 0; times < MAX_USERS && *p != -1; ++times)
	p = &(SHM->next_in_hash[*p]);

    if (times >= MAX_USERS)
    {
	// abort_bbs(0);
	fprintf(stderr, "add_to_uhash: exceed max users.\r\n");
	exit(0);
    }

    SHM->next_in_hash[*p = n] = -1;
}

void
remove_from_uhash(int n)
{
/*
 * note: after remove_from_uhash(), you should add_to_uhash() (likely with a
 * different name)
 */
    int             h = StringHash(SHM->userid[n])%(1<<HASH_BITS);
    int            *p = &(SHM->hash_head[h]);
    int             times;

    for (times = 0; times < MAX_USERS && (*p != -1 && *p != n); ++times)
	p = &(SHM->next_in_hash[*p]);

    if (times >= MAX_USERS)
    {
	// abort_bbs(0);
	fprintf(stderr, "remove_from_uhash: current SHM exceed max users.\r\n");
	exit(0);
    }

    if (*p == n)
	*p = SHM->next_in_hash[n];
}

#if (1<<HASH_BITS)*10 < MAX_USERS
#warning "Suggest to use bigger HASH_BITS for better searchuser() performance,"
#warning "searchuser() average chaining MAX_USERS/(1<<HASH_BITS) times."
#endif

int
dosearchuser(const char *userid, char *rightid)
{
    int             h, p, times;
    STATINC(STAT_SEARCHUSER);
    h = StringHash(userid)%(1<<HASH_BITS);
    p = SHM->hash_head[h];

    for (times = 0; times < MAX_USERS && p != -1 && p < MAX_USERS ; ++times) {
	if (strcasecmp(SHM->userid[p], userid) == 0) {
	    if(userid[0] && rightid) strcpy(rightid, SHM->userid[p]);
	    return p + 1;
	}
	p = SHM->next_in_hash[p];
    }

    return 0;
}

int
searchuser(const char *userid, char *rightid)
{
    if(userid[0]=='\0')
	return 0;
    return dosearchuser(userid, rightid);
}

char           *
getuserid(int num)
{
    if (--num >= 0 && num < MAX_USERS)
	return ((char *)SHM->userid[num]);
    return NULL;
}

void
setuserid(int num, const char *userid)
{
    if (num > 0 && num <= MAX_USERS) {
/*  Ptt: it may cause problems
	if (num > SHM->number)
	    SHM->number = num;
	else
*/
        remove_from_uhash(num - 1);
	add_to_uhash(num - 1, userid);
    }
}

userinfo_t     *
search_ulist_pid(int pid)
{
    register int    i = 0, j, start = 0, end = SHM->UTMPnumber - 1;
    int *ulist;
    register userinfo_t *u;
    if (end == -1)
	return NULL;
    ulist = SHM->sorted[SHM->currsorted][8];
    for (i = ((start + end) / 2);; i = (start + end) / 2) {
	u = &SHM->uinfo[ulist[i]];
	j = pid - u->pid;
	if (!j) {
	    return u;
	}
	if (end == start) {
	    break;
	} else if (i == start) {
	    i = end;
	    start = end;
	} else if (j > 0)
	    start = i;
	else
	    end = i;
    }
    return 0;
}

userinfo_t     *
search_ulistn(int uid, int unum)
{
    register int    i = 0, j, start = 0, end = SHM->UTMPnumber - 1;
    int *ulist;
    register userinfo_t *u;
    if (end == -1)
	return NULL;
    ulist = SHM->sorted[SHM->currsorted][7];
    for (i = ((start + end) / 2);; i = (start + end) / 2) {
	u = &SHM->uinfo[ulist[i]];
	j = uid - u->uid;
	if (j == 0) {
	    for (; i > 0 && uid == SHM->uinfo[ulist[i - 1]].uid; --i)
		;/* ����Ĥ@�� */
	    // piaip Tue Jan  8 09:28:03 CST 2008
	    // many people bugged about that their utmp have invalid
	    // entry on record.
	    // we found them caused by crash process (DEBUGSLEEPING) which
	    // may occupy utmp entries even after process was killed.
	    // because the memory is invalid, it is not safe for those process
	    // to wipe their utmp entry. it should be done by some external
	    // daemon.
	    // however, let's make a little workaround here...
	    for (; unum > 0 && i >= 0 && ulist[i] >= 0 &&
		    SHM->uinfo[ulist[i]].uid == uid; unum--, i++)
	    {
		if (SHM->uinfo[ulist[i]].mode == DEBUGSLEEPING)
		    unum ++;
	    }
	    if (unum == 0 && i > 0 && ulist[i-1] >= 0 &&
		    SHM->uinfo[ulist[i-1]].uid == uid)
		return &SHM->uinfo[ulist[i-1]];
	    /*
	    if ( i + unum - 1 >= 0 &&
		 (ulist[i + unum - 1] >= 0 &&
		  uid == SHM->uinfo[ulist[i + unum - 1]].uid ) )
		return &SHM->uinfo[ulist[i + unum - 1]];
		*/
	    break;		/* �W�L�d�� */
	}
	if (end == start) {
	    break;
	} else if (i == start) {
	    i = end;
	    start = end;
	} else if (j > 0)
	    start = i;
	else
	    end = i;
    }
    return 0;
}

userinfo_t     *
search_ulist_userid(const char *userid)
{
    register int    i = 0, j, start = 0, end = SHM->UTMPnumber - 1;
    int *ulist;
    register userinfo_t * u;
    if (end == -1)
	return NULL;
    ulist = SHM->sorted[SHM->currsorted][0];
    for (i = ((start + end) / 2);; i = (start + end) / 2) {
	u = &SHM->uinfo[ulist[i]];
	j = strcasecmp(userid, u->userid);
	if (!j) {
	    return u;
	}
	if (end == start) {
	    break;
	} else if (i == start) {
	    i = end;
	    start = end;
	} else if (j > 0)
	    start = i;
	else
	    end = i;
    }
    return 0;
}

/*
 * section - money cache
 */
int
setumoney(int uid, int money)
{
    SHM->money[uid - 1] = money;
    passwd_update_money(uid);
    return SHM->money[uid - 1];
}

int
deumoney(int uid, int money)
{
    if (uid <= 0 || uid > MAX_USERS){
	fprintf(stderr, "internal error: deumoney(%d, %d)\r\n", uid, money);
	return -1;
    }

    if (money < 0 && moneyof(uid) < -money)
	return setumoney(uid, 0);
    else
	return setumoney(uid, SHM->money[uid - 1] + money);
}

/*
 * section - board cache
 */
void touchbtotal(int bid) {
    assert(0<=bid-1 && bid-1<MAX_BOARD);
    SHM->total[bid - 1] = 0;
    SHM->lastposttime[bid - 1] = 0;
}

/**
 * qsort comparison function - �ӪO�W�Ƨ�
 */
static int
cmpboardname(const void * i, const void * j)
{
    return strcasecmp(bcache[*(int*)i].brdname, bcache[*(int*)j].brdname);
}

/**
 * qsort comparison function - ���Ӹs�ձƧǡB�P�@�Ӹs�դ��̪O�W��
 */
static int
cmpboardclass(const void * i, const void * j)
{
    boardheader_t *brd1 = &bcache[*(int*)i], *brd2 = &bcache[*(int*)j];
    int cmp;

    cmp=strncmp(brd1->title, brd2->title, 4);
    if(cmp!=0) return cmp;
    return strcasecmp(brd1->brdname, brd2->brdname);
}


void
sort_bcache(void)
{
    int             i;
    /* critical section �ɶq���n�I�s  */
    /* �u���s�W �β����ݪO �ݭn�I�s�� */
    if(SHM->Bbusystate) {
	sleep(1);
	return;
    }
    SHM->Bbusystate = 1;
    for (i = 0; i < SHM->Bnumber; i++) {
	SHM->bsorted[BRD_GROUP_LL_TYPE_NAME][i] = i;
	SHM->bsorted[BRD_GROUP_LL_TYPE_CLASS][i] = i;
    }
    qsort(SHM->bsorted[BRD_GROUP_LL_TYPE_NAME],
	  SHM->Bnumber, sizeof(int), cmpboardname);
    qsort(SHM->bsorted[BRD_GROUP_LL_TYPE_CLASS],
	  SHM->Bnumber, sizeof(int), cmpboardclass);

    for (i = 0; i < SHM->Bnumber; i++) {
	bcache[i].firstchild[BRD_GROUP_LL_TYPE_NAME] = 0;
	bcache[i].firstchild[BRD_GROUP_LL_TYPE_CLASS] = 0;
    }
    SHM->Bbusystate = 0;
}

void
reload_bcache(void)
{
    int     i, fd;
    pid_t   pid;
    for( i = 0 ; i < 10 && SHM->Bbusystate ; ++i ){
	fprintf(stderr, "SHM->Bbusystate is currently locked (value: %d). "
	       "please wait... \r\n", SHM->Bbusystate);
	sleep(1);
    }

    SHM->Bbusystate = 1;
    if ((fd = open(fn_board, O_RDONLY)) > 0) {
	SHM->Bnumber =
	    read(fd, bcache, MAX_BOARD * sizeof(boardheader_t)) /
	    sizeof(boardheader_t);
	close(fd);
    }
    memset(SHM->lastposttime, 0, MAX_BOARD * sizeof(time4_t));
    memset(SHM->total, 0, MAX_BOARD * sizeof(int));

    /* ���Ҧ� boards ��Ƨ�s��A�]�w uptime */
    SHM->Buptime = SHM->Btouchtime;
    // log_usies("CACHE", "reload bcache");
    fprintf(stderr, "cache: reload bcache\r\n");
    SHM->Bbusystate = 0;
    sort_bcache();

    fprintf(stderr, "load bottom in background\r\n");
    if( (pid = fork()) > 0 )
	return;
    setproctitle("loading bottom");
    for( i = 0 ; i < MAX_BOARD ; ++i )
	if( SHM->bcache[i].brdname[0] ){
	    char    fn[PATHLEN];
	    int n;
	    setbfile(fn, SHM->bcache[i].brdname, FN_DIR_BOTTOM);
	    n = get_num_records(fn, sizeof(fileheader_t));
	    if( n > 5 )
		n = 5;
	    SHM->n_bottom[i] = n;
	}
    fprintf(stderr, "load bottom done\r\n");
    if( pid == 0 )
	exit(0);
    // if pid == -1 should be returned
}

void resolve_boards(void)
{
    while (SHM->Buptime < SHM->Btouchtime) {
	reload_bcache();
    }
}

int num_boards(void)
{
    return SHM->Bnumber;
}

void addbrd_touchcache(void)
{
    SHM->Bnumber++;
    reset_board(num_boards());
    sort_bcache();
}

void
reset_board(int bid) /* XXXbid: from 1 */
{				/* Ptt: �o�˴N���ΦѬOtouch board�F */
    int             fd;
    boardheader_t  *bhdr;

    if (--bid < 0)
	return;
    assert(0<=bid && bid<MAX_BOARD);
    if (SHM->Bbusystate || COMMON_TIME - SHM->busystate_b[bid] < 10) {
	safe_sleep(1);
    } else {
	SHM->busystate_b[bid] = COMMON_TIME;

	bhdr = bcache;
	bhdr += bid;
	if ((fd = open(fn_board, O_RDONLY)) >= 0) {
	    lseek(fd, (off_t) (bid * sizeof(boardheader_t)), SEEK_SET);
	    read(fd, bhdr, sizeof(boardheader_t));
	    close(fd);
	}
	SHM->busystate_b[bid] = 0;

	buildBMcache(bid + 1); /* XXXbid */
    }
}

void
resolve_board_group(const int gid, const int type)
{
    boardheader_t  *bptr, *currbptr, *parent;
    int             bid, n, childcount = 0;
    int             boardcount;
    assert(0<=type && type<2);
    assert(0<= gid-1 && gid-1<MAX_BOARD);
    currbptr = parent = &bcache[gid - 1];
    boardcount = num_boards();
    assert(0<=boardcount && boardcount<=MAX_BOARD);
    for (n = 0; n < boardcount; ++n) {
	bid = SHM->bsorted[type][n]+1;
	if( bid<=0 || !(bptr = getbcache(bid))
		|| bptr->brdname[0] == '\0' )
	    continue;
	if (bptr->gid == gid) {
	    if (currbptr == parent)
		currbptr->firstchild[type] = bid;
	    else {
		currbptr->next[type] = bid;
		currbptr->parent = gid;
	    }
	    childcount++;
	    currbptr = bptr;
	}
    }
    parent->childcount = childcount;
    if (currbptr == parent) // no child
	currbptr->firstchild[type] = -1;
    else // the last child
	currbptr->next[type] = -1;
}

void
setbottomtotal(int bid)
{
    boardheader_t  *bh = getbcache(bid);
    char            fname[PATHLEN];
    int             n;

    assert(0<=bid-1 && bid-1<MAX_BOARD);
    if(!bh->brdname[0]) return;
    setbfile(fname, bh->brdname, FN_DIR ".bottom");
    n = get_num_records(fname, sizeof(fileheader_t));
    if(n>5)
      {
#ifdef DEBUG_BOTTOM
        log_file("fix_bottom", LOG_CREAT | LOG_VF, "%s n:%d\n", fname, n);
#endif
        unlink(fname);
        SHM->n_bottom[bid-1]=0;
      }
    else
        SHM->n_bottom[bid-1]=n;
}

void
setbtotal(int bid)
{
    boardheader_t  *bh = getbcache(bid);
    struct stat     st;
    char            genbuf[PATHLEN];
    int             num, fd;

    assert(0<=bid-1 && bid-1<MAX_BOARD);
    setbfile(genbuf, bh->brdname, FN_DIR);
    if ((fd = open(genbuf, O_RDONLY)) < 0)
	return;			/* .DIR���F */
    fstat(fd, &st);
    num = st.st_size / sizeof(fileheader_t);
    assert(0<=bid-1 && bid-1<MAX_BOARD);
    SHM->total[bid - 1] = num;

    if (num > 0) {
	lseek(fd, (off_t) (num - 1) * sizeof(fileheader_t), SEEK_SET);
	if (read(fd, genbuf, FNLEN) >= 0) {
#ifdef FN_SAFEDEL_PREFIX_LEN
            if (strncmp(genbuf, FN_SAFEDEL, FN_SAFEDEL_PREFIX_LEN) == 0)
                SHM->lastposttime[bid - 1] = 0;
            else
#endif
	    SHM->lastposttime[bid - 1] = (time4_t) atoi(&genbuf[2]);
	}
    } else
	SHM->lastposttime[bid - 1] = 0;
    close(fd);
}

void
touchbpostnum(int bid, int delta)
{
    int            *total = &SHM->total[bid - 1];
    assert(0<=bid-1 && bid-1<MAX_BOARD);
    if (*total)
	*total += delta;
}

int
getbnum(const char *bname)
{
    register int    i = 0, j, start = 0, end = SHM->Bnumber - 1;
    int *blist = SHM->bsorted[0];
    if(SHM->Bbusystate)
	sleep(1);
    for (i = ((start + end) / 2);; i = (start + end) / 2) {
	if (!(j = strcasecmp(bname, bcache[blist[i]].brdname)))
	    return (int)(blist[i] + 1);
	if (end == start) {
	    break;
	} else if (i == start) {
	    i = end;
	    start = end;
	} else if (j > 0)
	    start = i;
	else
	    end = i;
    }
    return 0;
}

int
parseBMlist(const char *input, int uids[MAX_BMs]) {
    int i, uid;
    char *ptr, *strtok_pos;
    char s[(IDLEN + 1) * MAX_BMs];

    strlcpy(s, input, sizeof(s));
    // reset BM list
    for (i = 0; i < MAX_BMs; i++)
        uids[i] = -1;

    for (i = 0 ; s[i] != 0 ; ++i)
	if (!isalpha((int)s[i]) && !isdigit((int)s[i]))
            s[i] = ' ';
    for (ptr = strtok_r(s, " ", &strtok_pos), i = 0;
	 i < MAX_BMs && ptr != NULL;
         ptr = strtok_r(NULL, " ", &strtok_pos))
        if((uid = searchuser(ptr, NULL)) != 0)
            uids[i++] = uid;
    return i;
}


void
buildBMcache(int bid) /* bid starts from 1 */
{
    assert(0<=bid-1 && bid-1<MAX_BOARD);
    parseBMlist(getbcache(bid)->BM, SHM->BMcache[bid - 1]);
}

/*
 * section - PTT cache (adbanner cache?)
 * �ʺA�ݪO�P�䥦
 */
int 
filter_aggressive(const char*s)
{
    (void)s;
    if (
	/*
	strstr(s, "���B������A����ĳ�ʦr�y") != NULL ||
	*/
	0
	)
	return 1;
    return 0;
}

int 
filter_dirtywords(const char*s)
{
    if (
	strstr(s, "�F�A�Q") != NULL ||
	0)
	return 1;
    return 0;
}

#define AGGRESSIVE_FN ".aggressive"
static char drop_aggressive = 0;

void 
load_aggressive_state()
{
    if (dashf(AGGRESSIVE_FN))
	drop_aggressive = 1;
    else
	drop_aggressive = 0;
}

void 
set_aggressive_state(int s)
{
    FILE *fp = NULL;
    if (s)
    {
	fp = fopen(AGGRESSIVE_FN, "wb");
	fclose(fp);
    } else {
	remove(AGGRESSIVE_FN);
    }
}

/* cache for �ʺA�ݪO */
void
reload_pttcache(void)
{
    if (SHM->Pbusystate)
	safe_sleep(1);
    else {			/* jochang: temporary workaround */
	fileheader_t    item, subitem;
	char            pbuf[256], buf[256];
	FILE           *fp, *fp1, *fp2;
	int             id, aggid, rawid;

	SHM->Pbusystate = 1;
	SHM->last_film = 0;
	bzero(SHM->notes, sizeof(SHM->notes));
	setapath(pbuf, BN_NOTE);
	setadir(buf, pbuf);

	load_aggressive_state();
	id = aggid = rawid = 0; // effective count, aggressive count, total (raw) count

	if ((fp = fopen(buf, "r"))) {

	    // .DIR loop
	    while (fread(&item, sizeof(item), 1, fp)) {

		int chkagg = 0; // should we check aggressive?
		int is_ordersong_dir = 0;

		if (item.title[3] != '<' || item.title[8] != '>')
		    continue;

#define ORDERSONG_FOLDERNAME	"<�I�q>"
		if (strncmp(item.title+3, ORDERSONG_FOLDERNAME, strlen(ORDERSONG_FOLDERNAME)) == 0)
		    is_ordersong_dir = 1;

#ifdef BN_NOTE_AGGCHKDIR
		// TODO aggressive: only count '<�I�q>' section
		if (strncmp(item.title+3, BN_NOTE_AGGCHKDIR, strlen(BN_NOTE_AGGCHKDIR)) == 0)
		    chkagg = 1;
#endif
		snprintf(buf, sizeof(buf), "%s/%s/" FN_DIR,
			pbuf, item.filename);

		if (!(fp1 = fopen(buf, "r")))
		    continue;

		// file loop
		while (fread(&subitem, sizeof(subitem), 1, fp1)) {

		    snprintf(buf, sizeof(buf),
			    "%s/%s/%s", pbuf, item.filename,
			    subitem.filename);

		    if (!(fp2 = fopen(buf, "r")))
			continue;

		    fread(SHM->notes[id], sizeof(char), sizeof(SHM->notes[0]), fp2);
		    SHM->notes[id][sizeof(SHM->notes[0]) - 1] = 0;
		    rawid ++;

		    // filtering
		    if (filter_dirtywords(SHM->notes[id]))
		    {
			memset(SHM->notes[id], 0, sizeof(SHM->notes[0]));
			rawid --;
		    }
		    else if (chkagg && filter_aggressive(SHM->notes[id]))
		    {
			aggid++;
			// handle aggressive notes by last detemined state
			if (drop_aggressive)
			    memset(SHM->notes[id], 0, sizeof(SHM->notes[0]));
			else
			    id++;
			// Debug purpose
			// fprintf(stderr, "found aggressive: %s\r\n", buf);
		    } 
		    else 
		    {
			id++;
		    }

		    fclose(fp2);
		    if (id >= MAX_ADBANNER)
			break;

		} // end of file loop
		fclose(fp1);

		if (is_ordersong_dir)
		    SHM->last_usong = id - 1;

		if (id >= MAX_ADBANNER)
		    break;

	    } // end of .DIR loop
	    fclose(fp);

	    // decide next aggressive state
	    if (rawid && aggid*3 >= rawid) // if aggressive exceed 1/3
		set_aggressive_state(1);
	    else
		set_aggressive_state(0);

	    // fprintf(stderr, "id(%d)/agg(%d)/raw(%d)\r\n",
	    //	    id, aggid, rawid);
	}
	SHM->last_film = id - 1;

	/* ���Ҧ���Ƨ�s��A�]�w uptime */

	SHM->Puptime = SHM->Ptouchtime;
	// log_usies("CACHE", "reload pttcache");
	fprintf(stderr, "cache: reload pttcache\r\n");
	SHM->Pbusystate = 0;
    }
}

void
resolve_garbage(void)
{
    int             count = 0;

    while (SHM->Puptime < SHM->Ptouchtime) {	/* ����while�� */
	reload_pttcache();
	if (count++ > 10 && SHM->Pbusystate) {
	    /*
	     * Ptt: �o��|�����D  load�W�L10 ��|�Ҧ��iloop��process tate = 0
	     * �o�˷|�Ҧ�prcosee���|�bload �ʺA�ݪO �|�y��load�j�W
	     * ���S���γo��function���� �U�@load passwd�ɪ�process���F
	     * �S�S���H��L �Ѷ}  �P�˪����D�o�ͦbreload passwd
	     */
	    SHM->Pbusystate = 0;
	    // log_usies("CACHE", "refork Ptt dead lock");
	    fprintf(stderr, "cache: refork Ptt dead lock\r\n");
	}
    }
}

/*
 * section - from host (deprecated by fromd / logind?)
 * cache for from host �P�̦h�W�u�H�� 
 */
void
reload_fcache(void)
{
    if (SHM->Fbusystate)
	safe_sleep(1);
    else {
	SHM->Fbusystate = 1;
	SHM->max_user = 0;

	/* ���Ҧ���Ƨ�s��A�]�w uptime */
	SHM->Fuptime = SHM->Ftouchtime;
	// log_usies("CACHE", "reload fcache");
	fprintf(stderr, "cache: reload from cache\r\n");
	SHM->Fbusystate = 0;
    }
}

void
resolve_fcache(void)
{
    while (SHM->Fuptime < SHM->Ftouchtime)
	reload_fcache();
}

/*
 * section - hbfl (hidden board friend list)
 */
void
hbflreload(int bid)
{
    int             hbfl[MAX_FRIEND + 1], i, num, uid;
    char            buf[128];
    FILE           *fp;

    assert(0<=bid-1 && bid-1<MAX_BOARD);
    memset(hbfl, 0, sizeof(hbfl));
    setbfile(buf, bcache[bid - 1].brdname, fn_visable);
    if ((fp = fopen(buf, "r")) != NULL) {
	for (num = 1; num <= MAX_FRIEND; ++num) {
	    if (fgets(buf, sizeof(buf), fp) == NULL)
		break;
	    for (i = 0; buf[i] != 0; ++i)
		if (buf[i] == ' ') {
		    buf[i] = 0;
		    break;
		}
	    if (strcasecmp(STR_GUEST, buf) == 0 ||
		(uid = searchuser(buf, NULL)) == 0) {
		--num;
		continue;
	    }
	    hbfl[num] = uid;
	}
	fclose(fp);
    }
    hbfl[0] = COMMON_TIME;
    memcpy(SHM->hbfl[bid-1], hbfl, sizeof(hbfl));
}

/* �O�_�q�L�O�ʹ���. �p�G�b�O�ͦW�椤���ܶǦ^ 1, �_�h�� 0 */
int
is_hidden_board_friend(int bid, int uid)
{
    int             i;

    assert(0<=bid-1 && bid-1<MAX_BOARD);
    if (SHM->hbfl[bid-1][0] < login_start_time - HBFLexpire)
	hbflreload(bid);
    for (i = 1; SHM->hbfl[bid-1][i] != 0 && i <= MAX_FRIEND; ++i) {
	if (SHM->hbfl[bid-1][i] == uid)
	    return 1;
    }
    return 0;
}

/*
 * section - cooldown
 */
#ifdef USE_COOLDOWN

void add_cooldowntime(int uid, int min)
{
    // Ptt: I will use the number below 15 seconds.
    time4_t base= now > SHM->cooldowntime[uid - 1]? 
                    now : SHM->cooldowntime[uid - 1];
    base += min*60;
    base &= 0xFFFFFFF0;

    SHM->cooldowntime[uid - 1] = base;
}
void add_posttimes(int uid, int times)
{
  if((SHM->cooldowntime[uid - 1] & 0xF) + times <0xF)
       SHM->cooldowntime[uid - 1] += times;
  else
       SHM->cooldowntime[uid - 1] |= 0xF;
}

#endif
