package inputport

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
)

type BusinessTypeUsecase interface {
	GetAll(ctx context.Context) ([]querymodel.BusinessTypeQuery, error)
}
