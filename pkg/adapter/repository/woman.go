package repository

import (
	"context"
	"strings"

	womanMapper "golang-trainning-backend/pkg/adapter/mapper/woman"
	"golang-trainning-backend/pkg/domain/entity"
	"golang-trainning-backend/pkg/usecase/outputport"
	"golang-trainning-backend/pkg/domain/collection"
	"golang-trainning-backend/pkg/helper"
	"golang-trainning-backend/pkg/infrastructure/model"
	"golang-trainning-backend/pkg/usecase/query"

	"gorm.io/gorm"
)

type womanRepository struct {
	db *gorm.DB
}

func NewWomanRepository(db *gorm.DB) outputport.WomanRepository {
	return &womanRepository{db: db}
}

const womanJoinSQL = `
	SELECT
		w.id, w.company_id, w.name, w.age, w.birthplace, w.blood_type, w.hobby, w.is_active,
		w.created_at, w.updated_at, w.deleted_at,
		wi.id AS wi_id, wi.woman_id AS wi_woman_id, wi.path AS wi_path,
		wi.created_at AS wi_created_at, wi.updated_at AS wi_updated_at,
		iaw.id AS iaw_id, iaw.store_id AS iaw_store_id, iaw.woman_id AS iaw_woman_id,
		iaw.expires_at AS iaw_expires_at, iaw.created_at AS iaw_created_at
	FROM women w
	LEFT JOIN woman_store_assignments wsa ON wsa.woman_id = w.id
	LEFT JOIN woman_images wi ON wi.woman_id = w.id
	LEFT JOIN immediate_available_women iaw ON iaw.woman_id = w.id AND iaw.expires_at > NOW()
	WHERE w.deleted_at IS NULL`

func buildWomanWhereClause(conditions []query.Condition) (string, []any) {
	tablePrefix := func(column string) string {
		switch column {
		case "store_id":
			return "wsa." + column
		default:
			return "w." + column
		}
	}

	var clauses []string
	var args []any

	for _, c := range conditions {
		col := tablePrefix(c.Column)
		switch c.Kind {
		case query.KindWhere:
			clauses = append(clauses, col+" = ?")
			args = append(args, c.Value)
		case query.KindWhereIn:
			clauses = append(clauses, col+" IN ?")
			args = append(args, c.Value)
		case query.KindWhereBetween:
			clauses = append(clauses, col+" BETWEEN ? AND ?")
			args = append(args, c.From, c.To)
		case query.KindWhereNotIn:
			clauses = append(clauses, col+" NOT IN ?")
			args = append(args, c.Value)
		}
	}

	if len(clauses) == 0 {
		return "", nil
	}
	return " AND " + strings.Join(clauses, " AND "), args
}

func groupWomen(rows []map[string]any) []*entity.Woman {
	womanMap := make(map[uint]*entity.Woman)
	womanOrder := make([]uint, 0)
	imageSeen := make(map[uint]map[uint]bool)
	imagesByID := make(map[uint][]entity.WomanImage)

	for _, row := range rows {
		womanID := helper.ToUint(row["id"])

		if _, ok := womanMap[womanID]; !ok {
			womanMap[womanID] = &entity.Woman{
				ID:        womanID,
				CompanyID: helper.ToUint(row["company_id"]),
				Name:       helper.ToString(row["name"]),
				Age:        helper.ToIntPtr(row["age"]),
				Birthplace: helper.ToStringPtr(row["birthplace"]),
				BloodType:  helper.ToStringPtr(row["blood_type"]),
				Hobby:      helper.ToStringPtr(row["hobby"]),
				IsActive:   helper.ToBool(row["is_active"]),
				CreatedAt:  helper.ToTimePtr(row["created_at"]),
				UpdatedAt:  helper.ToTimePtr(row["updated_at"]),
				DeletedAt:  helper.ToTimePtr(row["deleted_at"]),
			}
			imageSeen[womanID] = make(map[uint]bool)
			womanOrder = append(womanOrder, womanID)
		}

		imageID := helper.ToUint(row["wi_id"])
		if imageID != 0 && !imageSeen[womanID][imageID] {
			imageSeen[womanID][imageID] = true
			imagesByID[womanID] = append(imagesByID[womanID], entity.WomanImage{
				ID:        imageID,
				WomanID:   helper.ToUint(row["wi_woman_id"]),
				Path:      helper.ToString(row["wi_path"]),
				CreatedAt: helper.ToTimePtr(row["wi_created_at"]),
				UpdatedAt: helper.ToTimePtr(row["wi_updated_at"]),
			})
		}

		if iawID := helper.ToUint(row["iaw_id"]); iawID != 0 {
			womanMap[womanID].ImmediateAvailableWoman = &entity.ImmediateAvailableWoman{
				ID:        iawID,
				StoreID:   helper.ToUint(row["iaw_store_id"]),
				WomanID:   helper.ToUint(row["iaw_woman_id"]),
				ExpiresAt: helper.ToTime(row["iaw_expires_at"]),
				CreatedAt: helper.ToTime(row["iaw_created_at"]),
			}
		}
	}

	result := make([]*entity.Woman, len(womanOrder))
	for i, id := range womanOrder {
		w := womanMap[id]
		w.Images = collection.NewCollection(imagesByID[id])
		result[i] = w
	}
	return result
}

