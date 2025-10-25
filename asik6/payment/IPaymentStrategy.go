package payment

type IPaymentStrategy interface {
	Pay(req PayRequest) (PayResult, error)
}
