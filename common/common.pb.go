// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: thorchain/v1/common/common.proto

package common

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Asset struct {
	Chain  Chain  `protobuf:"bytes,1,opt,name=chain,proto3,casttype=Chain" json:"chain,omitempty"`
	Symbol Symbol `protobuf:"bytes,2,opt,name=symbol,proto3,casttype=Symbol" json:"symbol,omitempty"`
	Ticker Ticker `protobuf:"bytes,3,opt,name=ticker,proto3,casttype=Ticker" json:"ticker,omitempty"`
	Synth  bool   `protobuf:"varint,4,opt,name=synth,proto3" json:"synth,omitempty"`
}

func (m *Asset) Reset()      { *m = Asset{} }
func (*Asset) ProtoMessage() {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_89e4c42e2083e3ed, []int{0}
}
func (m *Asset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return m.Size()
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

type Coin struct {
	Asset    Asset                                   `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset"`
	Amount   github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"amount"`
	Decimals int64                                   `protobuf:"varint,3,opt,name=decimals,proto3" json:"decimals,omitempty"`
}

func (m *Coin) Reset()      { *m = Coin{} }
func (*Coin) ProtoMessage() {}
func (*Coin) Descriptor() ([]byte, []int) {
	return fileDescriptor_89e4c42e2083e3ed, []int{1}
}
func (m *Coin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Coin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Coin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Coin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coin.Merge(m, src)
}
func (m *Coin) XXX_Size() int {
	return m.Size()
}
func (m *Coin) XXX_DiscardUnknown() {
	xxx_messageInfo_Coin.DiscardUnknown(m)
}

var xxx_messageInfo_Coin proto.InternalMessageInfo

// PubKeySet contains two pub keys , secp256k1 and ed25519
type PubKeySet struct {
	Secp256k1 PubKey `protobuf:"bytes,1,opt,name=secp256k1,proto3,casttype=PubKey" json:"secp256k1,omitempty"`
	Ed25519   PubKey `protobuf:"bytes,2,opt,name=ed25519,proto3,casttype=PubKey" json:"ed25519,omitempty"`
}

func (m *PubKeySet) Reset()      { *m = PubKeySet{} }
func (*PubKeySet) ProtoMessage() {}
func (*PubKeySet) Descriptor() ([]byte, []int) {
	return fileDescriptor_89e4c42e2083e3ed, []int{2}
}
func (m *PubKeySet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PubKeySet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PubKeySet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PubKeySet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubKeySet.Merge(m, src)
}
func (m *PubKeySet) XXX_Size() int {
	return m.Size()
}
func (m *PubKeySet) XXX_DiscardUnknown() {
	xxx_messageInfo_PubKeySet.DiscardUnknown(m)
}

var xxx_messageInfo_PubKeySet proto.InternalMessageInfo

type Tx struct {
	ID          TxID    `protobuf:"bytes,1,opt,name=id,proto3,casttype=TxID" json:"id,omitempty"`
	Chain       Chain   `protobuf:"bytes,2,opt,name=chain,proto3,casttype=Chain" json:"chain,omitempty"`
	FromAddress Address `protobuf:"bytes,3,opt,name=from_address,json=fromAddress,proto3,casttype=Address" json:"from_address,omitempty"`
	ToAddress   Address `protobuf:"bytes,4,opt,name=to_address,json=toAddress,proto3,casttype=Address" json:"to_address,omitempty"`
	Coins       Coins   `protobuf:"bytes,5,rep,name=coins,proto3,castrepeated=Coins" json:"coins"`
	Gas         Gas     `protobuf:"bytes,6,rep,name=gas,proto3,castrepeated=Gas" json:"gas"`
	Memo        string  `protobuf:"bytes,7,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (m *Tx) Reset()      { *m = Tx{} }
func (*Tx) ProtoMessage() {}
func (*Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_89e4c42e2083e3ed, []int{3}
}
func (m *Tx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tx.Merge(m, src)
}
func (m *Tx) XXX_Size() int {
	return m.Size()
}
func (m *Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_Tx.DiscardUnknown(m)
}

var xxx_messageInfo_Tx proto.InternalMessageInfo

type Fee struct {
	Coins      Coins                                   `protobuf:"bytes,1,rep,name=coins,proto3,castrepeated=Coins" json:"coins"`
	PoolDeduct github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,2,opt,name=pool_deduct,json=poolDeduct,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"pool_deduct"`
}

func (m *Fee) Reset()      { *m = Fee{} }
func (*Fee) ProtoMessage() {}
func (*Fee) Descriptor() ([]byte, []int) {
	return fileDescriptor_89e4c42e2083e3ed, []int{4}
}
func (m *Fee) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Fee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Fee.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Fee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fee.Merge(m, src)
}
func (m *Fee) XXX_Size() int {
	return m.Size()
}
func (m *Fee) XXX_DiscardUnknown() {
	xxx_messageInfo_Fee.DiscardUnknown(m)
}

var xxx_messageInfo_Fee proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Asset)(nil), "common.Asset")
	proto.RegisterType((*Coin)(nil), "common.Coin")
	proto.RegisterType((*PubKeySet)(nil), "common.PubKeySet")
	proto.RegisterType((*Tx)(nil), "common.Tx")
	proto.RegisterType((*Fee)(nil), "common.Fee")
}

func init() { proto.RegisterFile("thorchain/v1/common/common.proto", fileDescriptor_89e4c42e2083e3ed) }

var fileDescriptor_89e4c42e2083e3ed = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0xfd, 0xbb, 0xf5, 0x73, 0xcb, 0x70, 0xaa, 0x84, 0xd5, 0xc1, 0x8e, 0xac, 0x4a, 0x14,
	0x10, 0xb1, 0x12, 0x14, 0x24, 0xc6, 0xa6, 0x15, 0x55, 0x61, 0xa9, 0xae, 0x61, 0x61, 0xa9, 0x1c,
	0xfb, 0x48, 0xac, 0xc4, 0xbe, 0x28, 0x77, 0x41, 0xcd, 0xc6, 0xc4, 0xc0, 0xc4, 0x02, 0x7f, 0x04,
	0x7f, 0x49, 0xc6, 0x8e, 0x15, 0x83, 0xa1, 0xce, 0x7f, 0x91, 0x09, 0x9d, 0x7d, 0x69, 0x90, 0x00,
	0x09, 0x31, 0xf9, 0xde, 0xf7, 0x7d, 0xde, 0xbb, 0xef, 0xdd, 0xf3, 0x41, 0x83, 0x0f, 0xe9, 0x34,
	0x1e, 0x46, 0x69, 0x1e, 0xbe, 0x6b, 0x85, 0x31, 0xcd, 0x32, 0x9a, 0xcb, 0x4f, 0x73, 0x32, 0xa5,
	0x9c, 0x22, 0xab, 0x8e, 0xf6, 0xf7, 0x06, 0x74, 0x40, 0x2b, 0x29, 0x14, 0xab, 0x3a, 0x1b, 0x7c,
	0x50, 0xc1, 0x3c, 0x62, 0x8c, 0x70, 0xe4, 0x83, 0x59, 0xf5, 0x71, 0xd5, 0x86, 0x7a, 0x68, 0x77,
	0xed, 0x55, 0xe1, 0x9b, 0xc7, 0x42, 0xc0, 0xb5, 0x8e, 0x02, 0xb0, 0xd8, 0x3c, 0xeb, 0xd3, 0xb1,
	0xab, 0x55, 0x04, 0xac, 0x0a, 0xdf, 0xba, 0xa8, 0x14, 0x2c, 0x33, 0x82, 0xe1, 0x69, 0x3c, 0x22,
	0x53, 0x57, 0xdf, 0x30, 0xbd, 0x4a, 0xc1, 0x32, 0x83, 0xf6, 0xc0, 0x64, 0xf3, 0x9c, 0x0f, 0x5d,
	0xa3, 0xa1, 0x1e, 0x6e, 0xe3, 0x3a, 0x08, 0xbe, 0xa8, 0x60, 0x1c, 0xd3, 0x34, 0x47, 0x0f, 0xc1,
	0x8c, 0x84, 0xa1, 0xca, 0x87, 0xd3, 0xde, 0x6d, 0xca, 0xd3, 0x54, 0x2e, 0xbb, 0xc6, 0xa2, 0xf0,
	0x15, 0x5c, 0x13, 0xe8, 0x14, 0xac, 0x28, 0xa3, 0xb3, 0x9c, 0x4b, 0x47, 0xa1, 0x48, 0x7e, 0x2b,
	0xfc, 0x07, 0x83, 0x94, 0x0f, 0x67, 0x7d, 0x51, 0x19, 0xc6, 0x94, 0x65, 0x94, 0xc9, 0xcf, 0x13,
	0x96, 0x8c, 0x42, 0x3e, 0x9f, 0x10, 0xd6, 0x7c, 0x9d, 0xe6, 0x1c, 0xcb, 0x72, 0xb4, 0x0f, 0xdb,
	0x09, 0x89, 0xd3, 0x2c, 0x1a, 0xb3, 0xca, 0xb8, 0x8e, 0xef, 0xe2, 0x60, 0x08, 0xf6, 0xf9, 0xac,
	0xff, 0x8a, 0xcc, 0x2f, 0x08, 0x47, 0x1d, 0xb0, 0x19, 0x89, 0x27, 0xed, 0xce, 0xb3, 0x51, 0x4b,
	0x5e, 0xd4, 0xfd, 0xb2, 0xf0, 0xed, 0x8b, 0xb5, 0x28, 0xce, 0x5b, 0xe3, 0x78, 0x43, 0xa2, 0x03,
	0xd8, 0x22, 0x49, 0xbb, 0xd3, 0x69, 0x3d, 0xff, 0xf5, 0xee, 0x24, 0xb7, 0x4e, 0x05, 0x9f, 0x35,
	0xd0, 0x7a, 0x57, 0xc8, 0x03, 0x2d, 0x4d, 0x64, 0xf3, 0x7b, 0x65, 0xe1, 0x6b, 0x67, 0x27, 0xab,
	0xc2, 0x37, 0x7a, 0x57, 0x67, 0x27, 0x58, 0x4b, 0x93, 0xcd, 0xa0, 0xb4, 0xbf, 0x0c, 0xaa, 0x09,
	0x3b, 0x6f, 0xa7, 0x34, 0xbb, 0x8c, 0x92, 0x64, 0x4a, 0x18, 0x93, 0xa3, 0x70, 0x56, 0x85, 0xbf,
	0x75, 0x54, 0x4b, 0xd8, 0x11, 0x80, 0x0c, 0xd0, 0x23, 0x00, 0x4e, 0xef, 0x68, 0xe3, 0x77, 0xda,
	0xe6, 0x74, 0xcd, 0xb6, 0xc0, 0x8c, 0x69, 0x9a, 0x33, 0xd7, 0x6c, 0xe8, 0x87, 0x4e, 0x7b, 0x67,
	0x3d, 0x1d, 0x31, 0xba, 0xee, 0xae, 0xb8, 0xff, 0xaf, 0xdf, 0x7d, 0x53, 0x44, 0x0c, 0xd7, 0x24,
	0x7a, 0x0c, 0xfa, 0x20, 0x62, 0xae, 0xf5, 0x87, 0x02, 0x47, 0x16, 0xe8, 0xa7, 0x11, 0xc3, 0x82,
	0x42, 0x08, 0x8c, 0x8c, 0x64, 0xd4, 0xdd, 0x12, 0x2e, 0x70, 0xb5, 0x0e, 0x3e, 0xaa, 0xa0, 0xbf,
	0x20, 0x64, 0xb3, 0xb7, 0xfa, 0xcf, 0x7b, 0x9f, 0x83, 0x33, 0xa1, 0x74, 0x7c, 0x99, 0x90, 0x64,
	0x16, 0xff, 0xf7, 0x6f, 0x02, 0xa2, 0xc7, 0x49, 0xd5, 0xa2, 0xfb, 0x72, 0x71, 0xeb, 0x29, 0x37,
	0xb7, 0x9e, 0xf2, 0xbe, 0xf4, 0x94, 0x45, 0xe9, 0xa9, 0xd7, 0xa5, 0xa7, 0xfe, 0x28, 0x3d, 0xf5,
	0xd3, 0xd2, 0x53, 0xae, 0x97, 0x9e, 0x72, 0xb3, 0xf4, 0x94, 0x37, 0x07, 0x83, 0x94, 0x8f, 0xa3,
	0xba, 0xf5, 0xe6, 0x85, 0x8a, 0x55, 0x4e, 0x13, 0x22, 0x1f, 0x68, 0xdf, 0xaa, 0xde, 0xe0, 0xd3,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x41, 0xd3, 0x47, 0x89, 0xc5, 0x03, 0x00, 0x00,
}

func (m *Asset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Asset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Asset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Synth {
		i--
		if m.Synth {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Ticker) > 0 {
		i -= len(m.Ticker)
		copy(dAtA[i:], m.Ticker)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Ticker)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Coin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Coin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Coin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Decimals != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCommon(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Asset.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCommon(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PubKeySet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PubKeySet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PubKeySet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ed25519) > 0 {
		i -= len(m.Ed25519)
		copy(dAtA[i:], m.Ed25519)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Ed25519)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Secp256k1) > 0 {
		i -= len(m.Secp256k1)
		copy(dAtA[i:], m.Secp256k1)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Secp256k1)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Tx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Tx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Memo) > 0 {
		i -= len(m.Memo)
		copy(dAtA[i:], m.Memo)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Memo)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Gas) > 0 {
		for iNdEx := len(m.Gas) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Gas[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommon(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommon(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Fee) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fee) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Fee) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.PoolDeduct.Size()
		i -= size
		if _, err := m.PoolDeduct.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCommon(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommon(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Asset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Ticker)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.Synth {
		n += 2
	}
	return n
}

func (m *Coin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Asset.Size()
	n += 1 + l + sovCommon(uint64(l))
	l = m.Amount.Size()
	n += 1 + l + sovCommon(uint64(l))
	if m.Decimals != 0 {
		n += 1 + sovCommon(uint64(m.Decimals))
	}
	return n
}

func (m *PubKeySet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Secp256k1)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Ed25519)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *Tx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	if len(m.Gas) > 0 {
		for _, e := range m.Gas {
			l = e.Size()
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	l = len(m.Memo)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *Fee) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	l = m.PoolDeduct.Size()
	n += 1 + l + sovCommon(uint64(l))
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Asset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: Asset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Asset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = Chain(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = Symbol(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ticker = Ticker(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Synth", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
			m.Synth = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
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
func (m *Coin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: Coin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Coin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Asset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
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
func (m *PubKeySet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: PubKeySet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PubKeySet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secp256k1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secp256k1 = PubKey(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ed25519", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ed25519 = PubKey(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
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
func (m *Tx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: Tx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = TxID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = Chain(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = Address(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = Address(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gas = append(m.Gas, Coin{})
			if err := m.Gas[len(m.Gas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Memo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
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
func (m *Fee) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: Fee: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fee: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolDeduct", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PoolDeduct.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)
