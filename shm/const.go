package shm

// #include <sys/shm.h>
//
// /***
//  * XXX hack for SHM_HUGETLB as 0
//  * SHM_HUGETLB is considered as flag
//  * flag | SHM_HUGETLB == flag => SHM_HUGETLB is not effective.
//  * flag & SHM_HUGETLB = 0 ==> SHM_HUGETLB is not set.
//  ***/
// #ifndef SHM_HUGETLB
// #define SHM_HUGETLB 0
// #endif
import "C"

const (
	IpcCreat   = C.IPC_CREAT
	IpcExcl    = C.IPC_EXCL
	ShmHugeTlb = C.SHM_HUGETLB
)
