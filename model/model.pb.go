// Code generated by protoc-gen-go.
// source: model.proto
// DO NOT EDIT!

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	model.proto
	cluster.proto

It has these top-level messages:
	Instance
	InstanceState
	FailureMessage
	Template
	NoneTemplate
	LxcTemplate
	VmwareTemplate
	NullTemplate
	Console
	ExecutorNode
	SchedulerNode
	NodeState
*/
package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InstanceState_State int32

const (
	InstanceState_REGISTERED   InstanceState_State = 0
	InstanceState_QUEUED       InstanceState_State = 1
	InstanceState_STARTING     InstanceState_State = 2
	InstanceState_RUNNING      InstanceState_State = 3
	InstanceState_STOPPING     InstanceState_State = 4
	InstanceState_STOPPED      InstanceState_State = 5
	InstanceState_REBOOTING    InstanceState_State = 6
	InstanceState_SHUTTINGDOWN InstanceState_State = 7
	InstanceState_TERMINATED   InstanceState_State = 8
	InstanceState_FAILED       InstanceState_State = 9
)

var InstanceState_State_name = map[int32]string{
	0: "REGISTERED",
	1: "QUEUED",
	2: "STARTING",
	3: "RUNNING",
	4: "STOPPING",
	5: "STOPPED",
	6: "REBOOTING",
	7: "SHUTTINGDOWN",
	8: "TERMINATED",
	9: "FAILED",
}
var InstanceState_State_value = map[string]int32{
	"REGISTERED":   0,
	"QUEUED":       1,
	"STARTING":     2,
	"RUNNING":      3,
	"STOPPING":     4,
	"STOPPED":      5,
	"REBOOTING":    6,
	"SHUTTINGDOWN": 7,
	"TERMINATED":   8,
	"FAILED":       9,
}

func (x InstanceState_State) String() string {
	return proto.EnumName(InstanceState_State_name, int32(x))
}
func (InstanceState_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type FailureMessage_ErrorType int32

const (
	FailureMessage_FAILED_BOOT      FailureMessage_ErrorType = 0
	FailureMessage_FAILED_START     FailureMessage_ErrorType = 1
	FailureMessage_FAILED_STOP      FailureMessage_ErrorType = 2
	FailureMessage_FAILED_REBOOT    FailureMessage_ErrorType = 3
	FailureMessage_FAILED_TERMINATE FailureMessage_ErrorType = 4
)

var FailureMessage_ErrorType_name = map[int32]string{
	0: "FAILED_BOOT",
	1: "FAILED_START",
	2: "FAILED_STOP",
	3: "FAILED_REBOOT",
	4: "FAILED_TERMINATE",
}
var FailureMessage_ErrorType_value = map[string]int32{
	"FAILED_BOOT":      0,
	"FAILED_START":     1,
	"FAILED_STOP":      2,
	"FAILED_REBOOT":    3,
	"FAILED_TERMINATE": 4,
}

func (x FailureMessage_ErrorType) String() string {
	return proto.EnumName(FailureMessage_ErrorType_name, int32(x))
}
func (FailureMessage_ErrorType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type NullTemplate_CrashStage int32

const (
	NullTemplate_NONE    NullTemplate_CrashStage = 0
	NullTemplate_START   NullTemplate_CrashStage = 1
	NullTemplate_STOP    NullTemplate_CrashStage = 2
	NullTemplate_CREATE  NullTemplate_CrashStage = 3
	NullTemplate_DESTROY NullTemplate_CrashStage = 4
	NullTemplate_REBOOT  NullTemplate_CrashStage = 5
)

var NullTemplate_CrashStage_name = map[int32]string{
	0: "NONE",
	1: "START",
	2: "STOP",
	3: "CREATE",
	4: "DESTROY",
	5: "REBOOT",
}
var NullTemplate_CrashStage_value = map[string]int32{
	"NONE":    0,
	"START":   1,
	"STOP":    2,
	"CREATE":  3,
	"DESTROY": 4,
	"REBOOT":  5,
}

func (x NullTemplate_CrashStage) String() string {
	return proto.EnumName(NullTemplate_CrashStage_name, int32(x))
}
func (NullTemplate_CrashStage) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

type Instance struct {
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	SlaveId string `protobuf:"bytes,2,opt,name=slave_id" json:"slave_id,omitempty"`
	// string resource_id = 3; // Obsolete
	LastState     *InstanceState             `protobuf:"bytes,4,opt,name=last_state" json:"last_state,omitempty"`
	CreatedAt     *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=created_at" json:"created_at,omitempty"`
	Template      *Template                  `protobuf:"bytes,6,opt,name=template" json:"template,omitempty"`
	LatestFailure *FailureMessage            `protobuf:"bytes,7,opt,name=latest_failure,json=latestFailure" json:"latest_failure,omitempty"`
}

func (m *Instance) Reset()                    { *m = Instance{} }
func (m *Instance) String() string            { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()               {}
func (*Instance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Instance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Instance) GetSlaveId() string {
	if m != nil {
		return m.SlaveId
	}
	return ""
}

func (m *Instance) GetLastState() *InstanceState {
	if m != nil {
		return m.LastState
	}
	return nil
}

func (m *Instance) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Instance) GetTemplate() *Template {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *Instance) GetLatestFailure() *FailureMessage {
	if m != nil {
		return m.LatestFailure
	}
	return nil
}

type InstanceState struct {
	State     InstanceState_State        `protobuf:"varint,1,opt,name=state,enum=model.InstanceState_State" json:"state,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *InstanceState) Reset()                    { *m = InstanceState{} }
func (m *InstanceState) String() string            { return proto.CompactTextString(m) }
func (*InstanceState) ProtoMessage()               {}
func (*InstanceState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *InstanceState) GetState() InstanceState_State {
	if m != nil {
		return m.State
	}
	return InstanceState_REGISTERED
}

func (m *InstanceState) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type FailureMessage struct {
	ErrorType FailureMessage_ErrorType   `protobuf:"varint,1,opt,name=error_type,json=errorType,enum=model.FailureMessage_ErrorType" json:"error_type,omitempty"`
	FailedAt  *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=failed_at,json=failedAt" json:"failed_at,omitempty"`
}

func (m *FailureMessage) Reset()                    { *m = FailureMessage{} }
func (m *FailureMessage) String() string            { return proto.CompactTextString(m) }
func (*FailureMessage) ProtoMessage()               {}
func (*FailureMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FailureMessage) GetErrorType() FailureMessage_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return FailureMessage_FAILED_BOOT
}

func (m *FailureMessage) GetFailedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.FailedAt
	}
	return nil
}

