package pb

import "air-grating-pms-backend/api/order/types"

func (m *ProductSetInfo) Api() *types.ProductSetInfo {
	return &types.ProductSetInfo{
		Id:     m.Id,
		Remark: m.Remark,
	}
}

func (m *ProductSetList) ApiList() *[]types.ProductSetInfo {
	var ret []types.ProductSetInfo
	for _, i := range m.List {
		ret = append(ret, *i.Api())
	}
	return &ret
}
