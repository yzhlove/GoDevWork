package main

import (
	"testing"
)

func Test_LFU(t *testing.T) {

	LFU := New(3)
	LFU.Set("1", 1)
	LFU.Set("2", 2)
	LFU.Set("3", 3)
	t.Log("=================================")
	LFU.ShowValues()
	LFU.Set("4", 4)
	t.Log("=================================")
	LFU.ShowValues()
	LFU.Get("3")
	LFU.Get("3")
	t.Log("=================================")
	LFU.ShowValues()
	//LFU.ShowBuckets()
	LFU.Get("2")
	//t.Log("=================================")
	//LFU.ShowBuckets()
	LFU.Set("5", 5)
	t.Log("=================================")
	LFU.ShowValues()
	LFU.Set("6", 6)
	t.Log("=================================")
	LFU.ShowValues()
}
