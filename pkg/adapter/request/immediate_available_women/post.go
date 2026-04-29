package immediate_available_women

type Post struct {
	StoreID uint `param:"id"      validate:"required"`
	WomanID uint `json:"woman_id" validate:"required"`
}
