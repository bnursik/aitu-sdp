package abstractfactory

type SuvCar struct{}

func (SuvCar) Model() string      { 
	return "SUV" 
}
func (SuvCar) SeatCount() int     { 
	return 7 
}
func (SuvCar) DailyRate() float64 { 
	return 89.50 
}

type FamilyInsurance struct{}

func (FamilyInsurance) Name() string          { 
	return "Family Insurance" 
}
func (FamilyInsurance) DailyRate() float64    { 
	return 12.50 
}
func (FamilyInsurance) CoverageSummary() string { 
	return "Liability + collision + child-seat rider" 
}

type MidGps struct{}

func (MidGps) Tier() string       { 
	return "Mid-tier GPS" 
}
func (MidGps) DailyRate() float64 { 
	return 5.00 
}

type FamilySuvPackageFactory struct {
	factoryName string
}

func NewFamilySuvPackageFactory(name string) *FamilySuvPackageFactory {
	return &FamilySuvPackageFactory{factoryName: name}
}

func (f *FamilySuvPackageFactory) CreateCar() ICar             { 
	return SuvCar{} 
}
func (f *FamilySuvPackageFactory) CreateInsurance() IInsurance { 
	return FamilyInsurance{} 
}
func (f *FamilySuvPackageFactory) CreateGps() IGpsDevice       { 
	return MidGps{} 
}
func (f *FamilySuvPackageFactory) PackageName() string         { 
	return f.factoryName 
}
