package decorator

type ChildSeatDecorator struct {
	inner          IPricer
	PerDaySeatFee  int 
}

func NewChildSeatDecorator(inner IPricer, perDay int) *ChildSeatDecorator {
	return &ChildSeatDecorator{inner: inner, PerDaySeatFee: perDay}
}
