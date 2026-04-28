package inputport

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
)

type ContractPlanUsecase interface {
	GetAll(ctx context.Context) ([]querymodel.ContractPlanQuery, error)
}
