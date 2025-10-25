package payment

import "fmt"

type CardPayment struct{}

func (c *CardPayment) Pay(req PayRequest) (PayResult, error) {
	return PayResult{
		PaymentID: fmt.Sprintf("card_%s", req.IdempotencyKey),
		Status:    "authorized",
	}, nil
}
