package controller

import (
	"net/http"

	requestDistricts "golang-trainning-backend/pkg/adapter/request/districts"
	"golang-trainning-backend/pkg/usecase/inputport"
)

type districtController struct {
	districtUsecase inputport.DistrictUsecase
}

type District interface {
	GetDistricts(c Context) error
}

func NewDistrictController(u inputport.DistrictUsecase) District {
	return &districtController{districtUsecase: u}
}

func (dc *districtController) GetDistricts(ctx Context) error {
	var req requestDistricts.List
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	districts, err := dc.districtUsecase.GetByPrefectureID(ctx.Request().Context(), req.PrefectureID)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, districts)
}
