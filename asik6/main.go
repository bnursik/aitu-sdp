package main

import (
	"fmt"

	"asik6/payment"
)

func main() {
	// Start with Card strategy
	card := &payment.CardPayment{}
	pp := payment.NewPaymentProcessor(card)

	r1, _ := pp.Pay(payment.PayRequest{
		Amount:         149.90,
		Currency:       "KZT",
		IdempotencyKey: "r1",
	})
	fmt.Printf("Card result: %+v\n", r1)

	// Switch to Kaspi strategy at runtime
	kaspi := &payment.KaspiPayment{}
	pp.SetStrategy(kaspi)

	r2, _ := pp.Pay(payment.PayRequest{
		Amount:         149.90,
		Currency:       "KZT",
		IdempotencyKey: "r2",
	})
	fmt.Printf("Kaspi result: %+v\n", r2)
}
