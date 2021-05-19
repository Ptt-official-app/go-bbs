package bbs

import (
	"strings"
	"testing"
)

func TestFilterStringANSI(t *testing.T) {
	src := "[1;31m→ [33mpichu2[m[33m:推"
	expected := "→ pichu2:推"

	dst := FilterStringANSI(src)

	if strings.Compare(expected, dst) != 0 {
		t.Errorf("FilterStringANSI doesn't filter ANSI CSI code, \nexpected: \n%s, \ngot: \n%s", expected, dst)
	}
}

func BenchmarkStringANSI(b *testing.B) {
	src := "[1;31m→ [33mpichu2[m[33m:推"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = FilterStringANSI(src)
	}
}
