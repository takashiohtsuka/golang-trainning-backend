package interactor

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
	"golang-trainning-backend/pkg/usecase/inputport"
	"golang-trainning-backend/pkg/usecase/outputport"
)

type prefectureUsecase struct {
	prefectureRepository outputport.PrefectureRepository
}

func NewPrefectureUsecase(r outputport.PrefectureRepository) inputport.PrefectureUsecase {
	return &prefectureUsecase{prefectureRepository: r}
}

func (u *prefectureUsecase) GetByRegionID(ctx context.Context, regionID uint, businessTypeIDs []uint, contractPlanIDs []uint) ([]querymodel.PrefectureQuery, error) {
	return u.prefectureRepository.FindByRegionID(ctx, regionID, businessTypeIDs, contractPlanIDs)
}
