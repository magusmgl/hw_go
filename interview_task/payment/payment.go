package payment

import (
	"errors"
	"fmt"
)

type PaymentProccessor interface {
	Proccess(amount float64) error
	Verify(amount float64) bool
}

type CreditCardProcessor struct {
	Limit float64
}

func (c *CreditCardProcessor) Proccess(amount float64) error {
	if amount > c.Limit {
		return errors.New("credit card limit exceed")
	}
	fmt.Printf("Processed payment of $%.2f using credit card", amount)
	return nil
}

func (c *CreditCardProcessor) Verify(amount float64) bool {
	return amount <= c.Limit
}

func ExecutePayment(processor PaymentProccessor, amount float64) {
	if processor.Verify(amount) {
		if err := processor.Proccess(amount); err != nil {
			fmt.Println("Error: ", err)
		}
	} else {
		fmt.Println("Verifiation failed: ", amount)
	}
}
