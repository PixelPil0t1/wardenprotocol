// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: warden/async/v1beta1/ve.proto

package v1beta1

import (
	fmt "fmt"
	types "github.com/cometbft/cometbft/abci/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type AsyncInjectedTx struct {
	// All the vote extensions gathered for this block.
	//
	// FIXME: slinky also does that, so technically we're duplicating information
	// and wasting block space.
	ExtendedVotesInfo []types.ExtendedVoteInfo `protobuf:"bytes,1,rep,name=extended_votes_info,json=extendedVotesInfo,proto3" json:"extended_votes_info"`
}

func (m *AsyncInjectedTx) Reset()         { *m = AsyncInjectedTx{} }
func (m *AsyncInjectedTx) String() string { return proto.CompactTextString(m) }
func (*AsyncInjectedTx) ProtoMessage()    {}
func (*AsyncInjectedTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3e24cf4461cf67a, []int{0}
}
func (m *AsyncInjectedTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AsyncInjectedTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AsyncInjectedTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AsyncInjectedTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AsyncInjectedTx.Merge(m, src)
}
func (m *AsyncInjectedTx) XXX_Size() int {
	return m.Size()
}
func (m *AsyncInjectedTx) XXX_DiscardUnknown() {
	xxx_messageInfo_AsyncInjectedTx.DiscardUnknown(m)
}

var xxx_messageInfo_AsyncInjectedTx proto.InternalMessageInfo

func (m *AsyncInjectedTx) GetExtendedVotesInfo() []types.ExtendedVoteInfo {
	if m != nil {
		return m.ExtendedVotesInfo
	}
	return nil
}

