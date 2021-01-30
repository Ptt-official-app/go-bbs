package cache

// +build unix

import (
	"fmt"

	syscall "golang.org/x/sys/unix"
)

func openMmap(fd int, size int) ([]byte, error) {
	b, err := syscall.Mmap(fd, 0, size, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		fmt.Println("mmap error", err)
		return nil, fmt.Errorf("mmap error:", err)
	}
	return b, err

}

func closeMmap(buf []byte) error {
	return syscall.Munmap(buf)
}
