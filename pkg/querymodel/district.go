package querymodel

type DistrictQuery interface {
	IsNil() bool
	GetID() uint
	GetName() string
	GetPrefectureID() uint
}

type District struct {
	ID           uint
	Name         string
	PrefectureID uint
}

func (d *District) IsNil() bool           { return d.ID == 0 }
func (d *District) GetID() uint           { return d.ID }
func (d *District) GetName() string       { return d.Name }
func (d *District) GetPrefectureID() uint { return d.PrefectureID }

type NilDistrict struct{}

func (d *NilDistrict) IsNil() bool           { return true }
func (d *NilDistrict) GetID() uint           { return 0 }
func (d *NilDistrict) GetName() string       { return "" }
func (d *NilDistrict) GetPrefectureID() uint { return 0 }
