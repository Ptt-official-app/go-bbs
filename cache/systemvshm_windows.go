// +build windows

package cache

import (
	"fmt"
)

type sizeT uint64

type ShmidDs struct {
	ShmSegsz sizeT
}

func shmget(key int, size int, flag int) (int, error) {
	return 0, fmt.Errorf("windows do not implement shmget")
}

func shmat(shmid int, shmaddr uintptr, shmflg int) (uintptr, error) {
	return 0, fmt.Errorf("windows do not implement shmat")
}

func shmdt(shmaddr uintptr) (int, error) {
	return 0, fmt.Errorf("windows do not implement shmdt")
}

func shmctl(shmid int, cmd int, buf *ShmidDs) (int, error) {
	return 0, fmt.Errorf("windows do not implement shmctl")
}
