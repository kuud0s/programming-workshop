package main

import (
	"fmt"
	"log"
	"math/rand" 
	"net/http"
	"strconv"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<h1>Главная</h1>")
		fmt.Fprintf(w, `<a href="/converter"><button>Конвертер валют</button></a> `)
		fmt.Fprintf(w, `<a href="/randomizer"><button>Генератор рандомных чисел</button></a>`)
	})

	http.HandleFunc("/converter", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		inputAmount := r.URL.Query().Get("amount")
		inputChoice := r.URL.Query().Get("choice")

		fmt.Fprint(w, `
			<h2>Конвертер валют</h2>
			<form action="/converter" method="GET">
				Введи сумму: <input type="text" name="amount" required><br><br>
				<input type="radio" name="choice" value="1" checked> RUB в USD<br>
				<input type="radio" name="choice" value="2"> RUB в EUR<br>
				<input type="radio" name="choice" value="3"> USD в RUB<br>
				<input type="radio" name="choice" value="4"> EUR в RUB<br><br>

				<button type="submit">Сконвертировать</button>
			</form><br>`
		)

		if inputAmount != "" && inputChoice != "" {
			choice, errChoice := strconv.Atoi(inputChoice)
			amount, errAmount := strconv.ParseFloat(inputAmount, 64)

			if errChoice != nil || choice < 1 || choice > 4 {
				fmt.Fprint(w, "Ошибка: введи число от 1 до 4")
				return
			}
			if errAmount != nil || amount < 0 {
				fmt.Fprint(w, "Ошибка: введи неотрицательное число")
				return
			}

			if choice == 1 {
				fmt.Fprintf(w, "<b>%.2f RUB = %.2f USD</b>\n", amount, rubToUsd(amount))
			} else if choice == 2 {
				fmt.Fprintf(w, "<b>%.2f RUB = %.2f EUR</b>\n", amount, rubToEur(amount))
			} else if choice == 3 {
				fmt.Fprintf(w, "<b>%.2f USD = %.2f RUB</b>\n", amount, usdToRub(amount))
			} else if choice == 4 {
				fmt.Fprintf(w, "<b>%.2f EUR = %.2f RUB</b>\n", amount, eurToRub(amount))
			}
		}
		fmt.Fprint(w, `<br><br><a href="/"><button>На главную</button></a>`)
	})

	http.HandleFunc("/randomizer", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		fmt.Fprint(w, `
			<h2>Генератор рандомных чисел</h2>
			<form action="/randomizer" method="GET">
				<input type="hidden" name="run" value="true">

				<button type="submit">Выдать случайное число</button>
			</form><br>`
		)

		if r.URL.Query().Get("run") == "true" {
			randomNumber := rand.Intn(100) + 1 
			fmt.Fprintf(w, "<b>Результат: %d</b>", randomNumber)
		}

		fmt.Fprint(w, `<br><br><a href="/"><button>На главную</button></a>`)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}