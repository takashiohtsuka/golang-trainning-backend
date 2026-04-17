package interactor

import (
	"context"

	"golang-trainning-backend/pkg/domain/entity"
	"golang-trainning-backend/pkg/usecase/input"
	"golang-trainning-backend/pkg/usecase/inputport"
	backendoutputport "golang-trainning-backend/pkg/usecase/outputport"
	"golang-trainning-backend/pkg/usecase/outputport"
)

type CompanyUsecase struct {
	companyRepository backendoutputport.CompanyRepository
	uow               outputport.UnitOfWork
}

func NewCompanyUsecase(
	companyRepository backendoutputport.CompanyRepository,
	uow outputport.UnitOfWork,
) inputport.CompanyUsecase {
	return &CompanyUsecase{companyRepository, uow}
}

func (u *CompanyUsecase) Create(ctx context.Context, i input.CreateCompanyInput) error {
	company := &entity.Company{
		Name:     i.Name,
		Rank:     i.Rank,
		IsActive: i.IsActive,
	}
	return u.uow.Do(ctx, func() error {
		return u.companyRepository.Create(ctx, company)
	})
}
