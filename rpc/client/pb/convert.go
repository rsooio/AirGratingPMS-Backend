package pb

import "air-grating-pms-backend/api/client/types"

func (m *ClientInfo) Api() *types.ClientInfo {
	return &types.ClientInfo{
		Id:          m.Id,
		WorkshopId:  m.WorkshopId,
		Name:        m.Name,
		PhoneNumber: m.PhoneNumber,
		Email:       m.Email,
		Address:     m.Address,
		Remark:      m.Remark,
		Version:     m.Version,
	}
}

func (m *ClientList) ApiList() *[]types.ClientInfo {
	var ret []types.ClientInfo
	for _, i := range m.List {
		ret = append(ret, *i.Api())
	}
	return &ret
}
