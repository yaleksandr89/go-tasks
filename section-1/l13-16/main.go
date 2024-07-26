package main

import (
	"fmt"
)

func main() {
	var number, digit int
	fmt.Scan(&number, &digit)

	var result int
	multiplier := 1

	for number > 0 {
		currentDigit := number % 10
		number /= 10

		if currentDigit != digit {
			result += currentDigit * multiplier
			multiplier *= 10
		}
	}

	fmt.Println(result)
}
