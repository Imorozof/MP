package main

import "fmt"

func main() {

	var (
		a, b, c int
	)
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	c = a / 30
	b = 2 * (a % 30)

	fmt.Println("It is", c, "hours", b, "minutes.")
}
