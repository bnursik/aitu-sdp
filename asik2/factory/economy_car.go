package factory

import "fmt"

type EconomyCar struct{}

func (EconomyCar) Model() string      { 
	return "Economy" 
}
func (EconomyCar) DailyRate() float64 { 
	return 39.99 
}
func (EconomyCar) SeatCount() int     { 
	return 4 
}
func (EconomyCar EconomyCar) String() string   { 
	return fmt.Sprintf("%s: $%.2f/day, Seats: %d", EconomyCar.Model(), EconomyCar.DailyRate(), EconomyCar.SeatCount()) 
}
