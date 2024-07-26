package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	var ending string
	switch {
	case N%10 == 1 && N%100 != 11:
		ending = "korova"
		break
	case (N%10 >= 2 && N%10 <= 4) && (N%100 < 10 || N%100 >= 20):
		ending = "korovy"
		break
	default:
		ending = "korov"
	}

	fmt.Printf("%d %s\n", N, ending)
}
