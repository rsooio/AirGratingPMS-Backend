package pb

import "air-grating-pms-backend/api/staffer/types"

func (m *StafferInfo) Api() *types.StafferInfo {
	return &types.StafferInfo{
		Id:          m.Id,
		Username:    m.Username,
		WorkshopId:  m.WorkshopId,
		Name:        m.Name,
		Role:        m.Role,
		Gender:      m.Gender,
		PhoneNumber: m.PhoneNumber,
		Email:       m.Email,
		Address:     m.Address,
		Remark:      m.Remark,
	}
}

func (m *StafferList) ApiList() *[]types.StafferInfo {
	var ret []types.StafferInfo
	for _, i := range m.List {
		ret = append(ret, *i.Api())
	}
	return &ret
}
