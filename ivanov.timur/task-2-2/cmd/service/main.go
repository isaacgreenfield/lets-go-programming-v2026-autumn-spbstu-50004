package main

import (
	"fmt"

	"task-2-2/internal/intheap"
)

func main() {
	var amount int

	_, err := fmt.Scan(&amount)
	if err != nil {
		fmt.Println(err)

		return
	}

	intHeap := intheap.IntHeap{}

	for range amount {
		var value int

		_, err := fmt.Scan(&value)
		if err != nil {
			fmt.Println(err)

			return
		}

		if err := intHeap.Push(value); err != nil {
			fmt.Println(err)

			return
		}
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
