package pb

import "air-grating-pms-backend/api/production/types"

func (m *ProductionPlanInfo) Api() *types.ProductionPlanInfo {
	return &types.ProductionPlanInfo{
		Id:         m.Id,
		WorkshopId: m.WorkshopId,
		State:      m.State,
		Remark:     m.Remark,
		Version:    m.Version,
	}
}

func (m *ProductionPlanList) ApiList() *[]types.ProductionPlanInfo {
	var ret []types.ProductionPlanInfo
	for _, i := range m.List {
		ret = append(ret, *i.Api())
	}
	return &ret
}
