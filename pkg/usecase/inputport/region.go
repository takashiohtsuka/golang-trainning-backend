package inputport

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
)

type RegionUsecase interface {
	GetAll(ctx context.Context) ([]querymodel.RegionQuery, error)
}
