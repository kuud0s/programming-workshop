package main

import "fmt"

func main() {
	ticket := 423126

	digit1 := ticket / 100000
	digit2 := (ticket / 10000) % 10
	digit3 := (ticket / 1000) % 10
	digit4 := (ticket / 100) % 10
	digit5 := (ticket / 10) % 10
	digit6 := ticket % 10

	sumFirst := digit1 + digit2 + digit3
	sumLast := digit4 + digit5 + digit6

	if sumFirst == sumLast {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
