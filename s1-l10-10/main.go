package main

import "fmt"

func main() {
	var num1, num2 string
	fmt.Scan(&num1, &num2)

	// Перебираем все цифры первого числа
	for i := 0; i < len(num1); i++ {
		digit1 := num1[i]

		// Проверяем, если цифра присутствует во втором числе
		for j := 0; j < len(num2); j++ {
			if digit1 == num2[j] {
				// Выводим цифру только если она еще не была выведена
				isPrinted := false
				for k := 0; k < i; k++ {
					if num1[k] == digit1 {
						isPrinted = true
						break
					}
				}
				if !isPrinted {
					if i > 0 {
						fmt.Print(" ")
					}
					fmt.Print(string(digit1))
				}
				break
			}
		}
	}
	fmt.Println()
}
