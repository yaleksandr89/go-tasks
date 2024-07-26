package main

import (
	"fmt"
)

func main() {
	// Объявляем и инициализируем массив из 5 элементов
	var numbers [5]int
	for i := 0; i < 5; i++ {
		fmt.Scan(&numbers[i])
	}

	// Инициализируем переменную max первым элементом массива
	max := numbers[0]

	// Ищем максимальное значение в массиве
	for i := 1; i < 5; i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}

	// Выводим максимальное значение
	fmt.Println(max)
}
