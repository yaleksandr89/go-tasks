package main

import (
	"fmt"
)

func main() {
	var year int
	fmt.Scan(&year)

	if year < 0 || year > 10000 {
		fmt.Println("Ошибка: введите положительное, не превышающее 10000 число.")
		return
	}

	if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
