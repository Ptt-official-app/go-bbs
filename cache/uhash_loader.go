package cache

import (
	"io"
	"os"
	"reflect"
	"unsafe"

	"github.com/PichuChen/go-bbs/cmbbs"
	"github.com/PichuChen/go-bbs/cmsys"
	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
	log "github.com/sirupsen/logrus"
)

//LoadUHash
//Load user-hash into SHM.
func LoadUHash() (err error) {
	if Shm == nil {
		return ErrShmNotInit
	}

	// line: 58
	number := int32(0)
	Shm.ReadAt(
		unsafe.Offsetof(Shm.Number),
		unsafe.Sizeof(Shm.Number),
		unsafe.Pointer(&number),
	)

	loaded := int32(0)
	Shm.ReadAt(
		unsafe.Offsetof(Shm.Loaded),
		unsafe.Sizeof(Shm.Loaded),
		unsafe.Pointer(&loaded),
	)

	//XXX in case it's not assumed zero, this becomes a race...
	if number == 0 && loaded == 0 {
		// line: 60
		fillUHash(false)

		// line: 61
		zeroByte := '\x00'
		Shm.WriteAt(
			unsafe.Offsetof(Shm.TodayIs),
			unsafe.Sizeof(Shm.TodayIs[0]),
			unsafe.Pointer(&zeroByte),
		)

		// line: 62
		loaded = 1
		Shm.WriteAt(
			unsafe.Offsetof(Shm.Loaded),
			unsafe.Sizeof(Shm.Loaded),
			unsafe.Pointer(&loaded),
		)
	} else {
		// line: 65
		fillUHash(true)
	}

	return nil
}

var (
	uHashLoaderInvalidUserID = 0
)

func fillUHash(isOnfly bool) error {
	log.Infof("fillUHash: start: isOnfly: %v", isOnfly)
	initFillUHash(isOnfly)

	file, err := os.Open(ptttype.FN_PASSWD)
	if err != nil {
		return err
	}

	usernum := int32(0)

	uHashLoaderInvalidUserID = 0
	for ; ; usernum++ {
		userecRaw, eachErr := ptttype.NewUserecRawWithFile(file)
		if eachErr != nil {
			// io.EOF is reading correctly to the end the file.
			if eachErr == io.EOF {
				break
			}

			err = eachErr
			break
		}

		userecRawAddToUHash(usernum, userecRaw, isOnfly)
	}

	Shm.WriteAt(
		unsafe.Offsetof(Shm.Number),
		unsafe.Sizeof(Shm.Number),
		unsafe.Pointer(&usernum),
	)
	return nil
}

func userecRawAddToUHash(usernum int32, userecRaw *ptttype.UserecRaw, isOnfly bool) {
	// uhash use userid="" to denote free slot for new register
	// However, such entries will have the same hash key.
	// So we skip most of invalid userid to prevent lots of hash collision.
	if !cmbbs.IsValidUserID(&userecRaw.UserID) {
		// dirty hack, preserve few slot for new register
		uHashLoaderInvalidUserID++
		if uHashLoaderInvalidUserID > 1000 {
			return
		}
	}

	h := cmsys.StringHash(userecRaw.UserID[:]) % (1 << ptttype.HASH_BITS)

	shmUserID := [ptttype.IDLEN + 1]byte{}
	Shm.ReadAt(
		unsafe.Offsetof(Shm.Userid)+ptttype.USER_ID_SZ*uintptr(usernum),
		ptttype.USER_ID_SZ,
		unsafe.Pointer(&shmUserID),
	)

	offsetNextInHash := unsafe.Offsetof(Shm.NextInHash)

	if !isOnfly || !reflect.DeepEqual(userecRaw.UserID, shmUserID) {
		Shm.WriteAt(
			unsafe.Offsetof(Shm.Userid)+ptttype.USER_ID_SZ*uintptr(usernum),
			ptttype.USER_ID_SZ,
			unsafe.Pointer(&userecRaw.UserID),
		)

		Shm.WriteAt(
			unsafe.Offsetof(Shm.Money)+types.INT32_SZ*uintptr(usernum),
			types.INT32_SZ,
			unsafe.Pointer(&userecRaw.Money),
		)

		if ptttype.USE_COOLDOWN {
			zero := types.Time4(0)
			Shm.WriteAt(
				unsafe.Offsetof(Shm.CooldownTime)+types.TIME4_SZ*uintptr(usernum),
				types.TIME4_SZ,
				unsafe.Pointer(&zero),
			)
		}
		log.Debugf("UHashLoader.userecRawAddToUHash: add info: usernum: %v id: %v shmUserID: %v", usernum, string(userecRaw.UserID[:]), string(shmUserID[:]))
	}

	p := h
	val := int32(0)
	offsetHashHead := unsafe.Offsetof(Shm.HashHead)
	//offsetNextInHash := unsafe.Offsetof(Shm.NextInHash)
	isFirst := true

	Shm.ReadAt(
		offsetHashHead+types.INT32_SZ*uintptr(p),
		types.INT32_SZ,
		unsafe.Pointer(&val),
	)

	l := 0
	for val >= 0 && val < ptttype.MAX_USERS {
		if isOnfly && val == usernum { // already in hash
			return
		}

		l++
		// go to next
		// 1. setting p as val
		// 2. get val from next_in_hash[p]
		p = uint32(val)
		Shm.ReadAt(
			offsetNextInHash+types.INT32_SZ*uintptr(p),
			types.INT32_SZ,
			unsafe.Pointer(&val),
		)

		isFirst = false
	}

	// set next in hash as n
	offset := offsetHashHead
	if !isFirst {
		offset = offsetNextInHash
	}
	val = usernum
	Shm.WriteAt(
		offset+types.INT32_SZ*uintptr(p),
		types.INT32_SZ,
		unsafe.Pointer(&val),
	)

	log.Infof("UHashLoader.userecRawAddToUHash: added level: %v p: %v hash: %v usernum: %v [%v] val: %v in hash isHashHead: %v", l, p, h, usernum, string(userecRaw.UserID[:]), val, isFirst)

	// set next in hash as -1
	p = uint32(val)
	val = -1
	Shm.WriteAt(
		offsetNextInHash+types.INT32_SZ*uintptr(p),
		types.INT32_SZ,
		unsafe.Pointer(&val),
	)
	log.Debugf("UHashLoader.userecRawAddToUHash: added NextInHash: usernum: %v p: %v val: %v isFirst: %v", usernum, p, val, isFirst)
}

