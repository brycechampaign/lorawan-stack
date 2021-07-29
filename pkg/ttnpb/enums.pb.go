// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/enums.proto

package ttnpb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	math "math"
	strconv "strconv"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DownlinkPathConstraint int32

const (
	// Indicates that the gateway can be selected for downlink without constraints by the Network Server.
	DOWNLINK_PATH_CONSTRAINT_NONE DownlinkPathConstraint = 0
	// Indicates that the gateway can be selected for downlink only if no other or better gateway can be selected.
	DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER DownlinkPathConstraint = 1
	// Indicates that this gateway will never be selected for downlink, even if that results in no available downlink path.
	DOWNLINK_PATH_CONSTRAINT_NEVER DownlinkPathConstraint = 2
)

var DownlinkPathConstraint_name = map[int32]string{
	0: "DOWNLINK_PATH_CONSTRAINT_NONE",
	1: "DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER",
	2: "DOWNLINK_PATH_CONSTRAINT_NEVER",
}

var DownlinkPathConstraint_value = map[string]int32{
	"DOWNLINK_PATH_CONSTRAINT_NONE":         0,
	"DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER": 1,
	"DOWNLINK_PATH_CONSTRAINT_NEVER":        2,
}

func (DownlinkPathConstraint) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e36318a1e2f407cb, []int{0}
}

// State enum defines states that an entity can be in.
type State int32

const (
	// Denotes that the entity has been requested and is pending review by an admin.
	STATE_REQUESTED State = 0
	// Denotes that the entity has been reviewed and approved by an admin.
	STATE_APPROVED State = 1
	// Denotes that the entity has been reviewed and rejected by an admin.
	STATE_REJECTED State = 2
	// Denotes that the entity has been flagged and is pending review by an admin.
	STATE_FLAGGED State = 3
	// Denotes that the entity has been reviewed and suspended by an admin.
	STATE_SUSPENDED State = 4
)

var State_name = map[int32]string{
	0: "STATE_REQUESTED",
	1: "STATE_APPROVED",
	2: "STATE_REJECTED",
	3: "STATE_FLAGGED",
	4: "STATE_SUSPENDED",
}

var State_value = map[string]int32{
	"STATE_REQUESTED": 0,
	"STATE_APPROVED":  1,
	"STATE_REJECTED":  2,
	"STATE_FLAGGED":   3,
	"STATE_SUSPENDED": 4,
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e36318a1e2f407cb, []int{1}
}

type ClusterRole int32

const (
	ClusterRole_NONE                         ClusterRole = 0
	ClusterRole_ENTITY_REGISTRY              ClusterRole = 1
	ClusterRole_ACCESS                       ClusterRole = 2
	ClusterRole_GATEWAY_SERVER               ClusterRole = 3
	ClusterRole_NETWORK_SERVER               ClusterRole = 4
	ClusterRole_APPLICATION_SERVER           ClusterRole = 5
	ClusterRole_JOIN_SERVER                  ClusterRole = 6
	ClusterRole_CRYPTO_SERVER                ClusterRole = 7
	ClusterRole_DEVICE_TEMPLATE_CONVERTER    ClusterRole = 8
	ClusterRole_DEVICE_CLAIMING_SERVER       ClusterRole = 9
	ClusterRole_GATEWAY_CONFIGURATION_SERVER ClusterRole = 10
	ClusterRole_QR_CODE_GENERATOR            ClusterRole = 11
	ClusterRole_PACKET_BROKER_AGENT          ClusterRole = 12
	ClusterRole_DEVICE_REPOSITORY            ClusterRole = 13
)

var ClusterRole_name = map[int32]string{
	0:  "NONE",
	1:  "ENTITY_REGISTRY",
	2:  "ACCESS",
	3:  "GATEWAY_SERVER",
	4:  "NETWORK_SERVER",
	5:  "APPLICATION_SERVER",
	6:  "JOIN_SERVER",
	7:  "CRYPTO_SERVER",
	8:  "DEVICE_TEMPLATE_CONVERTER",
	9:  "DEVICE_CLAIMING_SERVER",
	10: "GATEWAY_CONFIGURATION_SERVER",
	11: "QR_CODE_GENERATOR",
	12: "PACKET_BROKER_AGENT",
	13: "DEVICE_REPOSITORY",
}

