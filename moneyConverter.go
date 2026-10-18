package main

import (
	"fmt"
	"strconv"
	"strings"
)

func rubToUsd(rub float64) float64 {
	return rub * 0.012
}

func rubToEur(rub float64) float64 {
	return rub * 0.011
}

func usdToRub(usd float64) float64 {
	return usd * 83.87
}

func eurToRub(eur float64) float64 {
	return eur * 93.83
}

func main() {
	var input string

	for {
		fmt.Println("\n1) RUB в USD")
		fmt.Println("2) RUB в EUR")
		fmt.Println("3) USD в RUB")
		fmt.Println("4) EUR в RUB")
		fmt.Print("Выбери номер: ")

		if _, err := fmt.Scan(&input); err != nil {
			return
		}

		choice, err := strconv.Atoi(input)

		if err != nil || choice < 1 || choice > 4 {
			fmt.Println("Ошибка: введи число от 1 до 4")
			continue
		}

		var amount float64

		for {
			fmt.Print("Введи сумму: ")

			if _, err := fmt.Scan(&input); err != nil {
				return
			}

			amount, err = strconv.ParseFloat(input, 64)

			if err == nil && amount >= 0 {
				break
			}

			fmt.Println("Ошибка: введи неотрицательное число")
		}

		if choice == 1 {
			fmt.Printf("%.2f RUB = %.2f USD\n", amount, rubToUsd(amount))
		} else if choice == 2 {
			fmt.Printf("%.2f RUB = %.2f EUR\n", amount, rubToEur(amount))
		} else if choice == 3 {
			fmt.Printf("%.2f USD = %.2f RUB\n", amount, usdToRub(amount))
		} else if choice == 4 {
			fmt.Printf("%.2f EUR = %.2f RUB\n", amount, eurToRub(amount))
		}

		for {
			fmt.Print("Сконвертировать ещё yes/no? ")

			if _, err := fmt.Scan(&input); err != nil {
				return
			}

			answer := strings.ToLower(input)

			if answer == "yes" || answer == "y" {
				break
			}

			if answer == "no" || answer == "n" {
				fmt.Println("Программа завершила работу")
				return
			}

			fmt.Println("Ошибка: введи yes или no")
		}
	}
}
