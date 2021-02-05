// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as/external/api/serviceProfile.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateServiceProfileRequest struct {
	// Service-profile object to create.
	ServiceProfile       *ServiceProfile `protobuf:"bytes,1,opt,name=service_profile,json=serviceProfile,proto3" json:"service_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateServiceProfileRequest) Reset()         { *m = CreateServiceProfileRequest{} }
func (m *CreateServiceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*CreateServiceProfileRequest) ProtoMessage()    {}
func (*CreateServiceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f18fca066634371, []int{0}
}

func (m *CreateServiceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceProfileRequest.Unmarshal(m, b)
}
func (m *CreateServiceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceProfileRequest.Marshal(b, m, deterministic)
}
func (m *CreateServiceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceProfileRequest.Merge(m, src)
}
func (m *CreateServiceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_CreateServiceProfileRequest.Size(m)
}
func (m *CreateServiceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceProfileRequest proto.InternalMessageInfo

func (m *CreateServiceProfileRequest) GetServiceProfile() *ServiceProfile {
	if m != nil {
		return m.ServiceProfile
	}
	return nil
}

type CreateServiceProfileResponse struct {
	// Service-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServiceProfileResponse) Reset()         { *m = CreateServiceProfileResponse{} }
func (m *CreateServiceProfileResponse) String() string { return proto.CompactTextString(m) }
func (*CreateServiceProfileResponse) ProtoMessage()    {}
func (*CreateServiceProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f18fca066634371, []int{1}
}

func (m *CreateServiceProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServiceProfileResponse.Unmarshal(m, b)
}
func (m *CreateServiceProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServiceProfileResponse.Marshal(b, m, deterministic)
}
func (m *CreateServiceProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServiceProfileResponse.Merge(m, src)
}
func (m *CreateServiceProfileResponse) XXX_Size() int {
	return xxx_messageInfo_CreateServiceProfileResponse.Size(m)
}
func (m *CreateServiceProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServiceProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServiceProfileResponse proto.InternalMessageInfo

func (m *CreateServiceProfileResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetServiceProfileRequest struct {
	// Service-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceProfileRequest) Reset()         { *m = GetServiceProfileRequest{} }
func (m *GetServiceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetServiceProfileRequest) ProtoMessage()    {}
func (*GetServiceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f18fca066634371, []int{2}
}

func (m *GetServiceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceProfileRequest.Unmarshal(m, b)
}
func (m *GetServiceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceProfileRequest.Marshal(b, m, deterministic)
}
func (m *GetServiceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceProfileRequest.Merge(m, src)
}
func (m *GetServiceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetServiceProfileRequest.Size(m)
}
func (m *GetServiceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceProfileRequest proto.InternalMessageInfo

func (m *GetServiceProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetServiceProfileResponse struct {
	// Service-profile object.
	ServiceProfile *ServiceProfile `protobuf:"bytes,1,opt,name=service_profile,json=serviceProfile,proto3" json:"service_profile,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetServiceProfileResponse) Reset()         { *m = GetServiceProfileResponse{} }
func (m *GetServiceProfileResponse) String() string { return proto.CompactTextString(m) }
func (*GetServiceProfileResponse) ProtoMessage()    {}
func (*GetServiceProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f18fca066634371, []int{3}
}

