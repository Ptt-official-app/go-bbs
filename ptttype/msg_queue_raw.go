package ptttype

import (
	"unsafe"

	"github.com/PichuChen/go-bbs/types"
)

type MsgQueueRaw struct {
	//Require updating SHM_VERSION if MSG_QUEUE_RAW_SZ is changed.
	Pid        types.Pid_t
	UserID     [IDLEN + 1]byte
	LastCallIn [76]byte
	MsgMode    int32
}

//Require updating SHM_VERSION if MSG_QUEUE_RAW_SZ is changed.
const MSG_QUEUE_RAW_SZ = unsafe.Sizeof(MsgQueueRaw{})
