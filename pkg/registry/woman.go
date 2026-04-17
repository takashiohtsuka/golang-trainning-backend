package registry

import (
	"golang-trainning-backend/pkg/adapter/controller"
	backendrepository "golang-trainning-backend/pkg/adapter/repository"
	"golang-trainning-backend/pkg/usecase/interactor"
	"golang-trainning-backend/pkg/adapter/repository"
)

func (r *registry) NewWomanController() controller.Woman {
	u := interactor.NewWomanUsecase(
		backendrepository.NewWomanRepository(r.db),
		backendrepository.NewCompanyRepository(r.db),
		backendrepository.NewStoreRepository(r.db),
		repository.NewUnitOfWork(r.db),
		r.storage,
	)
	return controller.NewWomanController(u)
}
