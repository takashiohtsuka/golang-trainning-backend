package stores

type Get struct {
	ID uint `param:"id" validate:"required,min=1"`
}
