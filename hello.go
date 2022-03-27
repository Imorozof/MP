package main

import "fmt"

func main() {
	var (
		i, a    int
		b, c, d bool
	)
	a = 0
	fmt.Scan(&i)
	b = i%400 == a
	c = i%4 == a
	d = i%100 != a
	switch {
	case b == true || (c == true && d == true):
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
