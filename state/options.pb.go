// Code generated by protoc-gen-go. DO NOT EDIT.
// source: state/options.proto

package state // import "github.com/tcncloud/protoc-gen-state/state"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StateMessageType int32

const (
	StateMessageType_REDUX_STATE   StateMessageType = 0
	StateMessageType_CUSTOM_ACTION StateMessageType = 1
	StateMessageType_EXTERNAL_LINK StateMessageType = 2
)

var StateMessageType_name = map[int32]string{
	0: "REDUX_STATE",
	1: "CUSTOM_ACTION",
	2: "EXTERNAL_LINK",
}
var StateMessageType_value = map[string]int32{
	"REDUX_STATE":   0,
	"CUSTOM_ACTION": 1,
	"EXTERNAL_LINK": 2,
}

func (x StateMessageType) Enum() *StateMessageType {
	p := new(StateMessageType)
	*p = x
	return p
}
func (x StateMessageType) String() string {
	return proto.EnumName(StateMessageType_name, int32(x))
}
func (x *StateMessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(StateMessageType_value, data, "StateMessageType")
	if err != nil {
		return err
	}
	*x = StateMessageType(value)
	return nil
}
func (StateMessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_options_64966814c91bee01, []int{0}
}

type OutputTypes int32

const (
	OutputTypes_redux3 OutputTypes = 0
	OutputTypes_redux4 OutputTypes = 1
	OutputTypes_mobx   OutputTypes = 2
)

var OutputTypes_name = map[int32]string{
	0: "redux3",
	1: "redux4",
	2: "mobx",
}
var OutputTypes_value = map[string]int32{
	"redux3": 0,
	"redux4": 1,
	"mobx":   2,
}

func (x OutputTypes) Enum() *OutputTypes {
	p := new(OutputTypes)
	*p = x
	return p
}
func (x OutputTypes) String() string {
	return proto.EnumName(OutputTypes_name, int32(x))
}
func (x *OutputTypes) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(OutputTypes_value, data, "OutputTypes")
	if err != nil {
		return err
	}
	*x = OutputTypes(value)
	return nil
}
func (OutputTypes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_options_64966814c91bee01, []int{1}
}

