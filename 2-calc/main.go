package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operation := getOperation()
	floatSlice := getFloatSlice()
	result := getResult(operation, floatSlice)
	fmt.Printf("Результат операции %v равен %0.2f", operation, result)

}

func getFloatSlice() []float64 {
	floatSlice := make([]float64, 0, 2)
	var inputString string
	fmt.Print("Введите числа через ',': ")
	fmt.Scan(&inputString)
	for _, value := range strings.Split(inputString, ",") {
		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Printf("Ошибка парсинга %s: %v", value, err)
			continue
		}

		floatSlice = append(floatSlice, num)
	}
	return floatSlice
}

func getOperation() string {
	var operationName string
	for {
		fmt.Println("Введите операцию: ")
		fmt.Scan(&operationName)

		if operationName != "AVG" &&
			operationName != "SUM" &&
			operationName != "MED" {
			continue
		}
		break
	}
	return operationName
}

func getResult(operation string, floatSlice []float64) float64 {
	var result = 0.0
	switch operation {
	case "SUM":
		for _, value := range floatSlice {
			result += value
		}
	case "AVG":
		var sum float64
		for _, value := range floatSlice {
			sum += value
		}
		sliceLength := len(floatSlice)
		result = sum / float64(sliceLength)

	case "MED":
		sort.Float64s(floatSlice)
		sliceLength := len(floatSlice)
		if sliceLength%2 == 0 {
			index := sliceLength/2 - 1
			result = (floatSlice[index] + floatSlice[index+1]) / 2
		} else {
			index := (sliceLength - 1) / 2
			result = floatSlice[index]
		}
	}
	return result
}
