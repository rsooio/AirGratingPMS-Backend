package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	AuthRpc       zrpc.RpcClientConf
	StafferRpc    zrpc.RpcClientConf
	WorkshopRpc   zrpc.RpcClientConf
	EnterpriseRpc zrpc.RpcClientConf
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
