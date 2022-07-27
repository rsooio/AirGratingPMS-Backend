package utils

import (
	"air-grating-pms-backend/model/workshop"
	"air-grating-pms-backend/rpc/workshop/types"

	"github.com/zeromicro/go-zero/core/mr"
)

func Convert(list *[]*workshop.Workshop) (*types.WorkshopList, error) {
	out, _ := mr.MapReduce(func(source chan<- interface{}) {
		for _, ws := range *list {
			source <- ws
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		ws := item.(*workshop.Workshop)
		writer.Write(&types.WorkshopInfoWithId{
			Id:           ws.Id,
			EnterpriseId: ws.EnterpriseId,
			ManagerId:    ws.ManagerId.Int64,
			Name:         ws.Name,
			Address:      ws.Address.String,
			PhoneNumber:  ws.PhoneNumber.String,
			Remark:       ws.Remark.String,
			Version:      ws.Version,
		})
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
		var out []*types.WorkshopInfoWithId
		for ws := range pipe {
			out = append(out, ws.(*types.WorkshopInfoWithId))
		}
		writer.Write(out)
	})

	return &types.WorkshopList{List: out.([]*types.WorkshopInfoWithId)}, nil
}
