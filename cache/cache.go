package cache

import (
	"bytes"
	"unsafe"

	"github.com/PichuChen/go-bbs/cmsys"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

//AttachSHM
//Should be used after loadUHash (shmctl init) is done.
//Should be used only once in the beginning of the program.
func AttachSHM() error {
	if Shm != nil {
		return nil
	}

	return NewSHM(types.Key_t(ptttype.SHM_KEY), ptttype.USE_HUGETLB, false)
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
func SearchUser(userID string, isReturn bool) (usernum int32, rightID string, err error) {
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
func doSearchUser(userID string, isReturn bool) (usernum int32, rightID string, err error) {
	userIDBytes := &[ptttype.IDLEN + 1]byte{}
	copy(userIDBytes[:], []byte(userID))

	var rightIDBytes *[ptttype.IDLEN + 1]byte = nil
	if isReturn {
		rightIDBytes = &[ptttype.IDLEN + 1]byte{}
	}

	usernum, err = doSearchUserRaw(userIDBytes, rightIDBytes)
	if err != nil {
		return 0, "", err
	}

	rightID = ""
	if isReturn {
		rightID = types.CstrToString(rightIDBytes[:])
	}

	return usernum, rightID, nil
}

func doSearchUserRaw(userID *[ptttype.IDLEN + 1]byte, rightID *[ptttype.IDLEN + 1]byte) (int32, error) {
	// XXX we should have 0 as non-exists.
	//     currently the reason why it's ok is because the probability of collision on 0 is low.

	StatInc(ptttype.STAT_SEARCHUSER)
	h := cmsys.StringHash(userID[:]) % (1 << ptttype.HASH_BITS)

	//p = SHM->hash_head[h]  //line: 219
	p := int32(0)
	Shm.ReadAt(
		unsafe.Offsetof(Shm.HashHead)+types.INT32_SZ*uintptr(h),
		types.INT32_SZ,
		unsafe.Pointer(&p),
	)

	shmUserID := [ptttype.IDLEN + 1]byte{}

	for times := 0; times < ptttype.MAX_USERS && p != -1 && p < ptttype.MAX_USERS; times++ {
		//if (strcasecmp(SHM->userid[p], userid) == 0)  //line: 222
		Shm.ReadAt(
			unsafe.Offsetof(Shm.Userid)+ptttype.USER_ID_SZ*uintptr(p),
			ptttype.USER_ID_SZ,
			unsafe.Pointer(&shmUserID),
		)
		if bytes.Compare(bytes.ToUpper(userID[:]), bytes.ToUpper(shmUserID[:])) == 0 {
			if userID[0] != 0 && rightID != nil {
				copy(rightID[:], shmUserID[:])
			}
			return p + 1, nil
		}
		Shm.ReadAt(
			unsafe.Offsetof(Shm.NextInHash)+types.INT32_SZ*uintptr(p),
			types.INT32_SZ,
			unsafe.Pointer(&p),
		)
	}

	return 0, nil
}
