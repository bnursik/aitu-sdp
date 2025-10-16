package facade

type IPricingFacade interface {
	GetTotal(req QuoteRequest) (int, error)
}
