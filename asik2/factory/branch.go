package factory

type IRentalBranch interface {
	Name() string
	MakeCar() ICar // factory method
}