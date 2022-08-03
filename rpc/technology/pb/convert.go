package pb

import "air-grating-pms-backend/api/technology/types"

func (m *TechnologyInfo) Api() *types.TechnologyInfo {
	return &types.TechnologyInfo{
		Id:         m.Id,
		WorkshopId: m.WorkshopId,
		Name:       m.Name,
		Info:       m.Info,
		Remark:     m.Remark,
	}
}

func (m *TechnologyList) ApiList() *[]types.TechnologyInfo {
	var ret []types.TechnologyInfo
	for _, i := range m.List {
		ret = append(ret, *i.Api())
	}
	return &ret
}
