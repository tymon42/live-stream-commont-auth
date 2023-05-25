package core

import (
	"context"

	"gorm.io/gorm"
)

type (
	// AccessKey represents a api access key for developers
	AccessKey struct {
		gorm.Model
		// Buid is Bilibili user's UID
		Buid int `gorm:"type:int;default:0;not null" json:"buid"`
		// Key is the key, UUID
		Key string `gorm:"type:varchar(36);not null" json:"key"`
	}

	// AccessKeyStore interface
	AccessKeyStore interface {
		Save(ctx context.Context, accessKey *AccessKey) error
		Delete(ctx context.Context, buid int, accessKey string) error
		FindByKey(ctx context.Context, key string) (*AccessKey, error)
		ListByBuid(ctx context.Context, buid int) ([]*AccessKey, error)
	}
)
