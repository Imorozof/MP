package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	i = i % 4
	fmt.Println(i)
}
