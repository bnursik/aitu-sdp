package decorator

func (d *ChildSeatDecorator) Total(in QuoteInput) (int, error) {
	base, err := d.inner.Total(in)
	if err != nil {
		return 0, err
	}
	return base + in.Days*d.PerDaySeatFee, nil
}
