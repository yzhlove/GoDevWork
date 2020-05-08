package signal

import "testing"

func Test_Do(t *testing.T) {

	var g Group
	v, err := g.Do("key", func() (interface{}, error) {
		return "bar", nil
	})

	t.Log(v, " = ", err)

}
