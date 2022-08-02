package pb

import (
	"air-grating-pms-backend/api/order/types"
)

func (m *ProductInfo) Api() *types.ProductInfo {
	return &types.ProductInfo{
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

func (m *ProductList) ApiList() *[]types.ProductInfo {
	var ret []types.ProductInfo
	for _, i := range m.List {
		ret = append(ret, *i.Api())
	}
	return &ret
}
