package company

import (
	"time"

	"golang-trainning-backend/pkg/domain/entity"
	"golang-trainning-backend/pkg/infrastructure/model"

	"gorm.io/gorm"
)

func ToEntity(m *model.Company) *entity.Company {
	return &entity.Company{
		ID:        m.ID,
		Name:      m.Name,
		Rank:      m.Rank,
		IsActive:  m.IsActive,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		DeletedAt: toTimePtr(m.DeletedAt),
	}
}

func ToOrmModel(e *entity.Company) *model.Company {
	return &model.Company{
		ID:        e.ID,
		Name:      e.Name,
		Rank:      e.Rank,
		IsActive:  e.IsActive,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		DeletedAt: toDeletedAt(e.DeletedAt),
	}
}

func toTimePtr(d gorm.DeletedAt) *time.Time {
	if d.Valid {
		return &d.Time
	}
	return nil
}

func toDeletedAt(t *time.Time) gorm.DeletedAt {
	if t != nil {
		return gorm.DeletedAt{Time: *t, Valid: true}
	}
	return gorm.DeletedAt{}
}
