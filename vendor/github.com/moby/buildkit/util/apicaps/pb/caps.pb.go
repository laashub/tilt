// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: caps.proto

package moby_buildkit_v1_apicaps

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// APICap defines a capability supported by the service
type APICap struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Enabled              bool     `protobuf:"varint,2,opt,name=Enabled,proto3" json:"Enabled,omitempty"`
	Deprecated           bool     `protobuf:"varint,3,opt,name=Deprecated,proto3" json:"Deprecated,omitempty"`
	DisabledReason       string   `protobuf:"bytes,4,opt,name=DisabledReason,proto3" json:"DisabledReason,omitempty"`
	DisabledReasonMsg    string   `protobuf:"bytes,5,opt,name=DisabledReasonMsg,proto3" json:"DisabledReasonMsg,omitempty"`
	DisabledAlternative  string   `protobuf:"bytes,6,opt,name=DisabledAlternative,proto3" json:"DisabledAlternative,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APICap) Reset()         { *m = APICap{} }
func (m *APICap) String() string { return proto.CompactTextString(m) }
func (*APICap) ProtoMessage()    {}
func (*APICap) Descriptor() ([]byte, []int) {
	return fileDescriptor_e19c39d9fcb89b83, []int{0}
}
func (m *APICap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *APICap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_APICap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *APICap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APICap.Merge(m, src)
}
func (m *APICap) XXX_Size() int {
	return m.Size()
}
func (m *APICap) XXX_DiscardUnknown() {
	xxx_messageInfo_APICap.DiscardUnknown(m)
}

var xxx_messageInfo_APICap proto.InternalMessageInfo

func (m *APICap) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *APICap) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *APICap) GetDeprecated() bool {
	if m != nil {
		return m.Deprecated
	}
	return false
}

func (m *APICap) GetDisabledReason() string {
	if m != nil {
		return m.DisabledReason
	}
	return ""
}

func (m *APICap) GetDisabledReasonMsg() string {
	if m != nil {
		return m.DisabledReasonMsg
	}
	return ""
}

func (m *APICap) GetDisabledAlternative() string {
	if m != nil {
		return m.DisabledAlternative
	}
	return ""
}

func init() {
	proto.RegisterType((*APICap)(nil), "moby.buildkit.v1.apicaps.APICap")
}

func init() { proto.RegisterFile("caps.proto", fileDescriptor_e19c39d9fcb89b83) }

var fileDescriptor_e19c39d9fcb89b83 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4e, 0x2c, 0x28,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xc8, 0xcd, 0x4f, 0xaa, 0xd4, 0x4b, 0x2a, 0xcd,
	0xcc, 0x49, 0xc9, 0xce, 0x2c, 0xd1, 0x2b, 0x33, 0xd4, 0x4b, 0x2c, 0xc8, 0x04, 0xc9, 0x4b, 0xe9,
	0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7, 0xeb,
	0x83, 0x35, 0x24, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0x31, 0x48, 0xe9, 0x16, 0x23,
	0x17, 0x9b, 0x63, 0x80, 0xa7, 0x73, 0x62, 0x81, 0x10, 0x1f, 0x17, 0x93, 0xa7, 0x8b, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x93, 0xa7, 0x8b, 0x90, 0x04, 0x17, 0xbb, 0x6b, 0x5e, 0x62, 0x52,
	0x4e, 0x6a, 0x8a, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x8c, 0x2b, 0x24, 0xc7, 0xc5, 0xe5,
	0x92, 0x5a, 0x50, 0x94, 0x9a, 0x9c, 0x58, 0x92, 0x9a, 0x22, 0xc1, 0x0c, 0x96, 0x44, 0x12, 0x11,
	0x52, 0xe3, 0xe2, 0x73, 0xc9, 0x2c, 0x06, 0xab, 0x0d, 0x4a, 0x4d, 0x2c, 0xce, 0xcf, 0x93, 0x60,
	0x01, 0x9b, 0x8a, 0x26, 0x2a, 0xa4, 0xc3, 0x25, 0x88, 0x2a, 0xe2, 0x5b, 0x9c, 0x2e, 0xc1, 0x0a,
	0x56, 0x8a, 0x29, 0x21, 0x64, 0xc0, 0x25, 0x0c, 0x13, 0x74, 0xcc, 0x29, 0x49, 0x2d, 0xca, 0x4b,
	0x2c, 0xc9, 0x2c, 0x4b, 0x95, 0x60, 0x03, 0xab, 0xc7, 0x26, 0xe5, 0xc4, 0x73, 0xe2, 0x91, 0x1c,
	0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x26, 0xb1, 0x81, 0x7d, 0x6c, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x02, 0x2d, 0x9e, 0x91, 0x48, 0x01, 0x00, 0x00,
}

func (m *APICap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *APICap) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCaps(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.Enabled {
		dAtA[i] = 0x10
		i++
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Deprecated {
		dAtA[i] = 0x18
		i++
		if m.Deprecated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.DisabledReason) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCaps(dAtA, i, uint64(len(m.DisabledReason)))
		i += copy(dAtA[i:], m.DisabledReason)
	}
	if len(m.DisabledReasonMsg) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintCaps(dAtA, i, uint64(len(m.DisabledReasonMsg)))
		i += copy(dAtA[i:], m.DisabledReasonMsg)
	}
	if len(m.DisabledAlternative) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintCaps(dAtA, i, uint64(len(m.DisabledAlternative)))
		i += copy(dAtA[i:], m.DisabledAlternative)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintCaps(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *APICap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovCaps(uint64(l))
	}
	if m.Enabled {
		n += 2
	}
	if m.Deprecated {
		n += 2
	}
	l = len(m.DisabledReason)
	if l > 0 {
		n += 1 + l + sovCaps(uint64(l))
	}
	l = len(m.DisabledReasonMsg)
	if l > 0 {
		n += 1 + l + sovCaps(uint64(l))
	}
	l = len(m.DisabledAlternative)
	if l > 0 {
		n += 1 + l + sovCaps(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCaps(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCaps(x uint64) (n int) {
	return sovCaps(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *APICap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCaps
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
			return fmt.Errorf("proto: APICap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: APICap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaps
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
				return ErrInvalidLengthCaps
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCaps
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Enabled = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deprecated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaps
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deprecated = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisabledReason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaps
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
				return ErrInvalidLengthCaps
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCaps
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisabledReason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisabledReasonMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaps
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
				return ErrInvalidLengthCaps
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCaps
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisabledReasonMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisabledAlternative", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaps
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
				return ErrInvalidLengthCaps
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCaps
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisabledAlternative = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCaps(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCaps
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCaps
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCaps(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCaps
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
					return 0, ErrIntOverflowCaps
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCaps
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
				return 0, ErrInvalidLengthCaps
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCaps
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCaps
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCaps(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCaps
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCaps = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCaps   = fmt.Errorf("proto: integer overflow")
)
