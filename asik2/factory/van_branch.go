package factory

type VanBranch struct {
	branchName string
}

func NewVanBranch(name string) *VanBranch {
	return &VanBranch{branchName: name}
}

func (v *VanBranch) Name() string {
	return v.branchName
}

func (v *VanBranch) MakeCar() ICar {
	return VanCar{}
}
