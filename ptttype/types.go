package ptttype

import "unsafe"

type UserID [IDLEN + 1]byte

const USER_ID_SZ = unsafe.Sizeof(UserID{})
