
#include <sys/stat.h>
#include <sys/mman.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#include <assert.h>
#include <stdlib.h>
#include "cmsys.h"

#define BUFSIZE 512

/* Functions for fixed size record operations */

int
get_num_records(const char *fpath, size_t size)
{
    struct stat st;

    if (stat(fpath, &st) == -1)
	return 0;

    return st.st_size / size;
}