func (m *GetServiceProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceProfileResponse.Unmarshal(m, b)
}
func (m *GetServiceProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceProfileResponse.Marshal(b, m, deterministic)
}
func (m *GetServiceProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceProfileResponse.Merge(m, src)
}
func (m *GetServiceProfileResponse) XXX_Size() int {
	return xxx_messageInfo_GetServiceProfileResponse.Size(m)
}
func (m *GetServiceProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceProfileResponse proto.InternalMessageInfo

func (m *GetServiceProfileResponse) GetServiceProfile() *ServiceProfile {
	if m != nil {
		return m.ServiceProfile
	}
	return nil
}

func (m *GetServiceProfileResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GetServiceProfileResponse) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type UpdateServiceProfileRequest struct {
	// Service-profile object to update.
	ServiceProfile       *ServiceProfile `protobuf:"bytes,1,opt,name=service_profile,json=serviceProfile,proto3" json:"service_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateServiceProfileRequest) Reset()         { *m = UpdateServiceProfileRequest{} }
func (m *UpdateServiceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateServiceProfileRequest) ProtoMessage()    {}
func (*UpdateServiceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f18fca066634371, []int{4}
}

func (m *UpdateServiceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateServiceProfileRequest.Unmarshal(m, b)
}
func (m *UpdateServiceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateServiceProfileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateServiceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateServiceProfileRequest.Merge(m, src)
}
func (m *UpdateServiceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateServiceProfileRequest.Size(m)
}
func (m *UpdateServiceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateServiceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateServiceProfileRequest proto.InternalMessageInfo

func (m *UpdateServiceProfileRequest) GetServiceProfile() *ServiceProfile {
	if m != nil {
		return m.ServiceProfile
	}
	return nil
}

type DeleteServiceProfileRequest struct {
	// Service-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteServiceProfileRequest) Reset()         { *m = DeleteServiceProfileRequest{} }
func (m *DeleteServiceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteServiceProfileRequest) ProtoMessage()    {}
func (*DeleteServiceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f18fca066634371, []int{5}
}

func (m *DeleteServiceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteServiceProfileRequest.Unmarshal(m, b)
}
func (m *DeleteServiceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteServiceProfileRequest.Marshal(b, m, deterministic)
}
func (m *DeleteServiceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteServiceProfileRequest.Merge(m, src)
}
func (m *DeleteServiceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteServiceProfileRequest.Size(m)
}
func (m *DeleteServiceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteServiceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteServiceProfileRequest proto.InternalMessageInfo

func (m *DeleteServiceProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListServiceProfileRequest struct {
	// Max number of items to return.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Organization id to filter on.
	OrganizationId int64 `protobuf:"varint,3,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Network-server id to filter on.
	NetworkServerId      int64    `protobuf:"varint,4,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListServiceProfileRequest) Reset()         { *m = ListServiceProfileRequest{} }
func (m *ListServiceProfileRequest) String() string { return proto.CompactTextString(m) }
func (*ListServiceProfileRequest) ProtoMessage()    {}
func (*ListServiceProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f18fca066634371, []int{6}
}

func (m *ListServiceProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServiceProfileRequest.Unmarshal(m, b)
}
func (m *ListServiceProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServiceProfileRequest.Marshal(b, m, deterministic)
}
func (m *ListServiceProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServiceProfileRequest.Merge(m, src)
}
func (m *ListServiceProfileRequest) XXX_Size() int {
	return xxx_messageInfo_ListServiceProfileRequest.Size(m)
}
func (m *ListServiceProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServiceProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListServiceProfileRequest proto.InternalMessageInfo

func (m *ListServiceProfileRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListServiceProfileRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListServiceProfileRequest) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *ListServiceProfileRequest) GetNetworkServerId() int64 {
	if m != nil {
		return m.NetworkServerId
	}
	return 0
}

type ServiceProfileListItem struct {
	// Service-profile ID (UUID string).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Service-profile name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Organization ID of the service-profile.
	OrganizationId int64 `protobuf:"varint,3,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Network-server ID of the service-profile.
	NetworkServerId int64 `protobuf:"varint,4,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Network-server name of the service-profile.
	NetworkServerName    string   `protobuf:"bytes,7,opt,name=network_server_name,json=networkServerName,proto3" json:"network_server_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceProfileListItem) Reset()         { *m = ServiceProfileListItem{} }
func (m *ServiceProfileListItem) String() string { return proto.CompactTextString(m) }
func (*ServiceProfileListItem) ProtoMessage()    {}
func (*ServiceProfileListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f18fca066634371, []int{7}
}

func (m *ServiceProfileListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceProfileListItem.Unmarshal(m, b)
}
func (m *ServiceProfileListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceProfileListItem.Marshal(b, m, deterministic)
}
func (m *ServiceProfileListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceProfileListItem.Merge(m, src)
}
func (m *ServiceProfileListItem) XXX_Size() int {
	return xxx_messageInfo_ServiceProfileListItem.Size(m)
}
func (m *ServiceProfileListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceProfileListItem.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceProfileListItem proto.InternalMessageInfo

func (m *ServiceProfileListItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceProfileListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceProfileListItem) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *ServiceProfileListItem) GetNetworkServerId() int64 {
	if m != nil {
		return m.NetworkServerId
	}
	return 0
}

func (m *ServiceProfileListItem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ServiceProfileListItem) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *ServiceProfileListItem) GetNetworkServerName() string {
	if m != nil {
		return m.NetworkServerName
	}
	return ""
}

type ListServiceProfileResponse struct {
	// Total number of service-profiles.
	TotalCount           int64                     `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Result               []*ServiceProfileListItem `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ListServiceProfileResponse) Reset()         { *m = ListServiceProfileResponse{} }
func (m *ListServiceProfileResponse) String() string { return proto.CompactTextString(m) }
func (*ListServiceProfileResponse) ProtoMessage()    {}
func (*ListServiceProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f18fca066634371, []int{8}
}

func (m *ListServiceProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServiceProfileResponse.Unmarshal(m, b)
}
func (m *ListServiceProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServiceProfileResponse.Marshal(b, m, deterministic)
}
func (m *ListServiceProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServiceProfileResponse.Merge(m, src)
}
func (m *ListServiceProfileResponse) XXX_Size() int {
	return xxx_messageInfo_ListServiceProfileResponse.Size(m)
}
func (m *ListServiceProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServiceProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListServiceProfileResponse proto.InternalMessageInfo

func (m *ListServiceProfileResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListServiceProfileResponse) GetResult() []*ServiceProfileListItem {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateServiceProfileRequest)(nil), "api.CreateServiceProfileRequest")
	proto.RegisterType((*CreateServiceProfileResponse)(nil), "api.CreateServiceProfileResponse")
	proto.RegisterType((*GetServiceProfileRequest)(nil), "api.GetServiceProfileRequest")
	proto.RegisterType((*GetServiceProfileResponse)(nil), "api.GetServiceProfileResponse")
	proto.RegisterType((*UpdateServiceProfileRequest)(nil), "api.UpdateServiceProfileRequest")
	proto.RegisterType((*DeleteServiceProfileRequest)(nil), "api.DeleteServiceProfileRequest")
	proto.RegisterType((*ListServiceProfileRequest)(nil), "api.ListServiceProfileRequest")
	proto.RegisterType((*ServiceProfileListItem)(nil), "api.ServiceProfileListItem")
	proto.RegisterType((*ListServiceProfileResponse)(nil), "api.ListServiceProfileResponse")
}

