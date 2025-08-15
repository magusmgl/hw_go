package main

import "fmt"

func main() {
	// trasaction := [3]{1,2,3}
	// banks := [2]{"Tinkoff", "Sberbank"}
	transaction := [5]int{1, 2, 3, 5, 6}
	partial := transaction[1:3]
	fmt.Println(partial)

	transactionNew := transaction
	transactionNew[0] = 30
	fmt.Println(transaction)
	fmt.Println(transactionNew)

	transactionPartial := transaction[1:]
	transactionPartialNew := transactionPartial[:1]
	transactionPartialNew[0] = 30
	fmt.Println(transactionPartial)
	fmt.Println(len(transactionPartial), cap(transactionPartial))
	fmt.Println(len(transactionPartialNew), cap(transactionPartialNew))
	getTransactions()
	sl := []int{1, 2}
	sl[2] = 3
}

func getTransactions() {
	transacitons := []float64{}
	for {
		transaction := scanTransaction()
		if transaction == 0 {
			break
		}
		transacitons = append(transacitons, transaction)
	}
	// fmt.Println(transacitons)
	balance := calculateBalance(transacitons)
	fmt.Printf("Ваш баланс: %.2f", balance)
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Println("Введите транзакцию или n для выхода: ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transacitons []float64) float64 {
	balance := 0.0
	for _, value := range transacitons {
		// fmt.Println(value)
		balance += value
	}
	return balance
}
