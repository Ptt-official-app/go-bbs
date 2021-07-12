package cache

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

const (
	IPCCreate = 0o0001000

	IPCRMID = 0
	IPCSet  = 1
	IPCStat = 2
)

type SHM struct {
	buf []byte
}

func CreateKey(key int, size int) (*SHM, error) {
	flag := 0
	if size != 0 {
		// create
		flag = IPCCreate | 0o600
	}

	shmID, err := Shmget(key, size, flag)
	if err != nil {
		return nil, fmt.Errorf("shmget error: %w", err)
	}
	// fmt.Println("shmid", shmID)
	// size = 4
	if size == 0 {

		ds := ShmidDs{}
		_, err = Shmctl(shmID, IPCStat, &ds)
		if err != nil {
			return nil, fmt.Errorf("shmctl error: %w", err)
		}
		// fmt.Printf("%+v\n", ds)
		size = int(ds.ShmSegsz)
	}

	// v := 0

	ptr, err := Shmat(shmID, uintptr(0), 0)
	if err != nil {
		return nil, fmt.Errorf("shmat error: %w", err)
	}
	// fmt.Println("ptr", *(*int)(ptr))
	// fmt.Println("size", size)
	b := []byte{}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	hdr.Cap = size
	hdr.Len = size
	hdr.Data = ptr

	ret := SHM{
		buf: b,
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
	shmID, err := Shmget(key, 0, 0)
	if err != nil {
		return fmt.Errorf("shmget error: %w", err)
	}

	// ds := ShmidDs{}
	_, err = Shmctl(shmID, IPCRMID, nil)
	if err != nil {
		return fmt.Errorf("shmctl error: %w", err)
	}
	return nil
}

func (s *SHM) Bytes() []byte {
	return s.buf
}

func (s *SHM) Close() error {
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s.buf))
	ptr := hdr.Data
	_, err := Shmdt(ptr)

	// release GC setting
	runtime.SetFinalizer(s, nil)
	if err != nil {
		return fmt.Errorf("shmdt error: %w", err)
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
