// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lottery/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgPlaceBet struct {
	Sender string     `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Fee    types.Coin `protobuf:"bytes,2,opt,name=fee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"fee"`
	Bet    types.Coin `protobuf:"bytes,3,opt,name=bet,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"bet"`
}

func (m *MsgPlaceBet) Reset()         { *m = MsgPlaceBet{} }
func (m *MsgPlaceBet) String() string { return proto.CompactTextString(m) }
func (*MsgPlaceBet) ProtoMessage()    {}
func (*MsgPlaceBet) Descriptor() ([]byte, []int) {
	return fileDescriptor_16a4f365e1c6455d, []int{0}
}
func (m *MsgPlaceBet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPlaceBet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPlaceBet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPlaceBet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlaceBet.Merge(m, src)
}
func (m *MsgPlaceBet) XXX_Size() int {
	return m.Size()
}
func (m *MsgPlaceBet) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlaceBet.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlaceBet proto.InternalMessageInfo

func (m *MsgPlaceBet) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgPlaceBet) GetFee() types.Coin {
	if m != nil {
		return m.Fee
	}
	return types.Coin{}
}

func (m *MsgPlaceBet) GetBet() types.Coin {
	if m != nil {
		return m.Bet
	}
	return types.Coin{}
}

type MsgPlaceBetResponse struct {
}

func (m *MsgPlaceBetResponse) Reset()         { *m = MsgPlaceBetResponse{} }
func (m *MsgPlaceBetResponse) String() string { return proto.CompactTextString(m) }
func (*MsgPlaceBetResponse) ProtoMessage()    {}
func (*MsgPlaceBetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_16a4f365e1c6455d, []int{1}
}
func (m *MsgPlaceBetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPlaceBetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPlaceBetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPlaceBetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPlaceBetResponse.Merge(m, src)
}
func (m *MsgPlaceBetResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgPlaceBetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPlaceBetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPlaceBetResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgPlaceBet)(nil), "game.lottery.MsgPlaceBet")
	proto.RegisterType((*MsgPlaceBetResponse)(nil), "game.lottery.MsgPlaceBetResponse")
}

func init() { proto.RegisterFile("lottery/tx.proto", fileDescriptor_16a4f365e1c6455d) }

var fileDescriptor_16a4f365e1c6455d = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xc9, 0x2f, 0x29,
	0x49, 0x2d, 0xaa, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x49, 0x4f,
	0xcc, 0x4d, 0xd5, 0x83, 0x0a, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x25, 0xf4, 0x41, 0x2c,
	0x88, 0x1a, 0x29, 0xc9, 0xe4, 0xfc, 0xe2, 0xdc, 0xfc, 0xe2, 0x78, 0x88, 0x04, 0x84, 0x03, 0x95,
	0x92, 0x83, 0xf0, 0xf4, 0x93, 0x12, 0x8b, 0x53, 0xf5, 0xcb, 0x0c, 0x93, 0x52, 0x4b, 0x12, 0x0d,
	0xf5, 0x93, 0xf3, 0x33, 0xf3, 0x20, 0xf2, 0x4a, 0x0f, 0x19, 0xb9, 0xb8, 0x7d, 0x8b, 0xd3, 0x03,
	0x72, 0x12, 0x93, 0x53, 0x9d, 0x52, 0x4b, 0x84, 0xc4, 0xb8, 0xd8, 0x8a, 0x53, 0xf3, 0x52, 0x52,
	0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0xa1, 0x18, 0x2e, 0xe6, 0xb4, 0xd4,
	0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x49, 0x3d, 0xa8, 0x1d, 0x20, 0x53, 0xf5, 0xa0,
	0xa6, 0xea, 0x39, 0xe7, 0x67, 0xe6, 0x39, 0xe9, 0x9f, 0xb8, 0x27, 0xcf, 0xb0, 0xea, 0xbe, 0xbc,
	0x7a, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0x2e, 0xd4, 0x41, 0x50, 0x4a, 0xb7,
	0x38, 0x25, 0x5b, 0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x18, 0xac, 0x21, 0x08, 0x64, 0x2c, 0xc8, 0xf4,
	0xa4, 0xd4, 0x12, 0x09, 0x66, 0xea, 0x9b, 0x9e, 0x94, 0x5a, 0xa2, 0x24, 0xca, 0x25, 0x8c, 0xe4,
	0xc5, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x23, 0x7f, 0x2e, 0x66, 0xdf, 0xe2, 0x74,
	0x21, 0x0f, 0x2e, 0x0e, 0xb8, 0xef, 0x25, 0xf5, 0x90, 0x43, 0x5b, 0x0f, 0x49, 0x97, 0x94, 0x22,
	0x4e, 0x29, 0x98, 0x81, 0x4e, 0x2e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0,
	0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10,
	0xa5, 0x85, 0xe4, 0xde, 0xec, 0xc4, 0xa2, 0x92, 0x8c, 0xcc, 0x6c, 0x63, 0x13, 0x03, 0x7d, 0x90,
	0x89, 0xfa, 0x15, 0xfa, 0xf0, 0x38, 0x07, 0xb9, 0x3b, 0x89, 0x0d, 0x1c, 0x31, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x78, 0x41, 0xd8, 0xa6, 0x0b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	PlaceBet(ctx context.Context, in *MsgPlaceBet, opts ...grpc.CallOption) (*MsgPlaceBetResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) PlaceBet(ctx context.Context, in *MsgPlaceBet, opts ...grpc.CallOption) (*MsgPlaceBetResponse, error) {
	out := new(MsgPlaceBetResponse)
	err := c.cc.Invoke(ctx, "/game.lottery.Msg/PlaceBet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	PlaceBet(context.Context, *MsgPlaceBet) (*MsgPlaceBetResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) PlaceBet(ctx context.Context, req *MsgPlaceBet) (*MsgPlaceBetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceBet not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_PlaceBet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPlaceBet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PlaceBet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.lottery.Msg/PlaceBet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PlaceBet(ctx, req.(*MsgPlaceBet))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "game.lottery.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PlaceBet",
			Handler:    _Msg_PlaceBet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lottery/tx.proto",
}

func (m *MsgPlaceBet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlaceBet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPlaceBet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Bet.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgPlaceBetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPlaceBetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPlaceBetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgPlaceBet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Fee.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.Bet.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgPlaceBetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgPlaceBet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgPlaceBet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlaceBet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgPlaceBetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgPlaceBetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPlaceBetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
