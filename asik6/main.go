package main

import (
	"fmt"

	"asik6/payment"
)

func main() {
	// Start with Card strategy
	card := &payment.CardPayment{}
	paymentProcessor := payment.NewPaymentProcessor(card)

	cardResult, _ := paymentProcessor.Pay(payment.PayRequest{
		Amount:         149.90,
		Currency:       "KZT",
		IdempotencyKey: "r1",
	})
	fmt.Printf("Card result: %+v\n", cardResult)

	// Switch to Kaspi strategy at runtime
	kaspi := &payment.KaspiPayment{}
	paymentProcessor.SetStrategy(kaspi)

	kaspiResult, _ := paymentProcessor.Pay(payment.PayRequest{
		Amount:         149.90,
		Currency:       "KZT",
		IdempotencyKey: "r2",
	})
	fmt.Printf("Kaspi result: %+v\n", kaspiResult)
}
