package outputport

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
	"golang-trainning-backend/pkg/usecase/query"
)

type StoreQueryRepository interface {
	FindByID(ctx context.Context, id uint) (querymodel.StoreQuery, error)
	FindAll(ctx context.Context, conditions []query.Condition) ([]querymodel.StoreQuery, error)
}
