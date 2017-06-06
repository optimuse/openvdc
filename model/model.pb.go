// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

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
	QemuTemplate
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
	//	*Template_Qemu
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
type Template_Qemu struct {
	Qemu *QemuTemplate `protobuf:"bytes,503,opt,name=qemu,oneof"`
}

func (*Template_None) isTemplate_Item() {}
func (*Template_Lxc) isTemplate_Item()  {}
func (*Template_Null) isTemplate_Item() {}
func (*Template_Qemu) isTemplate_Item() {}

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

func (m *Template) GetQemu() *QemuTemplate {
	if x, ok := m.GetItem().(*Template_Qemu); ok {
		return x.Qemu
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
		(*Template_Qemu)(nil),
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
	case *Template_Qemu:
		b.EncodeVarint(503<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Qemu); err != nil {
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
	case 503: // Item.qemu
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(QemuTemplate)
		err := b.DecodeMessage(msg)
		m.Item = &Template_Qemu{msg}
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
	case *Template_Qemu:
		s := proto.Size(x.Qemu)
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

type QemuTemplate struct {
	Vcpu        int32                     `protobuf:"varint,1,opt,name=vcpu" json:"vcpu,omitempty"`
	MemoryGb    int32                     `protobuf:"varint,2,opt,name=memory_gb" json:"memory_gb,omitempty"`
	MinVcpu     int32                     `protobuf:"varint,3,opt,name=min_vcpu" json:"min_vcpu,omitempty"`
	MinMemoryGb int32                     `protobuf:"varint,4,opt,name=min_memory_gb" json:"min_memory_gb,omitempty"`
	QemuImage   *QemuTemplate_Image       `protobuf:"bytes,5,opt,name=qemu_image" json:"qemu_image,omitempty"`
	Interfaces  []*QemuTemplate_Interface `protobuf:"bytes,6,rep,name=interfaces" json:"interfaces,omitempty"`
	NodeGroups  []string                  `protobuf:"bytes,7,rep,name=node_groups" json:"node_groups,omitempty"`
}

func (m *QemuTemplate) Reset()                    { *m = QemuTemplate{} }
func (m *QemuTemplate) String() string            { return proto.CompactTextString(m) }
func (*QemuTemplate) ProtoMessage()               {}
func (*QemuTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *QemuTemplate) GetVcpu() int32 {
	if m != nil {
		return m.Vcpu
	}
	return 0
}

func (m *QemuTemplate) GetMemoryGb() int32 {
	if m != nil {
		return m.MemoryGb
	}
	return 0
}

func (m *QemuTemplate) GetMinVcpu() int32 {
	if m != nil {
		return m.MinVcpu
	}
	return 0
}

func (m *QemuTemplate) GetMinMemoryGb() int32 {
	if m != nil {
		return m.MinMemoryGb
	}
	return 0
}

func (m *QemuTemplate) GetQemuImage() *QemuTemplate_Image {
	if m != nil {
		return m.QemuImage
	}
	return nil
}

func (m *QemuTemplate) GetInterfaces() []*QemuTemplate_Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *QemuTemplate) GetNodeGroups() []string {
	if m != nil {
		return m.NodeGroups
	}
	return nil
}

type QemuTemplate_Image struct {
	DownloadUrl string `protobuf:"bytes,1,opt,name=download_url" json:"download_url,omitempty"`
	ChksumType  string `protobuf:"bytes,2,opt,name=chksum_type" json:"chksum_type,omitempty"`
	Chksum      string `protobuf:"bytes,3,opt,name=chksum" json:"chksum,omitempty"`
	Format      string `protobuf:"bytes,4,opt,name=format" json:"format,omitempty"`
}

func (m *QemuTemplate_Image) Reset()                    { *m = QemuTemplate_Image{} }
func (m *QemuTemplate_Image) String() string            { return proto.CompactTextString(m) }
func (*QemuTemplate_Image) ProtoMessage()               {}
func (*QemuTemplate_Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *QemuTemplate_Image) GetDownloadUrl() string {
	if m != nil {
		return m.DownloadUrl
	}
	return ""
}

func (m *QemuTemplate_Image) GetChksumType() string {
	if m != nil {
		return m.ChksumType
	}
	return ""
}

func (m *QemuTemplate_Image) GetChksum() string {
	if m != nil {
		return m.Chksum
	}
	return ""
}

func (m *QemuTemplate_Image) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

type QemuTemplate_Interface struct {
	Type     string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Macaddr  string `protobuf:"bytes,2,opt,name=macaddr" json:"macaddr,omitempty"`
	Ipv4Addr string `protobuf:"bytes,3,opt,name=ipv4addr" json:"ipv4addr,omitempty"`
}

func (m *QemuTemplate_Interface) Reset()                    { *m = QemuTemplate_Interface{} }
func (m *QemuTemplate_Interface) String() string            { return proto.CompactTextString(m) }
func (*QemuTemplate_Interface) ProtoMessage()               {}
func (*QemuTemplate_Interface) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 1} }

