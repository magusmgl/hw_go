package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Data struct {
	Results []struct {
		InsurancePrice struct {
			TotalSumPlan float64 `json:"totalSumPlan"`
			TotalSumFact float64 `json:"totalSumFact"`
		} `json:"insurancePrice"`
		DeliveryService string `json:deliveryService"`
	} `json:"results"`
}

func main() {

	// Укажите здесь имя вашего JSON-файла
	jsonBytes, err := os.ReadFile("insurance-fee-with-dv-price-not-equal.json")
	if err != nil {
		log.Fatal(err)
	}
	// Определяем структуру
	var data Data

	// Распаковываем json в структуру
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(data.Results))
	for i, result := range data.Results {
		// plan, err := strconv.ParseFloat(result.InsurancePrice.TotalSumPlan, 64)
		if result.DeliveryService == "5Post" {
			plan := result.InsurancePrice.TotalSumPlan
			if err != nil {
				log.Printf("Ошибка конвертации TotalSumPlan для результата #%d: %v", i+1, err)
				continue // Переходим к следующему результату, если данные некорректны
			}
			fact := result.InsurancePrice.TotalSumFact
			// fact, err := strconv.ParseFloat(result.InsurancePrice.TotalSumFact, 64)
			if err != nil {
				log.Printf("Ошибка конвертации TotalSumFact для результата #%d: %v", i+1, err)
				continue // Переходим к следующему результату
			}

			diff := plan / fact

			if diff < 1.16 {
				fmt.Printf("Результат #%d: План = %.2f, Факт = %.2f, Разница = %.2f\n", i+1, plan, fact, diff)
			}
		} else {
			fmt.Println(result.DeliveryService)
		}

	}
}
