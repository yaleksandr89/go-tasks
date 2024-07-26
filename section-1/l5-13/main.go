//+= a // тоже самое, что и a = a + a
//a -= a // тоже самое, что и a = a - a
//a /= a // тоже самое, что и a = a / a
//a *= a // тоже самое, что и a = a * a

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	a *= a
	fmt.Println(a)
}
