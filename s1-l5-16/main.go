package main

import "fmt"

// В часах часовая стрелка делает полный оборот (360 градусов) за 12 часов. Это означает, что за один час стрелка поворачивается на 360 / 12 = 30 градусов.

func main() {
	var d int
	fmt.Scan(&d) // Считываем угол поворота в градусах

	h := d / 30                       // Количество целей часов
	remainingDegrees := d % 30        // Оставшиеся градусы
	m := (remainingDegrees * 60) / 30 // Количество целых минут

	fmt.Printf("It is %d hours %d minutes.\n", h, m)
}
