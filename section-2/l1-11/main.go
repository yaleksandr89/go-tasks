package main

import "fmt"

func sumInt(nums ...int) (int, int) {
	count := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return count, sum
}

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b) // Output: 2, 1

	// Пример с дополнительными числами
	c, d := sumInt(1, 2, 3, 4, 5)
	fmt.Println(c, d) // Output: 5, 15
}
