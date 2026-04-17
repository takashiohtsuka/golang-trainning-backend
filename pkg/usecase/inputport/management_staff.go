package inputport

import (
	"context"

	"golang-trainning-backend/pkg/usecase/input"
)

type ManagementStaffUsecase interface {
	Create(ctx context.Context, i input.CreateManagementStaffInput) error
}
