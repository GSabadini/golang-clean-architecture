package entity

type Roles struct {
	canTransfer bool
}

func NewRoles(canTransfer bool) Roles {
	return Roles{canTransfer: canTransfer}
}

func (r Roles) CanTransfer() bool {
	return r.canTransfer
}
