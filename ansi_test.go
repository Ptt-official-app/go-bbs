package bbs

import (
	"strings"
	"testing"
)

const strWithANSI = "\x1b[1;31m→ \x1b[33mpichu2\x1b[m\x1b[33m:推"

func TestFilterStringANSI(t *testing.T) {
	src := strWithANSI
	expected := "→ pichu2:推"

	dst := FilterStringANSI(src)

	if strings.Compare(expected, dst) != 0 {
		t.Errorf("FilterStringANSI doesn't filter ANSI CSI code, \nexpected: \n%s, \ngot: \n%s", expected, dst)
	}
}

func BenchmarkStringANSI(b *testing.B) {
	src := strWithANSI

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = FilterStringANSI(src)
	}
}
