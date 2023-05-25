package svc

import (
	pkg_db "github.com/leaper-one/pkg/db"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/config"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/core"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/store"
)

type ServiceContext struct {
	Config      config.Config
	DanmuAuthDB core.DanmuAuthStore
	BalanceDB   core.BalanceStore
	AccessKeyDB core.AccessKeyStore
}

func NewServiceContext(c config.Config, db_path *string) *ServiceContext {
	db, _ := pkg_db.InitSQLiteDB(*db_path, &core.DanmuAuth{}, &core.Balance{}, &core.AccessKey{})
	// db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
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
		AccessKeyDB: store.NewAccessKeyStore(&pkg_db.DB{
			Write: db,
			Read:  db,
		}),
	}
}
