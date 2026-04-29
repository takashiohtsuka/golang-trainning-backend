package repository

import (
	"context"

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

func (r *prefectureRepository) FindByRegionID(ctx context.Context, regionID uint) ([]querymodel.PrefectureQuery, error) {
	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(
		`SELECT id, name, region_id FROM prefectures WHERE region_id = ? ORDER BY id`, regionID,
	).Scan(&rows).Error; err != nil {
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
