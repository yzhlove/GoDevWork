package main

import "fmt"

func main() {

	var count int
	tids := []int{0}
	update := func() {
		if count >= 5 {
			return
		}
		count++
		tids = append(tids, count)
		fmt.Println("len(tids) => ", len(tids))
	}

	for i := 0; i < len(tids); i++ {
		fmt.Println("i => ", i, " v=> ", tids[i])
		update()
	}

}
