package decorator

type InsuranceDecorator struct {
	inner            IPricer
	PerDayInsurance  int 
}

func NewInsuranceDecorator(inner IPricer, perDay int) *InsuranceDecorator {
	return &InsuranceDecorator{inner: inner, PerDayInsurance: perDay}
}
