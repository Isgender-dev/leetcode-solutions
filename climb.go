package main

import (
	"fmt"
)

func main() {
	for {
		var x int
		fmt.Printf("Input:  ")
		fmt.Scan(&x)
		if x < 0 {
			fmt.Printf(" ERROR!!! \n")
			break
		}
		fmt.Printf("Output: %v\n", climb(x))
	}
}
func climb(x int) int {
	switch x {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	default:
		return climb(x-1) + climb(x-2)
	}
}
