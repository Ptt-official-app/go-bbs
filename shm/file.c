#include <stdio.h>
#include <stdlib.h> // random
#include <sys/file.h> // flock
#include <unistd.h>
#include <fcntl.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <limits.h>
#include <strings.h>
#include <dirent.h>
#include <string.h>
#include <sys/wait.h>

#include "cmsys.h"


/* ----------------------------------------------------- */
/* �ɮ��ˬd��ơG�ɮסB�ؿ��B�ݩ�                        */
/* ----------------------------------------------------- */

/**
 * �Ǧ^ fname ���ɮפj�p
 * @param fname
 */
off_t
dashs(const char *fname)
{
    struct stat     st;

    if (!stat(fname, &st))
	return st.st_size;
    else
	return -1;
}

/**
 * �Ǧ^ fname �� mtime
 * @param fname
 */
time4_t
dasht(const char *fname)
{
    struct stat     st;

    if (!stat(fname, &st))
	return st.st_mtime;
    else
	return -1;
}

/**
 * �Ǧ^ fname �� ctime
 * @param fname
 */
time4_t
dashc(const char *fname)
{
    struct stat     st;

    if (!stat(fname, &st))
	return st.st_ctime;
    else
	return -1;
}

/**
 * �Ǧ^ fname �O�_�� symbolic link
 * @param fname
 */
int
dashl(const char *fname)
{
    struct stat     st;

    return (lstat(fname, &st) == 0 && S_ISLNK(st.st_mode));
}

/**
 * �Ǧ^ fname �O�_���@�몺�ɮ�
 * @param fname
 */
int
dashf(const char *fname)
{
    struct stat     st;

    return (stat(fname, &st) == 0 && S_ISREG(st.st_mode));
}

/**
 * �Ǧ^ fname �O�_���ؿ�
 * @param fname
 */
int
dashd(const char *fname)
{
    struct stat     st;

    return (stat(fname, &st) == 0 && S_ISDIR(st.st_mode));
}

/* ----------------------------------------------------- */
/* �ɮ׾ާ@��ơG�ƻs�B�h���B���[                        */
/* ----------------------------------------------------- */

#define BUFFER_SIZE	8192
int copy_file_to_file(const char *src, const char *dst)
{
    char buf[BUFFER_SIZE];
    int fdr, fdw, len;

    if ((fdr = open(src, O_RDONLY)) < 0)
	return -1;

    if ((fdw = OpenCreate(dst, O_WRONLY | O_TRUNC)) < 0) {
	close(fdr);
	return -1;
    }

    while (1) {
	len = read(fdr, buf, sizeof(buf));
	if (len <= 0)
	    break;
	write(fdw, buf, len);
	if (len < BUFFER_SIZE)
	    break;
    }

    close(fdr);
    close(fdw);
    return 0;
}
#undef BUFFER_SIZE

int copy_file_to_dir(const char *src, const char *dst)
{
    char buf[PATH_MAX];
    char *slash;
    if ((slash = rindex(src, '/')) == NULL)
	snprintf(buf, sizeof(buf), "%s/%s", dst, src);
    else
	snprintf(buf, sizeof(buf), "%s/%s", dst, slash);
    return copy_file_to_file(src, buf);
}

int copy_dir_to_dir(const char *src, const char *dst)
{
    DIR *dir;
    struct dirent *entry;
    struct stat st;
    char buf[PATH_MAX], buf2[PATH_MAX];

    if (stat(dst, &st) < 0)
	if (Mkdir(dst) < 0)
	    return -1;

    if ((dir = opendir(src)) == NULL)
	return -1;

    while ((entry = readdir(dir)) != NULL) {
	if (strcmp(entry->d_name, ".") == 0 ||
	    strcmp(entry->d_name, "..") == 0)
	    continue;
	snprintf(buf, sizeof(buf), "%s/%s", src, entry->d_name);
	snprintf(buf2, sizeof(buf2), "%s/%s", dst, entry->d_name);
	if (stat(buf, &st) < 0)
	    continue;
	if (S_ISDIR(st.st_mode))
	    Mkdir(buf2);
	copy_file(buf, buf2);
    }

    closedir(dir);
    return 0;
}

/**
 * copy src to dst (recursively)
 * @param src and dst are file or dir
 * @return -1 if failed
 */
int copy_file(const char *src, const char *dst)
{
    struct stat st;

    if (stat(dst, &st) == 0 && S_ISDIR(st.st_mode)) {
	if (stat(src, &st) < 0)
	    return -1;
	
    	if (S_ISDIR(st.st_mode))
	    return copy_dir_to_dir(src, dst);
	else if (S_ISREG(st.st_mode))
	    return copy_file_to_dir(src, dst);
	return -1;
    }
    else if (stat(src, &st) == 0 && S_ISDIR(st.st_mode))
	return copy_dir_to_dir(src, dst);
    return copy_file_to_file(src, dst);
}

