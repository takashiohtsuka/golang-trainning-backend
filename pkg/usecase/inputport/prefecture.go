package inputport

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
)

type PrefectureUsecase interface {
	GetByRegionID(ctx context.Context, regionID uint) ([]querymodel.PrefectureQuery, error)
}
