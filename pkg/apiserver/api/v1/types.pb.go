/*
Copyright 2021 The KubeCarrier Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package v1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type ObjectReference struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectReference) Reset()         { *m = ObjectReference{} }
func (m *ObjectReference) String() string { return proto.CompactTextString(m) }
func (*ObjectReference) ProtoMessage()    {}
func (*ObjectReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

func (m *ObjectReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectReference.Unmarshal(m, b)
}
func (m *ObjectReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectReference.Marshal(b, m, deterministic)
}
func (m *ObjectReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectReference.Merge(m, src)
}
func (m *ObjectReference) XXX_Size() int {
	return xxx_messageInfo_ObjectReference.Size(m)
}
func (m *ObjectReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectReference.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectReference proto.InternalMessageInfo

func (m *ObjectReference) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CRDInformation struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ApiGroup             string           `protobuf:"bytes,2,opt,name=apiGroup,proto3" json:"apiGroup,omitempty"`
	Kind                 string           `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Plural               string           `protobuf:"bytes,4,opt,name=plural,proto3" json:"plural,omitempty"`
	Versions             []*CRDVersion    `protobuf:"bytes,5,rep,name=versions,proto3" json:"versions,omitempty"`
	Region               *ObjectReference `protobuf:"bytes,6,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CRDInformation) Reset()         { *m = CRDInformation{} }
func (m *CRDInformation) String() string { return proto.CompactTextString(m) }
func (*CRDInformation) ProtoMessage()    {}
func (*CRDInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{1}
}

func (m *CRDInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CRDInformation.Unmarshal(m, b)
}
func (m *CRDInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CRDInformation.Marshal(b, m, deterministic)
}
func (m *CRDInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CRDInformation.Merge(m, src)
}
func (m *CRDInformation) XXX_Size() int {
	return xxx_messageInfo_CRDInformation.Size(m)
}
func (m *CRDInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_CRDInformation.DiscardUnknown(m)
}

var xxx_messageInfo_CRDInformation proto.InternalMessageInfo

func (m *CRDInformation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CRDInformation) GetApiGroup() string {
	if m != nil {
		return m.ApiGroup
	}
	return ""
}

func (m *CRDInformation) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *CRDInformation) GetPlural() string {
	if m != nil {
		return m.Plural
	}
	return ""
}

func (m *CRDInformation) GetVersions() []*CRDVersion {
	if m != nil {
		return m.Versions
	}
	return nil
}

func (m *CRDInformation) GetRegion() *ObjectReference {
	if m != nil {
		return m.Region
	}
	return nil
}

type CRDVersion struct {
	// name is the version name, e.g. “v1”, “v2beta1”, etc.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// schema describes the schema used for validation, pruning, and defaulting of this version of the custom resource.
	Schema string `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	// storage indicates this version should be used when persisting custom resources to storage.
	// There must be exactly one version with storage=true.
	Storage              bool     `protobuf:"varint,3,opt,name=storage,proto3" json:"storage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CRDVersion) Reset()         { *m = CRDVersion{} }
func (m *CRDVersion) String() string { return proto.CompactTextString(m) }
func (*CRDVersion) ProtoMessage()    {}
func (*CRDVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{2}
}

func (m *CRDVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CRDVersion.Unmarshal(m, b)
}
func (m *CRDVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CRDVersion.Marshal(b, m, deterministic)
}
func (m *CRDVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CRDVersion.Merge(m, src)
}
func (m *CRDVersion) XXX_Size() int {
	return xxx_messageInfo_CRDVersion.Size(m)
}
func (m *CRDVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_CRDVersion.DiscardUnknown(m)
}

var xxx_messageInfo_CRDVersion proto.InternalMessageInfo

func (m *CRDVersion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CRDVersion) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *CRDVersion) GetStorage() bool {
	if m != nil {
		return m.Storage
	}
	return false
}

type Image struct {
	MediaType            string   `protobuf:"bytes,1,opt,name=mediaType,proto3" json:"mediaType,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{3}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetMediaType() string {
	if m != nil {
		return m.MediaType
	}
	return ""
}

func (m *Image) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ConditionStatus struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConditionStatus) Reset()         { *m = ConditionStatus{} }
func (m *ConditionStatus) String() string { return proto.CompactTextString(m) }
func (*ConditionStatus) ProtoMessage()    {}
func (*ConditionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{4}
}

func (m *ConditionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConditionStatus.Unmarshal(m, b)
}
func (m *ConditionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConditionStatus.Marshal(b, m, deterministic)
}
func (m *ConditionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConditionStatus.Merge(m, src)
}
func (m *ConditionStatus) XXX_Size() int {
	return xxx_messageInfo_ConditionStatus.Size(m)
}
func (m *ConditionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ConditionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ConditionStatus proto.InternalMessageInfo

func (m *ConditionStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*ObjectReference)(nil), "kubecarrier.api.v1.ObjectReference")
	proto.RegisterType((*CRDInformation)(nil), "kubecarrier.api.v1.CRDInformation")
	proto.RegisterType((*CRDVersion)(nil), "kubecarrier.api.v1.CRDVersion")
	proto.RegisterType((*Image)(nil), "kubecarrier.api.v1.Image")
	proto.RegisterType((*ConditionStatus)(nil), "kubecarrier.api.v1.ConditionStatus")
}

func init() {
	proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355)
}

var fileDescriptor_d938547f84707355 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0xfd, 0x88, 0xed, 0x54, 0x2c, 0xec, 0xa1, 0x2c, 0x22, 0x52, 0x22, 0x42, 0xbd, 0x04,
	0x5a, 0x4f, 0xea, 0xcd, 0x16, 0xa4, 0x27, 0x61, 0x15, 0x0f, 0xde, 0xa6, 0xc9, 0xb4, 0xae, 0x6d,
	0x76, 0x97, 0xcd, 0xa6, 0xd0, 0x1f, 0xec, 0xff, 0x90, 0x6c, 0xd3, 0x16, 0x34, 0xb7, 0xf7, 0x1e,
	0x6f, 0xde, 0xec, 0xdb, 0x81, 0x9e, 0xdb, 0x19, 0xca, 0x63, 0x63, 0xb5, 0xd3, 0x8c, 0xad, 0x8b,
	0x05, 0x25, 0x68, 0xad, 0x24, 0x1b, 0xa3, 0x91, 0xf1, 0x76, 0x1c, 0xdd, 0x42, 0xff, 0x75, 0xf1,
	0x4d, 0x89, 0x13, 0xb4, 0x24, 0x4b, 0x2a, 0x21, 0xc6, 0xa0, 0xa5, 0x30, 0x23, 0x1e, 0x0c, 0x83,
	0x51, 0x57, 0x78, 0x1c, 0xfd, 0x04, 0x70, 0x31, 0x15, 0xb3, 0xb9, 0x5a, 0x6a, 0x9b, 0xa1, 0x93,
	0x5a, 0xd5, 0xd9, 0xd8, 0x25, 0x74, 0xd0, 0xc8, 0x17, 0xab, 0x0b, 0xc3, 0x1b, 0x5e, 0x3f, 0xf2,
	0xd2, 0xbf, 0x96, 0x2a, 0xe5, 0xcd, 0xbd, 0xbf, 0xc4, 0x6c, 0x00, 0xa1, 0xd9, 0x14, 0x16, 0x37,
	0xbc, 0xe5, 0xd5, 0x8a, 0xb1, 0x47, 0xe8, 0x6c, 0xc9, 0xe6, 0x52, 0xab, 0x9c, 0xb7, 0x87, 0xcd,
	0x51, 0x6f, 0x72, 0x1d, 0xff, 0x7f, 0x7c, 0x3c, 0x15, 0xb3, 0x8f, 0xbd, 0x4d, 0x1c, 0xfd, 0xec,
	0x09, 0x42, 0x4b, 0x2b, 0xa9, 0x15, 0x0f, 0x87, 0xc1, 0xa8, 0x37, 0xb9, 0xa9, 0x9b, 0xfc, 0xd3,
	0x59, 0x54, 0x23, 0x91, 0x00, 0x38, 0x85, 0xd6, 0x56, 0x1c, 0x40, 0x98, 0x27, 0x5f, 0x94, 0x61,
	0x55, 0xb0, 0x62, 0x8c, 0xc3, 0x59, 0xee, 0xb4, 0xc5, 0x15, 0xf9, 0x86, 0x1d, 0x71, 0xa0, 0xd1,
	0x03, 0xb4, 0xe7, 0x19, 0xae, 0x88, 0x5d, 0x41, 0x37, 0xa3, 0x54, 0xe2, 0xfb, 0xce, 0x1c, 0x32,
	0x4f, 0x42, 0xb9, 0x2c, 0x45, 0xb7, 0x8f, 0x3d, 0x17, 0x1e, 0x47, 0x77, 0xd0, 0x9f, 0x6a, 0x95,
	0xca, 0xf2, 0xc3, 0xdf, 0x1c, 0xba, 0x22, 0xf7, 0xfb, 0x3d, 0xaa, 0x12, 0x2a, 0xf6, 0xdc, 0xfa,
	0x6c, 0x6c, 0xc7, 0x8b, 0xd0, 0x5f, 0xfa, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xa0, 0xcc,
	0xcd, 0xf8, 0x01, 0x00, 0x00,
}