#include <signal.h>
#include <errno.h>
int
Rename(const char *src, const char *dst)
{
    pid_t pid;
    sig_t s;
    int ret = -1;

    if (rename(src, dst) == 0)
	return 0;

    if (errno != EXDEV)
        return -1;

    // prevent malicious shell escapes
    if (strchr(src, ';') || strchr(dst, ';'))
	return -1;

    // because we need the return value, override the signal handler
    s = signal(SIGCHLD, SIG_DFL);
    pid = fork();

    if (pid == 0)
	execl("/bin/mv", "mv", "-f", src, dst, (char *)NULL);
    else if (pid > 0)
    {
	waitpid(pid, &ret, 0);
	ret = WEXITSTATUS(ret) == 0 ? 0 : -1;
    }

    signal(SIGCHLD, s);
    return ret;
}

int
Copy(const char *src, const char *dst)
{
    int fi, fo, bytes;
    char buf[8192];
    fi=open(src, O_RDONLY);
    if(fi<0) return -1;
    fo=OpenCreate(dst, O_WRONLY | O_TRUNC);
    if(fo<0) {close(fi); return -1;}
    while((bytes=read(fi, buf, sizeof(buf)))>0)
         write(fo, buf, bytes);
    close(fo);
    close(fi);
    return 0;  
}

int
CopyN(const char *src, const char *dst, int n)
{
    int fi, fo, bytes;
    char buf[8192];

    fi=open(src, O_RDONLY);
    if(fi<0) return -1;

    fo=OpenCreate(dst, O_WRONLY | O_TRUNC);
    if(fo<0) {close(fi); return -1;}

    while(n > 0 && (bytes=read(fi, buf, sizeof(buf)))>0)
    {
	 n -= bytes;
	 if (n < 0)
	     bytes += n;
         write(fo, buf, bytes);
    }
    close(fo);
    close(fi);
    return 0;  
}

/* append data from tail of src (starting point=off) to dst */
int
AppendTail(const char *src, const char *dst, int off)
{
    int fi, fo, bytes;
    char buf[8192];

    fi=open(src, O_RDONLY);
    if(fi<0) return -1;

    fo=OpenCreate(dst, O_WRONLY | O_APPEND);
    if(fo<0) {close(fi); return -1;}
    // flock(dst, LOCK_SH);
    
    if(off > 0)
	lseek(fi, (off_t)off, SEEK_SET);

    while((bytes=read(fi, buf, sizeof(buf)))>0)
    {
         write(fo, buf, bytes);
    }
    // flock(dst, LOCK_UN);
    close(fo);
    close(fi);
    return 0;  
}

/**
 * @param src	file
 * @param dst	file
 * @return 0	if success
 */
int
Link(const char *src, const char *dst)
{
    if (symlink(src, dst) == 0)
       return 0;

    return Copy(src, dst);
}

/**
 * @param src	file
 * @param dst	file
 * @return 0	if success
 */
int
HardLink(const char *src, const char *dst)
{
    if (link(src, dst) == 0)
       return 0;

    return Copy(src, dst);
}

/**
 * @param path	directory name
 * @return 0	if success
 */
int
Mkdir(const char *path)
{
    return mkdir(path, DEFAULT_FOLDER_CREATE_PERM);
}

/**
 * @param path	directory name
 * @param flags optional parameters
 * @return 0	if success
 */
int
OpenCreate(const char *path, int flags)
{
    return open(path, flags | O_CREAT, DEFAULT_FILE_CREATE_PERM);
}

/* ----------------------------------------------------- */
/* �ɮפ��e�B�z��ơG�H�u��v�����                      */
/* ----------------------------------------------------- */

#define LINEBUFSZ (PATH_MAX)
#define STR_SPACE " \t\n\r"


/**
 * �Ǧ^ file �ɪ����
 * @param file
 */
int file_count_line(const char *file)
{
    FILE           *fp;
    int             count = 0;
    char            buf[LINEBUFSZ];

    if ((fp = fopen(file, "r"))) {
	while (fgets(buf, sizeof(buf), fp)) {
	    if (strchr(buf, '\n') == NULL)
		continue;
	    count++;
	}
	fclose(fp);
    }
    return count;
}

/**
 * �N string append ���ɮ� file ��� (���[����)
 * @param file �n�Q append ����
 * @param string
 * @return ���\�Ǧ^ 0�A���ѶǦ^ -1�C
 */
