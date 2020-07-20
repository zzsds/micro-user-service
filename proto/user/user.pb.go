// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Enabled int32

const (
	Enabled_No  Enabled = 0
	Enabled_Yes Enabled = 1
)

var Enabled_name = map[int32]string{
	0: "No",
	1: "Yes",
}

var Enabled_value = map[string]int32{
	"No":  0,
	"Yes": 1,
}

func (x Enabled) String() string {
	return proto.EnumName(Enabled_name, int32(x))
}

func (Enabled) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

type Resource struct {
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Id                   int32                `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Nickname             string               `protobuf:"bytes,5,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Realname             string               `protobuf:"bytes,6,opt,name=realname,proto3" json:"realname,omitempty"`
	Mobile               string               `protobuf:"bytes,7,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email                string               `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	Enabled              Enabled              `protobuf:"varint,9,opt,name=enabled,proto3,enum=srv.user.Enabled" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Resource) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Resource) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Resource) GetRealname() string {
	if m != nil {
		return m.Realname
	}
	return ""
}

func (m *Resource) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Resource) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Resource) GetEnabled() Enabled {
	if m != nil {
		return m.Enabled
	}
	return Enabled_No
}

type Pagination struct {
	Total                int32    `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pagination) Reset()         { *m = Pagination{} }
func (m *Pagination) String() string { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()    {}
func (*Pagination) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Pagination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pagination.Unmarshal(m, b)
}
func (m *Pagination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pagination.Marshal(b, m, deterministic)
}
func (m *Pagination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pagination.Merge(m, src)
}
func (m *Pagination) XXX_Size() int {
	return xxx_messageInfo_Pagination.Size(m)
}
func (m *Pagination) XXX_DiscardUnknown() {
	xxx_messageInfo_Pagination.DiscardUnknown(m)
}

var xxx_messageInfo_Pagination proto.InternalMessageInfo

