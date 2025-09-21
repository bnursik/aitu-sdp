package abstractfactory

type LuxuryCar struct{}

func (LuxuryCar) Model() string      { 
	return "Luxury" 
}
func (LuxuryCar) SeatCount() int     { 
	return 5 
}
func (LuxuryCar) DailyRate() float64 {
	 return 149.00
	 }

type LuxuryInsurance struct{}

func (LuxuryInsurance) Name() string          { 
	return "Luxury Insurance" 
}
func (LuxuryInsurance) DailyRate() float64    { 
	return 24.00 
}
func (LuxuryInsurance) CoverageSummary() string { 
	return "Full coverage + low deductible + roadside assist" 
}

type PremiumGps struct{}

func (PremiumGps) Tier() string       { 
	return "Premium GPS"
 }
func (PremiumGps) DailyRate() float64 { 
	return 9.00 
}

type LuxuryPackageFactory struct {
	factoryName string
}

func NewLuxuryPackageFactory(name string) *LuxuryPackageFactory {
	return &LuxuryPackageFactory{factoryName: name}
}

func (f *LuxuryPackageFactory) CreateCar() ICar             { 
	return LuxuryCar{}
 }
func (f *LuxuryPackageFactory) CreateInsurance() IInsurance { 
	return LuxuryInsurance{} 
}
func (f *LuxuryPackageFactory) CreateGps() IGpsDevice       { 
	return PremiumGps{} 
}
func (f *LuxuryPackageFactory) PackageName() string         { 
	return f.factoryName 
}
