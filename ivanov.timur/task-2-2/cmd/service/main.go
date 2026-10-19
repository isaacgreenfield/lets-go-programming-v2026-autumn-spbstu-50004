package main

import "fmt"

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	value, ok := x.(int)
	if !ok {
		panic("IntHeap.Push: value is not int")
	}

	*h = append(*h, value)
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func (h *IntHeap) Sort() {
	length := len(*h)
	for i := range length - 1 {
		for j := range length - 1 - i {
			if (*h)[j] > (*h)[j+1] {
				h.Swap(j, j+1)
			}
		}
	}
}

func main() {
	var amount int

	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println(err)

		return
	}

	intHeap := IntHeap{}

	for range amount {
		var value int

		_, err := fmt.Scan(&value)
		if err != nil {
			fmt.Println(err)

			return
		}

		intHeap.Push(value)
	}

	var position int

	_, err = fmt.Scan(&position)
	if err != nil {
		fmt.Println(err)

		return
	}

	intHeap.Sort()

	if position <= 0 || position > intHeap.Len() {
		fmt.Println(-1)

		return
	}

	fmt.Println(intHeap[intHeap.Len()-position])
}
