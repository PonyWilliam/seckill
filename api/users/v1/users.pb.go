// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.9.0
// source: api/users/v1/users.proto

package v1

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

type CommonReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//用于对操作不需要返回消息的消息统一回复
	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CommonReply) Reset() {
	*x = CommonReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_users_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonReply) ProtoMessage() {}

func (x *CommonReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_users_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonReply.ProtoReflect.Descriptor instead.
func (*CommonReply) Descriptor() ([]byte, []int) {
	return file_api_users_v1_users_proto_rawDescGZIP(), []int{0}
}

func (x *CommonReply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//传递注册用户，区别就是不需要id,其次部分东西是系统生成的，如逾期记录等
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	IdCard   string `protobuf:"bytes,3,opt,name=id_card,json=idCard,proto3" json:"id_card,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Sex      int64  `protobuf:"varint,5,opt,name=sex,proto3" json:"sex,omitempty"` //性别，0男1女
	Age      int64  `protobuf:"varint,6,opt,name=age,proto3" json:"age,omitempty"`
	TrueName string `protobuf:"bytes,7,opt,name=true_name,json=trueName,proto3" json:"true_name,omitempty"`
	Carrer   int64  `protobuf:"varint,8,opt,name=carrer,proto3" json:"carrer,omitempty"` //职业类别,这个等会规定,通过枚举列举
}

func (x *UsersRequest) Reset() {
	*x = UsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_users_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersRequest) ProtoMessage() {}

func (x *UsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_users_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersRequest.ProtoReflect.Descriptor instead.
func (*UsersRequest) Descriptor() ([]byte, []int) {
	return file_api_users_v1_users_proto_rawDescGZIP(), []int{1}
}

func (x *UsersRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UsersRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UsersRequest) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *UsersRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UsersRequest) GetSex() int64 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *UsersRequest) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UsersRequest) GetTrueName() string {
	if x != nil {
		return x.TrueName
	}
	return ""
}

func (x *UsersRequest) GetCarrer() int64 {
	if x != nil {
		return x.Carrer
	}
	return 0
}

type UsersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//返回通用用户消息,主要是给用户本人和管理员看，设置权限访问
	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	IdCard   string `protobuf:"bytes,3,opt,name=id_card,json=idCard,proto3" json:"id_card,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Sex      int64  `protobuf:"varint,5,opt,name=sex,proto3" json:"sex,omitempty"` //性别，0男1女
	Age      int64  `protobuf:"varint,6,opt,name=age,proto3" json:"age,omitempty"`
	TrueName string `protobuf:"bytes,7,opt,name=true_name,json=trueName,proto3" json:"true_name,omitempty"`
	Carrer   int64  `protobuf:"varint,8,opt,name=carrer,proto3" json:"carrer,omitempty"` //职业类别,这个等会规定,通过枚举列举
	//下面是系统生成的
	Score int64 `protobuf:"varint,9,opt,name=score,proto3" json:"score,omitempty"`  //信誉分，默认100
	Count int64 `protobuf:"varint,10,opt,name=count,proto3" json:"count,omitempty"` //逾期次数，默认0
}

func (x *UsersReply) Reset() {
	*x = UsersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_users_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersReply) ProtoMessage() {}

func (x *UsersReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_users_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersReply.ProtoReflect.Descriptor instead.
func (*UsersReply) Descriptor() ([]byte, []int) {
	return file_api_users_v1_users_proto_rawDescGZIP(), []int{2}
}

func (x *UsersReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UsersReply) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UsersReply) GetIdCard() string {
	if x != nil {
		return x.IdCard
	}
	return ""
}

func (x *UsersReply) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UsersReply) GetSex() int64 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *UsersReply) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UsersReply) GetTrueName() string {
	if x != nil {
		return x.TrueName
	}
	return ""
}

func (x *UsersReply) GetCarrer() int64 {
	if x != nil {
		return x.Carrer
	}
	return 0
}

func (x *UsersReply) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *UsersReply) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CreateUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsersRequest *UsersRequest `protobuf:"bytes,1,opt,name=UsersRequest,proto3" json:"UsersRequest,omitempty"`
}

func (x *CreateUsersRequest) Reset() {
	*x = CreateUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_users_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUsersRequest) ProtoMessage() {}

