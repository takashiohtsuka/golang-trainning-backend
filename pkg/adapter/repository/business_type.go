package repository

import (
	"context"

	"golang-trainning-backend/pkg/helper"
	"golang-trainning-backend/pkg/querymodel"
	"golang-trainning-backend/pkg/usecase/outputport"

	"gorm.io/gorm"
)

type businessTypeRepository struct {
	db *gorm.DB
}

func NewBusinessTypeRepository(db *gorm.DB) outputport.BusinessTypeRepository {
	return &businessTypeRepository{db: db}
}

func (r *businessTypeRepository) FindAll(ctx context.Context) ([]querymodel.BusinessTypeQuery, error) {
	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(`SELECT id, code FROM business_types ORDER BY id`).Scan(&rows).Error; err != nil {
		return nil, err
	}

	result := make([]querymodel.BusinessTypeQuery, 0, len(rows))
	for _, row := range rows {
		result = append(result, &querymodel.BusinessType{
			ID:   helper.ToUint(row["id"]),
			Code: helper.ToString(row["code"]),
		})
	}
	return result, nil
}
