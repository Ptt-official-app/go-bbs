package shm

// #cgo CFLAGS: -Os -Wno-missing-field-initializers -pipe -I./include -Wno-parentheses-equality
// #cgo linux LDFLAGS: -lbsd
// #cgo darwin LDFLAGS: -liconv
// #include "cache.h"
// #include "proto.h"
// #include "var.h"
// int dosearchuserwrapper(void *userid, void *rightid) {
//   return dosearchuser((char *)userid, (char *)rightid);
// }
// void attach_SHM_wrapper() {
//   if(SHM != NULL) {
//	   return;
//   }
//   attach_SHM();
// }
import "C"
import (
	"unsafe"

	"github.com/PichuChen/go-bbs/ptttype"
)

//AttachSHM
//Should be used after loadUHash (shmctl init) is done.
//Should be used only once in the beginning of the program.
func AttachSHM() error {
	_, err := C.attach_SHM_wrapper()
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

	rightID := string(make([]byte, ptttype.IDLEN+1))
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

//SearchUserBig5
//Params:
//	userID: querying user-id. (CBytes provides only []byte)
//	isReturn
//
//Return:
//	int: usernum.
//  []byte: the user-id in shm (CBytes provides only []byte)
//	error: err.
func SearchUserBig5(userID []byte, isReturn bool) (int, []byte, error) {
	if userID[0] == 0 {
		return 0, nil, nil
	}
	return doSearchUserBig5(userID, isReturn)
}

//doSearchUserBig5
//Params:
//	userID: querying user-id. (CBytes provides only []byte)
//	isReturn
//
//Return:
//	int: usernum.
//  []byte: the user-id in shm (CBytes provides only []byte)
//	error: err.
func doSearchUserBig5(userID []byte, isReturn bool) (int, []byte, error) {
	cuserID := C.CBytes(userID[:])
	defer C.free(cuserID)

	var crightID unsafe.Pointer = nil
	if isReturn {
		rightID := make([]byte, ptttype.IDLEN+1)
		crightID = C.CBytes(rightID)
		defer C.free(crightID)
	}

	_, _ = C.syncnow()

	cusernum, err := C.dosearchuserwrapper(cuserID, crightID)
	if err != nil {
		return 0, nil, err
	}

	var returnID []byte = nil
	if isReturn {
		returnID = C.GoBytes(crightID, ptttype.IDLEN+1)
	}

	return int(cusernum), returnID, nil
}
