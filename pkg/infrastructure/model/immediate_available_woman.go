package model

import "time"

/*
GORMでマッピングされたModelの構造体
*/
type ImmediateAvailableWoman struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	StoreID   int       `json:"store_id"`
	WomanID   int       `json:"woman_id"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (ImmediateAvailableWoman) TableName() string { return "immediate_available_women" }
