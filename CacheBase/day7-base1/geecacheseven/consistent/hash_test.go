package consistent

import (
	"strconv"
	"testing"
)

func Test_Consistent(t *testing.T) {

	c := NewConsistent(5, func(data []byte) uint32 {
		n, _ := strconv.Atoi(string(data))
		return uint32(n)
	})

	c.Set("1", "2", "3")
	t.Log(c.keys)
	t.Log(c.hashMap)

	t.Log(c.Get("1"))
	t.Log(c.Get("2"))
	t.Log(c.Get("3"))
	t.Log()
	t.Log(c.Get("10"))
	t.Log(c.Get("20"))
	t.Log(c.Get("30"))
	t.Log(c.Get("40"))
	t.Log(c.Get("50"))
	t.Log(c.Get("60"))
	t.Log()
	t.Log(c.Get("111"))
	t.Log(c.Get("222"))
	t.Log(c.Get("333"))

	t.Log()
	c.Set("4", "5", "6", "7", "8", "9")
	t.Log(c.keys)
	t.Log(c.hashMap)
	t.Log()
	t.Log(c.Get("1"))
	t.Log(c.Get("2"))
	t.Log(c.Get("3"))
	t.Log()
	t.Log(c.Get("10"))
	t.Log(c.Get("20"))
	t.Log(c.Get("30"))
	t.Log(c.Get("40"))
	t.Log(c.Get("50"))
	t.Log(c.Get("60"))
	t.Log()
	t.Log(c.Get("111"))
	t.Log(c.Get("222"))
	t.Log(c.Get("333"))

}
