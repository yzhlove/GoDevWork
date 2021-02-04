package main

import "fmt"

func main() {
	//a := []int{1,2,3,4}
	b := []int{4, 8, 12, 16}
	setup(b)
}

// number {1,2,3,4}
func setup(probabilities []int) {

	n := len(probabilities)

	avg := 10
	l, m := 0, n-1

	workList := make([]int, n)
	for i, prob := range probabilities {
		if prob > avg {
			workList[m] = i
			m--
		} else {
			workList[l] = i
			l++
		}
	}

	for _, v := range workList {
		fmt.Print(v, probabilities[v], "\t")
	}
	fmt.Println("\n---------------------")

	prob := make([]int, n)
	copy(prob, probabilities)

	fmt.Println("prob=>", prob)
	fmt.Println("work=>", workList)

	for l != 0 && m != n-1 {
		less, more := workList[l-1], workList[m+1]
		fmt.Println("l=", l-1, " m=", m+1, "prob[less]=", prob[less], " prob[more]=", prob[more])
		prob[more] = prob[more] + prob[less] - avg
		l--
		if prob[more] < avg {
			workList[l] = more
			l++
			m++
		}
	}

	fmt.Println("prob => ", prob)

}
