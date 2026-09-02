package main

import "fmt"

func main() {
	number := 12345

	for number >= 10 {
		number = number / 10
	}

	fmt.Println(number)
}