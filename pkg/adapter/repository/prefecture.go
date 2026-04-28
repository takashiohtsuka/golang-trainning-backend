package repository

import (
	"context"
	"strings"

	"golang-trainning-backend/pkg/helper"
	"golang-trainning-backend/pkg/querymodel"
	"golang-trainning-backend/pkg/usecase/outputport"

	"gorm.io/gorm"
)

type prefectureRepository struct {
	db *gorm.DB
}

func NewPrefectureRepository(db *gorm.DB) outputport.PrefectureRepository {
	return &prefectureRepository{db: db}
}

func (r *prefectureRepository) FindByRegionID(ctx context.Context, regionID uint, businessTypeIDs []uint, contractPlanIDs []uint) ([]querymodel.PrefectureQuery, error) {
	needsJoin := len(businessTypeIDs) > 0 || len(contractPlanIDs) > 0

	var sql string
	var args []any

	if needsJoin {
		sql = `SELECT DISTINCT p.id, p.name, p.region_id
FROM prefectures p
INNER JOIN districts d ON d.prefecture_id = p.id
INNER JOIN stores s ON s.district_id = d.id
WHERE p.region_id = ?`
		args = append(args, regionID)

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
		sql += " ORDER BY p.id"
	} else {
		sql = `SELECT id, name, region_id FROM prefectures WHERE region_id = ? ORDER BY id`
		args = append(args, regionID)
	}

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(sql, args...).Scan(&rows).Error; err != nil {
		return nil, err
	}

	result := make([]querymodel.PrefectureQuery, 0, len(rows))
	for _, row := range rows {
		result = append(result, &querymodel.Prefecture{
			ID:       helper.ToUint(row["id"]),
			Name:     helper.ToString(row["name"]),
			RegionID: helper.ToUint(row["region_id"]),
		})
	}
	return result, nil
}
