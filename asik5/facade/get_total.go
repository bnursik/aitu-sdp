package facade

import (
	"asik5/decorator"
)

func (f *PricingFacade) GetTotal(req QuoteRequest) (int, error) {
	var p decorator.IPricer = &decorator.BaseQuote{PerDay: req.PerDay}

	// Compose decorators based on request flags
	if req.AddInsurance {
		p = decorator.NewInsuranceDecorator(p, req.InsPerDay)
	}
	if req.AddChildSeat {
		p = decorator.NewChildSeatDecorator(p, req.SeatPerDay)
	}

	amount, err := p.Total(decorator.QuoteInput{Days: req.Days})
	if err != nil {
		return 0, err
	}
	return amount, nil
}
