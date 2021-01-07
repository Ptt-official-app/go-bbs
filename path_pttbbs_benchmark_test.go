package bbs

import (
	"testing"
)

func BenchmarkGetUserFavoritePath(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetUserFavoritePath("/home/bbs", "SYSOP")
	}
}
