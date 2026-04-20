package main

import (
	"fmt"
	"strconv"
)

var num int

func mirror(Str string) string {
	if len(Str) == 0 {
		return ""
	}
	return mirror(Str[1:]) + Str[:1]
}

func mirrorDistance(n int) int {
	Str := strconv.Itoa(n)
	num, err := strconv.Atoi(mirror(Str))
	if err != nil {
		fmt.Println("ERROR:", err)
		return 0
	}
	return abs(num - n)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	for {
		fmt.Print("Enter the number: ")
		fmt.Scan(&num)
		fmt.Println("Result: ", mirrorDistance(num))

	}
}