// A vote extension coming from a validator. It contains results and votes for
// some futures.
type AsyncVoteExtension struct {
	Results []*VEResultItem `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Votes   []*VEVoteItem   `protobuf:"bytes,2,rep,name=votes,proto3" json:"votes,omitempty"`
}

func (m *AsyncVoteExtension) Reset()         { *m = AsyncVoteExtension{} }
func (m *AsyncVoteExtension) String() string { return proto.CompactTextString(m) }
func (*AsyncVoteExtension) ProtoMessage()    {}
func (*AsyncVoteExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3e24cf4461cf67a, []int{1}
}
func (m *AsyncVoteExtension) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AsyncVoteExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AsyncVoteExtension.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AsyncVoteExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AsyncVoteExtension.Merge(m, src)
}
func (m *AsyncVoteExtension) XXX_Size() int {
	return m.Size()
}
func (m *AsyncVoteExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_AsyncVoteExtension.DiscardUnknown(m)
}

var xxx_messageInfo_AsyncVoteExtension proto.InternalMessageInfo

func (m *AsyncVoteExtension) GetResults() []*VEResultItem {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *AsyncVoteExtension) GetVotes() []*VEVoteItem {
	if m != nil {
		return m.Votes
	}
	return nil
}

type VEResultItem struct {
	FutureId uint64 `protobuf:"varint,1,opt,name=future_id,json=futureId,proto3" json:"future_id,omitempty"`
	Output   []byte `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (m *VEResultItem) Reset()         { *m = VEResultItem{} }
func (m *VEResultItem) String() string { return proto.CompactTextString(m) }
func (*VEResultItem) ProtoMessage()    {}
func (*VEResultItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3e24cf4461cf67a, []int{2}
}
func (m *VEResultItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VEResultItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VEResultItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VEResultItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VEResultItem.Merge(m, src)
}
func (m *VEResultItem) XXX_Size() int {
	return m.Size()
}
func (m *VEResultItem) XXX_DiscardUnknown() {
	xxx_messageInfo_VEResultItem.DiscardUnknown(m)
}

var xxx_messageInfo_VEResultItem proto.InternalMessageInfo

func (m *VEResultItem) GetFutureId() uint64 {
	if m != nil {
		return m.FutureId
	}
	return 0
}

func (m *VEResultItem) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

type VEVoteItem struct {
	FutureId uint64         `protobuf:"varint,1,opt,name=future_id,json=futureId,proto3" json:"future_id,omitempty"`
	Vote     FutureVoteType `protobuf:"varint,2,opt,name=vote,proto3,enum=warden.async.v1beta1.FutureVoteType" json:"vote,omitempty"`
}

func (m *VEVoteItem) Reset()         { *m = VEVoteItem{} }
func (m *VEVoteItem) String() string { return proto.CompactTextString(m) }
func (*VEVoteItem) ProtoMessage()    {}
func (*VEVoteItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3e24cf4461cf67a, []int{3}
}
func (m *VEVoteItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VEVoteItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VEVoteItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VEVoteItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VEVoteItem.Merge(m, src)
}
func (m *VEVoteItem) XXX_Size() int {
	return m.Size()
}
func (m *VEVoteItem) XXX_DiscardUnknown() {
	xxx_messageInfo_VEVoteItem.DiscardUnknown(m)
}

var xxx_messageInfo_VEVoteItem proto.InternalMessageInfo

func (m *VEVoteItem) GetFutureId() uint64 {
	if m != nil {
		return m.FutureId
	}
	return 0
}

func (m *VEVoteItem) GetVote() FutureVoteType {
	if m != nil {
		return m.Vote
	}
	return FutureVoteType_VOTE_TYPE_UNSPECIFIED
}

func init() {
	proto.RegisterType((*AsyncInjectedTx)(nil), "warden.async.v1beta1.AsyncInjectedTx")
	proto.RegisterType((*AsyncVoteExtension)(nil), "warden.async.v1beta1.AsyncVoteExtension")
	proto.RegisterType((*VEResultItem)(nil), "warden.async.v1beta1.VEResultItem")
	proto.RegisterType((*VEVoteItem)(nil), "warden.async.v1beta1.VEVoteItem")
}

func init() { proto.RegisterFile("warden/async/v1beta1/ve.proto", fileDescriptor_e3e24cf4461cf67a) }

var fileDescriptor_e3e24cf4461cf67a = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4f, 0x6b, 0xe2, 0x40,
	0x14, 0xcf, 0xb8, 0xae, 0xbb, 0x3b, 0x2b, 0xbb, 0x6c, 0x56, 0x4a, 0x50, 0x9a, 0xc6, 0xd0, 0x83,
	0x97, 0xce, 0xa0, 0x85, 0xd2, 0x43, 0x2f, 0xb5, 0x28, 0xe4, 0x1a, 0xc4, 0x42, 0x2f, 0x36, 0x7f,
	0x46, 0x1b, 0xd1, 0x99, 0x90, 0x4c, 0xac, 0x7e, 0x85, 0x9e, 0xfa, 0xb1, 0x3c, 0x7a, 0xec, 0xa9,
	0x14, 0xfd, 0x22, 0x65, 0x26, 0x23, 0x4a, 0xb1, 0xbd, 0xcd, 0x7b, 0xbf, 0x7f, 0x8f, 0x37, 0x0f,
	0x1e, 0x3f, 0x7a, 0x49, 0x48, 0x28, 0xf6, 0xd2, 0x05, 0x0d, 0xf0, 0xac, 0xe9, 0x13, 0xee, 0x35,
	0xf1, 0x8c, 0xa0, 0x38, 0x61, 0x9c, 0xe9, 0x95, 0x1c, 0x46, 0x12, 0x46, 0x0a, 0xae, 0x56, 0x46,
	0x6c, 0xc4, 0x24, 0x01, 0x8b, 0x57, 0xce, 0xad, 0xd6, 0x38, 0xa1, 0x21, 0x49, 0xa6, 0x11, 0xe5,
	0xd8, 0xf3, 0x83, 0x08, 0xf3, 0x45, 0x4c, 0x52, 0x05, 0xd6, 0x0f, 0xe6, 0x0c, 0x33, 0x9e, 0x25,
	0x2a, 0xcb, 0x1e, 0xc3, 0xbf, 0xd7, 0x02, 0x75, 0xe8, 0x98, 0x04, 0x9c, 0x84, 0xbd, 0xb9, 0x7e,
	0x0b, 0xff, 0x93, 0xb9, 0xb4, 0x0d, 0x07, 0x33, 0xc6, 0x49, 0x3a, 0x88, 0xe8, 0x90, 0x19, 0xc0,
	0xfa, 0xd6, 0xf8, 0xdd, 0xaa, 0xa3, 0x5d, 0x20, 0x12, 0x81, 0xa8, 0xa3, 0xb8, 0x7d, 0xc6, 0x89,
	0x43, 0x87, 0xac, 0x5d, 0x5c, 0xbe, 0x9e, 0x68, 0xee, 0x3f, 0xb2, 0xd7, 0x4f, 0x05, 0x60, 0x3f,
	0x01, 0xa8, 0xcb, 0x30, 0xd1, 0x92, 0xb2, 0x34, 0x62, 0x54, 0xbf, 0x82, 0x3f, 0x12, 0x92, 0x66,
	0x13, 0x9e, 0xaa, 0x0c, 0x1b, 0x1d, 0x5a, 0x00, 0xea, 0x77, 0x5c, 0x49, 0x73, 0x38, 0x99, 0xba,
	0x5b, 0x89, 0x7e, 0x01, 0xbf, 0xcb, 0x21, 0x8d, 0x82, 0xd4, 0x5a, 0x9f, 0x69, 0xe5, 0x78, 0x42,
	0x99, 0xd3, 0xed, 0x1b, 0x58, 0xde, 0x37, 0xd4, 0x6b, 0xf0, 0x57, 0xbe, 0x98, 0x41, 0x14, 0x1a,
	0xc0, 0x02, 0x8d, 0xa2, 0xfb, 0x33, 0x6f, 0x38, 0xa1, 0x7e, 0x04, 0x4b, 0x2c, 0xe3, 0x71, 0xc6,
	0x8d, 0x82, 0x05, 0x1a, 0x65, 0x57, 0x55, 0x76, 0x00, 0xe1, 0xce, 0xf9, 0x6b, 0x8b, 0x4b, 0x58,
	0x14, 0xc1, 0xd2, 0xe0, 0x4f, 0xeb, 0xf4, 0xf0, 0x98, 0x5d, 0xc9, 0x16, 0x86, 0xbd, 0x45, 0x4c,
	0x5c, 0xa9, 0x68, 0xdf, 0x2f, 0xd7, 0x26, 0x58, 0xad, 0x4d, 0xf0, 0xb6, 0x36, 0xc1, 0xf3, 0xc6,
	0xd4, 0x56, 0x1b, 0x53, 0x7b, 0xd9, 0x98, 0xda, 0x5d, 0x77, 0x14, 0xf1, 0x87, 0xcc, 0x47, 0x01,
	0x9b, 0xe2, 0xdc, 0xef, 0x4c, 0xfe, 0x6a, 0xc0, 0x26, 0xaa, 0xfe, 0x50, 0xe2, 0xb9, 0xba, 0x05,
	0x79, 0x25, 0xdb, 0x8b, 0xf0, 0x4b, 0x92, 0x76, 0xfe, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x46, 0xd4,
	0x19, 0x03, 0x98, 0x02, 0x00, 0x00,
}

