// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: product_set.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_product_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_product_set_proto_rawDescGZIP(), []int{0}
}

type ProductSetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	OrderId int64  `protobuf:"varint,2,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	Remark  string `protobuf:"bytes,3,opt,name=Remark,proto3" json:"Remark,omitempty"`
}

func (x *ProductSetInfo) Reset() {
	*x = ProductSetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_set_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSetInfo) ProtoMessage() {}

func (x *ProductSetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_product_set_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSetInfo.ProtoReflect.Descriptor instead.
func (*ProductSetInfo) Descriptor() ([]byte, []int) {
	return file_product_set_proto_rawDescGZIP(), []int{1}
}

func (x *ProductSetInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductSetInfo) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *ProductSetInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type InsertResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *InsertResp) Reset() {
	*x = InsertResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_set_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertResp) ProtoMessage() {}

func (x *InsertResp) ProtoReflect() protoreflect.Message {
	mi := &file_product_set_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertResp.ProtoReflect.Descriptor instead.
func (*InsertResp) Descriptor() ([]byte, []int) {
	return file_product_set_proto_rawDescGZIP(), []int{2}
}

func (x *InsertResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_set_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_set_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_product_set_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindOneByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *FindOneByIdReq) Reset() {
	*x = FindOneByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_set_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneByIdReq) ProtoMessage() {}

func (x *FindOneByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_set_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneByIdReq.ProtoReflect.Descriptor instead.
func (*FindOneByIdReq) Descriptor() ([]byte, []int) {
	return file_product_set_proto_rawDescGZIP(), []int{4}
}

func (x *FindOneByIdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindListByOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId    int64 `protobuf:"varint,1,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	PageSize   int32 `protobuf:"varint,2,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	PageNumber int32 `protobuf:"varint,3,opt,name=PageNumber,proto3" json:"PageNumber,omitempty"`
}

func (x *FindListByOrderReq) Reset() {
	*x = FindListByOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_set_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindListByOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindListByOrderReq) ProtoMessage() {}

func (x *FindListByOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_product_set_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindListByOrderReq.ProtoReflect.Descriptor instead.
func (*FindListByOrderReq) Descriptor() ([]byte, []int) {
	return file_product_set_proto_rawDescGZIP(), []int{5}
}

func (x *FindListByOrderReq) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *FindListByOrderReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FindListByOrderReq) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

type ProductSetList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64             `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	List  []*ProductSetInfo `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
}

func (x *ProductSetList) Reset() {
	*x = ProductSetList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_set_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductSetList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSetList) ProtoMessage() {}

func (x *ProductSetList) ProtoReflect() protoreflect.Message {
	mi := &file_product_set_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSetList.ProtoReflect.Descriptor instead.
func (*ProductSetList) Descriptor() ([]byte, []int) {
	return file_product_set_proto_rawDescGZIP(), []int{6}
}

func (x *ProductSetList) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ProductSetList) GetList() []*ProductSetInfo {
	if x != nil {
		return x.List
	}
	return nil
}

var File_product_set_proto protoreflect.FileDescriptor

var file_product_set_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x65, 0x74, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x1c, 0x0a, 0x0a,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x1b, 0x0a, 0x09, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x4f,
	0x6e, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x6a, 0x0a, 0x12, 0x46, 0x69, 0x6e,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12,
	0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x50, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x56, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xcd, 0x02,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x65, 0x74, 0x12, 0x3c, 0x0a, 0x06,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x65, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x65, 0x74, 0x2e,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x32, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x65,
	0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x37,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x65,
	0x74, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x45, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x4f,
	0x6e, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x65, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x65, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4d,
	0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x65, 0x74, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_set_proto_rawDescOnce sync.Once
	file_product_set_proto_rawDescData = file_product_set_proto_rawDesc
)

func file_product_set_proto_rawDescGZIP() []byte {
	file_product_set_proto_rawDescOnce.Do(func() {
		file_product_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_set_proto_rawDescData)
	})
	return file_product_set_proto_rawDescData
}

var file_product_set_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_product_set_proto_goTypes = []interface{}{
	(*Empty)(nil),              // 0: productset.Empty
	(*ProductSetInfo)(nil),     // 1: productset.ProductSetInfo
	(*InsertResp)(nil),         // 2: productset.InsertResp
	(*DeleteReq)(nil),          // 3: productset.DeleteReq
	(*FindOneByIdReq)(nil),     // 4: productset.FindOneByIdReq
	(*FindListByOrderReq)(nil), // 5: productset.FindListByOrderReq
	(*ProductSetList)(nil),     // 6: productset.ProductSetList
}
var file_product_set_proto_depIdxs = []int32{
	1, // 0: productset.ProductSetList.List:type_name -> productset.ProductSetInfo
	1, // 1: productset.productset.Insert:input_type -> productset.ProductSetInfo
	3, // 2: productset.productset.Delete:input_type -> productset.DeleteReq
	1, // 3: productset.productset.Update:input_type -> productset.ProductSetInfo
	4, // 4: productset.productset.FindOneById:input_type -> productset.FindOneByIdReq
	5, // 5: productset.productset.FindListByOrder:input_type -> productset.FindListByOrderReq
	2, // 6: productset.productset.Insert:output_type -> productset.InsertResp
	0, // 7: productset.productset.Delete:output_type -> productset.Empty
	0, // 8: productset.productset.Update:output_type -> productset.Empty
	1, // 9: productset.productset.FindOneById:output_type -> productset.ProductSetInfo
	6, // 10: productset.productset.FindListByOrder:output_type -> productset.ProductSetList
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_product_set_proto_init() }
func file_product_set_proto_init() {
	if File_product_set_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_set_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_set_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSetInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_set_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_set_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_set_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneByIdReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_set_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindListByOrderReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_product_set_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductSetList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_product_set_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_set_proto_goTypes,
		DependencyIndexes: file_product_set_proto_depIdxs,
		MessageInfos:      file_product_set_proto_msgTypes,
	}.Build()
	File_product_set_proto = out.File
	file_product_set_proto_rawDesc = nil
	file_product_set_proto_goTypes = nil
	file_product_set_proto_depIdxs = nil
}
