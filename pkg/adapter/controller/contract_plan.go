package controller

import (
	"net/http"

	"golang-trainning-backend/pkg/usecase/inputport"
)

type contractPlanController struct {
	contractPlanUsecase inputport.ContractPlanUsecase
}

type ContractPlan interface {
	GetContractPlans(c Context) error
}

func NewContractPlanController(u inputport.ContractPlanUsecase) ContractPlan {
	return &contractPlanController{contractPlanUsecase: u}
}

func (cc *contractPlanController) GetContractPlans(ctx Context) error {
	contractPlans, err := cc.contractPlanUsecase.GetAll(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, contractPlans)
}
