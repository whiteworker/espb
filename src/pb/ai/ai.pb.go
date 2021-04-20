// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ai/ai.proto

package ai

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/emptypb"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MachineAnswerRequest struct {
	Question             string   `protobuf:"bytes,1,opt,name=Question,proto3" json:"Question,omitempty" example:"你今天吃饭了吗?"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MachineAnswerRequest) Reset()         { *m = MachineAnswerRequest{} }
func (m *MachineAnswerRequest) String() string { return proto.CompactTextString(m) }
func (*MachineAnswerRequest) ProtoMessage()    {}
func (*MachineAnswerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_589b0027aa00517e, []int{0}
}
func (m *MachineAnswerRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MachineAnswerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MachineAnswerRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MachineAnswerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MachineAnswerRequest.Merge(m, src)
}
func (m *MachineAnswerRequest) XXX_Size() int {
	return m.Size()
}
func (m *MachineAnswerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MachineAnswerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MachineAnswerRequest proto.InternalMessageInfo

type MachineAnswerResponse struct {
	Answer               string   `protobuf:"bytes,1,opt,name=Answer,proto3" json:"Answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MachineAnswerResponse) Reset()         { *m = MachineAnswerResponse{} }
func (m *MachineAnswerResponse) String() string { return proto.CompactTextString(m) }
func (*MachineAnswerResponse) ProtoMessage()    {}
func (*MachineAnswerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_589b0027aa00517e, []int{1}
}
func (m *MachineAnswerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MachineAnswerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MachineAnswerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MachineAnswerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MachineAnswerResponse.Merge(m, src)
}
func (m *MachineAnswerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MachineAnswerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MachineAnswerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MachineAnswerResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MachineAnswerRequest)(nil), "AI.service.v1.MachineAnswerRequest")
	proto.RegisterType((*MachineAnswerResponse)(nil), "AI.service.v1.MachineAnswerResponse")
}

func init() { proto.RegisterFile("ai/ai.proto", fileDescriptor_589b0027aa00517e) }

var fileDescriptor_589b0027aa00517e = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcc, 0xd4, 0x4f,
	0xcc, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x75, 0xf4, 0xd4, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x33, 0x94, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x4a, 0x2a, 0x4d, 0x03, 0xf3, 0xc0,
	0x1c, 0x30, 0x0b, 0xa2, 0x5b, 0x4a, 0x3a, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x15, 0xa1, 0x2a, 0x35,
	0xb7, 0xa0, 0xa4, 0x12, 0x2a, 0x29, 0x03, 0x95, 0x4c, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb,
	0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xc8, 0x2a, 0x45, 0x70, 0x89, 0xf8, 0x26, 0x26,
	0x67, 0x64, 0xe6, 0xa5, 0x3a, 0xe6, 0x15, 0x97, 0xa7, 0x16, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16,
	0x97, 0x08, 0x39, 0x70, 0x71, 0x04, 0x82, 0x18, 0x99, 0xf9, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x4e, 0x2a, 0x9f, 0xee, 0xc9, 0x2b, 0xa4, 0x56, 0x24, 0xe6, 0x16, 0xe4, 0xa4, 0x5a, 0x29,
	0x3d, 0xd9, 0xbb, 0xe0, 0xc9, 0xee, 0xae, 0xa7, 0x4b, 0x56, 0x3e, 0x9d, 0xd0, 0xfc, 0x72, 0xe9,
	0xda, 0x27, 0xbb, 0xda, 0x9e, 0x4e, 0x98, 0x6e, 0xaf, 0x14, 0x04, 0xd7, 0xa5, 0xa4, 0xcf, 0x25,
	0x8a, 0x66, 0x72, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x18, 0x17, 0x1b, 0x44, 0x04, 0x62,
	0x70, 0x10, 0x94, 0x67, 0x54, 0xc8, 0xc5, 0xe4, 0xe8, 0x29, 0x94, 0xcd, 0xc5, 0x8b, 0xa2, 0x4d,
	0x48, 0x59, 0x0f, 0x25, 0x6c, 0xf4, 0xb0, 0x39, 0x57, 0x4a, 0x05, 0xbf, 0x22, 0x88, 0xcd, 0x4a,
	0xfc, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0xe2, 0x54, 0x62, 0xd1, 0x4f, 0x2c, 0xce, 0xb6, 0x62, 0xd4,
	0x72, 0x92, 0x38, 0xf1, 0x50, 0x8e, 0xe1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x8a, 0x29, 0x31, 0x33, 0x89, 0x0d, 0x1c, 0x3c,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0x60, 0x46, 0x5b, 0xa6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AIClient is the client API for AI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AIClient interface {
	MachineAnswer(ctx context.Context, in *MachineAnswerRequest, opts ...grpc.CallOption) (*MachineAnswerResponse, error)
}

type aIClient struct {
	cc *grpc.ClientConn
}

func NewAIClient(cc *grpc.ClientConn) AIClient {
	return &aIClient{cc}
}

func (c *aIClient) MachineAnswer(ctx context.Context, in *MachineAnswerRequest, opts ...grpc.CallOption) (*MachineAnswerResponse, error) {
	out := new(MachineAnswerResponse)
	err := c.cc.Invoke(ctx, "/AI.service.v1.AI/MachineAnswer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AIServer is the server API for AI service.
type AIServer interface {
	MachineAnswer(context.Context, *MachineAnswerRequest) (*MachineAnswerResponse, error)
}

// UnimplementedAIServer can be embedded to have forward compatible implementations.
type UnimplementedAIServer struct {
}

func (*UnimplementedAIServer) MachineAnswer(ctx context.Context, req *MachineAnswerRequest) (*MachineAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MachineAnswer not implemented")
}

func RegisterAIServer(s *grpc.Server, srv AIServer) {
	s.RegisterService(&_AI_serviceDesc, srv)
}

func _AI_MachineAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachineAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AIServer).MachineAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AI.service.v1.AI/MachineAnswer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AIServer).MachineAnswer(ctx, req.(*MachineAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AI.service.v1.AI",
	HandlerType: (*AIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MachineAnswer",
			Handler:    _AI_MachineAnswer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ai/ai.proto",
}

func (m *MachineAnswerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MachineAnswerRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MachineAnswerRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Question) > 0 {
		i -= len(m.Question)
		copy(dAtA[i:], m.Question)
		i = encodeVarintAi(dAtA, i, uint64(len(m.Question)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MachineAnswerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MachineAnswerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MachineAnswerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Answer) > 0 {
		i -= len(m.Answer)
		copy(dAtA[i:], m.Answer)
		i = encodeVarintAi(dAtA, i, uint64(len(m.Answer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAi(dAtA []byte, offset int, v uint64) int {
	offset -= sovAi(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MachineAnswerRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Question)
	if l > 0 {
		n += 1 + l + sovAi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MachineAnswerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Answer)
	if l > 0 {
		n += 1 + l + sovAi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAi(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAi(x uint64) (n int) {
	return sovAi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MachineAnswerRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAi
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
			return fmt.Errorf("proto: MachineAnswerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MachineAnswerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Question", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAi
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
				return ErrInvalidLengthAi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Question = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAi
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
func (m *MachineAnswerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAi
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
			return fmt.Errorf("proto: MachineAnswerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MachineAnswerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Answer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAi
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
				return ErrInvalidLengthAi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Answer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAi
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
func skipAi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAi
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
					return 0, ErrIntOverflowAi
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
					return 0, ErrIntOverflowAi
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
				return 0, ErrInvalidLengthAi
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAi
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAi
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAi        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAi          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAi = fmt.Errorf("proto: unexpected end of group")
)
