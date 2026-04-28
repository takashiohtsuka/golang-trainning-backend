package querymodel

type PrefectureQuery interface {
	IsNil() bool
	GetID() uint
	GetName() string
	GetRegionID() uint
}

type Prefecture struct {
	ID       uint
	Name     string
	RegionID uint
}

func (p *Prefecture) IsNil() bool       { return p.ID == 0 }
func (p *Prefecture) GetID() uint       { return p.ID }
func (p *Prefecture) GetName() string   { return p.Name }
func (p *Prefecture) GetRegionID() uint { return p.RegionID }

type NilPrefecture struct{}

func (p *NilPrefecture) IsNil() bool       { return true }
func (p *NilPrefecture) GetID() uint       { return 0 }
func (p *NilPrefecture) GetName() string   { return "" }
func (p *NilPrefecture) GetRegionID() uint { return 0 }
