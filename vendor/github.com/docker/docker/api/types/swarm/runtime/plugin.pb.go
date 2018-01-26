// Code generated by protoc-gen-gogo.
// source: plugin.proto
// DO NOT EDIT!

/*
	Package runtime is a generated protocol buffer package.

	It is generated from these files:
		plugin.proto

	It has these top-level messages:
		PluginSpec
		PluginPrivilege
*/
package runtime

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// PluginSpec defines the base payload which clients can specify for creating
// a service with the plugin runtime.
type PluginSpec struct ***REMOVED***
	Name       string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remote     string             `protobuf:"bytes,2,opt,name=remote,proto3" json:"remote,omitempty"`
	Privileges []*PluginPrivilege `protobuf:"bytes,3,rep,name=privileges" json:"privileges,omitempty"`
	Disabled   bool               `protobuf:"varint,4,opt,name=disabled,proto3" json:"disabled,omitempty"`
***REMOVED***

func (m *PluginSpec) Reset()                    ***REMOVED*** *m = PluginSpec***REMOVED******REMOVED*** ***REMOVED***
func (m *PluginSpec) String() string            ***REMOVED*** return proto.CompactTextString(m) ***REMOVED***
func (*PluginSpec) ProtoMessage()               ***REMOVED******REMOVED***
func (*PluginSpec) Descriptor() ([]byte, []int) ***REMOVED*** return fileDescriptorPlugin, []int***REMOVED***0***REMOVED*** ***REMOVED***

func (m *PluginSpec) GetName() string ***REMOVED***
	if m != nil ***REMOVED***
		return m.Name
	***REMOVED***
	return ""
***REMOVED***

func (m *PluginSpec) GetRemote() string ***REMOVED***
	if m != nil ***REMOVED***
		return m.Remote
	***REMOVED***
	return ""
***REMOVED***

func (m *PluginSpec) GetPrivileges() []*PluginPrivilege ***REMOVED***
	if m != nil ***REMOVED***
		return m.Privileges
	***REMOVED***
	return nil
***REMOVED***

func (m *PluginSpec) GetDisabled() bool ***REMOVED***
	if m != nil ***REMOVED***
		return m.Disabled
	***REMOVED***
	return false
***REMOVED***

// PluginPrivilege describes a permission the user has to accept
// upon installing a plugin.
type PluginPrivilege struct ***REMOVED***
	Name        string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Value       []string `protobuf:"bytes,3,rep,name=value" json:"value,omitempty"`
***REMOVED***

func (m *PluginPrivilege) Reset()                    ***REMOVED*** *m = PluginPrivilege***REMOVED******REMOVED*** ***REMOVED***
func (m *PluginPrivilege) String() string            ***REMOVED*** return proto.CompactTextString(m) ***REMOVED***
func (*PluginPrivilege) ProtoMessage()               ***REMOVED******REMOVED***
func (*PluginPrivilege) Descriptor() ([]byte, []int) ***REMOVED*** return fileDescriptorPlugin, []int***REMOVED***1***REMOVED*** ***REMOVED***

func (m *PluginPrivilege) GetName() string ***REMOVED***
	if m != nil ***REMOVED***
		return m.Name
	***REMOVED***
	return ""
***REMOVED***

func (m *PluginPrivilege) GetDescription() string ***REMOVED***
	if m != nil ***REMOVED***
		return m.Description
	***REMOVED***
	return ""
***REMOVED***

func (m *PluginPrivilege) GetValue() []string ***REMOVED***
	if m != nil ***REMOVED***
		return m.Value
	***REMOVED***
	return nil
***REMOVED***

func init() ***REMOVED***
	proto.RegisterType((*PluginSpec)(nil), "PluginSpec")
	proto.RegisterType((*PluginPrivilege)(nil), "PluginPrivilege")
***REMOVED***
func (m *PluginSpec) Marshal() (dAtA []byte, err error) ***REMOVED***
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return dAtA[:n], nil
***REMOVED***

