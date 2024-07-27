package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	result := vote(x, y, z)
	fmt.Println(result)
}

func vote(x int, y int, z int) int {
	sum := x + y + z
	if sum >= 2 {
		return 1
	}
	return 0
}
