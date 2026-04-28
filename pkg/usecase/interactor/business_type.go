package interactor

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
	"golang-trainning-backend/pkg/usecase/inputport"
	"golang-trainning-backend/pkg/usecase/outputport"
)

type businessTypeUsecase struct {
	businessTypeRepository outputport.BusinessTypeRepository
}

func NewBusinessTypeUsecase(r outputport.BusinessTypeRepository) inputport.BusinessTypeUsecase {
	return &businessTypeUsecase{businessTypeRepository: r}
}

func (u *businessTypeUsecase) GetAll(ctx context.Context) ([]querymodel.BusinessTypeQuery, error) {
	return u.businessTypeRepository.FindAll(ctx)
}
