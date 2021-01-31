package cache

import (
	"testing"
)

func TestNewCacheWithMmap(t *testing.T) {

	data, err := CreateMmap("./test", 20)
	data.Buf()[0] = 42
	if err != nil {
		t.Logf("err should be nil, got: %v", err)
	}

	data.Close()

	cache, err := NewCache("file:./test")
	if err != nil {
		t.Logf("err should be nil, got: %v", err)
	}

	if cache.Buf()[0] != 42 {
		t.Errorf("cache buf should be %v, got %v", 42, cache.Buf()[0])
	}
	cache.Buf()[0] = 43

	err = cache.Close()

	RemoveMmap("./test")
}

func TestNewCacheWithSHM(t *testing.T) {

	data, err := CreateKey(10, 4)
	data.Buf()[0] = 42
	if err != nil {
		t.Logf("err should be nil, got: %v", err)
	}

	if data.Buf()[0] != 42 {
		t.Errorf("data buf should be %v, got %v", 42, data.Buf()[0])
	}
	data.Close()

	cache, err := NewCache("shmkey:10")
	if err != nil {
		t.Logf("err should be nil, got: %v", err)
	}

	if cache.Buf()[0] != 42 {
		t.Errorf("cache buf should be %v, got %v", 42, cache.Buf()[0])
	}
	cache.Buf()[0] = 43

	err = cache.Close()

	RemoveKey(10)
}