func (m *QemuTemplate_Interface) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *QemuTemplate_Interface) GetMacaddr() string {
	if m != nil {
		return m.Macaddr
	}
	return ""
}

func (m *QemuTemplate_Interface) GetIpv4Addr() string {
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

func (m *NullTemplate) GetVcpu() int32 {
	if m != nil {
		return m.Vcpu
	}
	return 0
}

func (m *NullTemplate) GetMemoryGb() int32 {
	if m != nil {
		return m.MemoryGb
	}
	return 0
}

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
	proto.RegisterType((*QemuTemplate)(nil), "model.QemuTemplate")
	proto.RegisterType((*QemuTemplate_Image)(nil), "model.QemuTemplate.Image")
	proto.RegisterType((*QemuTemplate_Interface)(nil), "model.QemuTemplate.Interface")
	proto.RegisterType((*NullTemplate)(nil), "model.NullTemplate")
	proto.RegisterEnum("model.InstanceState_State", InstanceState_State_name, InstanceState_State_value)
	proto.RegisterEnum("model.FailureMessage_ErrorType", FailureMessage_ErrorType_name, FailureMessage_ErrorType_value)
	proto.RegisterEnum("model.NullTemplate_CrashStage", NullTemplate_CrashStage_name, NullTemplate_CrashStage_value)
}