func init() {
	proto.RegisterFile("as/external/api/serviceProfile.proto", fileDescriptor_4f18fca066634371)
}

var fileDescriptor_4f18fca066634371 = []byte{
	// 687 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdf, 0x4e, 0xd4, 0x4e,
	0x14, 0xce, 0x6e, 0x97, 0xfe, 0xc2, 0x21, 0x81, 0x30, 0xfc, 0x40, 0xe8, 0xf2, 0xb7, 0x31, 0x91,
	0x6c, 0x42, 0x9b, 0x40, 0xbc, 0x90, 0x78, 0x83, 0x60, 0x08, 0x46, 0x8d, 0x29, 0x7a, 0xa3, 0x17,
	0x9b, 0x61, 0x7b, 0xba, 0x3b, 0xa1, 0xed, 0xd4, 0x76, 0x8a, 0xa2, 0xe1, 0xc6, 0x57, 0xf0, 0xde,
	0xd7, 0xf0, 0xca, 0x67, 0xf0, 0xc2, 0x57, 0xf0, 0x41, 0xcc, 0x4c, 0xa7, 0x91, 0xad, 0xed, 0x1a,
	0x35, 0xdc, 0xf5, 0xcc, 0x7c, 0xe7, 0x9c, 0xef, 0x9c, 0x6f, 0xce, 0x29, 0xdc, 0xa6, 0x99, 0x8b,
	0x6f, 0x05, 0xa6, 0x31, 0x0d, 0x5d, 0x9a, 0x30, 0x37, 0xc3, 0xf4, 0x82, 0x0d, 0xf0, 0x59, 0xca,
	0x03, 0x16, 0xa2, 0x93, 0xa4, 0x5c, 0x70, 0x62, 0xd0, 0x84, 0x59, 0xab, 0x43, 0xce, 0x87, 0x21,
	0x2a, 0x14, 0x8d, 0x63, 0x2e, 0xa8, 0x60, 0x3c, 0xce, 0x0a, 0x88, 0xb5, 0xa1, 0x6f, 0x95, 0x75,
	0x96, 0x07, 0xae, 0x60, 0x11, 0x66, 0x82, 0x46, 0x89, 0x06, 0x74, 0xab, 0x00, 0x8c, 0x12, 0x71,
	0xa9, 0x2f, 0xd7, 0xab, 0x34, 0x92, 0x22, 0xbf, 0x8e, 0x6e, 0xbf, 0x82, 0xee, 0x61, 0x8a, 0x54,
	0xe0, 0xe9, 0x18, 0x3d, 0x0f, 0x5f, 0xe7, 0x98, 0x09, 0x72, 0x1f, 0xe6, 0x34, 0xef, 0xbe, 0x76,
	0x5c, 0x6e, 0x6d, 0xb6, 0xb6, 0x67, 0x76, 0x17, 0x1c, 0x9a, 0x30, 0xa7, 0xe2, 0x34, 0x3b, 0x5e,
	0xa3, 0xed, 0xc0, 0x6a, 0x7d, 0xf0, 0x2c, 0xe1, 0x71, 0x86, 0x64, 0x16, 0xda, 0xcc, 0x57, 0x01,
	0xa7, 0xbd, 0x36, 0xf3, 0xed, 0x1e, 0x2c, 0x1f, 0xa3, 0xa8, 0x67, 0x52, 0xc5, 0x7e, 0x6d, 0xc1,
	0x4a, 0x0d, 0x58, 0x47, 0xfe, 0x27, 0xde, 0xe4, 0x1e, 0xc0, 0x40, 0xf1, 0xf6, 0xfb, 0x54, 0x2c,
	0xb7, 0x95, 0xa3, 0xe5, 0x14, 0x6d, 0x76, 0xca, 0x36, 0x3b, 0xcf, 0x4b, 0x1d, 0xbc, 0x69, 0x8d,
	0x3e, 0x10, 0xd2, 0x35, 0x4f, 0xfc, 0xd2, 0xd5, 0xf8, 0xbd, 0xab, 0x46, 0x1f, 0x08, 0x29, 0xc5,
	0x0b, 0x65, 0xdc, 0x84, 0x14, 0x3b, 0xd0, 0x3d, 0xc2, 0x10, 0x9b, 0x82, 0x57, 0xbb, 0xfb, 0xa9,
	0x05, 0x2b, 0x8f, 0x59, 0xd6, 0xa0, 0xc5, 0xff, 0x30, 0x15, 0xb2, 0x88, 0x09, 0xe5, 0x60, 0x78,
	0x85, 0x41, 0x96, 0xc0, 0xe4, 0x41, 0x90, 0x61, 0xd1, 0x31, 0xc3, 0xd3, 0x16, 0xb9, 0x03, 0x73,
	0x3c, 0x1d, 0xd2, 0x98, 0xbd, 0x53, 0xef, 0xba, 0xcf, 0x7c, 0xd5, 0x17, 0xc3, 0x9b, 0xbd, 0x7e,
	0x7c, 0x72, 0x44, 0x7a, 0x30, 0x1f, 0xa3, 0x78, 0xc3, 0xd3, 0xf3, 0xbe, 0x64, 0x8f, 0xa9, 0x84,
	0x76, 0x14, 0x74, 0x4e, 0x5f, 0x9c, 0xaa, 0xf3, 0x93, 0x23, 0xfb, 0x73, 0x1b, 0x96, 0xc6, 0xc9,
	0x49, 0xba, 0x27, 0x02, 0xa3, 0x6a, 0x2d, 0x84, 0x40, 0x27, 0xa6, 0x11, 0x2a, 0x56, 0xd3, 0x9e,
	0xfa, 0xbe, 0x11, 0x4e, 0x95, 0x67, 0x33, 0xf5, 0xf7, 0xcf, 0xc6, 0xfc, 0x83, 0x67, 0x43, 0x1c,
	0x58, 0xa8, 0x30, 0x54, 0xd5, 0xfe, 0xa7, 0xaa, 0x9d, 0x1f, 0xe3, 0xf8, 0x94, 0x46, 0x68, 0xa7,
	0x60, 0xd5, 0x29, 0xab, 0x07, 0x67, 0x03, 0x66, 0x04, 0x17, 0x34, 0xec, 0x0f, 0x78, 0x1e, 0x97,
	0x02, 0x83, 0x3a, 0x3a, 0x94, 0x27, 0x64, 0x0f, 0xcc, 0x14, 0xb3, 0x3c, 0x94, 0x2a, 0x1b, 0xdb,
	0x33, 0xbb, 0xdd, 0x9a, 0xd7, 0x57, 0x4a, 0xe1, 0x69, 0xe8, 0xee, 0x97, 0x0e, 0x2c, 0x8e, 0x43,
	0xb4, 0x45, 0x38, 0x98, 0xc5, 0x8a, 0x20, 0x9b, 0x2a, 0xd0, 0x84, 0x65, 0x64, 0x6d, 0x4d, 0x40,
	0x14, 0xf4, 0xed, 0xcd, 0x0f, 0xdf, 0xbe, 0x7f, 0x6c, 0x5b, 0xf6, 0xe2, 0xf5, 0x95, 0xbb, 0x53,
	0xee, 0xbc, 0xfd, 0x56, 0x8f, 0x8c, 0xc0, 0x38, 0x46, 0x41, 0xd6, 0x54, 0xac, 0xa6, 0x6d, 0x63,
	0xad, 0x37, 0x5d, 0xeb, 0x3c, 0xb6, 0xca, 0xb3, 0x4a, 0xac, 0xda, 0x3c, 0xee, 0x7b, 0xe6, 0x5f,
	0x91, 0x4b, 0x30, 0x8b, 0x79, 0xd6, 0xa5, 0x4d, 0x18, 0x6e, 0x6b, 0xe9, 0x17, 0xad, 0x1f, 0xca,
	0x25, 0x6e, 0xdf, 0x55, 0x79, 0x5c, 0xab, 0xd7, 0x90, 0xa7, 0xb2, 0x11, 0x1c, 0xe6, 0x5f, 0xc9,
	0x22, 0x03, 0x30, 0x8b, 0x69, 0xd7, 0xa9, 0x27, 0x8c, 0x7e, 0x63, 0x6a, 0x5d, 0x62, 0x6f, 0x52,
	0x89, 0x01, 0x74, 0xa4, 0xd6, 0xa4, 0x68, 0x57, 0xe3, 0xc2, 0xb0, 0x36, 0x1a, 0xef, 0x75, 0x3f,
	0xd7, 0x54, 0xb2, 0x5b, 0xa4, 0x5e, 0xb7, 0x07, 0x7d, 0xd8, 0x62, 0xdc, 0x19, 0x8c, 0x58, 0x9a,
	0x64, 0x82, 0x0e, 0xce, 0x55, 0x38, 0x9a, 0x39, 0xe5, 0x9f, 0x4d, 0xda, 0x2f, 0xf7, 0x87, 0x4c,
	0x8c, 0xf2, 0x33, 0x67, 0xc0, 0x23, 0xf7, 0x11, 0x7d, 0x92, 0xa7, 0xc9, 0xe8, 0x34, 0x62, 0xee,
	0x4f, 0xaf, 0x1d, 0x19, 0x7b, 0xc8, 0xdd, 0x8b, 0x3d, 0xb7, 0xf2, 0x57, 0x3c, 0x33, 0x55, 0xf1,
	0x7b, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x95, 0x28, 0x34, 0x1b, 0xb6, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceProfileServiceClient is the client API for ServiceProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceProfileServiceClient interface {
	// Create creates the given service-profile.
	Create(ctx context.Context, in *CreateServiceProfileRequest, opts ...grpc.CallOption) (*CreateServiceProfileResponse, error)
	// Get returns the service-profile matching the given id.
	Get(ctx context.Context, in *GetServiceProfileRequest, opts ...grpc.CallOption) (*GetServiceProfileResponse, error)
	// Update updates the given serviceprofile.
	Update(ctx context.Context, in *UpdateServiceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete deletes the service-profile matching the given id.
	Delete(ctx context.Context, in *DeleteServiceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// List lists the available service-profiles.
	List(ctx context.Context, in *ListServiceProfileRequest, opts ...grpc.CallOption) (*ListServiceProfileResponse, error)
}

type serviceProfileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceProfileServiceClient(cc grpc.ClientConnInterface) ServiceProfileServiceClient {
	return &serviceProfileServiceClient{cc}
}

func (c *serviceProfileServiceClient) Create(ctx context.Context, in *CreateServiceProfileRequest, opts ...grpc.CallOption) (*CreateServiceProfileResponse, error) {
	out := new(CreateServiceProfileResponse)
	err := c.cc.Invoke(ctx, "/api.ServiceProfileService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProfileServiceClient) Get(ctx context.Context, in *GetServiceProfileRequest, opts ...grpc.CallOption) (*GetServiceProfileResponse, error) {
	out := new(GetServiceProfileResponse)
	err := c.cc.Invoke(ctx, "/api.ServiceProfileService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProfileServiceClient) Update(ctx context.Context, in *UpdateServiceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.ServiceProfileService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProfileServiceClient) Delete(ctx context.Context, in *DeleteServiceProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.ServiceProfileService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceProfileServiceClient) List(ctx context.Context, in *ListServiceProfileRequest, opts ...grpc.CallOption) (*ListServiceProfileResponse, error) {
	out := new(ListServiceProfileResponse)
	err := c.cc.Invoke(ctx, "/api.ServiceProfileService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceProfileServiceServer is the server API for ServiceProfileService service.
type ServiceProfileServiceServer interface {
	// Create creates the given service-profile.
	Create(context.Context, *CreateServiceProfileRequest) (*CreateServiceProfileResponse, error)
	// Get returns the service-profile matching the given id.
	Get(context.Context, *GetServiceProfileRequest) (*GetServiceProfileResponse, error)
	// Update updates the given serviceprofile.
	Update(context.Context, *UpdateServiceProfileRequest) (*empty.Empty, error)
	// Delete deletes the service-profile matching the given id.
	Delete(context.Context, *DeleteServiceProfileRequest) (*empty.Empty, error)
	// List lists the available service-profiles.
	List(context.Context, *ListServiceProfileRequest) (*ListServiceProfileResponse, error)
}

// UnimplementedServiceProfileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceProfileServiceServer struct {
}

func (*UnimplementedServiceProfileServiceServer) Create(ctx context.Context, req *CreateServiceProfileRequest) (*CreateServiceProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedServiceProfileServiceServer) Get(ctx context.Context, req *GetServiceProfileRequest) (*GetServiceProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedServiceProfileServiceServer) Update(ctx context.Context, req *UpdateServiceProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedServiceProfileServiceServer) Delete(ctx context.Context, req *DeleteServiceProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedServiceProfileServiceServer) List(ctx context.Context, req *ListServiceProfileRequest) (*ListServiceProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterServiceProfileServiceServer(s *grpc.Server, srv ServiceProfileServiceServer) {
	s.RegisterService(&_ServiceProfileService_serviceDesc, srv)
}

func _ServiceProfileService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProfileServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceProfileService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProfileServiceServer).Create(ctx, req.(*CreateServiceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProfileService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProfileServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceProfileService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProfileServiceServer).Get(ctx, req.(*GetServiceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProfileService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServiceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProfileServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceProfileService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProfileServiceServer).Update(ctx, req.(*UpdateServiceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProfileService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProfileServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceProfileService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProfileServiceServer).Delete(ctx, req.(*DeleteServiceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceProfileService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServiceProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceProfileServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ServiceProfileService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceProfileServiceServer).List(ctx, req.(*ListServiceProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ServiceProfileService",
	HandlerType: (*ServiceProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ServiceProfileService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ServiceProfileService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ServiceProfileService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ServiceProfileService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ServiceProfileService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as/external/api/serviceProfile.proto",
}
