// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/tcp_proxy/v2/tcp_proxy.proto

package envoy_config_filter_network_tcp_proxy_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/accesslog/v2"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type TcpProxy struct {
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*TcpProxy_Cluster
	//	*TcpProxy_WeightedClusters
	ClusterSpecifier      isTcpProxy_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	MetadataMatch         *core.Metadata              `protobuf:"bytes,9,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	IdleTimeout           *duration.Duration          `protobuf:"bytes,8,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	DownstreamIdleTimeout *duration.Duration          `protobuf:"bytes,3,opt,name=downstream_idle_timeout,json=downstreamIdleTimeout,proto3" json:"downstream_idle_timeout,omitempty"`
	UpstreamIdleTimeout   *duration.Duration          `protobuf:"bytes,4,opt,name=upstream_idle_timeout,json=upstreamIdleTimeout,proto3" json:"upstream_idle_timeout,omitempty"`
	AccessLog             []*v2.AccessLog             `protobuf:"bytes,5,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	DeprecatedV1          *TcpProxy_DeprecatedV1      `protobuf:"bytes,6,opt,name=deprecated_v1,json=deprecatedV1,proto3" json:"deprecated_v1,omitempty"` // Deprecated: Do not use.
	MaxConnectAttempts    *wrappers.UInt32Value       `protobuf:"bytes,7,opt,name=max_connect_attempts,json=maxConnectAttempts,proto3" json:"max_connect_attempts,omitempty"`
	HashPolicy            []*_type.HashPolicy         `protobuf:"bytes,11,rep,name=hash_policy,json=hashPolicy,proto3" json:"hash_policy,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                    `json:"-"`
	XXX_unrecognized      []byte                      `json:"-"`
	XXX_sizecache         int32                       `json:"-"`
}

func (m *TcpProxy) Reset()         { *m = TcpProxy{} }
func (m *TcpProxy) String() string { return proto.CompactTextString(m) }
func (*TcpProxy) ProtoMessage()    {}
func (*TcpProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0}
}

func (m *TcpProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy.Unmarshal(m, b)
}
func (m *TcpProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy.Marshal(b, m, deterministic)
}
func (m *TcpProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy.Merge(m, src)
}
func (m *TcpProxy) XXX_Size() int {
	return xxx_messageInfo_TcpProxy.Size(m)
}
func (m *TcpProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy proto.InternalMessageInfo

func (m *TcpProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

type isTcpProxy_ClusterSpecifier interface {
	isTcpProxy_ClusterSpecifier()
}

type TcpProxy_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

type TcpProxy_WeightedClusters struct {
	WeightedClusters *TcpProxy_WeightedCluster `protobuf:"bytes,10,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*TcpProxy_Cluster) isTcpProxy_ClusterSpecifier() {}

func (*TcpProxy_WeightedClusters) isTcpProxy_ClusterSpecifier() {}

func (m *TcpProxy) GetClusterSpecifier() isTcpProxy_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *TcpProxy) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*TcpProxy_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *TcpProxy) GetWeightedClusters() *TcpProxy_WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*TcpProxy_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (m *TcpProxy) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func (m *TcpProxy) GetIdleTimeout() *duration.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetDownstreamIdleTimeout() *duration.Duration {
	if m != nil {
		return m.DownstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetUpstreamIdleTimeout() *duration.Duration {
	if m != nil {
		return m.UpstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetAccessLog() []*v2.AccessLog {
	if m != nil {
		return m.AccessLog
	}
	return nil
}

// Deprecated: Do not use.
func (m *TcpProxy) GetDeprecatedV1() *TcpProxy_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

func (m *TcpProxy) GetMaxConnectAttempts() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnectAttempts
	}
	return nil
}

func (m *TcpProxy) GetHashPolicy() []*_type.HashPolicy {
	if m != nil {
		return m.HashPolicy
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TcpProxy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TcpProxy_Cluster)(nil),
		(*TcpProxy_WeightedClusters)(nil),
	}
}

// Deprecated: Do not use.
type TcpProxy_DeprecatedV1 struct {
	Routes               []*TcpProxy_DeprecatedV1_TCPRoute `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *TcpProxy_DeprecatedV1) Reset()         { *m = TcpProxy_DeprecatedV1{} }
func (m *TcpProxy_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_DeprecatedV1) ProtoMessage()    {}
func (*TcpProxy_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0, 0}
}

func (m *TcpProxy_DeprecatedV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Unmarshal(m, b)
}
func (m *TcpProxy_DeprecatedV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Marshal(b, m, deterministic)
}
func (m *TcpProxy_DeprecatedV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_DeprecatedV1.Merge(m, src)
}
func (m *TcpProxy_DeprecatedV1) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Size(m)
}
func (m *TcpProxy_DeprecatedV1) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_DeprecatedV1.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_DeprecatedV1 proto.InternalMessageInfo

