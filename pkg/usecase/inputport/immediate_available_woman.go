package inputport

import "context"

type ImmediateAvailableWomanUsecase interface {
	Create(ctx context.Context, storeID, womanID uint) error
	Delete(ctx context.Context, storeID, womanID uint) error
}
