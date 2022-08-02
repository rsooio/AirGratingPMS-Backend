package order

import (
	"air-grating-pms-backend/rpc/order/pb"
)

type OrderList []*Order

func (m Order) Rpc() *pb.OrderInfo {
	return &pb.OrderInfo{
		Id:                m.Id,
		EnterpriseId:      m.EnterpriseId,
		WorkshopId:        m.WorkshopId,
		ClientId:          m.ClientId,
		ProductionPlanId:  m.ProductionPlanId,
		State:             m.State,
		Address:           m.Address,
		Linkman:           m.Linkman,
		PhoneNumber:       m.PhoneNumber,
		Email:             m.Email,
		CorrespondingCode: m.CorrespondingCode,
		Remark:            m.Remark,
	}
}

func (m OrderList) RpcList() *[]*pb.OrderInfo {
	var ret []*pb.OrderInfo
	for _, i := range m {
		ret = append(ret, i.Rpc())
	}
	return &ret
}
