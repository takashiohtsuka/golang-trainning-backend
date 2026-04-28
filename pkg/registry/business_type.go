package registry

import (
	"golang-trainning-backend/pkg/adapter/controller"
	"golang-trainning-backend/pkg/adapter/repository"
	"golang-trainning-backend/pkg/usecase/interactor"
)

func (r *registry) NewBusinessTypeController() controller.BusinessType {
	u := interactor.NewBusinessTypeUsecase(
		repository.NewBusinessTypeRepository(r.db),
	)
	return controller.NewBusinessTypeController(u)
}
