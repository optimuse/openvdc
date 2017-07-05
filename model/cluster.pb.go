// Code generated by protoc-gen-go.
// source: cluster.proto
// DO NOT EDIT!

package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Console_Transport int32

const (
	Console_SSH Console_Transport = 0
)

var Console_Transport_name = map[int32]string{
	0: "SSH",
}
var Console_Transport_value = map[string]int32{
	"SSH": 0,
}

func (x Console_Transport) String() string {
	return proto.EnumName(Console_Transport_name, int32(x))
}
func (Console_Transport) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

type NodeState_State int32

const (
	NodeState_REGISTERED NodeState_State = 0
)

var NodeState_State_name = map[int32]string{
	0: "REGISTERED",
}
var NodeState_State_value = map[string]int32{
	"REGISTERED": 0,
}

func (x NodeState_State) String() string {
	return proto.EnumName(NodeState_State_name, int32(x))
}
func (NodeState_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3, 0} }

type Console struct {
	Type     Console_Transport `protobuf:"varint,1,opt,name=type,enum=model.Console_Transport" json:"type,omitempty"`
	BindAddr string            `protobuf:"bytes,2,opt,name=bind_addr" json:"bind_addr,omitempty"`
}

func (m *Console) Reset()                    { *m = Console{} }
func (m *Console) String() string            { return proto.CompactTextString(m) }
func (*Console) ProtoMessage()               {}
func (*Console) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Console) GetType() Console_Transport {
	if m != nil {
		return m.Type
	}
	return Console_SSH
}

func (m *Console) GetBindAddr() string {
	if m != nil {
		return m.BindAddr
	}
	return ""
}

type ExecutorNode struct {
	Id        string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created_at" json:"created_at,omitempty"`
	Console   *Console                   `protobuf:"bytes,3,opt,name=console" json:"console,omitempty"`
	GrpcAddr  string                     `protobuf:"bytes,4,opt,name=grpc_addr" json:"grpc_addr,omitempty"`
	LastState *NodeState                 `protobuf:"bytes,5,opt,name=last_state" json:"last_state,omitempty"`
}

func (m *ExecutorNode) Reset()                    { *m = ExecutorNode{} }
func (m *ExecutorNode) String() string            { return proto.CompactTextString(m) }
func (*ExecutorNode) ProtoMessage()               {}
func (*ExecutorNode) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ExecutorNode) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ExecutorNode) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ExecutorNode) GetConsole() *Console {
	if m != nil {
		return m.Console
	}
	return nil
}

func (m *ExecutorNode) GetGrpcAddr() string {
	if m != nil {
		return m.GrpcAddr
	}
	return ""
}

func (m *ExecutorNode) GetLastState() *NodeState {
	if m != nil {
		return m.LastState
	}
	return nil
}

type SchedulerNode struct {
	Id        string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *SchedulerNode) Reset()                    { *m = SchedulerNode{} }
