package registry

import (
	"golang-trainning-backend/pkg/adapter/controller"
	"golang-trainning-backend/pkg/adapter/repository"
	"golang-trainning-backend/pkg/usecase/interactor"
)

func (r *registry) NewPrefectureController() controller.Prefecture {
	u := interactor.NewPrefectureUsecase(
		repository.NewPrefectureRepository(r.db),
	)
	return controller.NewPrefectureController(u)
}
