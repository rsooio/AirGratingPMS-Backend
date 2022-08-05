package production_plan

import "air-grating-pms-backend/rpc/production_plan/pb"

type ProductionPlanList []*ProductionPlan

func (m *ProductionPlan) Rpc() *pb.ProductionPlanInfo {
	return &pb.ProductionPlanInfo{
		Id:           m.Id,
		EnterpriseId: m.EnterpriseId,
		WorkshopId:   m.WorkshopId,
		State:        m.State,
		Remark:       m.Remark,
		Version:      m.Version,
	}
}

func (m *ProductionPlanList) RpcList() *[]*pb.ProductionPlanInfo {
	var ret []*pb.ProductionPlanInfo
	for _, i := range *m {
		ret = append(ret, i.Rpc())
	}
	return &ret
}
