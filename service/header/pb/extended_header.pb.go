// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: extended_header.proto

package header_pb

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	da "github.com/tendermint/tendermint/proto/tendermint/da"
	types "github.com/tendermint/tendermint/proto/tendermint/types"
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

type ExtendedHeader struct {
	Header       *types.Header              `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Commit       *types.Commit              `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	ValidatorSet *types.ValidatorSet        `protobuf:"bytes,3,opt,name=validator_set,json=validatorSet,proto3" json:"validator_set,omitempty"`
	Dah          *da.DataAvailabilityHeader `protobuf:"bytes,4,opt,name=dah,proto3" json:"dah,omitempty"`
}

func (m *ExtendedHeader) Reset()         { *m = ExtendedHeader{} }
func (m *ExtendedHeader) String() string { return proto.CompactTextString(m) }
func (*ExtendedHeader) ProtoMessage()    {}
func (*ExtendedHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_c13a6e9f483d098b, []int{0}
}
func (m *ExtendedHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtendedHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtendedHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtendedHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtendedHeader.Merge(m, src)
}
func (m *ExtendedHeader) XXX_Size() int {
	return m.Size()
}
func (m *ExtendedHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtendedHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ExtendedHeader proto.InternalMessageInfo

func (m *ExtendedHeader) GetHeader() *types.Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ExtendedHeader) GetCommit() *types.Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *ExtendedHeader) GetValidatorSet() *types.ValidatorSet {
	if m != nil {
		return m.ValidatorSet
	}
	return nil
}

func (m *ExtendedHeader) GetDah() *da.DataAvailabilityHeader {
	if m != nil {
		return m.Dah
	}
	return nil
}