func (m *TcpProxy_DeprecatedV1) GetRoutes() []*TcpProxy_DeprecatedV1_TCPRoute {
	if m != nil {
		return m.Routes
	}
	return nil
}

type TcpProxy_DeprecatedV1_TCPRoute struct {
	Cluster              string            `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	DestinationIpList    []*core.CidrRange `protobuf:"bytes,2,rep,name=destination_ip_list,json=destinationIpList,proto3" json:"destination_ip_list,omitempty"`
	DestinationPorts     string            `protobuf:"bytes,3,opt,name=destination_ports,json=destinationPorts,proto3" json:"destination_ports,omitempty"`
	SourceIpList         []*core.CidrRange `protobuf:"bytes,4,rep,name=source_ip_list,json=sourceIpList,proto3" json:"source_ip_list,omitempty"`
	SourcePorts          string            `protobuf:"bytes,5,opt,name=source_ports,json=sourcePorts,proto3" json:"source_ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) Reset()         { *m = TcpProxy_DeprecatedV1_TCPRoute{} }
func (m *TcpProxy_DeprecatedV1_TCPRoute) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_DeprecatedV1_TCPRoute) ProtoMessage()    {}
func (*TcpProxy_DeprecatedV1_TCPRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0, 0, 0}
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Unmarshal(m, b)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Marshal(b, m, deterministic)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Merge(m, src)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Size(m)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute proto.InternalMessageInfo

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetDestinationIpList() []*core.CidrRange {
	if m != nil {
		return m.DestinationIpList
	}
	return nil
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetDestinationPorts() string {
	if m != nil {
		return m.DestinationPorts
	}
	return ""
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetSourceIpList() []*core.CidrRange {
	if m != nil {
		return m.SourceIpList
	}
	return nil
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetSourcePorts() string {
	if m != nil {
		return m.SourcePorts
	}
	return ""
}

type TcpProxy_WeightedCluster struct {
	Clusters             []*TcpProxy_WeightedCluster_ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *TcpProxy_WeightedCluster) Reset()         { *m = TcpProxy_WeightedCluster{} }
func (m *TcpProxy_WeightedCluster) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_WeightedCluster) ProtoMessage()    {}
func (*TcpProxy_WeightedCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0, 1}
}

func (m *TcpProxy_WeightedCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Unmarshal(m, b)
}
func (m *TcpProxy_WeightedCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Marshal(b, m, deterministic)
}
func (m *TcpProxy_WeightedCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_WeightedCluster.Merge(m, src)
}
func (m *TcpProxy_WeightedCluster) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Size(m)
}
func (m *TcpProxy_WeightedCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_WeightedCluster.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_WeightedCluster proto.InternalMessageInfo

func (m *TcpProxy_WeightedCluster) GetClusters() []*TcpProxy_WeightedCluster_ClusterWeight {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type TcpProxy_WeightedCluster_ClusterWeight struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weight               uint32         `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	MetadataMatch        *core.Metadata `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) Reset() {
	*m = TcpProxy_WeightedCluster_ClusterWeight{}
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_WeightedCluster_ClusterWeight) ProtoMessage()    {}
func (*TcpProxy_WeightedCluster_ClusterWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0, 1, 0}
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Unmarshal(m, b)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Marshal(b, m, deterministic)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Merge(m, src)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Size(m)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight proto.InternalMessageInfo

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func init() {
	proto.RegisterType((*TcpProxy)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy")
	proto.RegisterType((*TcpProxy_DeprecatedV1)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.DeprecatedV1")
	proto.RegisterType((*TcpProxy_DeprecatedV1_TCPRoute)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.DeprecatedV1.TCPRoute")
	proto.RegisterType((*TcpProxy_WeightedCluster)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.WeightedCluster")
	proto.RegisterType((*TcpProxy_WeightedCluster_ClusterWeight)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.WeightedCluster.ClusterWeight")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/tcp_proxy/v2/tcp_proxy.proto", fileDescriptor_1f6b35dbcbad27ba)
}

var fileDescriptor_1f6b35dbcbad27ba = []byte{
	// 859 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0x5e, 0x27, 0xfd, 0x91, 0x4e, 0x12, 0x68, 0x67, 0x59, 0xd6, 0x64, 0xab, 0xdd, 0x2e, 0xa7,
	0x88, 0x15, 0x36, 0x9b, 0x4a, 0x08, 0x21, 0x24, 0x54, 0x77, 0x0f, 0x2d, 0x6a, 0xa5, 0x60, 0x95,
	0xdd, 0xa3, 0x35, 0xb5, 0x5f, 0x9c, 0x01, 0xdb, 0x33, 0xcc, 0x8c, 0x93, 0xf4, 0xc2, 0x91, 0xeb,
	0x02, 0x27, 0x2e, 0xfc, 0x93, 0x9c, 0xd0, 0x1e, 0x10, 0xf2, 0xcc, 0x38, 0x71, 0xbb, 0x41, 0x5b,
	0xe8, 0xa9, 0x9e, 0xf7, 0xde, 0xf7, 0x7d, 0xaf, 0xef, 0x57, 0xd0, 0x17, 0x50, 0xcc, 0xd8, 0x95,
	0x1f, 0xb3, 0x62, 0x42, 0x53, 0x7f, 0x42, 0x33, 0x05, 0xc2, 0x2f, 0x40, 0xcd, 0x99, 0xf8, 0xc1,
	0x57, 0x31, 0x8f, 0xb8, 0x60, 0x8b, 0x2b, 0x7f, 0x36, 0x5a, 0x3d, 0x3c, 0x2e, 0x98, 0x62, 0x78,
	0xa8, 0x91, 0x9e, 0x41, 0x7a, 0x06, 0xe9, 0x59, 0xa4, 0xb7, 0x0a, 0x9e, 0x8d, 0x06, 0x4f, 0x8c,
	0x06, 0xe1, 0xb4, 0xe2, 0x89, 0x99, 0x00, 0x9f, 0x24, 0x89, 0x00, 0x29, 0x0d, 0xd5, 0x60, 0xff,
	0xed, 0x80, 0x4b, 0x22, 0xc1, 0x7a, 0x3f, 0x5b, 0x97, 0x22, 0x89, 0x63, 0x90, 0x32, 0x63, 0x69,
	0x85, 0x58, 0x3e, 0xae, 0xf3, 0xa9, 0x2b, 0x0e, 0xfe, 0x94, 0xc8, 0x69, 0xc4, 0x59, 0x46, 0x63,
	0x9b, 0xf8, 0xe0, 0x71, 0xca, 0x58, 0x9a, 0x81, 0xaf, 0x5f, 0x97, 0xe5, 0xc4, 0x4f, 0x4a, 0x41,
	0x14, 0x65, 0xc5, 0xbf, 0xf9, 0xe7, 0x82, 0x70, 0x0e, 0xa2, 0xce, 0xf6, 0x71, 0x99, 0x70, 0xe2,
	0x93, 0xa2, 0x60, 0x4a, 0xc3, 0xa4, 0x9f, 0xd3, 0x54, 0x10, 0x55, 0xe7, 0xfb, 0x70, 0x46, 0x32,
	0x9a, 0x10, 0x05, 0x7e, 0xfd, 0x61, 0x1c, 0x1f, 0xff, 0xdc, 0x43, 0x9d, 0x8b, 0x98, 0x8f, 0xab,
	0xba, 0xe0, 0x21, 0xea, 0x4a, 0x45, 0x54, 0xc4, 0x05, 0x4c, 0xe8, 0xc2, 0x75, 0x0e, 0x9c, 0xe1,
	0x4e, 0xb0, 0xfd, 0x26, 0xd8, 0x10, 0xad, 0x03, 0x27, 0x44, 0x95, 0x6f, 0xac, 0x5d, 0x78, 0x80,
	0xb6, 0xe3, 0xac, 0x94, 0x0a, 0x84, 0xdb, 0xaa, 0xa2, 0x4e, 0xee, 0x85, 0xb5, 0x01, 0xff, 0x88,
	0xf6, 0xe6, 0x40, 0xd3, 0xa9, 0x82, 0x24, 0xb2, 0x36, 0xe9, 0xa2, 0x03, 0x67, 0xd8, 0x1d, 0x05,
	0xde, 0x6d, 0x1b, 0xe4, 0xd5, 0x49, 0x79, 0xaf, 0x2c, 0xd7, 0xb1, 0xa1, 0x3a, 0xb9, 0x17, 0xee,
	0xce, 0xaf, 0x9b, 0x24, 0x0e, 0xd0, 0x7b, 0x39, 0x28, 0x92, 0x10, 0x45, 0xa2, 0x9c, 0xa8, 0x78,
	0xea, 0xee, 0x68, 0xbd, 0x47, 0x56, 0x8f, 0x70, 0x5a, 0x71, 0x56, 0x5d, 0xf4, 0xce, 0x6d, 0x60,
	0xd8, 0xaf, 0x21, 0xe7, 0x15, 0x02, 0x7f, 0x85, 0x7a, 0x34, 0xc9, 0x20, 0x52, 0x34, 0x07, 0x56,
	0x2a, 0xb7, 0xa3, 0x19, 0x3e, 0xf2, 0x4c, 0xe5, 0xbd, 0xba, 0xf2, 0xde, 0x0b, 0xdb, 0x99, 0xb0,
	0x5b, 0x85, 0x5f, 0x98, 0x68, 0xfc, 0x2d, 0x7a, 0x98, 0xb0, 0x79, 0x21, 0x95, 0x00, 0x92, 0x47,
	0xd7, 0x88, 0xda, 0xef, 0x22, 0x7a, 0xb0, 0x42, 0x9e, 0x36, 0x28, 0xcf, 0xd1, 0x83, 0x92, 0xaf,
	0x23, 0xdc, 0x78, 0x17, 0xe1, 0xfd, 0x1a, 0xd7, 0xa4, 0xfb, 0x06, 0x21, 0x33, 0x93, 0x51, 0xc6,
	0x52, 0x77, 0xf3, 0xa0, 0x3d, 0xec, 0x8e, 0x9e, 0xad, 0xed, 0xc7, 0x6a, 0x74, 0x67, 0x23, 0xef,
	0x48, 0x3f, 0xce, 0x58, 0x1a, 0xee, 0x90, 0xfa, 0x13, 0x4f, 0x51, 0x3f, 0x01, 0x2e, 0x20, 0x26,
	0x55, 0x93, 0x67, 0xcf, 0xdd, 0x2d, 0x9d, 0xd2, 0xd7, 0xff, 0xa3, 0xbd, 0x2f, 0x96, 0x3c, 0x2f,
	0x9f, 0x07, 0x2d, 0xd7, 0x09, 0x7b, 0x49, 0xc3, 0x82, 0x5f, 0xa1, 0x0f, 0x72, 0xb2, 0x88, 0x62,
	0x56, 0x14, 0x10, 0xab, 0x88, 0x28, 0x05, 0x39, 0x57, 0xd2, 0xdd, 0xd6, 0x82, 0xfb, 0x6f, 0xd5,
	0xe0, 0xbb, 0xd3, 0x42, 0x1d, 0x8e, 0x5e, 0x92, 0xac, 0x04, 0x3d, 0xb9, 0x9f, 0xb4, 0x86, 0x4e,
	0x88, 0x73, 0xb2, 0x38, 0x36, 0x0c, 0x47, 0x96, 0x00, 0x1f, 0xa1, 0x6e, 0x63, 0x0d, 0xdd, 0xae,
	0xae, 0xc7, 0x87, 0xf6, 0x1f, 0xa8, 0xb6, 0xd4, 0x3b, 0x21, 0x72, 0x3a, 0xd6, 0xde, 0xa0, 0xf3,
	0x26, 0xd8, 0xfc, 0xcd, 0x69, 0xed, 0x3a, 0x21, 0x9a, 0x2e, 0xad, 0x83, 0x5f, 0xda, 0xa8, 0xd7,
	0x4c, 0x1f, 0x7f, 0x8f, 0xb6, 0x04, 0x2b, 0x15, 0x48, 0xd7, 0xd1, 0x74, 0x27, 0x77, 0xac, 0x87,
	0x77, 0x71, 0x3c, 0x0e, 0x2b, 0xc2, 0x3a, 0x81, 0x8e, 0x13, 0x5a, 0x85, 0xc1, 0xeb, 0x16, 0xea,
	0xd4, 0x6e, 0xfc, 0x74, 0xb5, 0x8e, 0x37, 0x96, 0x76, 0xb9, 0x95, 0x67, 0xe8, 0x7e, 0x02, 0x52,
	0xd1, 0x42, 0x8f, 0x48, 0x44, 0x79, 0x94, 0x51, 0xa9, 0xdc, 0x96, 0x4e, 0x74, 0x7f, 0xcd, 0x9e,
	0x1c, 0xd3, 0x44, 0x84, 0xa4, 0x48, 0x21, 0xdc, 0x6b, 0x00, 0x4f, 0xf9, 0x19, 0x95, 0x0a, 0x3f,
	0x43, 0x4d, 0x63, 0xc4, 0x99, 0x50, 0x52, 0x0f, 0xfa, 0x4e, 0xb8, 0xdb, 0x70, 0x8c, 0x2b, 0x7b,
	0xb5, 0x9d, 0x92, 0x95, 0x22, 0x86, 0xa5, 0xea, 0xc6, 0x2d, 0x54, 0x7b, 0x06, 0x63, 0x05, 0x9f,
	0x22, 0xfb, 0xb6, 0x5a, 0x9b, 0x5a, 0xab, 0x6b, 0x6c, 0x5a, 0xe6, 0xcb, 0x96, 0xeb, 0x0c, 0xfe,
	0x68, 0xa1, 0xf7, 0x6f, 0x1c, 0x0c, 0x3c, 0x43, 0x9d, 0xe5, 0x19, 0x32, 0x7d, 0x19, 0xdf, 0xfd,
	0x0c, 0x79, 0xf6, 0xaf, 0x31, 0x37, 0xfa, 0xb3, 0xd4, 0x1a, 0xfc, 0xea, 0xa0, 0xfe, 0xb5, 0x28,
	0xfc, 0x08, 0x6d, 0x14, 0x24, 0x87, 0x9b, 0x3d, 0xd2, 0x46, 0xfc, 0x04, 0x6d, 0x99, 0xbb, 0xa6,
	0x2f, 0x6a, 0x7f, 0x35, 0xbd, 0xd6, 0xbc, 0xe6, 0xc8, 0xb5, 0xff, 0xeb, 0x91, 0x0b, 0x5c, 0xb4,
	0x67, 0xf3, 0x8b, 0x24, 0x87, 0x98, 0x4e, 0x28, 0x08, 0xdc, 0xfe, 0x2b, 0x70, 0x82, 0x9f, 0xfe,
	0xfc, 0xfd, 0xef, 0xd7, 0x9b, 0x3e, 0xfe, 0xd4, 0x90, 0xc1, 0x42, 0x41, 0x21, 0xab, 0x5f, 0x12,
	0x5b, 0x1e, 0xb9, 0xae, 0x3e, 0x87, 0xe8, 0x73, 0xca, 0x8c, 0xbc, 0xb1, 0xdc, 0xb6, 0xae, 0x41,
	0xbf, 0x2e, 0xec, 0xb8, 0xda, 0xdc, 0xb1, 0x73, 0xb9, 0xa5, 0x57, 0xf8, 0xf0, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xee, 0xf0, 0x63, 0x95, 0xfd, 0x07, 0x00, 0x00,
}
