package main

import "fmt"

func main() {
	var (
		num      int
		MinValue = 10
		MaxValue = 100
	)

	for {
		fmt.Scan(&num)

		switch {
		case num < MinValue:
			continue
		case num > MaxValue:
			return
		default:
			fmt.Println(num)
		}
	}
}
