package interactor

import (
	"context"
	"errors"
	"time"

	"golang-trainning-backend/pkg/apperror"
	"golang-trainning-backend/pkg/domain/entity"
	"golang-trainning-backend/pkg/usecase/inputport"
	"golang-trainning-backend/pkg/usecase/outputport"
	"golang-trainning-backend/pkg/usecase/query"
)

type immediateAvailableWomanUsecase struct {
	storeRepo outputport.StoreRepository
	womanRepo outputport.WomanRepository
}

func NewImmediateAvailableWomanUsecase(
	storeRepo outputport.StoreRepository,
	womanRepo outputport.WomanRepository,
) inputport.ImmediateAvailableWomanUsecase {
	return &immediateAvailableWomanUsecase{
		storeRepo: storeRepo,
		womanRepo: womanRepo,
	}
}

func (u *immediateAvailableWomanUsecase) Delete(ctx context.Context, storeID, womanID uint) error {
	return u.womanRepo.DeleteImmediateAvailable(ctx, storeID, womanID)
}

func (u *immediateAvailableWomanUsecase) Create(ctx context.Context, storeID, womanID uint) error {
	// 1. 店舗取得（契約プランコード）
	store, err := u.storeRepo.FindOne(ctx, []query.Condition{
		query.Where("id", storeID),
	})
	if err != nil {
		return err
	}
	if store.IsNil() {
		return apperror.NewNotFoundException("store not found")
	}

	// 2. 女性が店舗に所属しているか確認（非アクティブ除外）
	woman, err := u.womanRepo.FindOne(ctx, []query.Condition{
		query.Where("id", womanID),
		query.Where("store_id", storeID),
		query.Where("is_active", true),
	})
	if err != nil {
		return err
	}
	if woman.IsNil() {
		return apperror.NewNotFoundException("woman does not belong to this store or is not active")
	}

	// 3. 有効な即予約リスト取得
	activeList, err := u.womanRepo.FindActiveImmediateAvailableByStore(ctx, storeID)
	if err != nil {
		return err
	}

	// 4. 上限チェック
	if len(activeList) >= store.GetContractPlan().GetImmediateAvailableLimit() {
		return errors.New("immediate available limit reached for this contract plan")
	}

	// 5. 重複チェック
	for _, iaw := range activeList {
		if iaw.GetWomanID() == womanID {
			return errors.New("woman is already set as immediate available")
		}
	}

	// 6. 作成
	return u.womanRepo.CreateImmediateAvailable(ctx, &entity.ImmediateAvailableWoman{
		StoreID:   storeID,
		WomanID:   womanID,
		ExpiresAt: time.Now().Add(entity.ImmediateAvailableExpiry),
	})
}
