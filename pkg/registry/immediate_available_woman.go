package registry

import (
	"golang-trainning-backend/pkg/adapter/controller"
	"golang-trainning-backend/pkg/adapter/repository"
	"golang-trainning-backend/pkg/usecase/interactor"
)

func (r *registry) NewImmediateAvailableWomanController() controller.ImmediateAvailableWoman {
	u := interactor.NewImmediateAvailableWomanUsecase(
		repository.NewStoreRepository(r.db),
		repository.NewWomanRepository(r.db),
	)
	return controller.NewImmediateAvailableWomanController(u)
}
