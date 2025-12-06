package main

import (
	"fmt"
	"strings"
)

type convertorMap = map[string]float64
type currMap = map[string]convertorMap

func main() {
	const usdToRub float64 = 78.38
	const usdToEur float64 = 0.87
	eurToRub := usdToRub / usdToEur

	usdMap := convertorMap{"EUR": usdToEur, "RUB": usdToRub}
	eurMap := convertorMap{"USD": 1 / usdToEur, "RUB": eurToRub}
	rubMap := convertorMap{"USD": 1 / usdToRub, "EUR": 1 / eurToRub}

	currencyMap := currMap{
		"USD": usdMap,
		"EUR": eurMap,
		"RUB": rubMap}

	fmt.Println("___ Конвертор валюты ___")
	originalCurrency := getOriginalCurrency(&currencyMap)
	amountCurrency := getSumToConvert()
	targetCurrency := getTargetCurrency(originalCurrency, &currencyMap)
	res := convertCurrency(originalCurrency, amountCurrency, targetCurrency, &currencyMap)
	fmt.Printf("\n%0.f", res)
}

func getOriginalCurrency(currencyMap *currMap) string {
	var originalCurrency string
	for {
		fmt.Print("Введите исходную валюту (RUB, USD, EUR): ")
		fmt.Scan(&originalCurrency)
		if _, ok := (*currencyMap)[originalCurrency]; !ok {
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

func getTargetCurrency(originalCurrency string, currencyMap *currMap) string {
	var targetCurrency string
	validTargetMap := (*currencyMap)[originalCurrency]
	availableCurrency := make([]string, 0, len(validTargetMap))
	for curr := range validTargetMap {
		availableCurrency = append(availableCurrency, curr)
	}

	for {
		fmt.Printf("Введите целевую валюту (%s): ", strings.Join(availableCurrency, ", "))
		fmt.Scan(&targetCurrency)
		if _, ok := validTargetMap[targetCurrency]; ok {
			return targetCurrency
		}
		fmt.Println("Не поддерживаемая валюта")
	}
}

func convertCurrency(originalCurrency string, amountCurrency float64, targetCurrency string, currencyMap *currMap) float64 {
	res := (*currencyMap)[originalCurrency][targetCurrency] * amountCurrency
	return res
}
