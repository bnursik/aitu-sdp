package main

import (
	"fmt"

	"asik4/pricing"
	"asik4/vehicles"
)

func main() {
	// Implementations (the "how")
	flat := &pricing.FlatRate{PerDay: 10000}
	perKm := &pricing.PerKm{BasePerDay: 3000, PerKm: 50}

	// Abstractions (the "what") bridged to implementations
	economy := vehicles.NewEconomyVehicle(flat) 
	suv := vehicles.NewSuvVehicle(perKm)        

	q1, _ := economy.Quote(pricing.PricingInput{Days: 3, Km: 50})   
	q2, _ := suv.Quote(pricing.PricingInput{Days: 2, Km: 180})      

	fmt.Printf("Economy (FlatRate) 3 days: %.2f Tenge\n", q1)
	fmt.Printf("SUV (PerKm) 2 days / 180 km: %.2f Tenge\n", q2)

	economyPerKm := vehicles.NewEconomyVehicle(perKm)
	q3, _ := economyPerKm.Quote(pricing.PricingInput{Days: 3, Km: 120})
	fmt.Printf("Economy (PerKm) 3 days / 120 km: %.2f Tenge\n", q3)
}
