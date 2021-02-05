package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"strings"
)

func main() {
	writeGob()
	readGob()
}

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Address   []*Address
	Remark    string
}

func (a *Address) String() string {
	return fmt.Sprintf("[Address] Type:%s <City:%s Country:%s>", a.Type, a.City, a.Country)
}

func (v *VCard) String() string {
	var buf strings.Builder
	buf.WriteString(fmt.Sprintf("[VCard] FirstName:%s LastName:%s ", v.FirstName, v.LastName))
	for _, address := range v.Address {
		buf.WriteString(fmt.Sprintf(" <Address:%s> ", address))
	}
	buf.WriteString(fmt.Sprintf(" Remark:%s", v.Remark))
	return buf.String()
}

func readGob() {

	f, err := os.Open("vcard.gob")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cards []*VCard
	if decoder := gob.NewDecoder(f); decoder != nil {
		if err := decoder.Decode(&cards); err != nil {
			panic(err)
		}
	}

	for _, card := range cards {
		fmt.Println("card => ", card)
	}
}

func writeGob() {

	f, err := os.Create("vcard.gob")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var buf bytes.Buffer
	if encoder := gob.NewEncoder(&buf); encoder != nil {
		if err := encoder.Encode(creatorGobData()); err != nil {
			panic(err)
		}
	}

	if _, err := f.Write(buf.Bytes()); err != nil {
		panic(err)
	}
}

func creatorGobData() []*VCard {
	cards := make([]*VCard, 0, 2)
	cards = append(cards, &VCard{FirstName: "stack", LastName: "tony", Remark: "nornal", Address: []*Address{
		&Address{Type: "CA", City: "Shanghai", Country: "Chain"},
		&Address{Type: "CA", City: "Beijing", Country: "Chain"},
		&Address{Type: "CA", City: "Xiamen", Country: "Chain"},
	}})
	cards = append(cards, &VCard{FirstName: "stack", LastName: "tony", Remark: "nornal", Address: []*Address{
		&Address{Type: "CF", City: "AAAA", Country: "USA"},
		&Address{Type: "CF", City: "BBBB", Country: "Chain"},
		&Address{Type: "CF", City: "CCCC", Country: "Chain"},
	}})
	return cards
}
