package outputport

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
)

type PrefectureRepository interface {
	FindByRegionID(ctx context.Context, regionID uint, businessTypeIDs []uint, contractPlanIDs []uint) ([]querymodel.PrefectureQuery, error)
}