type Template struct {
	TemplateUri string `protobuf:"bytes,1,opt,name=template_uri" json:"template_uri,omitempty"`
	// Types that are valid to be assigned to Item:
	//	*Template_None
	//	*Template_Lxc
	//	*Template_Null
	//	*Template_Vmware
	Item      isTemplate_Item            `protobuf_oneof:"Item"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *Template) Reset()                    { *m = Template{} }
func (m *Template) String() string            { return proto.CompactTextString(m) }
func (*Template) ProtoMessage()               {}
func (*Template) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isTemplate_Item interface {
	isTemplate_Item()
}

type Template_None struct {
	None *NoneTemplate `protobuf:"bytes,500,opt,name=none,oneof"`
}
type Template_Lxc struct {
	Lxc *LxcTemplate `protobuf:"bytes,501,opt,name=lxc,oneof"`
}
type Template_Null struct {
	Null *NullTemplate `protobuf:"bytes,502,opt,name=null,oneof"`
}
type Template_Vmware struct {
	Vmware *VmwareTemplate `protobuf:"bytes,503,opt,name=vmware,oneof"`
}

func (*Template_None) isTemplate_Item()   {}
func (*Template_Lxc) isTemplate_Item()    {}
func (*Template_Null) isTemplate_Item()   {}
func (*Template_Vmware) isTemplate_Item() {}

func (m *Template) GetItem() isTemplate_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *Template) GetTemplateUri() string {
	if m != nil {
		return m.TemplateUri
	}
	return ""
}

func (m *Template) GetNone() *NoneTemplate {
	if x, ok := m.GetItem().(*Template_None); ok {
		return x.None
	}
	return nil
}

func (m *Template) GetLxc() *LxcTemplate {
	if x, ok := m.GetItem().(*Template_Lxc); ok {
		return x.Lxc
	}
	return nil
}

func (m *Template) GetNull() *NullTemplate {
	if x, ok := m.GetItem().(*Template_Null); ok {
		return x.Null
	}
	return nil
}

func (m *Template) GetVmware() *VmwareTemplate {
	if x, ok := m.GetItem().(*Template_Vmware); ok {
		return x.Vmware
	}
	return nil
}

func (m *Template) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Template) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Template_OneofMarshaler, _Template_OneofUnmarshaler, _Template_OneofSizer, []interface{}{
		(*Template_None)(nil),
		(*Template_Lxc)(nil),
		(*Template_Null)(nil),
		(*Template_Vmware)(nil),
	}
}

func _Template_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Template)
	// Item
	switch x := m.Item.(type) {
	case *Template_None:
		b.EncodeVarint(500<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.None); err != nil {
			return err
		}
	case *Template_Lxc:
		b.EncodeVarint(501<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Lxc); err != nil {
			return err
		}
	case *Template_Null:
		b.EncodeVarint(502<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Null); err != nil {
			return err
		}
	case *Template_Vmware:
		b.EncodeVarint(503<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Vmware); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Template.Item has unexpected type %T", x)
	}
	return nil
}

func _Template_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Template)
	switch tag {
	case 500: // Item.none
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NoneTemplate)
		err := b.DecodeMessage(msg)
		m.Item = &Template_None{msg}
		return true, err
	case 501: // Item.lxc
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LxcTemplate)
		err := b.DecodeMessage(msg)
		m.Item = &Template_Lxc{msg}
		return true, err
	case 502: // Item.null
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NullTemplate)
		err := b.DecodeMessage(msg)
		m.Item = &Template_Null{msg}
		return true, err
	case 503: // Item.vmware
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(VmwareTemplate)
		err := b.DecodeMessage(msg)
		m.Item = &Template_Vmware{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Template_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Template)
	// Item
	switch x := m.Item.(type) {
	case *Template_None:
		s := proto.Size(x.None)
		n += proto.SizeVarint(500<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Template_Lxc:
		s := proto.Size(x.Lxc)
		n += proto.SizeVarint(501<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Template_Null:
		s := proto.Size(x.Null)
		n += proto.SizeVarint(502<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Template_Vmware:
		s := proto.Size(x.Vmware)
		n += proto.SizeVarint(503<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NoneTemplate struct {
}

func (m *NoneTemplate) Reset()                    { *m = NoneTemplate{} }
func (m *NoneTemplate) String() string            { return proto.CompactTextString(m) }
func (*NoneTemplate) ProtoMessage()               {}
func (*NoneTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type LxcTemplate struct {
	Vcpu        int32                    `protobuf:"varint,1,opt,name=vcpu" json:"vcpu,omitempty"`
	MemoryGb    int32                    `protobuf:"varint,2,opt,name=memory_gb" json:"memory_gb,omitempty"`
	MinVcpu     int32                    `protobuf:"varint,3,opt,name=min_vcpu" json:"min_vcpu,omitempty"`
	MinMemoryGb int32                    `protobuf:"varint,4,opt,name=min_memory_gb" json:"min_memory_gb,omitempty"`
	LxcImage    *LxcTemplate_Image       `protobuf:"bytes,5,opt,name=lxc_image" json:"lxc_image,omitempty"`
	Interfaces  []*LxcTemplate_Interface `protobuf:"bytes,6,rep,name=interfaces" json:"interfaces,omitempty"`
	LxcTemplate *LxcTemplate_Template    `protobuf:"bytes,7,opt,name=lxc_template" json:"lxc_template,omitempty"`
	NodeGroups  []string                 `protobuf:"bytes,8,rep,name=node_groups" json:"node_groups,omitempty"`
}

func (m *LxcTemplate) Reset()                    { *m = LxcTemplate{} }
func (m *LxcTemplate) String() string            { return proto.CompactTextString(m) }
func (*LxcTemplate) ProtoMessage()               {}
func (*LxcTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LxcTemplate) GetVcpu() int32 {
	if m != nil {
		return m.Vcpu
	}
	return 0
}

func (m *LxcTemplate) GetMemoryGb() int32 {
	if m != nil {
		return m.MemoryGb
	}
	return 0
}

func (m *LxcTemplate) GetMinVcpu() int32 {
	if m != nil {
		return m.MinVcpu
	}
	return 0
}

func (m *LxcTemplate) GetMinMemoryGb() int32 {
	if m != nil {
		return m.MinMemoryGb
	}
	return 0
}

func (m *LxcTemplate) GetLxcImage() *LxcTemplate_Image {
	if m != nil {
		return m.LxcImage
	}
	return nil
}

func (m *LxcTemplate) GetInterfaces() []*LxcTemplate_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *LxcTemplate) GetLxcTemplate() *LxcTemplate_Template {
	if m != nil {
		return m.LxcTemplate
	}
	return nil
}


func (m *LxcTemplate) GetNodeGroups() []string {
	if m != nil {
		return m.NodeGroups
	}
	return nil
}

type LxcTemplate_Image struct {
	DownloadUrl string `protobuf:"bytes,1,opt,name=download_url" json:"download_url,omitempty"`
	ChksumType  string `protobuf:"bytes,2,opt,name=chksum_type" json:"chksum_type,omitempty"`
	Chksum      string `protobuf:"bytes,3,opt,name=chksum" json:"chksum,omitempty"`
}

func (m *LxcTemplate_Image) Reset()                    { *m = LxcTemplate_Image{} }
func (m *LxcTemplate_Image) String() string            { return proto.CompactTextString(m) }
func (*LxcTemplate_Image) ProtoMessage()               {}
func (*LxcTemplate_Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }


func (m *LxcTemplate_Image) GetDownloadUrl() string {
	if m != nil {
		return m.DownloadUrl
	}
	return ""
}

func (m *LxcTemplate_Image) GetChksumType() string {
	if m != nil {
		return m.ChksumType
	}
	return ""
}

func (m *LxcTemplate_Image) GetChksum() string {
	if m != nil {
		return m.Chksum
	}
	return ""
}

type LxcTemplate_Interface struct {
	Type     string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Macaddr  string `protobuf:"bytes,2,opt,name=macaddr" json:"macaddr,omitempty"`
	Ipv4Addr string `protobuf:"bytes,3,opt,name=ipv4addr" json:"ipv4addr,omitempty"`
}

func (m *LxcTemplate_Interface) Reset()                    { *m = LxcTemplate_Interface{} }
func (m *LxcTemplate_Interface) String() string            { return proto.CompactTextString(m) }
func (*LxcTemplate_Interface) ProtoMessage()               {}
func (*LxcTemplate_Interface) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 1} }

func (m *LxcTemplate_Interface) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *LxcTemplate_Interface) GetMacaddr() string {
	if m != nil {
		return m.Macaddr
	}
	return ""
}

func (m *LxcTemplate_Interface) GetIpv4Addr() string {
	if m != nil {
		return m.Ipv4Addr
	}
	return ""
}

type LxcTemplate_Template struct {
	// Template specifies the name of the template.
	Template string `protobuf:"bytes,1,opt,name=template" json:"template,omitempty"`
	// Backend specifies the type of the backend.
	Backend int32 `protobuf:"varint,2,opt,name=backend" json:"backend,omitempty"`
	// Distro specifies the name of the distribution.
	Distro string `protobuf:"bytes,3,opt,name=distro" json:"distro,omitempty"`
	// Release specifies the name/version of the distribution.
	Release string `protobuf:"bytes,4,opt,name=release" json:"release,omitempty"`
	// Arch specified the architecture of the container.
	Arch string `protobuf:"bytes,5,opt,name=arch" json:"arch,omitempty"`
	// Variant specifies the variant of the image (default: "default").
	Variant string `protobuf:"bytes,6,opt,name=variant" json:"variant,omitempty"`
	// Image server (default: "images.linuxcontainers.org").
	Server string `protobuf:"bytes,7,opt,name=server" json:"server,omitempty"`
	// GPG keyid (default: 0x...).
	KeyId string `protobuf:"bytes,8,opt,name=key_id" json:"key_id,omitempty"`
	// GPG keyserver to use.
	KeyServer string `protobuf:"bytes,9,opt,name=key_server" json:"key_server,omitempty"`
	// Disable GPG validation (not recommended).
	DisableGpgValidation bool `protobuf:"varint,10,opt,name=disable_gpg_validation" json:"disable_gpg_validation,omitempty"`
	// Flush the local copy (if present).
	FlushCache bool `protobuf:"varint,11,opt,name=flush_cache" json:"flush_cache,omitempty"`
	// Force the use of the local copy even if expired.
	ForceCache bool `protobuf:"varint,12,opt,name=force_cache" json:"force_cache,omitempty"`
	// ExtraArgs provides a way to specify template specific args.
	ExtraArgs []string `protobuf:"bytes,13,rep,name=extra_args" json:"extra_args,omitempty"`
}

func (m *LxcTemplate_Template) Reset()                    { *m = LxcTemplate_Template{} }
func (m *LxcTemplate_Template) String() string            { return proto.CompactTextString(m) }
func (*LxcTemplate_Template) ProtoMessage()               {}
func (*LxcTemplate_Template) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 2} }

func (m *LxcTemplate_Template) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *LxcTemplate_Template) GetBackend() int32 {
	if m != nil {
		return m.Backend
	}
	return 0
}

func (m *LxcTemplate_Template) GetDistro() string {
	if m != nil {
		return m.Distro
	}
	return ""
}

func (m *LxcTemplate_Template) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func (m *LxcTemplate_Template) GetArch() string {
	if m != nil {
		return m.Arch
	}
	return ""
}

func (m *LxcTemplate_Template) GetVariant() string {
	if m != nil {
		return m.Variant
	}
	return ""
}

func (m *LxcTemplate_Template) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *LxcTemplate_Template) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *LxcTemplate_Template) GetKeyServer() string {
	if m != nil {
		return m.KeyServer
	}
	return ""
}

func (m *LxcTemplate_Template) GetDisableGpgValidation() bool {
	if m != nil {
		return m.DisableGpgValidation
	}
	return false
}

func (m *LxcTemplate_Template) GetFlushCache() bool {
	if m != nil {
		return m.FlushCache
	}
	return false
}

func (m *LxcTemplate_Template) GetForceCache() bool {
	if m != nil {
		return m.ForceCache
	}
	return false
}

func (m *LxcTemplate_Template) GetExtraArgs() []string {
	if m != nil {
		return m.ExtraArgs
	}
	return nil
}

type VmwareTemplate struct {
	Vcpu        int32                       `protobuf:"varint,1,opt,name=vcpu" json:"vcpu,omitempty"`
	MemoryGb    int32                       `protobuf:"varint,2,opt,name=memory_gb" json:"memory_gb,omitempty"`
	MinVcpu     int32                       `protobuf:"varint,3,opt,name=min_vcpu" json:"min_vcpu,omitempty"`
	MinMemoryGb int32                       `protobuf:"varint,4,opt,name=min_memory_gb" json:"min_memory_gb,omitempty"`
	VmwareImage *VmwareTemplate_Image       `protobuf:"bytes,5,opt,name=vmware_image" json:"vmware_image,omitempty"`
	Interfaces  []*VmwareTemplate_Interface `protobuf:"bytes,6,rep,name=interfaces" json:"interfaces,omitempty"`
	NodeGroups  []string                    `protobuf:"bytes,7,rep,name=node_groups" json:"node_groups,omitempty"`
}

func (m *VmwareTemplate) Reset()                    { *m = VmwareTemplate{} }
func (m *VmwareTemplate) String() string            { return proto.CompactTextString(m) }
func (*VmwareTemplate) ProtoMessage()               {}
func (*VmwareTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *VmwareTemplate) GetVcpu() int32 {
	if m != nil {
		return m.Vcpu
	}
	return 0
}

func (m *VmwareTemplate) GetMemoryGb() int32 {
	if m != nil {
		return m.MemoryGb
	}
	return 0
}

func (m *VmwareTemplate) GetMinVcpu() int32 {
	if m != nil {
		return m.MinVcpu
	}
	return 0
}

func (m *VmwareTemplate) GetMinMemoryGb() int32 {
	if m != nil {
		return m.MinMemoryGb
	}
	return 0
}

func (m *VmwareTemplate) GetNodeGroups() []string {
	if m != nil {
		return m.NodeGroups
	}
	return nil
}

func (m *VmwareTemplate) GetVmwareImage() *VmwareTemplate_Image {
	if m != nil {
		return m.VmwareImage
	}
	return nil
}

func (m *VmwareTemplate) GetInterfaces() []*VmwareTemplate_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

type VmwareTemplate_Image struct {
	DownloadUrl string `protobuf:"bytes,1,opt,name=download_url" json:"download_url,omitempty"`
	ChksumType  string `protobuf:"bytes,2,opt,name=chksum_type" json:"chksum_type,omitempty"`
	Chksum      string `protobuf:"bytes,3,opt,name=chksum" json:"chksum,omitempty"`
}

func (m *VmwareTemplate_Image) Reset()                    { *m = VmwareTemplate_Image{} }
func (m *VmwareTemplate_Image) String() string            { return proto.CompactTextString(m) }
func (*VmwareTemplate_Image) ProtoMessage()               {}
func (*VmwareTemplate_Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *VmwareTemplate_Image) GetDownloadUrl() string {
	if m != nil {
		return m.DownloadUrl
	}
	return ""
}

func (m *VmwareTemplate_Image) GetChksumType() string {
	if m != nil {
		return m.ChksumType
	}
	return ""
}

func (m *VmwareTemplate_Image) GetChksum() string {
	if m != nil {
		return m.Chksum
	}
	return ""
}

type VmwareTemplate_Interface struct {
	Type     string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Macaddr  string `protobuf:"bytes,2,opt,name=macaddr" json:"macaddr,omitempty"`
	Ipv4Addr string `protobuf:"bytes,3,opt,name=ipv4addr" json:"ipv4addr,omitempty"`
}

func (m *VmwareTemplate_Interface) Reset()                    { *m = VmwareTemplate_Interface{} }
func (m *VmwareTemplate_Interface) String() string            { return proto.CompactTextString(m) }
func (*VmwareTemplate_Interface) ProtoMessage()               {}
func (*VmwareTemplate_Interface) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 1} }

func (m *VmwareTemplate_Interface) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *VmwareTemplate_Interface) GetMacaddr() string {
	if m != nil {
		return m.Macaddr
	}
	return ""
}

func (m *VmwareTemplate_Interface) GetIpv4Addr() string {
	if m != nil {
		return m.Ipv4Addr
	}
	return ""
}

type NullTemplate struct {
	Vcpu       int32                   `protobuf:"varint,1,opt,name=vcpu" json:"vcpu,omitempty"`
	MemoryGb   int32                   `protobuf:"varint,2,opt,name=memory_gb" json:"memory_gb,omitempty"`
	CrashStage NullTemplate_CrashStage `protobuf:"varint,3,opt,name=crash_stage,enum=model.NullTemplate_CrashStage" json:"crash_stage,omitempty"`
	NodeGroups []string                `protobuf:"bytes,4,rep,name=node_groups" json:"node_groups,omitempty"`
}

func (m *NullTemplate) Reset()                    { *m = NullTemplate{} }
func (m *NullTemplate) String() string            { return proto.CompactTextString(m) }
func (*NullTemplate) ProtoMessage()               {}
func (*NullTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *NullTemplate) GetCrashStage() NullTemplate_CrashStage {
	if m != nil {
		return m.CrashStage
	}
	return NullTemplate_NONE
}

func (m *NullTemplate) GetNodeGroups() []string {
	if m != nil {
		return m.NodeGroups
	}
	return nil
}

func init() {
	proto.RegisterType((*Instance)(nil), "model.Instance")
	proto.RegisterType((*InstanceState)(nil), "model.InstanceState")
	proto.RegisterType((*FailureMessage)(nil), "model.FailureMessage")
	proto.RegisterType((*Template)(nil), "model.Template")
	proto.RegisterType((*NoneTemplate)(nil), "model.NoneTemplate")
	proto.RegisterType((*LxcTemplate)(nil), "model.LxcTemplate")
	proto.RegisterType((*LxcTemplate_Image)(nil), "model.LxcTemplate.Image")
	proto.RegisterType((*LxcTemplate_Interface)(nil), "model.LxcTemplate.Interface")
	proto.RegisterType((*LxcTemplate_Template)(nil), "model.LxcTemplate.Template")
	proto.RegisterType((*VmwareTemplate)(nil), "model.VmwareTemplate")
	proto.RegisterType((*VmwareTemplate_Image)(nil), "model.VmwareTemplate.Image")
	proto.RegisterType((*VmwareTemplate_Interface)(nil), "model.VmwareTemplate.Interface")
	proto.RegisterType((*NullTemplate)(nil), "model.NullTemplate")
	proto.RegisterEnum("model.InstanceState_State", InstanceState_State_name, InstanceState_State_value)
	proto.RegisterEnum("model.FailureMessage_ErrorType", FailureMessage_ErrorType_name, FailureMessage_ErrorType_value)
	proto.RegisterEnum("model.NullTemplate_CrashStage", NullTemplate_CrashStage_name, NullTemplate_CrashStage_value)
}

func init() { proto.RegisterFile("model.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xdc, 0x56, 0xdb, 0x8e, 0x1b, 0x45,
	0x10, 0x8d, 0xef, 0x9e, 0xf2, 0x25, 0xa6, 0x09, 0xd1, 0xc8, 0x09, 0xd9, 0x68, 0x84, 0x04, 0x02,
	0xc9, 0x46, 0x10, 0x81, 0x04, 0x08, 0xd8, 0xcd, 0x3a, 0x89, 0xa5, 0xc4, 0x5e, 0xda, 0x5e, 0x10,
	0xbc, 0x8c, 0xda, 0x33, 0xbd, 0xf6, 0x68, 0xe7, 0x62, 0xcd, 0x8c, 0x9d, 0xdd, 0x7f, 0xe0, 0x0d,
	0x29, 0x7f, 0x80, 0xf8, 0x33, 0x7e, 0x00, 0x08, 0xaf, 0x54, 0x5f, 0xc6, 0x1e, 0xef, 0x3a, 0x12,
	0x97, 0x07, 0xa4, 0x3c, 0x4d, 0xd7, 0xe9, 0x53, 0xd5, 0x55, 0xd5, 0xd5, 0x55, 0x03, 0x8d, 0x20,
	0x72, 0xb9, 0xdf, 0x5b, 0xc6, 0x51, 0x1a, 0x91, 0x8a, 0x14, 0xba, 0x9f, 0xcf, 0xbd, 0x74, 0xb1,
	0x9a, 0xf5, 0x9c, 0x28, 0xe8, 0xcf, 0x23, 0x9f, 0x85, 0xf3, 0xbe, 0xdc, 0x9f, 0xad, 0xce, 0xfa,
	0xcb, 0xf4, 0x72, 0xc9, 0x93, 0x7e, 0xea, 0x05, 0x3c, 0x49, 0x59, 0xb0, 0xdc, 0xae, 0x94, 0x0d,
	0xeb, 0xa7, 0x22, 0xd4, 0x87, 0x21, 0x22, 0xa1, 0xc3, 0x49, 0x1b, 0x8a, 0x9e, 0x6b, 0x16, 0xee,
	0x17, 0xde, 0x33, 0x28, 0xae, 0x48, 0x17, 0xea, 0x89, 0xcf, 0xd6, 0xdc, 0x46, 0xb4, 0x28, 0xd1,
	0x8d, 0x4c, 0x1e, 0x00, 0xf8, 0x2c, 0x49, 0x6d, 0x54, 0x4d, 0xb9, 0x59, 0xc6, 0xdd, 0xc6, 0x47,
	0xb7, 0x7a, 0xca, 0xbd, 0xcc, 0xe0, 0x44, 0xec, 0xd1, 0x1c, 0x8f, 0x7c, 0x06, 0xe0, 0xc4, 0x1c,
	0x57, 0xae, 0xcd, 0x52, 0xb3, 0x22, 0xb5, 0xba, 0xbd, 0x79, 0x14, 0xcd, 0x7d, 0xde, 0xcb, 0xbc,
	0xee, 0x4d, 0x33, 0x27, 0x69, 0x8e, 0x4d, 0x3e, 0x80, 0x7a, 0xca, 0x83, 0xa5, 0x2f, 0xce, 0xab,
	0x4a, 0xcd, 0x9b, 0xfa, 0xbc, 0xa9, 0x86, 0xe9, 0x86, 0x40, 0xbe, 0x80, 0xb6, 0xf8, 0xe2, 0xc1,
	0x67, 0xcc, 0xf3, 0x57, 0x31, 0x37, 0x6b, 0x52, 0xe5, 0x2d, 0xad, 0xf2, 0x48, 0xa1, 0xcf, 0x78,
	0x92, 0xb0, 0x39, 0xa7, 0x2d, 0x45, 0xd6, 0xa8, 0xf5, 0xa2, 0x08, 0xad, 0x9d, 0x20, 0xc8, 0x87,
	0x50, 0x51, 0x91, 0x8a, 0xec, 0xb4, 0xd1, 0xe7, 0x3d, 0x91, 0xf6, 0x54, 0xbc, 0x95, 0x7d, 0xa1,
	0x16, 0xff, 0x49, 0xa8, 0xd6, 0x8b, 0x02, 0x54, 0xd4, 0xb9, 0x6d, 0x00, 0x3a, 0x78, 0x3c, 0x9c,
	0x4c, 0x07, 0x74, 0x70, 0xdc, 0xb9, 0x41, 0x00, 0xaa, 0xdf, 0x9c, 0x0e, 0x4e, 0x71, 0x5d, 0x20,
	0x4d, 0xa8, 0x4f, 0xa6, 0x87, 0x74, 0x3a, 0x1c, 0x3d, 0xee, 0x14, 0x49, 0x03, 0x6a, 0xf4, 0x74,
	0x34, 0x12, 0x42, 0x49, 0x6d, 0x8d, 0x4f, 0x4e, 0x84, 0x54, 0x16, 0x5b, 0x52, 0x42, 0xad, 0x0a,
	0x69, 0x81, 0x41, 0x07, 0x47, 0xe3, 0xb1, 0x54, 0xab, 0x92, 0x0e, 0x34, 0x27, 0x4f, 0x4e, 0xa7,
	0x42, 0x3a, 0x1e, 0x7f, 0x37, 0xea, 0xd4, 0xc4, 0x91, 0x78, 0xda, 0xb3, 0xe1, 0xe8, 0x70, 0x8a,
	0x0a, 0x75, 0x71, 0xe4, 0xa3, 0xc3, 0xe1, 0x53, 0x5c, 0x1b, 0xd6, 0x6f, 0x05, 0x68, 0xef, 0xa6,
	0x8e, 0x7c, 0x09, 0xc0, 0xe3, 0x38, 0x8a, 0x6d, 0x51, 0x6a, 0x3a, 0x3d, 0x07, 0x7b, 0xb3, 0xdc,
	0x1b, 0x08, 0xde, 0x14, 0x69, 0xd4, 0xe0, 0xd9, 0x92, 0x7c, 0x0a, 0x86, 0xb8, 0xa2, 0xbf, 0x9b,
	0xa6, 0xba, 0x22, 0x1f, 0xa6, 0xd6, 0x02, 0x8c, 0x8d, 0x41, 0x72, 0x13, 0x1a, 0xca, 0x49, 0x5b,
	0x84, 0x86, 0x89, 0xc2, 0xb8, 0x34, 0x20, 0x73, 0x84, 0xe9, 0xda, 0x52, 0x44, 0x32, 0x30, 0x63,
	0x6f, 0x40, 0x4b, 0x03, 0x2a, 0x21, 0x98, 0xb7, 0x5b, 0xd0, 0xd1, 0xd0, 0x26, 0x05, 0x9d, 0xb2,
	0xf5, 0x33, 0x3e, 0x92, 0xac, 0xc6, 0x88, 0x05, 0xcd, 0xac, 0xca, 0xec, 0x55, 0xec, 0xe9, 0xe7,
	0xb2, 0x83, 0x91, 0xf7, 0xa1, 0x1c, 0x46, 0x21, 0x37, 0x7f, 0x2f, 0xc9, 0x78, 0xde, 0xd4, 0xe9,
	0x18, 0x21, 0x96, 0xd9, 0x79, 0x72, 0x83, 0x4a, 0x0e, 0x79, 0x17, 0x4a, 0xfe, 0x85, 0x63, 0xfe,
	0xa1, 0xa8, 0x44, 0x53, 0x9f, 0x5e, 0x38, 0x39, 0xa6, 0x60, 0x48, 0xa3, 0x2b, 0xdf, 0x37, 0x5f,
	0x5e, 0x31, 0x8a, 0xd8, 0x8e, 0x51, 0x94, 0xb1, 0x5c, 0xab, 0xeb, 0xe0, 0x39, 0xc3, 0xb2, 0xff,
	0xb3, 0xb4, 0x53, 0xf7, 0xdf, 0x4a, 0x34, 0xc7, 0xd7, 0xbc, 0xff, 0x52, 0xae, 0x47, 0x55, 0x28,
	0x0f, 0x31, 0x7e, 0xab, 0x0d, 0xcd, 0x7c, 0x88, 0xd6, 0x2f, 0x35, 0x68, 0xe4, 0x02, 0x21, 0x04,
	0xca, 0x6b, 0x67, 0xb9, 0x92, 0x29, 0xab, 0x50, 0xb9, 0x26, 0x77, 0xc1, 0x08, 0x78, 0x10, 0xc5,
	0x97, 0xf6, 0x7c, 0x26, 0x8f, 0xad, 0xd0, 0x2d, 0x20, 0x3a, 0x50, 0xe0, 0x85, 0xb6, 0xd4, 0x2a,
	0xc9, 0xcd, 0x8d, 0x4c, 0xde, 0x81, 0x96, 0x58, 0x6f, 0xb5, 0xcb, 0x92, 0xb0, 0x0b, 0x92, 0x4f,
	0xc0, 0xc0, 0xe4, 0xd9, 0x5e, 0x80, 0x05, 0xa8, 0x1b, 0x8e, 0x79, 0x3d, 0xc7, 0xbd, 0xa1, 0xd8,
	0xa7, 0x5b, 0x2a, 0x36, 0x10, 0xf0, 0xc2, 0x94, 0xc7, 0x67, 0xcc, 0xe1, 0x09, 0xf6, 0x9b, 0x12,
	0x2a, 0xde, 0xdd, 0xa7, 0x98, 0x91, 0x68, 0x8e, 0x4f, 0xbe, 0x82, 0xa6, 0x30, 0xb5, 0xe9, 0x57,
	0xaa, 0xf9, 0xdc, 0xd9, 0xa3, 0xbf, 0xe9, 0x5d, 0x3b, 0x0a, 0xe4, 0x3e, 0x34, 0x42, 0xe4, 0xda,
	0xf3, 0x38, 0x5a, 0x2d, 0x13, 0xb3, 0x8e, 0xe7, 0x1b, 0x34, 0x0f, 0x75, 0x39, 0x54, 0xa4, 0xd3,
	0xa2, 0x20, 0xdd, 0xe8, 0x79, 0xe8, 0x47, 0xcc, 0xc5, 0xe2, 0xf3, 0xb3, 0x82, 0xcc, 0x63, 0xc2,
	0x9c, 0xb3, 0x38, 0x4f, 0x56, 0x81, 0x7a, 0xa5, 0xaa, 0x99, 0xe7, 0x21, 0x72, 0x1b, 0xaa, 0x4a,
	0x94, 0x79, 0x36, 0xa8, 0x96, 0xba, 0xa7, 0x60, 0x6c, 0x42, 0x14, 0x17, 0xb8, 0x79, 0xe5, 0x06,
	0x95, 0x6b, 0x62, 0x42, 0x2d, 0x60, 0x0e, 0x73, 0xdd, 0x58, 0x9b, 0xcd, 0x44, 0x71, 0x79, 0xde,
	0x72, 0xfd, 0x40, 0x6e, 0x29, 0xa3, 0x1b, 0xb9, 0xfb, 0x63, 0x29, 0xf7, 0xa4, 0xba, 0xb9, 0xce,
	0xae, 0x4c, 0x6f, 0x1b, 0x39, 0x9a, 0x9f, 0x31, 0xe7, 0x9c, 0x87, 0xae, 0xae, 0x8e, 0x4c, 0x14,
	0x1e, 0xbb, 0x5e, 0x92, 0xc6, 0x51, 0xe6, 0xb1, 0x92, 0x84, 0x46, 0xcc, 0x7d, 0xce, 0x12, 0x35,
	0x96, 0xd0, 0x21, 0x2d, 0x0a, 0xf7, 0x59, 0xec, 0x2c, 0x64, 0x19, 0xa0, 0xfb, 0x62, 0x2d, 0xd8,
	0x6b, 0x16, 0x7b, 0x2c, 0x4c, 0xe5, 0x50, 0x41, 0xb6, 0x16, 0x85, 0xfd, 0x84, 0xc7, 0x6b, 0x1e,
	0xcb, 0xdb, 0x43, 0xfb, 0x4a, 0x12, 0xf8, 0x39, 0xbf, 0x14, 0x33, 0xb1, 0xae, 0x70, 0x25, 0x91,
	0x7b, 0x00, 0x62, 0xa5, 0x75, 0x0c, 0xb9, 0x97, 0x43, 0xb0, 0x12, 0x6f, 0xa3, 0x87, 0x6c, 0xe6,
	0xe3, 0x15, 0x2e, 0xe7, 0xf6, 0x9a, 0xf9, 0x9e, 0xcb, 0x52, 0x2f, 0x0a, 0x4d, 0x40, 0x6e, 0x9d,
	0xbe, 0x62, 0x57, 0xdc, 0xdd, 0x99, 0xbf, 0x4a, 0x16, 0xb6, 0xc3, 0x9c, 0x05, 0x37, 0x1b, 0x92,
	0x9c, 0x87, 0x24, 0x23, 0x8a, 0x1d, 0xae, 0x19, 0x4d, 0xcd, 0xd8, 0x42, 0xc2, 0x37, 0x7e, 0x91,
	0xc6, 0xcc, 0x66, 0xf1, 0x3c, 0x31, 0x5b, 0xb2, 0x9a, 0x72, 0x88, 0xf5, 0x6b, 0x09, 0xda, 0xbb,
	0xad, 0xe1, 0x7f, 0x79, 0xac, 0xf8, 0x6c, 0x54, 0x3b, 0xda, 0x79, 0xaf, 0x77, 0xf6, 0xf6, 0x2e,
	0xfd, 0x64, 0x77, 0x14, 0xd0, 0xc0, 0xf5, 0x57, 0x7b, 0xf0, 0x0a, 0xf5, 0xbd, 0x0f, 0xf7, 0xca,
	0xbb, 0xab, 0xbd, 0x66, 0xef, 0xce, 0x7a, 0x59, 0xc0, 0x1e, 0x9d, 0x9b, 0x18, 0xff, 0xe2, 0x9a,
	0xbf, 0xc6, 0x98, 0x62, 0x86, 0xc5, 0x87, 0x83, 0x00, 0xef, 0xa8, 0x24, 0x27, 0xfe, 0xbd, 0x3d,
	0xd3, 0xa8, 0xf7, 0x50, 0xd0, 0x26, 0x82, 0x45, 0xf3, 0x2a, 0x57, 0x93, 0x5c, 0xbe, 0x96, 0x64,
	0xeb, 0x04, 0x60, 0xab, 0x4c, 0xea, 0x50, 0x1e, 0x8d, 0x47, 0x03, 0x9c, 0xea, 0x06, 0xfe, 0x17,
	0xe9, 0x71, 0x8e, 0xa0, 0x9e, 0xe3, 0xf8, 0x83, 0xf2, 0x90, 0x0e, 0xc4, 0xa8, 0x2e, 0x89, 0x5f,
	0x9d, 0xe3, 0xc1, 0x64, 0x4a, 0xc7, 0xdf, 0xe3, 0x7f, 0x0f, 0x6e, 0xe8, 0xc9, 0x5e, 0x39, 0x3a,
	0xf8, 0xe1, 0xed, 0xdc, 0x7f, 0x32, 0xbb, 0x48, 0x16, 0xfd, 0x68, 0xc9, 0xc3, 0xb5, 0xeb, 0xf4,
	0xa5, 0xe7, 0xb3, 0xaa, 0x1c, 0x72, 0x1f, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0x3d, 0xbc, 0x2d,
	0x96, 0x63, 0x0b, 0x00, 0x00,
}
