#include "var.h"

time4_t         login_start_time, last_login_time;

char * const fn_board = FN_BOARD;
const char * const fn_visable = FN_VISABLE;

time4_t         now;

SHM_t          *SHM;
boardheader_t  *bcache;
