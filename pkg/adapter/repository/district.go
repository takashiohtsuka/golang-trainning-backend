package repository

import (
	"context"
	"strings"

	"golang-trainning-backend/pkg/helper"
	"golang-trainning-backend/pkg/querymodel"
	"golang-trainning-backend/pkg/usecase/outputport"

	"gorm.io/gorm"
)

type districtRepository struct {
	db *gorm.DB
}

func NewDistrictRepository(db *gorm.DB) outputport.DistrictRepository {
	return &districtRepository{db: db}
}

func (r *districtRepository) FindByPrefectureID(ctx context.Context, prefectureID uint, businessTypeIDs []uint, contractPlanIDs []uint) ([]querymodel.DistrictQuery, error) {
	needsJoin := len(businessTypeIDs) > 0 || len(contractPlanIDs) > 0

	var sql string
	var args []any

	if needsJoin {
		sql = `SELECT DISTINCT d.id, d.name, d.prefecture_id
FROM districts d
INNER JOIN stores s ON s.district_id = d.id
WHERE d.prefecture_id = ?`
		args = append(args, prefectureID)

		if len(businessTypeIDs) > 0 {
			placeholders := strings.Repeat("?,", len(businessTypeIDs))
			placeholders = placeholders[:len(placeholders)-1]
			sql += " AND s.business_type_id IN (" + placeholders + ")"
			for _, id := range businessTypeIDs {
				args = append(args, id)
			}
		}
		if len(contractPlanIDs) > 0 {
			placeholders := strings.Repeat("?,", len(contractPlanIDs))
			placeholders = placeholders[:len(placeholders)-1]
			sql += " AND s.contract_plan_id IN (" + placeholders + ")"
			for _, id := range contractPlanIDs {
				args = append(args, id)
			}
		}
		sql += " ORDER BY d.id"
	} else {
		sql = `SELECT id, name, prefecture_id FROM districts WHERE prefecture_id = ? ORDER BY id`
		args = append(args, prefectureID)
	}

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(sql, args...).Scan(&rows).Error; err != nil {
		return nil, err
	}

	result := make([]querymodel.DistrictQuery, 0, len(rows))
	for _, row := range rows {
		result = append(result, &querymodel.District{
			ID:           helper.ToUint(row["id"]),
			Name:         helper.ToString(row["name"]),
			PrefectureID: helper.ToUint(row["prefecture_id"]),
		})
	}
	return result, nil
}
