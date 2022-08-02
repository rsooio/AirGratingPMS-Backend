package product_test

import (
	"air-grating-pms-backend/model/product"
	"fmt"
	"testing"
	"time"
)

var mock = &product.Product{
	Id:           1,
	ProductSetId: 22,
	TechnologyId: 333,
	Length:       4444,
	Width:        55555,
	UnitPrice:    666666,
	Quantity:     7777777,
	CreateTime:   time.Now(),
	UpdateTime:   time.Now().Add(time.Hour * 24),
	Remark:       "hello, world",
	Version:      88888888,
}

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

	// model := product.NewProductModel(conn, cache)

	// model.Insert(context.Background(), &product.Product{
	// 	Id:           0,
	// 	ProductSetId: 0,
	// 	TechnologyId: 0,
	// 	Length:       0,
	// 	Width:        0,
	// 	UnitPrice:    0,
	// 	Quantity:     0,
	// 	CreateTime:   time.Time{},
	// 	UpdateTime:   time.Time{},
	// 	Remark:       sql.NullString{},
	// 	Version:      0,
	// })

	// t := time.Now().UnixMilli()
	// model.Update(context.Background(), &product_set.ProductSet{
	// 	Id:      1,
	// 	Version: 7,
	// })
	// fmt.Println(time.Now().UnixMilli() - t)

	size := 8
	fmt.Println(1 << size)
	list := make(product.ProductList, 0, 1<<size)
	list = append(list, mock)
	for i := 0; i < size; i++ {
		list = append(list, list...)
	}
	// fmt.Println(list)

	t := time.Now().UnixMicro()
	_ = list.RpcList()
	fmt.Println(time.Now().UnixMicro() - t)
}
