package factory

import "fmt"

type RentalContract struct {
	Branch string
	Car    ICar
	Days   int
	Total  float64
}

func Rent(branch IRentalBranch, days int) RentalContract {
	car := branch.MakeCar()
	return RentalContract{
		Branch: branch.Name(),
		Car:    car,
		Days:   days,
		Total:  car.DailyRate() * float64(days),
	}
}

func (c RentalContract) String() string {
	return fmt.Sprintf(
		"Branch: %s\nCar: %s | $%.2f/day | %d seats\nDays: %d\nTotal: $%.2f",
		c.Branch,
		c.Car.Model(), c.Car.DailyRate(), c.Car.SeatCount(),
		c.Days,
		c.Total,
	)
}