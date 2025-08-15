package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var menuFunc = map[string]func([]float64) float64{
	"AVG": countAvg,
	"MED": countMed,
	"SUM": countSum,
}

func main() {
	currentFunc := getFunction()
	floatSlice := getFloatSlice()
	result := currentFunc(floatSlice)
	fmt.Printf("Результат операции равен %0.2f", result)

}

func getFloatSlice() []float64 {
	floatSlice := make([]float64, 0, 2)
	var inputString string
	fmt.Print("Введите числа через ',': ")
	fmt.Scan(&inputString)
	for _, value := range strings.Split(inputString, ",") {
		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Printf("Ошибка парсинга %s: %v\n", value, err)
			continue
		}

		floatSlice = append(floatSlice, num)
	}
	return floatSlice
}

func getFunction() func([]float64) float64 {
	var operationName string
	for {
		fmt.Println("Введите операцию: ")
		fmt.Scan(&operationName)
		currentFunc := menuFunc[operationName]
		if currentFunc == nil {
			continue
		}
		return currentFunc
		// if operationName != "AVG" &&
		// 	operationName != "SUM" &&
		// 	operationName != "MED" {
		// 	continue
		// }
		// break
	}
}

// func getResult(operation string, floatSlice []float64) float64 {
// 	currentFunc := menuFunc[operation]
// 	return currentFunc(floatSlice)

// var result = 0.0
// switch operation {
// case "SUM":
// 	for _, value := range floatSlice {
// 		result += value
// 	}
// case "AVG":
// 	var sum float64
// 	for _, value := range floatSlice {
// 		sum += value
// 	}
// 	sliceLength := len(floatSlice)
// 	result = sum / float64(sliceLength)

// case "MED":
// 	sort.Float64s(floatSlice)
// 	sliceLength := len(floatSlice)
// 	if sliceLength%2 == 0 {
// 		index := sliceLength/2 - 1
// 		result = (floatSlice[index] + floatSlice[index+1]) / 2
// 	} else {
// 		index := (sliceLength - 1) / 2
// 		result = floatSlice[index]
// 	}
// }
// return result
// }

func countSum(floatSlice []float64) float64 {
	var result = 0.0
	for _, value := range floatSlice {
		result += value
	}
	return result
}

func countMed(floatSlice []float64) float64 {
	var result = 0.0
	sort.Float64s(floatSlice)
	sliceLength := len(floatSlice)
	if sliceLength%2 == 0 {
		index := sliceLength/2 - 1
		result = (floatSlice[index] + floatSlice[index+1]) / 2
	} else {
		index := (sliceLength - 1) / 2
		result = floatSlice[index]
	}
	return result
}

func countAvg(floatSlice []float64) float64 {
	var result = 0.0
	var sum float64
	for _, value := range floatSlice {
		sum += value
	}
	sliceLength := len(floatSlice)
	result = sum / float64(sliceLength)
	return result
}
