package store

import (
	"context"
	"errors"

	"github.com/leaper-one/pkg/db"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/core"
	"gorm.io/gorm"
)

func NewAccessKeyStore(db *db.DB) core.AccessKeyStore {
	return &accessKeyStore{db: db}
}

type accessKeyStore struct {
	db *db.DB
}

// toUpdateParams returns the params to update
func (a *accessKeyStore) toUpdateParams(accessKey *core.AccessKey) map[string]interface{} {
	return map[string]interface{}{
		"buid": accessKey.Buid,
		"key":  accessKey.Key,
	}
}

func (a *accessKeyStore) update(db *db.DB, accessKey *core.AccessKey) (int64, error) {
	updates := a.toUpdateParams(accessKey)
	tx := db.Update().Model(accessKey).Where("id = ?", accessKey.ID).Updates(updates)
	return tx.RowsAffected, tx.Error
}

func (a *accessKeyStore) Save(ctx context.Context, accessKey *core.AccessKey) error {
	return a.db.Tx(func(tx *db.DB) error {
		var rows int64
		var err error
		rows, err = a.update(tx, accessKey)
		if err != nil {
			return err
		}

		if rows == 0 {
			return tx.Update().Create(accessKey).Error
		}

		return nil
	})
}

func (a *accessKeyStore) Delete(ctx context.Context, buid int) error {
	return a.db.Update().Where("buid = ?", buid).Delete(&core.AccessKey{}).Error
}

func (a *accessKeyStore) FindByKey(ctx context.Context, key string) (*core.AccessKey, error) {
	var accessKey core.AccessKey
	err := a.db.Update().Where("key = ?", key).First(&accessKey).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &accessKey, err
}

func (a *accessKeyStore) ListByBuid(ctx context.Context, buid int) ([]*core.AccessKey, error) {
	var accessKeys []*core.AccessKey
	err := a.db.Update().Where("buid = ?", buid).Find(&accessKeys).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return accessKeys, err
}
