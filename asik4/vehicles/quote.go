package vehicles

import "asik4/pricing"

func (v *Vehicle) Quote(in pricing.PricingInput) (float64, error) {
	return v.strategy.ComputeTotal(in)
}
