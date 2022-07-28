package convert

import (
	"air-grating-pms-backend/api/staffer/internal/types"
	"air-grating-pms-backend/rpc/staffer/pb"

	"github.com/zeromicro/go-zero/core/mr"
)

func SingleConvert(info *pb.StafferInfo) *types.Staffer {
	return &types.Staffer{
		Id:          info.Id,
		Username:    info.Username,
		WorkshopId:  info.WorkshopId,
		Name:        info.Name,
		Role:        info.Role,
		Gender:      info.Gender,
		PhoneNumber: info.PhoneNumber,
		Email:       info.Email,
		Address:     info.Address,
		Remark:      info.Remark,
	}
}

func ListConvert(list *pb.StafferList) (*[]types.Staffer, error) {
	r, _ := mr.MapReduce(func(source chan<- interface{}) {
		for _, sf := range list.StafferList {
			source <- sf
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		sf := item.(*pb.StafferInfo)
		writer.Write(SingleConvert(sf))
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
		var sfs []types.Staffer
		for sf := range pipe {
			sfs = append(sfs, *sf.(*types.Staffer))
		}
		writer.Write(&sfs)
	})

	return r.(*[]types.Staffer), nil
}
