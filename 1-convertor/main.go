package main

import "fmt"

func main() {
	const usdToEur float64 = 0.87
	const usdToRub float64 = 78.38
	// eurToRub := usdToRub / usdToEur

	originalCurrency, targetCurrency, amountCurrency := getUserInput()
	convertCurrency(originalCurrency, targetCurrency, amountCurrency)
}

func getUserInput() (string, string, float64) {
	var originalCurrency string
	var targetCurrency string
	var amountCurrency float64
	fmt.Print("Введите исходную валюту: ")
	fmt.Scan(&originalCurrency)
	fmt.Print("Введите целевую валюту: ")
	fmt.Scan(&targetCurrency)
	fmt.Print("Введите количество валюты для конвертации")
	fmt.Scan(&amountCurrency)
	return originalCurrency, targetCurrency, amountCurrency
}

func convertCurrency(originalCurrency string, targetCurrency string, amountCurrency float64) {
}
