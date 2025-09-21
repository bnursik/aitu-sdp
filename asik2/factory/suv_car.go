package factory

import "fmt"

type SuvCar struct{}

func (s SuvCar) Model() string     { 
	return "SUV" 
}
func (s SuvCar) DailyRate() float64 { 
	return 79.50 
}
func (s SuvCar) SeatCount() int     { 
	return 7 
}
func (SuvCar SuvCar) String() string {
	return fmt.Sprintf("%s: $%.2f/day, Seats: %d", SuvCar.Model(), SuvCar.DailyRate(), SuvCar.SeatCount())
}