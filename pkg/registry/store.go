package registry

import (
	"golang-trainning-backend/pkg/adapter/controller"
	backendrepository "golang-trainning-backend/pkg/adapter/repository"
	"golang-trainning-backend/pkg/usecase/interactor"
	"golang-trainning-backend/pkg/adapter/repository"
)

func (r *registry) NewStoreController() controller.Store {
	u := interactor.NewStoreUsecase(
		backendrepository.NewStoreRepository(r.db),
		backendrepository.NewCompanyRepository(r.db),
		repository.NewUnitOfWork(r.db),
	)
	q := interactor.NewStoreQueryUsecase(
		backendrepository.NewStoreQueryRepository(r.db),
	)
	return controller.NewStoreController(u, q)
}
