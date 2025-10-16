package decorator

func (d *InsuranceDecorator) Total(in QuoteInput) (int, error) {
	base, err := d.inner.Total(in)
	if err != nil {
		return 0, err
	}
	return base + in.Days*d.PerDayInsurance, nil
}
