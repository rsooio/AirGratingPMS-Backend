package pb

import "air-grating-pms-backend/api/organization/types"

func (m *WorkshopInfo) Api() *types.WorkshopInfo {
	return &types.WorkshopInfo{
		Id:          m.Id,
		Name:        m.Name,
		Address:     m.Address,
		PhoneNumber: m.PhoneNumber,
		ManagerId:   m.ManagerId,
		Remark:      m.Remark,
	}
}

func (m *WorkshopList) ApiList() *[]types.WorkshopInfo {
	var ret []types.WorkshopInfo
	for _, i := range m.List {
		ret = append(ret, *i.Api())
	}
	return &ret
}
