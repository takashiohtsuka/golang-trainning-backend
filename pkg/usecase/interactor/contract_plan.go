package interactor

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
	"golang-trainning-backend/pkg/usecase/inputport"
	"golang-trainning-backend/pkg/usecase/outputport"
)

type contractPlanUsecase struct {
	contractPlanRepository outputport.ContractPlanRepository
}

func NewContractPlanUsecase(r outputport.ContractPlanRepository) inputport.ContractPlanUsecase {
	return &contractPlanUsecase{contractPlanRepository: r}
}

func (u *contractPlanUsecase) GetAll(ctx context.Context) ([]querymodel.ContractPlanQuery, error) {
	return u.contractPlanRepository.FindAll(ctx)
}
