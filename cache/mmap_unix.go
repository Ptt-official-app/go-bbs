// +build unix darwin

package cache

import (
	"fmt"

	syscall "golang.org/x/sys/unix"
)

func openMmap(fd int, size int) ([]byte, error) {
	b, err := syscall.Mmap(fd, 0, size, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		return nil, fmt.Errorf("mmap error: %v", err)
	}
	return b, err

}

func closeMmap(buf []byte) error {
	return syscall.Munmap(buf)
}
