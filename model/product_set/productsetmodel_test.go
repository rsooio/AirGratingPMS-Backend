package product_set_test

import (
	"testing"
)

func TestMain(m *testing.M) {
	// conn := sqlx.NewMysql("air:@tcp(localhost)/air?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai")
	// cache := cache.CacheConf{
	// 	cache.NodeConf{
	// 		RedisConf: redis.RedisConf{
	// 			Host: "localhost:6379",
	// 			Pass: "",
	// 			Type: "node",
	// 		},
	// 		Weight: 1,
	// 	},
	// }

	// model := product_set.NewProductSetModel(conn, cache)

	// model.Insert(context.Background(), &product_set.ProductSet{
	// 	OrderId: 4,
	// 	Remark:  sql.NullString{String: "hello, world", Valid: true},
	// 	Version: 5,
	// })

	// t := time.Now().UnixMilli()
	// model.Update(context.Background(), &product_set.ProductSet{
	// 	Id:     1,
	// 	Remark: sql.NullString{String: "", Valid: true},
	// })
	// fmt.Println(time.Now().UnixMilli() - t)
}
