package abstractfactory

type IInsurance interface {
	Name() string
	DailyRate() float64
	CoverageSummary() string
}
