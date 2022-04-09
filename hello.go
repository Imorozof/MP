package main

import "fmt"

func main() {
	var x, p, y, n int
	fmt.Scan(&x, &p, &y)
	for n = 0; x <= y; n++ {
		x += x * p / 100
	}
	fmt.Println(n)
}
