package production_plan

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

const (
	_ int64 = iota
	WaitProduction
	InProduction
	ProductionDone
)
