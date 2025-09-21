package factory

type ICar interface {
	Model() string
	DailyRate() float64
	SeatCount() int
	String() string
}
