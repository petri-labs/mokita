// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/streamswap/v1/state.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Sale struct {
	// Destination for the earned token_in
	Treasury string `protobuf:"bytes,1,opt,name=treasury,proto3" json:"treasury,omitempty"`
	Id       uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// token_out is a token denom to be bootstraped. May be referred as base
	// currency, or a sale token.
	TokenOut string `protobuf:"bytes,3,opt,name=token_out,json=tokenOut,proto3" json:"token_out,omitempty"`
	// token_in is a token denom used to buy sale tokens (`token_out`). May be
	// referred as quote_currency or payment token.
	TokenIn string `protobuf:"bytes,4,opt,name=token_in,json=tokenIn,proto3" json:"token_in,omitempty"`
	// total number of `tokens_out` to be sold during the continuous sale.
	TokenOutSupply github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=token_out_supply,json=tokenOutSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"token_out_supply"`
	// start time when the token emission starts.
	StartTime time.Time `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	// end time when the token emission ends. Can't be bigger than start +
	// 139years (to avoid round overflow)
	EndTime time.Time `protobuf:"bytes,7,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time"`
	// Round number when the sale was last time updated.
	Round int64 `protobuf:"varint,8,opt,name=round,proto3" json:"round,omitempty"`
	// Last round of the Sale;
	EndRound int64 `protobuf:"varint,9,opt,name=end_round,json=endRound,proto3" json:"end_round,omitempty"`
	// amout of remaining token_out to sell
	OutRemaining github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=out_remaining,json=outRemaining,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"out_remaining"`
	// amount of token_out sold
	OutSold github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=out_sold,json=outSold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"out_sold"`
	// out token per share
	OutPerShare github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,12,opt,name=out_per_share,json=outPerShare,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"out_per_share"`
	// total amount of currently staked coins (token_in) but not spent coins.
	Staked github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,13,opt,name=staked,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"staked"`
	// total amount of earned coins (token_in)
	Income github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,14,opt,name=income,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"income"`
	// total amount of shares
	Shares github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,15,opt,name=shares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"shares"`
	// Name for the sale.
	Name string `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`
	// URL with sale and project details.
	Url string `protobuf:"bytes,21,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *Sale) Reset()         { *m = Sale{} }
func (m *Sale) String() string { return proto.CompactTextString(m) }
func (*Sale) ProtoMessage()    {}
func (*Sale) Descriptor() ([]byte, []int) {
	return fileDescriptor_7602494b1425cb83, []int{0}
}
func (m *Sale) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sale) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sale.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Sale) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sale.Merge(m, src)
}
func (m *Sale) XXX_Size() int {
	return m.Size()
}
func (m *Sale) XXX_DiscardUnknown() {
	xxx_messageInfo_Sale.DiscardUnknown(m)
}

var xxx_messageInfo_Sale proto.InternalMessageInfo

// UserPosition represents user account in a sale
type UserPosition struct {
	Shares github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=shares,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"shares"`
	// total number of currently staked tokens
	Staked github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=staked,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"staked"`
	// last token/share ratio
	OutPerShare github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=out_per_share,json=outPerShare,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"out_per_share"`
	// amount of token_in spent
	Spent github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=spent,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"spent"`
	// Amount of accumulated, not withdrawn, purchased tokens (token_out)
	Purchased github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=purchased,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"purchased"`
}

func (m *UserPosition) Reset()         { *m = UserPosition{} }
func (m *UserPosition) String() string { return proto.CompactTextString(m) }
func (*UserPosition) ProtoMessage()    {}
func (*UserPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_7602494b1425cb83, []int{1}
}
func (m *UserPosition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserPosition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPosition.Merge(m, src)
}
func (m *UserPosition) XXX_Size() int {
	return m.Size()
}
func (m *UserPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPosition.DiscardUnknown(m)
}

