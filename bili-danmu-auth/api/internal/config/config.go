package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	DanmuAuth struct {
		VCodeExpire           int // in seconds
		DevloperVCodePrefix   string
		NormalUserVCodePrefix string
		InitialBalance        int // the initial balance for new devloper
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Worker struct {
		ApiKey string
	}
}
