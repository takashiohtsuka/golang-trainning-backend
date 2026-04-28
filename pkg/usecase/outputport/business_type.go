package outputport

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
)

type BusinessTypeRepository interface {
	FindAll(ctx context.Context) ([]querymodel.BusinessTypeQuery, error)
}
