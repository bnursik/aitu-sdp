package concretevisitors

import (
	"fmt"

	"asik8/pkg/rental/vehicle"
)

var _ vehicle.IVisitor = (*MaintenanceVisitor)(nil)

type MaintenanceVisitor struct {
	Report []string
}

func NewMaintenanceVisitor() *MaintenanceVisitor { return &MaintenanceVisitor{} }

func (v *MaintenanceVisitor) VisitEconomy(c *vehicle.EconomyCar) {
	v.Report = append(v.Report, fmt.Sprintf("Economy: basic inspection, seats=%d", c.Seats))
}

func (v *MaintenanceVisitor) VisitSuv(c *vehicle.SuvCar) {
	task := "standard inspection"
	if c.Awd {
		task = "AWD drivetrain check + standard inspection"
	}
	v.Report = append(v.Report, fmt.Sprintf("SUV: %s", task))
}
