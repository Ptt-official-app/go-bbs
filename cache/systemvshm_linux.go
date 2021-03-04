// +build linux

package cache

type keyT uint32
type uidT uint32
type gidT uint32
type sizeT uint64
type shmattT uint64

type ShmidDs struct {
	x struct {
		_key keyT
		uid  uidT
		gid  gidT
		cuid uidT
		cgid gidT
		mode int16
		_seq int16
	}
	ShmSegsz    sizeT
	shmAtime    int64
	shmDtime    int64
	shmCtime    int64
	ShmCpid     int32
	ShmLpid     int32
	ShmNattach  shmattT
	shmInternal int64
}
