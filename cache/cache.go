package cache

import (
	"fmt"
	"strconv"
	"strings"
)

type Cache interface {
	Bytes() []byte
	Close() error
}

// NewCache returns Cache (SHM) by connectionString, connectionString indicate the shm location
// with uri format  eg. shmkey:1228 or file:/tmp/ramdisk/bbs.shm
func NewCache(connectionString string) (Cache, error) {
	s := strings.Split(connectionString, ":")
	if len(s) == 1 {
		// default is mmap
		return OpenMmap(s[0])
	}
	scheme := s[0]
	switch scheme {
	case "shmkey":
		key, err := strconv.Atoi(s[1])
		if err != nil {
			return nil, fmt.Errorf("atoi error: %w", err)
		}
		return OpenKey(key)
	case "file":
		return OpenMmap(s[1])
	default:
		return nil, fmt.Errorf("unsupport scheme: %v", scheme)
	}
}
