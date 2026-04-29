package immediate_available_women

type Delete struct {
	StoreID uint `param:"id"        validate:"required"`
	WomanID uint `param:"woman_id"  validate:"required"`
}
