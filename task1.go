package main

import "fmt"

func main() {
	number := 5

	if number > 0 {
		fmt.Println("Число положительное")
	} else if number < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Ноль")
	}
}