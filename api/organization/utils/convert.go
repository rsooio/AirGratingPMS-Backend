package utils

import (
	"air-grating-pms-backend/api/organization/internal/types"
	"air-grating-pms-backend/rpc/workshop/pb"

	"github.com/zeromicro/go-zero/core/mr"
)

func Convert(list *pb.WorkshopList) (*[]types.Workshop, error) {
	r, _ := mr.MapReduce(func(source chan<- interface{}) {
		for _, ws := range list.List {
			source <- ws
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		ws := item.(*pb.WorkshopInfo)
		writer.Write(&types.Workshop{
			Id:          ws.Id,
			Name:        ws.Name,
			Address:     ws.Address,
			PhoneNumber: ws.PhoneNumber,
			ManagerId:   ws.ManagerId,
			Remark:      ws.Remark,
		})
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
		var wss []types.Workshop
		for ws := range pipe {
			wss = append(wss, *ws.(*types.Workshop))
		}
		writer.Write(&wss)
	})

	return r.(*[]types.Workshop), nil
}
