package entity

import "time"

const ImmediateAvailableExpiry = 10 * time.Minute

type ImmediateAvailableWoman struct {
	ID        uint
	StoreID   uint
	WomanID   uint
	ExpiresAt time.Time
	CreatedAt time.Time
}

func (i *ImmediateAvailableWoman) IsNil() bool     { return i.ID == 0 }
func (i *ImmediateAvailableWoman) GetID() uint      { return i.ID }
func (i *ImmediateAvailableWoman) GetStoreID() uint { return i.StoreID }
func (i *ImmediateAvailableWoman) GetWomanID() uint { return i.WomanID }
func (i *ImmediateAvailableWoman) IsExpired() bool  { return time.Now().After(i.ExpiresAt) }

type NilImmediateAvailableWoman struct{}

func (n *NilImmediateAvailableWoman) IsNil() bool     { return true }
func (n *NilImmediateAvailableWoman) GetID() uint      { return 0 }
func (n *NilImmediateAvailableWoman) GetStoreID() uint { return 0 }
func (n *NilImmediateAvailableWoman) GetWomanID() uint { return 0 }
func (n *NilImmediateAvailableWoman) IsExpired() bool  { return true }
