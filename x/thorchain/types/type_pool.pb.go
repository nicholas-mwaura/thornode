// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: thorchain/v1/x/thorchain/types/type_pool.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	common "gitlab.com/thorchain/thornode/common"
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

// |    State    | Swap | Add   | Withdraw  | Refunding |
// | ----------- | ---- | ----- | --------- | --------- |
// | `staged`    | no   | yes   | yes       | Refund Invalid Add/Remove Liquidity && all Swaps |
// | `available` | yes  | yes   | yes       | Refund Invalid Tx |
// | `suspended` | no   | no    | no        | Refund all |
type PoolStatus int32

const (
	PoolStatus_UnknownPoolStatus PoolStatus = 0
	PoolStatus_Available         PoolStatus = 1
	PoolStatus_Staged            PoolStatus = 2
	PoolStatus_Suspended         PoolStatus = 3
)

var PoolStatus_name = map[int32]string{
	0: "UnknownPoolStatus",
	1: "Available",
	2: "Staged",
	3: "Suspended",
}

var PoolStatus_value = map[string]int32{
	"UnknownPoolStatus": 0,
	"Available":         1,
	"Staged":            2,
	"Suspended":         3,
}

func (x PoolStatus) String() string {
	return proto.EnumName(PoolStatus_name, int32(x))
}

func (PoolStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_90e6f303297a9bb4, []int{0}
}

type Pool struct {
	BalanceRune         github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,1,opt,name=balance_rune,json=balanceRune,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"balance_rune"`
	BalanceAsset        github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,2,opt,name=balance_asset,json=balanceAsset,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"balance_asset"`
	Asset               common.Asset                            `protobuf:"bytes,3,opt,name=asset,proto3" json:"asset"`
	LPUnits             github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,4,opt,name=LP_units,json=lPUnits,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"LP_units"`
	Status              PoolStatus                              `protobuf:"varint,5,opt,name=status,proto3,enum=types.PoolStatus" json:"status,omitempty"`
	Decimals            int64                                   `protobuf:"varint,6,opt,name=decimals,proto3" json:"decimals,omitempty"`
	SynthUnits          github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,7,opt,name=synth_units,json=synthUnits,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"synth_units"`
	PendingInboundRune  github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,8,opt,name=pending_inbound_rune,json=pendingInboundRune,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"pending_inbound_rune"`
	PendingInboundAsset github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,9,opt,name=pending_inbound_asset,json=pendingInboundAsset,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"pending_inbound_asset"`
}

func (m *Pool) Reset()      { *m = Pool{} }
func (*Pool) ProtoMessage() {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_90e6f303297a9bb4, []int{0}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("types.PoolStatus", PoolStatus_name, PoolStatus_value)
	proto.RegisterType((*Pool)(nil), "types.Pool")
}

func init() {
	proto.RegisterFile("thorchain/v1/x/thorchain/types/type_pool.proto", fileDescriptor_90e6f303297a9bb4)
}

