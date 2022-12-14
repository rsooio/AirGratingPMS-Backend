// Code generated by goctl. DO NOT EDIT.
package types

type InsertOrderReq struct {
	WorkshopId        int64  `json:"workshop_id,optional"`
	ClientId          int64  `json:"client_id"`
	Address           string `json:"address"`
	Linkman           string `json:"linkman"`
	PhoneNumber       string `json:"phone_number"`
	Email             string `json:"email"`
	CorrespondingCode string `json:"corresponding_code"`
	Remark            string `json:"remark"`
}

type InsertOrderReply struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}

type UpdateOrderInfoReq struct {
	Id                int64  `json:"id"`
	WorkshopId        int64  `json:"workshop_id,optional"`
	ClientId          int64  `json:"client_id,optional"`
	Address           string `json:"address,optional"`
	Linkman           string `json:"linkman,optional"`
	PhoneNumber       string `json:"phone_number,optional"`
	Email             string `json:"email,optional"`
	CorrespondingCode string `json:"corresponding_code"`
	Remark            string `json:"remark,optional"`
}

type UpdateOrderInfoReply struct {
	Message       string   `json:"message"`
	ChangedFields []string `json:"changed_fields"`
}

type DeleteOrderReq struct {
	Id int64 `json:"id"`
}

type DeleteOrderReply struct {
	Message string `json:"message"`
}

type OrderInfo struct {
	Id                int64  `json:"id"`
	WorkshopId        int64  `json:"workshop_id,optional"`
	ClientId          int64  `json:"client_id"`
	Address           string `json:"address"`
	Linkman           string `json:"linkman"`
	PhoneNumber       string `json:"phone_number"`
	Email             string `json:"email"`
	CorrespondingCode string `json:"corresponding_code"`
	Remark            string `json:"remark"`
}

type GetOrderListByEnterpriseReq struct {
	PageNumber int32 `json:"pn"`
	PageSize   int32 `json:"ps"`
}

type GetOrderListByEnterpriseReply struct {
	Message   string      `json:"message"`
	Count     int64       `json:"count"`
	OrderList []OrderInfo `json:"order_list"`
}

type GetOrderListByWorkshopReq struct {
	WorkshopId int64 `json:"workshop_id,optional"`
	PageNumber int32 `json:"pn"`
	PageSize   int32 `json:"ps"`
}

type GetOrderListByWorkshopReply struct {
	Message   string      `json:"message"`
	Count     int64       `json:"count"`
	OrderList []OrderInfo `json:"order_list"`
}

type GetOrderListByProductionPlanReq struct {
	ProductionPlanId int64 `json:"pid"`
	PageNumber       int32 `json:"pn"`
	PageSize         int32 `json:"ps"`
}

type GetOrderListByProductionPlanReply struct {
	Message     string      `json:"message"`
	Count       int64       `json:"count"`
	ProductList []OrderInfo `json:"product_list"`
}

type InsertProductSetReq struct {
	OrderId int64  `json:"order_id"`
	Remark  string `json:"remark"`
}

type InsertProductSetReply struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}

type DeleteProductSetReq struct {
	Id int64 `json:"id"`
}

type DeleteProductSetReply struct {
	Message string `json:"message"`
}

type UpdateProductSetReq struct {
	Id     int64  `json:"id"`
	Remark string `json:"remark,optional"`
}

type UpdateProductSetReply struct {
	Message       string   `json:"message"`
	ChangedFields []string `json:"changed_fields"`
}

type GetProductSetListByOrderReq struct {
	OrderId    int64 `json:"order_id"`
	PageNumber int32 `json:"pn"`
	PageSize   int32 `json:"ps"`
}

type ProductSetInfo struct {
	Id     int64  `json:"id"`
	Remark string `json:"remark"`
}

type GetProductSetListByOrderReply struct {
	Message        string           `json:"message"`
	Count          int64            `json:"count"`
	ProductSetList []ProductSetInfo `json:"product_set_list"`
}

type InsertProductReq struct {
	ProductSetId int64  `json:"product_set_id"`
	TechnologyId int64  `json:"technology_id"`
	Length       string `json:"length"`
	Width        string `json:"width"`
	UnitPrice    string `json:"unit_price"`
	Quantity     int64  `json:"quantity"`
	Remark       string `json:"remark"`
}

type InsertProductReply struct {
	Id      int64  `json:"id"`
	Message string `json:"message"`
}

type DeleteProductReq struct {
	Id int64 `json:"id"`
}

type DeleteProductReply struct {
	Message string `json:"message"`
}

type UpdateProductReq struct {
	Id           int64  `json:"id"`
	ProductSetId int64  `json:"product_set_id,optional"`
	TechnologyId int64  `json:"technology_id,optional"`
	Length       string `json:"length,optional"`
	Width        string `json:"width,optional"`
	UnitPrice    string `json:"unit_price,optional"`
	Quantity     int64  `json:"quantity,optional"`
	Remark       string `json:"remark,optional"`
}

type UpdateProductReply struct {
	Message       string   `json:"message"`
	ChangedFields []string `json:"changed_fields"`
}

type ProductInfo struct {
	Id           int64  `json:"id"`
	ProductSetId int64  `json:"product_set_id"`
	TechnologyId int64  `json:"technology_id"`
	Length       string `json:"length"`
	Width        string `json:"width"`
	UnitPrice    string `json:"unit_price"`
	Quantity     int64  `json:"quantity"`
	Remark       string `json:"remark"`
}

type GetProductListByProductSetReq struct {
	ProductSetId int64 `json:"product_set_id"`
	PageNumber   int32 `json:"pn"`
	PageSize     int32 `json:"ps"`
}

type GetProductListByProductSetReply struct {
	Message     string        `json:"message"`
	Count       int64         `json:"count"`
	ProductList []ProductInfo `json:"product_list"`
}
