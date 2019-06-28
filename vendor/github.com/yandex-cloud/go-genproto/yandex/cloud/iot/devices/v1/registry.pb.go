// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/iot/devices/v1/registry.proto

package devices

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"
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

type Registry struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FolderId             string               `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Name                 string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Labels               map[string]string    `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Registry) Reset()         { *m = Registry{} }
func (m *Registry) String() string { return proto.CompactTextString(m) }
func (*Registry) ProtoMessage()    {}
func (*Registry) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c05472a87f1ea4, []int{0}
}

func (m *Registry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Registry.Unmarshal(m, b)
}
func (m *Registry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Registry.Marshal(b, m, deterministic)
}
func (m *Registry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Registry.Merge(m, src)
}
func (m *Registry) XXX_Size() int {
	return xxx_messageInfo_Registry.Size(m)
}
func (m *Registry) XXX_DiscardUnknown() {
	xxx_messageInfo_Registry.DiscardUnknown(m)
}

var xxx_messageInfo_Registry proto.InternalMessageInfo

func (m *Registry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Registry) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *Registry) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Registry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Registry) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Registry) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type RegistryCertificate struct {
	RegistryId           string               `protobuf:"bytes,1,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
	Fingerprint          string               `protobuf:"bytes,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	CertificateData      string               `protobuf:"bytes,3,opt,name=certificate_data,json=certificateData,proto3" json:"certificate_data,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RegistryCertificate) Reset()         { *m = RegistryCertificate{} }
func (m *RegistryCertificate) String() string { return proto.CompactTextString(m) }
func (*RegistryCertificate) ProtoMessage()    {}
func (*RegistryCertificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c05472a87f1ea4, []int{1}
}

func (m *RegistryCertificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryCertificate.Unmarshal(m, b)
}
func (m *RegistryCertificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryCertificate.Marshal(b, m, deterministic)
}
func (m *RegistryCertificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryCertificate.Merge(m, src)
}
func (m *RegistryCertificate) XXX_Size() int {
	return xxx_messageInfo_RegistryCertificate.Size(m)
}
func (m *RegistryCertificate) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryCertificate.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryCertificate proto.InternalMessageInfo

func (m *RegistryCertificate) GetRegistryId() string {
	if m != nil {
		return m.RegistryId
	}
	return ""
}

func (m *RegistryCertificate) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *RegistryCertificate) GetCertificateData() string {
	if m != nil {
		return m.CertificateData
	}
	return ""
}

func (m *RegistryCertificate) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Registry)(nil), "yandex.cloud.iot.devices.v1.Registry")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.iot.devices.v1.Registry.LabelsEntry")
	proto.RegisterType((*RegistryCertificate)(nil), "yandex.cloud.iot.devices.v1.RegistryCertificate")
}

func init() {
	proto.RegisterFile("yandex/cloud/iot/devices/v1/registry.proto", fileDescriptor_39c05472a87f1ea4)
}

