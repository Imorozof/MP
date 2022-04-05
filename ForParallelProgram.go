package main

import "fmt"

func main() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)
	for result := c; (result%d) == 0 && result <= n; result += c {
	}
	if result%d != 0 {
		fmt.Println(result)
	}
}
