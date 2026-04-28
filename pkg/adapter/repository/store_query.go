package repository

import (
	"context"

	storeMapper "golang-trainning-backend/pkg/adapter/mapper/store"
	"golang-trainning-backend/pkg/querymodel"
	"golang-trainning-backend/pkg/usecase/outputport"
	"golang-trainning-backend/pkg/usecase/query"

	"gorm.io/gorm"
)

type storeQueryRepository struct {
	db *gorm.DB
}

func NewStoreQueryRepository(db *gorm.DB) outputport.StoreQueryRepository {
	return &storeQueryRepository{db: db}
}

const storeQueryBaseSQL = `
	SELECT
		s.id,
		s.company_id,
		s.name,
		s.is_active,
		s.open_status,
		bt.code  AS business_type_code,
		cp.code  AS contract_plan_code,
		d.id     AS district_id,
		d.name   AS district_name,
		p.id     AS prefecture_id,
		p.name   AS prefecture_name,
		r.id     AS region_id,
		r.name   AS region_name
	FROM stores s
	JOIN business_types bt  ON s.business_type_id  = bt.id
	JOIN contract_plans cp  ON s.contract_plan_id  = cp.id
	JOIN districts d        ON s.district_id       = d.id
	JOIN prefectures p      ON d.prefecture_id     = p.id
	JOIN regions r          ON p.region_id         = r.id
	WHERE s.deleted_at IS NULL`

func (r *storeQueryRepository) FindByID(ctx context.Context, id uint) (querymodel.StoreQuery, error) {
	sql := storeQueryBaseSQL + " AND s.id = ?"
	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(sql, id).Scan(&rows).Error; err != nil {
		return &querymodel.NilStoreQueryModel{}, err
	}
	if len(rows) == 0 {
		return &querymodel.NilStoreQueryModel{}, nil
	}
	return storeMapper.ToQueryModel(rows[0]), nil
}

func (r *storeQueryRepository) FindAll(ctx context.Context, conditions []query.Condition) ([]querymodel.StoreQuery, error) {
	whereClause, args := buildWhereClause(conditions)
	sql := storeQueryBaseSQL + whereClause

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(sql, args...).Scan(&rows).Error; err != nil {
		return nil, err
	}

	result := make([]querymodel.StoreQuery, 0, len(rows))
	for _, row := range rows {
		result = append(result, storeMapper.ToQueryModel(row))
	}
	return result, nil
}
