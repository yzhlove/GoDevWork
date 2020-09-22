package dh

import "testing"

func TestDH(t *testing.T) {

	X1, E1 := DHExchange()
	X2, E2 := DHExchange()

	t.Log(X1, E1)
	t.Log(X2, E2)

	KEY1 := DHKey(X1, E2)
	KEY2 := DHKey(X2, E1)

	t.Log(KEY1, ",", KEY2)

	if KEY1.Cmp(KEY2) != 0 {
		t.Error("dh error")
	}
}

func BenchmarkDH(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X1, E1 := DHExchange()
		X2, E2 := DHExchange()
		DHKey(X1, E2)
		DHKey(X2, E1)
	}
}
