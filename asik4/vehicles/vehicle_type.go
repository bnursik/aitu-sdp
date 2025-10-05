package vehicles

import (
	"asik4/pricing"
)

type Vehicle struct {
	Name     string
	strategy pricing.IPricingStrategy
}
