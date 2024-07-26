package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	if N == 0 {
		fmt.Println(0)
		return
	}

	binaryRepresentation := ""

	for N > 0 {
		remainder := N % 2
		binaryRepresentation = fmt.Sprintf("%d", remainder) + binaryRepresentation
		N = N / 2
	}

	fmt.Println(binaryRepresentation)
}
