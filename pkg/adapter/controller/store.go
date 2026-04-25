package controller

import (
	"errors"
	"net/http"
	"strconv"

	requestStores "golang-trainning-backend/pkg/adapter/request/stores"
	"golang-trainning-backend/pkg/apperror"
	"golang-trainning-backend/pkg/usecase/inputport"
)


type storeController struct {
	storeUsecase      inputport.StoreUsecase
	storeQueryUsecase inputport.StoreQueryUsecase
}

type Store interface {
	GetStore(c Context) error
	CreateStore(c Context) error
	UpdateStore(c Context) error
}

func NewStoreController(u inputport.StoreUsecase, q inputport.StoreQueryUsecase) Store {
	return &storeController{u, q}
}

func (sc *storeController) GetStore(ctx Context) error {
	var req requestStores.Get
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	store, err := sc.storeQueryUsecase.GetByID(ctx.Request().Context(), req.ID)
	if err != nil {
		var notFound *apperror.NotFoundException
		if errors.As(err, &notFound) {
			return ctx.JSON(http.StatusNotFound, map[string]string{"error": notFound.Error()})
		}
		return err
	}

	return ctx.JSON(http.StatusOK, store)
}

func (sc *storeController) CreateStore(ctx Context) error {
	var req requestStores.Post
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := sc.storeUsecase.Create(ctx.Request().Context(), req.ToInput()); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, nil)
}

func (sc *storeController) UpdateStore(ctx Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	var req requestStores.Put
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := sc.storeUsecase.Update(ctx.Request().Context(), req.ToInput(uint(id))); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, nil)
}
