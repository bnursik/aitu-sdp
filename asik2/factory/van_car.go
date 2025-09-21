package factory

import "fmt"

type VanCar struct{}

func (v VanCar) Model() string     { 
	return "Van" 
}
func (v VanCar) DailyRate() float64 { 
	return 89.00 
}
func (v VanCar) SeatCount() int     { 
	return 9 
}
func (VanCar VanCar) String() string {
	return fmt.Sprintf("%s: $%.2f/day, Seats: %d", VanCar.Model(), VanCar.DailyRate(), VanCar.SeatCount())
}