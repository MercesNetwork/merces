// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planet/twitter_coins.proto

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

type TwitterCoins struct {
	Index  string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Amount int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *TwitterCoins) Reset()         { *m = TwitterCoins{} }
func (m *TwitterCoins) String() string { return proto.CompactTextString(m) }
func (*TwitterCoins) ProtoMessage()    {}
func (*TwitterCoins) Descriptor() ([]byte, []int) {
	return fileDescriptor_17f3a5f122a338ef, []int{0}
}
func (m *TwitterCoins) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TwitterCoins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TwitterCoins.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TwitterCoins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TwitterCoins.Merge(m, src)
}
func (m *TwitterCoins) XXX_Size() int {
	return m.Size()
}
func (m *TwitterCoins) XXX_DiscardUnknown() {
	xxx_messageInfo_TwitterCoins.DiscardUnknown(m)
}

var xxx_messageInfo_TwitterCoins proto.InternalMessageInfo

func (m *TwitterCoins) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *TwitterCoins) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func init() {
	proto.RegisterType((*TwitterCoins)(nil), "mercestoken.planet.planet.TwitterCoins")
}

func init() { proto.RegisterFile("planet/twitter_coins.proto", fileDescriptor_17f3a5f122a338ef) }

var fileDescriptor_17f3a5f122a338ef = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0xc8, 0x49, 0xcc,
	0x4b, 0x2d, 0xd1, 0x2f, 0x29, 0xcf, 0x2c, 0x29, 0x49, 0x2d, 0x8a, 0x4f, 0xce, 0xcf, 0xcc, 0x2b,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xcc, 0x4d, 0x2d, 0x4a, 0x4e, 0x2d, 0x2e, 0xc9,
	0xcf, 0x4e, 0xcd, 0xd3, 0x83, 0xa8, 0x83, 0x52, 0x4a, 0x36, 0x5c, 0x3c, 0x21, 0x10, 0x1d, 0xce,
	0x20, 0x0d, 0x42, 0x22, 0x5c, 0xac, 0x99, 0x79, 0x29, 0xa9, 0x15, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x10, 0x8e, 0x90, 0x18, 0x17, 0x5b, 0x62, 0x6e, 0x7e, 0x69, 0x5e, 0x89, 0x04, 0x93,
	0x02, 0xa3, 0x06, 0x73, 0x10, 0x94, 0xe7, 0xe4, 0x76, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72,
	0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7,
	0x72, 0x0c, 0x51, 0x3a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xbe,
	0x60, 0xdb, 0x43, 0x40, 0xb6, 0xeb, 0x43, 0x5d, 0x59, 0x01, 0x63, 0x94, 0x54, 0x16, 0xa4, 0x16,
	0x27, 0xb1, 0x81, 0xdd, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xf7, 0xe1, 0x2f, 0xc5,
	0x00, 0x00, 0x00,
}

func (m *TwitterCoins) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TwitterCoins) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TwitterCoins) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintTwitterCoins(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTwitterCoins(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTwitterCoins(dAtA []byte, offset int, v uint64) int {
	offset -= sovTwitterCoins(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TwitterCoins) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTwitterCoins(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovTwitterCoins(uint64(m.Amount))
	}
	return n
}

func sovTwitterCoins(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTwitterCoins(x uint64) (n int) {
	return sovTwitterCoins(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TwitterCoins) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTwitterCoins
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
			return fmt.Errorf("proto: TwitterCoins: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TwitterCoins: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwitterCoins
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
				return ErrInvalidLengthTwitterCoins
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTwitterCoins
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTwitterCoins
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTwitterCoins(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTwitterCoins
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
func skipTwitterCoins(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTwitterCoins
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
					return 0, ErrIntOverflowTwitterCoins
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
					return 0, ErrIntOverflowTwitterCoins
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
				return 0, ErrInvalidLengthTwitterCoins
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTwitterCoins
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTwitterCoins
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTwitterCoins        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTwitterCoins          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTwitterCoins = fmt.Errorf("proto: unexpected end of group")
)
