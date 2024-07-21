package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Scan(&num)

	if num <= 0 {
		fmt.Printf("число %.2f не подходит\n", num)
	} else if num > 100 {
		fmt.Printf("%e\n", num)
	} else {
		squared := math.Pow(num, 2)
		fmt.Printf("%.4f\n", squared)
	}
}
