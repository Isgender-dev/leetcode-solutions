package main

import "fmt"

func main() {
	for i := 1; i <= 30; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Printf("Fizz, ")
		} else if i%5 == 0 && i%3 != 0 {
			fmt.Printf("Buzz, ")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Printf("FizzBuzz, ")
		} else {
			fmt.Printf("%v, ", i)
		}
	}
}
