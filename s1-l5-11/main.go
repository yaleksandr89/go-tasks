// Ключевые слова: break, case, chan, const, continue, default, defer, else, fallthrough, for, func, go,
//                 goto, if, import, interface, map, package, range, return, select, struct, switch, type, var

package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)
	a = a*2 + 100

	fmt.Println(a)
}
