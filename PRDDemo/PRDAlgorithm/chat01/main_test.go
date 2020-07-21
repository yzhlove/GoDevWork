package main

import "testing"

func TestP2C(t *testing.T) {
	for i := 1; i <= 100; i++ {
		c := float64(i) / float64(100)
		rs := P2C(c)
		t.Logf("概率:%.2f%% \tC值:%.15f\n", c*100, rs)
	}
}
