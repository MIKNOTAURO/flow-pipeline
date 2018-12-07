// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flow.proto

package flowprotob

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

type FlowMessage_FlowType int32

const (
	FlowMessage_FLOWUNKNOWN FlowMessage_FlowType = 0
	FlowMessage_NFV9        FlowMessage_FlowType = 9
	FlowMessage_IPFIX       FlowMessage_FlowType = 10
	FlowMessage_SFLOW       FlowMessage_FlowType = 5
)

var FlowMessage_FlowType_name = map[int32]string{
	0:  "FLOWUNKNOWN",
	9:  "NFV9",
	10: "IPFIX",
	5:  "SFLOW",
}
var FlowMessage_FlowType_value = map[string]int32{
	"FLOWUNKNOWN": 0,
	"NFV9":        9,
	"IPFIX":       10,
	"SFLOW":       5,
}

func (x FlowMessage_FlowType) String() string {
	return proto.EnumName(FlowMessage_FlowType_name, int32(x))
}
func (FlowMessage_FlowType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_flow_e561816bb870e53b, []int{0, 0}
}

// To be deprecated
type FlowMessage_IPType int32

const (
	FlowMessage_IPUNKNOWN FlowMessage_IPType = 0
	FlowMessage_IPv4      FlowMessage_IPType = 4
	FlowMessage_IPv6      FlowMessage_IPType = 6
)

var FlowMessage_IPType_name = map[int32]string{
	0: "IPUNKNOWN",
	4: "IPv4",
	6: "IPv6",
}
var FlowMessage_IPType_value = map[string]int32{
	"IPUNKNOWN": 0,
	"IPv4":      4,
	"IPv6":      6,
}

func (x FlowMessage_IPType) String() string {
	return proto.EnumName(FlowMessage_IPType_name, int32(x))
}
func (FlowMessage_IPType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_flow_e561816bb870e53b, []int{0, 1}
}

