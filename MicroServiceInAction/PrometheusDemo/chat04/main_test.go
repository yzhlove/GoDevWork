package main

import (
	"net/http"
	"testing"
	"time"
)

func Test_Handle(t *testing.T) {

	for i := 0; i < 1e6; i++ {
		if _, err := http.Get("http://localhost:1234/handle"); err != nil {
			t.Error(err)
			continue
		}
		time.Sleep(50 * time.Millisecond)
	}

}
