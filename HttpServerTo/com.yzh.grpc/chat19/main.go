package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {

	ps := &Person{
		Name:   "yzh",
		Age:    16,
		Emails: []string{"lcm_520@live.com", "lcmm5201314@gmail.com"},
		Phones: []*PhoneNumber{
			&PhoneNumber{
				Number: "12138",
				Type:   PhoneType_MOBILE,
			},
			&PhoneNumber{
				Number: "22138",
				Type:   PhoneType_HOME,
			},
			&PhoneNumber{
				Number: "32138",
				Type:   PhoneType_WORK,
			},
		},
	}

	data, err := proto.Marshal(ps)
	if err != nil {
		panic(err)
	}

	newPerson := &Person{}
	if err := proto.Unmarshal(data, newPerson); err != nil {
		panic(err)
	}
	fmt.Println("person ==> ", newPerson)
}
