// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/projects/project.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/lavanet/lava/v5/x/plans/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ProjectKey_Type int32

const (
	ProjectKey_NONE      ProjectKey_Type = 0
	ProjectKey_ADMIN     ProjectKey_Type = 1
	ProjectKey_DEVELOPER ProjectKey_Type = 2
)

var ProjectKey_Type_name = map[int32]string{
	0: "NONE",
	1: "ADMIN",
	2: "DEVELOPER",
}

var ProjectKey_Type_value = map[string]int32{
	"NONE":      0,
	"ADMIN":     1,
	"DEVELOPER": 2,
}

func (x ProjectKey_Type) String() string {
	return proto.EnumName(ProjectKey_Type_name, int32(x))
}

func (ProjectKey_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9027839604ae2915, []int{1, 0}
}

type Project struct {
	Index              string        `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Subscription       string        `protobuf:"bytes,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Enabled            bool          `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	ProjectKeys        []ProjectKey  `protobuf:"bytes,5,rep,name=project_keys,json=projectKeys,proto3" json:"project_keys"`
	AdminPolicy        *types.Policy `protobuf:"bytes,6,opt,name=admin_policy,json=adminPolicy,proto3" json:"admin_policy"`
	UsedCu             uint64        `protobuf:"varint,7,opt,name=used_cu,json=usedCu,proto3" json:"used_cu,omitempty"`
	SubscriptionPolicy *types.Policy `protobuf:"bytes,8,opt,name=subscription_policy,json=subscriptionPolicy,proto3" json:"subscription_policy"`
	Snapshot           uint64        `protobuf:"varint,9,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_9027839604ae2915, []int{0}
}
func (m *Project) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Project.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return m.Size()
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Project) GetSubscription() string {
	if m != nil {
		return m.Subscription
	}
	return ""
}

func (m *Project) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Project) GetProjectKeys() []ProjectKey {
	if m != nil {
		return m.ProjectKeys
	}
	return nil
}

func (m *Project) GetAdminPolicy() *types.Policy {
	if m != nil {
		return m.AdminPolicy
	}
	return nil
}

func (m *Project) GetUsedCu() uint64 {
	if m != nil {
		return m.UsedCu
	}
	return 0
}

func (m *Project) GetSubscriptionPolicy() *types.Policy {
	if m != nil {
		return m.SubscriptionPolicy
	}
	return nil
}

func (m *Project) GetSnapshot() uint64 {
	if m != nil {
		return m.Snapshot
	}
	return 0
}

type ProjectKey struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key"`
	Kinds uint32 `protobuf:"varint,4,opt,name=kinds,proto3" json:"kinds"`
}

