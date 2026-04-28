package querymodel

type RegionQuery interface {
	IsNil() bool
	GetID() uint
	GetName() string
}

type Region struct {
	ID   uint
	Name string
}

func (r *Region) IsNil() bool     { return r.ID == 0 }
func (r *Region) GetID() uint     { return r.ID }
func (r *Region) GetName() string { return r.Name }

type NilRegion struct{}

func (r *NilRegion) IsNil() bool     { return true }
func (r *NilRegion) GetID() uint     { return 0 }
func (r *NilRegion) GetName() string { return "" }
