package cache

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
)

func TestLoadUHash(t *testing.T) {
	setupTest()
	defer teardownTest()

	err := NewSHM(types.Key_t(ptttype.SHM_KEY), ptttype.USE_HUGETLB, true)
	if err != nil {
		return
	}
	defer CloseSHM()

	wantHashHead := [1 << ptttype.HASH_BITS]int32{}
	wantNextInHash := [ptttype.MAX_USERS]int32{}
	for idx := range wantHashHead {
		wantHashHead[idx] = -1
	}
	wantHashHead[29935] = 0 //SYSOP
	wantHashHead[56375] = 1 //CodingMan
	wantHashHead[36994] = 2 //pichu
	wantHashHead[15845] = 3 //Kahou
	wantHashHead[22901] = 4 //Kahou2
	wantHashHead[35] = 5    //""

	wantNextInHash[0] = -1
	wantNextInHash[1] = -1
	wantNextInHash[2] = -1
	wantNextInHash[3] = -1
	wantNextInHash[4] = -1
	for idx := 5; idx < 49; idx++ {
		wantNextInHash[idx] = int32(idx + 1)
	}
	wantNextInHash[49] = -1

	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error = nil
			if err = LoadUHash(); (err != nil) != tt.wantErr {
				t.Errorf("loadUHash() error = %v, wantErr %v", err, tt.wantErr)
			}

			hashHead := [1 << ptttype.HASH_BITS]int32{}
			nextInHash := [ptttype.MAX_USERS]int32{}

			Shm.ReadAt(
				unsafe.Offsetof(Shm.HashHead),
				unsafe.Sizeof(Shm.HashHead),
				unsafe.Pointer(&hashHead),
			)

			for idx, each := range hashHead {
				if each != wantHashHead[idx] {
					t.Errorf("loadUHash() (%v) hashHead: %v expected: %v", idx, each, wantHashHead[idx])
					break
				}
			}

			Shm.ReadAt(
				unsafe.Offsetof(Shm.NextInHash),
				unsafe.Sizeof(Shm.NextInHash),
				unsafe.Pointer(&nextInHash),
			)

			if !reflect.DeepEqual(nextInHash, wantNextInHash) {
				t.Errorf("loadUHash() nextInHash: %v expected: %v", nextInHash, wantNextInHash)

			}

		})
	}
}
