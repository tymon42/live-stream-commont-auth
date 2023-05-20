package core

import (
	"context"

	"gorm.io/gorm"
)

type (
	// Balance represents the balence of a developer
	Balance struct {
		gorm.Model
		// Buid is Bilibili user's UID
		Buid uint64 `gorm:"type:int;default:0;not null;unique" json:"buid"`
		// Balance is the balence of the developer
		Balance int `gorm:"default:100;not null" json:"balance"`
	}

	// BalenceStore interface
	BalenceStore interface {
		Save(ctx context.Context, balance *Balance) error
		// FindByBuid returns the balance of the buid
		FindByBuid(ctx context.Context, buid uint64) (*Balance, error)
		// Deposit adds amount to the balance by buid
		Deposit(ctx context.Context, balance *Balance, amount int) error
	}
)
