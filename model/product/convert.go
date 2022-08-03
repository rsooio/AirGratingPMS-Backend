package product

import (
	"air-grating-pms-backend/rpc/product/pb"
)

type ProductList []*Product

func (m *Product) Rpc() *pb.ProductInfo {
	return &pb.ProductInfo{
		Id:           m.Id,
		ProductSetId: m.ProductSetId,
		TechnologyId: m.TechnologyId,
		Length:       m.Length,
		Width:        m.Width,
		UnitPrice:    m.UnitPrice,
		Quantity:     m.Quantity,
		Remark:       m.Remark,
	}
}

func (m *ProductList) RpcList() *[]*pb.ProductInfo {
	var ret []*pb.ProductInfo
	for _, i := range *m {
		ret = append(ret, i.Rpc())
	}
	return &ret
}
