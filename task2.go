package main

import "fmt"

func main() {
	number := 111

	first := number / 100
	second := (number / 10) % 10
	third := number % 10

	if first != second && first != third && second != third {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}