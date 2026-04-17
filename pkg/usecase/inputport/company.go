package inputport

import (
	"context"

	"golang-trainning-backend/pkg/usecase/input"
)

type CompanyUsecase interface {
	Create(ctx context.Context, i input.CreateCompanyInput) error
}
