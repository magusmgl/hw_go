package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Reocver ", r)
		}
	}
	fmt.Println("___ Калькулятор индекса массы тела ___")
	for {
		userKg, userHeight := getUserInput()
		IMT, err := calculateIMT(userKg, userHeight)
		if err != nil {
			fmt.Println("Не заданы параметры для расчета")
			continue
			panic("Не заданы параметры для расчета")
		}
		outputResult(IMT)
		isRepeatCalculation := checkRepeatCalculation()
		if !isRepeatCalculation {
			break
		}
	}
	// const IMTPower = 2
	// var userHeight float64 // 0.
	// var userKg float64
	// // fmt.Println("___ Калькулятор индекса массы тела ___")
	// fmt.Print("Введите свой рост (в сантиметрах): ")
	// fmt.Scan(&userHeight)
	// fmt.Print("Введите свой вес (в кг): ")
	// fmt.Scan(&userKg)
	// IMT := userKg / math.Pow(userHeight/100, IMTPower)
	// fmt.Printf("Ваш индекс массы тела: %v", IMT)
	// fmt.Printf("Ваш индекс массы тела: %.0f", IMT)

	// result := fmt.Sprintf("Ваш индекс массы тела: %.0f", IMT)
	// fmt.Print(result)

	// isLean := IMT < 16

	// if IMT < 16 {
	// 	fmt.Println("У вас сильный дефицит массы веса")
	// } else if IMT < 18.5 {
	// 	fmt.Println("У вас дефицит массы веса")
	// } else if IMT < 25 {
	// 	fmt.Println("У вас нормальный вес")
	// } else if IMT < 30 {
	// 	fmt.Println("У вас избыточный вес")
	// } else {
	// 	fmt.Println("У вас степень ожирения")
	// }
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Print(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы веса")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы веса")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

// ver 1 calculateIMT

// func calculateIMT(userKg float64, userHeight float64) float64 {
// 	const IMTPower = 2
// 	IMT := userKg / math.Pow(userHeight/100, IMTPower)
// 	return IMT
// }

// ver 2 calculateIMT
func calculateIMT(userKg float64, userHeight float64) (float64, error) {
	if userKg <= 0 || userHeight <= 0 {
		return 0, errors.New("no params error")
	}
	IMT := userKg / math.Pow(userHeight/100, IMTPower)
	return IMT, nil
}

func getUserInput() (float64, float64) {
	var userHeight float64 // 0.
	var userKg float64
	fmt.Print("Введите свой рост (в сантиметрах): ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес (в кг): ")
	fmt.Scan(&userKg)
	return userKg, userHeight
}

func checkRepeatCalculation() bool {
	var userChoice string
	fmt.Print("Вы хотите сделть еще расчет (y/n): ")
	fmt.Scan(&userChoice)
	if userChoice == "y" || userChoice == "Y" {
		return true
	}
	return false

}
