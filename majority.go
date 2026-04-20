package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	var count int
	var x int
	for _, v := range nums {
		if count == 0 {
			x = v
		}
		if v == x {
			count++
		} else {
			count--
		}
	}

	return x
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Printf("Most repeated number: %v \n", majorityElement(nums))
}
