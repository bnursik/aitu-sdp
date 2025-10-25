package payment

import "fmt"

type PaymentProcessor struct {
	strategy IPaymentStrategy
}

func NewPaymentProcessor(s IPaymentStrategy) *PaymentProcessor {
	return &PaymentProcessor{strategy: s}
}

func (p *PaymentProcessor) SetStrategy(s IPaymentStrategy) {
	p.strategy = s
}

func (p *PaymentProcessor) Pay(req PayRequest) (PayResult, error) {
	if p.strategy == nil {
		return PayResult{}, fmt.Errorf("no strategy set")
	}
	return p.strategy.Pay(req)
}
