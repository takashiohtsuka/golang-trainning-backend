package inputport

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
)

type StoreQueryUsecase interface {
	GetByID(ctx context.Context, id uint) (querymodel.StoreQuery, error)
}
