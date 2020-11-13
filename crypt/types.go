package crypt

import "unsafe"

type desCBlock [8]uint8

const DES_KEY_SZ = unsafe.Sizeof(desCBlock{})

//desKeySchedule is used for uint32-computation.
//not-sure how to have *[32]uint32 pointed to [16][8]uint8
//set desKeySchedule directly to [32]uint32
type desKeySchedule [32]uint32 // 16 * 8 = 128 ([16]desCBlock)