func init() { proto.RegisterFile("model.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x56, 0x5b, 0x8f, 0xdb, 0xc4,
	0x17, 0xdf, 0xdc, 0xed, 0x93, 0x4b, 0xfd, 0x9f, 0x7f, 0xa9, 0x4c, 0xe8, 0x25, 0xb2, 0x90, 0x58,
	0x81, 0x94, 0xa0, 0xa5, 0x2a, 0x02, 0xca, 0x65, 0xdb, 0xb8, 0x6d, 0xa4, 0x36, 0x69, 0x27, 0x59,
	0x21, 0x78, 0xb1, 0x26, 0xf6, 0xac, 0x63, 0xad, 0x6f, 0x8c, 0xed, 0xb0, 0xfb, 0xc0, 0x37, 0xe0,
	0x0d, 0xa9, 0xaf, 0x3c, 0xf2, 0x01, 0x29, 0xe5, 0x15, 0xcd, 0x78, 0x9c, 0x38, 0xbb, 0x41, 0xe2,
	0x22, 0xe0, 0x29, 0x73, 0x7e, 0xf3, 0x3b, 0x67, 0xce, 0x39, 0x3e, 0x97, 0x40, 0x3b, 0x88, 0x1c,
	0xea, 0x0f, 0x63, 0x16, 0xa5, 0x11, 0x6a, 0x08, 0xa1, 0xff, 0x89, 0xeb, 0xa5, 0xab, 0x6c, 0x39,
	0xb4, 0xa3, 0x60, 0xe4, 0x46, 0x3e, 0x09, 0xdd, 0x91, 0xb8, 0x5f, 0x66, 0xa7, 0xa3, 0x38, 0xbd,
	0x88, 0x69, 0x32, 0x4a, 0xbd, 0x80, 0x26, 0x29, 0x09, 0xe2, 0xed, 0x29, 0xb7, 0x61, 0xfc, 0x50,
	0x05, 0x65, 0x12, 0x26, 0x29, 0x09, 0x6d, 0x8a, 0x7a, 0x50, 0xf5, 0x1c, 0xbd, 0x32, 0xa8, 0x1c,
	0xaa, 0xb8, 0xea, 0x39, 0xa8, 0x0f, 0x4a, 0xe2, 0x93, 0x35, 0xb5, 0x3c, 0x47, 0xaf, 0x0a, 0x74,
	0x23, 0xa3, 0xbb, 0x00, 0x3e, 0x49, 0x52, 0x2b, 0x49, 0x49, 0x4a, 0xf5, 0xfa, 0xa0, 0x72, 0xd8,
	0x3e, 0xba, 0x3e, 0xcc, 0xdd, 0x2b, 0x0c, 0xce, 0xf9, 0x1d, 0x2e, 0xf1, 0xd0, 0xc7, 0x00, 0x36,
	0xa3, 0x24, 0xa5, 0x8e, 0x45, 0x52, 0xbd, 0x21, 0xb4, 0xfa, 0x43, 0x37, 0x8a, 0x5c, 0x9f, 0x0e,
	0x0b, 0xaf, 0x87, 0x8b, 0xc2, 0x49, 0x5c, 0x62, 0xa3, 0xf7, 0x40, 0x49, 0x69, 0x10, 0xfb, 0xfc,
	0xbd, 0xa6, 0xd0, 0xbc, 0x26, 0xdf, 0x5b, 0x48, 0x18, 0x6f, 0x08, 0xe8, 0x3e, 0xf4, 0xf8, 0x6f,
	0x92, 0x5a, 0xa7, 0xc4, 0xf3, 0x33, 0x46, 0xf5, 0x96, 0x50, 0x79, 0x43, 0xaa, 0x3c, 0xca, 0xd1,
	0x67, 0x34, 0x49, 0x88, 0x4b, 0x71, 0x37, 0x27, 0x4b, 0xd4, 0x78, 0x59, 0x85, 0xee, 0x4e, 0x10,
	0xe8, 0x7d, 0x68, 0xe4, 0x91, 0xf2, 0xec, 0xf4, 0x8e, 0xfa, 0xfb, 0x22, 0x1d, 0xe6, 0xf1, 0x36,
	0xf6, 0x85, 0x5a, 0xfd, 0x33, 0xa1, 0x1a, 0x2f, 0x2b, 0xd0, 0xc8, 0xdf, 0xed, 0x01, 0x60, 0xf3,
	0xf1, 0x64, 0xbe, 0x30, 0xb1, 0x39, 0xd6, 0x0e, 0x10, 0x40, 0xf3, 0xc5, 0x89, 0x79, 0x62, 0x8e,
	0xb5, 0x0a, 0xea, 0x80, 0x32, 0x5f, 0x1c, 0xe3, 0xc5, 0x64, 0xfa, 0x58, 0xab, 0xa2, 0x36, 0xb4,
	0xf0, 0xc9, 0x74, 0xca, 0x85, 0x5a, 0x7e, 0x35, 0x7b, 0xfe, 0x9c, 0x4b, 0x75, 0x7e, 0x25, 0x24,
	0x73, 0xac, 0x35, 0x50, 0x17, 0x54, 0x6c, 0x3e, 0x98, 0xcd, 0x84, 0x5a, 0x13, 0x69, 0xd0, 0x99,
	0x3f, 0x39, 0x59, 0x70, 0x69, 0x3c, 0xfb, 0x72, 0xaa, 0xb5, 0xf8, 0x93, 0x0b, 0x13, 0x3f, 0x9b,
	0x4c, 0x8f, 0x17, 0xe6, 0x58, 0x53, 0xf8, 0x93, 0x8f, 0x8e, 0x27, 0x4f, 0xcd, 0xb1, 0xa6, 0x1a,
	0x3f, 0x57, 0xa0, 0xb7, 0x9b, 0x3a, 0xf4, 0x19, 0x00, 0x65, 0x2c, 0x62, 0x16, 0x2f, 0x35, 0x99,
	0x9e, 0x3b, 0x7b, 0xb3, 0x3c, 0x34, 0x39, 0x6f, 0x71, 0x11, 0x53, 0xac, 0xd2, 0xe2, 0x88, 0x3e,
	0x04, 0x95, 0x7f, 0xa2, 0x3f, 0x9a, 0x26, 0x25, 0x27, 0x1f, 0xa7, 0xc6, 0x0a, 0xd4, 0x8d, 0x41,
	0x74, 0x0d, 0xda, 0xb9, 0x93, 0x16, 0x0f, 0x4d, 0x3b, 0xe0, 0x71, 0x49, 0x40, 0xe4, 0x48, 0xab,
	0x94, 0x28, 0x3c, 0x19, 0x5a, 0x15, 0xfd, 0x0f, 0xba, 0x12, 0xc8, 0x13, 0xa2, 0xd5, 0xd0, 0x75,
	0xd0, 0x24, 0xb4, 0x49, 0x81, 0x56, 0x37, 0x7e, 0xac, 0x82, 0x52, 0xd4, 0x18, 0x32, 0xa0, 0x53,
	0x54, 0x99, 0x95, 0x31, 0x4f, 0xb6, 0xcb, 0x0e, 0x86, 0xde, 0x85, 0x7a, 0x18, 0x85, 0x54, 0x7f,
	0x55, 0x13, 0xf1, 0xfc, 0x5f, 0xa6, 0x63, 0x1a, 0x85, 0xb4, 0xb0, 0xf3, 0xe4, 0x00, 0x0b, 0x0e,
	0x7a, 0x07, 0x6a, 0xfe, 0xb9, 0xad, 0xff, 0x92, 0x53, 0x91, 0xa4, 0x3e, 0x3d, 0xb7, 0x4b, 0x4c,
	0xce, 0x10, 0x46, 0x33, 0xdf, 0xd7, 0x5f, 0x5f, 0x32, 0x9a, 0xf9, 0xfe, 0x8e, 0xd1, 0xcc, 0xf7,
	0x39, 0xf7, 0x1b, 0x1a, 0x64, 0xfa, 0xaf, 0xbb, 0xdc, 0x17, 0x34, 0xc8, 0xca, 0x5c, 0xce, 0xf9,
	0x3b, 0x85, 0xfa, 0xa0, 0x09, 0xf5, 0x49, 0x4a, 0x03, 0xa3, 0x07, 0x9d, 0x72, 0x70, 0xc6, 0x4f,
	0x2d, 0x68, 0x97, 0x42, 0x40, 0x08, 0xea, 0x6b, 0x3b, 0xce, 0x44, 0xb2, 0x1a, 0x58, 0x9c, 0xd1,
	0x4d, 0x50, 0x03, 0x1a, 0x44, 0xec, 0xc2, 0x72, 0x97, 0xe2, 0xd9, 0x06, 0xde, 0x02, 0x7c, 0xf6,
	0x04, 0x5e, 0x68, 0x09, 0xad, 0x9a, 0xb8, 0xdc, 0xc8, 0xe8, 0x6d, 0xe8, 0xf2, 0xf3, 0x56, 0xbb,
	0x2e, 0x08, 0xbb, 0x20, 0xba, 0x07, 0xaa, 0x7f, 0x6e, 0x5b, 0x5e, 0x40, 0x5c, 0x2a, 0x47, 0x8d,
	0x7e, 0x35, 0xbb, 0xc3, 0x09, 0xbf, 0xc7, 0x5b, 0x2a, 0xba, 0x0f, 0xe0, 0x85, 0x29, 0x65, 0xa7,
	0xc4, 0xa6, 0x89, 0xde, 0x1c, 0xd4, 0x0e, 0xdb, 0x47, 0x37, 0xf7, 0x29, 0x16, 0x24, 0x5c, 0xe2,
	0xa3, 0xcf, 0xa1, 0xc3, 0x4d, 0x6d, 0x26, 0x55, 0x3e, 0x76, 0xde, 0xda, 0xa3, 0xbf, 0x99, 0x5a,
	0x3b, 0x0a, 0x68, 0x00, 0xed, 0x30, 0x72, 0xa8, 0xe5, 0xb2, 0x28, 0x8b, 0x13, 0x5d, 0x19, 0xd4,
	0x0e, 0x55, 0x5c, 0x86, 0xfa, 0x14, 0x1a, 0xc2, 0x69, 0x5e, 0x8a, 0x4e, 0xf4, 0x6d, 0xe8, 0x47,
	0xc4, 0xb1, 0x32, 0xe6, 0x17, 0xa5, 0x58, 0xc6, 0xb8, 0x39, 0x7b, 0x75, 0x96, 0x64, 0x41, 0xde,
	0x9f, 0xf9, 0x18, 0x2f, 0x43, 0xe8, 0x06, 0x34, 0x73, 0x51, 0xe4, 0x59, 0xc5, 0x52, 0xea, 0x9f,
	0x80, 0xba, 0x09, 0x91, 0x7f, 0xc0, 0x4d, 0x7f, 0xab, 0x58, 0x9c, 0x91, 0x0e, 0xad, 0x80, 0xd8,
	0xc4, 0x71, 0x98, 0x34, 0x5b, 0x88, 0xfc, 0xe3, 0x79, 0xf1, 0xfa, 0xae, 0xb8, 0xca, 0x8d, 0x6e,
	0xe4, 0xfe, 0xf7, 0xb5, 0x52, 0x33, 0xf5, 0x4b, 0x33, 0x3d, 0x37, 0xbd, 0x1d, 0xe1, 0x3a, 0xb4,
	0x96, 0xc4, 0x3e, 0xa3, 0xa1, 0x23, 0xab, 0xa3, 0x10, 0xb9, 0xc7, 0x8e, 0x97, 0xa4, 0x2c, 0x2a,
	0x3c, 0xce, 0x25, 0xae, 0xc1, 0xa8, 0x4f, 0x49, 0x92, 0x2f, 0x24, 0x15, 0x17, 0x22, 0x77, 0x9f,
	0x30, 0x7b, 0x25, 0xca, 0x40, 0xc5, 0xe2, 0xcc, 0xd9, 0x6b, 0xc2, 0x3c, 0x12, 0xa6, 0x62, 0x9d,
	0xa8, 0xb8, 0x10, 0xb9, 0xfd, 0x84, 0xb2, 0x35, 0x65, 0xe2, 0xeb, 0xa9, 0x58, 0x4a, 0x1c, 0x3f,
	0xa3, 0x17, 0x7c, 0x1b, 0x2a, 0x39, 0x9e, 0x4b, 0xe8, 0x36, 0x00, 0x3f, 0x49, 0x1d, 0x55, 0xdc,
	0x95, 0x10, 0x74, 0x0f, 0x6e, 0x38, 0x5e, 0x42, 0x96, 0x3e, 0xb5, 0xdc, 0xd8, 0xb5, 0xd6, 0xc4,
	0xf7, 0x1c, 0x92, 0x7a, 0x51, 0xa8, 0xc3, 0xa0, 0x72, 0xa8, 0xe0, 0xdf, 0xb9, 0xe5, 0xdf, 0xee,
	0xd4, 0xcf, 0x92, 0x95, 0x65, 0x13, 0x7b, 0x45, 0xf5, 0xb6, 0x20, 0x97, 0x21, 0xc1, 0x88, 0x98,
	0x4d, 0x25, 0xa3, 0x23, 0x19, 0x5b, 0x88, 0xfb, 0x46, 0xcf, 0x53, 0x46, 0x2c, 0xc2, 0xdc, 0x44,
	0xef, 0x8a, 0x6a, 0x2a, 0x21, 0xc6, 0xab, 0x1a, 0x74, 0xca, 0x63, 0xe1, 0x3f, 0x69, 0xd5, 0x8f,
	0x00, 0xf8, 0x28, 0xda, 0xe9, 0xd5, 0x37, 0xf7, 0xcc, 0x2c, 0xd9, 0xac, 0x25, 0x32, 0xfa, 0x74,
	0x4f, 0xb7, 0xde, 0xda, 0xab, 0xba, 0xb7, 0x5d, 0x2f, 0x75, 0x5b, 0xeb, 0x6a, 0xb7, 0x7d, 0xf7,
	0xaf, 0x74, 0x1b, 0xc7, 0x4f, 0x23, 0x16, 0x90, 0x54, 0x96, 0xae, 0x94, 0xfe, 0xa1, 0x2e, 0x34,
	0x5e, 0x57, 0xa0, 0x53, 0xde, 0x1c, 0x7f, 0xe1, 0xb3, 0x7f, 0x01, 0x6d, 0x9b, 0x91, 0x64, 0xc5,
	0xff, 0xda, 0xb9, 0x54, 0xbc, 0xd0, 0x3b, 0xba, 0xbd, 0x67, 0x2b, 0x0d, 0x1f, 0x72, 0xda, 0x9c,
	0xb3, 0x70, 0x59, 0xe5, 0x72, 0xf2, 0xeb, 0x57, 0x92, 0x6f, 0x3c, 0x07, 0xd8, 0x2a, 0x23, 0x05,
	0xea, 0xd3, 0xd9, 0xd4, 0xd4, 0x0e, 0x90, 0x0a, 0x8d, 0x62, 0xad, 0x2b, 0x50, 0x97, 0xfb, 0x1c,
	0xa0, 0xf9, 0x10, 0x9b, 0x7c, 0x65, 0xd7, 0xf8, 0x5f, 0x9e, 0xb1, 0x39, 0x5f, 0xe0, 0xd9, 0x57,
	0x5a, 0x9d, 0x5f, 0xc8, 0x0d, 0xdf, 0x78, 0x70, 0xe7, 0xeb, 0x5b, 0xa5, 0xff, 0xcb, 0xe4, 0x3c,
	0x59, 0x8d, 0xa2, 0x98, 0x86, 0x6b, 0xc7, 0x1e, 0x09, 0xcf, 0x97, 0x4d, 0xb1, 0xf2, 0x3e, 0xf8,
	0x2d, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x20, 0xe4, 0xe7, 0x6b, 0x0b, 0x00, 0x00,
}
