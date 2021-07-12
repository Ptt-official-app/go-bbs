package crypt

type desCBlock [8]uint8

// desKeySchedule is used for uint32-computation.
// not-sure how to have *[32]uint32 pointed to [16][8]uint8
// set desKeySchedule directly to [32]uint32
type desKeySchedule [32]uint32 // 16 * 8 = 128 ([16]desCBlock)
