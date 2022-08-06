package main

import (
	"flag"
	"fmt"

	"air-grating-pms-backend/api/client/internal/config"
	"air-grating-pms-backend/api/client/internal/handler"
	"air-grating-pms-backend/api/client/internal/svc"
	"air-grating-pms-backend/common/middleware"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/client-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	server.Use(middleware.AuthenticationMiddleware(&c.AuthRpc))
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
