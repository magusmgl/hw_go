package stock

import (
	"time"
)

func UpdateProductStock() <-chan map[string]int {
	stockUpdates := make(chan map[string]int)

	go func() {
		defer close(stockUpdates)
		currentStock := map[string]int{
			"Apples":  50,
			"Bananas": 30,
			"Oranges": 20,
			"Grapes":  15,
		}

		for i := 0; i < 5; i++ {
			newStock := make(map[string]int)

			for product, quantity := range currentStock {
				newStock[product] = int(float64(quantity) * 0.95)
			}
			stockUpdates <- newStock

			currentStock = newStock
			time.Sleep(150 * time.Millisecond)
		}
	}()

	return stockUpdates
}
