// +build linux darwin unix

package cache

import (
	"testing"
)

func TestOpenSHM(t *testing.T) {
	data, err := CreateKey(10, 4)
	data.Bytes()[0] = 42
	if err != nil {
		t.Logf("err should be nil, got: %v", err)
	}

	if data.Bytes()[0] != 42 {
		t.Errorf("data buf should be %v, got %v", 42, data.Bytes()[0])
	}
	data.Close()

	data, err = OpenKey(10)
	if err != nil {
		t.Logf("err should be nil, got: %v", err)
	}

	if data.Bytes()[0] != 42 {
		t.Errorf("data buf should be %v, got %v", 42, data.Bytes()[0])
	}
	data.Bytes()[0] = 43

	err = data.Close()
	if err != nil {
		t.Logf("err should be nil, got: %v", err)
	}

	data, err = OpenKey(10)
	if err != nil {
		t.Logf("err should be nil, got: %v", err)
	}

	if data.Bytes()[0] != 43 {
		t.Errorf("data buf should be %v, got %v", 43, data.Bytes()[0])
	}
	RemoveKey(10)
}
