package shm

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/PichuChen/go-bbs/types"
	"github.com/sirupsen/logrus"
)

type testStruct struct {
	A int32
	B testStruct2
}

const TEST_STRUCT_SZ = unsafe.Sizeof(testStruct{})

type testStruct2 struct {
	C [10]uint8
}

func TestCreateShm(t *testing.T) {
	type args struct {
		key           types.Key_t
		size          types.Size_t
		is_usehugetlb bool
	}
	tests := []struct {
		name        string
		args        args
		wantShmid   int
		wantShmaddr unsafe.Pointer
		wantIsNew   bool
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			args:      args{testShmKey, 100, false},
			wantShmid: 0,
			wantIsNew: true,
		},
		{
			args:      args{testShmKey, 100, false},
			wantShmid: 0,
			wantIsNew: false,
		},
	}

	firstGoodShmID := 0
	for _, tt := range tests {
		if firstGoodShmID != 0 {
			tt.wantShmid = firstGoodShmID
		}

		t.Run(tt.name, func(t *testing.T) {
			gotShmid, _, gotIsNew, err := CreateShm(tt.args.key, tt.args.size, tt.args.is_usehugetlb)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateShm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantShmid != 0 && gotShmid != tt.wantShmid {
				t.Errorf("CreateShm() gotShmid = %v, expected %v", gotShmid, tt.wantShmid)
			}
			if gotIsNew != tt.wantIsNew {
				t.Errorf("CreateShm() gotIsNew = %v, expected %v", gotIsNew, tt.wantIsNew)
			}
		})
	}
	CloseShm(firstGoodShmID)
}

func TestCloseShm(t *testing.T) {
	gotShmid, _, _, _ := CreateShm(testShmKey, 100, false)

	type args struct {
		shmid int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{gotShmid},
		},
		{
			args:    args{gotShmid},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CloseShm(tt.args.shmid); (err != nil) != tt.wantErr {
				t.Errorf("CloseShm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOpenShm(t *testing.T) {
	shmid, _, _, _ := CreateShm(testShmKey, 100, false)
	defer CloseShm(shmid)

	type args struct {
		key           types.Key_t
		size          types.Size_t
		is_usehugetlb bool
	}
	tests := []struct {
		name        string
		args        args
		wantShmid   int
		wantShmaddr unsafe.Pointer
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			name:      "same size",
			args:      args{testShmKey, 100, false},
			wantShmid: shmid,
		},
		{
			name:      "same size 2",
			args:      args{testShmKey, 100, false},
			wantShmid: shmid,
		},
		{
			name:      "diff size",
			args:      args{testShmKey, 500, false},
			wantShmid: -1,
			wantErr:   true,
		},
		{
			name:      "diff key",
			args:      args{testShmKey + 1, 500, false},
			wantShmid: -1,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotShmid, _, err := OpenShm(tt.args.key, tt.args.size, tt.args.is_usehugetlb)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenShm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotShmid != tt.wantShmid {
				t.Errorf("OpenShm() gotShmid = %v, expected %v", gotShmid, tt.wantShmid)
			}
		})
	}

}

func TestReadAt(t *testing.T) {
	shmid, shmaddr, _, _ := CreateShm(testShmKey, 40, false)
	defer CloseShm(shmid)

	_, shmaddr2, _ := OpenShm(testShmKey, 40, false)

	empty := [40]byte{}
	emptyptr := unsafe.Pointer(&empty)
	WriteAt(shmaddr, 0, 40, emptyptr)

	test1 := &testStruct{}
	test1.A = 10
	copy(test1.B.C[:], []byte("0123456789"))
	test1ptr := unsafe.Pointer(test1)
	WriteAt(shmaddr, 12, types.Size_t(TEST_STRUCT_SZ), test1ptr)
	logrus.Infof("TEST_STRUCT_SZ: %v", TEST_STRUCT_SZ)

	test2 := &testStruct{}
	test2ptr := unsafe.Pointer(test2)

	test3 := &[40]byte{}
	test3ptr := unsafe.Pointer(test3)

	test4 := &[40]byte{}
	test4ptr := unsafe.Pointer(test4)

	want3 := &[40]byte{}
	want3[12] = 10
	copy(want3[12+4:], []byte("0123456789"))

	type args struct {
		shmaddr unsafe.Pointer
		offset  int
		size    types.Size_t
		outptr  unsafe.Pointer
	}
	tests := []struct {
		name     string
		args     args
		read     interface{}
		expected interface{}
	}{
		// TODO: Add test cases.
		{
			name:     "read test2 (testStruct)",
			args:     args{shmaddr: shmaddr, offset: 12, size: types.Size_t(TEST_STRUCT_SZ), outptr: test2ptr},
			read:     test2,
			expected: test1,
		},
		{
			name:     "read test3 (bytes)",
			args:     args{shmaddr: shmaddr, offset: 0, size: 40, outptr: test3ptr},
			read:     test3,
			expected: want3,
		},
		{
			name:     "read test4 (bytes) from shmaddr2",
			args:     args{shmaddr: shmaddr2, offset: 0, size: 40, outptr: test4ptr},
			read:     test4,
			expected: want3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadAt(tt.args.shmaddr, tt.args.offset, tt.args.size, tt.args.outptr)

			if !reflect.DeepEqual(tt.read, tt.expected) {
				t.Errorf("ReadAt() read: %v expected: %v", tt.read, tt.expected)
			}
		})
	}

}

