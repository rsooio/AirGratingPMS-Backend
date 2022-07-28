package utils

import (
	"air-grating-pms-backend/model/staffer"
	"air-grating-pms-backend/rpc/staffer/pb"

	"github.com/zeromicro/go-zero/core/mr"
)

func ListConvert(list []*staffer.Staffer) (*pb.StafferList, error) {
	r, _ := mr.MapReduce(func(source chan<- interface{}) {
		for _, staffer := range list {
			source <- staffer
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		info := item.(*staffer.Staffer)
		writer.Write(SingleConvert(info))
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
		var staffers []*pb.StafferInfo
		for stafferInfoWithId := range pipe {
			staffers = append(staffers, stafferInfoWithId.(*pb.StafferInfo))
		}
		writer.Write(staffers)
	})

	return &pb.StafferList{StafferList: r.([]*pb.StafferInfo)}, nil
}

func SingleConvert(info *staffer.Staffer) *pb.StafferInfo {
	return &pb.StafferInfo{
		Id:             info.Id,
		EnterpriseId:   info.EnterpriseId,
		WorkshopId:     info.WorkshopId,
		Username:       info.Username,
		Role:           info.Role,
		Name:           info.Name,
		HashedPassword: info.HashedPassword,
		Gender:         info.Gender.String,
		PhoneNumber:    info.PhoneNumber.String,
		Email:          info.Email.String,
		Address:        info.Address.String,
		ExpireTime:     info.ExpireTime,
		Remark:         info.Remark.String,
		Version:        info.Version,
	}
}
