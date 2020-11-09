package cache

var (
	//////////
	// !!! We should have only 1 Shm.
	Shm *SHM = nil
)

const (
	// from https://github.com/ptt/pttbbs/blob/master/include/pttstruct.h
	// commit: 6bdd36898bde207683a441cdffe2981e95de5b20
	SHM_VERSION = 4842
)
