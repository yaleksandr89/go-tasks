package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	found := false
	maxMultiple := 0

	for i := b; i >= a; i-- {
		if i%7 == 0 {
			maxMultiple = i
			found = true
			break
		}
	}

	if found {
		fmt.Println(maxMultiple)
	} else {
		fmt.Println("NO")
	}
}
