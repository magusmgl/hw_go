package main

import "fmt"

func main() {
	const usdToEur float64 = 0.87
	const usdToRub float64 = 78.38
	eurToRub := usdToRub / usdToEur
	fmt.Print("Курс EUR в  RUB  - ", eurToRub)
}
