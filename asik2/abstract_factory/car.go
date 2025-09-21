package abstractfactory

type ICar interface {
	Model() string
	SeatCount() int
	DailyRate() float64
}
