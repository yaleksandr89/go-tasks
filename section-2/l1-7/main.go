package main

import "fmt"

func main() {
	result := minimumFromFour()

	fmt.Println(result)
}

func minimumFromFour() int {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if a < b && a < c && a < d {
		return a
	} else if b < a && b < c && b < d {
		return b
	} else if c < a && c < b && c < d {
		return c
	} else {
		return d
	}
}
