package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print("numRows = ")
	fmt.Scan(&x)

	n := [][]int{
		{1},
		{1, 1},
	}
	if x >= 1 {
		fmt.Println(n[:1])
	}
	if x >= 2 {
		fmt.Println(n[:2])
		for len(n) < x {
			a := []int{1}
			z := n[len(n)-1]
			//for j:= len(n);
			for i := 0; i < len(z)-1; i++ {
				sum := z[i] + z[i+1]
				a = append(a, sum)
			}
			a = append(a, 1)
			n = append(n, a)
			fmt.Println(n)
		}
	}
}
