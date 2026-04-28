package controller

import (
	"net/http"

	"golang-trainning-backend/pkg/usecase/inputport"
)

type businessTypeController struct {
	businessTypeUsecase inputport.BusinessTypeUsecase
}

type BusinessType interface {
	GetBusinessTypes(c Context) error
}

func NewBusinessTypeController(u inputport.BusinessTypeUsecase) BusinessType {
	return &businessTypeController{businessTypeUsecase: u}
}

func (bc *businessTypeController) GetBusinessTypes(ctx Context) error {
	businessTypes, err := bc.businessTypeUsecase.GetAll(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, businessTypes)
}
