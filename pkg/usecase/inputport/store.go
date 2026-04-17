package inputport

import (
	"context"

	"golang-trainning-backend/pkg/usecase/input"
)

type StoreUsecase interface {
	Create(ctx context.Context, i input.CreateStoreInput) error
	Update(ctx context.Context, i input.UpdateStoreInput) error
}
