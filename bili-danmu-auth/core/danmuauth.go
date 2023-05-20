package core

import (
	"context"

	"gorm.io/gorm"
)

type (
	// Auth ...
	DanmuAuth struct {
		gorm.Model
		// Buid is Bilibili user's UID
		Buid int `gorm:"type:int;default:0;not null" json:"buid"`
		// UUID is Client instance's ID
		UUID string `gorm:"type:varchar(36)" json:"uuid"`
		// VCode is Client verification code
		VCode string `gorm:"type:varchar(20)" json:"v_code"`
		// IsVerified is Client verification status
		VerifiedCount int `gorm:"default:0;not null" json:"verified_count"`
	}
	// AuthStore ...
	DanmuAuthStore interface {
		Save(ctx context.Context, danmuAuth *DanmuAuth) error
		SaveVCode(ctx context.Context, danmuAuth *DanmuAuth, vcode string) error
		SaveVerifiedCount(ctx context.Context, danmuAuth *DanmuAuth, verifiedCount int) error
		Delete(ctx context.Context, id uint) error
		FindByUUIDBuidVCode(ctx context.Context, uuid string, buid uint, vCode string) (*DanmuAuth, error)
		FindByUUIDBuid(ctx context.Context, uuid string, buid uint) (*DanmuAuth, error)
		FindByBuidVCode(ctx context.Context, buid uint, VCode string) (*DanmuAuth, error)
	}
)