type FlowMessage struct {
	Type         FlowMessage_FlowType `protobuf:"varint,1,opt,name=Type,proto3,enum=flowprotob.FlowMessage_FlowType" json:"Type,omitempty"`
	TimeRecvd    uint64               `protobuf:"varint,2,opt,name=TimeRecvd,proto3" json:"TimeRecvd,omitempty"`
	SamplingRate uint64               `protobuf:"varint,3,opt,name=SamplingRate,proto3" json:"SamplingRate,omitempty"`
	SequenceNum  uint32               `protobuf:"varint,4,opt,name=SequenceNum,proto3" json:"SequenceNum,omitempty"`
	// Found inside packet
	TimeFlow uint64 `protobuf:"varint,5,opt,name=TimeFlow,proto3" json:"TimeFlow,omitempty"`
	// Source/destination addresses
	SrcIP     []byte             `protobuf:"bytes,6,opt,name=SrcIP,proto3" json:"SrcIP,omitempty"`
	DstIP     []byte             `protobuf:"bytes,7,opt,name=DstIP,proto3" json:"DstIP,omitempty"`
	IPversion FlowMessage_IPType `protobuf:"varint,8,opt,name=IPversion,proto3,enum=flowprotob.FlowMessage_IPType" json:"IPversion,omitempty"`
	// Size of the sampled packet
	Bytes   uint64 `protobuf:"varint,9,opt,name=Bytes,proto3" json:"Bytes,omitempty"`
	Packets uint64 `protobuf:"varint,10,opt,name=Packets,proto3" json:"Packets,omitempty"`
	// Routing information
	RouterAddr []byte `protobuf:"bytes,11,opt,name=RouterAddr,proto3" json:"RouterAddr,omitempty"`
	NextHop    []byte `protobuf:"bytes,12,opt,name=NextHop,proto3" json:"NextHop,omitempty"`
	NextHopAS  uint32 `protobuf:"varint,13,opt,name=NextHopAS,proto3" json:"NextHopAS,omitempty"`
	// Autonomous system information
	SrcAS uint32 `protobuf:"varint,14,opt,name=SrcAS,proto3" json:"SrcAS,omitempty"`
	DstAS uint32 `protobuf:"varint,15,opt,name=DstAS,proto3" json:"DstAS,omitempty"`
	// Prefix size
	SrcNet uint32 `protobuf:"varint,16,opt,name=SrcNet,proto3" json:"SrcNet,omitempty"`
	DstNet uint32 `protobuf:"varint,17,opt,name=DstNet,proto3" json:"DstNet,omitempty"`
	// Interfaces
	SrcIf uint32 `protobuf:"varint,18,opt,name=SrcIf,proto3" json:"SrcIf,omitempty"`
	DstIf uint32 `protobuf:"varint,19,opt,name=DstIf,proto3" json:"DstIf,omitempty"`
	// Layer 4 protocol
	Proto uint32 `protobuf:"varint,20,opt,name=Proto,proto3" json:"Proto,omitempty"`
	// Port for UDP and TCP
	SrcPort uint32 `protobuf:"varint,21,opt,name=SrcPort,proto3" json:"SrcPort,omitempty"`
	DstPort uint32 `protobuf:"varint,22,opt,name=DstPort,proto3" json:"DstPort,omitempty"`
	// IP and TCP special flags
	IPTos            uint32 `protobuf:"varint,23,opt,name=IPTos,proto3" json:"IPTos,omitempty"`
	ForwardingStatus uint32 `protobuf:"varint,24,opt,name=ForwardingStatus,proto3" json:"ForwardingStatus,omitempty"`
	IPTTL            uint32 `protobuf:"varint,25,opt,name=IPTTL,proto3" json:"IPTTL,omitempty"`
	TCPFlags         uint32 `protobuf:"varint,26,opt,name=TCPFlags,proto3" json:"TCPFlags,omitempty"`
	// Ethernet information
	SrcMac uint64 `protobuf:"varint,27,opt,name=SrcMac,proto3" json:"SrcMac,omitempty"`
	DstMac uint64 `protobuf:"varint,28,opt,name=DstMac,proto3" json:"DstMac,omitempty"`
	// 802.1q VLAN in sampled packet
	VlanId uint32 `protobuf:"varint,29,opt,name=VlanId,proto3" json:"VlanId,omitempty"`
	// Layer 3 protocol (IPv4/IPv6/ARP/...)
	Etype    uint32 `protobuf:"varint,30,opt,name=Etype,proto3" json:"Etype,omitempty"`
	IcmpType uint32 `protobuf:"varint,31,opt,name=IcmpType,proto3" json:"IcmpType,omitempty"`
	IcmpCode uint32 `protobuf:"varint,32,opt,name=IcmpCode,proto3" json:"IcmpCode,omitempty"`
	// Vlan
	SrcVlan uint32 `protobuf:"varint,33,opt,name=SrcVlan,proto3" json:"SrcVlan,omitempty"`
	DstVlan uint32 `protobuf:"varint,34,opt,name=DstVlan,proto3" json:"DstVlan,omitempty"`
	// Fragments (IPv4/IPv6)
	FragmentId           uint32   `protobuf:"varint,35,opt,name=FragmentId,proto3" json:"FragmentId,omitempty"`
	FragmentOffset       uint32   `protobuf:"varint,36,opt,name=FragmentOffset,proto3" json:"FragmentOffset,omitempty"`
	IPv6FlowLabel        uint32   `protobuf:"varint,37,opt,name=IPv6FlowLabel,proto3" json:"IPv6FlowLabel,omitempty"`
	SrcCountry           string   `protobuf:"bytes,100,opt,name=SrcCountry,proto3" json:"SrcCountry,omitempty"`
	DstCountry           string   `protobuf:"bytes,101,opt,name=DstCountry,proto3" json:"DstCountry,omitempty"`
	SrcASDB              uint32   `protobuf:"varint,102,opt,name=SrcASDB,proto3" json:"SrcASDB,omitempty"`
	DstASDB              uint32   `protobuf:"varint,103,opt,name=DstASDB,proto3" json:"DstASDB,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowMessage) Reset()         { *m = FlowMessage{} }
func (m *FlowMessage) String() string { return proto.CompactTextString(m) }
func (*FlowMessage) ProtoMessage()    {}
func (*FlowMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_flow_e561816bb870e53b, []int{0}
}
func (m *FlowMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowMessage.Unmarshal(m, b)
}
func (m *FlowMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowMessage.Marshal(b, m, deterministic)
}
func (dst *FlowMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowMessage.Merge(dst, src)
}
func (m *FlowMessage) XXX_Size() int {
	return xxx_messageInfo_FlowMessage.Size(m)
}
func (m *FlowMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowMessage.DiscardUnknown(m)
}

var xxx_messageInfo_FlowMessage proto.InternalMessageInfo

func (m *FlowMessage) GetType() FlowMessage_FlowType {
	if m != nil {
		return m.Type
	}
	return FlowMessage_FLOWUNKNOWN
}

func (m *FlowMessage) GetTimeRecvd() uint64 {
	if m != nil {
		return m.TimeRecvd
	}
	return 0
}

func (m *FlowMessage) GetSamplingRate() uint64 {
	if m != nil {
		return m.SamplingRate
	}
	return 0
}

func (m *FlowMessage) GetSequenceNum() uint32 {
	if m != nil {
		return m.SequenceNum
	}
	return 0
}

func (m *FlowMessage) GetTimeFlow() uint64 {
	if m != nil {
		return m.TimeFlow
	}
	return 0
}

func (m *FlowMessage) GetSrcIP() []byte {
	if m != nil {
		return m.SrcIP
	}
	return nil
}

func (m *FlowMessage) GetDstIP() []byte {
	if m != nil {
		return m.DstIP
	}
	return nil
}

func (m *FlowMessage) GetIPversion() FlowMessage_IPType {
	if m != nil {
		return m.IPversion
	}
	return FlowMessage_IPUNKNOWN
}

func (m *FlowMessage) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

func (m *FlowMessage) GetPackets() uint64 {
	if m != nil {
		return m.Packets
	}
	return 0
}

func (m *FlowMessage) GetRouterAddr() []byte {
	if m != nil {
		return m.RouterAddr
	}
	return nil
}

func (m *FlowMessage) GetNextHop() []byte {
	if m != nil {
		return m.NextHop
	}
	return nil
}

func (m *FlowMessage) GetNextHopAS() uint32 {
	if m != nil {
		return m.NextHopAS
	}
	return 0
}

func (m *FlowMessage) GetSrcAS() uint32 {
	if m != nil {
		return m.SrcAS
	}
	return 0
}

func (m *FlowMessage) GetDstAS() uint32 {
	if m != nil {
		return m.DstAS
	}
	return 0
}

func (m *FlowMessage) GetSrcNet() uint32 {
	if m != nil {
		return m.SrcNet
	}
	return 0
}

func (m *FlowMessage) GetDstNet() uint32 {
	if m != nil {
		return m.DstNet
	}
	return 0
}

func (m *FlowMessage) GetSrcIf() uint32 {
	if m != nil {
		return m.SrcIf
	}
	return 0
}

func (m *FlowMessage) GetDstIf() uint32 {
	if m != nil {
		return m.DstIf
	}
	return 0
}

func (m *FlowMessage) GetProto() uint32 {
	if m != nil {
		return m.Proto
	}
	return 0
}

func (m *FlowMessage) GetSrcPort() uint32 {
	if m != nil {
		return m.SrcPort
	}
	return 0
}

func (m *FlowMessage) GetDstPort() uint32 {
	if m != nil {
		return m.DstPort
	}
	return 0
}

func (m *FlowMessage) GetIPTos() uint32 {
	if m != nil {
		return m.IPTos
	}
	return 0
}

func (m *FlowMessage) GetForwardingStatus() uint32 {
	if m != nil {
		return m.ForwardingStatus
	}
	return 0
}

func (m *FlowMessage) GetIPTTL() uint32 {
	if m != nil {
		return m.IPTTL
	}
	return 0
}

func (m *FlowMessage) GetTCPFlags() uint32 {
	if m != nil {
		return m.TCPFlags
	}
	return 0
}

func (m *FlowMessage) GetSrcMac() uint64 {
	if m != nil {
		return m.SrcMac
	}
	return 0
}

func (m *FlowMessage) GetDstMac() uint64 {
	if m != nil {
		return m.DstMac
	}
	return 0
}

func (m *FlowMessage) GetVlanId() uint32 {
	if m != nil {
		return m.VlanId
	}
	return 0
}

func (m *FlowMessage) GetEtype() uint32 {
	if m != nil {
		return m.Etype
	}
	return 0
}

func (m *FlowMessage) GetIcmpType() uint32 {
	if m != nil {
		return m.IcmpType
	}
	return 0
}

func (m *FlowMessage) GetIcmpCode() uint32 {
	if m != nil {
		return m.IcmpCode
	}
	return 0
}

func (m *FlowMessage) GetSrcVlan() uint32 {
	if m != nil {
		return m.SrcVlan
	}
	return 0
}

func (m *FlowMessage) GetDstVlan() uint32 {
	if m != nil {
		return m.DstVlan
	}
	return 0
}

func (m *FlowMessage) GetFragmentId() uint32 {
	if m != nil {
		return m.FragmentId
	}
	return 0
}

func (m *FlowMessage) GetFragmentOffset() uint32 {
	if m != nil {
		return m.FragmentOffset
	}
	return 0
}

func (m *FlowMessage) GetIPv6FlowLabel() uint32 {
	if m != nil {
		return m.IPv6FlowLabel
	}
	return 0
}

func (m *FlowMessage) GetSrcCountry() string {
	if m != nil {
		return m.SrcCountry
	}
	return ""
}

func (m *FlowMessage) GetDstCountry() string {
	if m != nil {
		return m.DstCountry
	}
	return ""
}

func (m *FlowMessage) GetSrcASDB() uint32 {
	if m != nil {
		return m.SrcASDB
	}
	return 0
}

func (m *FlowMessage) GetDstASDB() uint32 {
	if m != nil {
		return m.DstASDB
	}
	return 0
}

func init() {
	proto.RegisterType((*FlowMessage)(nil), "flowprotob.FlowMessage")
	proto.RegisterEnum("flowprotob.FlowMessage_FlowType", FlowMessage_FlowType_name, FlowMessage_FlowType_value)
	proto.RegisterEnum("flowprotob.FlowMessage_IPType", FlowMessage_IPType_name, FlowMessage_IPType_value)
}

func init() { proto.RegisterFile("flow.proto", fileDescriptor_flow_e561816bb870e53b) }

var fileDescriptor_flow_e561816bb870e53b = []byte{
	// 689 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xd1, 0x73, 0xd2, 0x40,
	0x10, 0xc6, 0x45, 0x29, 0x85, 0xa3, 0xb4, 0xf1, 0xac, 0x75, 0xad, 0x15, 0x23, 0x56, 0x87, 0xb1,
	0x33, 0x3c, 0x68, 0xa7, 0x33, 0x8e, 0xbe, 0x40, 0x91, 0x31, 0x23, 0x4d, 0x33, 0x09, 0xb6, 0xbe,
	0x86, 0xe4, 0x92, 0x61, 0x0c, 0x09, 0x26, 0x07, 0xb5, 0x7f, 0x86, 0xff, 0xb1, 0xb3, 0x7b, 0x49,
	0xa0, 0x3a, 0xbe, 0xdd, 0xf7, 0xfb, 0xf6, 0xf6, 0xd8, 0x8f, 0xdc, 0x31, 0x16, 0x44, 0xc9, 0x4d,
	0x6f, 0x91, 0x26, 0x32, 0xe1, 0xb4, 0xa6, 0xe5, 0xb4, 0xf3, 0x9b, 0xb1, 0xe6, 0x28, 0x4a, 0x6e,
	0x2e, 0x44, 0x96, 0xb9, 0xa1, 0xe0, 0xa7, 0xac, 0x3a, 0xb9, 0x5d, 0x08, 0xa8, 0xe8, 0x95, 0xee,
	0xee, 0x3b, 0xbd, 0xb7, 0x2e, 0xed, 0x6d, 0x94, 0xd1, 0x1a, 0xeb, 0x6c, 0xaa, 0xe6, 0x47, 0xac,
	0x31, 0x99, 0xcd, 0x85, 0x2d, 0xbc, 0x95, 0x0f, 0xf7, 0xf5, 0x4a, 0xb7, 0x6a, 0xaf, 0x01, 0xef,
	0xb0, 0x1d, 0xc7, 0x9d, 0x2f, 0xa2, 0x59, 0x1c, 0xda, 0xae, 0x14, 0xf0, 0x80, 0x0a, 0xee, 0x30,
	0xae, 0xb3, 0xa6, 0x23, 0x7e, 0x2e, 0x45, 0xec, 0x09, 0x73, 0x39, 0x87, 0xaa, 0x5e, 0xe9, 0xb6,
	0xec, 0x4d, 0xc4, 0x0f, 0x59, 0x1d, 0x5b, 0xe2, 0xc9, 0xb0, 0x45, 0x1d, 0x4a, 0xcd, 0xf7, 0xd9,
	0x96, 0x93, 0x7a, 0x86, 0x05, 0x35, 0xbd, 0xd2, 0xdd, 0xb1, 0x95, 0x40, 0x3a, 0xcc, 0xa4, 0x61,
	0xc1, 0xb6, 0xa2, 0x24, 0xf8, 0x27, 0xd6, 0x30, 0xac, 0x95, 0x48, 0xb3, 0x59, 0x12, 0x43, 0x9d,
	0xc6, 0x6c, 0xff, 0x6f, 0x4c, 0xc3, 0xa2, 0x21, 0xd7, 0x1b, 0xb0, 0xe7, 0xe0, 0x56, 0x8a, 0x0c,
	0x1a, 0xf4, 0x13, 0x94, 0xe0, 0xc0, 0xb6, 0x2d, 0xd7, 0xfb, 0x21, 0x64, 0x06, 0x8c, 0x78, 0x21,
	0x79, 0x9b, 0x31, 0x3b, 0x59, 0x4a, 0x91, 0xf6, 0x7d, 0x3f, 0x85, 0x26, 0xfd, 0x90, 0x0d, 0x82,
	0x3b, 0x4d, 0xf1, 0x4b, 0x7e, 0x49, 0x16, 0xb0, 0x43, 0x66, 0x21, 0x31, 0xd3, 0x7c, 0xd9, 0x77,
	0xa0, 0x45, 0x79, 0xac, 0x41, 0x3e, 0x71, 0xdf, 0x81, 0x5d, 0x72, 0x94, 0xc8, 0x27, 0xee, 0x3b,
	0xb0, 0xa7, 0x28, 0x09, 0x7e, 0xc0, 0x6a, 0x4e, 0xea, 0x99, 0x42, 0x82, 0x46, 0x38, 0x57, 0xc8,
	0x87, 0x99, 0x44, 0xfe, 0x50, 0x71, 0xa5, 0x8a, 0x34, 0x03, 0xe0, 0x65, 0x6f, 0x23, 0x28, 0xd2,
	0x0c, 0xe0, 0x51, 0xd9, 0x5b, 0x51, 0x0b, 0x73, 0x83, 0x7d, 0x45, 0x49, 0xe0, 0x54, 0x4e, 0xea,
	0x59, 0x49, 0x2a, 0xe1, 0x31, 0xf1, 0x42, 0xa2, 0x33, 0xcc, 0x24, 0x39, 0x07, 0xca, 0xc9, 0x25,
	0x76, 0x32, 0xac, 0x49, 0x92, 0xc1, 0x13, 0xd5, 0x89, 0x04, 0x7f, 0xcb, 0xb4, 0x51, 0x92, 0xde,
	0xb8, 0xa9, 0x3f, 0x8b, 0x43, 0x47, 0xba, 0x72, 0x99, 0x01, 0x50, 0xc1, 0x3f, 0x3c, 0xef, 0x30,
	0x19, 0xc3, 0xd3, 0xb2, 0xc3, 0x64, 0x4c, 0xdf, 0xcd, 0xb9, 0x35, 0x8a, 0xdc, 0x30, 0x83, 0x43,
	0x32, 0x4a, 0x9d, 0x27, 0x73, 0xe1, 0x7a, 0xf0, 0x8c, 0xfe, 0xb6, 0x5c, 0xe5, 0xc9, 0x20, 0x3f,
	0x52, 0x5c, 0x29, 0xe4, 0x57, 0x91, 0x1b, 0x1b, 0x3e, 0x3c, 0x57, 0x89, 0x29, 0x85, 0x27, 0x7f,
	0x96, 0x78, 0x6d, 0xda, 0xea, 0x64, 0x12, 0x78, 0xb2, 0xe1, 0xcd, 0x17, 0x74, 0x9f, 0x5e, 0xa8,
	0x93, 0x0b, 0x5d, 0x78, 0xe7, 0x89, 0x2f, 0x40, 0x5f, 0x7b, 0xa8, 0xf3, 0xf4, 0xb0, 0x35, 0xbc,
	0x2c, 0xd3, 0x43, 0x99, 0xa7, 0x47, 0x4e, 0xa7, 0x4c, 0x8f, 0x9c, 0x36, 0x63, 0xa3, 0xd4, 0x0d,
	0xe7, 0x22, 0x96, 0x86, 0x0f, 0xaf, 0xc8, 0xdc, 0x20, 0xfc, 0x0d, 0xdb, 0x2d, 0xd4, 0x65, 0x10,
	0x64, 0x42, 0xc2, 0x31, 0xd5, 0xfc, 0x45, 0xf9, 0x31, 0x6b, 0x19, 0xd6, 0xea, 0x0c, 0x2f, 0xc1,
	0xd8, 0x9d, 0x8a, 0x08, 0x5e, 0x53, 0xd9, 0x5d, 0x88, 0xa7, 0x39, 0xa9, 0x77, 0x9e, 0x2c, 0x63,
	0x99, 0xde, 0x82, 0xaf, 0x57, 0xba, 0x0d, 0x7b, 0x83, 0xa0, 0x3f, 0xcc, 0x64, 0xe1, 0x0b, 0xe5,
	0xaf, 0x49, 0x3e, 0x61, 0xdf, 0x19, 0x0e, 0x20, 0x28, 0x27, 0x44, 0x99, 0x4f, 0x48, 0x4e, 0x58,
	0x4e, 0x88, 0xb2, 0xf3, 0x91, 0xd5, 0x8b, 0x57, 0x87, 0xef, 0xb1, 0xe6, 0x68, 0x7c, 0x79, 0xfd,
	0xcd, 0xfc, 0x6a, 0x5e, 0x5e, 0x9b, 0xda, 0x3d, 0x5e, 0x67, 0x55, 0x73, 0x74, 0xf5, 0x41, 0x6b,
	0xf0, 0x06, 0x7e, 0x04, 0x23, 0xe3, 0xbb, 0xc6, 0x70, 0xe9, 0x60, 0x99, 0xb6, 0xd5, 0x39, 0x61,
	0x35, 0x75, 0x97, 0x79, 0x0b, 0xaf, 0xff, 0x9d, 0x8d, 0x86, 0xb5, 0x3a, 0xd5, 0xaa, 0xf9, 0xea,
	0x4c, 0xab, 0x0d, 0x4e, 0xd8, 0xa1, 0x97, 0xcc, 0x7b, 0x5e, 0x94, 0x2c, 0xfd, 0x20, 0x72, 0x53,
	0xd1, 0x8b, 0x85, 0xa4, 0x27, 0xc2, 0x0d, 0xc3, 0x41, 0x6b, 0xe3, 0x81, 0xb0, 0xa6, 0xd3, 0x1a,
	0x3d, 0x1b, 0xef, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x79, 0x20, 0x35, 0x61, 0x05, 0x00,
	0x00,
}
