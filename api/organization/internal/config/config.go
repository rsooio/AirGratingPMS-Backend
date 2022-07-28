package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	AuthRPC       zrpc.RpcClientConf
	StafferRPC    zrpc.RpcClientConf
	WorkshopRPC   zrpc.RpcClientConf
	EnterpriseRPC zrpc.RpcClientConf
	CacheRedis    cache.CacheConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	DTM struct {
		Server string
	}
	Mysql struct {
		DataSource string
	}
}
