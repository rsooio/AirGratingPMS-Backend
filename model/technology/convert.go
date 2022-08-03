package technology

import "air-grating-pms-backend/rpc/technology/pb"

type TechnologyList []*Technology

func (m *Technology) Rpc() *pb.TechnologyInfo {
	return &pb.TechnologyInfo{
		Id:           m.Id,
		EnterpriseId: m.EnterpriseId,
		WorkshopId:   m.WorkshopId,
		Name:         m.Name,
		Info:         m.Info,
		Remark:       m.Remark,
		Version:      m.Version,
	}
}

func (m *TechnologyList) RpcList() *[]*pb.TechnologyInfo {
	var ret []*pb.TechnologyInfo
	for _, i := range *m {
		ret = append(ret, i.Rpc())
	}
	return &ret
}
