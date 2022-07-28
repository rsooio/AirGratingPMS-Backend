package middleware

import (
	"air-grating-pms-backend/rpc/authentication/types/authentication"

	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/zrpc"
)

func AuthenticationMiddleware(conf *zrpc.RpcClientConf) rest.Middleware {
	authRpc := authentication.NewAuthenticationClient(zrpc.MustNewClient(*conf).Conn())
	fmt.Println("Middleware init")
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			role := ctx.Value("rol")
			if role == nil {
				role = "unknown"
			}
			fmt.Printf("role: %s\n", role)
			resp, err := authRpc.Enforce(r.Context(), &authentication.AuthReq{
				Subject: role.(string),
				Object:  r.URL.Path,
				Action:  r.Method,
			})
			if err != nil {
				httpx.Error(w, err)
				return
			}
			if !resp.Permission {
				w.WriteHeader(http.StatusUnauthorized)
				details, _ := httputil.DumpRequest(r, true)
				logx.Errorf("authorize failed: %s\n=> %+v", "access denied", string(details))
			}

			next(w, r.WithContext(ctx))
		}
	}
}
