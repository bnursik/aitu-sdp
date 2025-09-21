package abstractfactory

type IGpsDevice interface {
	Tier() string
	DailyRate() float64
}
