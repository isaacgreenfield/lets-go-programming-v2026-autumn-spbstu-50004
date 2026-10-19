package main

import "fmt"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Sort() {
	for i := 0; i < len(*h)-1; i++ {
		for j := 0; j < len(*h)-1-i; j++ {
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

	h := IntHeap{}
	for range amount {
		var tmp int

		_, err := fmt.Scan(&tmp)
		if err != nil {
			fmt.Println(err)

			return
		}

		h.Push(tmp)
	}

	_, err = fmt.Scan(&amount)
	if err != nil {
		fmt.Println(err)

		return
	}

	h.Sort()
	if amount > h.Len() {
		fmt.Println(-1)

		return
	} else {
		fmt.Println(h[amount-1])
	}
}
