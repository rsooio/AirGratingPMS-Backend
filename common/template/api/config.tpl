package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	{{.authImport}}
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	AuthRpc    zrpc.RpcClientConf

	{{.auth}}
	{{.jwtTrans}}
}
