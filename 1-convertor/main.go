package main

import (
	"fmt"
)

func main() {
	fmt.Println("___ Конвертор валюты ___")
	originalCurrency, targetCurrency, amountCurrency := getUserInput()
	res := convertCurrency(originalCurrency, targetCurrency, amountCurrency)
	fmt.Printf("\n%0.f", res)
}

func getUserInput() (string, float64, string) {
	originalCurrency := getOriginalCurrency()
	amountCurrency := getSumToConvert()
	targetCurrency := getTargetCurrency(originalCurrency)

	return originalCurrency, amountCurrency, targetCurrency
}

func getOriginalCurrency() string {
	var originalCurrency string
	for {
		fmt.Print("Введите исходную валюту (RUB, USD, EUR): ")
		fmt.Scan(&originalCurrency)
		if originalCurrency != "RUB" && originalCurrency != "USD" && originalCurrency != "EUR" {
			fmt.Println("Не поддерживаемый тип валюты")
			continue
		}
		return originalCurrency
	}
}

func getSumToConvert() float64 {
	var amountCurrency float64
	for {
		fmt.Print("Введите количество валюты для конвертации: ")
		fmt.Scan(&amountCurrency)
		if amountCurrency <= 0 {
			fmt.Println("Неверное значение суммы для конвертации")
			continue
		}
		return amountCurrency
	}
}

func getTargetCurrency(originalCurrency string) string {
	var targetCurrency string
	switch originalCurrency {
	case "RUB":
		for {
			fmt.Print("Введите целевую валюту (USD, EUR): ")
			fmt.Scan(&targetCurrency)
			if targetCurrency != "USD" && targetCurrency != "EUR" {
				fmt.Println("Введено неправильный тип валюты")
				continue
			}
			break
		}
	case "USD":
		for {
			fmt.Print("Введите целевую валюту (RUB, EUR): ")
			fmt.Scan(&targetCurrency)
			if targetCurrency != "RUB" && targetCurrency != "EUR" {
				fmt.Println("Введено неправильный тип валюты")
				continue
			}
			break
		}
	default:
		for {
			fmt.Print("Введите целевую валюту (RUB, USD): ")
			fmt.Scan(&targetCurrency)
			if targetCurrency != "RUB" && targetCurrency != "USD" {
				fmt.Println("Введено неправильный тип валюты")
				continue
			}
			break
		}
	}
	return targetCurrency
}

func convertCurrency(originalCurrency string, amountCurrency float64, targetCurrency string) float64 {
	const usdToEur float64 = 0.87
	const usdToRub float64 = 78.38
	eurToRub := usdToRub / usdToEur
	var res float64
	switch {
	case originalCurrency == "RUB" && targetCurrency == "EUR":
		res = amountCurrency / eurToRub
	case originalCurrency == "RUB" && targetCurrency == "USD":
		res = amountCurrency / usdToRub
	case originalCurrency == "USD" && targetCurrency == "RUB":
		res = amountCurrency * usdToRub
	case originalCurrency == "EUR" && targetCurrency == "RUB":
		res = amountCurrency * eurToRub
	case originalCurrency == "USD" && targetCurrency == "EUR":
		res = amountCurrency / usdToEur
	case originalCurrency == "EUR" && targetCurrency == "USD":
		res = amountCurrency * usdToEur
	}
	return res
}
