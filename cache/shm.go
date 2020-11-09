package cache

import (
    "unsafe"

    "github.com/PichuChen/go-bbs/shm"
    "github.com/PichuChen/go-bbs/types"
    log "github.com/sirupsen/logrus"
)

type SHM struct {
    Shmid   int
    IsNew   bool
    Shmaddr unsafe.Pointer

    SHMRaw //dummy variable to get the offset and size of the shm-fields.
}

func NewSHM(key types.Key_t, isUseHugeTlb bool, isCreate bool) error {
    if Shm != nil {
        return ErrInvalidOp
    }

    shmid := int(0)
    var shmaddr unsafe.Pointer = nil
    isNew := false
    var err error = nil

    size := types.Size_t(SHM_RAW_SZ)
    log.Infof("NewSHM: SHM_RAW_SZ: %v", SHM_RAW_SZ)

    if isCreate {
        shmid, shmaddr, isNew, err = shm.CreateShm(key, size, isUseHugeTlb)
        if err != nil {
            return err
        }
    } else {
        shmid, shmaddr, err = shm.OpenShm(key, size, isUseHugeTlb)
        if err != nil {
            return err
        }
    }

    Shm = &SHM{
        Shmid:   shmid,
        IsNew:   isNew,
        Shmaddr: shmaddr,
    }

    if isNew {
        in_version := SHM_VERSION
        in_size := int32(SHM_RAW_SZ)
        in_number := int32(0)
        in_loaded := int32(0)
        Shm.WriteAt(
            unsafe.Offsetof(Shm.Version),
            unsafe.Sizeof(Shm.Version),
            unsafe.Pointer(&in_version),
        )
        Shm.WriteAt(
            unsafe.Offsetof(Shm.Size),
            unsafe.Sizeof(Shm.Size),
            unsafe.Pointer(&in_size),
        )
        Shm.WriteAt(
            unsafe.Offsetof(Shm.Number),
            unsafe.Sizeof(Shm.Number),
            unsafe.Pointer(&in_number),
        )
        Shm.WriteAt(
            unsafe.Offsetof(Shm.Loaded),
            unsafe.Sizeof(Shm.Loaded),
            unsafe.Pointer(&in_loaded),
        )
    }

    // version and size should be fixed.
    Shm.ReadAt(
        unsafe.Offsetof(Shm.Version),
        unsafe.Sizeof(Shm.Version),
        unsafe.Pointer(&Shm.Version),
    )
    Shm.ReadAt(
        unsafe.Offsetof(Shm.Size),
        unsafe.Sizeof(Shm.Size),
        unsafe.Pointer(&Shm.Size),
    )

    // verify version
    if Shm.Version != SHM_VERSION {
        log.Errorf("NewSHM: version not match: key: %v Shm.Version: %v SHM_VERSION: %v isCreate: %v isNew: %v", key, Shm.Version, SHM_VERSION, isCreate, isNew)
        CloseSHM()
        return ErrShmVersion
    }
    if isCreate && !isNew {
        log.Warnf("NewSHM: is expected to create, but not: key: %v", key)
    }

    return nil
}

//Close
//
//XXX [WARNING] know what you are doing before using Close!.
//This is to be able to close the shared mem for the completeness of the mem-usage.
//However, in production, we create shm without the need of closing the shm.
//
//We simply use ipcrm to delete the shm if necessary.
//
//Currently used only in test.
func CloseSHM() error {
    if Shm == nil {
        // Already Closed
        log.Errorf("CloseSHM: already closed")
        return ErrShmNotInit
    }

    err := Shm.Close()
    if err != nil {
        log.Errorf("CloseSHM: unable to close: e: %v", err)
        return err
    }

    Shm = nil

    log.Infof("CloseSHM: done")

    return nil
}

//Close
//
//XXX [WARNING] know what you are doing before using Close!.
//This is to be able to close the shared mem for the completeness of the mem-usage.
//However, in production, we create shm without the need of closing the shm.
//
//We simply use ipcrm to delete the shm if necessary.
//
//Currently used only in test.
func (s *SHM) Close() error {
    if !IsTest {
        return ErrInvalidOp
    }
    return shm.CloseShm(s.Shmid)
}

//ReadAt
//
//Require precalculated offset and size and outptr to efficiently get the data.
//See tests for exact usage.
//[!!!] If we are reading from the array, make sure that have unit-size * n in the size.
//
//Params
//  offsetOfSHMRawComponent: offset from SHMRaw
//  size: size of the variable, usually can be referred from SHMRaw
//        [!!!]If we are reading from the array, make sure that have unit-size * n in the size.
//  outptr: the ptr of the object to read.
func (s *SHM) ReadAt(offsetOfSHMRawComponent uintptr, size uintptr, outptr unsafe.Pointer) {
    shm.ReadAt(s.Shmaddr, int(offsetOfSHMRawComponent), types.Size_t(size), outptr)
}

//WriteAt
//
//Require recalculated offset and size and outptr to efficiently get the data.
//See tests for exact usage.
//[!!!]If we are reading from the array, make sure that have unit-size * n in the size.
//
//Params
//  offsetOfSHMRawComponent: offset from SHMRaw
//  size: size of the variable
//        [!!!]If we are reading from the array, make sure that have unit-size * n in the size.
//  inptr: the ptr of the object to write.
func (s *SHM) WriteAt(offsetOfSHMRawComponent uintptr, size uintptr, inptr unsafe.Pointer) {
    shm.WriteAt(s.Shmaddr, int(offsetOfSHMRawComponent), types.Size_t(size), inptr)
}

func (s *SHM) IncUint32(offsetOfSHMRawComponent uintptr) {
    shm.IncUint32(s.Shmaddr, int(offsetOfSHMRawComponent))
}
