// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type TestModeKeys int32

const (
	TestModeKeys_api_test TestModeKeys = 0
)

var TestModeKeys_name = map[int32]string{
	0: "api_test",
}

var TestModeKeys_value = map[string]int32{
	"api_test": 0,
}

func (x TestModeKeys) String() string {
	return proto.EnumName(TestModeKeys_name, int32(x))
}

func (TestModeKeys) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{0}
}

// Logging verbosity level
type LogLevel_LogLevel int32

const (
	LogLevel_DEBUG    LogLevel_LogLevel = 0
	LogLevel_INFO     LogLevel_LogLevel = 1
	LogLevel_WARNING  LogLevel_LogLevel = 2
	LogLevel_ERROR    LogLevel_LogLevel = 3
	LogLevel_CRITICAL LogLevel_LogLevel = 4
	LogLevel_FATAL    LogLevel_LogLevel = 5
)

var LogLevel_LogLevel_name = map[int32]string{
	0: "DEBUG",
	1: "INFO",
	2: "WARNING",
	3: "ERROR",
	4: "CRITICAL",
	5: "FATAL",
}

var LogLevel_LogLevel_value = map[string]int32{
	"DEBUG":    0,
	"INFO":     1,
	"WARNING":  2,
	"ERROR":    3,
	"CRITICAL": 4,
	"FATAL":    5,
}

func (x LogLevel_LogLevel) String() string {
	return proto.EnumName(LogLevel_LogLevel_name, int32(x))
}

func (LogLevel_LogLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{2, 0}
}

// Administrative State
type AdminState_AdminState int32

const (
	// The administrative state of the device is unknown
	AdminState_UNKNOWN AdminState_AdminState = 0
	// The device is pre-provisioned into Voltha, but not contacted by it
	AdminState_PREPROVISIONED AdminState_AdminState = 1
	// The device is enabled for activation and operation
	AdminState_ENABLED AdminState_AdminState = 2
	// The device is disabled and shall not perform its intended forwarding
	// functions other than being available for re-activation.
	AdminState_DISABLED AdminState_AdminState = 3
	// The device is in the state of image download
	AdminState_DOWNLOADING_IMAGE AdminState_AdminState = 4
	// The device is marked to be deleted
	AdminState_DELETED AdminState_AdminState = 5
)

var AdminState_AdminState_name = map[int32]string{
	0: "UNKNOWN",
	1: "PREPROVISIONED",
	2: "ENABLED",
	3: "DISABLED",
	4: "DOWNLOADING_IMAGE",
	5: "DELETED",
}

var AdminState_AdminState_value = map[string]int32{
	"UNKNOWN":           0,
	"PREPROVISIONED":    1,
	"ENABLED":           2,
	"DISABLED":          3,
	"DOWNLOADING_IMAGE": 4,
	"DELETED":           5,
}

func (x AdminState_AdminState) String() string {
	return proto.EnumName(AdminState_AdminState_name, int32(x))
}

func (AdminState_AdminState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{6, 0}
}

// Operational Status
type OperStatus_OperStatus int32

const (
	// The status of the device is unknown at this point
	OperStatus_UNKNOWN OperStatus_OperStatus = 0
	// The device has been discovered, but not yet activated
	OperStatus_DISCOVERED OperStatus_OperStatus = 1
	// The device is being activated (booted, rebooted, upgraded, etc.)
	OperStatus_ACTIVATING OperStatus_OperStatus = 2
	// Service impacting tests are being conducted
	OperStatus_TESTING OperStatus_OperStatus = 3
	// The device is up and active
	OperStatus_ACTIVE OperStatus_OperStatus = 4
	// The device has failed and cannot fulfill its intended role
	OperStatus_FAILED OperStatus_OperStatus = 5
)

var OperStatus_OperStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "DISCOVERED",
	2: "ACTIVATING",
	3: "TESTING",
	4: "ACTIVE",
	5: "FAILED",
}

var OperStatus_OperStatus_value = map[string]int32{
	"UNKNOWN":    0,
	"DISCOVERED": 1,
	"ACTIVATING": 2,
	"TESTING":    3,
	"ACTIVE":     4,
	"FAILED":     5,
}

func (x OperStatus_OperStatus) String() string {
	return proto.EnumName(OperStatus_OperStatus_name, int32(x))
}

