// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: warden/async/v1beta1/plugin.proto

package v1beta1

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Plugin represents an extension for adding new capabilities to the
// blockchain. As part of the x/async module, a plugin executes Tasks offchain,
// asynchronously.
//
// Validators register zero, one, or more, Plugins they're willing to execute.
//
// Plugins are divided in two categories: 1st party plugins that are shipped
// together with the Warden Protocol node binary, and 3rd party plugins
// developed by others.
type Plugin struct {
	// Unique ID of the Plugin.
	//
	// For 1st party plugins, it will just be its name, e.g.:
	// - foo
	// - bar
	// - http
	//
	// For 3rd party plugins, it will be a combination of the address of the creator and its name, e.g.:
	// - 0x4838B108FCe9647Bdf1A7877BF73cE8B0BAD5f97:foo
	// - 0x73f7b9124B2cD361cC0f7654998953E2a251dd58:foo
	// - 0x4838B108FCe9647Bdf1A7877BF73cE8B0BAD5f97:bar
	// - 0x73f7b9124B2cD361cC0f7654998953E2a251dd58:http
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Creator of the plugin. In case of 1st party plugins, this will be empty.
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	// Human-readable description of what this plugin can do.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Fees for creating and executing tasks.
	Fee PluginFee `protobuf:"bytes,4,opt,name=fee,proto3" json:"fee"`
	// Timeout for adding a result to tasks of this plugin.
	// After this timeout elapses, an automated error result will be added on
	// behalf of the solver.
	//
	// The maximum value for this timeout is defined as an x/async module
	// parameter.
	Timeout time.Duration `protobuf:"bytes,5,opt,name=timeout,proto3,stdduration" json:"timeout"`
}

func (m *Plugin) Reset()         { *m = Plugin{} }
func (m *Plugin) String() string { return proto.CompactTextString(m) }
func (*Plugin) ProtoMessage()    {}
func (*Plugin) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa771fd44141666a, []int{0}
}
func (m *Plugin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Plugin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Plugin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Plugin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plugin.Merge(m, src)
}
func (m *Plugin) XXX_Size() int {
	return m.Size()
}
func (m *Plugin) XXX_DiscardUnknown() {
	xxx_messageInfo_Plugin.DiscardUnknown(m)
}

var xxx_messageInfo_Plugin proto.InternalMessageInfo

func (m *Plugin) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Plugin) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Plugin) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Plugin) GetFee() PluginFee {
	if m != nil {
		return m.Fee
	}
	return PluginFee{}
}

func (m *Plugin) GetTimeout() time.Duration {
	if m != nil {
		return m.Timeout
	}
	return 0
}

// Fees for creating and executing concrete task.
type PluginFee struct {
	// The percentage of the fee that goes to the creator of the plugin. The rest goes to the executor of the task.
	// Expressed as a number in the range [0, 1].
	PluginCreatorRewardInPercent cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=plugin_creator_reward_in_percent,json=pluginCreatorRewardInPercent,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"plugin_creator_reward_in_percent"`
	// Fee for creating and executing new Task.
	Fee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=fee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee"`
}

func (m *PluginFee) Reset()         { *m = PluginFee{} }
func (m *PluginFee) String() string { return proto.CompactTextString(m) }
func (*PluginFee) ProtoMessage()    {}
func (*PluginFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa771fd44141666a, []int{1}
}
func (m *PluginFee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PluginFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PluginFee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PluginFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginFee.Merge(m, src)
}
func (m *PluginFee) XXX_Size() int {
	return m.Size()
}
func (m *PluginFee) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginFee.DiscardUnknown(m)
}

var xxx_messageInfo_PluginFee proto.InternalMessageInfo

func (m *PluginFee) GetFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Fee
	}
	return nil
}

func init() {
	proto.RegisterType((*Plugin)(nil), "warden.async.v1beta1.Plugin")
	proto.RegisterType((*PluginFee)(nil), "warden.async.v1beta1.PluginFee")
}

func init() { proto.RegisterFile("warden/async/v1beta1/plugin.proto", fileDescriptor_fa771fd44141666a) }