func (m *Pagination) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Pagination) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Pagination) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type UserList struct {
	Data                 []*Resource `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetData() []*Resource {
	if m != nil {
		return m.Data
	}
	return nil
}

type LoginRequest struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type LoginResponse struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type RegisterRequest struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type RegisterResponse struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{6}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

func init() {
	proto.RegisterEnum("srv.user.Enabled", Enabled_name, Enabled_value)
	proto.RegisterType((*Resource)(nil), "srv.user.Resource")
	proto.RegisterType((*Pagination)(nil), "srv.user.Pagination")
	proto.RegisterType((*UserList)(nil), "srv.user.UserList")
	proto.RegisterType((*LoginRequest)(nil), "srv.user.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "srv.user.LoginResponse")
	proto.RegisterType((*RegisterRequest)(nil), "srv.user.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "srv.user.RegisterResponse")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x51, 0x6f, 0x94, 0x4c,
	0x14, 0x5d, 0x96, 0x85, 0x85, 0xbb, 0xfd, 0xfa, 0xad, 0x93, 0x6a, 0x90, 0x17, 0x11, 0x8d, 0x21,
	0x9a, 0xb0, 0x09, 0x7d, 0xf1, 0xb5, 0xb1, 0x5b, 0x53, 0x53, 0x8d, 0x99, 0xe8, 0x83, 0x8f, 0xb3,
	0xcb, 0x95, 0x4c, 0x04, 0x06, 0x99, 0xc1, 0xa8, 0xbf, 0xc5, 0xbf, 0xe1, 0xff, 0x33, 0x33, 0x40,
	0xa9, 0x49, 0x1b, 0x93, 0xf6, 0x65, 0x73, 0xcf, 0x3d, 0x67, 0xce, 0xce, 0x39, 0x13, 0xe0, 0x7e,
	0xd3, 0x0a, 0x25, 0x36, 0x9d, 0xc4, 0xd6, 0xfc, 0xa4, 0x06, 0x13, 0x4f, 0xb6, 0xdf, 0x52, 0x8d,
	0xc3, 0x47, 0x85, 0x10, 0x45, 0x89, 0x1b, 0xb3, 0xdf, 0x75, 0x9f, 0x37, 0x8a, 0x57, 0x28, 0x15,
	0xab, 0x9a, 0x5e, 0x1a, 0xff, 0x9e, 0x83, 0x47, 0x51, 0x8a, 0xae, 0xdd, 0x23, 0x79, 0x09, 0xfe,
	0xbe, 0x45, 0xa6, 0x30, 0x3f, 0x51, 0x81, 0x15, 0x59, 0xc9, 0x2a, 0x0b, 0xd3, 0xde, 0x21, 0x1d,
	0x1d, 0xd2, 0x0f, 0xa3, 0x03, 0x9d, 0xc4, 0xfa, 0x64, 0xd7, 0xe4, 0xc3, 0xc9, 0xf9, 0xbf, 0x4f,
	0x5e, 0x8a, 0xc9, 0x21, 0xcc, 0x79, 0x1e, 0xd8, 0x91, 0x95, 0x38, 0x74, 0xce, 0x73, 0x42, 0x60,
	0x51, 0xb3, 0x0a, 0x83, 0x45, 0x64, 0x25, 0x3e, 0x35, 0x33, 0x09, 0xc1, 0xab, 0xf9, 0xfe, 0x8b,
	0xd9, 0x3b, 0x66, 0x7f, 0x89, 0x35, 0xd7, 0x22, 0x2b, 0x0d, 0xe7, 0xf6, 0xdc, 0x88, 0xc9, 0x03,
	0x70, 0x2b, 0xb1, 0xe3, 0x25, 0x06, 0x4b, 0xc3, 0x0c, 0x88, 0x1c, 0x81, 0x83, 0x15, 0xe3, 0x65,
	0xe0, 0x99, 0x75, 0x0f, 0xc8, 0x0b, 0x58, 0x62, 0xcd, 0x76, 0x25, 0xe6, 0x81, 0x1f, 0x59, 0xc9,
	0x61, 0x76, 0x2f, 0x1d, 0x7b, 0x4c, 0xb7, 0x3d, 0x41, 0x47, 0x45, 0xfc, 0x06, 0xe0, 0x3d, 0x2b,
	0x78, 0xcd, 0x14, 0x17, 0xb5, 0x36, 0x54, 0x42, 0xb1, 0xd2, 0x94, 0xe6, 0xd0, 0x1e, 0xe8, 0x28,
	0x0d, 0x2b, 0xd0, 0xf4, 0xe1, 0x50, 0x33, 0xeb, 0x9d, 0xe4, 0x3f, 0x71, 0x08, 0x6c, 0xe6, 0x38,
	0x03, 0xef, 0xa3, 0xc4, 0xf6, 0x82, 0x4b, 0x45, 0x9e, 0xc1, 0x22, 0x67, 0x8a, 0x05, 0x56, 0x64,
	0x27, 0xab, 0x8c, 0x4c, 0x37, 0x18, 0x1f, 0x89, 0x1a, 0x3e, 0x8e, 0xe0, 0xe0, 0x42, 0x14, 0xbc,
	0xa6, 0xf8, 0xb5, 0x43, 0xa9, 0xc8, 0x1a, 0x6c, 0xc9, 0x7e, 0x98, 0xff, 0xf7, 0xa9, 0x1e, 0xe3,
	0xc7, 0xf0, 0xdf, 0xa0, 0x90, 0x8d, 0xa8, 0x25, 0x5e, 0x23, 0x79, 0x02, 0xff, 0x53, 0x2c, 0xb8,
	0x54, 0xd8, 0xde, 0xec, 0xf3, 0x14, 0xd6, 0x93, 0xe8, 0x26, 0xab, 0xe7, 0x21, 0x2c, 0x87, 0x8e,
	0x88, 0x0b, 0xf3, 0x77, 0x62, 0x3d, 0x23, 0x4b, 0xb0, 0x3f, 0xa1, 0x5c, 0x5b, 0xd9, 0x2f, 0x1b,
	0x16, 0x3a, 0x20, 0x39, 0x06, 0xe7, 0xbc, 0xce, 0xf1, 0x3b, 0x39, 0x9a, 0x72, 0x4d, 0x2d, 0x86,
	0x57, 0xd2, 0x8e, 0x7d, 0xc4, 0x33, 0xb2, 0x05, 0x38, 0xe3, 0x75, 0xfe, 0xb6, 0x7f, 0xba, 0x87,
	0x57, 0x1b, 0xf9, 0xeb, 0xea, 0x61, 0x78, 0x1d, 0xd5, 0x5f, 0x38, 0x9e, 0x91, 0x53, 0xf0, 0xb5,
	0xcd, 0xd6, 0x3c, 0xf5, 0xad, 0x5d, 0x4e, 0xc0, 0xd5, 0x2e, 0xe7, 0xa7, 0xb7, 0xb7, 0x78, 0x0d,
	0x07, 0xaf, 0xcc, 0x77, 0x73, 0xd7, 0x44, 0x67, 0xb0, 0xea, 0x8d, 0xee, 0x96, 0x69, 0xe7, 0x9a,
	0x0f, 0xf4, 0xf8, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0x57, 0xb9, 0xc8, 0x4d, 0x04, 0x00,
	0x00,
}
