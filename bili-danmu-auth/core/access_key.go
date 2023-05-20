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
		Buid uint `gorm:"type:int;default:0;not null" json:"buid"`
		// Key is the key, UUID
		Key string `gorm:"type:varchar(36);not null" json:"key"`
		// IsEnabled is the key's status
		IsEnabled bool `gorm:"default:false;not null" json:"is_enabled"`
	}

	// AccessKeyStore interface
	AccessKeyStore interface {
		Save(ctx context.Context, accessKey *AccessKey) error
		Delete(ctx context.Context, id uint) error
		FindByKey(ctx context.Context, key string) (*AccessKey, error)
		ListByBuid(ctx context.Context, buid uint) ([]*AccessKey, error)
	}
)
