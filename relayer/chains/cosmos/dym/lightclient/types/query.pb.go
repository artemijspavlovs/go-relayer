// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymensionxyz/dymension/lightclient/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryRollappCanonChannelRequest struct {
	RollappId string `protobuf:"bytes,1,opt,name=rollappId,proto3" json:"rollappId,omitempty"`
}

func (m *QueryRollappCanonChannelRequest) Reset()         { *m = QueryRollappCanonChannelRequest{} }
func (m *QueryRollappCanonChannelRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRollappCanonChannelRequest) ProtoMessage()    {}
func (*QueryRollappCanonChannelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a51f5810cc34625d, []int{0}
}
func (m *QueryRollappCanonChannelRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRollappCanonChannelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRollappCanonChannelRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRollappCanonChannelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRollappCanonChannelRequest.Merge(m, src)
}
func (m *QueryRollappCanonChannelRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRollappCanonChannelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRollappCanonChannelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRollappCanonChannelRequest proto.InternalMessageInfo

func (m *QueryRollappCanonChannelRequest) GetRollappId() string {
	if m != nil {
		return m.RollappId
	}
	return ""
}

type QueryRollappCanonChannelResponse struct {
	// hub side
	HubChannelId string `protobuf:"bytes,1,opt,name=hub_channel_id,json=hubChannelId,proto3" json:"hub_channel_id,omitempty"`
	// rollapp side ('counterparty')
	RollappChannelId string `protobuf:"bytes,2,opt,name=rollapp_channel_id,json=rollappChannelId,proto3" json:"rollapp_channel_id,omitempty"`
}

func (m *QueryRollappCanonChannelResponse) Reset()         { *m = QueryRollappCanonChannelResponse{} }
func (m *QueryRollappCanonChannelResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRollappCanonChannelResponse) ProtoMessage()    {}
func (*QueryRollappCanonChannelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a51f5810cc34625d, []int{1}
}
func (m *QueryRollappCanonChannelResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRollappCanonChannelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRollappCanonChannelResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRollappCanonChannelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRollappCanonChannelResponse.Merge(m, src)
}
func (m *QueryRollappCanonChannelResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRollappCanonChannelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRollappCanonChannelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRollappCanonChannelResponse proto.InternalMessageInfo

func (m *QueryRollappCanonChannelResponse) GetHubChannelId() string {
	if m != nil {
		return m.HubChannelId
	}
	return ""
}

func (m *QueryRollappCanonChannelResponse) GetRollappChannelId() string {
	if m != nil {
		return m.RollappChannelId
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryRollappCanonChannelRequest)(nil), "dymensionxyz.dymension.lightclient.QueryRollappCanonChannelRequest")
	proto.RegisterType((*QueryRollappCanonChannelResponse)(nil), "dymensionxyz.dymension.lightclient.QueryRollappCanonChannelResponse")
}

func init() {
	proto.RegisterFile("dymensionxyz/dymension/lightclient/query.proto", fileDescriptor_a51f5810cc34625d)
}

var fileDescriptor_a51f5810cc34625d = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x5e, 0x06, 0x0a, 0x0b, 0x22, 0x12, 0x3d, 0xe8, 0x18, 0x75, 0x14, 0x0f, 0x1e, 0xa4, 0x01,
	0x3d, 0xcb, 0xc0, 0xb9, 0xc3, 0x8e, 0xee, 0xe8, 0x65, 0xa4, 0x6d, 0x6c, 0x03, 0x69, 0x5e, 0xd7,
	0xa4, 0x62, 0x15, 0x2f, 0xfe, 0x02, 0xc1, 0x3f, 0xe5, 0x71, 0xe0, 0xc5, 0xa3, 0x6c, 0x82, 0x67,
	0xff, 0x81, 0xac, 0xed, 0xba, 0x09, 0xca, 0x04, 0x4f, 0xc9, 0x7b, 0xdf, 0xf7, 0xbe, 0x2f, 0xf9,
	0x78, 0xd8, 0xf1, 0xb3, 0x88, 0x2b, 0x2d, 0x40, 0xdd, 0x64, 0xb7, 0xb4, 0x2a, 0xa8, 0x14, 0x41,
	0x68, 0x3c, 0x29, 0xb8, 0x32, 0x74, 0x94, 0xf2, 0x24, 0x73, 0xe2, 0x04, 0x0c, 0x10, 0x7b, 0x99,
	0xbf, 0x18, 0x76, 0x96, 0xf8, 0xcd, 0x9d, 0x00, 0x02, 0xc8, 0xe9, 0x74, 0x76, 0x2b, 0x26, 0x9b,
	0xad, 0x00, 0x20, 0x90, 0x9c, 0xb2, 0x58, 0x50, 0xa6, 0x14, 0x18, 0x66, 0x04, 0x28, 0x5d, 0xa2,
	0x7b, 0x25, 0x9a, 0x57, 0x6e, 0x7a, 0x45, 0x99, 0x2a, 0x2d, 0xed, 0x0e, 0xde, 0xbf, 0x98, 0xbd,
	0x60, 0x00, 0x52, 0xb2, 0x38, 0xee, 0x32, 0x05, 0xaa, 0x1b, 0x32, 0xa5, 0xb8, 0x1c, 0xf0, 0x51,
	0xca, 0xb5, 0x21, 0x2d, 0xdc, 0x48, 0x0a, 0xb4, 0xef, 0xef, 0xa2, 0x36, 0x3a, 0x6c, 0x0c, 0x16,
	0x0d, 0xfb, 0x1a, 0xb7, 0x7f, 0x17, 0xd0, 0x31, 0x28, 0xcd, 0xc9, 0x01, 0xde, 0x0c, 0x53, 0x77,
	0xe8, 0x15, 0xed, 0xa1, 0x98, 0xcb, 0x6c, 0x84, 0xa9, 0x5b, 0x72, 0xfb, 0x3e, 0x39, 0xc2, 0xa4,
	0x94, 0x5d, 0x66, 0xd6, 0x73, 0xe6, 0x56, 0x89, 0x54, 0xec, 0xe3, 0x4f, 0x84, 0xd7, 0x72, 0x63,
	0xf2, 0x81, 0xf0, 0xf6, 0x0f, 0xee, 0xa4, 0xeb, 0xac, 0x8e, 0xd3, 0x59, 0xf1, 0xf9, 0xe6, 0xf9,
	0xff, 0x44, 0x8a, 0x00, 0xec, 0xde, 0xc3, 0xcb, 0xfb, 0x53, 0xbd, 0x43, 0x4e, 0xe9, 0x1f, 0x36,
	0xc2, 0x9b, 0x29, 0xcc, 0x23, 0xa0, 0x77, 0x55, 0xd4, 0xf7, 0x67, 0xc3, 0xe7, 0x89, 0x85, 0xc6,
	0x13, 0x0b, 0xbd, 0x4d, 0x2c, 0xf4, 0x38, 0xb5, 0x6a, 0xe3, 0xa9, 0x55, 0x7b, 0x9d, 0x5a, 0xb5,
	0xcb, 0x5e, 0x20, 0x4c, 0x98, 0xba, 0x8e, 0x07, 0x11, 0xf5, 0x40, 0x47, 0xa0, 0x69, 0xc2, 0x25,
	0xcb, 0x78, 0x52, 0x9d, 0x5e, 0xc8, 0x84, 0xd2, 0x73, 0xd4, 0xcf, 0xa2, 0x6f, 0xa6, 0x26, 0x8b,
	0xb9, 0x76, 0xd7, 0xf3, 0xa5, 0x38, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x22, 0x4c, 0x78, 0xe2,
	0xb9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	RollappCanonChannel(ctx context.Context, in *QueryRollappCanonChannelRequest, opts ...grpc.CallOption) (*QueryRollappCanonChannelResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) RollappCanonChannel(ctx context.Context, in *QueryRollappCanonChannelRequest, opts ...grpc.CallOption) (*QueryRollappCanonChannelResponse, error) {
	out := new(QueryRollappCanonChannelResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.lightclient.Query/RollappCanonChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	RollappCanonChannel(context.Context, *QueryRollappCanonChannelRequest) (*QueryRollappCanonChannelResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) RollappCanonChannel(ctx context.Context, req *QueryRollappCanonChannelRequest) (*QueryRollappCanonChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollappCanonChannel not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_RollappCanonChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRollappCanonChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RollappCanonChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.lightclient.Query/RollappCanonChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RollappCanonChannel(ctx, req.(*QueryRollappCanonChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dymensionxyz.dymension.lightclient.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RollappCanonChannel",
			Handler:    _Query_RollappCanonChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dymensionxyz/dymension/lightclient/query.proto",
}

func (m *QueryRollappCanonChannelRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRollappCanonChannelRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRollappCanonChannelRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RollappId) > 0 {
		i -= len(m.RollappId)
		copy(dAtA[i:], m.RollappId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.RollappId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryRollappCanonChannelResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRollappCanonChannelResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRollappCanonChannelResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RollappChannelId) > 0 {
		i -= len(m.RollappChannelId)
		copy(dAtA[i:], m.RollappChannelId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.RollappChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.HubChannelId) > 0 {
		i -= len(m.HubChannelId)
		copy(dAtA[i:], m.HubChannelId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.HubChannelId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryRollappCanonChannelRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RollappId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryRollappCanonChannelResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HubChannelId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.RollappChannelId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRollappCanonChannelRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRollappCanonChannelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRollappCanonChannelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryRollappCanonChannelResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRollappCanonChannelResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRollappCanonChannelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HubChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HubChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollappChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollappChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
