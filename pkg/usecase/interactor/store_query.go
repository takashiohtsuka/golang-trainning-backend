package interactor

import (
	"context"

	"golang-trainning-backend/pkg/apperror"
	"golang-trainning-backend/pkg/querymodel"
	"golang-trainning-backend/pkg/usecase/inputport"
	"golang-trainning-backend/pkg/usecase/outputport"
)

type storeQueryUsecase struct {
	storeQueryRepository outputport.StoreQueryRepository
}

func NewStoreQueryUsecase(r outputport.StoreQueryRepository) inputport.StoreQueryUsecase {
	return &storeQueryUsecase{storeQueryRepository: r}
}

func (u *storeQueryUsecase) GetByID(ctx context.Context, id uint) (querymodel.StoreQuery, error) {
	store, err := u.storeQueryRepository.FindByID(ctx, id)
	if err != nil {
		return &querymodel.NilStoreQueryModel{}, err
	}
	if store.IsNil() {
		return &querymodel.NilStoreQueryModel{}, apperror.NewNotFoundException("store not found")
	}
	return store, nil
}