func (x *CreateUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_users_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUsersRequest.ProtoReflect.Descriptor instead.
func (*CreateUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_users_v1_users_proto_rawDescGZIP(), []int{3}
}

func (x *CreateUsersRequest) GetUsersRequest() *UsersRequest {
	if x != nil {
		return x.UsersRequest
	}
	return nil
}

type CreateUsersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply *CommonReply `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *CreateUsersReply) Reset() {
	*x = CreateUsersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_users_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUsersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUsersReply) ProtoMessage() {}

func (x *CreateUsersReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_users_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUsersReply.ProtoReflect.Descriptor instead.
func (*CreateUsersReply) Descriptor() ([]byte, []int) {
	return file_api_users_v1_users_proto_rawDescGZIP(), []int{4}
}

func (x *CreateUsersReply) GetReply() *CommonReply {
	if x != nil {
		return x.Reply
	}
	return nil
}

type UpdateUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	Phone    string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Carrer   int64  `protobuf:"varint,3,opt,name=carrer,proto3" json:"carrer,omitempty"` //只有这3个参数可以改
}

func (x *UpdateUsersRequest) Reset() {
	*x = UpdateUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_users_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUsersRequest) ProtoMessage() {}

func (x *UpdateUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_users_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUsersRequest.ProtoReflect.Descriptor instead.
func (*UpdateUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_users_v1_users_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUsersRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UpdateUsersRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateUsersRequest) GetCarrer() int64 {
	if x != nil {
		return x.Carrer
	}
	return 0
}

type UpdateUsersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply *CommonReply `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *UpdateUsersReply) Reset() {
	*x = UpdateUsersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_users_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUsersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUsersReply) ProtoMessage() {}

func (x *UpdateUsersReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_users_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUsersReply.ProtoReflect.Descriptor instead.
func (*UpdateUsersReply) Descriptor() ([]byte, []int) {
	return file_api_users_v1_users_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUsersReply) GetReply() *CommonReply {
	if x != nil {
		return x.Reply
	}
	return nil
}

type DeleteUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//由管理员向这个接口发送请求
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUsersRequest) Reset() {
	*x = DeleteUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_users_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUsersRequest) ProtoMessage() {}

