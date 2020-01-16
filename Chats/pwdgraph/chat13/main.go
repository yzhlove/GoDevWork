package main

import "fmt"

var Dict = [...]string{
	"A", "B", "C", "D", "E", "F", "G", "H",
	"J", "K", "L", "M", "N", "P", "Q", "R",
	"S", "T", "U", "V", "W", "X", "Y", "Z",
	"2", "3", "4", "5", "6", "7", "8", "9",
}

func main() {

	var code uint64 = 1214113291710763023

	for {
		fmt.Print(Dict[code&uint64(0x1F)])
		code >>= 5
		if code == 0 {
			break
		}
	}

}
