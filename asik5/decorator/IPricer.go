package decorator

type IPricer interface {
	Total(in QuoteInput) (int, error)
}