var xxx_messageInfo_UserPosition proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Sale)(nil), "osmosis.streamswap.v1.Sale")
	proto.RegisterType((*UserPosition)(nil), "osmosis.streamswap.v1.UserPosition")
}

func init() { proto.RegisterFile("osmosis/streamswap/v1/state.proto", fileDescriptor_7602494b1425cb83) }

var fileDescriptor_7602494b1425cb83 = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcb, 0x6a, 0xdb, 0x40,
	0x14, 0x86, 0x2d, 0xdb, 0xf1, 0x65, 0x72, 0x69, 0x18, 0x12, 0x98, 0xb8, 0x20, 0xbb, 0x59, 0x14,
	0x6f, 0x22, 0xe1, 0x16, 0xba, 0x2d, 0xa4, 0xa5, 0x10, 0x28, 0x34, 0xc8, 0x29, 0x94, 0x6e, 0xc4,
	0xd8, 0x9a, 0xca, 0xc2, 0xd2, 0x8c, 0x98, 0x8b, 0x5b, 0xbf, 0x40, 0xd7, 0x79, 0x2c, 0x2f, 0xb3,
	0x2c, 0x5d, 0xa4, 0xad, 0xfd, 0x00, 0x7d, 0x85, 0x32, 0x33, 0x72, 0x6c, 0xe8, 0x2a, 0xf2, 0xca,
	0x73, 0x2e, 0xff, 0x77, 0x0e, 0xbf, 0xcc, 0x01, 0xcf, 0x98, 0xc8, 0x98, 0x48, 0x84, 0x2f, 0x24,
	0x27, 0x38, 0x13, 0x5f, 0x71, 0xee, 0xcf, 0x06, 0xbe, 0x90, 0x58, 0x12, 0x2f, 0xe7, 0x4c, 0x32,
	0x78, 0x5a, 0xb4, 0x78, 0x9b, 0x16, 0x6f, 0x36, 0xe8, 0x74, 0x63, 0xc6, 0xe2, 0x94, 0xf8, 0xa6,
	0x69, 0xa4, 0xbe, 0xf8, 0x32, 0xc9, 0x88, 0x90, 0x38, 0xcb, 0xad, 0xae, 0x73, 0x36, 0x36, 0xc2,
	0xd0, 0x44, 0xbe, 0x0d, 0x8a, 0xd2, 0x49, 0xcc, 0x62, 0x66, 0xf3, 0xfa, 0x65, 0xb3, 0xe7, 0x7f,
	0x1b, 0xa0, 0x3e, 0xc4, 0x29, 0x81, 0x1d, 0xd0, 0xd2, 0xa3, 0x84, 0xe2, 0x73, 0xe4, 0xf4, 0x9c,
	0x7e, 0x3b, 0x78, 0x88, 0xe1, 0x11, 0xa8, 0x26, 0x11, 0xaa, 0xf6, 0x9c, 0x7e, 0x3d, 0xa8, 0x26,
	0x11, 0x7c, 0x0a, 0xda, 0x92, 0x4d, 0x09, 0x0d, 0x99, 0x92, 0xa8, 0x56, 0x34, 0xeb, 0xc4, 0x07,
	0x25, 0xe1, 0x19, 0xb0, 0xef, 0x30, 0xa1, 0xa8, 0x6e, 0x6a, 0x4d, 0x13, 0x5f, 0x51, 0xf8, 0x09,
	0x1c, 0x3f, 0xe8, 0x42, 0xa1, 0xf2, 0x3c, 0x9d, 0xa3, 0x3d, 0xdd, 0x72, 0xe9, 0x2d, 0xee, 0xbb,
	0x95, 0x9f, 0xf7, 0xdd, 0xe7, 0x71, 0x22, 0x27, 0x6a, 0xe4, 0x8d, 0x59, 0x56, 0x6c, 0x5f, 0xfc,
	0x5c, 0x88, 0x68, 0xea, 0xcb, 0x79, 0x4e, 0x84, 0x77, 0x45, 0x65, 0x70, 0xb4, 0x1e, 0x37, 0x34,
	0x14, 0xf8, 0x06, 0x00, 0x21, 0x31, 0x97, 0xa1, 0x36, 0x04, 0x35, 0x7a, 0x4e, 0x7f, 0xff, 0x45,
	0xc7, 0xb3, 0x6e, 0x79, 0x6b, 0xb7, 0xbc, 0x9b, 0xb5, 0x5b, 0x97, 0x2d, 0x3d, 0xef, 0xf6, 0x57,
	0xd7, 0x09, 0xda, 0x46, 0xa7, 0x2b, 0xf0, 0x35, 0x68, 0x11, 0x1a, 0x59, 0x44, 0xf3, 0x11, 0x88,
	0x26, 0xa1, 0x91, 0x01, 0x9c, 0x80, 0x3d, 0xce, 0x14, 0x8d, 0x50, 0xab, 0xe7, 0xf4, 0x6b, 0x81,
	0x0d, 0xb4, 0x5b, 0x1a, 0x6b, 0x2b, 0x6d, 0x53, 0xd1, 0x73, 0x02, 0x53, 0x1c, 0x82, 0x43, 0x6d,
	0x06, 0x27, 0x19, 0x4e, 0x68, 0x42, 0x63, 0x04, 0x4a, 0xf9, 0x71, 0xc0, 0x94, 0x0c, 0xd6, 0x0c,
	0x78, 0x05, 0x5a, 0xc6, 0x61, 0x96, 0x46, 0x68, 0xbf, 0x14, 0xaf, 0xc9, 0x94, 0x1c, 0xb2, 0x34,
	0x82, 0x81, 0xdd, 0x2f, 0x27, 0x3c, 0x14, 0x13, 0xcc, 0x09, 0x3a, 0x28, 0xc5, 0xdb, 0x67, 0x4a,
	0x5e, 0x13, 0x3e, 0xd4, 0x08, 0xf8, 0x0e, 0x34, 0x84, 0xc4, 0x53, 0x12, 0xa1, 0xc3, 0x52, 0xb0,
	0x42, 0xad, 0x39, 0x09, 0x1d, 0xb3, 0x8c, 0xa0, 0xa3, 0x72, 0x1c, 0xab, 0x36, 0xfb, 0xe8, 0xc5,
	0x04, 0x7a, 0x52, 0x72, 0x1f, 0xa3, 0x86, 0x10, 0xd4, 0x29, 0xce, 0x08, 0x3a, 0x31, 0xff, 0x7a,
	0xf3, 0x86, 0xc7, 0xa0, 0xa6, 0x78, 0x8a, 0x4e, 0x4d, 0x4a, 0x3f, 0xcf, 0xbf, 0xd7, 0xc0, 0xc1,
	0x47, 0x41, 0xf8, 0x35, 0x13, 0x89, 0x4c, 0x18, 0xdd, 0x1a, 0xef, 0xec, 0x34, 0x7e, 0x63, 0x6b,
	0x75, 0x27, 0x5b, 0xff, 0xfb, 0xe4, 0xb5, 0xdd, 0x3f, 0xf9, 0x5b, 0xb0, 0x27, 0x72, 0x42, 0xa5,
	0xbd, 0x08, 0x8f, 0x66, 0x59, 0x31, 0x7c, 0x0f, 0xda, 0xb9, 0xe2, 0xe3, 0x09, 0x16, 0x24, 0x2a,
	0x79, 0x38, 0x36, 0x80, 0xcb, 0x9b, 0xc5, 0x1f, 0xb7, 0xb2, 0x58, 0xba, 0xce, 0xdd, 0xd2, 0x75,
	0x7e, 0x2f, 0x5d, 0xe7, 0x76, 0xe5, 0x56, 0xee, 0x56, 0x6e, 0xe5, 0xc7, 0xca, 0xad, 0x7c, 0x7e,
	0xb5, 0x05, 0x2c, 0x8e, 0xf1, 0x45, 0x8a, 0x47, 0x62, 0x1d, 0xf8, 0xb3, 0xc1, 0xc0, 0xff, 0xb6,
	0x7d, 0xc2, 0xcd, 0x90, 0x51, 0xc3, 0x9c, 0x8a, 0x97, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x89,
	0x74, 0xc2, 0x26, 0xe5, 0x05, 0x00, 0x00,
}

