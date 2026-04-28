package controller

import (
	"net/http"

	requestPrefectures "golang-trainning-backend/pkg/adapter/request/prefectures"
	"golang-trainning-backend/pkg/usecase/inputport"
)

type prefectureController struct {
	prefectureUsecase inputport.PrefectureUsecase
}

type Prefecture interface {
	GetPrefectures(c Context) error
}

func NewPrefectureController(u inputport.PrefectureUsecase) Prefecture {
	return &prefectureController{prefectureUsecase: u}
}

func (pc *prefectureController) GetPrefectures(ctx Context) error {
	var req requestPrefectures.List
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	prefectures, err := pc.prefectureUsecase.GetByRegionID(ctx.Request().Context(), req.RegionID, req.BusinessTypeIDs, req.ContractPlanIDs)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, prefectures)
}
