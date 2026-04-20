package main

import (
	"fmt"
)

func maxProfit(price []int) int {
	minV := price[0]
	maxProfit := 0

	for i := 0; i < len(price); i++ {
		if price[i] < minV {
			minV = price[i]
		} else {
			profit := price[i] - minV
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

func main() {
	fmt.Println("Output: ", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println()
	fmt.Println("Output: ", maxProfit([]int{7, 6, 4, 3, 1, 1}))
	fmt.Println()
	fmt.Println("Output: ", maxProfit([]int{7, 7, 7, 7, 7, 7}))
	fmt.Println()
	fmt.Println("Output: ", maxProfit([]int{5, 2, 6, 1, 4, 3}))
}
