package repository

import (
	"context"

	"golang-trainning-backend/pkg/helper"
	"golang-trainning-backend/pkg/querymodel"
	"golang-trainning-backend/pkg/usecase/outputport"

	"gorm.io/gorm"
)

type contractPlanRepository struct {
	db *gorm.DB
}

func NewContractPlanRepository(db *gorm.DB) outputport.ContractPlanRepository {
	return &contractPlanRepository{db: db}
}

func (r *contractPlanRepository) FindAll(ctx context.Context) ([]querymodel.ContractPlanQuery, error) {
	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(`SELECT id, code FROM contract_plans ORDER BY id`).Scan(&rows).Error; err != nil {
		return nil, err
	}

	result := make([]querymodel.ContractPlanQuery, 0, len(rows))
	for _, row := range rows {
		result = append(result, &querymodel.ContractPlan{
			ID:   helper.ToUint(row["id"]),
			Code: helper.ToString(row["code"]),
		})
	}
	return result, nil
}