func (m *SchedulerNode) String() string            { return proto.CompactTextString(m) }
func (*SchedulerNode) ProtoMessage()               {}
func (*SchedulerNode) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *SchedulerNode) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SchedulerNode) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type NodeState struct {
	State     NodeState_State            `protobuf:"varint,1,opt,name=state,enum=model.NodeState_State" json:"state,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *NodeState) Reset()                    { *m = NodeState{} }
func (m *NodeState) String() string            { return proto.CompactTextString(m) }
func (*NodeState) ProtoMessage()               {}
func (*NodeState) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *NodeState) GetState() NodeState_State {
	if m != nil {
		return m.State
	}
	return NodeState_REGISTERED
}

func (m *NodeState) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Console)(nil), "model.Console")
	proto.RegisterType((*ExecutorNode)(nil), "model.ExecutorNode")
	proto.RegisterType((*SchedulerNode)(nil), "model.SchedulerNode")
	proto.RegisterType((*NodeState)(nil), "model.NodeState")
	proto.RegisterEnum("model.Console_Transport", Console_Transport_name, Console_Transport_value)
	proto.RegisterEnum("model.NodeState_State", NodeState_State_name, NodeState_State_value)
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4f, 0xb3, 0x40,
	0x10, 0xc6, 0x4b, 0x5b, 0xde, 0x86, 0x79, 0x2d, 0x69, 0x36, 0x46, 0x49, 0xa3, 0xd1, 0x70, 0xea,
	0xc1, 0x2c, 0xa6, 0xde, 0xf4, 0xa6, 0x12, 0xf5, 0xe2, 0x61, 0xe9, 0x49, 0x0f, 0x0d, 0x2c, 0x2b,
	0x25, 0x02, 0x4b, 0x96, 0xc5, 0xd4, 0x2f, 0xe1, 0xa7, 0xf3, 0x03, 0x09, 0xbb, 0xfd, 0xa7, 0x57,
	0xbd, 0x6d, 0x66, 0x9e, 0x79, 0x9e, 0xdf, 0x64, 0x16, 0x86, 0x34, 0xab, 0x2b, 0xc9, 0x04, 0x2e,
	0x05, 0x97, 0x1c, 0x99, 0x39, 0x8f, 0x59, 0x36, 0xbe, 0x4a, 0x52, 0xb9, 0xa8, 0x23, 0x4c, 0x79,
	0xee, 0x25, 0x3c, 0x0b, 0x8b, 0xc4, 0x53, 0xfd, 0xa8, 0x7e, 0xf1, 0x4a, 0xf9, 0x5e, 0xb2, 0xca,
	0x93, 0x69, 0xce, 0x2a, 0x19, 0xe6, 0xe5, 0xf6, 0xa5, 0x3d, 0xdc, 0x57, 0x18, 0xdc, 0xf0, 0xa2,
	0xe2, 0x19, 0x43, 0x67, 0xd0, 0x6f, 0xd5, 0x8e, 0x71, 0x6a, 0x4c, 0xec, 0xa9, 0x83, 0x95, 0x3b,
	0x5e, 0x75, 0xf1, 0x4c, 0x84, 0x45, 0x55, 0x72, 0x21, 0x89, 0x52, 0xa1, 0x23, 0xb0, 0xa2, 0xb4,
	0x88, 0xe7, 0x61, 0x1c, 0x0b, 0xa7, 0xdb, 0x8c, 0x58, 0x64, 0x5b, 0x70, 0xf7, 0xc1, 0xda, 0x0c,
	0xa0, 0x01, 0xf4, 0x82, 0xe0, 0x7e, 0xd4, 0x71, 0x3f, 0x0d, 0xd8, 0xf3, 0x97, 0x8c, 0xd6, 0x92,
	0x8b, 0xc7, 0xc6, 0x1d, 0xd9, 0xd0, 0x4d, 0x63, 0x15, 0x68, 0x91, 0xe6, 0x85, 0x2e, 0x01, 0xa8,
	0x60, 0xa1, 0x64, 0x8d, 0x8d, 0x54, 0xae, 0xff, 0xa7, 0x63, 0x9c, 0x70, 0x9e, 0x34, 0x04, 0xeb,
	0xa5, 0xf0, 0x6c, 0xbd, 0x03, 0xd9, 0x51, 0xa3, 0x09, 0x0c, 0xa8, 0x66, 0x75, 0x7a, 0x6a, 0xd0,
	0xfe, 0xbe, 0x01, 0x59, 0xb7, 0x5b, 0xf4, 0x44, 0x94, 0x54, 0xa3, 0xf7, 0x35, 0xfa, 0xa6, 0x80,
	0xce, 0x01, 0xb2, 0xb0, 0x92, 0xf3, 0x26, 0x41, 0x32, 0xc7, 0x54, 0x56, 0xa3, 0x95, 0x55, 0x0b,
	0x1d, 0xb4, 0x75, 0xb2, 0xa3, 0x71, 0x9f, 0x61, 0x18, 0xd0, 0x05, 0x8b, 0xeb, 0x8c, 0xfd, 0xf9,
	0x5a, 0xee, 0x87, 0x01, 0xd6, 0x26, 0xb6, 0xb9, 0x91, 0xa9, 0xb9, 0xf4, 0x91, 0x0e, 0x7e, 0x72,
	0x61, 0x4d, 0xa7, 0x45, 0xbf, 0xca, 0x3d, 0x04, 0x53, 0x47, 0xda, 0x00, 0xc4, 0xbf, 0x7b, 0x08,
	0x66, 0x3e, 0xf1, 0x6f, 0x47, 0x9d, 0xeb, 0x93, 0xa7, 0xe3, 0x9d, 0x0f, 0x17, 0x2e, 0xab, 0x85,
	0xc7, 0x4b, 0x56, 0xbc, 0xc5, 0xd4, 0x53, 0x30, 0xd1, 0x3f, 0xe5, 0x7c, 0xf1, 0x15, 0x00, 0x00,
	0xff, 0xff, 0xc4, 0xc8, 0x99, 0x8c, 0xae, 0x02, 0x00, 0x00,
}
