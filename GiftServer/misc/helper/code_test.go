package helper

import (
	"testing"
)

func Test_EncodeAndDecode(t *testing.T) {

	id := 1
	num := int64(223344)

	code := Encode(uint32(id), num)
	t.Log("code => ", code)

	old, ok := Decode(code)
	t.Log("id => ", old, " ok => ", ok)

	if !ok {
		t.Error(ok)
	}

}
