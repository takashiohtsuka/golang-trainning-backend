package repository

import (
	"context"

	"golang-trainning-backend/pkg/helper"
	"golang-trainning-backend/pkg/querymodel"
	"golang-trainning-backend/pkg/usecase/outputport"

	"gorm.io/gorm"
)

type regionRepository struct {
	db *gorm.DB
}

func NewRegionRepository(db *gorm.DB) outputport.RegionRepository {
	return &regionRepository{db: db}
}

func (r *regionRepository) FindAll(ctx context.Context) ([]querymodel.RegionQuery, error) {
	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(`SELECT id, name FROM regions ORDER BY id`).Scan(&rows).Error; err != nil {
		return nil, err
	}

	result := make([]querymodel.RegionQuery, 0, len(rows))
	for _, row := range rows {
		result = append(result, &querymodel.Region{
			ID:   helper.ToUint(row["id"]),
			Name: helper.ToString(row["name"]),
		})
	}
	return result, nil
}