var ClusterRole_value = map[string]int32{
	"NONE":                         0,
	"ENTITY_REGISTRY":              1,
	"ACCESS":                       2,
	"GATEWAY_SERVER":               3,
	"NETWORK_SERVER":               4,
	"APPLICATION_SERVER":           5,
	"JOIN_SERVER":                  6,
	"CRYPTO_SERVER":                7,
	"DEVICE_TEMPLATE_CONVERTER":    8,
	"DEVICE_CLAIMING_SERVER":       9,
	"GATEWAY_CONFIGURATION_SERVER": 10,
	"QR_CODE_GENERATOR":            11,
	"PACKET_BROKER_AGENT":          12,
	"DEVICE_REPOSITORY":            13,
}

func (ClusterRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e36318a1e2f407cb, []int{2}
}

func init() {
	proto.RegisterEnum("ttn.lorawan.v3.DownlinkPathConstraint", DownlinkPathConstraint_name, DownlinkPathConstraint_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.DownlinkPathConstraint", DownlinkPathConstraint_name, DownlinkPathConstraint_value)
	proto.RegisterEnum("ttn.lorawan.v3.State", State_name, State_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.State", State_name, State_value)
	proto.RegisterEnum("ttn.lorawan.v3.ClusterRole", ClusterRole_name, ClusterRole_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.ClusterRole", ClusterRole_name, ClusterRole_value)
}

func init() { proto.RegisterFile("lorawan-stack/api/enums.proto", fileDescriptor_e36318a1e2f407cb) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/enums.proto", fileDescriptor_e36318a1e2f407cb)
}

var fileDescriptor_e36318a1e2f407cb = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x4e, 0xdb, 0x40,
	0x10, 0xc6, 0xe3, 0x10, 0x28, 0x5d, 0x0a, 0x2c, 0x8b, 0x4a, 0x55, 0x54, 0xac, 0xb6, 0x52, 0x0f,
	0x45, 0x22, 0x3e, 0xf0, 0x04, 0xcb, 0x7a, 0x30, 0x26, 0x61, 0xd7, 0x8c, 0x97, 0xa0, 0xf4, 0xb2,
	0x72, 0x50, 0x9a, 0x44, 0x09, 0x76, 0xe4, 0x6c, 0xe0, 0xda, 0x63, 0x8f, 0xbc, 0x43, 0x2f, 0x7d,
	0x0c, 0x1e, 0xa3, 0xd7, 0x86, 0x4b, 0x8f, 0x7d, 0x84, 0xca, 0xf9, 0x53, 0xc4, 0x81, 0xdb, 0xce,
	0x6f, 0xbf, 0x99, 0xf9, 0x3e, 0x69, 0xc8, 0xde, 0x20, 0xcb, 0x93, 0xdb, 0x24, 0x3d, 0x18, 0xd9,
	0xe4, 0xaa, 0xef, 0x25, 0xc3, 0x9e, 0xd7, 0x4e, 0xc7, 0xd7, 0xa3, 0xea, 0x30, 0xcf, 0x6c, 0xc6,
	0x36, 0xac, 0x4d, 0xab, 0x73, 0x49, 0xf5, 0xe6, 0x70, 0xf7, 0xa0, 0xd3, 0xb3, 0xdd, 0x71, 0xab,
	0x7a, 0x95, 0x5d, 0x7b, 0x9d, 0xac, 0x93, 0x79, 0x53, 0x59, 0x6b, 0xfc, 0x75, 0x5a, 0x4d, 0x8b,
	0xe9, 0x6b, 0xd6, 0xbe, 0x7f, 0xe7, 0x90, 0x1d, 0x3f, 0xbb, 0x4d, 0x07, 0xbd, 0xb4, 0x1f, 0x25,
	0xb6, 0x2b, 0xb2, 0x74, 0x64, 0xf3, 0xa4, 0x97, 0x5a, 0xf6, 0x81, 0xec, 0xf9, 0xea, 0x52, 0xd6,
	0x43, 0x59, 0x33, 0x11, 0xd7, 0x27, 0x46, 0x28, 0x19, 0x6b, 0xe4, 0xa1, 0xd4, 0x46, 0x2a, 0x09,
	0xb4, 0xc4, 0x3e, 0x93, 0x4f, 0xcf, 0x4a, 0x22, 0x84, 0x63, 0x40, 0xa3, 0xf4, 0x09, 0x20, 0x75,
	0xd8, 0x47, 0xe2, 0x3e, 0x3f, 0x0d, 0x1a, 0x80, 0xb4, 0xbc, 0x5b, 0xf9, 0xfe, 0xc3, 0x2d, 0xed,
	0xe7, 0x64, 0x39, 0xb6, 0x89, 0x6d, 0xb3, 0x6d, 0xb2, 0x19, 0x6b, 0xae, 0xc1, 0x20, 0x9c, 0x5f,
	0x40, 0xac, 0xc1, 0xa7, 0x25, 0xc6, 0xc8, 0xc6, 0x0c, 0xf2, 0x28, 0x42, 0xd5, 0x00, 0x9f, 0x3a,
	0x8f, 0x0c, 0xe1, 0x14, 0x44, 0xa1, 0x2b, 0xb3, 0x2d, 0xb2, 0x3e, 0x63, 0xc7, 0x75, 0x1e, 0x04,
	0xe0, 0xd3, 0xa5, 0xc7, 0x79, 0xf1, 0x45, 0x1c, 0x81, 0xf4, 0xc1, 0xa7, 0x95, 0xf9, 0xce, 0xfb,
	0x32, 0x59, 0x13, 0x83, 0xf1, 0xc8, 0xb6, 0x73, 0xcc, 0x06, 0x6d, 0xb6, 0x4a, 0x2a, 0xf3, 0x88,
	0xdb, 0x64, 0x13, 0xa4, 0x0e, 0x75, 0xd3, 0x20, 0x04, 0x61, 0xac, 0xb1, 0x49, 0x1d, 0x46, 0xc8,
	0x0a, 0x17, 0x02, 0xe2, 0x98, 0x96, 0x8b, 0xe5, 0x01, 0xd7, 0x70, 0xc9, 0x9b, 0x26, 0x06, 0x2c,
	0x82, 0x2c, 0x15, 0x4c, 0x82, 0xbe, 0x54, 0x58, 0x5b, 0xb0, 0x0a, 0xdb, 0x21, 0x8c, 0x47, 0x51,
	0x3d, 0x14, 0x5c, 0x87, 0x4a, 0x2e, 0xf8, 0x32, 0xdb, 0x24, 0x6b, 0xa7, 0x2a, 0xfc, 0x0f, 0x56,
	0x0a, 0xe7, 0x02, 0x9b, 0x91, 0x56, 0x0b, 0xf4, 0x82, 0xed, 0x91, 0xb7, 0x3e, 0x34, 0x42, 0x01,
	0x46, 0xc3, 0x59, 0x54, 0x2f, 0x32, 0x08, 0x25, 0x1b, 0x80, 0x1a, 0x90, 0xae, 0xb2, 0x5d, 0xb2,
	0x33, 0xff, 0x16, 0x75, 0x1e, 0x9e, 0x85, 0x32, 0x58, 0xb4, 0xbe, 0x64, 0xef, 0xc9, 0xbb, 0x85,
	0x3d, 0xa1, 0xe4, 0x71, 0x18, 0x5c, 0xe0, 0x13, 0x03, 0x84, 0xbd, 0x26, 0x5b, 0xe7, 0x68, 0x84,
	0xf2, 0xc1, 0x04, 0x20, 0x01, 0xb9, 0x56, 0x48, 0xd7, 0xd8, 0x1b, 0xb2, 0x1d, 0x71, 0x51, 0x03,
	0x6d, 0x8e, 0x50, 0xd5, 0x00, 0x0d, 0x0f, 0x40, 0x6a, 0xfa, 0xaa, 0xd0, 0xcf, 0xb7, 0x21, 0x44,
	0x2a, 0x0e, 0xb5, 0xc2, 0x26, 0x5d, 0x3f, 0x3a, 0xfb, 0xf5, 0xdb, 0x2d, 0x7d, 0x9b, 0xb8, 0xce,
	0xcf, 0x89, 0xeb, 0xfc, 0x99, 0xb8, 0xa5, 0xbf, 0x13, 0xd7, 0xb9, 0x7b, 0x70, 0x4b, 0xf7, 0x0f,
	0xae, 0xf3, 0xc5, 0xeb, 0x64, 0x55, 0xdb, 0x6d, 0xdb, 0x6e, 0x2f, 0xed, 0x8c, 0xaa, 0x69, 0xdb,
	0xde, 0x66, 0x79, 0xdf, 0x7b, 0x7a, 0xda, 0x37, 0x87, 0xde, 0xb0, 0xdf, 0xf1, 0xac, 0x4d, 0x87,
	0xad, 0xd6, 0xca, 0xf4, 0x3e, 0x0f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8d, 0xc7, 0x2a, 0x7e,
	0xff, 0x02, 0x00, 0x00,
}

func (x DownlinkPathConstraint) String() string {
	s, ok := DownlinkPathConstraint_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x State) String() string {
	s, ok := State_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ClusterRole) String() string {
	s, ok := ClusterRole_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
