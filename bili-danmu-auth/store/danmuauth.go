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

func toUpdateParams(danmuAuth *core.DanmuAuth) map[string]interface{} {
	return map[string]interface{}{
		"buid":           danmuAuth.Buid,
		"uuid":           danmuAuth.UUID,
		"v_code":         danmuAuth.VCode,
		"verified_count": danmuAuth.VerifiedCount,
	}
}

func update(db *db.DB, danmuAuth *core.DanmuAuth) (int64, error) {
	updates := toUpdateParams(danmuAuth)
	tx := db.Update().Model(danmuAuth).Where("id = ?", danmuAuth.ID).Updates(updates)
	return tx.RowsAffected, tx.Error
}

func (s *danmuAuthStore) Save(ctx context.Context, danmuAuth *core.DanmuAuth) error {
	return s.db.Tx(func(tx *db.DB) error {
		var rows int64
		var err error
		rows, err = update(tx, danmuAuth)
		if err != nil {
			return err
		}

		if rows == 0 {
			return tx.Update().Create(danmuAuth).Error
		}

		return nil
	})
}

func (s *danmuAuthStore) SaveVCode(ctx context.Context, danmuAuth *core.DanmuAuth, vcode string) error {
	return s.db.Tx(func(tx *db.DB) error {
		newTx := tx.Update().Model(danmuAuth).Update("v_code", vcode)
		if newTx.Error != nil {
			return newTx.Error
		}

		if newTx.RowsAffected == 0 {
			return tx.Update().Create(danmuAuth).Error
		}

		return nil
	})
}

func (s *danmuAuthStore) SaveVerifiedCount(ctx context.Context, danmuAuth *core.DanmuAuth, verifiedCount int) error {
	return s.db.Tx(func(tx *db.DB) error {
		newTx := tx.Update().Model(danmuAuth).Update("verified_count", verifiedCount)
		if newTx.Error != nil {
			return newTx.Error
		}

		if newTx.RowsAffected == 0 {
			return tx.Update().Create(danmuAuth).Error
		}

		return nil
	})
}

func (s *danmuAuthStore) Delete(ctx context.Context, id uint) error {
	return s.db.Tx(func(tx *db.DB) error {
		return tx.Update().WithContext(ctx).Delete(&core.DanmuAuth{}, id).Error
	})
}

func (s *danmuAuthStore) FindByUUIDBuidVCode(ctx context.Context, uuid string, buid uint, vCode string) (*core.DanmuAuth, error) {
	var danmuAuth core.DanmuAuth
	_10minAgo := time.Now().Add(-10 * time.Minute)
	err := s.db.View().WithContext(ctx).Where("uuid = ? AND buid = ? AND v_code = ? AND created_at > ?", uuid, buid, vCode, _10minAgo).Last(&danmuAuth).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &danmuAuth, nil
}

func (s *danmuAuthStore) FindByBuidVCode(ctx context.Context, buid uint, VCode string) (*core.DanmuAuth, error) {
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

func (s *danmuAuthStore) FindByUUIDBuid(ctx context.Context, uuid string, buid uint) (*core.DanmuAuth, error) {
	var danmuAuth core.DanmuAuth
	_10minAgo := time.Now().Add(-10 * time.Minute)
	err := s.db.View().WithContext(ctx).Where("uuid = ? AND buid = ? AND created_at > ?", uuid, buid, _10minAgo).Last(&danmuAuth).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &danmuAuth, nil
}

func (s *danmuAuthStore) FindByBuid(ctx context.Context, buid uint) (*core.DanmuAuth, error) {
	var danmuAuth *core.DanmuAuth
	_10minAgo := time.Now().Add(-10 * time.Minute)
	err := s.db.View().WithContext(ctx).Where("buid = ? AND created_at > ?", buid, _10minAgo).Last(&danmuAuth).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return danmuAuth, nil
}
