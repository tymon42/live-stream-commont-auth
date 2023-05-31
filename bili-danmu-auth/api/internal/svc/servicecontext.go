package svc

import (
	"time"

	pkg_db "github.com/leaper-one/pkg/db"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/api/internal/config"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/core"
	"github.com/tymon42/live-stream-commont-auth/bili-danmu-auth/store"

	limiter "github.com/leaper-one/pkg/limiter"
)

type ServiceContext struct {
	Config      config.Config
	DanmuAuthDB core.DanmuAuthStore
	BalanceDB   core.BalanceStore
	AccessKeyDB core.AccessKeyStore
	Limiter     *limiter.SlidingWindowLimiter
}

var (
	second time.Duration = 1000000000
)

func NewServiceContext(c config.Config, db_path *string) *ServiceContext {
	db, _ := pkg_db.InitSQLiteDB(*db_path, &core.DanmuAuth{}, &core.Balance{}, &core.AccessKey{})
	// db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	// 1 second in nanoseconds

	// init SlidingWindowLimiter
	limiter, err := limiter.NewSlidingWindowLimiter(500, 60*second, 1)
	if err != nil {
		return nil
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
		AccessKeyDB: store.NewAccessKeyStore(&pkg_db.DB{
			Write: db,
			Read:  db,
		}),
		Limiter: limiter,
	}
}
