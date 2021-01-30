package cache

// #include <sys/ipc.h>
// #include <sys/shm.h>
import "C"

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

// shmget: get shared memort area identifier, please see man 2 shmget
func Shmget(key int, size int, flag int) (int, error) {
	return shmget(key, size, flag)
}

// shmat: map/unmap shared memory, plase see man 2 shmat.
// golang gc won't affect on this.
func Shmat(shmid int, shmaddr uintptr, shmflg int) (uintptr, error) {
	return shmat(shmid, shmaddr, shmflg)
}

// shmdt: map/unmap shared memory, plase see man 2 shmdt.
func Shmdt(shmaddr uintptr) (int, error) {
	return shmdt(shmaddr)
}

// shmctl: shared memory control operations
func Shmctl(shmid int, cmd int, buf *ShmidDs) (int, error) {
	return shmctl(shmid, cmd, buf)
}

type ShmidDs struct {
	x struct {
		uid  int32
		gid  int32
		cuid int32
		cgid int32
		mode int32
		_seq int16
		_key uint16
	}
	ShmSegsz    int32
	ShmLpid     int32
	ShmCpid     int32
	ShmNattach  int16
	shmAtime    int64
	shmDtime    int64
	shmCtime    int64
	shmInternal int64
}

const (
	IPCCreate = 00001000

	IPCRMID = 0
	IPCSet  = 1
	IPCStat = 2
)

type SHM struct {
	Buf []byte
}

func CreateKey(key int, size int) (*SHM, error) {
	flag := 0
	if size != 0 {
		// create
		flag = IPCCreate | 0600
	}

	shmId, err := Shmget(key, size, flag)
	if err != nil {
		return nil, fmt.Errorf("shmget error: %v", err)
	}
	// fmt.Println("shmid", shmId)
	// size = 4
	if size == 0 {

		ds := ShmidDs{}
		_, err = Shmctl(shmId, IPCStat, &ds)
		if err != nil {
			return nil, fmt.Errorf("shmctl error: %v", err)
		}
		// fmt.Println(ds)
		size = int(ds.ShmSegsz)
	}

	// v := 0

	ptr, err := Shmat(shmId, uintptr(0), 0)
	if err != nil {
		return nil, fmt.Errorf("shmat error: %v", err)
	}
	// fmt.Println("ptr", *(*int)(ptr))
	// fmt.Println("size", size)
	b := []byte{}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	hdr.Cap = size
	hdr.Len = size
	hdr.Data = ptr

	ret := SHM{
		Buf: b,
	}

	// For GC
	runtime.SetFinalizer(&ret, (*SHM).Close)
	return &ret, nil
}

// OpenKey open a shm which is already exists.
func OpenKey(key int) (*SHM, error) {
	return CreateKey(key, 0)
}
func RemoveKey(key int) error {
	shmId, err := Shmget(key, 0, 0)
	if err != nil {
		return fmt.Errorf("shmget error: %v", err)
	}

	// ds := ShmidDs{}
	_, err = Shmctl(shmId, IPCRMID, nil)
	if err != nil {
		return fmt.Errorf("shmctl error: %v", err)
	}
	return nil
}
func (m *SHM) Close() error {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&m.Buf))
	ptr := hdr.Data
	_, err := Shmdt(ptr)

	// release GC setting
	runtime.SetFinalizer(m, nil)
	if err != nil {
		return fmt.Errorf("shmdt error: %v", err)
	}
	return nil
}

// func shmNew(key int, size int) int {

// 	shmid, err := Shmget(key, size, 0)

// 	shmptr := C.shmat(C.int(shmid), unsafe.Pointer(nil), 0)
// 	fmt.Printf("nowValue = %d, pos = %v %v\n", *(*int)(shmptr), (*int)(shmptr), shmptr)

// 	return int(shmid)

// }

// void * shm_new(const int key, const int size) {
// 	shmid = shmget(key, size, 0);
// 	printf("shmid: %d %s \n", shmid, strerror(errno));
// 	if(shmid == -1){
// 		// Create
// 		shmid = shmget(key, size, IPC_CREAT | 0600);
// 		if(shmid == -1){
// 			printf("shmid: %d %s \n", shmid, strerror(errno));
// 			exit(-1);
// 		}
// 	}

// 	void * shmptr = (void *) shmat(shmid, NULL, 0);
// 	if((int)shmptr == -1){
// 		printf("shmat: %p %s \n", shmptr, strerror(errno));
// 		exit(-2);
// 	}
// 	printf("nowValue = %d, pos: %p\n", *(int*)shmptr, shmptr);
// 	*(int*)shmptr = *(int*)shmptr + 1;
// 	return shmptr;
// }

// func main() {
// 	shmNew(3, 4)
// 	fmt.Println("OK")
// }
