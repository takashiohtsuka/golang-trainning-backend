package outputport

import (
	"context"

	"golang-trainning-backend/pkg/domain/collection"
	"golang-trainning-backend/pkg/domain/entity"
	"golang-trainning-backend/pkg/usecase/query"
)

type WomanRepository interface {
	FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.WomanEntity], error)
	FindOne(ctx context.Context, conditions []query.Condition) (entity.WomanEntity, error)
	Create(ctx context.Context, w *entity.Woman) (uint, error)
	Update(ctx context.Context, w *entity.Woman) error
	SaveImage(ctx context.Context, womanID uint, path string) error

	FindActiveImmediateAvailableByStore(ctx context.Context, storeID uint) ([]*entity.ImmediateAvailableWoman, error)
	CreateImmediateAvailable(ctx context.Context, iaw *entity.ImmediateAvailableWoman) error
	DeleteImmediateAvailable(ctx context.Context, storeID, womanID uint) error
}
