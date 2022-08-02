package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	AuthRpc       zrpc.RpcClientConf
	OrderRpc      zrpc.RpcClientConf
	ProductRpc    zrpc.RpcClientConf
	ProductSetRpc zrpc.RpcClientConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
