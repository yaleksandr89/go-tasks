package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	// Создаем массив для хранения элементов
	numbers := make([]int, N)

	// Считываем элементы массива
	for i := 0; i < N; i++ {
		fmt.Scan(&numbers[i])
	}

	// Подсчитываем количество положительных чисел
	count := 0
	for i := 0; i < N; i++ {
		if numbers[i] > 0 {
			count++
		}
	}

	// Выводим количество положительных чисел
	fmt.Println(count)
}
