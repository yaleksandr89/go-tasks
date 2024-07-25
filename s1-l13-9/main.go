package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	minValue := math.MaxInt
	countMix := 0

	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)

		if num < minValue {
			minValue = num
			countMix = 1
		} else if num == minValue {
			countMix++
		}
	}

	fmt.Println(countMix)
}
