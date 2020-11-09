/* $Id$ */
#ifndef INCLUDE_PROTO_H
#define INCLUDE_PROTO_H

#ifdef __GNUC__
#define GCC_CHECK_FORMAT(a,b) __attribute__ ((format (printf, a, b)))
#define GCC_NORETURN          __attribute__ ((__noreturn__))
#define GCC_UNUSED            __attribute__ ((__unused__))
#else
#define GCC_CHECK_FORMAT(a,b)
#define GCC_NORETURN
#define GCC_UNUSED
#endif

/* stuff */
void syncnow(void);

#endif