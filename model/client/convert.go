package client

import "air-grating-pms-backend/rpc/client/pb"

type ClientList []*Client

func (m *Client) Rpc() *pb.ClientInfo {
	return &pb.ClientInfo{
		Id:           m.Id,
		EnterpriseId: m.EnterpriseId,
		WorkshopId:   m.WorkshopId,
		Name:         m.Name,
		PhoneNumber:  m.PhoneNumber,
		Email:        m.Email,
		Address:      m.Address,
		Remark:       m.Remark,
		Version:      m.Version,
	}
}

func (m *ClientList) RpcList() *[]*pb.ClientInfo {
	var ret []*pb.ClientInfo
	for _, i := range *m {
		ret = append(ret, i.Rpc())
	}
	return &ret
}