func (r *womanRepository) FindAll(ctx context.Context, conditions []query.Condition) (collection.Collection[entity.WomanEntity], error) {
	where, args := buildWomanWhereClause(conditions)

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(womanJoinSQL+where, args...).Scan(&rows).Error; err != nil {
		return collection.NewCollection[entity.WomanEntity](nil), err
	}

	women := groupWomen(rows)
	items := make([]entity.WomanEntity, len(women))
	for i, w := range women {
		items[i] = w
	}
	return collection.NewCollection(items), nil
}

func (r *womanRepository) FindOne(ctx context.Context, conditions []query.Condition) (entity.WomanEntity, error) {
	where, args := buildWomanWhereClause(conditions)

	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(womanJoinSQL+where, args...).Scan(&rows).Error; err != nil {
		return &entity.NilWoman{}, err
	}

	women := groupWomen(rows)
	if len(women) == 0 {
		return &entity.NilWoman{}, nil
	}
	return women[0], nil
}

func (r *womanRepository) Create(ctx context.Context, w *entity.Woman) (uint, error) {
	m := womanMapper.ToOrmModel(w)
	if err := r.db.WithContext(ctx).Create(m).Error; err != nil {
		return 0, err
	}
	return m.ID, nil
}

func (r *womanRepository) Update(ctx context.Context, w *entity.Woman) error {
	return r.db.WithContext(ctx).Save(womanMapper.ToOrmModel(w)).Error
}

func (r *womanRepository) SaveImage(ctx context.Context, womanID uint, path string) error {
	return r.db.WithContext(ctx).Create(&model.WomanImage{WomanID: womanID, Path: path}).Error
}

func (r *womanRepository) FindActiveImmediateAvailableByStore(ctx context.Context, storeID uint) ([]*entity.ImmediateAvailableWoman, error) {
	var rows []map[string]any
	if err := r.db.WithContext(ctx).Raw(
		`SELECT id, store_id, woman_id, expires_at, created_at
		 FROM immediate_available_women
		 WHERE store_id = ? AND expires_at > NOW()
		 ORDER BY id`,
		storeID,
	).Scan(&rows).Error; err != nil {
		return nil, err
	}

	result := make([]*entity.ImmediateAvailableWoman, 0, len(rows))
	for _, row := range rows {
		result = append(result, &entity.ImmediateAvailableWoman{
			ID:        helper.ToUint(row["id"]),
			StoreID:   helper.ToUint(row["store_id"]),
			WomanID:   helper.ToUint(row["woman_id"]),
			ExpiresAt: helper.ToTime(row["expires_at"]),
			CreatedAt: helper.ToTime(row["created_at"]),
		})
	}
	return result, nil
}

func (r *womanRepository) CreateImmediateAvailable(ctx context.Context, iaw *entity.ImmediateAvailableWoman) error {
	return r.db.WithContext(ctx).Exec(
		`INSERT INTO immediate_available_women (store_id, woman_id, expires_at) VALUES (?, ?, ?)`,
		iaw.StoreID, iaw.WomanID, iaw.ExpiresAt,
	).Error
}

func (r *womanRepository) DeleteImmediateAvailable(ctx context.Context, storeID, womanID uint) error {
	return r.db.WithContext(ctx).Exec(
		`DELETE FROM immediate_available_women WHERE store_id = ? AND woman_id = ?`,
		storeID, womanID,
	).Error
}
