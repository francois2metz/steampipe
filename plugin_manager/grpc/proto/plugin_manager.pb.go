// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: plugin_manager.proto

package proto

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

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connection string `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetConnection() string {
	if x != nil {
		return x.Connection
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reattach *ReattachConfig `protobuf:"bytes,1,opt,name=reattach,proto3" json:"reattach,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetReattach() *ReattachConfig {
	if x != nil {
		return x.Reattach
	}
	return nil
}

type ShutdownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ShutdownRequest) Reset() {
	*x = ShutdownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownRequest) ProtoMessage() {}

func (x *ShutdownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownRequest.ProtoReflect.Descriptor instead.
func (*ShutdownRequest) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{2}
}

type ShutdownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ShutdownResponse) Reset() {
	*x = ShutdownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownResponse) ProtoMessage() {}

func (x *ShutdownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownResponse.ProtoReflect.Descriptor instead.
func (*ShutdownResponse) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{3}
}

type ReattachConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol        string   `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	ProtocolVersion int64    `protobuf:"varint,2,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
	Addr            *NetAddr `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	Pid             int64    `protobuf:"varint,4,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *ReattachConfig) Reset() {
	*x = ReattachConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReattachConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReattachConfig) ProtoMessage() {}

func (x *ReattachConfig) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReattachConfig.ProtoReflect.Descriptor instead.
func (*ReattachConfig) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{4}
}

func (x *ReattachConfig) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ReattachConfig) GetProtocolVersion() int64 {
	if x != nil {
		return x.ProtocolVersion
	}
	return 0
}

func (x *ReattachConfig) GetAddr() *NetAddr {
	if x != nil {
		return x.Addr
	}
	return nil
}

func (x *ReattachConfig) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type NetAddr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string `protobuf:"bytes,1,opt,name=Network,proto3" json:"Network,omitempty"` // name of the network (for example, "tcp", "udp")
	Address string `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"` // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
}

func (x *NetAddr) Reset() {
	*x = NetAddr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetAddr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetAddr) ProtoMessage() {}

func (x *NetAddr) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetAddr.ProtoReflect.Descriptor instead.
func (*NetAddr) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{5}
}

func (x *NetAddr) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *NetAddr) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ConnectionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugin          string `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	PluginShortName string `protobuf:"bytes,2,opt,name=plugin_short_name,json=pluginShortName,proto3" json:"plugin_short_name,omitempty"`
	Config          string `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ConnectionConfig) Reset() {
	*x = ConnectionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionConfig) ProtoMessage() {}

func (x *ConnectionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionConfig.ProtoReflect.Descriptor instead.
func (*ConnectionConfig) Descriptor() ([]byte, []int) {
	return file_plugin_manager_proto_rawDescGZIP(), []int{6}
}

func (x *ConnectionConfig) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

func (x *ConnectionConfig) GetPluginShortName() string {
	if x != nil {
		return x.PluginShortName
	}
	return ""
}

func (x *ConnectionConfig) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

var File_plugin_manager_proto protoreflect.FileDescriptor

var file_plugin_manager_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x72, 0x65,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x08, 0x72, 0x65, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x22, 0x11, 0x0a,
	0x0f, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x12, 0x0a, 0x10, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x22,
	0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x52, 0x04, 0x61, 0x64,
	0x64, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x70, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x6e, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12,
	0x2a, 0x0a, 0x11, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x32, 0x7e, 0x0a, 0x0d, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e,
	0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_manager_proto_rawDescOnce sync.Once
	file_plugin_manager_proto_rawDescData = file_plugin_manager_proto_rawDesc
)

func file_plugin_manager_proto_rawDescGZIP() []byte {
	file_plugin_manager_proto_rawDescOnce.Do(func() {
		file_plugin_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_manager_proto_rawDescData)
	})
	return file_plugin_manager_proto_rawDescData
}

var file_plugin_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_plugin_manager_proto_goTypes = []interface{}{
	(*GetRequest)(nil),       // 0: proto.GetRequest
	(*GetResponse)(nil),      // 1: proto.GetResponse
	(*ShutdownRequest)(nil),  // 2: proto.ShutdownRequest
	(*ShutdownResponse)(nil), // 3: proto.ShutdownResponse
	(*ReattachConfig)(nil),   // 4: proto.ReattachConfig
	(*NetAddr)(nil),          // 5: proto.NetAddr
	(*ConnectionConfig)(nil), // 6: proto.ConnectionConfig
}
var file_plugin_manager_proto_depIdxs = []int32{
	4, // 0: proto.GetResponse.reattach:type_name -> proto.ReattachConfig
	5, // 1: proto.ReattachConfig.addr:type_name -> proto.NetAddr
	0, // 2: proto.PluginManager.Get:input_type -> proto.GetRequest
	2, // 3: proto.PluginManager.Shutdown:input_type -> proto.ShutdownRequest
	1, // 4: proto.PluginManager.Get:output_type -> proto.GetResponse
	3, // 5: proto.PluginManager.Shutdown:output_type -> proto.ShutdownResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_plugin_manager_proto_init() }
func file_plugin_manager_proto_init() {
	if File_plugin_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugin_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_plugin_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_plugin_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShutdownRequest); i {
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
		file_plugin_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShutdownResponse); i {
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
		file_plugin_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReattachConfig); i {
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
		file_plugin_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetAddr); i {
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
		file_plugin_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionConfig); i {
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
			RawDescriptor: file_plugin_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugin_manager_proto_goTypes,
		DependencyIndexes: file_plugin_manager_proto_depIdxs,
		MessageInfos:      file_plugin_manager_proto_msgTypes,
	}.Build()
	File_plugin_manager_proto = out.File
	file_plugin_manager_proto_rawDesc = nil
	file_plugin_manager_proto_goTypes = nil
	file_plugin_manager_proto_depIdxs = nil
}
