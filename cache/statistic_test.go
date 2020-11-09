package cache

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/PichuChen/go-bbs/ptttype"
	"github.com/PichuChen/go-bbs/types"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func TestStatInc(t *testing.T) {
	setupTest()
	defer teardownTest()

	err := NewSHM(testShmKey, false, true)
	if err != nil {
		log.Errorf("TestStatInc: unable to NewSHM: e: %v", err)
		return
	}
	defer CloseSHM()

	type args struct {
		stats uintptr
	}
	tests := []struct {
		name     string
		args     args
		expected uint32
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{stats: ptttype.STAT_BOARDREC},
			expected: 1,
		},
		{
			args:     args{stats: ptttype.STAT_BOARDREC},
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StatInc(tt.args.stats); (err != nil) != tt.wantErr {
				t.Errorf("StatInc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		out := uint32(0)
		Shm.ReadAt(
			unsafe.Offsetof(Shm.Statistic)+types.UINT32_SZ*ptttype.STAT_BOARDREC,
			unsafe.Sizeof(Shm.Statistic[ptttype.STAT_BOARDREC]),
			unsafe.Pointer(&out),
		)

		if !reflect.DeepEqual(out, tt.expected) {
			t.Errorf("StatInc() out: %v expected: %v", out, tt.expected)
		}
	}
}

func TestReadStat(t *testing.T) {
	setupTest()
	defer teardownTest()

	err := NewSHM(testShmKey, false, true)
	if err != nil {
		log.Errorf("TestReadStat: unable to NewSHM: e: %v", err)
		return
	}
	defer CloseSHM()

	CleanStat()

	_ = StatInc(ptttype.STAT_BOARDREC)
	_ = StatInc(ptttype.STAT_BOARDREC)
	_ = StatInc(ptttype.STAT_GAMBLE)

	type args struct {
		stats uintptr
	}
	tests := []struct {
		name     string
		args     args
		expected uint32
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			args:     args{ptttype.STAT_BOARDREC},
			expected: 2,
		},
		{
			args:     args{ptttype.STAT_BOARDREC},
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadStat(tt.args.stats)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadStat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("ReadStat() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestCleanStat(t *testing.T) {
	setupTest()
	defer teardownTest()

	err := NewSHM(testShmKey, false, true)
	if err != nil {
		logrus.Errorf("CleanStat: unable to new shm: e: %v", err)
		return
	}
	defer CloseSHM()

	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CleanStat()

			statsBoardRec, _ := ReadStat(ptttype.STAT_BOARDREC)
			if statsBoardRec != 0 {
				t.Errorf("CleanStat() statsBoardRec: %v", statsBoardRec)
			}
		})
	}
}
