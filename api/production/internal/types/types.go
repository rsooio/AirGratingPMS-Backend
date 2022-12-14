// Code generated by goctl. DO NOT EDIT.
package types

type InsertProductionPlanReq struct {
	WorkshopId int64  `json:"workshop_id,optional"`
	Remark     string `json:"remark,optional"`
}

type InsertProductionPlanReply struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}

type DeleteProductionPlanReq struct {
	Id int64 `json:"id"`
}

type DeleteProductionPlanReply struct {
	Message string `json:"message"`
}

type UpdateProductionPlanReq struct {
	Id         int64  `json:"id"`
	WorkshopId int64  `json:"workshop_id,optional"`
	Remark     string `json:"remark,optional"`
}

type UpdateProductionPlanReply struct {
	Message string `json:"message"`
}

type GetProductionPlanInfoByIdReq struct {
	Id int64 `json:"id"`
}

type ProductionPlanInfo struct {
	Id         int64  `json:"id"`
	WorkshopId int64  `json:"workshop_id"`
	State      int64  `json:"state"`
	Remark     string `json:"remark"`
	Version    int64  `json:"version"`
}

type GetProductionPlanListByEnterpriseReq struct {
	PageNumber int32 `json:"pn"`
	PageSize   int32 `json:"ps"`
}

type GetProductionPlanListByEnterpriseReply struct {
	Message string               `json:"message"`
	Count   int64                `json:"count"`
	List    []ProductionPlanInfo `json:"list"`
}

type GetProductionPlanListByWorkshopReq struct {
	WorkshopId int64 `json:"workshop_id"`
	PageNumber int32 `json:"pn"`
	PageSize   int32 `json:"ps"`
}

type GetProductionPlanListByWorkshopReply struct {
	Message string               `json:"message"`
	Count   int64                `json:"count"`
	List    []ProductionPlanInfo `json:"list"`
}

type StartProductionPlanReq struct {
	Id int64 `json:"id"`
}

type StartProductionPlanReply struct {
	Message string `json:"message"`
}

type CompleteProductionPlanReq struct {
	Id int64 `json:"id"`
}

type CompleteProductionPlanReply struct {
	Message string `json:"message"`
}

type CompleteOrderProductionReq struct {
	Id int64 `json:"id"`
}

type CompleteOrderProductionReply struct {
	Message string `json:"message"`
}

type SetOrderProductionPlanReq struct {
	Id               int64 `json:"id"`
	ProductionPlanId int64 `json:"pid"`
}

type SetOrderProductionPlanReply struct {
	Message string `json:"message"`
}
