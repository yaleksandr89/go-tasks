package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	// Начальная степень двойки
	power := 1

	for power <= N {
		fmt.Printf("%d ", power)
		power *= 2
	}

	fmt.Println()
}
