package outputport

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
)

type DistrictRepository interface {
	FindByPrefectureID(ctx context.Context, prefectureID uint) ([]querymodel.DistrictQuery, error)
}
