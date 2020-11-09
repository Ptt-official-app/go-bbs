package shm

// #cgo CFLAGS: -Os -Wno-missing-field-initializers -pipe -I./include -Wno-parentheses-equality
// #cgo linux LDFLAGS: -lbsd
// #cgo darwin LDFLAGS: -liconv
// #include "cache.h"
import "C"
import "unsafe"

func AttachSHM() error {
	_, err := C.attach_SHM()
	return err
}

func DoSearchUser(userID string, rightID string) (int, error) {
	cuserID := C.CString(userID)
	defer C.free(unsafe.Pointer(cuserID))

	crightID := C.CString(rightID)
	defer C.free(unsafe.Pointer(crightID))

	cusernum, err := C.dosearchuser(cuserID, crightID)
	if err != nil {
		return 0, err
	}

	return int(cusernum), nil
}
