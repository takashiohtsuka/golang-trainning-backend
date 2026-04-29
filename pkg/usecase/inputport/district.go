package inputport

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
)

type DistrictUsecase interface {
	GetByPrefectureID(ctx context.Context, prefectureID uint) ([]querymodel.DistrictQuery, error)
}