func (m *AsyncInjectedTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AsyncInjectedTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AsyncInjectedTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExtendedVotesInfo) > 0 {
		for iNdEx := len(m.ExtendedVotesInfo) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExtendedVotesInfo[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVe(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AsyncVoteExtension) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AsyncVoteExtension) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AsyncVoteExtension) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVe(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Results) > 0 {
		for iNdEx := len(m.Results) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Results[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVe(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *VEResultItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VEResultItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VEResultItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Output) > 0 {
		i -= len(m.Output)
		copy(dAtA[i:], m.Output)
		i = encodeVarintVe(dAtA, i, uint64(len(m.Output)))
		i--
		dAtA[i] = 0x12
	}
	if m.FutureId != 0 {
		i = encodeVarintVe(dAtA, i, uint64(m.FutureId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *VEVoteItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VEVoteItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VEVoteItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Vote != 0 {
		i = encodeVarintVe(dAtA, i, uint64(m.Vote))
		i--
		dAtA[i] = 0x10
	}
	if m.FutureId != 0 {
		i = encodeVarintVe(dAtA, i, uint64(m.FutureId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVe(dAtA []byte, offset int, v uint64) int {
	offset -= sovVe(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AsyncInjectedTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ExtendedVotesInfo) > 0 {
		for _, e := range m.ExtendedVotesInfo {
			l = e.Size()
			n += 1 + l + sovVe(uint64(l))
		}
	}
	return n
}

func (m *AsyncVoteExtension) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, e := range m.Results {
			l = e.Size()
			n += 1 + l + sovVe(uint64(l))
		}
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovVe(uint64(l))
		}
	}
	return n
}

func (m *VEResultItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FutureId != 0 {
		n += 1 + sovVe(uint64(m.FutureId))
	}
	l = len(m.Output)
	if l > 0 {
		n += 1 + l + sovVe(uint64(l))
	}
	return n
}

func (m *VEVoteItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FutureId != 0 {
		n += 1 + sovVe(uint64(m.FutureId))
	}
	if m.Vote != 0 {
		n += 1 + sovVe(uint64(m.Vote))
	}
	return n
}

func sovVe(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVe(x uint64) (n int) {
	return sovVe(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AsyncInjectedTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVe
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
			return fmt.Errorf("proto: AsyncInjectedTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AsyncInjectedTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtendedVotesInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVe
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
				return ErrInvalidLengthVe
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVe
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtendedVotesInfo = append(m.ExtendedVotesInfo, types.ExtendedVoteInfo{})
			if err := m.ExtendedVotesInfo[len(m.ExtendedVotesInfo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVe(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVe
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
func (m *AsyncVoteExtension) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVe
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
			return fmt.Errorf("proto: AsyncVoteExtension: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AsyncVoteExtension: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVe
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
				return ErrInvalidLengthVe
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVe
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, &VEResultItem{})
			if err := m.Results[len(m.Results)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVe
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
				return ErrInvalidLengthVe
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVe
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, &VEVoteItem{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVe(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVe
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
func (m *VEResultItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVe
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
			return fmt.Errorf("proto: VEResultItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VEResultItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FutureId", wireType)
			}
			m.FutureId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVe
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FutureId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Output", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVe
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
				return ErrInvalidLengthVe
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthVe
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Output = append(m.Output[:0], dAtA[iNdEx:postIndex]...)
			if m.Output == nil {
				m.Output = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVe(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVe
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
func (m *VEVoteItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVe
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
			return fmt.Errorf("proto: VEVoteItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VEVoteItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FutureId", wireType)
			}
			m.FutureId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVe
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FutureId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vote", wireType)
			}
			m.Vote = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVe
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vote |= FutureVoteType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVe(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVe
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
func skipVe(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVe
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
					return 0, ErrIntOverflowVe
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
					return 0, ErrIntOverflowVe
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
				return 0, ErrInvalidLengthVe
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVe
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVe
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVe        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVe          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVe = fmt.Errorf("proto: unexpected end of group")
)