var fileDescriptor_fa771fd44141666a = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x73, 0x4d, 0x69, 0xd5, 0x2b, 0x20, 0x61, 0x65, 0x70, 0x03, 0xb2, 0x4d, 0xa7, 0xa8,
	0x52, 0xee, 0x94, 0x32, 0x30, 0xe3, 0x46, 0x91, 0x90, 0x18, 0x2a, 0xb3, 0xb1, 0x18, 0xfb, 0x7c,
	0x75, 0x4f, 0x8d, 0x7d, 0xd6, 0xdd, 0x19, 0xc8, 0x7f, 0xc1, 0x88, 0xf8, 0x0b, 0x10, 0x53, 0x87,
	0xfe, 0x11, 0x1d, 0xab, 0x4e, 0x88, 0xa1, 0x45, 0xc9, 0xd0, 0x81, 0x3f, 0x81, 0x05, 0xdd, 0x0f,
	0x47, 0x15, 0x62, 0x49, 0xfc, 0xee, 0xde, 0xf7, 0xf3, 0xde, 0xfb, 0xbe, 0x83, 0xcf, 0x3f, 0x66,
	0xa2, 0xa0, 0x35, 0xce, 0xe4, 0xa2, 0x26, 0xf8, 0xc3, 0x24, 0xa7, 0x2a, 0x9b, 0xe0, 0x66, 0xde,
	0x96, 0xac, 0x46, 0x8d, 0xe0, 0x8a, 0x7b, 0x03, 0x9b, 0x82, 0x4c, 0x0a, 0x72, 0x29, 0xc3, 0x27,
	0x59, 0xc5, 0x6a, 0x8e, 0xcd, 0xaf, 0x4d, 0x1c, 0x06, 0x84, 0xcb, 0x8a, 0x4b, 0x9c, 0x67, 0x92,
	0xae, 0x51, 0x84, 0x77, 0xa0, 0xe1, 0x9e, 0xbd, 0x4f, 0x4d, 0x84, 0x6d, 0xe0, 0xae, 0x06, 0x25,
	0x2f, 0xb9, 0x3d, 0xd7, 0x5f, 0x1d, 0xb0, 0xe4, 0xbc, 0x9c, 0x53, 0x6c, 0xa2, 0xbc, 0x3d, 0xc1,
	0x45, 0x2b, 0x32, 0xc5, 0xb8, 0x03, 0xee, 0xff, 0x06, 0x70, 0xeb, 0xd8, 0xb4, 0xea, 0x3d, 0x86,
	0x1b, 0xac, 0xf0, 0x41, 0x04, 0x46, 0x3b, 0xc9, 0x06, 0x2b, 0xbc, 0x43, 0xb8, 0x4d, 0x04, 0xcd,
	0x14, 0x17, 0xfe, 0x86, 0x3e, 0x8c, 0xfd, 0xeb, 0x8b, 0xf1, 0xc0, 0xd5, 0x7c, 0x55, 0x14, 0x82,
	0x4a, 0xf9, 0x56, 0x09, 0x56, 0x97, 0x49, 0x97, 0xe8, 0x45, 0x70, 0xb7, 0xa0, 0x92, 0x08, 0xd6,
	0xe8, 0x1a, 0x7e, 0xdf, 0xc0, 0xee, 0x1f, 0x79, 0x2f, 0x61, 0xff, 0x84, 0x52, 0x7f, 0x33, 0x02,
	0xa3, 0xdd, 0xc3, 0x10, 0xfd, 0xcf, 0x18, 0x64, 0x1b, 0x9a, 0x51, 0x1a, 0x6f, 0x5e, 0xde, 0x84,
	0xbd, 0x44, 0x2b, 0xbc, 0x18, 0x6e, 0x2b, 0x56, 0x51, 0xde, 0x2a, 0xff, 0x81, 0x11, 0xef, 0x21,
	0x3b, 0x1b, 0xea, 0x66, 0x43, 0x53, 0x37, 0x5b, 0xfc, 0x48, 0xcb, 0xbe, 0xdc, 0x86, 0xe0, 0xdb,
	0xdd, 0xf9, 0x01, 0x48, 0x3a, 0xe1, 0xfe, 0x1f, 0x00, 0x77, 0xd6, 0x70, 0x6f, 0x01, 0x23, 0xbb,
	0xa5, 0xd4, 0xb5, 0x9f, 0x0a, 0xaa, 0xfb, 0x49, 0x59, 0x9d, 0x36, 0x54, 0x10, 0x5a, 0x2b, 0x6b,
	0x47, 0x3c, 0xd1, 0xbc, 0x9f, 0x37, 0xe1, 0x53, 0x3b, 0xbd, 0x2c, 0xce, 0x10, 0xe3, 0xb8, 0xca,
	0xd4, 0x29, 0x7a, 0x43, 0xcb, 0x8c, 0x2c, 0xa6, 0x94, 0x5c, 0x5f, 0x8c, 0xa1, 0x33, 0x67, 0x4a,
	0x49, 0xf2, 0xcc, 0xa2, 0x8f, 0x2c, 0x39, 0x31, 0xe0, 0xd7, 0xf5, 0xb1, 0xc5, 0x7a, 0xd2, 0xba,
	0xd0, 0x8f, 0xfa, 0x66, 0x10, 0xa7, 0xd3, 0x5b, 0x5f, 0x9b, 0x70, 0xc4, 0x59, 0x1d, 0xcf, 0x74,
	0xe1, 0xef, 0xb7, 0xe1, 0xa8, 0x64, 0xea, 0xb4, 0xcd, 0x11, 0xe1, 0x95, 0xdb, 0xba, 0xfb, 0x1b,
	0xcb, 0xe2, 0x0c, 0xab, 0x45, 0x43, 0xa5, 0x11, 0xc8, 0xaf, 0x77, 0xe7, 0x07, 0x0f, 0xe7, 0xa6,
	0xa7, 0x54, 0xbf, 0x1b, 0x69, 0x1d, 0xd0, 0xd5, 0xe2, 0xf7, 0x97, 0xcb, 0x00, 0x5c, 0x2d, 0x03,
	0xf0, 0x6b, 0x19, 0x80, 0xcf, 0xab, 0xa0, 0x77, 0xb5, 0x0a, 0x7a, 0x3f, 0x56, 0x41, 0xef, 0xdd,
	0xec, 0x1e, 0xde, 0x6e, 0x64, 0x6c, 0x4c, 0x25, 0x7c, 0xee, 0xe2, 0x7f, 0x42, 0xfc, 0xc9, 0x3d,
	0x77, 0x53, 0xba, 0x7b, 0xa9, 0xf9, 0x96, 0x49, 0x7b, 0xf1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x98,
	0xde, 0x79, 0xaa, 0x13, 0x03, 0x00, 0x00,
}

