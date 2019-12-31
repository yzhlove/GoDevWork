package storage

import (
	"fmt"
	"testing"
)

func TestStorage_LoadAuth(t *testing.T) {

	auths := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

	for i, v := range auths[3:] {
		fmt.Println(i, " - ", v)
	}

}