func (x *DeleteUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_users_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUsersRequest.ProtoReflect.Descriptor instead.
func (*DeleteUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_users_v1_users_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteUsersRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteUsersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reply *CommonReply `protobuf:"bytes,1,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (x *DeleteUsersReply) Reset() {
	*x = DeleteUsersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_users_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUsersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUsersReply) ProtoMessage() {}

func (x *DeleteUsersReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_users_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUsersReply.ProtoReflect.Descriptor instead.
func (*DeleteUsersReply) Descriptor() ([]byte, []int) {
	return file_api_users_v1_users_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteUsersReply) GetReply() *CommonReply {
	if x != nil {
		return x.Reply
	}
	return nil
}

type GetUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//这个也要鉴权，注意权限控制
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUsersRequest) Reset() {
	*x = GetUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_users_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRequest) ProtoMessage() {}

func (x *GetUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_users_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRequest.ProtoReflect.Descriptor instead.
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_users_v1_users_proto_rawDescGZIP(), []int{9}
}

func (x *GetUsersRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUsersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserReply *UsersReply `protobuf:"bytes,1,opt,name=userReply,proto3" json:"userReply,omitempty"`
}

func (x *GetUsersReply) Reset() {
	*x = GetUsersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_users_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersReply) ProtoMessage() {}

func (x *GetUsersReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_users_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersReply.ProtoReflect.Descriptor instead.
func (*GetUsersReply) Descriptor() ([]byte, []int) {
	return file_api_users_v1_users_proto_rawDescGZIP(), []int{10}
}

func (x *GetUsersReply) GetUserReply() *UsersReply {
	if x != nil {
		return x.UserReply
	}
	return nil
}

type ListUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListUsersRequest) Reset() {
	*x = ListUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_users_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersRequest) ProtoMessage() {}

func (x *ListUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_users_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersRequest.ProtoReflect.Descriptor instead.
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return file_api_users_v1_users_proto_rawDescGZIP(), []int{11}
}

type ListUsersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsersReply []*UsersReply `protobuf:"bytes,1,rep,name=usersReply,proto3" json:"usersReply,omitempty"`
}

func (x *ListUsersReply) Reset() {
	*x = ListUsersReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_users_v1_users_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersReply) ProtoMessage() {}

func (x *ListUsersReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_users_v1_users_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersReply.ProtoReflect.Descriptor instead.
func (*ListUsersReply) Descriptor() ([]byte, []int) {
	return file_api_users_v1_users_proto_rawDescGZIP(), []int{12}
}

func (x *ListUsersReply) GetUsersReply() []*UsersReply {
	if x != nil {
		return x.UsersReply
	}
	return nil
}

var File_api_users_v1_users_proto protoreflect.FileDescriptor

var file_api_users_v1_users_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x33, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xce, 0x01,
	0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x43, 0x61, 0x72, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x72, 0x75,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x72,
	0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x72, 0x65, 0x72,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x61, 0x72, 0x72, 0x65, 0x72, 0x22, 0xec,
	0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x64, 0x5f,
	0x63, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x43, 0x61,
	0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x72, 0x75, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x72, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72,
	0x72, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x61, 0x72, 0x72, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x54, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x5e, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x72, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x63, 0x61, 0x72, 0x72, 0x65, 0x72, 0x22, 0x43, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x05,
	0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x43, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x38, 0x0a, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x32, 0x8d, 0x03, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x4f,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x20, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x4f, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x4f, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x46, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x49, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x42, 0x29, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x17, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_users_v1_users_proto_rawDescOnce sync.Once
	file_api_users_v1_users_proto_rawDescData = file_api_users_v1_users_proto_rawDesc
)

func file_api_users_v1_users_proto_rawDescGZIP() []byte {
	file_api_users_v1_users_proto_rawDescOnce.Do(func() {
		file_api_users_v1_users_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_users_v1_users_proto_rawDescData)
	})
	return file_api_users_v1_users_proto_rawDescData
}

var file_api_users_v1_users_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_users_v1_users_proto_goTypes = []interface{}{
	(*CommonReply)(nil),        // 0: api.users.v1.CommonReply
	(*UsersRequest)(nil),       // 1: api.users.v1.UsersRequest
	(*UsersReply)(nil),         // 2: api.users.v1.UsersReply
	(*CreateUsersRequest)(nil), // 3: api.users.v1.CreateUsersRequest
	(*CreateUsersReply)(nil),   // 4: api.users.v1.CreateUsersReply
	(*UpdateUsersRequest)(nil), // 5: api.users.v1.UpdateUsersRequest
	(*UpdateUsersReply)(nil),   // 6: api.users.v1.UpdateUsersReply
	(*DeleteUsersRequest)(nil), // 7: api.users.v1.DeleteUsersRequest
	(*DeleteUsersReply)(nil),   // 8: api.users.v1.DeleteUsersReply
	(*GetUsersRequest)(nil),    // 9: api.users.v1.GetUsersRequest
	(*GetUsersReply)(nil),      // 10: api.users.v1.GetUsersReply
	(*ListUsersRequest)(nil),   // 11: api.users.v1.ListUsersRequest
	(*ListUsersReply)(nil),     // 12: api.users.v1.ListUsersReply
}
var file_api_users_v1_users_proto_depIdxs = []int32{
	1,  // 0: api.users.v1.CreateUsersRequest.UsersRequest:type_name -> api.users.v1.UsersRequest
	0,  // 1: api.users.v1.CreateUsersReply.reply:type_name -> api.users.v1.CommonReply
	0,  // 2: api.users.v1.UpdateUsersReply.reply:type_name -> api.users.v1.CommonReply
	0,  // 3: api.users.v1.DeleteUsersReply.reply:type_name -> api.users.v1.CommonReply
	2,  // 4: api.users.v1.GetUsersReply.userReply:type_name -> api.users.v1.UsersReply
	2,  // 5: api.users.v1.ListUsersReply.usersReply:type_name -> api.users.v1.UsersReply
	3,  // 6: api.users.v1.Users.CreateUsers:input_type -> api.users.v1.CreateUsersRequest
	5,  // 7: api.users.v1.Users.UpdateUsers:input_type -> api.users.v1.UpdateUsersRequest
	7,  // 8: api.users.v1.Users.DeleteUsers:input_type -> api.users.v1.DeleteUsersRequest
	9,  // 9: api.users.v1.Users.GetUsers:input_type -> api.users.v1.GetUsersRequest
	11, // 10: api.users.v1.Users.ListUsers:input_type -> api.users.v1.ListUsersRequest
	4,  // 11: api.users.v1.Users.CreateUsers:output_type -> api.users.v1.CreateUsersReply
	6,  // 12: api.users.v1.Users.UpdateUsers:output_type -> api.users.v1.UpdateUsersReply
	8,  // 13: api.users.v1.Users.DeleteUsers:output_type -> api.users.v1.DeleteUsersReply
	10, // 14: api.users.v1.Users.GetUsers:output_type -> api.users.v1.GetUsersReply
	12, // 15: api.users.v1.Users.ListUsers:output_type -> api.users.v1.ListUsersReply
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_users_v1_users_proto_init() }
func file_api_users_v1_users_proto_init() {
	if File_api_users_v1_users_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_users_v1_users_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonReply); i {
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
		file_api_users_v1_users_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersRequest); i {
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
		file_api_users_v1_users_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersReply); i {
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
		file_api_users_v1_users_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUsersRequest); i {
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
		file_api_users_v1_users_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUsersReply); i {
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
		file_api_users_v1_users_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUsersRequest); i {
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
		file_api_users_v1_users_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUsersReply); i {
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
		file_api_users_v1_users_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUsersRequest); i {
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
		file_api_users_v1_users_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUsersReply); i {
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
		file_api_users_v1_users_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersRequest); i {
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
		file_api_users_v1_users_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersReply); i {
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
		file_api_users_v1_users_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersRequest); i {
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
		file_api_users_v1_users_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersReply); i {
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
			RawDescriptor: file_api_users_v1_users_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_users_v1_users_proto_goTypes,
		DependencyIndexes: file_api_users_v1_users_proto_depIdxs,
		MessageInfos:      file_api_users_v1_users_proto_msgTypes,
	}.Build()
	File_api_users_v1_users_proto = out.File
	file_api_users_v1_users_proto_rawDesc = nil
	file_api_users_v1_users_proto_goTypes = nil
	file_api_users_v1_users_proto_depIdxs = nil
}