func (OperStatus_OperStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{7, 0}
}

// Connectivity Status
type ConnectStatus_ConnectStatus int32

const (
	// The device connectivity status is unknown
	ConnectStatus_UNKNOWN ConnectStatus_ConnectStatus = 0
	// The device cannot be reached by Voltha
	ConnectStatus_UNREACHABLE ConnectStatus_ConnectStatus = 1
	// There is live communication between device and Voltha
	ConnectStatus_REACHABLE ConnectStatus_ConnectStatus = 2
)

var ConnectStatus_ConnectStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "UNREACHABLE",
	2: "REACHABLE",
}

var ConnectStatus_ConnectStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"UNREACHABLE": 1,
	"REACHABLE":   2,
}

func (x ConnectStatus_ConnectStatus) String() string {
	return proto.EnumName(ConnectStatus_ConnectStatus_name, int32(x))
}

func (ConnectStatus_ConnectStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{8, 0}
}

type OperationResp_OperationReturnCode int32

const (
	OperationResp_OPERATION_SUCCESS     OperationResp_OperationReturnCode = 0
	OperationResp_OPERATION_FAILURE     OperationResp_OperationReturnCode = 1
	OperationResp_OPERATION_UNSUPPORTED OperationResp_OperationReturnCode = 2
)

var OperationResp_OperationReturnCode_name = map[int32]string{
	0: "OPERATION_SUCCESS",
	1: "OPERATION_FAILURE",
	2: "OPERATION_UNSUPPORTED",
}

var OperationResp_OperationReturnCode_value = map[string]int32{
	"OPERATION_SUCCESS":     0,
	"OPERATION_FAILURE":     1,
	"OPERATION_UNSUPPORTED": 2,
}

func (x OperationResp_OperationReturnCode) String() string {
	return proto.EnumName(OperationResp_OperationReturnCode_name, int32(x))
}

func (OperationResp_OperationReturnCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{9, 0}
}

// Convey a resource identifier
type ID struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{0}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Represents a list of IDs
type IDs struct {
	Items                []*ID    `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDs) Reset()         { *m = IDs{} }
func (m *IDs) String() string { return proto.CompactTextString(m) }
func (*IDs) ProtoMessage()    {}
func (*IDs) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{1}
}

func (m *IDs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDs.Unmarshal(m, b)
}
func (m *IDs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDs.Marshal(b, m, deterministic)
}
func (m *IDs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDs.Merge(m, src)
}
func (m *IDs) XXX_Size() int {
	return xxx_messageInfo_IDs.Size(m)
}
func (m *IDs) XXX_DiscardUnknown() {
	xxx_messageInfo_IDs.DiscardUnknown(m)
}

var xxx_messageInfo_IDs proto.InternalMessageInfo

func (m *IDs) GetItems() []*ID {
	if m != nil {
		return m.Items
	}
	return nil
}

type LogLevel struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogLevel) Reset()         { *m = LogLevel{} }
func (m *LogLevel) String() string { return proto.CompactTextString(m) }
func (*LogLevel) ProtoMessage()    {}
func (*LogLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{2}
}

func (m *LogLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogLevel.Unmarshal(m, b)
}
func (m *LogLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogLevel.Marshal(b, m, deterministic)
}
func (m *LogLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogLevel.Merge(m, src)
}
func (m *LogLevel) XXX_Size() int {
	return xxx_messageInfo_LogLevel.Size(m)
}
func (m *LogLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_LogLevel.DiscardUnknown(m)
}

var xxx_messageInfo_LogLevel proto.InternalMessageInfo

type Logging struct {
	Level                LogLevel_LogLevel `protobuf:"varint,1,opt,name=level,proto3,enum=common.LogLevel_LogLevel" json:"level,omitempty"`
	PackageName          string            `protobuf:"bytes,2,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	ComponentName        string            `protobuf:"bytes,3,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Logging) Reset()         { *m = Logging{} }
func (m *Logging) String() string { return proto.CompactTextString(m) }
func (*Logging) ProtoMessage()    {}
func (*Logging) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{3}
}

