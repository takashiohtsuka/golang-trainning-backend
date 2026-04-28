package inputport

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
)

type DistrictUsecase interface {
	GetByPrefectureID(ctx context.Context, prefectureID uint, businessTypeIDs []uint, contractPlanIDs []uint) ([]querymodel.DistrictQuery, error)
}
