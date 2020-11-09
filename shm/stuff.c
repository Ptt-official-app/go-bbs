#include "cmsys.h"
#include "var.h"

// synchronize 'now'
void syncnow(void)
{
	now = time(0);
}
