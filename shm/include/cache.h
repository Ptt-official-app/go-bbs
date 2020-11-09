#ifndef __BBS_CACHE_H__
#define __BBS_CACHE_H__

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
#include <sys/mman.h>
#include <sys/sem.h>
#include <unistd.h>
#include <errno.h>
#include <fcntl.h>
#include "cmsys.h"
#include "cmbbs.h"
#include "common.h"
#include "var.h"

#include "modes.h" // for DEBUGSLEEPING

#endif // __BBS_CACHE_H__