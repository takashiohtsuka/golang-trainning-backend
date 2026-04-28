package registry

import (
	"golang-trainning-backend/pkg/adapter/controller"
	"golang-trainning-backend/pkg/adapter/repository"
	"golang-trainning-backend/pkg/usecase/interactor"
)

func (r *registry) NewContractPlanController() controller.ContractPlan {
	u := interactor.NewContractPlanUsecase(
		repository.NewContractPlanRepository(r.db),
	)
	return controller.NewContractPlanController(u)
}