func (m *Sale) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sale) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Sale) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintState(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xaa
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintState(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa2
	}
	{
		size := m.Shares.Size()
		i -= size
		if _, err := m.Shares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x7a
	{
		size := m.Income.Size()
		i -= size
		if _, err := m.Income.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x72
	{
		size := m.Staked.Size()
		i -= size
		if _, err := m.Staked.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	{
		size := m.OutPerShare.Size()
		i -= size
		if _, err := m.OutPerShare.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	{
		size := m.OutSold.Size()
		i -= size
		if _, err := m.OutSold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.OutRemaining.Size()
		i -= size
		if _, err := m.OutRemaining.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.EndRound != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.EndRound))
		i--
		dAtA[i] = 0x48
	}
	if m.Round != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x40
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintState(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x3a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintState(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x32
	{
		size := m.TokenOutSupply.Size()
		i -= size
		if _, err := m.TokenOutSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.TokenIn) > 0 {
		i -= len(m.TokenIn)
		copy(dAtA[i:], m.TokenIn)
		i = encodeVarintState(dAtA, i, uint64(len(m.TokenIn)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TokenOut) > 0 {
		i -= len(m.TokenOut)
		copy(dAtA[i:], m.TokenOut)
		i = encodeVarintState(dAtA, i, uint64(len(m.TokenOut)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Treasury) > 0 {
		i -= len(m.Treasury)
		copy(dAtA[i:], m.Treasury)
		i = encodeVarintState(dAtA, i, uint64(len(m.Treasury)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserPosition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserPosition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserPosition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Purchased.Size()
		i -= size
		if _, err := m.Purchased.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Spent.Size()
		i -= size
		if _, err := m.Spent.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.OutPerShare.Size()
		i -= size
		if _, err := m.OutPerShare.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Staked.Size()
		i -= size
		if _, err := m.Staked.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Shares.Size()
		i -= size
		if _, err := m.Shares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintState(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Sale) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Treasury)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovState(uint64(m.Id))
	}
	l = len(m.TokenOut)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.TokenIn)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = m.TokenOutSupply.Size()
	n += 1 + l + sovState(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovState(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovState(uint64(l))
	if m.Round != 0 {
		n += 1 + sovState(uint64(m.Round))
	}
	if m.EndRound != 0 {
		n += 1 + sovState(uint64(m.EndRound))
	}
	l = m.OutRemaining.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.OutSold.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.OutPerShare.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.Staked.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.Income.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.Shares.Size()
	n += 1 + l + sovState(uint64(l))
	l = len(m.Name)
	if l > 0 {
		n += 2 + l + sovState(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 2 + l + sovState(uint64(l))
	}
	return n
}

func (m *UserPosition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Shares.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.Staked.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.OutPerShare.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.Spent.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.Purchased.Size()
	n += 1 + l + sovState(uint64(l))
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Sale) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: Sale: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sale: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Treasury", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Treasury = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenOut = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenIn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOutSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenOutSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndRound", wireType)
			}
			m.EndRound = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndRound |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutRemaining", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OutRemaining.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutSold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OutSold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutPerShare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OutPerShare.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staked", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Staked.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Income", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Income.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Shares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 20:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 21:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *UserPosition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: UserPosition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserPosition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Shares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staked", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Staked.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutPerShare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OutPerShare.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Purchased", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Purchased.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)