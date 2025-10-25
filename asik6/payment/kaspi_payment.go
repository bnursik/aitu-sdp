package payment

import "fmt"

type KaspiPayment struct{}

func (k *KaspiPayment) Pay(req PayRequest) (PayResult, error) {
	return PayResult{
		PaymentID: fmt.Sprintf("kaspi_%s", req.IdempotencyKey),
		Status:    "authorized",
	}, nil
}
