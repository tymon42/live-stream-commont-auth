package svc

import (
	"github.com/glebarez/sqlite"
	pkg_db "github.com/leaper-one/pkg/db"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/config"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/core"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/store"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	DanmuAuthDB core.DanmuAuthStore
	BalanceDB   core.BalanceStore
}

func NewServiceContext(c config.Config) *ServiceContext {
	// In-Memory Databases: https://www.sqlite.org/inmemorydb.html
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		logx.Errorf("open db failed, err: %v", err)
	}
	err = db.AutoMigrate(&core.DanmuAuth{}, &core.Balance{})
	if err != nil {
		logx.Infof("auto migrate failed, err: %v", err)
	}
	return &ServiceContext{
		Config: c,
		DanmuAuthDB: store.NewDanmuAuthStore(&pkg_db.DB{
			Write: db,
			Read:  db,
		}),
		BalanceDB: store.NewBalanceStore(&pkg_db.DB{
			Write: db,
			Read:  db,
		}),
	}
}
