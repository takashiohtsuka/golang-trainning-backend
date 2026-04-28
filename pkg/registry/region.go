package registry

import (
	"golang-trainning-backend/pkg/adapter/controller"
	"golang-trainning-backend/pkg/adapter/repository"
	"golang-trainning-backend/pkg/usecase/interactor"
)

func (r *registry) NewRegionController() controller.Region {
	u := interactor.NewRegionUsecase(
		repository.NewRegionRepository(r.db),
	)
	return controller.NewRegionController(u)
}
