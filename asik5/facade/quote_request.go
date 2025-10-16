package facade

type QuoteRequest struct {
	PerDay int
	Days   int

	AddInsurance bool  
	InsPerDay    int

	AddChildSeat bool
	SeatPerDay   int
}