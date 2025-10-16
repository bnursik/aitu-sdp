package decorator

import "errors"

func (b *BaseQuote) Total(in QuoteInput) (int, error) {
	if in.Days <= 0 {
		return 0, errors.New("invalid days: must be > 0")
	}
	return in.Days * b.PerDay, nil
}
