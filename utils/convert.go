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
