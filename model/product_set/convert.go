package product_set

import (
	"air-grating-pms-backend/rpc/product_set/pb"
)

type ProductSetList []*ProductSet

func (m *ProductSet) Rpc() *pb.ProductSetInfo {
	return &pb.ProductSetInfo{
		Id:      m.Id,
		OrderId: m.OrderId,
		Remark:  m.Remark,
	}
}

func (m *ProductSetList) RpcList() *[]*pb.ProductSetInfo {
	var ret []*pb.ProductSetInfo
	for _, i := range *m {
		ret = append(ret, i.Rpc())
	}
	return &ret
}
