package concretevisitors

import (
	"fmt"

	"asik8/pkg/rental/vehicle"
)

var _ vehicle.IVisitor = (*PricingVisitor)(nil)

type PricingVisitor struct {
	Days  int
	Total int
	Notes []string
}

func NewPricingVisitor(days int) *PricingVisitor {
	return &PricingVisitor{Days: days}
}

func (v *PricingVisitor) VisitEconomy(c *vehicle.EconomyCar) {
	price := v.Days * c.DailyBase
	v.Total += price
	v.Notes = append(v.Notes,
		fmt.Sprintf("Economy × %d days = %d", v.Days, price))
}

func (v *PricingVisitor) VisitSuv(c *vehicle.SuvCar) {
	base := v.Days * c.DailyBase

	if c.Awd {
		base = int(float64(base) * 1.10)
	}

	v.Total += base
	v.Notes = append(v.Notes,
		fmt.Sprintf("SUV × %d days (AWD=%v) = %d", v.Days, c.Awd, base))
}
