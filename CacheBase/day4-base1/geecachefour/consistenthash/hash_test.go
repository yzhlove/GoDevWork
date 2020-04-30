package consistenthash

import (
	"fmt"
	"geecachefive/consistenthash"
	"strconv"
	"testing"
)

func Test_Hash(t *testing.T) {
	c := consistenthash.NewConsistentHash(3, func(data []byte) uint32 {
		i, _ := strconv.Atoi(string(data))
		return uint32(i)
	})
	c.Set("6", "4", "2")
	fmt.Println(c.keys)
	fmt.Println(c.hashMap)

	testCase := map[string]string{
		"2":  "2",
		"11": "2",
		"23": "4",
		"27": "2",
	}

	for k, v := range testCase {
		if c.Get(k) != v {
			t.Errorf("get err: %v ", k)
			return
		}
	}

	c.Set("8")
	fmt.Println(c.keys)
	fmt.Println(c.hashMap)

	t.Log(c.Get("27"))
}
