// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/Transaction.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// A single signed transaction, including all its signatures. The SignatureList will have a Signature for each Key in the transaction, either explicit or implicit, in the order that they appear in the transaction. For example, a CryptoTransfer will first have a Signature corresponding to the Key for the paying account, followed by a Signature corresponding to the Key for each account that is sending or receiving cryptocurrency in the transfer. Each Transaction should not have more than 50 levels.
// The SignatureList field is deprecated and succeeded by SignatureMap.
type Transaction struct {
	// Types that are valid to be assigned to BodyData:
	//	*Transaction_Body
	//	*Transaction_BodyBytes
	BodyData             isTransaction_BodyData `protobuf_oneof:"bodyData"`
	Sigs                 *SignatureList         `protobuf:"bytes,2,opt,name=sigs,proto3" json:"sigs,omitempty"` // Deprecated: Do not use.
	SigMap               *SignatureMap          `protobuf:"bytes,3,opt,name=sigMap,proto3" json:"sigMap,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c62dd3117cbb23, []int{0}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

type isTransaction_BodyData interface {
	isTransaction_BodyData()
}

type Transaction_Body struct {
	Body *TransactionBody `protobuf:"bytes,1,opt,name=body,proto3,oneof"`
}

type Transaction_BodyBytes struct {
	BodyBytes []byte `protobuf:"bytes,4,opt,name=bodyBytes,proto3,oneof"`
}

func (*Transaction_Body) isTransaction_BodyData() {}

func (*Transaction_BodyBytes) isTransaction_BodyData() {}

func (m *Transaction) GetBodyData() isTransaction_BodyData {
	if m != nil {
		return m.BodyData
	}
	return nil
}

// Deprecated: Do not use.
func (m *Transaction) GetBody() *TransactionBody {
	if x, ok := m.GetBodyData().(*Transaction_Body); ok {
		return x.Body
	}
	return nil
}

func (m *Transaction) GetBodyBytes() []byte {
	if x, ok := m.GetBodyData().(*Transaction_BodyBytes); ok {
		return x.BodyBytes
	}
	return nil
}

// Deprecated: Do not use.
func (m *Transaction) GetSigs() *SignatureList {
	if m != nil {
		return m.Sigs
	}
	return nil
}

func (m *Transaction) GetSigMap() *SignatureMap {
	if m != nil {
		return m.SigMap
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Transaction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Transaction_Body)(nil),
		(*Transaction_BodyBytes)(nil),
	}
}

func init() {
	proto.RegisterType((*Transaction)(nil), "proto.Transaction")
}

func init() {
	proto.RegisterFile("proto/Transaction.proto", fileDescriptor_f4c62dd3117cbb23)
}

var fileDescriptor_f4c62dd3117cbb23 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8d, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0x97, 0x39, 0x87, 0x66, 0x9e, 0xaa, 0xcc, 0x50, 0x41, 0x86, 0xa7, 0x82, 0x2e, 0x15,
	0xfd, 0x07, 0xc1, 0x43, 0x0f, 0x0e, 0xa4, 0xee, 0xe4, 0xed, 0x6b, 0x1b, 0x92, 0x20, 0x6b, 0x42,
	0xbe, 0xec, 0xd0, 0x3f, 0xe8, 0xef, 0x92, 0x26, 0x43, 0x65, 0x9e, 0x5e, 0xf8, 0x9e, 0xe7, 0x7d,
	0x3f, 0x7a, 0xed, 0xbc, 0x0d, 0xb6, 0xdc, 0x7a, 0xe8, 0x11, 0xda, 0x60, 0x6c, 0xcf, 0xe3, 0x25,
	0x3b, 0x8d, 0x91, 0x2f, 0x13, 0x17, 0x80, 0xa6, 0xdd, 0x0e, 0x4e, 0x62, 0xc2, 0xf9, 0xcd, 0xbf,
	0x9e, 0xb0, 0xdd, 0x90, 0xe0, 0xdd, 0x17, 0xa1, 0x8b, 0x3f, 0x24, 0x7b, 0xa4, 0xb3, 0xc6, 0x76,
	0x03, 0x23, 0x2b, 0x52, 0x2c, 0x9e, 0x96, 0xc9, 0xe2, 0x47, 0x5d, 0x31, 0x65, 0xa4, 0x9a, 0xd4,
	0xd1, 0xcc, 0x6e, 0xe9, 0xf9, 0x98, 0x62, 0x08, 0x12, 0xd9, 0x6c, 0x45, 0x8a, 0x8b, 0x6a, 0x52,
	0xff, 0x9e, 0xb2, 0x07, 0x3a, 0x43, 0xa3, 0x90, 0x4d, 0xe3, 0xe2, 0xd5, 0x61, 0xf1, 0xdd, 0xa8,
	0x1e, 0xc2, 0xde, 0xcb, 0x57, 0x83, 0x61, 0xdc, 0xab, 0xa3, 0x95, 0xdd, 0xd3, 0x39, 0x1a, 0xb5,
	0x01, 0xc7, 0x4e, 0xa2, 0x7f, 0x79, 0xec, 0x6f, 0xc0, 0xd5, 0x07, 0x45, 0x50, 0x7a, 0x36, 0xfe,
	0x79, 0x81, 0x00, 0xa2, 0xa2, 0x79, 0x6b, 0x77, 0x5c, 0xcb, 0x4e, 0x7a, 0xe0, 0x1a, 0x50, 0x2b,
	0x0f, 0x4e, 0xa7, 0xfa, 0x1b, 0xf9, 0x28, 0x94, 0x09, 0x7a, 0xdf, 0xf0, 0xd6, 0xee, 0xca, 0x1f,
	0x5a, 0x26, 0x7d, 0x8d, 0xdd, 0xe7, 0x5a, 0xd9, 0x32, 0xba, 0xcd, 0x3c, 0xc6, 0xf3, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x31, 0xf7, 0x0b, 0xc2, 0x70, 0x01, 0x00, 0x00,
}
