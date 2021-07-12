// +build windows

package cache

import (
	"fmt"
)

func openMmap(fd int, size int) ([]byte, error) {
	return nil, fmt.Errorf("TODO: implement open mmap")
}

func closeMmap(buf []byte) error {
	return fmt.Errorf("TODO: implement close mmap")
}
