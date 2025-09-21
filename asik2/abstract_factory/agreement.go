package abstractfactory

import "fmt"

type RentalAgreement struct {
	PackageName string
	Car         ICar
	Insurance   IInsurance
	Gps         IGpsDevice
	Days        int
	Total       float64
}

func BuildAgreement(factory IRentalPackageFactory, days int) RentalAgreement {
	car := factory.CreateCar()
	insurance := factory.CreateInsurance()
	gps := factory.CreateGps()

	total := (car.DailyRate() + insurance.DailyRate() + gps.DailyRate()) * float64(days)

	return RentalAgreement{
		PackageName: factory.PackageName(),
		Car:         car,
		Insurance:   insurance,
		Gps:         gps,
		Days:        days,
		Total:       total,
	}
}

func (a RentalAgreement) String() string {
	return fmt.Sprintf(
		"Package: %s\nCar: %s (%d seats) $%.2f/day\nInsurance: %s $%.2f/day\nGPS: %s $%.2f/day\nDays: %d\nTotal: $%.2f",
		a.PackageName,
		a.Car.Model(), a.Car.SeatCount(), a.Car.DailyRate(),
		a.Insurance.Name(), a.Insurance.DailyRate(),
		a.Gps.Tier(), a.Gps.DailyRate(),
		a.Days,
		a.Total,
	)
}
