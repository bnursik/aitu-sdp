package vehicles

import "asik4/pricing"

func NewEconomyVehicle(strategy pricing.IPricingStrategy) *Vehicle {
	return &Vehicle{Name: "Economy", strategy: strategy}
}
