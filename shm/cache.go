package shm

// #cgo CFLAGS: -Os -Wno-missing-field-initializers -pipe -I./include -Wno-parentheses-equality
// #cgo linux LDFLAGS: -lbsd
// #cgo darwin LDFLAGS: -liconv
// #include "cache.h"
// #include "proto.h"
import "C"
import "unsafe"

//AttachSHM
//Should be used after loadUHash (shmctl init) is done.
//Should be used only once in the beginning of the program.
func AttachSHM() error {
	_, err := C.attach_SHM()
	return err
}

//SearchUser
//Params:
//	userID: querying user-id.
//	isReturn: is return the user-id in the shm.
//
//Return:
//	int: usernum.
//	string: the userID in shm.
//	error: err.
func SearchUser(userID string, isReturn bool) (int, string, error) {
	if len(userID) == 0 {
		return 0, "", nil
	}
	return doSearchUser(userID, isReturn)
}

//doSearchUser
//Params:
//	userID
//	isReturn
//
//Return:
//	int: usernum.
//	string: the userID in shm.
//	error: err.
func doSearchUser(userID string, isReturn bool) (int, string, error) {
	cuserID := C.CString(userID)
	defer C.free(unsafe.Pointer(cuserID))

	rightID := ""

	crightID := C.CString(rightID)
	defer C.free(unsafe.Pointer(crightID))

	_, _ = C.syncnow()

	crightID2 := crightID
	if !isReturn {
		crightID2 = nil
	}

	cusernum, err := C.dosearchuser(cuserID, crightID2)
	if err != nil {
		return 0, "", err
	}

	returnID := ""
	if isReturn {
		returnID = C.GoString(crightID2)
	}

	return int(cusernum), returnID, nil
}
