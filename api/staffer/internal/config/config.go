package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	DefaultPassword string
	AuthRPC         zrpc.RpcClientConf
	StafferRPC      zrpc.RpcClientConf
	CacheRedis      cache.CacheConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Mysql struct {
		DataSource string
	}
}
