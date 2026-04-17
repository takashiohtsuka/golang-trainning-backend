package outputport

import (
	"context"

	"golang-trainning-backend/pkg/domain/entity"
	"golang-trainning-backend/pkg/domain/collection"
	"golang-trainning-backend/pkg/usecase/query"
)

type BlogRepository interface {
	FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.BlogEntity], error)
	FindOne(ctx context.Context, conditions []query.Condition) (entity.BlogEntity, error)
	Create(ctx context.Context, b *entity.Blog) error
	Update(ctx context.Context, b *entity.Blog) error
	Delete(ctx context.Context, id uint) error
}
