package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func main() {
	setup([]float64{0.1, 0.3, 0.2, 0.4})
}

type ele struct {
	prob float64
	idx  int
}

type Alias struct {
	table []ele
}

func (a *Alias) Pick() int {
	return a.PickWithRand(nil)
}

func (a *Alias) PickWithRand(r *rand.Rand) int {
	var rnd float64

	if r == nil {
		rnd = rand.Float64() * float64(len(a.table))
	} else {
		rnd = r.Float64() * float64(len(a.table))
	}

	cloumn := int(rnd)

	if rnd-float64(cloumn) < a.table[cloumn].prob {
		return cloumn
	}

	return a.table[cloumn].idx
}

func setup(probabilities []float64) ([]ele, error) {
	n := len(probabilities)
	if n == 0 {
		return nil, errors.New("probabilities must not be nil")
	}

	avg := 1.0 / float64(n)
	fmt.Println("avg => ", avg)
	l, m := 0, n-1

	workList := make([]int, n)
	//(0.1,0.3,0.2,0.4)
	for i, prob := range probabilities {
		if prob < 0 {
			return nil, errors.New("probability should not be negative")
		}
		if prob > avg {
			workList[m] = i
			m--
		} else {
			workList[l] = i
			l++
		}
	}

	fmt.Println("l = ", l, " m = ", m)

	for _, v := range workList {
		fmt.Println("workList real => ", probabilities[v])
	}

	fmt.Println("prob list ==> ", workList)

	eles := make([]ele, n)
	prob := make([]float64, n)
	copy(prob, probabilities)
	fmt.Println("====================\n")
	for l != 0 && m != n-1 {
		less, more := workList[l-1], workList[m+1]
		eles[less] = ele{prob: prob[less] * float64(n), idx: more}
		fmt.Println("ele => ", eles[less])
		fmt.Println("more ", more, " prob[more] ", prob[more], " less ", less, " prob[less]", prob[less])
		prob[more] = prob[more] + prob[less] - avg
		fmt.Println("after more => ", prob[more])
		l--

		if prob[more] < avg {
			fmt.Println("====> workList[l] = ", more, " l = ", l, " l++", l+1, " m++", m+1)
			workList[l] = more
			l++
			m++
		}
	}

	fmt.Println("workList => ", workList)
	fmt.Println("prob => ", prob)
	fmt.Println("prolist => ", probabilities)

	for ; l != 0; l-- {
		eles[workList[l-1]] = ele{prob: 1}
	}

	for ; m != n-1; m++ {
		eles[workList[m+1]] = ele{prob: 1}
	}

	fmt.Println("eles => ", eles)

	return nil, nil
}
