package core

import "testing"

func BenchmarkGet(b *testing.B) {
	repeated := make(map[int64]struct{}, 1024)
	for i := 0; i < b.N; i++ {
		uuid := Get()
		if _, ok := repeated[uuid]; ok {
			b.Error("repeated.")
		} else {
			repeated[uuid] = struct{}{}
		}
	}
	b.Log("ok.")
}
