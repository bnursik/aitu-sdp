package main

import (
	"fmt"

	concretevisitors "asik8/pkg/rental/concrete_visitors"
	"asik8/pkg/rental/vehicle"
)

func main() {
	fleet := []vehicle.IElement{
		&vehicle.EconomyCar{DailyBase: 39.0, Seats: 4},
		&vehicle.SuvCar{DailyBase: 75.0, Awd: true},
	}

	// 1) Pricing behavior via Visitor
	pricer := concretevisitors.NewPricingVisitor(3)
	for _, e := range fleet {
		e.Accept(pricer)
	}
	fmt.Println("Pricing:")
	for _, n := range pricer.Notes {
		fmt.Println(" -", n)
	}
	fmt.Printf("Total = %d\n\n", pricer.Total)

	// 2) Maintenance behavior via Visitor
	mt := concretevisitors.NewMaintenanceVisitor()
	for _, e := range fleet {
		e.Accept(mt)
	}
	fmt.Println("Maintenance:")
	for _, r := range mt.Report {
		fmt.Println(" -", r)
	}
}
