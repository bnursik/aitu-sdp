package vehicle

type EconomyCar struct {
	DailyBase int
	Seats     int
}

func (c *EconomyCar) Accept(v IVisitor) { v.VisitEconomy(c) }
func (c *EconomyCar) TypeName() string  { return "EconomyCar" }
