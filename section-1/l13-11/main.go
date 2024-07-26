package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		sum := 0

		for N > 0 {
			sum += N % 10
			N /= 10
		}

		N = sum
	}

	fmt.Println(N)
}
