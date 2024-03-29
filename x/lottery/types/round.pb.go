// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lottery/round.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Round struct {
	Val uint64 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (m *Round) Reset()         { *m = Round{} }
func (m *Round) String() string { return proto.CompactTextString(m) }
func (*Round) ProtoMessage()    {}
func (*Round) Descriptor() ([]byte, []int) {
	return fileDescriptor_3541da3b982cac01, []int{0}
}
func (m *Round) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Round) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Round.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Round) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Round.Merge(m, src)
}
func (m *Round) XXX_Size() int {
	return m.Size()
}
func (m *Round) XXX_DiscardUnknown() {
	xxx_messageInfo_Round.DiscardUnknown(m)
}

var xxx_messageInfo_Round proto.InternalMessageInfo

func (m *Round) GetVal() uint64 {
	if m != nil {
		return m.Val
	}
	return 0
}

func init() {
	proto.RegisterType((*Round)(nil), "game.lottery.Round")
}

func init() { proto.RegisterFile("lottery/round.proto", fileDescriptor_3541da3b982cac01) }

var fileDescriptor_3541da3b982cac01 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0xc9, 0x2f, 0x29,
	0x49, 0x2d, 0xaa, 0xd4, 0x2f, 0xca, 0x2f, 0xcd, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x49, 0x4f, 0xcc, 0x4d, 0xd5, 0x83, 0xca, 0x28, 0x49, 0x72, 0xb1, 0x06, 0x81, 0x24, 0x85,
	0x04, 0xb8, 0x98, 0xcb, 0x12, 0x73, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0x40, 0x4c, 0x27,
	0x97, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63,
	0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x4e, 0x2c, 0x2a, 0xc9, 0xc8, 0xcc, 0x36, 0x36,
	0x31, 0xd0, 0x07, 0x19, 0xac, 0x5f, 0xa1, 0x0f, 0xb3, 0xb4, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89,
	0x0d, 0x6c, 0xab, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x94, 0xfc, 0x5b, 0xae, 0x8c, 0x00, 0x00,
	0x00,
}

func (m *Round) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Round) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Round) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Val != 0 {
		i = encodeVarintRound(dAtA, i, uint64(m.Val))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRound(dAtA []byte, offset int, v uint64) int {
	offset -= sovRound(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Round) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Val != 0 {
		n += 1 + sovRound(uint64(m.Val))
	}
	return n
}

func sovRound(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRound(x uint64) (n int) {
	return sovRound(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Round) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRound
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
			return fmt.Errorf("proto: Round: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Round: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
			}
			m.Val = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRound
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Val |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRound(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRound
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
func skipRound(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRound
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
					return 0, ErrIntOverflowRound
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
					return 0, ErrIntOverflowRound
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
				return 0, ErrInvalidLengthRound
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRound
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRound
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRound        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRound          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRound = fmt.Errorf("proto: unexpected end of group")
)
