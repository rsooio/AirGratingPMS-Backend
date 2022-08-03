package workshop

import "air-grating-pms-backend/rpc/workshop/pb"

type WorkshopList []*Workshop

func (m *Workshop) Rpc() *pb.WorkshopInfo {
	return &pb.WorkshopInfo{
		Id:           m.Id,
		EnterpriseId: m.EnterpriseId,
		ManagerId:    m.ManagerId,
		Name:         m.Name,
		Address:      m.Address,
		PhoneNumber:  m.PhoneNumber,
		Remark:       m.Remark,
		Version:      m.Version,
	}
}

func (m *WorkshopList) RpcList() *[]*pb.WorkshopInfo {
	var ret []*pb.WorkshopInfo
	for _, i := range *m {
		ret = append(ret, i.Rpc())
	}
	return &ret
}
