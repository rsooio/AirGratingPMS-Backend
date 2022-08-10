package utils

import "context"

func GetEnterpriseId(ctx context.Context) int64 {
	return ctx.Value("eid").(int64)
}

func GetWorkshopId(ctx context.Context) int64 {
	return ctx.Value("wid").(int64)
}

func GetStafferId(ctx context.Context) int64 {
	return ctx.Value("uid").(int64)
}

func GetRole(ctx context.Context) string {
	return ctx.Value("rol").(string)
}

func SelectWorkshopId(ctx context.Context, input int64) int64 {
	if GetRole(ctx) != "boss" {
		return GetWorkshopId(ctx)
	}
	return input
}

func ZeroSelectWorkshopId(ctx context.Context, input int64) int64 {
	if GetRole(ctx) != "boss" {
		return 0
	}
	return input
}

func SelectErrorCode(ctxs ...context.Context) int {
	if len(ctxs) == 0 {
		return -1
	}
	if code, ok := ctxs[0].Value("errCode").(int); ok && code != 0 {
		return code
	}
	return -1
}

func SetErrorCode(ctx context.Context, code int) context.Context {
	return context.WithValue(ctx, "errCode", code)
}
