package main

import (
	"fmt"
)

func main() {
	var depCount int

	_, err := fmt.Scan(&depCount)
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	for i := 0; i < depCount; i++ {
		var workCount int

		_, err := fmt.Scan(&workCount)
		if err != nil {
			fmt.Println("Error:", err)

			return
		}

		minTemp := 15
		maxTemp := 30

		for j := 0; j < workCount; j++ {
			var pointer string
			var temp int

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
				fmt.Println(maxTemp)
			}
		}
	}
}
