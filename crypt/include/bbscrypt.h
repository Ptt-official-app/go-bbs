// from https://github.com/ptt/pttbbs/blob/master/common/sys/crypt.c
// pttbbs commit: 6bdd36898bde207683a441cdffe2981e95de5b20

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
