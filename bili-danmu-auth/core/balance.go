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
		Buid int `gorm:"type:int;default:0;not null;unique" json:"buid"`
		// Balance is the balence of the developer
		Balance int `gorm:"default:100;not null" json:"balance"`
	}

	// BalanceStore interface
	BalanceStore interface {
		Save(ctx context.Context, balance *Balance) error
		// FindByBuid returns the balance of the buid
		FindByBuid(ctx context.Context, buid int) (*Balance, error)
		// DecrBalance decr the balance by buid
		DecrBalance(ctx context.Context, balance *Balance, amount int) error
		// Charge adds amount to the balance by buid
		Charge(ctx context.Context, balance *Balance, amount int) error
	}
)
