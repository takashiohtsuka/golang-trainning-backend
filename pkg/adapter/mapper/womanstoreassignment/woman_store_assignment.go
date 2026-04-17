package womanstoreassignment

import (
	"golang-trainning-backend/pkg/domain/entity"
	"golang-trainning-backend/pkg/infrastructure/model"
)

func ToEntity(m *model.WomanStoreAssignment) *entity.WomanStoreAssignment {
	return &entity.WomanStoreAssignment{
		ID:        m.ID,
		StoreID:   m.StoreID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToOrmModel(e *entity.WomanStoreAssignment, womanID uint) *model.WomanStoreAssignment {
	return &model.WomanStoreAssignment{
		ID:        e.ID,
		WomanID:   womanID,
		StoreID:   e.StoreID,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
