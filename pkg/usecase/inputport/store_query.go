package inputport

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
	"golang-trainning-backend/pkg/usecase/query"
)

type StoreQueryUsecase interface {
	GetByID(ctx context.Context, id uint) (querymodel.StoreQuery, error)
	GetList(ctx context.Context, conditions []query.Condition) ([]querymodel.StoreQuery, error)
}
