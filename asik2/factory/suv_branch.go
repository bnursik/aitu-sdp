package factory

type SuvBranch struct{
	branchName string
}

func NewSuvBranch(name string) *SuvBranch {
	return &SuvBranch{branchName: name}
}

func (SuvBranch *SuvBranch) Name() string {
	return SuvBranch.branchName
}

func (SuvBranch *SuvBranch) MakeCar() ICar {
	return SuvCar{}
}