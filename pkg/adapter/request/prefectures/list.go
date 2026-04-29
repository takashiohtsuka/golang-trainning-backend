package prefectures

type List struct {
	RegionID uint `query:"region_id" validate:"required,min=1"`
}