type StringFieldOptions struct {
	Create               *string  `protobuf:"bytes,1,opt,name=create" json:"create,omitempty"`
	Get                  *string  `protobuf:"bytes,2,opt,name=get" json:"get,omitempty"`
	List                 *string  `protobuf:"bytes,3,opt,name=list" json:"list,omitempty"`
	Update               *string  `protobuf:"bytes,4,opt,name=update" json:"update,omitempty"`
	Delete               *string  `protobuf:"bytes,5,opt,name=delete" json:"delete,omitempty"`
	Custom               *string  `protobuf:"bytes,6,opt,name=custom" json:"custom,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringFieldOptions) Reset()         { *m = StringFieldOptions{} }
func (m *StringFieldOptions) String() string { return proto.CompactTextString(m) }
func (*StringFieldOptions) ProtoMessage()    {}
func (*StringFieldOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_options_64966814c91bee01, []int{0}
}
func (m *StringFieldOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringFieldOptions.Unmarshal(m, b)
}
func (m *StringFieldOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringFieldOptions.Marshal(b, m, deterministic)
}
func (dst *StringFieldOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringFieldOptions.Merge(dst, src)
}
func (m *StringFieldOptions) XXX_Size() int {
	return xxx_messageInfo_StringFieldOptions.Size(m)
}
func (m *StringFieldOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_StringFieldOptions.DiscardUnknown(m)
}

var xxx_messageInfo_StringFieldOptions proto.InternalMessageInfo

func (m *StringFieldOptions) GetCreate() string {
	if m != nil && m.Create != nil {
		return *m.Create
	}
	return ""
}

func (m *StringFieldOptions) GetGet() string {
	if m != nil && m.Get != nil {
		return *m.Get
	}
	return ""
}

func (m *StringFieldOptions) GetList() string {
	if m != nil && m.List != nil {
		return *m.List
	}
	return ""
}

func (m *StringFieldOptions) GetUpdate() string {
	if m != nil && m.Update != nil {
		return *m.Update
	}
	return ""
}

func (m *StringFieldOptions) GetDelete() string {
	if m != nil && m.Delete != nil {
		return *m.Delete
	}
	return ""
}

func (m *StringFieldOptions) GetCustom() string {
	if m != nil && m.Custom != nil {
		return *m.Custom
	}
	return ""
}

type IntFieldOptions struct {
	Create               *int64   `protobuf:"varint,1,opt,name=create" json:"create,omitempty"`
	Get                  *int64   `protobuf:"varint,2,opt,name=get" json:"get,omitempty"`
	List                 *int64   `protobuf:"varint,3,opt,name=list" json:"list,omitempty"`
	Update               *int64   `protobuf:"varint,4,opt,name=update" json:"update,omitempty"`
	Delete               *int64   `protobuf:"varint,5,opt,name=delete" json:"delete,omitempty"`
	Custom               *int64   `protobuf:"varint,6,opt,name=custom" json:"custom,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntFieldOptions) Reset()         { *m = IntFieldOptions{} }
func (m *IntFieldOptions) String() string { return proto.CompactTextString(m) }
func (*IntFieldOptions) ProtoMessage()    {}
func (*IntFieldOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_options_64966814c91bee01, []int{1}
}
func (m *IntFieldOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntFieldOptions.Unmarshal(m, b)
}
func (m *IntFieldOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntFieldOptions.Marshal(b, m, deterministic)
}
func (dst *IntFieldOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntFieldOptions.Merge(dst, src)
}
func (m *IntFieldOptions) XXX_Size() int {
	return xxx_messageInfo_IntFieldOptions.Size(m)
}
func (m *IntFieldOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_IntFieldOptions.DiscardUnknown(m)
}

var xxx_messageInfo_IntFieldOptions proto.InternalMessageInfo

func (m *IntFieldOptions) GetCreate() int64 {
	if m != nil && m.Create != nil {
		return *m.Create
	}
	return 0
}

func (m *IntFieldOptions) GetGet() int64 {
	if m != nil && m.Get != nil {
		return *m.Get
	}
	return 0
}

func (m *IntFieldOptions) GetList() int64 {
	if m != nil && m.List != nil {
		return *m.List
	}
	return 0
}

func (m *IntFieldOptions) GetUpdate() int64 {
	if m != nil && m.Update != nil {
		return *m.Update
	}
	return 0
}

func (m *IntFieldOptions) GetDelete() int64 {
	if m != nil && m.Delete != nil {
		return *m.Delete
	}
	return 0
}

func (m *IntFieldOptions) GetCustom() int64 {
	if m != nil && m.Custom != nil {
		return *m.Custom
	}
	return 0
}

type StateMessageOptions struct {
	Type                 *StateMessageType `protobuf:"varint,1,req,name=type,enum=state.StateMessageType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StateMessageOptions) Reset()         { *m = StateMessageOptions{} }
func (m *StateMessageOptions) String() string { return proto.CompactTextString(m) }
func (*StateMessageOptions) ProtoMessage()    {}
func (*StateMessageOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_options_64966814c91bee01, []int{2}
}
func (m *StateMessageOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateMessageOptions.Unmarshal(m, b)
}
func (m *StateMessageOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateMessageOptions.Marshal(b, m, deterministic)
}
func (dst *StateMessageOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateMessageOptions.Merge(dst, src)
}
func (m *StateMessageOptions) XXX_Size() int {
	return xxx_messageInfo_StateMessageOptions.Size(m)
}
func (m *StateMessageOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_StateMessageOptions.DiscardUnknown(m)
}

var xxx_messageInfo_StateMessageOptions proto.InternalMessageInfo

func (m *StateMessageOptions) GetType() StateMessageType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return StateMessageType_REDUX_STATE
}

type StateFieldOptions struct {
	Timeout              *int64              `protobuf:"varint,1,opt,name=timeout" json:"timeout,omitempty"`
	Retries              *int64              `protobuf:"varint,2,opt,name=retries" json:"retries,omitempty"`
	Method               *StringFieldOptions `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
	MethodTimeout        *IntFieldOptions    `protobuf:"bytes,4,opt,name=method_timeout,json=methodTimeout" json:"method_timeout,omitempty"`
	MethodRetries        *IntFieldOptions    `protobuf:"bytes,5,opt,name=method_retries,json=methodRetries" json:"method_retries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *StateFieldOptions) Reset()         { *m = StateFieldOptions{} }
func (m *StateFieldOptions) String() string { return proto.CompactTextString(m) }
func (*StateFieldOptions) ProtoMessage()    {}
func (*StateFieldOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_options_64966814c91bee01, []int{3}
}
func (m *StateFieldOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateFieldOptions.Unmarshal(m, b)
}
func (m *StateFieldOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateFieldOptions.Marshal(b, m, deterministic)
}
func (dst *StateFieldOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateFieldOptions.Merge(dst, src)
}
func (m *StateFieldOptions) XXX_Size() int {
	return xxx_messageInfo_StateFieldOptions.Size(m)
}
func (m *StateFieldOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_StateFieldOptions.DiscardUnknown(m)
}

var xxx_messageInfo_StateFieldOptions proto.InternalMessageInfo

func (m *StateFieldOptions) GetTimeout() int64 {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return 0
}

func (m *StateFieldOptions) GetRetries() int64 {
	if m != nil && m.Retries != nil {
		return *m.Retries
	}
	return 0
}

func (m *StateFieldOptions) GetMethod() *StringFieldOptions {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *StateFieldOptions) GetMethodTimeout() *IntFieldOptions {
	if m != nil {
		return m.MethodTimeout
	}
	return nil
}

func (m *StateFieldOptions) GetMethodRetries() *IntFieldOptions {
	if m != nil {
		return m.MethodRetries
	}
	return nil
}

type StateFileOptions struct {
	DefaultTimeout *int64 `protobuf:"varint,1,opt,name=default_timeout,json=defaultTimeout" json:"default_timeout,omitempty"`
	DefaultRetries *int64 `protobuf:"varint,2,opt,name=default_retries,json=defaultRetries" json:"default_retries,omitempty"`
	// turn on debug logging
	Debug *bool `protobuf:"varint,3,opt,name=debug" json:"debug,omitempty"`
	// port to be used for api calls in epics
	Port *int64 `protobuf:"varint,4,opt,name=port" json:"port,omitempty"`
	// debounce time for api calls in epics
	Debounce *int64 `protobuf:"varint,5,opt,name=debounce" json:"debounce,omitempty"`
	// sets custom import path for proto typescript files
	ProtocTsPath *string `protobuf:"bytes,6,opt,name=protoc_ts_path,json=protocTsPath" json:"protoc_ts_path,omitempty"`
	// static hostname string for api calls
	Hostname *string `protobuf:"bytes,7,opt,name=hostname" json:"hostname,omitempty"`
	// hostname for api calls location in redux
	HostnameLocation *string `protobuf:"bytes,8,opt,name=hostname_location,json=hostnameLocation" json:"hostname_location,omitempty"`
	// add authorization bearer header using the token at this location in redux
	AuthTokenLocation    *string      `protobuf:"bytes,9,opt,name=auth_token_location,json=authTokenLocation" json:"auth_token_location,omitempty"`
	OutputType           *OutputTypes `protobuf:"varint,10,req,name=output_type,json=outputType,enum=state.OutputTypes" json:"output_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StateFileOptions) Reset()         { *m = StateFileOptions{} }
func (m *StateFileOptions) String() string { return proto.CompactTextString(m) }
func (*StateFileOptions) ProtoMessage()    {}
func (*StateFileOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_options_64966814c91bee01, []int{4}
}
func (m *StateFileOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StateFileOptions.Unmarshal(m, b)
}
func (m *StateFileOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StateFileOptions.Marshal(b, m, deterministic)
}
func (dst *StateFileOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateFileOptions.Merge(dst, src)
}
func (m *StateFileOptions) XXX_Size() int {
	return xxx_messageInfo_StateFileOptions.Size(m)
}
func (m *StateFileOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_StateFileOptions.DiscardUnknown(m)
}

var xxx_messageInfo_StateFileOptions proto.InternalMessageInfo

func (m *StateFileOptions) GetDefaultTimeout() int64 {
	if m != nil && m.DefaultTimeout != nil {
		return *m.DefaultTimeout
	}
	return 0
}

func (m *StateFileOptions) GetDefaultRetries() int64 {
	if m != nil && m.DefaultRetries != nil {
		return *m.DefaultRetries
	}
	return 0
}

func (m *StateFileOptions) GetDebug() bool {
	if m != nil && m.Debug != nil {
		return *m.Debug
	}
	return false
}

func (m *StateFileOptions) GetPort() int64 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *StateFileOptions) GetDebounce() int64 {
	if m != nil && m.Debounce != nil {
		return *m.Debounce
	}
	return 0
}

func (m *StateFileOptions) GetProtocTsPath() string {
	if m != nil && m.ProtocTsPath != nil {
		return *m.ProtocTsPath
	}
	return ""
}

func (m *StateFileOptions) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *StateFileOptions) GetHostnameLocation() string {
	if m != nil && m.HostnameLocation != nil {
		return *m.HostnameLocation
	}
	return ""
}

