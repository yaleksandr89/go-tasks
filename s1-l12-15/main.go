package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	// Создаем массив с N элементами
	numbers := make([]int, N)

	// Считываем элементы массива
	for i := 0; i < N; i++ {
		fmt.Scan(&numbers[i])
	}

	// Выводим элементы массива с четными индексами
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			fmt.Print(numbers[i], " ")
		}
	}
	fmt.Println()
}
