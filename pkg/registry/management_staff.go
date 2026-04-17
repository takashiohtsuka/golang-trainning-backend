package registry

import (
	"golang-trainning-backend/pkg/adapter/controller"
	backendrepository "golang-trainning-backend/pkg/adapter/repository"
	"golang-trainning-backend/pkg/usecase/interactor"
	"golang-trainning-backend/pkg/adapter/repository"
)

func (r *registry) NewManagementStaffController() controller.ManagementStaff {
	u := interactor.NewManagementStaffUsecase(
		backendrepository.NewManagementStaffRepository(r.db),
		backendrepository.NewCompanyRepository(r.db),
		backendrepository.NewStoreRepository(r.db),
		repository.NewUnitOfWork(r.db),
	)
	return controller.NewManagementStaffController(u)
}