var fileDescriptor_39c05472a87f1ea4 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0xd2, 0x6e, 0xb5, 0x9d, 0x48, 0xb0, 0x32, 0x1c, 0xa2, 0xae, 0xd0, 0x56, 0x7b, 0x2a,
	0x48, 0x6b, 0xab, 0xcb, 0x85, 0x85, 0x13, 0x5f, 0x42, 0x95, 0x38, 0x45, 0x9c, 0xb8, 0x54, 0x6e,
	0x3c, 0x0d, 0x23, 0xd2, 0xb8, 0x72, 0xa6, 0x11, 0xfd, 0x53, 0xfc, 0x04, 0x7e, 0x1b, 0xaa, 0xed,
	0x40, 0xe0, 0xb0, 0xd2, 0xde, 0x26, 0xcf, 0x6f, 0x9e, 0xdf, 0x7b, 0x31, 0xbc, 0x38, 0xea, 0xc6,
	0xe0, 0x0f, 0x55, 0xd6, 0xf6, 0x60, 0x14, 0x59, 0x56, 0x06, 0x3b, 0x2a, 0xb1, 0x55, 0xdd, 0x52,
	0x39, 0xac, 0xa8, 0x65, 0x77, 0x94, 0x7b, 0x67, 0xd9, 0x8a, 0xcb, 0xc0, 0x95, 0x9e, 0x2b, 0xc9,
	0xb2, 0x8c, 0x5c, 0xd9, 0x2d, 0x67, 0x57, 0x95, 0xb5, 0x55, 0x8d, 0xca, 0x53, 0x37, 0x87, 0xad,
	0x62, 0xda, 0x61, 0xcb, 0x7a, 0xb7, 0x0f, 0xdb, 0xb3, 0x67, 0xff, 0xdc, 0xd4, 0xe9, 0x9a, 0x8c,
	0x66, 0xb2, 0x4d, 0x38, 0xbe, 0xfe, 0x99, 0xc2, 0x79, 0x11, 0xef, 0x13, 0x8f, 0x20, 0x25, 0x93,
	0x27, 0xf3, 0x64, 0x31, 0x2d, 0x52, 0x32, 0xe2, 0x12, 0xa6, 0x5b, 0x5b, 0x1b, 0x74, 0x6b, 0x32,
	0x79, 0xea, 0xe1, 0xf3, 0x00, 0xac, 0x8c, 0xb8, 0x03, 0x28, 0x1d, 0x6a, 0x46, 0xb3, 0xd6, 0x9c,
	0x8f, 0xe6, 0xc9, 0x22, 0xbb, 0x9d, 0xc9, 0x60, 0x47, 0xf6, 0x76, 0xe4, 0x97, 0xde, 0x4e, 0x31,
	0x8d, 0xec, 0xb7, 0x2c, 0x04, 0x8c, 0x1b, 0xbd, 0xc3, 0x7c, 0xec, 0x25, 0xfd, 0x2c, 0xe6, 0x90,
	0x19, 0x6c, 0x4b, 0x47, 0xfb, 0x93, 0xbb, 0xfc, 0xcc, 0x1f, 0x0d, 0x21, 0xb1, 0x82, 0x49, 0xad,
	0x37, 0x58, 0xb7, 0xf9, 0x64, 0x3e, 0x5a, 0x64, 0xb7, 0x4b, 0x79, 0x4f, 0x31, 0xb2, 0x0f, 0x25,
	0x3f, 0xfb, 0x9d, 0x8f, 0x0d, 0xbb, 0x63, 0x11, 0x05, 0x66, 0x77, 0x90, 0x0d, 0x60, 0x71, 0x01,
	0xa3, 0xef, 0x78, 0x8c, 0xc1, 0x4f, 0xa3, 0x78, 0x0a, 0x67, 0x9d, 0xae, 0x0f, 0x18, 0x53, 0x87,
	0x8f, 0xd7, 0xe9, 0xab, 0xe4, 0xfa, 0x57, 0x02, 0x4f, 0x7a, 0xed, 0xf7, 0xe8, 0x98, 0xb6, 0x54,
	0x6a, 0x46, 0x71, 0x05, 0x59, 0xff, 0xdf, 0xd6, 0x7f, 0x4a, 0x84, 0x1e, 0x5a, 0x99, 0x53, 0xc0,
	0x2d, 0x35, 0x15, 0xba, 0xbd, 0xa3, 0x86, 0xa3, 0xf0, 0x10, 0x12, 0xcf, 0xe1, 0xa2, 0xfc, 0xab,
	0xb8, 0x36, 0x9a, 0xb5, 0xef, 0x75, 0x5a, 0x3c, 0x1e, 0xe0, 0x1f, 0x34, 0xeb, 0xff, 0xca, 0x1f,
	0x3f, 0xa0, 0xfc, 0x77, 0xab, 0xaf, 0x9f, 0x2a, 0xe2, 0x6f, 0x87, 0x8d, 0x2c, 0xed, 0x4e, 0x85,
	0x0a, 0x6f, 0xc2, 0xeb, 0xa8, 0xec, 0x4d, 0x85, 0x8d, 0x5f, 0x57, 0xf7, 0x3c, 0xd0, 0x37, 0x71,
	0xdc, 0x4c, 0x3c, 0xf5, 0xe5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0xf1, 0x48, 0xbe, 0xce,
	0x02, 0x00, 0x00,
}