func TestWriteAt(t *testing.T) {
	shmid, shmaddr, _, _ := CreateShm(testShmKey, 40, false)
	defer CloseShm(shmid)

	empty := [40]byte{}
	emptyptr := unsafe.Pointer(&empty)

	empty1 := [40]byte{}

	WriteAt(shmaddr, 0, 40, emptyptr)

	read := [40]byte{}
	readptr := unsafe.Pointer(&read)

	test1 := &testStruct{}
	test1.A = 10
	copy(test1.B.C[:], []byte("0123456789"))
	test1ptr := unsafe.Pointer(test1)

	want1 := [40]byte{}
	want1[0] = 10
	copy(want1[0+4:], []byte("0123456789"))

	want2 := [40]byte{}
	copy(want2[:12], want1[:])
	want2[12] = 10
	copy(want2[12+4:], []byte("0123456789"))

	type args struct {
		shmaddr unsafe.Pointer
		offset  int
		size    types.Size_t
		inptr   unsafe.Pointer
	}
	tests := []struct {
		name     string
		args     args
		expected [40]byte
	}{
		// TODO: Add test cases.
		{
			name:     "empty (1)",
			args:     args{shmaddr, 0, 40, emptyptr},
			expected: empty1,
		},
		{
			name:     "test",
			args:     args{shmaddr, 0, types.Size_t(TEST_STRUCT_SZ), test1ptr},
			expected: want1,
		},
		{
			name:     "test-want2",
			args:     args{shmaddr, 12, types.Size_t(TEST_STRUCT_SZ), test1ptr},
			expected: want2,
		},
		{
			name:     "empty (2)",
			args:     args{shmaddr, 0, 40, emptyptr},
			expected: empty1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteAt(tt.args.shmaddr, tt.args.offset, tt.args.size, tt.args.inptr)

			ReadAt(shmaddr, 0, 40, readptr)
			if !reflect.DeepEqual(read, tt.expected) {
				t.Errorf("WriteAt() = %v expected: %v", read, tt.expected)
			}

		})
	}

}

func TestIncUint32(t *testing.T) {
	shmid, shmaddr, _, _ := CreateShm(testShmKey, 40, false)
	defer CloseShm(shmid)

	empty := [40]byte{}
	emptyptr := unsafe.Pointer(&empty)
	WriteAt(shmaddr, 0, 40, emptyptr)

	type args struct {
		shmaddr unsafe.Pointer
		offset  int
	}
	tests := []struct {
		name     string
		args     args
		expected uint32
	}{
		// TODO: Add test cases.
		{
			args:     args{shmaddr: shmaddr, offset: 1},
			expected: 1,
		},
		{
			args:     args{shmaddr: shmaddr, offset: 1},
			expected: 2,
		},
		{
			args:     args{shmaddr: shmaddr, offset: 2},
			expected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IncUint32(tt.args.shmaddr, tt.args.offset)

			out := uint32(0)
			ReadAt(tt.args.shmaddr, tt.args.offset, types.Size_t(unsafe.Sizeof(out)), unsafe.Pointer(&out))
			if !reflect.DeepEqual(out, tt.expected) {
				t.Errorf("IncUint32() out: %v expected: %v", out, tt.expected)
			}
		})
	}
}
