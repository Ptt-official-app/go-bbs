package cache

import "errors"

var (
	ErrInvalidOp      = errors.New("invalid op")
	ErrShmNotInit     = errors.New("shm not init")
	ErrShmAlreadyInit = errors.New("shm already init")
	ErrShmVersion     = errors.New("invalid shm version")

	ErrStats = errors.New("invalid stats")
)
