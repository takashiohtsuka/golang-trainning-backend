package controller

import (
	"net/http"

	"golang-trainning-backend/pkg/usecase/inputport"
)

type regionController struct {
	regionUsecase inputport.RegionUsecase
}

type Region interface {
	GetRegions(c Context) error
}

func NewRegionController(u inputport.RegionUsecase) Region {
	return &regionController{regionUsecase: u}
}

func (rc *regionController) GetRegions(ctx Context) error {
	regions, err := rc.regionUsecase.GetAll(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, regions)
}
