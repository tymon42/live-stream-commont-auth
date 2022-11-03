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
		Buid uint `gorm:"type:int;default:0;not null" json:"buid"`
		// UUID is Client instance's ID
		UUID string `gorm:"type:varchar(100)" json:"uuid"`
		// VCode is Client verification code
		VCode string `gorm:"type:varchar(100)" json:"v_code"`
		// IsVerified is Client verification status
		VerifiedCount int `gorm:"default:0;not null" json:"verified_count"`
	}
	// AuthStore ...
	DanmuAuthStore interface {
		Save(ctx context.Context, danmuAuth *DanmuAuth) error
		Delete(ctx context.Context, id uint) error
		FindByUUIDBuidVCode(ctx context.Context, uuid string, buid uint, vCode string) (*DanmuAuth, error)
		FindByBuidVCode(ctx context.Context, buid uint, VCode string) (*DanmuAuth, error)
		FindByBuid(ctx context.Context, buid uint) (*DanmuAuth, error)
	}
)
