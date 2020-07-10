package util

import "testing"

func BenchmarkGetUID(b *testing.B) {
	repeated := make(map[string]struct{}, 1e6)
	for i := 0; i < b.N; i++ {
		UID := GetUID()
		if _, ok := repeated[UID]; ok {
			b.Error("repeated")
		} else {
			repeated[UID] = struct{}{}
		}
	}
}
