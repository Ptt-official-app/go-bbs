package types

import (
	"unsafe"
)

type Time4 int32

type Pid_t uint32

type InAddr_t uint32

type Key_t int

type Size_t uint32

const INT32_SZ = unsafe.Sizeof(int32(0))

const UINT32_SZ = unsafe.Sizeof(uint32(0))

const TIME4_SZ = INT32_SZ
