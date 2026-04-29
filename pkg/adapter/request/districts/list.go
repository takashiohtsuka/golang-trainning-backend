package districts

type List struct {
	PrefectureID uint `query:"prefecture_id" validate:"required,min=1"`
}
