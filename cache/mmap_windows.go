// +build windows

package cache

import (
	"fmt"
)

func openMmap(fd int, size int) ([]byte, error) {
	return nil, fmt.Errorf("windows do not implement shmget")

}

func closeMmap(buf []byte) error {
	return fmt.Errorf("windows do not implement shmget")
}
