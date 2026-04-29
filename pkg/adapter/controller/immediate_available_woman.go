package controller

import (
	"errors"
	"net/http"

	requestIAW "golang-trainning-backend/pkg/adapter/request/immediate_available_women"
	"golang-trainning-backend/pkg/apperror"
	"golang-trainning-backend/pkg/usecase/inputport"
)

type immediateAvailableWomanController struct {
	usecase inputport.ImmediateAvailableWomanUsecase
}

type ImmediateAvailableWoman interface {
	Create(c Context) error
	Delete(c Context) error
}

func NewImmediateAvailableWomanController(u inputport.ImmediateAvailableWomanUsecase) ImmediateAvailableWoman {
	return &immediateAvailableWomanController{usecase: u}
}

func (ic *immediateAvailableWomanController) Create(ctx Context) error {
	var req requestIAW.Post
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := ic.usecase.Create(ctx.Request().Context(), req.StoreID, req.WomanID); err != nil {
		return handleUsecaseError(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, nil)
}

func (ic *immediateAvailableWomanController) Delete(ctx Context) error {
	var req requestIAW.Delete
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := ic.usecase.Delete(ctx.Request().Context(), req.StoreID, req.WomanID); err != nil {
		return handleUsecaseError(ctx, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func handleUsecaseError(ctx Context, err error) error {
	var notFound *apperror.NotFoundException
	if errors.As(err, &notFound) {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": notFound.Error()})
	}
	return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
}
