package main

import "fmt"

const PercentageBase = 100

func main() {
	// x — начальная сумма вклада.
	// p — процент ежегодного увеличения.
	// y — целевая сумма, которую нужно достичь.
	var x, p, y int
	fmt.Scan(&x, &p, &y)

	years := 0

	// Пока вклад меньше целевой суммы
	for x < y {
		// Рассчитываем процентное увеличение
		increase := (x * p) / PercentageBase
		// Увеличиваем сумму вклада
		x += increase
		// Увеличиваем счетчик лет
		years++
	}

	fmt.Println(years)
}
