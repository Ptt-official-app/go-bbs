// +build darwin

package cache

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
	ShmSegsz    uint64
	ShmLpid     int32
	ShmCpid     int32
	ShmNattach  int16
	shmAtime    int64
	shmDtime    int64
	shmCtime    int64
	shmInternal int64
}
