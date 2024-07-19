package main

import (
	"fmt"
)

func main() {
	var ticketNumber int
	fmt.Scan(&ticketNumber)

	// Проверка, что число трехзначное
	if ticketNumber < 100000 || ticketNumber > 999999 {
		fmt.Println("Ошибка: введите шестизначный номер билета.")
		return
	}

	firstDigit := ticketNumber / 100000
	secondDigit := (ticketNumber / 10000) % 10
	thirdDigit := (ticketNumber / 1000) % 10
	fourthDigit := (ticketNumber / 100) % 10
	fifthDigit := (ticketNumber / 10) % 10
	sixthDigit := ticketNumber % 10

	sumFirstThree := firstDigit + secondDigit + thirdDigit
	sumLastThree := fourthDigit + fifthDigit + sixthDigit

	if sumFirstThree == sumLastThree {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
