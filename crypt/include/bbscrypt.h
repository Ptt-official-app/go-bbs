// from https://github.com/ptt/pttbbs/blob/master/common/sys/crypt.c

#ifndef _BBSCRYPT_H_
#define _BBSCRYPT_H_

#include <stdlib.h>
#include <string.h>

#ifdef PERL5
char *des_crypt(char *buf, char *salt);
#else
char *fcrypt(char *buf, char *salt);
#endif

#endif
