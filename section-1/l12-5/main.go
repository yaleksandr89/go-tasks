package main

import "fmt"

func main() {
	// Создаем массив из 10 элементов типа uint8
	var workArray [10]uint8

	// Считываем 10 целых положительных чисел
	for i := 0; i < 10; i++ {
		var num uint8
		fmt.Scan(&num)

		workArray[i] = num
	}

	// Считываем 3 пары индексов и меняем соответствующие элементы местами
	for i := 0; i < 3; i++ {
		var index1, index2 int
		fmt.Scan(&index1, &index2)

		workArray[index1], workArray[index2] = workArray[index2], workArray[index1]
	}

	// Выводим элементы массива через пробел
	for i := 0; i < 10; i++ {
		fmt.Print(workArray[i], " ")
	}
}
