package photo

import (
	"golang-trainning-backend/pkg/domain/entity"
	"golang-trainning-backend/pkg/infrastructure/model"
)

func ToEntity(m *model.Photo) *entity.Photo {
	return &entity.Photo{
		ID:        m.ID,
		BlogID:    m.BlogID,
		URL:       m.URL,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToOrmModel(e *entity.Photo) *model.Photo {
	return &model.Photo{
		ID:        e.ID,
		BlogID:    e.BlogID,
		URL:       e.URL,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
