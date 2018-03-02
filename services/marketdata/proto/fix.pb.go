// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/ilovelili/sumoproto/services/marketdata/proto/fix.proto

/*
Package invast_sumo_srv_CurrenexMarketData is a generated protocol buffer package.

It is generated from these files:
	github.com/ilovelili/sumoproto/services/marketdata/proto/fix.proto

It has these top-level messages:
	Fix
*/
package invast_sumo_srv_CurrenexMarketData

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Fix struct {
	Fieldid      int32  `protobuf:"varint,1,opt,name=fieldid" json:"fieldid,omitempty"`
	Value        string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Type         string `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	Decodedvalue string `protobuf:"bytes,5,opt,name=decodedvalue" json:"decodedvalue,omitempty"`
	Time         int64  `protobuf:"varint,6,opt,name=time" json:"time,omitempty"`
}

func (m *Fix) Reset()                    { *m = Fix{} }
func (m *Fix) String() string            { return proto.CompactTextString(m) }
func (*Fix) ProtoMessage()               {}
func (*Fix) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Fix) GetFieldid() int32 {
	if m != nil {
		return m.Fieldid
	}
	return 0
}

func (m *Fix) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Fix) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Fix) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Fix) GetDecodedvalue() string {
	if m != nil {
		return m.Decodedvalue
	}
	return ""
}

func (m *Fix) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterType((*Fix)(nil), "invast.sumo.srv.CurrenexMarketData.Fix")
}

func init() {
	proto.RegisterFile("github.com/ilovelili/sumoproto/services/marketdata/proto/fix.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x31, 0x4e, 0xc4, 0x30,
	0x10, 0x45, 0x65, 0xb2, 0x59, 0x84, 0x45, 0x65, 0x51, 0xb8, 0x8c, 0x52, 0xa5, 0xb2, 0x0b, 0x6e,
	0x00, 0x88, 0x8e, 0x26, 0x37, 0x98, 0x8d, 0x67, 0x61, 0x84, 0x1d, 0x47, 0xb6, 0x63, 0x85, 0x83,
	0x70, 0x5f, 0x94, 0x89, 0x28, 0xb6, 0x7b, 0xff, 0x7d, 0xfd, 0x91, 0x46, 0xbe, 0x7c, 0x52, 0xf9,
	0x5a, 0x2f, 0x66, 0x8a, 0xc1, 0x92, 0x8f, 0x15, 0x3d, 0x79, 0xb2, 0x79, 0x0d, 0x71, 0x49, 0xb1,
	0x44, 0x9b, 0x31, 0x55, 0x9a, 0x30, 0xdb, 0x00, 0xe9, 0x1b, 0x8b, 0x83, 0x02, 0xf6, 0x68, 0xae,
	0xb4, 0x19, 0x26, 0xd5, 0xd3, 0x5c, 0x21, 0x17, 0xb3, 0xaf, 0x4c, 0x4e, 0xd5, 0xbc, 0xae, 0x29,
	0xe1, 0x8c, 0xdb, 0x07, 0x6f, 0xde, 0xa0, 0x40, 0xff, 0x2b, 0x64, 0xf3, 0x4e, 0x9b, 0xd2, 0xf2,
	0xfe, 0x4a, 0xe8, 0x1d, 0x39, 0x2d, 0x3a, 0x31, 0xb4, 0xe3, 0x7f, 0x54, 0x4f, 0xb2, 0xad, 0xe0,
	0x57, 0xd4, 0x77, 0x9d, 0x18, 0x1e, 0xc6, 0x23, 0x28, 0x25, 0x4f, 0x33, 0x04, 0xd4, 0x0d, 0x4b,
	0xe6, 0xdd, 0x95, 0x9f, 0x05, 0xf5, 0xe9, 0x70, 0x3b, 0xab, 0x5e, 0x3e, 0x3a, 0x9c, 0xa2, 0x43,
	0x77, 0x1c, 0x69, 0xb9, 0xbb, 0x71, 0xbc, 0xa3, 0x80, 0xfa, 0xdc, 0x89, 0xa1, 0x19, 0x99, 0x2f,
	0x67, 0x7e, 0xe1, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x77, 0x78, 0x08, 0x52, 0x08, 0x01, 0x00,
	0x00,
}