func initFillUHash(isOnfly bool) {
	if !isOnfly {
		toFillHashHead := [1 << ptttype.HASH_BITS]int32{}
		for idx := range toFillHashHead {
			toFillHashHead[idx] = -1
		}
		Shm.WriteAt(
			unsafe.Offsetof(Shm.HashHead),
			unsafe.Sizeof(Shm.HashHead),
			unsafe.Pointer(&toFillHashHead),
		)
	} else {
		for idx := uint32(0); idx < (1 << ptttype.HASH_BITS); idx++ {
			checkHash(idx)
		}
	}
}

func checkHash(h uint32) {
	// p as delegate-pointer to the Shm.
	// in the beginning, p is the indicator of HashHead.
	// after 1st for-loop, p is in nextInHash.
	// val as the corresponding *p

	// line: 71
	p := h
	val := int32(0)
	pval := &val
	valptr := unsafe.Pointer(pval)
	Shm.ReadAt(
		unsafe.Offsetof(Shm.HashHead)+types.INT32_SZ*uintptr(p),
		types.INT32_SZ,
		valptr,
	)
	if val != -1 {
		log.Infof("checkHash: after read at HashHead: h: %v p: %v val: %v", h, p, val)
	}

	// line: 72
	isFirst := true

	var offset uintptr = 0
	offsetHashHead := unsafe.Offsetof(Shm.HashHead)
	offsetNextInHash := unsafe.Offsetof(Shm.NextInHash)

	userID := [ptttype.IDLEN + 1]byte{}
	deep := 0
	for val != -1 {
		offset = offsetNextInHash
		if isFirst {
			offset = offsetHashHead
		}

		// check invalid pointer-val, set as -1  line: 74
		if val < -1 || val >= ptttype.MAX_USERS {
			log.Warnf("uhash_loader.checkHash: val invalid: h: %v p: %v val: %v isHead: %v", h, p, val, isFirst)
			*pval = -1
			Shm.WriteAt(
				offset+types.INT32_SZ*uintptr(p),
				types.INT32_SZ,
				valptr,
			)
			break
		}

		// get user-id: line: 75
		Shm.ReadAt(
			unsafe.Offsetof(Shm.Userid)+ptttype.USER_ID_SZ*uintptr(val),
			ptttype.USER_ID_SZ,
			unsafe.Pointer(&userID),
		)
		log.Infof("checkHash: (in-for-loop): after read userID: h: %v p: %v val: %v userID: %v", h, p, val, types.CstrToString(userID[:]))

		userIDHash := cmsys.StringHash(userID[:]) % (1 << ptttype.HASH_BITS)

		// check hash as expected line: 76
		if userIDHash != h {
			// XXX
			// the result of the userID does not fit the h (broken?).
			// XXX uhash_loader is used only 1-time when starting the service.
			next := int32(0)

			// get next from *p (val)
			Shm.ReadAt(
				offsetNextInHash+types.INT32_SZ*uintptr(val),
				types.INT32_SZ,
				unsafe.Pointer(&next),
			)
			log.Warnf("userID hash is not in the corresponding idx (to remove) (%v): userID: %v userIDHash: %v h: %v next: %v", deep, types.CstrToString(userID[:]), userIDHash, h, next)
			// remove current by setting current as the next, hopefully the next user can fit the userIDHash.
			*pval = next
			Shm.WriteAt(
				offset+types.INT32_SZ*uintptr(p),
				types.INT32_SZ,
				unsafe.Pointer(&next),
			)
		} else {
			// 1. p as val (pointer in NextInHash)
			// 2. update val as NextInHash[p]
			p = uint32(val)
			Shm.ReadAt(
				offsetNextInHash+types.INT32_SZ*uintptr(p),
				types.INT32_SZ,
				unsafe.Pointer(&val),
			)
			isFirst = false

			log.Infof("checkHash: (in-for-loop (match)): after read next: h: %v p: %v val: %v userID: %v isFirst: %v", h, p, val, types.CstrToString(userID[:]), isFirst)
		}

		// line: 87
		deep++
		if deep == 100 {
			log.Warnf("checkHash deadlock: deep: %v h: %v p: %v val: %v isFirst: %v", deep, h, p, val, isFirst)
			break
		}
	}
}
