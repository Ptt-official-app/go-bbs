package bbs

import (
	"bytes"
	"strings"
	"testing"
)

func TestFilterANSI(t *testing.T) {
	src := []byte("[1;31m→ [33mpichu2[m[33m:推")
	dst := make([]byte, len(src))
	expected := []byte("→ pichu2:推")

	dst = FilterANSI(dst, src)

	if bytes.Compare(expected, dst) != 0 {
		t.Errorf("FilterANSI doesn't filter the ANSI code, \nexpected: \n%s, \ngot: \n%s", expected, dst)
	}
}

func TestFilterStringANSI(t *testing.T) {
	src := "[1;31m→ [33mpichu2[m[33m:推"
	expected := "→ pichu2:推"

	dst := FilterStringANSI(src)

	if strings.Compare(expected, dst) != 0 {
		t.Errorf("FilterStringANSI doesn't filter ANSI CSI code, \nexpected: \n%s, \ngot: \n%s", expected, dst)
	}
}

func BenchmarkReflect(b *testing.B) {
	src := []byte("[1;31m→ [33mpichu2[m[33m:推")
	dst := make([]byte, len(src))

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		FilterANSI(dst, src)
	}
}