int file_append_line(const char *file, const char *string)
{
    FILE *fp;
    if ((fp = fopen(file, "a")) == NULL)
	return -1;
    flock(fileno(fp), LOCK_EX);
    fputs(string, fp);
    flock(fileno(fp), LOCK_UN);
    fclose(fp);
    return 0;
}

/**
 * �N "$key\n" append ���ɮ� file ���
 * @param file �n�Q append ����
 * @param key �S�����檺�r��
 * @return ���\�Ǧ^ 0�A���ѶǦ^ -1�C
 */
int file_append_record(const char *file, const char *key)
{
    FILE *fp;
    if (!key || !*key) return -1;
    if ((fp = fopen(file, "a")) == NULL)
	return -1;
    flock(fileno(fp), LOCK_EX);
    fputs(key, fp);
    fputs("\n", fp);
    flock(fileno(fp), LOCK_UN);
    fclose(fp);
    return 0;
}

/**
 * �Ǧ^�ɮ� file �� key �Ҧb���
 */
int file_find_record(const char *file, const char *key)
{
    FILE           *fp;
    char            buf[LINEBUFSZ], *ptr;
    int i = 0;

    if ((fp = fopen(file, "r")) == NULL)
	return 0;

    while (fgets(buf, LINEBUFSZ, fp)) {
	char *strtok_pos;
	i++;
	if ((ptr = strtok_r(buf, STR_SPACE, &strtok_pos)) && !strcasecmp(ptr, key)) {
	    fclose(fp);
	    return i;
	}
    }
    fclose(fp);
    return 0;
}

/**
 * �Ǧ^�ɮ� file ���O�_�� key
 */
int file_exist_record(const char *file, const char *key)
{
    return file_find_record(file, key) > 0 ? 1 : 0;
}

/**
 * �R���ɮ� file ���H string �}�Y����
 * @param file �n�B�z���ɮ�
 * @param string �M�䪺 key name
 * @param case_sensitive �O�_�n�B�z�j�p�g
 * @return ���\�Ǧ^ 0�A���ѶǦ^ -1�C
 */
int
file_delete_record(const char *file, const char *string, int case_sensitive)
{
    // TODO nfp �� tmpfile() ����n�H ���L Rename �|�ܺC...
    FILE *fp = NULL, *nfp = NULL;
    char fnew[PATH_MAX];
    char buf[LINEBUFSZ + 1];
    int ret = -1, i = 0;
    const size_t toklen = strlen(string);

    if (!toklen)
	return 0;

    do {
	snprintf(fnew, sizeof(fnew), "%s.%3.3X", file, (unsigned int)(random() & 0xFFF));
	if (access(fnew, 0) != 0)
	    break;
    } while (i++ < 10); // max tries = 10

    if (access(fnew, 0) == 0) return -1;    // cannot create temp file.

    i = 0;
    if ((fp = fopen(file, "r")) && (nfp = fopen(fnew, "w"))) {
	while (fgets(buf, sizeof(buf), fp))
	{
	    size_t klen = strcspn(buf, STR_SPACE);
	    if (toklen == klen)
	    {
		if (((case_sensitive && strncmp(buf, string, toklen) == 0) ||
		    (!case_sensitive && strncasecmp(buf, string, toklen) == 0)))
		{
		    // found line. skip it.
		    i++;
		    continue;
		}
	    }
	    // other wise, keep the line.
	    fputs(buf, nfp);
            // Fix broken records (ex, deprecated "distinct" files in format
            // "%s\0#%s\n".
            if (*buf && buf[strlen(buf) - 1] != '\n')
                fputc('\n', nfp);
	}
	fclose(nfp); nfp = NULL;
	if (i > 0)
	{
	    if(Rename(fnew, file) < 0)
		ret = -1;
	    else
		ret = 0;
	} else {
	    unlink(fnew);
	    ret = 0;
	}
    }
    if(fp)
	fclose(fp);
    if(nfp)
	fclose(nfp);
    return ret;
}

/**
 * ��C�@�� record �� func �o��ơC
 * @param file
 * @param func �B�z�C�� record �� handler�A���@ function pointer�C
 *        �Ĥ@�ӰѼƬO�ɮפ����@��A�ĤG�ӰѼƬ� info�C
 * @param info �@���B�~���ѼơC
 */
int file_foreach_entry(const char *file, int (*func)(char *, int), int info)
{
    char line[80];
    FILE *fp;

    if ((fp = fopen(file, "r")) == NULL)
	return -1;

    while (fgets(line, sizeof(line), fp)) {
	(*func)(line, info);
    }

    fclose(fp);
    return 0;
}
