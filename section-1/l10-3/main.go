package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	if A < 1 || A > 100 || B < 1 || B > 100 {
		fmt.Println("Ошибка: убедитесь, что 1 <= A < B <= 100.")
		return
	}

	sum := 0

	for i := A; i <= B; i++ {
		sum += i
	}

	fmt.Println(sum)
}
