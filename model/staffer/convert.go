package staffer

import "air-grating-pms-backend/rpc/staffer/pb"

type StafferList []*Staffer

func (m *Staffer) Rpc() *pb.StafferInfo {
	return &pb.StafferInfo{
		Id:             m.Id,
		EnterpriseId:   m.EnterpriseId,
		WorkshopId:     m.WorkshopId,
		Username:       m.Username,
		Role:           m.Role,
		Name:           m.Name,
		HashedPassword: m.HashedPassword,
		Gender:         m.Gender,
		PhoneNumber:    m.PhoneNumber,
		Email:          m.Email,
		Address:        m.Address,
		ExpireTime:     m.ExpireTime,
		Remark:         m.Remark,
		Version:        m.Version,
	}
}

func (m *StafferList) RpcList() *[]*pb.StafferInfo {
	var ret []*pb.StafferInfo
	for _, i := range *m {
		ret = append(ret, i.Rpc())
	}
	return &ret
}