var fileDescriptor_90e6f303297a9bb4 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xbf, 0x6e, 0xdb, 0x30,
	0x10, 0xc6, 0xa5, 0xf8, 0x4f, 0xec, 0x73, 0x5d, 0x38, 0x6c, 0x02, 0x08, 0x1e, 0x14, 0xa1, 0x4b,
	0x9d, 0x02, 0x95, 0xd0, 0xf4, 0x09, 0xe2, 0xad, 0x45, 0x0b, 0x18, 0x72, 0xbd, 0x74, 0x31, 0x28,
	0x89, 0x90, 0x88, 0x48, 0x3c, 0xc3, 0xa4, 0xd2, 0x66, 0xeb, 0xd8, 0xb1, 0x8f, 0xe5, 0x31, 0x63,
	0xd0, 0x21, 0x68, 0xec, 0x17, 0x29, 0x44, 0x2a, 0x75, 0xd3, 0x6c, 0x5a, 0x44, 0xf2, 0xee, 0xf8,
	0xbb, 0x4f, 0xfc, 0x70, 0xe0, 0xab, 0x0c, 0xd7, 0x71, 0x46, 0xb9, 0x08, 0xae, 0xde, 0x06, 0xdf,
	0x82, 0xfd, 0x51, 0x5d, 0xaf, 0x98, 0xd4, 0xdf, 0xe5, 0x0a, 0x31, 0xf7, 0x57, 0x6b, 0x54, 0x48,
	0x3a, 0x3a, 0x3c, 0xf6, 0x1e, 0x5d, 0x8b, 0xb1, 0x28, 0x50, 0xd4, 0x8b, 0x29, 0x1c, 0x1f, 0xa7,
	0x98, 0xa2, 0xde, 0x06, 0xd5, 0xce, 0x44, 0x5f, 0xfe, 0xe8, 0x40, 0x7b, 0x86, 0x98, 0x93, 0x10,
	0x9e, 0x45, 0x34, 0xa7, 0x22, 0x66, 0xcb, 0x75, 0x29, 0x98, 0x63, 0x7b, 0xf6, 0xa4, 0x3f, 0x0d,
	0x36, 0x77, 0xa7, 0xd6, 0xaf, 0xbb, 0xd3, 0x57, 0x29, 0x57, 0x59, 0x19, 0xf9, 0x31, 0x16, 0x41,
	0x8c, 0xb2, 0x40, 0x59, 0x2f, 0x6f, 0x64, 0x72, 0x69, 0x94, 0xf9, 0x0b, 0x2e, 0x54, 0x38, 0xa8,
	0x21, 0x61, 0x29, 0x18, 0xf9, 0x0c, 0xc3, 0x07, 0x26, 0x95, 0x92, 0x29, 0xe7, 0xa0, 0x19, 0xf4,
	0x41, 0xd9, 0x45, 0x05, 0x21, 0x67, 0xd0, 0x31, 0xb4, 0x96, 0x67, 0x4f, 0x06, 0xe7, 0x43, 0xbf,
	0xfe, 0x4d, 0x9d, 0x9d, 0xb6, 0x2b, 0x78, 0x68, 0x2a, 0xc8, 0x07, 0xe8, 0x7d, 0x9c, 0x2d, 0x4b,
	0xc1, 0x95, 0x74, 0xda, 0xcd, 0x7a, 0x1f, 0xe6, 0xb3, 0x45, 0x75, 0x9f, 0x9c, 0x41, 0x57, 0x2a,
	0xaa, 0x4a, 0xe9, 0x74, 0x3c, 0x7b, 0xf2, 0xfc, 0xfc, 0xc8, 0x37, 0x55, 0xd5, 0xeb, 0xcd, 0x75,
	0x22, 0xac, 0x0b, 0xc8, 0x18, 0x7a, 0x09, 0x8b, 0x79, 0x41, 0x73, 0xe9, 0x74, 0x3d, 0x7b, 0xd2,
	0x0a, 0xff, 0x9e, 0xc9, 0x0c, 0x06, 0xf2, 0x5a, 0xa8, 0xac, 0x56, 0x75, 0xd8, 0x4c, 0x15, 0x68,
	0x86, 0x11, 0x46, 0xe1, 0x78, 0xc5, 0x44, 0xc2, 0x45, 0xba, 0xe4, 0x22, 0xc2, 0x52, 0x24, 0xc6,
	0xc1, 0x5e, 0x33, 0x34, 0xa9, 0x61, 0xef, 0x0d, 0x4b, 0x1b, 0x19, 0xc3, 0xc9, 0xff, 0x2d, 0x8c,
	0x05, 0xfd, 0x66, 0x3d, 0x5e, 0x3c, 0xee, 0xa1, 0x9d, 0x7b, 0xfd, 0x09, 0x60, 0xff, 0x96, 0xe4,
	0x04, 0x8e, 0x16, 0xe2, 0x52, 0xe0, 0x57, 0xb1, 0x0f, 0x8e, 0x2c, 0x32, 0x84, 0xfe, 0xc5, 0x15,
	0xe5, 0x39, 0x8d, 0x72, 0x36, 0xb2, 0x09, 0x40, 0x77, 0xae, 0x68, 0xca, 0x92, 0xd1, 0x41, 0x95,
	0x9a, 0x97, 0xb2, 0x22, 0xb3, 0x64, 0xd4, 0x9a, 0x2e, 0x36, 0xf7, 0xae, 0x75, 0x7b, 0xef, 0x5a,
	0xdf, 0xb7, 0xae, 0xb5, 0xd9, 0xba, 0xf6, 0xcd, 0xd6, 0xb5, 0x7f, 0x6f, 0x5d, 0xfb, 0xe7, 0xce,
	0xb5, 0x6e, 0x76, 0xae, 0x75, 0xbb, 0x73, 0xad, 0x2f, 0x41, 0xca, 0x55, 0x4e, 0x8d, 0xe4, 0x7f,
	0x46, 0x2d, 0xc3, 0xb5, 0xc0, 0x84, 0x3d, 0x9d, 0xbf, 0xa8, 0xab, 0xe7, 0xe6, 0xdd, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa3, 0x76, 0x38, 0xdf, 0xa8, 0x03, 0x00, 0x00,
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.PendingInboundAsset.Size()
		i -= size
		if _, err := m.PendingInboundAsset.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypePool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.PendingInboundRune.Size()
		i -= size
		if _, err := m.PendingInboundRune.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypePool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.SynthUnits.Size()
		i -= size
		if _, err := m.SynthUnits.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypePool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if m.Decimals != 0 {
		i = encodeVarintTypePool(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x30
	}
	if m.Status != 0 {
		i = encodeVarintTypePool(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	{
		size := m.LPUnits.Size()
		i -= size
		if _, err := m.LPUnits.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypePool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Asset.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypePool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.BalanceAsset.Size()
		i -= size
		if _, err := m.BalanceAsset.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypePool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.BalanceRune.Size()
		i -= size
		if _, err := m.BalanceRune.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypePool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypePool(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypePool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BalanceRune.Size()
	n += 1 + l + sovTypePool(uint64(l))
	l = m.BalanceAsset.Size()
	n += 1 + l + sovTypePool(uint64(l))
	l = m.Asset.Size()
	n += 1 + l + sovTypePool(uint64(l))
	l = m.LPUnits.Size()
	n += 1 + l + sovTypePool(uint64(l))
	if m.Status != 0 {
		n += 1 + sovTypePool(uint64(m.Status))
	}
	if m.Decimals != 0 {
		n += 1 + sovTypePool(uint64(m.Decimals))
	}
	l = m.SynthUnits.Size()
	n += 1 + l + sovTypePool(uint64(l))
	l = m.PendingInboundRune.Size()
	n += 1 + l + sovTypePool(uint64(l))
	l = m.PendingInboundAsset.Size()
	n += 1 + l + sovTypePool(uint64(l))
	return n
}

func sovTypePool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypePool(x uint64) (n int) {
	return sovTypePool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypePool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BalanceRune", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypePool
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
				return ErrInvalidLengthTypePool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BalanceRune.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BalanceAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypePool
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
				return ErrInvalidLengthTypePool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BalanceAsset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypePool
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
				return ErrInvalidLengthTypePool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Asset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LPUnits", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypePool
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
				return ErrInvalidLengthTypePool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LPUnits.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypePool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= PoolStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypePool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SynthUnits", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypePool
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
				return ErrInvalidLengthTypePool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SynthUnits.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingInboundRune", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypePool
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
				return ErrInvalidLengthTypePool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PendingInboundRune.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingInboundAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypePool
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
				return ErrInvalidLengthTypePool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypePool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PendingInboundAsset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypePool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypePool
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypePool
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
func skipTypePool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypePool
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
					return 0, ErrIntOverflowTypePool
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
					return 0, ErrIntOverflowTypePool
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
				return 0, ErrInvalidLengthTypePool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypePool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypePool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypePool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypePool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypePool = fmt.Errorf("proto: unexpected end of group")
)
