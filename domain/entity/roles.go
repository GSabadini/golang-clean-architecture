package entity

type Roles struct {
	canTransfer bool
}

func (r Roles) CanTransfer() bool {
	return r.canTransfer
}
