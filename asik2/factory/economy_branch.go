package factory
	
type EconomyBranch struct{ 
	branchName string 
}

func NewEconomyBranch(name string) *EconomyBranch { 
	return &EconomyBranch{branchName: name} 
}
func (EconomyBranch *EconomyBranch) Name() string { 
	return EconomyBranch.branchName 
}
func (EconomyBranch *EconomyBranch) MakeCar() ICar {
	return EconomyCar{} 
}
