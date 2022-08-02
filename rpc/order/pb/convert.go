package pb

import (
	"air-grating-pms-backend/api/order/types"
)

func (m *OrderInfo) Api() *types.OrderInfo {
	return &types.OrderInfo{
		Id:          m.Id,
		WorkshopId:  m.WorkshopId,
		ClientId:    m.ClientId,
		Address:     m.Address,
		Linkman:     m.Linkman,
		PhoneNumber: m.PhoneNumber,
		Email:       m.Email,
		Remark:      m.Remark,
	}
}

func (m *OrderList) ApiList() *[]types.OrderInfo {
	var ret []types.OrderInfo
	for _, i := range m.List {
		ret = append(ret, *i.Api())
	}
	return &ret
}
