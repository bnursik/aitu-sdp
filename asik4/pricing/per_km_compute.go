package pricing

import "fmt"

func (s *PerKm) ComputeTotal(in PricingInput) (float64, error) {
	if in.Days <= 0 {
		return 0, fmt.Errorf("invalid days: %d", in.Days)
	}
	if in.Km < 0 {
		return 0, fmt.Errorf("invalid km: %v", in.Km)
	}
	return float64(in.Days)*s.BasePerDay + in.Km*s.PerKm, nil
}
