package main

import (
	"fmt"
)

func main() {
	var intVal int
	fmt.Scan(&intVal)

	a := intVal / 100       // Извлекаем первую цифру
	b := (intVal / 10) % 10 // Извлекаем вторую цифру
	c := intVal % 10        // Извлекаем первую цифру

	// Проверка, что число трехзначное
	if intVal < 100 || intVal > 999 {
		fmt.Println("Ошибка: введите трехзначное число.")
		return
	}

	// Проверяем все ли числа различны
	if a != b && b != c && a != c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
