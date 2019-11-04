package main

import "fmt"

type Depot struct {
	Id  uint32
	Num uint32
}

func (dpt *Depot) MulRef(num uint32) {
	dpt.Num -= num
}

func (dt Depot) Mul(num uint32) {
	dt.Num -= num
}

func showDeport(deports []Depot) {
	for _, deport := range deports {
		fmt.Printf("id => %d num:%d \n", deport.Id, deport.Num)
	}
	fmt.Println()
}

func showRefDeport(deports []*Depot) {
	for _, deport := range deports {
		fmt.Printf("id => %d num:%d \n", deport.Id, deport.Num)
	}
	fmt.Println()
}

func main() {

	dts := []Depot{{Id: 1, Num: 10}, {Id: 2, Num: 20}}
	dpts := []*Depot{{Id: 1, Num: 10}, {Id: 2, Num: 20}}

	for _, d := range dts {
		d.Mul(1)
	}
	showDeport(dts)

	for _, d := range dts {
		d.MulRef(1)
	}
	showDeport(dts)

	for _, d := range dpts {
		d.Mul(2)
	}
	showRefDeport(dpts)

	for _, d := range dpts {
		d.MulRef(2)
	}
	showRefDeport(dpts)

	tds := Depot{Id: 3, Num: 10}
	tds.Mul(2)
	fmt.Println("tds =>", tds)
	tds.MulRef(3)
	fmt.Println("tds =>", tds)

	fmt.Println()

	ptds := &Depot{Id: 3, Num: 10}
	ptds.Mul(2)
	fmt.Println("tds =>", ptds)
	ptds.MulRef(3)
	fmt.Println("tds =>", ptds)

}
