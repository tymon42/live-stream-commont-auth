package store

import (
	"context"
	"errors"
	"time"

	"github.com/leaper-one/pkg/db"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/core"
	"gorm.io/gorm"
)

func NewDanmuAuthStore(db *db.DB) core.DanmuAuthStore {
	return &danmuAuthStore{db: db}
}

type danmuAuthStore struct {
	db *db.DB
}

func (d *danmuAuthStore) toUpdateParams(danmuAuth *core.DanmuAuth) map[string]interface{} {
	return map[string]interface{}{
		"buid":           danmuAuth.Buid,
		"v_code":         danmuAuth.VCode,
		"client_id":      danmuAuth.ClientID,
		"verified_count": danmuAuth.VerifiedCount,
	}
}

func (d *danmuAuthStore) update(db *db.DB, danmuAuth *core.DanmuAuth) (int64, error) {
	updates := d.toUpdateParams(danmuAuth)
	tx := db.Update().Model(danmuAuth).Where("id = ?", danmuAuth.ID).Updates(updates)
	return tx.RowsAffected, tx.Error
}

func (s *danmuAuthStore) Save(ctx context.Context, danmuAuth *core.DanmuAuth) error {
	return s.db.Tx(func(tx *db.DB) error {
		var rows int64
		var err error
		rows, err = s.update(tx, danmuAuth)
		if err != nil {
			return err
		}

		if rows == 0 {
			return tx.Update().Create(danmuAuth).Error
		}

		return nil
	})
}

func (s *danmuAuthStore) AddVerifiedCount(ctx context.Context, danmuAuth *core.DanmuAuth) error {
	return s.db.Tx(func(tx *db.DB) error {
		newTx := tx.Update().Model(danmuAuth).Update("verified_count", danmuAuth.VerifiedCount+1)
		if newTx.Error != nil {
			return newTx.Error
		}

		if newTx.RowsAffected == 0 {
			return tx.Update().Create(danmuAuth).Error
		}

		return nil
	})
}

func (s *danmuAuthStore) FindByBuidVCode(ctx context.Context, buid int, VCode string) (*core.DanmuAuth, error) {
	var danmuAuth core.DanmuAuth
	_10minAgo := time.Now().Add(-10 * time.Minute)
	err := s.db.View().WithContext(ctx).Where("buid = ? AND v_code = ? AND created_at > ?", buid, VCode, _10minAgo).Last(&danmuAuth).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &danmuAuth, nil
}

func (s *danmuAuthStore) FindByClientID(ctx context.Context, client_id string) (*core.DanmuAuth, error) {
	var danmuAuth core.DanmuAuth
	_10minAgo := time.Now().Add(-10 * time.Minute)
	err := s.db.View().WithContext(ctx).Where("client_id = ? AND created_at > ?", client_id, _10minAgo).Last(&danmuAuth).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &danmuAuth, nil
}
