package main

import (
	"fmt"
)

func main() {
	var num, maxNum, count int

	for {
		fmt.Scan(&num)

		if num == 0 {
			break
		}

		if num > maxNum {
			maxNum = num
			count = 1
		} else if num == maxNum {
			count++
		}
	}

	fmt.Println(count)
}
