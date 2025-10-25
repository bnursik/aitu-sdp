package payment

type PayRequest struct {
	Amount         float64
	Currency       string
	IdempotencyKey string
}