func (m *PluginSpec) MarshalTo(dAtA []byte) (int, error) ***REMOVED***
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 ***REMOVED***
		dAtA[i] = 0xa
		i++
		i = encodeVarintPlugin(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	***REMOVED***
	if len(m.Remote) > 0 ***REMOVED***
		dAtA[i] = 0x12
		i++
		i = encodeVarintPlugin(dAtA, i, uint64(len(m.Remote)))
		i += copy(dAtA[i:], m.Remote)
	***REMOVED***
	if len(m.Privileges) > 0 ***REMOVED***
		for _, msg := range m.Privileges ***REMOVED***
			dAtA[i] = 0x1a
			i++
			i = encodeVarintPlugin(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil ***REMOVED***
				return 0, err
			***REMOVED***
			i += n
		***REMOVED***
	***REMOVED***
	if m.Disabled ***REMOVED***
		dAtA[i] = 0x20
		i++
		if m.Disabled ***REMOVED***
			dAtA[i] = 1
		***REMOVED*** else ***REMOVED***
			dAtA[i] = 0
		***REMOVED***
		i++
	***REMOVED***
	return i, nil
***REMOVED***

func (m *PluginPrivilege) Marshal() (dAtA []byte, err error) ***REMOVED***
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return dAtA[:n], nil
***REMOVED***

func (m *PluginPrivilege) MarshalTo(dAtA []byte) (int, error) ***REMOVED***
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 ***REMOVED***
		dAtA[i] = 0xa
		i++
		i = encodeVarintPlugin(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	***REMOVED***
	if len(m.Description) > 0 ***REMOVED***
		dAtA[i] = 0x12
		i++
		i = encodeVarintPlugin(dAtA, i, uint64(len(m.Description)))
		i += copy(dAtA[i:], m.Description)
	***REMOVED***
	if len(m.Value) > 0 ***REMOVED***
		for _, s := range m.Value ***REMOVED***
			dAtA[i] = 0x1a
			i++
			l = len(s)
			for l >= 1<<7 ***REMOVED***
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			***REMOVED***
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		***REMOVED***
	***REMOVED***
	return i, nil
***REMOVED***

func encodeFixed64Plugin(dAtA []byte, offset int, v uint64) int ***REMOVED***
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
***REMOVED***
func encodeFixed32Plugin(dAtA []byte, offset int, v uint32) int ***REMOVED***
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
***REMOVED***
func encodeVarintPlugin(dAtA []byte, offset int, v uint64) int ***REMOVED***
	for v >= 1<<7 ***REMOVED***
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	***REMOVED***
	dAtA[offset] = uint8(v)
	return offset + 1
***REMOVED***
func (m *PluginSpec) Size() (n int) ***REMOVED***
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 ***REMOVED***
		n += 1 + l + sovPlugin(uint64(l))
	***REMOVED***
	l = len(m.Remote)
	if l > 0 ***REMOVED***
		n += 1 + l + sovPlugin(uint64(l))
	***REMOVED***
	if len(m.Privileges) > 0 ***REMOVED***
		for _, e := range m.Privileges ***REMOVED***
			l = e.Size()
			n += 1 + l + sovPlugin(uint64(l))
		***REMOVED***
	***REMOVED***
	if m.Disabled ***REMOVED***
		n += 2
	***REMOVED***
	return n
***REMOVED***

func (m *PluginPrivilege) Size() (n int) ***REMOVED***
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 ***REMOVED***
		n += 1 + l + sovPlugin(uint64(l))
	***REMOVED***
	l = len(m.Description)
	if l > 0 ***REMOVED***
		n += 1 + l + sovPlugin(uint64(l))
	***REMOVED***
	if len(m.Value) > 0 ***REMOVED***
		for _, s := range m.Value ***REMOVED***
			l = len(s)
			n += 1 + l + sovPlugin(uint64(l))
		***REMOVED***
	***REMOVED***
	return n
***REMOVED***

func sovPlugin(x uint64) (n int) ***REMOVED***
	for ***REMOVED***
		n++
		x >>= 7
		if x == 0 ***REMOVED***
			break
		***REMOVED***
	***REMOVED***
	return n
***REMOVED***
func sozPlugin(x uint64) (n int) ***REMOVED***
	return sovPlugin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
***REMOVED***
func (m *PluginSpec) Unmarshal(dAtA []byte) error ***REMOVED***
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l ***REMOVED***
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 ***REMOVED***
			if shift >= 64 ***REMOVED***
				return ErrIntOverflowPlugin
			***REMOVED***
			if iNdEx >= l ***REMOVED***
				return io.ErrUnexpectedEOF
			***REMOVED***
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 ***REMOVED***
				break
			***REMOVED***
		***REMOVED***
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 ***REMOVED***
			return fmt.Errorf("proto: PluginSpec: wiretype end group for non-group")
		***REMOVED***
		if fieldNum <= 0 ***REMOVED***
			return fmt.Errorf("proto: PluginSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		***REMOVED***
		switch fieldNum ***REMOVED***
		case 1:
			if wireType != 2 ***REMOVED***
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			***REMOVED***
			var stringLen uint64
			for shift := uint(0); ; shift += 7 ***REMOVED***
				if shift >= 64 ***REMOVED***
					return ErrIntOverflowPlugin
				***REMOVED***
				if iNdEx >= l ***REMOVED***
					return io.ErrUnexpectedEOF
				***REMOVED***
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 ***REMOVED***
					break
				***REMOVED***
			***REMOVED***
			intStringLen := int(stringLen)
			if intStringLen < 0 ***REMOVED***
				return ErrInvalidLengthPlugin
			***REMOVED***
			postIndex := iNdEx + intStringLen
			if postIndex > l ***REMOVED***
				return io.ErrUnexpectedEOF
			***REMOVED***
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 ***REMOVED***
				return fmt.Errorf("proto: wrong wireType = %d for field Remote", wireType)
			***REMOVED***
			var stringLen uint64
			for shift := uint(0); ; shift += 7 ***REMOVED***
				if shift >= 64 ***REMOVED***
					return ErrIntOverflowPlugin
				***REMOVED***
				if iNdEx >= l ***REMOVED***
					return io.ErrUnexpectedEOF
				***REMOVED***
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 ***REMOVED***
					break
				***REMOVED***
			***REMOVED***
			intStringLen := int(stringLen)
			if intStringLen < 0 ***REMOVED***
				return ErrInvalidLengthPlugin
			***REMOVED***
			postIndex := iNdEx + intStringLen
			if postIndex > l ***REMOVED***
				return io.ErrUnexpectedEOF
			***REMOVED***
			m.Remote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 ***REMOVED***
				return fmt.Errorf("proto: wrong wireType = %d for field Privileges", wireType)
			***REMOVED***
			var msglen int
			for shift := uint(0); ; shift += 7 ***REMOVED***
				if shift >= 64 ***REMOVED***
					return ErrIntOverflowPlugin
				***REMOVED***
				if iNdEx >= l ***REMOVED***
					return io.ErrUnexpectedEOF
				***REMOVED***
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 ***REMOVED***
					break
				***REMOVED***
			***REMOVED***
			if msglen < 0 ***REMOVED***
				return ErrInvalidLengthPlugin
			***REMOVED***
			postIndex := iNdEx + msglen
			if postIndex > l ***REMOVED***
				return io.ErrUnexpectedEOF
			***REMOVED***
			m.Privileges = append(m.Privileges, &PluginPrivilege***REMOVED******REMOVED***)
			if err := m.Privileges[len(m.Privileges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil ***REMOVED***
				return err
			***REMOVED***
			iNdEx = postIndex
		case 4:
			if wireType != 0 ***REMOVED***
				return fmt.Errorf("proto: wrong wireType = %d for field Disabled", wireType)
			***REMOVED***
			var v int
			for shift := uint(0); ; shift += 7 ***REMOVED***
				if shift >= 64 ***REMOVED***
					return ErrIntOverflowPlugin
				***REMOVED***
				if iNdEx >= l ***REMOVED***
					return io.ErrUnexpectedEOF
				***REMOVED***
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 ***REMOVED***
					break
				***REMOVED***
			***REMOVED***
			m.Disabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil ***REMOVED***
				return err
			***REMOVED***
			if skippy < 0 ***REMOVED***
				return ErrInvalidLengthPlugin
			***REMOVED***
			if (iNdEx + skippy) > l ***REMOVED***
				return io.ErrUnexpectedEOF
			***REMOVED***
			iNdEx += skippy
		***REMOVED***
	***REMOVED***

	if iNdEx > l ***REMOVED***
		return io.ErrUnexpectedEOF
	***REMOVED***
	return nil
***REMOVED***
func (m *PluginPrivilege) Unmarshal(dAtA []byte) error ***REMOVED***
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l ***REMOVED***
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 ***REMOVED***
			if shift >= 64 ***REMOVED***
				return ErrIntOverflowPlugin
			***REMOVED***
			if iNdEx >= l ***REMOVED***
				return io.ErrUnexpectedEOF
			***REMOVED***
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 ***REMOVED***
				break
			***REMOVED***
		***REMOVED***
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 ***REMOVED***
			return fmt.Errorf("proto: PluginPrivilege: wiretype end group for non-group")
		***REMOVED***
		if fieldNum <= 0 ***REMOVED***
			return fmt.Errorf("proto: PluginPrivilege: illegal tag %d (wire type %d)", fieldNum, wire)
		***REMOVED***
		switch fieldNum ***REMOVED***
		case 1:
			if wireType != 2 ***REMOVED***
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			***REMOVED***
			var stringLen uint64
			for shift := uint(0); ; shift += 7 ***REMOVED***
				if shift >= 64 ***REMOVED***
					return ErrIntOverflowPlugin
				***REMOVED***
				if iNdEx >= l ***REMOVED***
					return io.ErrUnexpectedEOF
				***REMOVED***
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 ***REMOVED***
					break
				***REMOVED***
			***REMOVED***
			intStringLen := int(stringLen)
			if intStringLen < 0 ***REMOVED***
				return ErrInvalidLengthPlugin
			***REMOVED***
			postIndex := iNdEx + intStringLen
			if postIndex > l ***REMOVED***
				return io.ErrUnexpectedEOF
			***REMOVED***
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 ***REMOVED***
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			***REMOVED***
			var stringLen uint64
			for shift := uint(0); ; shift += 7 ***REMOVED***
				if shift >= 64 ***REMOVED***
					return ErrIntOverflowPlugin
				***REMOVED***
				if iNdEx >= l ***REMOVED***
					return io.ErrUnexpectedEOF
				***REMOVED***
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 ***REMOVED***
					break
				***REMOVED***
			***REMOVED***
			intStringLen := int(stringLen)
			if intStringLen < 0 ***REMOVED***
				return ErrInvalidLengthPlugin
			***REMOVED***
			postIndex := iNdEx + intStringLen
			if postIndex > l ***REMOVED***
				return io.ErrUnexpectedEOF
			***REMOVED***
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 ***REMOVED***
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			***REMOVED***
			var stringLen uint64
			for shift := uint(0); ; shift += 7 ***REMOVED***
				if shift >= 64 ***REMOVED***
					return ErrIntOverflowPlugin
				***REMOVED***
				if iNdEx >= l ***REMOVED***
					return io.ErrUnexpectedEOF
				***REMOVED***
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 ***REMOVED***
					break
				***REMOVED***
			***REMOVED***
			intStringLen := int(stringLen)
			if intStringLen < 0 ***REMOVED***
				return ErrInvalidLengthPlugin
			***REMOVED***
			postIndex := iNdEx + intStringLen
			if postIndex > l ***REMOVED***
				return io.ErrUnexpectedEOF
			***REMOVED***
			m.Value = append(m.Value, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil ***REMOVED***
				return err
			***REMOVED***
			if skippy < 0 ***REMOVED***
				return ErrInvalidLengthPlugin
			***REMOVED***
			if (iNdEx + skippy) > l ***REMOVED***
				return io.ErrUnexpectedEOF
			***REMOVED***
			iNdEx += skippy
		***REMOVED***
	***REMOVED***

	if iNdEx > l ***REMOVED***
		return io.ErrUnexpectedEOF
	***REMOVED***
	return nil
***REMOVED***
func skipPlugin(dAtA []byte) (n int, err error) ***REMOVED***
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l ***REMOVED***
		var wire uint64
		for shift := uint(0); ; shift += 7 ***REMOVED***
			if shift >= 64 ***REMOVED***
				return 0, ErrIntOverflowPlugin
			***REMOVED***
			if iNdEx >= l ***REMOVED***
				return 0, io.ErrUnexpectedEOF
			***REMOVED***
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 ***REMOVED***
				break
			***REMOVED***
		***REMOVED***
		wireType := int(wire & 0x7)
		switch wireType ***REMOVED***
		case 0:
			for shift := uint(0); ; shift += 7 ***REMOVED***
				if shift >= 64 ***REMOVED***
					return 0, ErrIntOverflowPlugin
				***REMOVED***
				if iNdEx >= l ***REMOVED***
					return 0, io.ErrUnexpectedEOF
				***REMOVED***
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 ***REMOVED***
					break
				***REMOVED***
			***REMOVED***
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 ***REMOVED***
				if shift >= 64 ***REMOVED***
					return 0, ErrIntOverflowPlugin
				***REMOVED***
				if iNdEx >= l ***REMOVED***
					return 0, io.ErrUnexpectedEOF
				***REMOVED***
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 ***REMOVED***
					break
				***REMOVED***
			***REMOVED***
			iNdEx += length
			if length < 0 ***REMOVED***
				return 0, ErrInvalidLengthPlugin
			***REMOVED***
			return iNdEx, nil
		case 3:
			for ***REMOVED***
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 ***REMOVED***
					if shift >= 64 ***REMOVED***
						return 0, ErrIntOverflowPlugin
					***REMOVED***
					if iNdEx >= l ***REMOVED***
						return 0, io.ErrUnexpectedEOF
					***REMOVED***
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 ***REMOVED***
						break
					***REMOVED***
				***REMOVED***
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 ***REMOVED***
					break
				***REMOVED***
				next, err := skipPlugin(dAtA[start:])
				if err != nil ***REMOVED***
					return 0, err
				***REMOVED***
				iNdEx = start + next
			***REMOVED***
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		***REMOVED***
	***REMOVED***
	panic("unreachable")
***REMOVED***

var (
	ErrInvalidLengthPlugin = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlugin   = fmt.Errorf("proto: integer overflow")
)

func init() ***REMOVED*** proto.RegisterFile("plugin.proto", fileDescriptorPlugin) ***REMOVED***

var fileDescriptorPlugin = []byte***REMOVED***
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xc8, 0x29, 0x4d,
	0xcf, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x6a, 0x63, 0xe4, 0xe2, 0x0a, 0x00, 0x0b,
	0x04, 0x17, 0xa4, 0x26, 0x0b, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x62, 0x5c, 0x6c, 0x45, 0xa9, 0xb9, 0xf9, 0x25, 0xa9, 0x12,
	0x4c, 0x60, 0x51, 0x28, 0x4f, 0xc8, 0x80, 0x8b, 0xab, 0xa0, 0x28, 0xb3, 0x2c, 0x33, 0x27, 0x35,
	0x3d, 0xb5, 0x58, 0x82, 0x59, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x40, 0x0f, 0x62, 0x58, 0x00, 0x4c,
	0x22, 0x08, 0x49, 0x8d, 0x90, 0x14, 0x17, 0x47, 0x4a, 0x66, 0x71, 0x62, 0x52, 0x4e, 0x6a, 0x8a,
	0x04, 0x8b, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x9c, 0xaf, 0x14, 0xcb, 0xc5, 0x8f, 0xa6, 0x15, 0xab,
	0x63, 0x14, 0xb8, 0xb8, 0x53, 0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0x32, 0xf3, 0xf3, 0xa0,
	0x2e, 0x42, 0x16, 0x12, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x05, 0xbb, 0x88, 0x33,
	0x08, 0xc2, 0x71, 0xe2, 0x39, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4,
	0x18, 0x93, 0xd8, 0xc0, 0x9e, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x84, 0xad, 0x79,
	0x0c, 0x01, 0x00, 0x00,
***REMOVED***