func (m *ProjectKey) Reset()         { *m = ProjectKey{} }
func (m *ProjectKey) String() string { return proto.CompactTextString(m) }
func (*ProjectKey) ProtoMessage()    {}
func (*ProjectKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_9027839604ae2915, []int{1}
}
func (m *ProjectKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProjectKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProjectKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProjectKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectKey.Merge(m, src)
}
func (m *ProjectKey) XXX_Size() int {
	return m.Size()
}
func (m *ProjectKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectKey.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectKey proto.InternalMessageInfo

func (m *ProjectKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ProjectKey) GetKinds() uint32 {
	if m != nil {
		return m.Kinds
	}
	return 0
}

type ProtoDeveloperData struct {
	ProjectID string `protobuf:"bytes,1,opt,name=projectID,proto3" json:"projectID,omitempty"`
}

func (m *ProtoDeveloperData) Reset()         { *m = ProtoDeveloperData{} }
func (m *ProtoDeveloperData) String() string { return proto.CompactTextString(m) }
func (*ProtoDeveloperData) ProtoMessage()    {}
func (*ProtoDeveloperData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9027839604ae2915, []int{2}
}
func (m *ProtoDeveloperData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtoDeveloperData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtoDeveloperData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtoDeveloperData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoDeveloperData.Merge(m, src)
}
func (m *ProtoDeveloperData) XXX_Size() int {
	return m.Size()
}
func (m *ProtoDeveloperData) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoDeveloperData.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoDeveloperData proto.InternalMessageInfo

func (m *ProtoDeveloperData) GetProjectID() string {
	if m != nil {
		return m.ProjectID
	}
	return ""
}

// used as a container struct for the subscription module
type ProjectData struct {
	Name        string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Enabled     bool          `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	ProjectKeys []ProjectKey  `protobuf:"bytes,4,rep,name=projectKeys,proto3" json:"projectKeys"`
	Policy      *types.Policy `protobuf:"bytes,5,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (m *ProjectData) Reset()         { *m = ProjectData{} }
func (m *ProjectData) String() string { return proto.CompactTextString(m) }
func (*ProjectData) ProtoMessage()    {}
func (*ProjectData) Descriptor() ([]byte, []int) {
	return fileDescriptor_9027839604ae2915, []int{3}
}
func (m *ProjectData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProjectData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProjectData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProjectData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectData.Merge(m, src)
}
func (m *ProjectData) XXX_Size() int {
	return m.Size()
}
func (m *ProjectData) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectData.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectData proto.InternalMessageInfo

func (m *ProjectData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProjectData) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *ProjectData) GetProjectKeys() []ProjectKey {
	if m != nil {
		return m.ProjectKeys
	}
	return nil
}

func (m *ProjectData) GetPolicy() *types.Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func init() {
	proto.RegisterEnum("lavanet.lava.projects.ProjectKey_Type", ProjectKey_Type_name, ProjectKey_Type_value)
	proto.RegisterType((*Project)(nil), "lavanet.lava.projects.Project")
	proto.RegisterType((*ProjectKey)(nil), "lavanet.lava.projects.ProjectKey")
	proto.RegisterType((*ProtoDeveloperData)(nil), "lavanet.lava.projects.ProtoDeveloperData")
	proto.RegisterType((*ProjectData)(nil), "lavanet.lava.projects.ProjectData")
}

func init() {
	proto.RegisterFile("lavanet/lava/projects/project.proto", fileDescriptor_9027839604ae2915)
}

var fileDescriptor_9027839604ae2915 = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4f, 0x8b, 0xda, 0x40,
	0x14, 0x77, 0x4c, 0xd4, 0xe4, 0xe9, 0x42, 0x98, 0x5a, 0x36, 0x95, 0x92, 0x58, 0x7b, 0x91, 0x16,
	0x12, 0xb0, 0x14, 0x7a, 0xad, 0xd5, 0x83, 0xdb, 0x56, 0x25, 0x94, 0x1e, 0xf6, 0x22, 0xd1, 0x0c,
	0x6e, 0xaa, 0x26, 0xc1, 0x89, 0xb2, 0xb9, 0xf5, 0x23, 0xf4, 0x63, 0x14, 0xfa, 0x25, 0x7a, 0xdc,
	0xe3, 0x1e, 0x7b, 0x92, 0xa2, 0x37, 0x3f, 0x45, 0x99, 0xcc, 0xb8, 0x9a, 0xb2, 0xb0, 0x7b, 0xc9,
	0xcc, 0x7b, 0xf3, 0x7b, 0xff, 0xf2, 0xfb, 0x3d, 0x78, 0x39, 0x77, 0xd7, 0x6e, 0x40, 0x62, 0x9b,
	0x9d, 0x76, 0xb4, 0x0c, 0xbf, 0x91, 0x49, 0x4c, 0x0f, 0x17, 0x2b, 0x5a, 0x86, 0x71, 0x88, 0x9f,
	0x0a, 0x90, 0xc5, 0x4e, 0xeb, 0x00, 0xaa, 0x55, 0xa7, 0xe1, 0x34, 0x4c, 0x11, 0x36, 0xbb, 0x71,
	0x70, 0xcd, 0xcc, 0x66, 0x9c, 0xbb, 0x01, 0xb5, 0xa3, 0x70, 0xee, 0x4f, 0x12, 0x0e, 0x68, 0xfc,
	0x92, 0xa0, 0x34, 0xe4, 0x39, 0x70, 0x15, 0x0a, 0x7e, 0xe0, 0x91, 0x6b, 0x1d, 0xd5, 0x51, 0x53,
	0x75, 0xb8, 0x81, 0x1b, 0x50, 0xa1, 0xab, 0x31, 0x9d, 0x2c, 0xfd, 0x28, 0xf6, 0xc3, 0x40, 0xcf,
	0xa7, 0x8f, 0x19, 0x1f, 0xd6, 0xa1, 0x44, 0x02, 0x77, 0x3c, 0x27, 0x9e, 0x2e, 0xd7, 0x51, 0x53,
	0x71, 0x0e, 0x26, 0xbe, 0x84, 0x8a, 0x68, 0x71, 0x34, 0x23, 0x09, 0xd5, 0x0b, 0x75, 0xa9, 0x59,
	0x6e, 0xbd, 0xb0, 0xee, 0x1d, 0xc2, 0x12, 0x9d, 0x7c, 0x24, 0x49, 0xbb, 0x7a, 0xb3, 0x31, 0x73,
	0xfb, 0x8d, 0x99, 0x09, 0x77, 0xca, 0xd1, 0x1d, 0x82, 0xe2, 0x01, 0x54, 0x5c, 0x6f, 0xe1, 0x07,
	0x23, 0x3e, 0x91, 0x5e, 0xac, 0xa3, 0x66, 0xb9, 0x55, 0xfb, 0x2f, 0x37, 0x9b, 0xd9, 0x1a, 0xa6,
	0x88, 0xb6, 0xc6, 0x12, 0x9e, 0xc6, 0x38, 0xe5, 0xd4, 0xe2, 0xcf, 0xf8, 0x1c, 0x4a, 0x2b, 0x4a,
	0xbc, 0xd1, 0x64, 0xa5, 0x97, 0xea, 0xa8, 0x29, 0x3b, 0x45, 0x66, 0x7e, 0x58, 0x61, 0x0f, 0x9e,
	0x9c, 0xce, 0x7b, 0x28, 0xa8, 0x3c, 0x58, 0xf0, 0x7c, 0xbf, 0x31, 0xef, 0x0b, 0x75, 0xf0, 0xa9,
	0x53, 0x94, 0xaf, 0x81, 0x42, 0x03, 0x37, 0xa2, 0x57, 0x61, 0xac, 0xab, 0x69, 0xfd, 0x3b, 0xfb,
	0x42, 0x56, 0x24, 0x4d, 0x6e, 0x7c, 0x47, 0x00, 0xc7, 0x7f, 0x84, 0x9f, 0x81, 0x34, 0x23, 0x09,
	0xa7, 0xab, 0x5d, 0xda, 0x6f, 0x4c, 0x66, 0x3a, 0xec, 0x83, 0x4d, 0x28, 0xcc, 0xfc, 0xc0, 0xa3,
	0x29, 0x1f, 0x67, 0x6d, 0x75, 0xbf, 0x31, 0xb9, 0xc3, 0xe1, 0x47, 0xe3, 0x15, 0xc8, 0x5f, 0x92,
	0x88, 0x60, 0x05, 0xe4, 0xfe, 0xa0, 0xdf, 0xd5, 0x72, 0x58, 0x85, 0xc2, 0xfb, 0xce, 0xe7, 0x5e,
	0x5f, 0x43, 0xf8, 0x0c, 0xd4, 0x4e, 0xf7, 0x6b, 0xf7, 0xd3, 0x60, 0xd8, 0x75, 0xb4, 0xfc, 0x85,
	0xac, 0xe4, 0x35, 0x49, 0xb4, 0xf0, 0x0e, 0xf0, 0x90, 0x29, 0xa7, 0x43, 0xd6, 0x64, 0x1e, 0x46,
	0x64, 0xd9, 0x71, 0x63, 0x17, 0x3f, 0x07, 0x55, 0x30, 0xd3, 0xeb, 0x08, 0xf9, 0x1c, 0x1d, 0x3c,
	0xbe, 0xf1, 0x1b, 0x41, 0x59, 0x34, 0x9f, 0xc6, 0x60, 0x90, 0x03, 0x77, 0x41, 0x04, 0x3c, 0xbd,
	0x9f, 0x0a, 0x49, 0xca, 0x0a, 0xa9, 0x07, 0xa7, 0xdc, 0xeb, 0xf2, 0x63, 0x75, 0x24, 0x33, 0x1d,
	0x65, 0x75, 0xd3, 0x82, 0xa2, 0x20, 0xb0, 0xf0, 0x10, 0x81, 0x8e, 0x40, 0xf2, 0x11, 0xda, 0xbd,
	0x9f, 0x5b, 0x03, 0xdd, 0x6c, 0x0d, 0x74, 0xbb, 0x35, 0xd0, 0xdf, 0xad, 0x81, 0x7e, 0xec, 0x8c,
	0xdc, 0xed, 0xce, 0xc8, 0xfd, 0xd9, 0x19, 0xb9, 0xcb, 0xd7, 0x53, 0x3f, 0xbe, 0x5a, 0x8d, 0xad,
	0x49, 0xb8, 0xb0, 0x33, 0x7b, 0xb7, 0x7e, 0x6b, 0x5f, 0x1f, 0xd7, 0x39, 0x4e, 0x22, 0x42, 0xc7,
	0xc5, 0x74, 0xff, 0xde, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xae, 0xd5, 0x2f, 0xf4, 0x03,
	0x00, 0x00,
}