func (m *Logging) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Logging.Unmarshal(m, b)
}
func (m *Logging) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Logging.Marshal(b, m, deterministic)
}
func (m *Logging) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Logging.Merge(m, src)
}
func (m *Logging) XXX_Size() int {
	return xxx_messageInfo_Logging.Size(m)
}
func (m *Logging) XXX_DiscardUnknown() {
	xxx_messageInfo_Logging.DiscardUnknown(m)
}

var xxx_messageInfo_Logging proto.InternalMessageInfo

func (m *Logging) GetLevel() LogLevel_LogLevel {
	if m != nil {
		return m.Level
	}
	return LogLevel_DEBUG
}

func (m *Logging) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *Logging) GetComponentName() string {
	if m != nil {
		return m.ComponentName
	}
	return ""
}

// For GetLogLevels(), select component to query
type LoggingComponent struct {
	ComponentName        string   `protobuf:"bytes,1,opt,name=component_name,json=componentName,proto3" json:"component_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoggingComponent) Reset()         { *m = LoggingComponent{} }
func (m *LoggingComponent) String() string { return proto.CompactTextString(m) }
func (*LoggingComponent) ProtoMessage()    {}
func (*LoggingComponent) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{4}
}

func (m *LoggingComponent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoggingComponent.Unmarshal(m, b)
}
func (m *LoggingComponent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoggingComponent.Marshal(b, m, deterministic)
}
func (m *LoggingComponent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoggingComponent.Merge(m, src)
}
func (m *LoggingComponent) XXX_Size() int {
	return xxx_messageInfo_LoggingComponent.Size(m)
}
func (m *LoggingComponent) XXX_DiscardUnknown() {
	xxx_messageInfo_LoggingComponent.DiscardUnknown(m)
}

var xxx_messageInfo_LoggingComponent proto.InternalMessageInfo

func (m *LoggingComponent) GetComponentName() string {
	if m != nil {
		return m.ComponentName
	}
	return ""
}

// For returning multiple log levels
type Loggings struct {
	Items                []*Logging `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Loggings) Reset()         { *m = Loggings{} }
func (m *Loggings) String() string { return proto.CompactTextString(m) }
func (*Loggings) ProtoMessage()    {}
func (*Loggings) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{5}
}

func (m *Loggings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Loggings.Unmarshal(m, b)
}
func (m *Loggings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Loggings.Marshal(b, m, deterministic)
}
func (m *Loggings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Loggings.Merge(m, src)
}
func (m *Loggings) XXX_Size() int {
	return xxx_messageInfo_Loggings.Size(m)
}
func (m *Loggings) XXX_DiscardUnknown() {
	xxx_messageInfo_Loggings.DiscardUnknown(m)
}

var xxx_messageInfo_Loggings proto.InternalMessageInfo

func (m *Loggings) GetItems() []*Logging {
	if m != nil {
		return m.Items
	}
	return nil
}

type AdminState struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdminState) Reset()         { *m = AdminState{} }
func (m *AdminState) String() string { return proto.CompactTextString(m) }
func (*AdminState) ProtoMessage()    {}
func (*AdminState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{6}
}

func (m *AdminState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminState.Unmarshal(m, b)
}
func (m *AdminState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminState.Marshal(b, m, deterministic)
}
func (m *AdminState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminState.Merge(m, src)
}
func (m *AdminState) XXX_Size() int {
	return xxx_messageInfo_AdminState.Size(m)
}
func (m *AdminState) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminState.DiscardUnknown(m)
}

var xxx_messageInfo_AdminState proto.InternalMessageInfo

type OperStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperStatus) Reset()         { *m = OperStatus{} }
func (m *OperStatus) String() string { return proto.CompactTextString(m) }
func (*OperStatus) ProtoMessage()    {}
func (*OperStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{7}
}

func (m *OperStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperStatus.Unmarshal(m, b)
}
func (m *OperStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperStatus.Marshal(b, m, deterministic)
}
func (m *OperStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperStatus.Merge(m, src)
}
func (m *OperStatus) XXX_Size() int {
	return xxx_messageInfo_OperStatus.Size(m)
}
func (m *OperStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_OperStatus.DiscardUnknown(m)
}

var xxx_messageInfo_OperStatus proto.InternalMessageInfo

type ConnectStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectStatus) Reset()         { *m = ConnectStatus{} }
func (m *ConnectStatus) String() string { return proto.CompactTextString(m) }
func (*ConnectStatus) ProtoMessage()    {}
func (*ConnectStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{8}
}

func (m *ConnectStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectStatus.Unmarshal(m, b)
}
func (m *ConnectStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectStatus.Marshal(b, m, deterministic)
}
func (m *ConnectStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectStatus.Merge(m, src)
}
func (m *ConnectStatus) XXX_Size() int {
	return xxx_messageInfo_ConnectStatus.Size(m)
}
func (m *ConnectStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectStatus proto.InternalMessageInfo

type OperationResp struct {
	// Return code
	Code OperationResp_OperationReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=common.OperationResp_OperationReturnCode" json:"code,omitempty"`
	// Additional Info
	AdditionalInfo       string   `protobuf:"bytes,2,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationResp) Reset()         { *m = OperationResp{} }
func (m *OperationResp) String() string { return proto.CompactTextString(m) }
func (*OperationResp) ProtoMessage()    {}
func (*OperationResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e3fd231961e826, []int{9}
}

func (m *OperationResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationResp.Unmarshal(m, b)
}
func (m *OperationResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationResp.Marshal(b, m, deterministic)
}
func (m *OperationResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationResp.Merge(m, src)
}
func (m *OperationResp) XXX_Size() int {
	return xxx_messageInfo_OperationResp.Size(m)
}
func (m *OperationResp) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationResp.DiscardUnknown(m)
}

var xxx_messageInfo_OperationResp proto.InternalMessageInfo

func (m *OperationResp) GetCode() OperationResp_OperationReturnCode {
	if m != nil {
		return m.Code
	}
	return OperationResp_OPERATION_SUCCESS
}

func (m *OperationResp) GetAdditionalInfo() string {
	if m != nil {
		return m.AdditionalInfo
	}
	return ""
}

func init() {
	proto.RegisterEnum("common.TestModeKeys", TestModeKeys_name, TestModeKeys_value)
	proto.RegisterEnum("common.LogLevel_LogLevel", LogLevel_LogLevel_name, LogLevel_LogLevel_value)
	proto.RegisterEnum("common.AdminState_AdminState", AdminState_AdminState_name, AdminState_AdminState_value)
	proto.RegisterEnum("common.OperStatus_OperStatus", OperStatus_OperStatus_name, OperStatus_OperStatus_value)
	proto.RegisterEnum("common.ConnectStatus_ConnectStatus", ConnectStatus_ConnectStatus_name, ConnectStatus_ConnectStatus_value)
	proto.RegisterEnum("common.OperationResp_OperationReturnCode", OperationResp_OperationReturnCode_name, OperationResp_OperationReturnCode_value)
	proto.RegisterType((*ID)(nil), "common.ID")
	proto.RegisterType((*IDs)(nil), "common.IDs")
	proto.RegisterType((*LogLevel)(nil), "common.LogLevel")
	proto.RegisterType((*Logging)(nil), "common.Logging")
	proto.RegisterType((*LoggingComponent)(nil), "common.LoggingComponent")
	proto.RegisterType((*Loggings)(nil), "common.Loggings")
	proto.RegisterType((*AdminState)(nil), "common.AdminState")
	proto.RegisterType((*OperStatus)(nil), "common.OperStatus")
	proto.RegisterType((*ConnectStatus)(nil), "common.ConnectStatus")
	proto.RegisterType((*OperationResp)(nil), "common.OperationResp")
}

func init() { proto.RegisterFile("voltha_protos/common.proto", fileDescriptor_c2e3fd231961e826) }

var fileDescriptor_c2e3fd231961e826 = []byte{
	// 661 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0x4d, 0x4f, 0xdb, 0x4a,
	0x14, 0x8d, 0xf3, 0x05, 0xdc, 0x90, 0xe0, 0x37, 0xef, 0x21, 0x01, 0x7a, 0x95, 0x52, 0x4b, 0x08,
	0xda, 0x0a, 0xa2, 0xd2, 0x55, 0xab, 0x76, 0x61, 0xec, 0x21, 0x1d, 0x61, 0xc6, 0xd1, 0xd8, 0x01,
	0xa9, 0x0b, 0x22, 0x13, 0x0f, 0xc6, 0x6a, 0x32, 0x63, 0xc5, 0x06, 0x89, 0x65, 0xa5, 0xfe, 0xbd,
	0xfe, 0x85, 0xfe, 0x86, 0xae, 0xba, 0xae, 0xc6, 0x76, 0x48, 0x52, 0xb1, 0xf3, 0x39, 0x73, 0xae,
	0xcf, 0xbd, 0xe7, 0x8e, 0x06, 0xf6, 0x1e, 0xe4, 0x24, 0xbb, 0x0b, 0x46, 0xc9, 0x4c, 0x66, 0x32,
	0xed, 0x8d, 0xe5, 0x74, 0x2a, 0xc5, 0x71, 0x8e, 0x50, 0xb3, 0x40, 0x7b, 0xdd, 0x55, 0xcd, 0x63,
	0x20, 0xa2, 0x91, 0x4c, 0xb2, 0x58, 0x8a, 0xb4, 0x50, 0x1a, 0xff, 0x41, 0x95, 0xd8, 0xa8, 0x03,
	0xd5, 0x38, 0xdc, 0xd1, 0xba, 0xda, 0xe1, 0x06, 0xab, 0xc6, 0xa1, 0x71, 0x00, 0x35, 0x62, 0xa7,
	0xa8, 0x0b, 0x8d, 0x38, 0xe3, 0xd3, 0x74, 0x47, 0xeb, 0xd6, 0x0e, 0x5b, 0x27, 0x70, 0x5c, 0x9a,
	0x10, 0x9b, 0x15, 0x07, 0xc6, 0x18, 0xd6, 0x1d, 0x19, 0x39, 0xfc, 0x81, 0x4f, 0x8c, 0xc1, 0xe2,
	0x1b, 0x6d, 0x40, 0xc3, 0xc6, 0xa7, 0xc3, 0xbe, 0x5e, 0x41, 0xeb, 0x50, 0x27, 0xf4, 0xcc, 0xd5,
	0x35, 0xd4, 0x82, 0xb5, 0x2b, 0x93, 0x51, 0x42, 0xfb, 0x7a, 0x55, 0x29, 0x30, 0x63, 0x2e, 0xd3,
	0x6b, 0x68, 0x13, 0xd6, 0x2d, 0x46, 0x7c, 0x62, 0x99, 0x8e, 0x5e, 0x57, 0x07, 0x67, 0xa6, 0x6f,
	0x3a, 0x7a, 0xe3, 0x43, 0xe3, 0xd7, 0xef, 0x1f, 0x2f, 0x2a, 0xc6, 0x77, 0x0d, 0xd6, 0x1c, 0x19,
	0x45, 0xb1, 0x88, 0x50, 0x0f, 0x1a, 0x13, 0xe5, 0x90, 0x37, 0xdb, 0x39, 0xd9, 0x9d, 0xb7, 0x34,
	0x77, 0x7e, 0xfa, 0x60, 0x85, 0x0e, 0xbd, 0x84, 0xcd, 0x24, 0x18, 0x7f, 0x0d, 0x22, 0x3e, 0x12,
	0xc1, 0x94, 0xef, 0x54, 0xf3, 0x21, 0x5b, 0x25, 0x47, 0x83, 0x29, 0x47, 0xfb, 0xd0, 0x19, 0xcb,
	0x69, 0x22, 0x05, 0x17, 0x59, 0x21, 0xaa, 0xe5, 0xa2, 0xf6, 0x13, 0xab, 0x64, 0xc6, 0x7b, 0xd0,
	0xcb, 0x2e, 0xac, 0x39, 0xff, 0x4c, 0xa9, 0xf6, 0x5c, 0xe9, 0xdb, 0x3c, 0x1a, 0x55, 0x9a, 0xa2,
	0xfd, 0xd5, 0x50, 0xb7, 0x96, 0x26, 0x50, 0x82, 0x79, 0xb2, 0xdf, 0x34, 0x00, 0x33, 0x9c, 0xc6,
	0xc2, 0xcb, 0x82, 0x8c, 0x1b, 0x93, 0x65, 0xa4, 0x92, 0x1c, 0xd2, 0x73, 0xea, 0x5e, 0x51, 0xbd,
	0x82, 0x10, 0x74, 0x06, 0x0c, 0x0f, 0x98, 0x7b, 0x49, 0x3c, 0xe2, 0x52, 0x6c, 0x17, 0x51, 0x63,
	0x6a, 0x9e, 0x3a, 0xd8, 0xd6, 0xab, 0x2a, 0x5f, 0x9b, 0x78, 0x05, 0xaa, 0xa1, 0x6d, 0xf8, 0xc7,
	0x76, 0xaf, 0xa8, 0xe3, 0x9a, 0x36, 0xa1, 0xfd, 0x11, 0xb9, 0x30, 0xfb, 0x58, 0xaf, 0xab, 0x0a,
	0x1b, 0x3b, 0xd8, 0xc7, 0xf6, 0x22, 0xf8, 0x14, 0xc0, 0x4d, 0xf8, 0x4c, 0x79, 0xde, 0xa7, 0xc6,
	0xf5, 0x32, 0x5a, 0x6d, 0xa1, 0x03, 0x60, 0x13, 0xcf, 0x72, 0x2f, 0x31, 0xcb, 0xed, 0x3b, 0x00,
	0xa6, 0xe5, 0x93, 0x4b, 0xd3, 0x2f, 0x96, 0xdd, 0x82, 0x35, 0x1f, 0x7b, 0x39, 0xa8, 0x21, 0x80,
	0x66, 0x7e, 0xa8, 0x5c, 0x01, 0x9a, 0x67, 0x26, 0x71, 0x96, 0x4d, 0x7d, 0x68, 0x5b, 0x52, 0x08,
	0x3e, 0xce, 0x4a, 0xdf, 0x8f, 0x7f, 0x11, 0xab, 0xd6, 0x5b, 0xd0, 0x1a, 0x52, 0x86, 0x4d, 0xeb,
	0xb3, 0x1a, 0x50, 0xd7, 0x50, 0x1b, 0x36, 0x16, 0xb0, 0x3a, 0xff, 0xeb, 0x4f, 0x0d, 0xda, 0xaa,
	0xfb, 0x40, 0x5d, 0x7e, 0xc6, 0xd3, 0x04, 0x7d, 0x82, 0xfa, 0x58, 0x86, 0xbc, 0xbc, 0x48, 0xaf,
	0xe6, 0x6b, 0x58, 0x11, 0x2d, 0xa3, 0xec, 0x7e, 0x26, 0x2c, 0x19, 0x72, 0x96, 0x97, 0xa1, 0x03,
	0xd8, 0x0a, 0xc2, 0x30, 0x56, 0x67, 0xc1, 0x64, 0x14, 0x8b, 0x5b, 0x59, 0x5e, 0xad, 0xce, 0x82,
	0x26, 0xe2, 0x56, 0x1a, 0xd7, 0xf0, 0xef, 0x33, 0x7f, 0x51, 0x6b, 0x70, 0x07, 0x98, 0x99, 0x3e,
	0x71, 0xe9, 0xc8, 0x1b, 0x5a, 0x16, 0xf6, 0x3c, 0xbd, 0xb2, 0x4a, 0xab, 0x68, 0x86, 0x4c, 0x0d,
	0xb5, 0x0b, 0xdb, 0x0b, 0x7a, 0x48, 0xbd, 0xe1, 0x60, 0xe0, 0x32, 0xb5, 0xab, 0xf9, 0x80, 0xaf,
	0xff, 0x87, 0x4d, 0x9f, 0xa7, 0xd9, 0x85, 0x0c, 0xf9, 0x39, 0x7f, 0x4c, 0xd5, 0xd2, 0x83, 0x24,
	0x1e, 0x65, 0x3c, 0xcd, 0xf4, 0xca, 0xe9, 0xd1, 0x97, 0x37, 0x51, 0x9c, 0xdd, 0xdd, 0xdf, 0xa8,
	0x31, 0x7b, 0x32, 0xe1, 0x62, 0x2c, 0x67, 0x61, 0xaf, 0x78, 0x1e, 0x8e, 0xca, 0xe7, 0x21, 0x92,
	0xe5, 0x2b, 0x72, 0xd3, 0xcc, 0x99, 0x77, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xd4, 0xbf,
	0xf3, 0x64, 0x04, 0x00, 0x00,
}