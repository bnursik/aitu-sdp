package vehicles

import "asik4/pricing"

func NewSuvVehicle(strategy pricing.IPricingStrategy) *Vehicle {
	return &Vehicle{Name: "SUV", strategy: strategy}
}
