// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dydxprotocol/clob/equity_tier_limit_config.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_dydxprotocol_v4_chain_protocol_dtypes "github.com/dydxprotocol/v4-chain/protocol/dtypes"
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

// Defines the set of equity tiers to limit how many open orders
// a subaccount is allowed to have.
type EquityTierLimitConfiguration struct {
	// How many short term stateful orders are allowed per equity tier.
	// Specifying 0 values disables this limit.
	ShortTermOrderEquityTiers []EquityTierLimit `protobuf:"bytes,1,rep,name=short_term_order_equity_tiers,json=shortTermOrderEquityTiers,proto3" json:"short_term_order_equity_tiers"`
	// How many open stateful orders are allowed per equity tier.
	// Specifying 0 values disables this limit.
	StatefulOrderEquityTiers []EquityTierLimit `protobuf:"bytes,2,rep,name=stateful_order_equity_tiers,json=statefulOrderEquityTiers,proto3" json:"stateful_order_equity_tiers"`
}

func (m *EquityTierLimitConfiguration) Reset()         { *m = EquityTierLimitConfiguration{} }
func (m *EquityTierLimitConfiguration) String() string { return proto.CompactTextString(m) }
func (*EquityTierLimitConfiguration) ProtoMessage()    {}
func (*EquityTierLimitConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb08cfe6323c23aa, []int{0}
}
func (m *EquityTierLimitConfiguration) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EquityTierLimitConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EquityTierLimitConfiguration.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EquityTierLimitConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EquityTierLimitConfiguration.Merge(m, src)
}
func (m *EquityTierLimitConfiguration) XXX_Size() int {
	return m.Size()
}
func (m *EquityTierLimitConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_EquityTierLimitConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_EquityTierLimitConfiguration proto.InternalMessageInfo

func (m *EquityTierLimitConfiguration) GetShortTermOrderEquityTiers() []EquityTierLimit {
	if m != nil {
		return m.ShortTermOrderEquityTiers
	}
	return nil
}

func (m *EquityTierLimitConfiguration) GetStatefulOrderEquityTiers() []EquityTierLimit {
	if m != nil {
		return m.StatefulOrderEquityTiers
	}
	return nil
}

// Defines an equity tier limit.
type EquityTierLimit struct {
	// The total net collateral in USDC quote quantums of equity required.
	UsdTncRequired github_com_dydxprotocol_v4_chain_protocol_dtypes.SerializableInt `protobuf:"bytes,1,opt,name=usd_tnc_required,json=usdTncRequired,proto3,customtype=github.com/dydxprotocol/v4-chain/protocol/dtypes.SerializableInt" json:"usd_tnc_required"`
	// What the limit is for `usd_tnc_required`.
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (m *EquityTierLimit) Reset()         { *m = EquityTierLimit{} }
func (m *EquityTierLimit) String() string { return proto.CompactTextString(m) }
func (*EquityTierLimit) ProtoMessage()    {}
func (*EquityTierLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb08cfe6323c23aa, []int{1}
}
func (m *EquityTierLimit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EquityTierLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EquityTierLimit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EquityTierLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EquityTierLimit.Merge(m, src)
}
func (m *EquityTierLimit) XXX_Size() int {
	return m.Size()
}
func (m *EquityTierLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_EquityTierLimit.DiscardUnknown(m)
}

var xxx_messageInfo_EquityTierLimit proto.InternalMessageInfo

func (m *EquityTierLimit) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterType((*EquityTierLimitConfiguration)(nil), "dydxprotocol.clob.EquityTierLimitConfiguration")
	proto.RegisterType((*EquityTierLimit)(nil), "dydxprotocol.clob.EquityTierLimit")
}

func init() {
	proto.RegisterFile("dydxprotocol/clob/equity_tier_limit_config.proto", fileDescriptor_eb08cfe6323c23aa)
}

var fileDescriptor_eb08cfe6323c23aa = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbb, 0x4e, 0xfb, 0x30,
	0x14, 0xc6, 0xe3, 0xfe, 0x2f, 0x83, 0xb9, 0x47, 0x1d, 0xc2, 0x2d, 0xad, 0x3a, 0x75, 0x21, 0x41,
	0x80, 0x98, 0x51, 0x11, 0x12, 0x48, 0x48, 0xa0, 0xd0, 0x89, 0xc5, 0x4a, 0x6c, 0x37, 0x35, 0x4a,
	0xe2, 0x62, 0x9f, 0xa0, 0x96, 0xa7, 0x60, 0xe7, 0x85, 0x3a, 0x76, 0x44, 0x0c, 0x15, 0x6a, 0x1f,
	0x81, 0x17, 0x40, 0x71, 0xb9, 0x94, 0xc2, 0x00, 0x5b, 0x72, 0xf4, 0x7d, 0xbf, 0xdf, 0x91, 0x8e,
	0xf1, 0x36, 0xeb, 0xb1, 0x6e, 0x47, 0x49, 0x90, 0x54, 0x26, 0x3e, 0x4d, 0x64, 0xe4, 0xf3, 0xeb,
	0x5c, 0x40, 0x8f, 0x80, 0xe0, 0x8a, 0x24, 0x22, 0x15, 0x40, 0xa8, 0xcc, 0x5a, 0x22, 0xf6, 0x4c,
	0xcc, 0x5e, 0x99, 0x6e, 0x78, 0x45, 0x63, 0xad, 0x1c, 0xcb, 0x58, 0x9a, 0x91, 0x5f, 0x7c, 0x4d,
	0x82, 0xb5, 0x67, 0x84, 0x37, 0x8e, 0x0c, 0xab, 0x29, 0xb8, 0x3a, 0x2d, 0x48, 0x87, 0x06, 0x94,
	0xab, 0x10, 0x84, 0xcc, 0xec, 0x2b, 0xbc, 0xa9, 0xdb, 0x52, 0x01, 0x01, 0xae, 0x52, 0x22, 0x15,
	0xe3, 0x8a, 0x4c, 0xc9, 0xb5, 0x83, 0xaa, 0x7f, 0xea, 0x73, 0x3b, 0x35, 0xef, 0x8b, 0xd1, 0x9b,
	0xe1, 0x36, 0xfe, 0xf6, 0x87, 0x15, 0x2b, 0x58, 0x35, 0xb8, 0x26, 0x57, 0xe9, 0x59, 0x01, 0xfb,
	0x08, 0x69, 0x3b, 0xc6, 0xeb, 0x1a, 0x42, 0xe0, 0xad, 0x3c, 0xf9, 0xce, 0x54, 0xfa, 0xa5, 0xc9,
	0x79, 0x83, 0xcd, 0x8a, 0x6a, 0xf7, 0x08, 0x2f, 0xcd, 0x74, 0x6c, 0x85, 0x97, 0x73, 0xcd, 0x08,
	0x64, 0x94, 0xa8, 0x42, 0xab, 0x38, 0x73, 0x50, 0x15, 0xd5, 0xe7, 0x1b, 0xc7, 0x05, 0xed, 0x71,
	0x58, 0x39, 0x88, 0x05, 0xb4, 0xf3, 0xc8, 0xa3, 0x32, 0xf5, 0x3f, 0x5d, 0xe4, 0x66, 0x6f, 0x8b,
	0xb6, 0x43, 0x91, 0xf9, 0xef, 0x13, 0x06, 0xbd, 0x0e, 0xd7, 0xde, 0x05, 0x57, 0x22, 0x4c, 0xc4,
	0x6d, 0x18, 0x25, 0xfc, 0x24, 0x83, 0x60, 0x31, 0xd7, 0xac, 0x99, 0xd1, 0xe0, 0x95, 0x6f, 0x97,
	0xf1, 0x3f, 0x73, 0x3c, 0xa7, 0x54, 0x45, 0xf5, 0x85, 0x60, 0xf2, 0xd3, 0x38, 0xef, 0x8f, 0x5c,
	0x34, 0x18, 0xb9, 0xe8, 0x69, 0xe4, 0xa2, 0xbb, 0xb1, 0x6b, 0x0d, 0xc6, 0xae, 0xf5, 0x30, 0x76,
	0xad, 0xcb, 0xfd, 0x9f, 0x6f, 0xd0, 0x9d, 0xbc, 0x13, 0xb3, 0x47, 0xf4, 0xdf, 0x8c, 0x77, 0x5f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x40, 0x8b, 0xef, 0x78, 0x49, 0x02, 0x00, 0x00,
}

