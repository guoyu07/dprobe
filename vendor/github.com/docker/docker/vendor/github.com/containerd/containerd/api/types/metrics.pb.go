// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/containerd/containerd/api/types/metrics.proto

package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/gogo/protobuf/types"
import _ "github.com/gogo/protobuf/types"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

type Metric struct ***REMOVED***
	Timestamp time.Time             `protobuf:"bytes,1,opt,name=timestamp,stdtime" json:"timestamp"`
	ID        string                `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Data      *google_protobuf1.Any `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
***REMOVED***

func (m *Metric) Reset()                    ***REMOVED*** *m = Metric***REMOVED******REMOVED*** ***REMOVED***
func (*Metric) ProtoMessage()               ***REMOVED******REMOVED***
func (*Metric) Descriptor() ([]byte, []int) ***REMOVED*** return fileDescriptorMetrics, []int***REMOVED***0***REMOVED*** ***REMOVED***

func init() ***REMOVED***
	proto.RegisterType((*Metric)(nil), "containerd.types.Metric")
***REMOVED***
func (m *Metric) Marshal() (dAtA []byte, err error) ***REMOVED***
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return dAtA[:n], nil
***REMOVED***

func (m *Metric) MarshalTo(dAtA []byte) (int, error) ***REMOVED***
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMetrics(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)))
	n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i:])
	if err != nil ***REMOVED***
		return 0, err
	***REMOVED***
	i += n1
	if len(m.ID) > 0 ***REMOVED***
		dAtA[i] = 0x12
		i++
		i = encodeVarintMetrics(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	***REMOVED***
	if m.Data != nil ***REMOVED***
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMetrics(dAtA, i, uint64(m.Data.Size()))
		n2, err := m.Data.MarshalTo(dAtA[i:])
		if err != nil ***REMOVED***
			return 0, err
		***REMOVED***
		i += n2
	***REMOVED***
	return i, nil
***REMOVED***

func encodeVarintMetrics(dAtA []byte, offset int, v uint64) int ***REMOVED***
	for v >= 1<<7 ***REMOVED***
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	***REMOVED***
	dAtA[offset] = uint8(v)
	return offset + 1
***REMOVED***
func (m *Metric) Size() (n int) ***REMOVED***
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovMetrics(uint64(l))
	l = len(m.ID)
	if l > 0 ***REMOVED***
		n += 1 + l + sovMetrics(uint64(l))
	***REMOVED***
	if m.Data != nil ***REMOVED***
		l = m.Data.Size()
		n += 1 + l + sovMetrics(uint64(l))
	***REMOVED***
	return n
***REMOVED***

func sovMetrics(x uint64) (n int) ***REMOVED***
	for ***REMOVED***
		n++
		x >>= 7
		if x == 0 ***REMOVED***
			break
		***REMOVED***
	***REMOVED***
	return n
***REMOVED***
func sozMetrics(x uint64) (n int) ***REMOVED***
	return sovMetrics(uint64((x << 1) ^ uint64((int64(x) >> 63))))
***REMOVED***
func (this *Metric) String() string ***REMOVED***
	if this == nil ***REMOVED***
		return "nil"
	***REMOVED***
	s := strings.Join([]string***REMOVED***`&Metric***REMOVED***`,
		`Timestamp:` + strings.Replace(strings.Replace(this.Timestamp.String(), "Timestamp", "google_protobuf2.Timestamp", 1), `&`, ``, 1) + `,`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Data:` + strings.Replace(fmt.Sprintf("%v", this.Data), "Any", "google_protobuf1.Any", 1) + `,`,
		`***REMOVED***`,
	***REMOVED***, "")
	return s
***REMOVED***
func valueToStringMetrics(v interface***REMOVED******REMOVED***) string ***REMOVED***
	rv := reflect.ValueOf(v)
	if rv.IsNil() ***REMOVED***
		return "nil"
	***REMOVED***
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
***REMOVED***
func (m *Metric) Unmarshal(dAtA []byte) error ***REMOVED***
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l ***REMOVED***
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 ***REMOVED***
			if shift >= 64 ***REMOVED***
				return ErrIntOverflowMetrics
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
			return fmt.Errorf("proto: Metric: wiretype end group for non-group")
		***REMOVED***
		if fieldNum <= 0 ***REMOVED***
			return fmt.Errorf("proto: Metric: illegal tag %d (wire type %d)", fieldNum, wire)
		***REMOVED***
		switch fieldNum ***REMOVED***
		case 1:
			if wireType != 2 ***REMOVED***
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			***REMOVED***
			var msglen int
			for shift := uint(0); ; shift += 7 ***REMOVED***
				if shift >= 64 ***REMOVED***
					return ErrIntOverflowMetrics
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
				return ErrInvalidLengthMetrics
			***REMOVED***
			postIndex := iNdEx + msglen
			if postIndex > l ***REMOVED***
				return io.ErrUnexpectedEOF
			***REMOVED***
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil ***REMOVED***
				return err
			***REMOVED***
			iNdEx = postIndex
		case 2:
			if wireType != 2 ***REMOVED***
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			***REMOVED***
			var stringLen uint64
			for shift := uint(0); ; shift += 7 ***REMOVED***
				if shift >= 64 ***REMOVED***
					return ErrIntOverflowMetrics
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
				return ErrInvalidLengthMetrics
			***REMOVED***
			postIndex := iNdEx + intStringLen
			if postIndex > l ***REMOVED***
				return io.ErrUnexpectedEOF
			***REMOVED***
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 ***REMOVED***
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			***REMOVED***
			var msglen int
			for shift := uint(0); ; shift += 7 ***REMOVED***
				if shift >= 64 ***REMOVED***
					return ErrIntOverflowMetrics
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
				return ErrInvalidLengthMetrics
			***REMOVED***
			postIndex := iNdEx + msglen
			if postIndex > l ***REMOVED***
				return io.ErrUnexpectedEOF
			***REMOVED***
			if m.Data == nil ***REMOVED***
				m.Data = &google_protobuf1.Any***REMOVED******REMOVED***
			***REMOVED***
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil ***REMOVED***
				return err
			***REMOVED***
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(dAtA[iNdEx:])
			if err != nil ***REMOVED***
				return err
			***REMOVED***
			if skippy < 0 ***REMOVED***
				return ErrInvalidLengthMetrics
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
func skipMetrics(dAtA []byte) (n int, err error) ***REMOVED***
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l ***REMOVED***
		var wire uint64
		for shift := uint(0); ; shift += 7 ***REMOVED***
			if shift >= 64 ***REMOVED***
				return 0, ErrIntOverflowMetrics
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
					return 0, ErrIntOverflowMetrics
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
					return 0, ErrIntOverflowMetrics
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
				return 0, ErrInvalidLengthMetrics
			***REMOVED***
			return iNdEx, nil
		case 3:
			for ***REMOVED***
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 ***REMOVED***
					if shift >= 64 ***REMOVED***
						return 0, ErrIntOverflowMetrics
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
				next, err := skipMetrics(dAtA[start:])
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
	ErrInvalidLengthMetrics = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetrics   = fmt.Errorf("proto: integer overflow")
)

func init() ***REMOVED***
	proto.RegisterFile("github.com/containerd/containerd/api/types/metrics.proto", fileDescriptorMetrics)
***REMOVED***

var fileDescriptorMetrics = []byte***REMOVED***
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0xcf, 0x2b, 0x49, 0xcc, 0xcc, 0x4b, 0x2d,
	0x4a, 0x41, 0x66, 0x26, 0x16, 0x64, 0xea, 0x97, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0xe7, 0xa6, 0x96,
	0x14, 0x65, 0x26, 0x17, 0xeb, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0x20, 0xd4, 0xe8, 0x81,
	0xe5, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x92, 0xfa, 0x20, 0x16, 0x44, 0x9d, 0x94, 0x64,
	0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x98, 0x97, 0x54, 0x9a, 0xa6, 0x9f, 0x98, 0x57, 0x09,
	0x95, 0x92, 0x47, 0x97, 0x2a, 0xc9, 0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x80, 0x28, 0x50,
	0xea, 0x63, 0xe4, 0x62, 0xf3, 0x05, 0xdb, 0x2a, 0xe4, 0xc4, 0xc5, 0x09, 0x97, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd2, 0x83, 0xe8, 0xd7, 0x83, 0xe9, 0xd7, 0x0b, 0x81, 0xa9, 0x70,
	0xe2, 0x38, 0x71, 0x4f, 0x9e, 0x61, 0xc2, 0x7d, 0x79, 0xc6, 0x20, 0x84, 0x36, 0x21, 0x31, 0x2e,
	0xa6, 0xcc, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0xb6, 0x47, 0xf7, 0xe4, 0x99, 0x3c,
	0x5d, 0x82, 0x98, 0x32, 0x53, 0x84, 0x34, 0xb8, 0x58, 0x52, 0x12, 0x4b, 0x12, 0x25, 0x98, 0xc1,
	0xc6, 0x8a, 0x60, 0x18, 0xeb, 0x98, 0x57, 0x19, 0x04, 0x56, 0xe1, 0xe4, 0x75, 0xe2, 0xa1, 0x1c,
	0xc3, 0x8d, 0x87, 0x72, 0x0c, 0x0d, 0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48,
	0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x28, 0x03, 0xe2, 0x03, 0xd2, 0x1a, 0x4c, 0x46, 0x30, 0x24,
	0xb1, 0x81, 0x6d, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xde, 0x0d, 0x02, 0xfe, 0x85, 0x01,
	0x00, 0x00,
***REMOVED***
