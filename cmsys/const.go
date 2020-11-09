package cmsys

const (
	//https://github.com/ptt/pttbbs/blob/master/include/fnv_hash.h
	//commit: 6bdd36898bde207683a441cdffe2981e95de5b20
	FNV1_32_INIT uint32 = 33554467
	FNV1_64_INIT uint64 = 0xcbf29ce484222325

	FNV_32_PRIME uint32 = 0x01000193
	FNV_64_PRIME uint64 = 0x100000001b3
)
