package bbs

const (
	//////////
	//pttstruch.h
	//////////
	PTT_IDLEN   = 12 /* Length of bid/uid */
	PTT_IPV4LEN = 15 /* a.b.c.d form */

	PTT_PASS_INPUT_LEN = 8 /* Length of valid input password length.
	   For DES, set to 8. */
	PTT_PASSLEN = 14 /* Length of encrypted passwd field */
	PTT_REGLEN  = 38 /* Length of registration data */

	PTT_REALNAMESZ = 20 /* Size of real-name field */
	PTT_NICKNAMESZ = 24 /* SIze of nick-name field */
	PTT_EMAILSZ    = 50 /* Size of email field */
	PTT_ADDRESSSZ  = 50 /* Size of address field */
	PTT_CAREERSZ   = 40 /* Size of career field */
	PTT_PHONESZ    = 20 /* Size of phone field */

	PTT_PASSWD_VERSION = 4194 /* passwd version */

	PTT_TTLEN = 64 /* Length of title */
	PTT_FNLEN = 28 /* Length of filename */
)

const (
	//////////
	//pttstruch.h: 292
	//////////
	PTT_FILE_LOCAL     = 0x01 /* local saved,  non-mail */
	PTT_FILE_READ      = 0x01 /* already read, mail only */
	PTT_FILE_MARKED    = 0x02 /* non-mail + mail */
	PTT_FILE_DIGEST    = 0x04 /* digest,       non-mail */
	PTT_FILE_REPLIED   = 0x04 /* replied,      mail only */
	PTT_FILE_BOTTOM    = 0x08 /* push_bottom,  non-mail */
	PTT_FILE_MULTI     = 0x08 /* multi send,   mail only */
	PTT_FILE_SOLVED    = 0x10 /* problem solved, sysop/BM non-mail only */
	PTT_FILE_HIDE      = 0x20 /* hide,	in announce */
	PTT_FILE_BID       = 0x20 /* bid,		in non-announce */
	PTT_FILE_BM        = 0x40 /* BM only,	in announce */
	PTT_FILE_VOTE      = 0x40 /* for vote,	in non-announce */
	PTT_FILE_ANONYMOUS = 0x80 /* anonymous file */
)
