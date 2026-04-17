package registry

import (
	"golang-trainning-backend/pkg/adapter/controller"
	backendrepository "golang-trainning-backend/pkg/adapter/repository"
	"golang-trainning-backend/pkg/usecase/interactor"
	"golang-trainning-backend/pkg/adapter/repository"
)

func (r *registry) NewCompanyController() controller.Company {
	u := interactor.NewCompanyUsecase(
		backendrepository.NewCompanyRepository(r.db),
		repository.NewUnitOfWork(r.db),
	)
	return controller.NewCompanyController(u)
}
