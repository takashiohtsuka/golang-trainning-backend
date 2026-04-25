package outputport

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
)

type StoreQueryRepository interface {
	FindByID(ctx context.Context, id uint) (querymodel.StoreQuery, error)
}