func (m *StateFileOptions) GetAuthTokenLocation() string {
	if m != nil && m.AuthTokenLocation != nil {
		return *m.AuthTokenLocation
	}
	return ""
}

func (m *StateFileOptions) GetOutputType() OutputTypes {
	if m != nil && m.OutputType != nil {
		return *m.OutputType
	}
	return OutputTypes_redux3
}

var E_StateOptions = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*StateMessageOptions)(nil),
	Field:         550002,
	Name:          "state.state_options",
	Tag:           "bytes,550002,opt,name=state_options,json=stateOptions",
	Filename:      "state/options.proto",
}

var E_StateFieldOptions = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*StateFieldOptions)(nil),
	Field:         550001,
	Name:          "state.state_field_options",
	Tag:           "bytes,550001,opt,name=state_field_options,json=stateFieldOptions",
	Filename:      "state/options.proto",
}

var E_StateFileOptions = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*StateFileOptions)(nil),
	Field:         550003,
	Name:          "state.state_file_options",
	Tag:           "bytes,550003,opt,name=state_file_options,json=stateFileOptions",
	Filename:      "state/options.proto",
}

func init() {
	proto.RegisterType((*StringFieldOptions)(nil), "state.StringFieldOptions")
	proto.RegisterType((*IntFieldOptions)(nil), "state.IntFieldOptions")
	proto.RegisterType((*StateMessageOptions)(nil), "state.StateMessageOptions")
	proto.RegisterType((*StateFieldOptions)(nil), "state.StateFieldOptions")
	proto.RegisterType((*StateFileOptions)(nil), "state.StateFileOptions")
	proto.RegisterEnum("state.StateMessageType", StateMessageType_name, StateMessageType_value)
	proto.RegisterEnum("state.OutputTypes", OutputTypes_name, OutputTypes_value)
	proto.RegisterExtension(E_StateOptions)
	proto.RegisterExtension(E_StateFieldOptions)
	proto.RegisterExtension(E_StateFileOptions)
}

