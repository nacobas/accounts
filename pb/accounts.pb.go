// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.1
// source: pb/accounts.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductCode   uint32 `protobuf:"varint,1,opt,name=product_code,json=productCode,proto3" json:"product_code,omitempty"`
	AccountId     string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AccountNumber string `protobuf:"bytes,3,opt,name=account_number,json=accountNumber,proto3" json:"account_number,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_accounts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_pb_accounts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_pb_accounts_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetProductCode() uint32 {
	if x != nil {
		return x.ProductCode
	}
	return 0
}

func (x *Account) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *Account) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

type OpenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductCode uint32 `protobuf:"varint,1,opt,name=product_code,json=productCode,proto3" json:"product_code,omitempty"`
}

func (x *OpenRequest) Reset() {
	*x = OpenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_accounts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenRequest) ProtoMessage() {}

func (x *OpenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_accounts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenRequest.ProtoReflect.Descriptor instead.
func (*OpenRequest) Descriptor() ([]byte, []int) {
	return file_pb_accounts_proto_rawDescGZIP(), []int{1}
}

func (x *OpenRequest) GetProductCode() uint32 {
	if x != nil {
		return x.ProductCode
	}
	return 0
}

type OpenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *OpenResponse) Reset() {
	*x = OpenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_accounts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenResponse) ProtoMessage() {}

func (x *OpenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_accounts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenResponse.ProtoReflect.Descriptor instead.
func (*OpenResponse) Descriptor() ([]byte, []int) {
	return file_pb_accounts_proto_rawDescGZIP(), []int{2}
}

func (x *OpenResponse) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

type CloseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *CloseRequest) Reset() {
	*x = CloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_accounts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseRequest) ProtoMessage() {}

func (x *CloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_accounts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseRequest.ProtoReflect.Descriptor instead.
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return file_pb_accounts_proto_rawDescGZIP(), []int{3}
}

func (x *CloseRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type CloseReponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *CloseReponse) Reset() {
	*x = CloseReponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_accounts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseReponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseReponse) ProtoMessage() {}

func (x *CloseReponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_accounts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseReponse.ProtoReflect.Descriptor instead.
func (*CloseReponse) Descriptor() ([]byte, []int) {
	return file_pb_accounts_proto_rawDescGZIP(), []int{4}
}

func (x *CloseReponse) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type SetRestrictionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level     int32  `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *SetRestrictionRequest) Reset() {
	*x = SetRestrictionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_accounts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRestrictionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRestrictionRequest) ProtoMessage() {}

func (x *SetRestrictionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_accounts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRestrictionRequest.ProtoReflect.Descriptor instead.
func (*SetRestrictionRequest) Descriptor() ([]byte, []int) {
	return file_pb_accounts_proto_rawDescGZIP(), []int{5}
}

func (x *SetRestrictionRequest) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *SetRestrictionRequest) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type SetRestrictionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SetRestrictionResponse) Reset() {
	*x = SetRestrictionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_accounts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRestrictionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRestrictionResponse) ProtoMessage() {}

func (x *SetRestrictionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_accounts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRestrictionResponse.ProtoReflect.Descriptor instead.
func (*SetRestrictionResponse) Descriptor() ([]byte, []int) {
	return file_pb_accounts_proto_rawDescGZIP(), []int{6}
}

func (x *SetRestrictionResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_pb_accounts_proto protoreflect.FileDescriptor

var file_pb_accounts_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x30, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x32, 0x0a, 0x0c, 0x4f, 0x70, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2d, 0x0a,
	0x0c, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x0c,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x15, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x16, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x9f, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e,
	0x12, 0x0c, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x0d, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x63, 0x6f, 0x62, 0x61, 0x73, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pb_accounts_proto_rawDescOnce sync.Once
	file_pb_accounts_proto_rawDescData = file_pb_accounts_proto_rawDesc
)

func file_pb_accounts_proto_rawDescGZIP() []byte {
	file_pb_accounts_proto_rawDescOnce.Do(func() {
		file_pb_accounts_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_accounts_proto_rawDescData)
	})
	return file_pb_accounts_proto_rawDescData
}

var file_pb_accounts_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_accounts_proto_goTypes = []interface{}{
	(*Account)(nil),                // 0: Account
	(*OpenRequest)(nil),            // 1: OpenRequest
	(*OpenResponse)(nil),           // 2: OpenResponse
	(*CloseRequest)(nil),           // 3: CloseRequest
	(*CloseReponse)(nil),           // 4: CloseReponse
	(*SetRestrictionRequest)(nil),  // 5: SetRestrictionRequest
	(*SetRestrictionResponse)(nil), // 6: SetRestrictionResponse
}
var file_pb_accounts_proto_depIdxs = []int32{
	0, // 0: OpenResponse.account:type_name -> Account
	1, // 1: AccountService.Open:input_type -> OpenRequest
	3, // 2: AccountService.Close:input_type -> CloseRequest
	5, // 3: AccountService.SetRestriction:input_type -> SetRestrictionRequest
	2, // 4: AccountService.Open:output_type -> OpenResponse
	4, // 5: AccountService.Close:output_type -> CloseReponse
	6, // 6: AccountService.SetRestriction:output_type -> SetRestrictionResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_accounts_proto_init() }
func file_pb_accounts_proto_init() {
	if File_pb_accounts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_accounts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_pb_accounts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenRequest); i {
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
		file_pb_accounts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenResponse); i {
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
		file_pb_accounts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseRequest); i {
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
		file_pb_accounts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseReponse); i {
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
		file_pb_accounts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRestrictionRequest); i {
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
		file_pb_accounts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRestrictionResponse); i {
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
			RawDescriptor: file_pb_accounts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_accounts_proto_goTypes,
		DependencyIndexes: file_pb_accounts_proto_depIdxs,
		MessageInfos:      file_pb_accounts_proto_msgTypes,
	}.Build()
	File_pb_accounts_proto = out.File
	file_pb_accounts_proto_rawDesc = nil
	file_pb_accounts_proto_goTypes = nil
	file_pb_accounts_proto_depIdxs = nil
}