func (m *EquityTierLimitConfiguration) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EquityTierLimitConfiguration) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EquityTierLimitConfiguration) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StatefulOrderEquityTiers) > 0 {
		for iNdEx := len(m.StatefulOrderEquityTiers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StatefulOrderEquityTiers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEquityTierLimitConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ShortTermOrderEquityTiers) > 0 {
		for iNdEx := len(m.ShortTermOrderEquityTiers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ShortTermOrderEquityTiers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEquityTierLimitConfig(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *EquityTierLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EquityTierLimit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EquityTierLimit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Limit != 0 {
		i = encodeVarintEquityTierLimitConfig(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.UsdTncRequired.Size()
		i -= size
		if _, err := m.UsdTncRequired.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEquityTierLimitConfig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintEquityTierLimitConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovEquityTierLimitConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EquityTierLimitConfiguration) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ShortTermOrderEquityTiers) > 0 {
		for _, e := range m.ShortTermOrderEquityTiers {
			l = e.Size()
			n += 1 + l + sovEquityTierLimitConfig(uint64(l))
		}
	}
	if len(m.StatefulOrderEquityTiers) > 0 {
		for _, e := range m.StatefulOrderEquityTiers {
			l = e.Size()
			n += 1 + l + sovEquityTierLimitConfig(uint64(l))
		}
	}
	return n
}

func (m *EquityTierLimit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.UsdTncRequired.Size()
	n += 1 + l + sovEquityTierLimitConfig(uint64(l))
	if m.Limit != 0 {
		n += 1 + sovEquityTierLimitConfig(uint64(m.Limit))
	}
	return n
}

func sovEquityTierLimitConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEquityTierLimitConfig(x uint64) (n int) {
	return sovEquityTierLimitConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EquityTierLimitConfiguration) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEquityTierLimitConfig
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
			return fmt.Errorf("proto: EquityTierLimitConfiguration: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EquityTierLimitConfiguration: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShortTermOrderEquityTiers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquityTierLimitConfig
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
				return ErrInvalidLengthEquityTierLimitConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEquityTierLimitConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShortTermOrderEquityTiers = append(m.ShortTermOrderEquityTiers, EquityTierLimit{})
			if err := m.ShortTermOrderEquityTiers[len(m.ShortTermOrderEquityTiers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatefulOrderEquityTiers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquityTierLimitConfig
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
				return ErrInvalidLengthEquityTierLimitConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEquityTierLimitConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatefulOrderEquityTiers = append(m.StatefulOrderEquityTiers, EquityTierLimit{})
			if err := m.StatefulOrderEquityTiers[len(m.StatefulOrderEquityTiers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEquityTierLimitConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEquityTierLimitConfig
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
func (m *EquityTierLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEquityTierLimitConfig
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
			return fmt.Errorf("proto: EquityTierLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EquityTierLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsdTncRequired", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquityTierLimitConfig
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
				return ErrInvalidLengthEquityTierLimitConfig
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEquityTierLimitConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UsdTncRequired.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEquityTierLimitConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEquityTierLimitConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEquityTierLimitConfig
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
func skipEquityTierLimitConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEquityTierLimitConfig
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
					return 0, ErrIntOverflowEquityTierLimitConfig
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
					return 0, ErrIntOverflowEquityTierLimitConfig
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
				return 0, ErrInvalidLengthEquityTierLimitConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEquityTierLimitConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEquityTierLimitConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEquityTierLimitConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEquityTierLimitConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEquityTierLimitConfig = fmt.Errorf("proto: unexpected end of group")
)
