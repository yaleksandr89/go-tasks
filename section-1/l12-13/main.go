package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	// Создаем срез с длиной N
	numbers := make([]int, N)

	// Считываем N целых чисел и сохраняем их в срез
	for i := 0; i < N; i++ {
		fmt.Scan(&numbers[i])
	}

	// Выводим 4-й элемент (элемент с индексом 3)
	fmt.Println(numbers[3])
}
