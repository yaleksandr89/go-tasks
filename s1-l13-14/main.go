package main

import (
	"fmt"
)

func main() {
	var A int
	fmt.Scan(&A)

	if A <= 1 {
		fmt.Println(-1)
		return
	}

	fib1, fib2 := 1, 1
	index := 2

	for fib2 < A {
		fib1, fib2 = fib2, fib1+fib2
		index++
	}

	if fib2 == A {
		fmt.Println(index)
	} else {
		fmt.Println(-1)
	}
}
