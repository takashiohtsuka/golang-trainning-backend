package outputport

import (
	"context"

	"golang-trainning-backend/pkg/querymodel"
)

type RegionRepository interface {
	FindAll(ctx context.Context) ([]querymodel.RegionQuery, error)
}
