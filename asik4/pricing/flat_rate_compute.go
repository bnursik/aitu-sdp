package pricing

import "fmt"

func (s *FlatRate) ComputeTotal(in PricingInput) (float64, error) {
	if in.Days <= 0 {
		return 0, fmt.Errorf("invalid days: %d", in.Days)
	}
	return float64(in.Days) * s.PerDay, nil
}
