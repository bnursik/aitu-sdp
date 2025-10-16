package main

import (
	"asik5/decorator"
	"asik5/facade"
	"fmt"
)

func main() {
	//decorator pattern
	fmt.Println("Decorator Pattern Example:")
	var price decorator.IPricer = &decorator.BaseQuote{PerDay: 10000}

	// Add-ons via Decorators (order does not matter for these additive fees)
	price = decorator.NewInsuranceDecorator(price, 3000) // +3000/day
	price = decorator.NewChildSeatDecorator(price, 2000)  // +2000/day

	total, err := price.Total(decorator.QuoteInput{Days: 3})
	if err != nil {
		panic(err)
	}
	fmt.Printf("3 days with insurance + child seat: %d Tenge\n", total)

	// child seat only
	price = decorator.NewChildSeatDecorator(price, 2000)

	total2, _ := price.Total(decorator.QuoteInput{Days: 2})
	fmt.Printf("2 days with child seat only: %d Tenge\n", total2)

	//facade pattern
	fmt.Println("\nFacade Pattern Example:")
	pf := facade.NewPricingFacade()
	req := facade.QuoteRequest{
		PerDay:       10000,
		Days:         3,
		AddInsurance: true,
		InsPerDay:    3000,
		AddChildSeat: true,
		SeatPerDay:   2000,
	}

	totalFacade, err := pf.GetTotal(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Total: %d Tenge\n", totalFacade)
}
