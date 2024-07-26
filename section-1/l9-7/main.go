package main

import (
	"fmt"
)

func main() {
	var intVal int
	fmt.Scan(&intVal)

	// Проверка, что число трехзначное
	if intVal < 0 || intVal > 10000 {
		fmt.Println("Ошибка: введите число от 0 до 10000.")
		return
	}

	// Делим введенное число, пока оно не станет меньше 10
	for intVal >= 10 {
		intVal /= 10
	}

	fmt.Println(intVal)
}