func (this *Project) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Project)
	if !ok {
		that2, ok := that.(Project)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Index != that1.Index {
		return false
	}
	if this.Subscription != that1.Subscription {
		return false
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	if len(this.ProjectKeys) != len(that1.ProjectKeys) {
		return false
	}
	for i := range this.ProjectKeys {
		if !this.ProjectKeys[i].Equal(&that1.ProjectKeys[i]) {
			return false
		}
	}
	if !this.AdminPolicy.Equal(that1.AdminPolicy) {
		return false
	}
	if this.UsedCu != that1.UsedCu {
		return false
	}
	if !this.SubscriptionPolicy.Equal(that1.SubscriptionPolicy) {
		return false
	}
	if this.Snapshot != that1.Snapshot {
		return false
	}
	return true
}
func (this *ProjectKey) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProjectKey)
	if !ok {
		that2, ok := that.(ProjectKey)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if this.Kinds != that1.Kinds {
		return false
	}
	return true
}
func (this *ProtoDeveloperData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProtoDeveloperData)
	if !ok {
		that2, ok := that.(ProtoDeveloperData)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ProjectID != that1.ProjectID {
		return false
	}
	return true
}
func (this *ProjectData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProjectData)
	if !ok {
		that2, ok := that.(ProjectData)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Enabled != that1.Enabled {
		return false
	}
	if len(this.ProjectKeys) != len(that1.ProjectKeys) {
		return false
	}
	for i := range this.ProjectKeys {
		if !this.ProjectKeys[i].Equal(&that1.ProjectKeys[i]) {
			return false
		}
	}
	if !this.Policy.Equal(that1.Policy) {
		return false
	}
	return true
}
func (m *Project) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Project) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Project) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Snapshot != 0 {
		i = encodeVarintProject(dAtA, i, uint64(m.Snapshot))
		i--
		dAtA[i] = 0x48
	}
	if m.SubscriptionPolicy != nil {
		{
			size, err := m.SubscriptionPolicy.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.UsedCu != 0 {
		i = encodeVarintProject(dAtA, i, uint64(m.UsedCu))
		i--
		dAtA[i] = 0x38
	}
	if m.AdminPolicy != nil {
		{
			size, err := m.AdminPolicy.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.ProjectKeys) > 0 {
		for iNdEx := len(m.ProjectKeys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProjectKeys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProject(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Subscription) > 0 {
		i -= len(m.Subscription)
		copy(dAtA[i:], m.Subscription)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Subscription)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProjectKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProjectKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProjectKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Kinds != 0 {
		i = encodeVarintProject(dAtA, i, uint64(m.Kinds))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProtoDeveloperData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtoDeveloperData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtoDeveloperData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProjectID) > 0 {
		i -= len(m.ProjectID)
		copy(dAtA[i:], m.ProjectID)
		i = encodeVarintProject(dAtA, i, uint64(len(m.ProjectID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProjectData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProjectData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProjectData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Policy != nil {
		{
			size, err := m.Policy.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProject(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ProjectKeys) > 0 {
		for iNdEx := len(m.ProjectKeys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProjectKeys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProject(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProject(dAtA []byte, offset int, v uint64) int {
	offset -= sovProject(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Project) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	l = len(m.Subscription)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	if len(m.ProjectKeys) > 0 {
		for _, e := range m.ProjectKeys {
			l = e.Size()
			n += 1 + l + sovProject(uint64(l))
		}
	}
	if m.AdminPolicy != nil {
		l = m.AdminPolicy.Size()
		n += 1 + l + sovProject(uint64(l))
	}
	if m.UsedCu != 0 {
		n += 1 + sovProject(uint64(m.UsedCu))
	}
	if m.SubscriptionPolicy != nil {
		l = m.SubscriptionPolicy.Size()
		n += 1 + l + sovProject(uint64(l))
	}
	if m.Snapshot != 0 {
		n += 1 + sovProject(uint64(m.Snapshot))
	}
	return n
}

func (m *ProjectKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	if m.Kinds != 0 {
		n += 1 + sovProject(uint64(m.Kinds))
	}
	return n
}

func (m *ProtoDeveloperData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ProjectID)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	return n
}

func (m *ProjectData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	if len(m.ProjectKeys) > 0 {
		for _, e := range m.ProjectKeys {
			l = e.Size()
			n += 1 + l + sovProject(uint64(l))
		}
	}
	if m.Policy != nil {
		l = m.Policy.Size()
		n += 1 + l + sovProject(uint64(l))
	}
	return n
}

func sovProject(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProject(x uint64) (n int) {
	return sovProject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Project) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProject
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Project: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Project: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscription", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subscription = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectKeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProjectKeys = append(m.ProjectKeys, ProjectKey{})
			if err := m.ProjectKeys[len(m.ProjectKeys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminPolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AdminPolicy == nil {
				m.AdminPolicy = &types.Policy{}
			}
			if err := m.AdminPolicy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsedCu", wireType)
			}
			m.UsedCu = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UsedCu |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubscriptionPolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SubscriptionPolicy == nil {
				m.SubscriptionPolicy = &types.Policy{}
			}
			if err := m.SubscriptionPolicy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Snapshot", wireType)
			}
			m.Snapshot = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Snapshot |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProjectKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProject
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProjectKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProjectKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kinds", wireType)
			}
			m.Kinds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Kinds |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProtoDeveloperData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProject
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProtoDeveloperData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtoDeveloperData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProjectID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProjectData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProject
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProjectData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProjectData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectKeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProjectKeys = append(m.ProjectKeys, ProjectKey{})
			if err := m.ProjectKeys[len(m.ProjectKeys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Policy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Policy == nil {
				m.Policy = &types.Policy{}
			}
			if err := m.Policy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProject
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipProject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProject
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProject
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowProject
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthProject
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProject
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProject
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProject        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProject          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProject = fmt.Errorf("proto: unexpected end of group")
)
