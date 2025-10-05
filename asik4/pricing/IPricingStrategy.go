package pricing

type IPricingStrategy interface {
	ComputeTotal(in PricingInput) (float64, error)
}
