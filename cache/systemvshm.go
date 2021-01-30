package cache

// #include <sys/ipc.h>
// #include <sys/shm.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func Shmget(key int, size int, flag int) int {
	return int(C.shmget(C.int(key), C.ulong(size), C.int(flag)))
}

func shmNew(key int, size int) int {

	shmid := Shmget(key, size, 0)

	shmptr := C.shmat(C.int(shmid), unsafe.Pointer(nil), 0)
	fmt.Printf("nowValue = %d, pos = %v %v\n", *(*int)(shmptr), (*int)(shmptr), shmptr)

	return int(shmid)

}

// func main() {
// 	shmNew(3, 4)
// 	fmt.Println("OK")
// }
