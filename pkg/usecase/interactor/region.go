package interactor

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
	"golang-trainning-backend/pkg/usecase/inputport"
	"golang-trainning-backend/pkg/usecase/outputport"
)

type regionUsecase struct {
	regionRepository outputport.RegionRepository
}

func NewRegionUsecase(r outputport.RegionRepository) inputport.RegionUsecase {
	return &regionUsecase{regionRepository: r}
}

func (u *regionUsecase) GetAll(ctx context.Context) ([]querymodel.RegionQuery, error) {
	return u.regionRepository.FindAll(ctx)
}
