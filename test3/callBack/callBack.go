package callback

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type CallBack struct {
	OrderNumber string `json:"OrderNumber"`
	RawData     struct {
		Data struct {
			Order struct {
				PaymentInfo struct {
					Tariff        string  `json:"tariff"`
					AssessedValue float64 `json:"assessedValue"`
				} `json:"paymentInfo"`
				ShippingAddress struct {
					City string `json:"city"`
				} `json:"shippingAddress"`
				ConsigneeAddress struct {
					City string `json:"city"`
				} `json:"consigneeAddress"`
				Packages []struct {
					Weight int32 `json:"weight"`
				} `json:"packages"`
			} `json:"order"`
			Services []struct {
				Name string  `json:"name"`
				Sum  float64 `json:"sum"`
			} `json:"services"`
		} `json:"data"`
	} `json:"RawData"`
}

type DataForTable struct {
	OrderNumber   string
	Tariff        string
	CityFrom      string
	CityTo        string
	Weight        int32
	AssessedValue float64
	BasePrice     float64
	InsurancePay  float64
}

func GetCallBack(filename string) ([]DataForTable, error) {
	var result []DataForTable
	jsonBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла %s: %w", filename, err)
	}
	var data []CallBack

	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return nil, fmt.Errorf("ошибка разбора JSON в файле %s: %w", filename, err)
	}

	startTime := time.Now()

	for _, order := range data {
		var basePrice float64
		var insurancePay float64

		for _, value := range order.RawData.Data.Services {
			if value.Name == "Услуга доставки" {
				basePrice = value.Sum
			}

			if value.Name == "Услуга страховки" {
				insurancePay = value.Sum
			}
		}

		var weight int32
		if len(order.RawData.Data.Order.Packages) > 0 {
			weight = order.RawData.Data.Order.Packages[0].Weight / 1000
		}

		tableRow := DataForTable{
			OrderNumber:   order.OrderNumber,
			Tariff:        order.RawData.Data.Order.PaymentInfo.Tariff,
			CityFrom:      order.RawData.Data.Order.ShippingAddress.City,
			CityTo:        order.RawData.Data.Order.ConsigneeAddress.City,
			Weight:        weight,
			AssessedValue: order.RawData.Data.Order.PaymentInfo.AssessedValue,
			InsurancePay:  float64(insurancePay),
			BasePrice:     float64(basePrice),
		}
		result = append(result, tableRow)
	}
	fmt.Println(time.Since(startTime))
	return result, nil
}
