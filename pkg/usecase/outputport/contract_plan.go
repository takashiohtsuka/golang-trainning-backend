package outputport

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
)

type ContractPlanRepository interface {
	FindAll(ctx context.Context) ([]querymodel.ContractPlanQuery, error)
}
