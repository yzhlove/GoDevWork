package main

import (
	"encoding/json"
	"fmt"
)

//go:generate msgp -io=false -tests=false

type Student struct {
	ID       uint32
	Name     string
	Work     string
	Age      uint16
	Birthday string
}

type Teacher struct {
	ID       uint32 `msg:"id,omitempty" json:"id,omitempty"`
	Name     string `msg:"name,omitempty" json:"name,omitempty"`
	Work     string `msg:"work,omitempty" json:"work,omitempty"`
	Age      uint16 `msg:"age,omitempty" json:"age,omitempty"`
	Birthday string `msg:"birthday,omitempty" json:"birthday,omitempty"`
}

func main() {

	s := Student{}
	sdata, err := s.MarshalMsg(nil)
	if err != nil {
		panic(err)
	}
	jsonsdata, _ := json.Marshal(s)
	fmt.Println("sdata size -> ", len(sdata), "json size -> ", len(jsonsdata), " sdata => ", string(sdata), sdata)

	t := Teacher{}
	tdata, err := t.MarshalMsg(nil)
	if err != nil {
		panic(err)
	}
	jsontdata, _ := json.Marshal(t)
	fmt.Println("tdata size -> ", len(tdata), " json size -> ", len(jsontdata), " tdata => ", string(tdata), tdata)
}
