package main

import (
	"fmt"
)

func main() {
	var seconds int
	fmt.Scan(&seconds)

	secondsInHours := 3600
	minutesInHours := 60

	hours := seconds / secondsInHours
	minutes := (seconds % secondsInHours) / minutesInHours

	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}
