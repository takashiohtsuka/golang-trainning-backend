package registry

import (
	"golang-trainning-backend/pkg/adapter/controller"
	"golang-trainning-backend/pkg/adapter/repository"
	"golang-trainning-backend/pkg/usecase/interactor"
)

func (r *registry) NewDistrictController() controller.District {
	u := interactor.NewDistrictUsecase(
		repository.NewDistrictRepository(r.db),
	)
	return controller.NewDistrictController(u)
}
