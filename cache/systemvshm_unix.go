// +build !windows

package cache

import (
	syscall "golang.org/x/sys/unix"
	"unsafe"
)

func shmget(key int, size int, flag int) (int, error) {
	r1, _, err := syscall.Syscall(syscall.SYS_SHMGET, uintptr(key), uintptr(size), uintptr(flag))
	if err != 0 {
		return 0, err
	}
	return int(r1), nil
}

func shmat(shmid int, shmaddr uintptr, shmflg int) (uintptr, error) {
	r1, _, err := syscall.Syscall(syscall.SYS_SHMAT, uintptr(shmid), uintptr(unsafe.Pointer(shmaddr)), uintptr(shmflg))
	if err != 0 {
		return 0, err
	}
	// XXX

	return r1, nil
}

func shmdt(shmaddr uintptr) (int, error) {
	r1, _, err := syscall.Syscall(syscall.SYS_SHMDT, shmaddr, uintptr(0), uintptr(0))
	if err != 0 {
		return int(r1), err
	}
	// XXX

	return int(r1), nil
}

func shmctl(shmid int, cmd int, buf *ShmidDs) (int, error) {
	r1, _, err := syscall.Syscall(syscall.SYS_SHMCTL, uintptr(shmid), uintptr(cmd), uintptr(unsafe.Pointer(buf)))
	if err != 0 {
		return int(r1), err
	}
	// XXX
	return int(r1), nil
}
