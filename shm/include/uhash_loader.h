#ifndef __BBS_UHASH_LOADER_H__
#define __BBS_UHASH_LOADER_H__

#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <signal.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/ipc.h>
#include <sys/mman.h>
#include <sys/shm.h>
#include <sys/sem.h>
#include <unistd.h>
#include <errno.h>
#include <fcntl.h>

#include "config.h"
#include "cmbbs.h"
#include "pttstruct.h"
#include "common.h"
#include "var.h"

void userec_add_to_uhash(int n, userec_t *id, int onfly);
int fill_uhash(int onfly);
int load_uhash(void);

#endif