package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)

	if num < 100 || num > 999 || num%10 == 0 {
		return
	}

	hundreds := num / 100
	tens := (num / 10) % 10
	units := num % 10

	reversed := (units * 100) + (tens * 10) + hundreds

	fmt.Println(reversed)
}
