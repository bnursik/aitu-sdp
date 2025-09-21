package abstractfactory

type IRentalPackageFactory interface {
	CreateCar() ICar
	CreateInsurance() IInsurance
	CreateGps() IGpsDevice
	PackageName() string
}
