package vehicle

type SuvCar struct {
	DailyBase int
	Awd       bool
}

func (c *SuvCar) Accept(v IVisitor) { v.VisitSuv(c) }
func (c *SuvCar) TypeName() string  { return "SuvCar" }
