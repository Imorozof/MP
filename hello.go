package main

import "fmt"

func main() {
	var newNumber int8
	for fmt.Scan(&newNumber); newNumber > 0; fmt.Scan(&newNumber) {
		if newNumber < 10 {
			continue
		} else if newNumber >= 10 && newNumber <= 100 {
			fmt.Println(newNumber)
		} else if newNumber > 100 {
			break
		}
	}
}
