package main

import (
	"fmt"
)

func main() {
	var number int

	for i := 1; i <= 10; i++ {
		number = i * i
		fmt.Println(number)
	}
}
