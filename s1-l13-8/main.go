package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	countZeroes := 0

	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)

		if num == 0 {
			countZeroes++
		}
	}

	fmt.Println(countZeroes)
}
