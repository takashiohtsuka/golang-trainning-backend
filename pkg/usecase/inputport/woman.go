package inputport

import (
	"context"

	"golang-trainning-backend/pkg/usecase/input"
)

type WomanUsecase interface {
	Create(ctx context.Context, i input.CreateWomanInput) error
	Update(ctx context.Context, i input.UpdateWomanInput) error
}