func (m *Plugin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Plugin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Plugin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.Timeout, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Timeout):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPlugin(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPlugin(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintPlugin(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintPlugin(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintPlugin(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PluginFee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PluginFee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PluginFee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fee) > 0 {
		for iNdEx := len(m.Fee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPlugin(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size := m.PluginCreatorRewardInPercent.Size()
		i -= size
		if _, err := m.PluginCreatorRewardInPercent.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPlugin(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintPlugin(dAtA []byte, offset int, v uint64) int {
	offset -= sovPlugin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Plugin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPlugin(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovPlugin(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovPlugin(uint64(l))
	}
	l = m.Fee.Size()
	n += 1 + l + sovPlugin(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Timeout)
	n += 1 + l + sovPlugin(uint64(l))
	return n
}

func (m *PluginFee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.PluginCreatorRewardInPercent.Size()
	n += 1 + l + sovPlugin(uint64(l))
	if len(m.Fee) > 0 {
		for _, e := range m.Fee {
			l = e.Size()
			n += 1 + l + sovPlugin(uint64(l))
		}
	}
	return n
}

func sovPlugin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPlugin(x uint64) (n int) {
	return sovPlugin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Plugin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlugin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Plugin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Plugin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlugin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlugin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlugin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlugin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlugin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.Timeout, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlugin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PluginFee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlugin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PluginFee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PluginFee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PluginCreatorRewardInPercent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPlugin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PluginCreatorRewardInPercent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPlugin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = append(m.Fee, types.Coin{})
			if err := m.Fee[len(m.Fee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPlugin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPlugin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlugin
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPlugin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPlugin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPlugin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPlugin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlugin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPlugin = fmt.Errorf("proto: unexpected end of group")
)
