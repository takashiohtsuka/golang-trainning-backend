package outputport

import (
	"context"

	"golang-trainning-backend/pkg/domain/entity"
	"golang-trainning-backend/pkg/domain/collection"
	"golang-trainning-backend/pkg/usecase/query"
)

type ManagementStaffRepository interface {
	FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.ManagementStaffEntity], error)
	FindOne(ctx context.Context, conditions []query.Condition) (entity.ManagementStaffEntity, error)
	Create(ctx context.Context, m *entity.ManagementStaff) error
	Update(ctx context.Context, m *entity.ManagementStaff) error
}
