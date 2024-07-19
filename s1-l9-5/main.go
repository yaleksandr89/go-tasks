package main

import (
	"fmt"
)

func main() {
	var intVal int
	fmt.Scan(&intVal)

	switch {
	case intVal > 0:
		fmt.Println("Число положительное")
	case intVal < 0:
		fmt.Println("Число отрицательное")
	case intVal == 0:
		fmt.Println("Ноль")
	}
}
