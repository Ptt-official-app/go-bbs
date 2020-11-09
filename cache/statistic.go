package cache

import (
	"unsafe"

	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

func StatInc(stats uintptr) error {
	err := validateStats(stats)
	if err != nil {
		return err
	}

	Shm.IncUint32(unsafe.Offsetof(Shm.Statistic) + types.UINT32_SZ*stats)

	return nil
}

func CleanStat() {
	in := [ptttype.STAT_MAX]uint32{}
	Shm.WriteAt(
		unsafe.Offsetof(Shm.Statistic),
		unsafe.Sizeof(Shm.Statistic),
		unsafe.Pointer(&in),
	)
}

func ReadStat(stats uintptr) (uint32, error) {
	err := validateStats(stats)
	if err != nil {
		return 0, err
	}

	out := uint32(0)
	Shm.ReadAt(
		unsafe.Offsetof(Shm.Statistic)+types.UINT32_SZ*stats,
		types.UINT32_SZ,
		unsafe.Pointer(&out),
	)

	return out, nil
}

func validateStats(stats uintptr) error {
	if Shm == nil {
		return ErrShmNotInit
	}

	if Shm.Version != SHM_VERSION {
		return ErrShmVersion
	}
	if stats < 0 || stats >= ptttype.STAT_MAX {
		return ErrStats
	}

	return nil
}
