// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: thorchain/v1/x/thorchain/types/msg_observed_txin.proto

package types

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

type MsgObservedTxIn struct {
	Txs    ObservedTxs                                   `protobuf:"bytes,1,rep,name=txs,proto3,castrepeated=ObservedTxs" json:"txs"`
	Signer github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=signer,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"signer,omitempty"`
}

func (m *MsgObservedTxIn) Reset()         { *m = MsgObservedTxIn{} }
func (m *MsgObservedTxIn) String() string { return proto.CompactTextString(m) }
func (*MsgObservedTxIn) ProtoMessage()    {}
func (*MsgObservedTxIn) Descriptor() ([]byte, []int) {
	return fileDescriptor_16bb506bbe99dc37, []int{0}
}
func (m *MsgObservedTxIn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgObservedTxIn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgObservedTxIn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgObservedTxIn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgObservedTxIn.Merge(m, src)
}
func (m *MsgObservedTxIn) XXX_Size() int {
	return m.Size()
}
func (m *MsgObservedTxIn) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgObservedTxIn.DiscardUnknown(m)
}

var xxx_messageInfo_MsgObservedTxIn proto.InternalMessageInfo

func (m *MsgObservedTxIn) GetTxs() ObservedTxs {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *MsgObservedTxIn) GetSigner() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Signer
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgObservedTxIn)(nil), "types.MsgObservedTxIn")
}

func init() {
	proto.RegisterFile("thorchain/v1/x/thorchain/types/msg_observed_txin.proto", fileDescriptor_16bb506bbe99dc37)
}

var fileDescriptor_16bb506bbe99dc37 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x2b, 0xc9, 0xc8, 0x2f,
	0x4a, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x33, 0xd4, 0xaf, 0xd0, 0x47, 0x70, 0x4b, 0x2a, 0x0b,
	0x52, 0x8b, 0xf5, 0x73, 0x8b, 0xd3, 0xe3, 0xf3, 0x93, 0x8a, 0x53, 0x8b, 0xca, 0x52, 0x53, 0xe2,
	0x4b, 0x2a, 0x32, 0xf3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x58, 0xc1, 0xd2, 0x52, 0xa6,
	0x04, 0xb4, 0x83, 0x48, 0x64, 0xfd, 0x10, 0xdd, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0xa6,
	0x3e, 0x88, 0x05, 0x11, 0x55, 0x9a, 0xc6, 0xc8, 0xc5, 0xef, 0x5b, 0x9c, 0xee, 0x0f, 0x55, 0x1e,
	0x52, 0xe1, 0x99, 0x27, 0x64, 0xc1, 0xc5, 0x5c, 0x52, 0x51, 0x2c, 0xc1, 0xa8, 0xc0, 0xac, 0xc1,
	0x6d, 0x24, 0xa8, 0x07, 0x36, 0x55, 0x0f, 0xa1, 0xc2, 0x49, 0xf8, 0xc4, 0x3d, 0x79, 0x86, 0x55,
	0xf7, 0xe5, 0xb9, 0x11, 0x62, 0xc5, 0x41, 0x20, 0x2d, 0x42, 0x9e, 0x5c, 0x6c, 0xc5, 0x99, 0xe9,
	0x79, 0xa9, 0x45, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c, 0x4e, 0x86, 0xbf, 0xee, 0xc9, 0xeb, 0xa6,
	0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0xe7, 0x17, 0xe7, 0xe6, 0x17,
	0x43, 0x29, 0xdd, 0xe2, 0x94, 0x6c, 0x88, 0x93, 0xf5, 0x1c, 0x93, 0x93, 0x1d, 0x53, 0x52, 0x8a,
	0x52, 0x8b, 0x8b, 0x83, 0xa0, 0x06, 0x38, 0x79, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c,
	0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1,
	0x1c, 0x43, 0x94, 0x7e, 0x7a, 0x66, 0x49, 0x4e, 0x22, 0xc4, 0x40, 0x24, 0xbf, 0x67, 0xe4, 0x17,
	0xe5, 0xe5, 0xa7, 0xa4, 0x62, 0x06, 0x48, 0x12, 0x1b, 0xd8, 0xab, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd1, 0x48, 0x0a, 0x49, 0x78, 0x01, 0x00, 0x00,
}

func (m *MsgObservedTxIn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgObservedTxIn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgObservedTxIn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgObservedTxin(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Txs) > 0 {
		for iNdEx := len(m.Txs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Txs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMsgObservedTxin(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsgObservedTxin(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgObservedTxin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgObservedTxIn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for _, e := range m.Txs {
			l = e.Size()
			n += 1 + l + sovMsgObservedTxin(uint64(l))
		}
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgObservedTxin(uint64(l))
	}
	return n
}

func sovMsgObservedTxin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgObservedTxin(x uint64) (n int) {
	return sovMsgObservedTxin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgObservedTxIn) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgObservedTxin
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
			return fmt.Errorf("proto: MsgObservedTxIn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgObservedTxIn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgObservedTxin
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
				return ErrInvalidLengthMsgObservedTxin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgObservedTxin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, ObservedTx{})
			if err := m.Txs[len(m.Txs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgObservedTxin
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
				return ErrInvalidLengthMsgObservedTxin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgObservedTxin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = append(m.Signer[:0], dAtA[iNdEx:postIndex]...)
			if m.Signer == nil {
				m.Signer = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgObservedTxin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMsgObservedTxin
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMsgObservedTxin
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
func skipMsgObservedTxin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgObservedTxin
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
					return 0, ErrIntOverflowMsgObservedTxin
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
					return 0, ErrIntOverflowMsgObservedTxin
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
				return 0, ErrInvalidLengthMsgObservedTxin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgObservedTxin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgObservedTxin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgObservedTxin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgObservedTxin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgObservedTxin = fmt.Errorf("proto: unexpected end of group")
)