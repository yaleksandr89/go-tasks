package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)

	hundreds := num / 100
	tens := (num / 10) % 10
	units := num % 10

	sum := hundreds + tens + units

	fmt.Println(sum)
}
