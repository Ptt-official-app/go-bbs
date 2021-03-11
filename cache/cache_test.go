// +build darwin linux unix

package cache

import (
	"testing"
)

func TestNewCacheWithMmap(t *testing.T) {

	data, err := CreateMmap("./test", 20)
	if err != nil {
		t.Logf("err should be nil, got: %v", err)
	}

	data.Bytes()[0] = 42

	data.Close()

	cache, err := NewCache("file:./test")
	if err != nil {
		t.Logf("err should be nil, got: %v", err)
	}

	if cache.Bytes()[0] != 42 {
		t.Errorf("cache buf should be %v, got %v", 42, cache.Bytes()[0])
	}
	cache.Bytes()[0] = 43

	err = cache.Close()
	if err != nil {
		t.Logf("err should be nil, got: %v", err)
	}

	RemoveMmap("./test")
}

func TestNewCacheWithSHM(t *testing.T) {

	data, err := CreateKey(10, 4)
	data.Bytes()[0] = 42
	if err != nil {
		t.Logf("err should be nil, got: %v", err)
	}

	if data.Bytes()[0] != 42 {
		t.Errorf("data buf should be %v, got %v", 42, data.Bytes()[0])
	}
	data.Close()

	cache, err := NewCache("shmkey:10")
	if err != nil {
		t.Logf("err should be nil, got: %v", err)
	}

	if len(cache.Bytes()) != 4 {
		t.Errorf("cache buf length not correct, expected 4, got: %v", len(cache.Bytes()))
	}

	if cache.Bytes()[0] != 42 {
		t.Errorf("cache buf should be %v, got %v", 42, cache.Bytes()[0])
	}
	cache.Bytes()[0] = 43

	err = cache.Close()

	RemoveKey(10)
}
