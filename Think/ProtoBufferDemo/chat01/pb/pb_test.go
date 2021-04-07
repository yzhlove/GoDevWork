package pb

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
)

func Test_Encode(t *testing.T) {

	user := &User{Id: 150, UID: 150}
	t.Log(user)
	data, err := proto.Marshal(user)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
	t.Log(fmt.Sprintf("%x\n", data))
}

func Test_Stu(t *testing.T) {

	s := Stu{Id: 150}
	t.Log(s)

}
