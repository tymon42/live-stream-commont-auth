package store

import (
	"context"
	"errors"

	"github.com/leaper-one/pkg/db"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/core"
	"gorm.io/gorm"
)

func NewBalanceStore(db *db.DB) core.BalenceStore {
	return &balanceStore{db: db}
}

type balanceStore struct {
	db *db.DB
}

// toUpdateParams returns the params to update
func (b *balanceStore) toUpdateParams(balance *core.Balance) map[string]interface{} {
	return map[string]interface{}{
		"buid":    balance.Buid,
		"balance": balance.Balance,
	}
}

func (b *balanceStore) update(db *db.DB, balance *core.Balance) (int64, error) {
	updates := b.toUpdateParams(balance)
	tx := db.Update().Model(balance).Where("id = ?", balance.ID).Updates(updates)
	return tx.RowsAffected, tx.Error
}

func (b *balanceStore) Save(ctx context.Context, balance *core.Balance) error {
	return b.db.Tx(func(tx *db.DB) error {
		var rows int64
		var err error
		rows, err = b.update(tx, balance)
		if err != nil {
			return err
		}

		if rows == 0 {
			return tx.Update().Create(balance).Error
		}

		return nil
	})
}

func (b *balanceStore) FindByBuid(ctx context.Context, buid uint64) (*core.Balance, error) {
	var balance core.Balance
	err := b.db.Update().Where("buid = ?", buid).First(&balance).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &balance, err
}

func (b *balanceStore) Charge(ctx context.Context, balance *core.Balance, amount int) error {
	balance.Balance += amount
	return b.db.Update().Model(balance).Where("id = ?", balance.ID).Update("balance", balance.Balance).Error
}
