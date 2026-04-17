package controller

import (
	"net/http"

	requestCompanies "golang-trainning-backend/pkg/adapter/request/companies"
	"golang-trainning-backend/pkg/usecase/inputport"
)

type companyController struct {
	companyUsecase inputport.CompanyUsecase
}

type Company interface {
	CreateCompany(c Context) error
}

func NewCompanyController(u inputport.CompanyUsecase) Company {
	return &companyController{u}
}

func (cc *companyController) CreateCompany(ctx Context) error {
	var req requestCompanies.Post
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := cc.companyUsecase.Create(ctx.Request().Context(), req.ToInput()); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, nil)
}
