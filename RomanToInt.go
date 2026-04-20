func romanToInt(s string) int {
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var total int
	for i := 0; i < len(s); i++ {
		a := values[s[i]]
		if i+1 < len(s) && a < values[s[i+1]] {
			total -= a
		} else {
			total += a
		}
	}
	return total
}

func main() {
	for {
		var roman string
		fmt.Print("Enter the Roman number: ")
		fmt.Scan(&roman)
		fmt.Println("Result:", RomanToInt(roman))

	}
}