type ExtendedHeaderRequest struct {
	Origin uint64 `protobuf:"varint,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Hash   []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Amount uint64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *ExtendedHeaderRequest) Reset()         { *m = ExtendedHeaderRequest{} }
func (m *ExtendedHeaderRequest) String() string { return proto.CompactTextString(m) }
func (*ExtendedHeaderRequest) ProtoMessage()    {}
func (*ExtendedHeaderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c13a6e9f483d098b, []int{1}
}
func (m *ExtendedHeaderRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtendedHeaderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtendedHeaderRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtendedHeaderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtendedHeaderRequest.Merge(m, src)
}
func (m *ExtendedHeaderRequest) XXX_Size() int {
	return m.Size()
}
func (m *ExtendedHeaderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtendedHeaderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExtendedHeaderRequest proto.InternalMessageInfo

func (m *ExtendedHeaderRequest) GetOrigin() uint64 {
	if m != nil {
		return m.Origin
	}
	return 0
}

func (m *ExtendedHeaderRequest) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ExtendedHeaderRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*ExtendedHeader)(nil), "header.pb.ExtendedHeader")
	proto.RegisterType((*ExtendedHeaderRequest)(nil), "header.pb.ExtendedHeaderRequest")
}

func init() { proto.RegisterFile("extended_header.proto", fileDescriptor_c13a6e9f483d098b) }

var fileDescriptor_c13a6e9f483d098b = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x3b, 0x7f, 0x43, 0xe1, 0x1f, 0xab, 0x8b, 0x81, 0xca, 0x50, 0x64, 0x28, 0x05, 0xc1,
	0x85, 0xa4, 0xa2, 0x0b, 0xd7, 0x5a, 0x05, 0xd7, 0x23, 0xb8, 0x71, 0x51, 0x6e, 0x9d, 0xc1, 0x0c,
	0x34, 0x99, 0x9a, 0xdc, 0x16, 0xfb, 0x16, 0x3e, 0x96, 0xcb, 0x2e, 0x5d, 0x4a, 0xf2, 0x06, 0x3e,
	0x81, 0x64, 0x66, 0xaa, 0x29, 0xc5, 0x4d, 0xc8, 0xe1, 0x7c, 0xe7, 0xe6, 0x9c, 0xd0, 0x9e, 0x7e,
	0x45, 0x9d, 0x29, 0xad, 0x26, 0x89, 0x06, 0xa5, 0xf3, 0x78, 0x9e, 0x5b, 0xb4, 0xec, 0xff, 0x46,
	0x4d, 0xfb, 0x47, 0xce, 0xcf, 0x53, 0x93, 0xe1, 0x08, 0x57, 0x73, 0x5d, 0xf8, 0xa7, 0x07, 0xfb,
	0x83, 0x1d, 0x77, 0x09, 0x33, 0xa3, 0x00, 0x6d, 0x38, 0xd5, 0x3f, 0x6d, 0x10, 0x0a, 0x46, 0x0a,
	0x10, 0x26, 0xb0, 0x04, 0x33, 0x83, 0xa9, 0x99, 0x19, 0x5c, 0x6d, 0x7d, 0x78, 0xf8, 0x45, 0xe8,
	0xc1, 0x6d, 0xa8, 0x74, 0xe7, 0x0c, 0x76, 0x46, 0x3b, 0x1e, 0xe1, 0x64, 0x40, 0x4e, 0xf6, 0xce,
	0x79, 0xfc, 0x7b, 0x31, 0xf6, 0x5d, 0x3c, 0x29, 0x03, 0x57, 0x27, 0x9e, 0x6c, 0x9a, 0x1a, 0xe4,
	0xff, 0xfe, 0x4a, 0x8c, 0x9d, 0x2f, 0x03, 0xc7, 0xc6, 0x74, 0xff, 0xa7, 0xf7, 0xa4, 0xd0, 0xc8,
	0xdb, 0x2e, 0x28, 0x76, 0x83, 0x0f, 0x1b, 0xec, 0x5e, 0xa3, 0xec, 0x2e, 0x1b, 0x8a, 0x5d, 0xd2,
	0xb6, 0x82, 0x84, 0x47, 0x2e, 0x7a, 0xdc, 0x8c, 0x2a, 0x88, 0x6f, 0x00, 0xe1, 0xaa, 0x31, 0x3b,
	0x54, 0xae, 0x13, 0xc3, 0x47, 0xda, 0xdb, 0xde, 0x2c, 0xf5, 0xcb, 0x42, 0x17, 0xc8, 0x0e, 0x69,
	0xc7, 0xe6, 0xe6, 0xd9, 0x64, 0x6e, 0x7a, 0x24, 0x83, 0x62, 0x8c, 0x46, 0x09, 0x14, 0x89, 0x9b,
	0xd7, 0x95, 0xee, 0xbd, 0x66, 0x21, 0xb5, 0x8b, 0xcc, 0x77, 0x8f, 0x64, 0x50, 0xd7, 0xfc, 0xbd,
	0x14, 0x64, 0x5d, 0x0a, 0xf2, 0x59, 0x0a, 0xf2, 0x56, 0x89, 0xd6, 0xba, 0x12, 0xad, 0x8f, 0x4a,
	0xb4, 0xa6, 0x1d, 0xf7, 0xcb, 0x2f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x71, 0x1d, 0xa7, 0x89,
	0x04, 0x02, 0x00, 0x00,
}

func (m *ExtendedHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtendedHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtendedHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Dah != nil {
		{
			size, err := m.Dah.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExtendedHeader(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.ValidatorSet != nil {
		{
			size, err := m.ValidatorSet.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExtendedHeader(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Commit != nil {
		{
			size, err := m.Commit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExtendedHeader(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExtendedHeader(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ExtendedHeaderRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtendedHeaderRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtendedHeaderRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintExtendedHeader(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintExtendedHeader(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Origin != 0 {
		i = encodeVarintExtendedHeader(dAtA, i, uint64(m.Origin))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintExtendedHeader(dAtA []byte, offset int, v uint64) int {
	offset -= sovExtendedHeader(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExtendedHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovExtendedHeader(uint64(l))
	}
	if m.Commit != nil {
		l = m.Commit.Size()
		n += 1 + l + sovExtendedHeader(uint64(l))
	}
	if m.ValidatorSet != nil {
		l = m.ValidatorSet.Size()
		n += 1 + l + sovExtendedHeader(uint64(l))
	}
	if m.Dah != nil {
		l = m.Dah.Size()
		n += 1 + l + sovExtendedHeader(uint64(l))
	}
	return n
}

func (m *ExtendedHeaderRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Origin != 0 {
		n += 1 + sovExtendedHeader(uint64(m.Origin))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovExtendedHeader(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovExtendedHeader(uint64(m.Amount))
	}
	return n
}

func sovExtendedHeader(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExtendedHeader(x uint64) (n int) {
	return sovExtendedHeader(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtendedHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtendedHeader
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
			return fmt.Errorf("proto: ExtendedHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtendedHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedHeader
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
				return ErrInvalidLengthExtendedHeader
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &types.Header{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedHeader
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
				return ErrInvalidLengthExtendedHeader
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Commit == nil {
				m.Commit = &types.Commit{}
			}
			if err := m.Commit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedHeader
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
				return ErrInvalidLengthExtendedHeader
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ValidatorSet == nil {
				m.ValidatorSet = &types.ValidatorSet{}
			}
			if err := m.ValidatorSet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dah", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedHeader
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
				return ErrInvalidLengthExtendedHeader
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Dah == nil {
				m.Dah = &da.DataAvailabilityHeader{}
			}
			if err := m.Dah.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExtendedHeader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExtendedHeader
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
func (m *ExtendedHeaderRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtendedHeader
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
			return fmt.Errorf("proto: ExtendedHeaderRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtendedHeaderRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			m.Origin = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Origin |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthExtendedHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipExtendedHeader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExtendedHeader
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
func skipExtendedHeader(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExtendedHeader
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
					return 0, ErrIntOverflowExtendedHeader
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
					return 0, ErrIntOverflowExtendedHeader
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
				return 0, ErrInvalidLengthExtendedHeader
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExtendedHeader
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExtendedHeader
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExtendedHeader        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExtendedHeader          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExtendedHeader = fmt.Errorf("proto: unexpected end of group")
)
