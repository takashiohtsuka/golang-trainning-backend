package registry

import (
	"golang-trainning-backend/pkg/adapter/controller"
	backendoutputport "golang-trainning-backend/pkg/usecase/outputport"

	"gorm.io/gorm"
)

type registry struct {
	db      *gorm.DB
	storage backendoutputport.StorageRepository
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(db *gorm.DB, storage backendoutputport.StorageRepository) Registry {
	return &registry{db, storage}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Company:         r.NewCompanyController(),
		Store:           r.NewStoreController(),
		ManagementStaff: r.NewManagementStaffController(),
		Woman:           r.NewWomanController(),
		Region:          r.NewRegionController(),
		Prefecture:      r.NewPrefectureController(),
		District:        r.NewDistrictController(),
		BusinessType:            r.NewBusinessTypeController(),
		ContractPlan:            r.NewContractPlanController(),
		ImmediateAvailableWoman: r.NewImmediateAvailableWomanController(),
	}
}
