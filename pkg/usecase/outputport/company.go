package outputport

import (
	"context"

	"golang-trainning-backend/pkg/domain/entity"
	"golang-trainning-backend/pkg/domain/collection"
	"golang-trainning-backend/pkg/usecase/query"
)

type CompanyRepository interface {
	FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.CompanyEntity], error)
	FindOne(ctx context.Context, conditions []query.Condition) (entity.CompanyEntity, error)
	Create(ctx context.Context, c *entity.Company) error
	Update(ctx context.Context, c *entity.Company) error
}
