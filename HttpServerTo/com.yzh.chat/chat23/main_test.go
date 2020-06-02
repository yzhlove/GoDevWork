package main

import "testing"

func Test_GetID(t *testing.T) {

	a := GetID()
	t.Log(a)
	b := GetID()
	t.Log(b)

}

func Benchmark_GetID(t *testing.B) {
	repeated := make(map[int64]struct{}, 1024)
	for i := 0; i < t.N; i++ {
		id := GetID()
		if _, ok := repeated[id]; ok {
			t.Error("repeated .", id, " length = ", len(repeated))
		} else {
			repeated[id] = struct{}{}
		}
	}
	t.Log("ok.")
}
