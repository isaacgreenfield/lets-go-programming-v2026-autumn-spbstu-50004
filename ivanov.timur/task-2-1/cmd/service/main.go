package main

import (
	"fmt"
)

func main() {
	var (
		depCount  int
		workCount int
		maxTemp   int8 = 30
		minTemp   int8 = 15
		temp      int8
	)

	_, err := fmt.Scan(&depCount)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	for range depCount {
		_, err = fmt.Scan(&workCount)
		if err != nil {
			fmt.Println("Error:", err)

			return
		}

		for ; workCount > 0; workCount-- {
			var pointer string

			_, err := fmt.Scan(&pointer, &temp)
			if err != nil {
				fmt.Println("Error:", err)

				return
			}

			switch pointer {
			case ">=":
				if temp > minTemp {
					minTemp = temp
				}
			case "<=":
				if temp < maxTemp {
					maxTemp = temp
				}
			default:
				fmt.Println("Wrong input")

				return
			}

			if minTemp > maxTemp {
				fmt.Println(-1)
			} else {
				fmt.Println(minTemp)
			}
		}
	}
}
