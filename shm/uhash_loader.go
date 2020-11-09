package shm

// #cgo CFLAGS: -Os -Wno-missing-field-initializers -pipe -I./include -Wno-parentheses-equality
// #cgo linux LDFLAGS: -lbsd
// #cgo darwin LDFLAGS: -liconv
// #include "uhash_loader.h"
import "C"
import (
	log "github.com/sirupsen/logrus"
)

func LoadHash() error {
	// Always got "file-exists error" in the error.
	// err is not an indicator of wrong-op in C.load_uhash.
	// use ret instead.
	ret, _ := C.load_uhash()
	log.Infof("after load_uhash: ret: %v", ret)
	if ret != 0 {
		return ErrInvalidOp
	}

	return nil
}