func init() { proto.RegisterFile("state/options.proto", fileDescriptor_options_64966814c91bee01) }

var fileDescriptor_options_64966814c91bee01 = []byte{
	// 714 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xd1, 0x4e, 0xdb, 0x4a,
	0x10, 0xc5, 0x71, 0x12, 0xc2, 0x04, 0x82, 0xb3, 0x5c, 0x81, 0x2f, 0xba, 0x57, 0xa5, 0x51, 0xa5,
	0x22, 0x10, 0x4e, 0x0b, 0x7d, 0xa2, 0xea, 0x03, 0xd0, 0x54, 0x8a, 0x0a, 0xa4, 0x72, 0x8c, 0x84,
	0xfa, 0x62, 0x39, 0xf6, 0xe2, 0x58, 0x75, 0xbc, 0x96, 0x3d, 0x96, 0xe0, 0x4b, 0x2a, 0xf5, 0x8b,
	0xfa, 0x3b, 0x6d, 0x5f, 0xfa, 0x54, 0x55, 0xbb, 0x5e, 0x9b, 0x24, 0x84, 0xaa, 0x2f, 0xd1, 0xce,
	0xcc, 0x19, 0x9f, 0x33, 0xe3, 0xe3, 0x0d, 0x6c, 0xa4, 0xe8, 0x20, 0xed, 0xb2, 0x18, 0x03, 0x16,
	0xa5, 0x46, 0x9c, 0x30, 0x64, 0xa4, 0x26, 0x92, 0xdb, 0x3b, 0x3e, 0x63, 0x7e, 0x48, 0xbb, 0x22,
	0x39, 0xca, 0x6e, 0xba, 0x1e, 0x4d, 0xdd, 0x24, 0x88, 0x91, 0x25, 0x39, 0xb0, 0xf3, 0x45, 0x01,
	0x32, 0xc4, 0x24, 0x88, 0xfc, 0x77, 0x01, 0x0d, 0xbd, 0x41, 0xfe, 0x14, 0xb2, 0x09, 0x75, 0x37,
	0xa1, 0x0e, 0x52, 0x5d, 0xd9, 0x51, 0x76, 0x57, 0x4c, 0x19, 0x11, 0x0d, 0x54, 0x9f, 0xa2, 0x5e,
	0x11, 0x49, 0x7e, 0x24, 0x04, 0xaa, 0x61, 0x90, 0xa2, 0xae, 0x8a, 0x94, 0x38, 0xf3, 0xee, 0x2c,
	0xf6, 0x78, 0x77, 0x35, 0xef, 0xce, 0x23, 0x9e, 0xf7, 0x68, 0x48, 0x91, 0xea, 0xb5, 0x3c, 0x9f,
	0x47, 0x82, 0x2d, 0x4b, 0x91, 0x4d, 0xf4, 0xba, 0x64, 0x13, 0x51, 0xe7, 0xb3, 0x02, 0xeb, 0xfd,
	0x08, 0xff, 0xa0, 0x4c, 0x5d, 0xa4, 0x4c, 0x7d, 0xa8, 0x4c, 0x5d, 0xa8, 0x4c, 0x7d, 0x44, 0x99,
	0xfa, 0x88, 0x32, 0xb5, 0x54, 0x76, 0x0a, 0x1b, 0x43, 0xbe, 0xe1, 0x0b, 0x9a, 0xa6, 0x8e, 0x4f,
	0x0b, 0x71, 0xfb, 0x50, 0xc5, 0xbb, 0x98, 0x4b, 0xab, 0xec, 0xb6, 0x0e, 0xb7, 0x0c, 0xf1, 0x16,
	0x8c, 0x69, 0xa4, 0x75, 0x17, 0x53, 0x53, 0x80, 0x3a, 0x3f, 0x15, 0x68, 0x8b, 0xd2, 0xcc, 0x7c,
	0x3a, 0x2c, 0x63, 0x30, 0xa1, 0x2c, 0x43, 0x39, 0x60, 0x11, 0xf2, 0x4a, 0x42, 0x31, 0x09, 0x68,
	0x2a, 0xa7, 0x2c, 0x42, 0xf2, 0x12, 0xea, 0x13, 0x8a, 0x63, 0xe6, 0x89, 0x59, 0x9b, 0x87, 0xff,
	0x96, 0xc4, 0xf3, 0x2f, 0xd6, 0x94, 0x40, 0xf2, 0x06, 0x5a, 0xf9, 0xc9, 0x2e, 0xd8, 0xaa, 0xa2,
	0x75, 0x53, 0xb6, 0xce, 0xad, 0xdd, 0x5c, 0xcb, 0xd1, 0x96, 0xd4, 0x72, 0xdf, 0x5e, 0x48, 0xaa,
	0xfd, 0x4d, 0xbb, 0x99, 0x83, 0x3b, 0xbf, 0x2a, 0xa0, 0xc9, 0xd1, 0xc3, 0x72, 0x79, 0xcf, 0x61,
	0xdd, 0xa3, 0x37, 0x4e, 0x16, 0xa2, 0x3d, 0xbb, 0x81, 0x96, 0x4c, 0x17, 0xe4, 0x53, 0xc0, 0xd9,
	0x85, 0x14, 0x40, 0x49, 0x43, 0xfe, 0x81, 0x9a, 0x47, 0x47, 0x99, 0x2f, 0xd6, 0xd2, 0x30, 0xf3,
	0x80, 0xfb, 0x22, 0x66, 0x09, 0x4a, 0x07, 0x88, 0x33, 0xd9, 0x86, 0x86, 0x47, 0x47, 0x2c, 0x8b,
	0xdc, 0xc2, 0x01, 0x65, 0x4c, 0x9e, 0x41, 0x4b, 0x7c, 0x2b, 0xae, 0x8d, 0xa9, 0x1d, 0x3b, 0x38,
	0x96, 0x2e, 0x5d, 0xcd, 0xb3, 0x56, 0xfa, 0xc1, 0xc1, 0x31, 0x7f, 0xc2, 0x98, 0xa5, 0x18, 0x39,
	0x13, 0xaa, 0x2f, 0x8b, 0x7a, 0x19, 0x93, 0x7d, 0x68, 0x17, 0x67, 0x3b, 0x64, 0xae, 0xc3, 0xe7,
	0xd5, 0x1b, 0x02, 0xa4, 0x15, 0x85, 0x73, 0x99, 0x27, 0x06, 0x6c, 0x38, 0x19, 0x8e, 0x6d, 0x64,
	0x9f, 0x68, 0x74, 0x0f, 0x5f, 0x11, 0xf0, 0x36, 0x2f, 0x59, 0xbc, 0x52, 0xe2, 0x8f, 0xa0, 0xc9,
	0x32, 0x8c, 0x33, 0xb4, 0x85, 0xf5, 0x40, 0x58, 0x8f, 0xc8, 0xf7, 0x30, 0x10, 0x15, 0x6e, 0xba,
	0xd4, 0x04, 0x56, 0x06, 0x7b, 0x7d, 0xb9, 0xff, 0x29, 0x57, 0x92, 0x75, 0x68, 0x9a, 0xbd, 0xb7,
	0x57, 0xd7, 0xf6, 0xd0, 0x3a, 0xb1, 0x7a, 0xda, 0x12, 0x69, 0xc3, 0xda, 0xd9, 0xd5, 0xd0, 0x1a,
	0x5c, 0xd8, 0x27, 0x67, 0x56, 0x7f, 0x70, 0xa9, 0x29, 0x3c, 0xd5, 0xbb, 0xb6, 0x7a, 0xe6, 0xe5,
	0xc9, 0xb9, 0x7d, 0xde, 0xbf, 0x7c, 0xaf, 0x55, 0xf6, 0xba, 0xd0, 0x9c, 0x62, 0x21, 0x00, 0xf5,
	0x84, 0x7a, 0xd9, 0xed, 0x91, 0xb6, 0x54, 0x9e, 0x5f, 0x69, 0x0a, 0x69, 0x40, 0x75, 0xc2, 0x46,
	0xb7, 0x5a, 0xe5, 0x78, 0x04, 0x6b, 0x42, 0x9c, 0x2d, 0xaf, 0x2c, 0xf2, 0xc4, 0xc8, 0xaf, 0x29,
	0xa3, 0xb8, 0xa6, 0x8c, 0xd9, 0xcf, 0x4a, 0xff, 0xfe, 0xf5, 0xa9, 0x70, 0xd7, 0xf6, 0x82, 0x0f,
	0xaa, 0x70, 0xd8, 0xaa, 0x28, 0xc9, 0xe8, 0x38, 0x90, 0xd7, 0xa2, 0x7d, 0xc3, 0x5d, 0x58, 0x32,
	0xfd, 0xff, 0x80, 0x69, 0xda, 0xa5, 0xfa, 0x37, 0xc9, 0xa3, 0x4f, 0xf3, 0xcc, 0xf8, 0xb8, 0x9d,
	0xce, 0xa7, 0x8e, 0x29, 0x90, 0x82, 0x2a, 0xbc, 0x9f, 0xe9, 0xbf, 0x05, 0x4c, 0xa5, 0xd5, 0xf5,
	0x1f, 0x92, 0x68, 0x6b, 0x96, 0xa8, 0x04, 0x98, 0x5a, 0x3a, 0x97, 0x39, 0x3d, 0xfc, 0xf8, 0xc2,
	0x0f, 0x70, 0x9c, 0x8d, 0x0c, 0x97, 0x4d, 0xba, 0xe8, 0x46, 0x6e, 0xc8, 0x32, 0x2f, 0xbf, 0xd9,
	0xdd, 0x03, 0x9f, 0x46, 0x07, 0xf9, 0xff, 0x80, 0xf8, 0x7d, 0x2d, 0x7e, 0x7f, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xf7, 0x27, 0xde, 0x6a, 0x1b, 0x06, 0x00, 0x00,
}
