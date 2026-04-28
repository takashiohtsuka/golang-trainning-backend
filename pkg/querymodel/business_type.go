package querymodel

type BusinessTypeQuery interface {
	IsNil() bool
	GetID() uint
	GetCode() string
}

type BusinessType struct {
	ID   uint
	Code string
}

func (b *BusinessType) IsNil() bool     { return b.ID == 0 }
func (b *BusinessType) GetID() uint     { return b.ID }
func (b *BusinessType) GetCode() string { return b.Code }

type NilBusinessType struct{}

func (b *NilBusinessType) IsNil() bool     { return true }
func (b *NilBusinessType) GetID() uint     { return 0 }
func (b *NilBusinessType) GetCode() string { return "" }
