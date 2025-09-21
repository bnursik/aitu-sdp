package abstractfactory


type EconomyCar struct{}

func (EconomyCar) Model() string      { 
	return "Economy" 
}
func (EconomyCar) SeatCount() int     { 
	return 4 
}
func (EconomyCar) DailyRate() float64 { 
	return 39.99 
}

type EconomyInsurance struct{}

func (EconomyInsurance) Name() string          { 
	return "Economy Insurance" 
}
func (EconomyInsurance) DailyRate() float64    { 
	return 6.50 
}
func (EconomyInsurance) CoverageSummary() string { 
	return "Liability + basic collision" 
}

type BasicGps struct{}

func (BasicGps) Tier() string        { 
	return "Basic GPS" 
}
func (BasicGps) DailyRate() float64  {
	 return 3.00
	}


type EconomyPackageFactory struct {
	factoryName string
}

func NewEconomyPackageFactory(name string) *EconomyPackageFactory {
	return &EconomyPackageFactory{factoryName: name}
}

func (f *EconomyPackageFactory) CreateCar() ICar{ 
	return EconomyCar{} 
}
func (f *EconomyPackageFactory) CreateInsurance() IInsurance{ 
	return EconomyInsurance{} 
}
func (f *EconomyPackageFactory) CreateGps() IGpsDevice{ 
	return BasicGps{} 
}
func (f *EconomyPackageFactory) PackageName() string{
	return f.factoryName 
}
