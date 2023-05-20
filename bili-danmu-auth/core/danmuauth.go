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
		// ClientID is Client instance's ID
		ClientID string `gorm:"type:varchar(36)" json:"uuid"`
		// VCode is Client verification code
		VCode string `gorm:"type:varchar(20)" json:"v_code"`
		// IsVerified is Client verification status
		VerifiedCount int `gorm:"default:0;not null" json:"verified_count"`
	}
	// AuthStore ...
	DanmuAuthStore interface {
		Save(ctx context.Context, danmuAuth *DanmuAuth) error
		AddVerifiedCount(ctx context.Context, danmuAuth *DanmuAuth) error
		FindByBuidVCode(ctx context.Context, buid int, VCode string) (*DanmuAuth, error)
		FindByClientID(ctx context.Context, client_id string) (*DanmuAuth, error)
	}
)
