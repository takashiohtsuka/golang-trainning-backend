package interactor

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
	"golang-trainning-backend/pkg/usecase/inputport"
	"golang-trainning-backend/pkg/usecase/outputport"
)

type districtUsecase struct {
	districtRepository outputport.DistrictRepository
}

func NewDistrictUsecase(r outputport.DistrictRepository) inputport.DistrictUsecase {
	return &districtUsecase{districtRepository: r}
}

func (u *districtUsecase) GetByPrefectureID(ctx context.Context, prefectureID uint, businessTypeIDs []uint, contractPlanIDs []uint) ([]querymodel.DistrictQuery, error) {
	return u.districtRepository.FindByPrefectureID(ctx, prefectureID, businessTypeIDs, contractPlanIDs)
}
