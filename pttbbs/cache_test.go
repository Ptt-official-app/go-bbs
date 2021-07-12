package pttbbs

import (
	"testing"
)

// type Cache struct {
// 	cache.Cache
// 	*MemoryMappingSetting
// 	cachePos
// }

type MockCache struct {
	buf []byte
}

func (c *MockCache) Bytes() []byte {
	return c.buf
}

func (c *MockCache) Close() error {
	return nil
}

func TestGetVersion(t *testing.T) {
	c := Cache{
		Cache: &MockCache{
			buf: []byte{12, 0, 0, 0},
		},
		MemoryMappingSetting: &MemoryMappingSetting{
			AlignmentBytes: 2,
			MaxUsers:       10,
			IDLen:          12,
		},
	}
	c.caculatePos()

	if c.Version() != 12 {
		t.Errorf("version excepted 1 got: %v", c.Version())
	}
}
