// Code generated by protoc-gen-go.
// source: bitchan.proto
// DO NOT EDIT!

/*
Package bitchan_pb is a generated protocol buffer package.

It is generated from these files:
	bitchan.proto

It has these top-level messages:
	BlockHeader
	BlockBody
	Transaction
	TransactionInput
	TransactionOutput
	Post
*/
package bitchan_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BlockHeader struct {
	PreviousBlockHeaderHash []byte  `protobuf:"bytes,1,opt,name=previous_block_header_hash,proto3" json:"previous_block_header_hash,omitempty"`
	BodyHash                []byte  `protobuf:"bytes,2,opt,name=body_hash,proto3" json:"body_hash,omitempty"`
	Timestamp               uint64  `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Nonce                   uint64  `protobuf:"varint,4,opt,name=nonce" json:"nonce,omitempty"`
	Difficulty              float32 `protobuf:"fixed32,5,opt,name=difficulty" json:"difficulty,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BlockHeader) GetPreviousBlockHeaderHash() []byte {
	if m != nil {
		return m.PreviousBlockHeaderHash
	}
	return nil
}

func (m *BlockHeader) GetBodyHash() []byte {
	if m != nil {
		return m.BodyHash
	}
	return nil
}

func (m *BlockHeader) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *BlockHeader) GetDifficulty() float32 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

type BlockBody struct {
	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *BlockBody) Reset()                    { *m = BlockBody{} }
func (m *BlockBody) String() string            { return proto.CompactTextString(m) }
func (*BlockBody) ProtoMessage()               {}
func (*BlockBody) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BlockBody) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type Transaction struct {
	TransactionInputs  []*TransactionInput  `protobuf:"bytes,1,rep,name=transaction_inputs" json:"transaction_inputs,omitempty"`
	TransactionOutputs []*TransactionOutput `protobuf:"bytes,2,rep,name=transaction_outputs" json:"transaction_outputs,omitempty"`
	BoardIdentifier    string               `protobuf:"bytes,3,opt,name=board_identifier" json:"board_identifier,omitempty"`
	PostHash           []byte               `protobuf:"bytes,4,opt,name=post_hash,proto3" json:"post_hash,omitempty"`
	// Set only when the post is successive message in the thread.
	ThreadTransactionHash []byte `protobuf:"bytes,5,opt,name=thread_transaction_hash,proto3" json:"thread_transaction_hash,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Transaction) GetTransactionInputs() []*TransactionInput {
	if m != nil {
		return m.TransactionInputs
	}
	return nil
}

func (m *Transaction) GetTransactionOutputs() []*TransactionOutput {
	if m != nil {
		return m.TransactionOutputs
	}
	return nil
}

func (m *Transaction) GetBoardIdentifier() string {
	if m != nil {
		return m.BoardIdentifier
	}
	return ""
}

func (m *Transaction) GetPostHash() []byte {
	if m != nil {
		return m.PostHash
	}
	return nil
}

func (m *Transaction) GetThreadTransactionHash() []byte {
	if m != nil {
		return m.ThreadTransactionHash
	}
	return nil
}

type TransactionInput struct {
	PreviousTransactionHash  []byte `protobuf:"bytes,1,opt,name=previous_transaction_hash,proto3" json:"previous_transaction_hash,omitempty"`
	PreviousTransactionIndex int32  `protobuf:"varint,2,opt,name=previous_transaction_index" json:"previous_transaction_index,omitempty"`
	Signature                []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	PubicKey                 []byte `protobuf:"bytes,4,opt,name=pubic_key,proto3" json:"pubic_key,omitempty"`
}

func (m *TransactionInput) Reset()                    { *m = TransactionInput{} }
func (m *TransactionInput) String() string            { return proto.CompactTextString(m) }
func (*TransactionInput) ProtoMessage()               {}
func (*TransactionInput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TransactionInput) GetPreviousTransactionHash() []byte {
	if m != nil {
		return m.PreviousTransactionHash
	}
	return nil
}

func (m *TransactionInput) GetPreviousTransactionIndex() int32 {
	if m != nil {
		return m.PreviousTransactionIndex
	}
	return 0
}

func (m *TransactionInput) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *TransactionInput) GetPubicKey() []byte {
	if m != nil {
		return m.PubicKey
	}
	return nil
}

type TransactionOutput struct {
	Amount        int64  `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
	PublicKeyHash []byte `protobuf:"bytes,2,opt,name=public_key_hash,proto3" json:"public_key_hash,omitempty"`
}

func (m *TransactionOutput) Reset()                    { *m = TransactionOutput{} }
func (m *TransactionOutput) String() string            { return proto.CompactTextString(m) }
func (*TransactionOutput) ProtoMessage()               {}
func (*TransactionOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TransactionOutput) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TransactionOutput) GetPublicKeyHash() []byte {
	if m != nil {
		return m.PublicKeyHash
	}
	return nil
}

type Post struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Mail      string `protobuf:"bytes,2,opt,name=mail" json:"mail,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
	Timestamp uint64 `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	// Set only when the post is successive message in the thread.
	PreviousTransactionHash []byte `protobuf:"bytes,5,opt,name=previous_transaction_hash,proto3" json:"previous_transaction_hash,omitempty"`
	// Set only when the post is the first message in the thread.
	ThreadTitle string `protobuf:"bytes,6,opt,name=thread_title" json:"thread_title,omitempty"`
}

func (m *Post) Reset()                    { *m = Post{} }
func (m *Post) String() string            { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()               {}
func (*Post) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Post) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Post) GetMail() string {
	if m != nil {
		return m.Mail
	}
	return ""
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Post) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Post) GetPreviousTransactionHash() []byte {
	if m != nil {
		return m.PreviousTransactionHash
	}
	return nil
}

func (m *Post) GetThreadTitle() string {
	if m != nil {
		return m.ThreadTitle
	}
	return ""
}

func init() {
	proto.RegisterType((*BlockHeader)(nil), "bitchan_pb.BlockHeader")
	proto.RegisterType((*BlockBody)(nil), "bitchan_pb.BlockBody")
	proto.RegisterType((*Transaction)(nil), "bitchan_pb.Transaction")
	proto.RegisterType((*TransactionInput)(nil), "bitchan_pb.TransactionInput")
	proto.RegisterType((*TransactionOutput)(nil), "bitchan_pb.TransactionOutput")
	proto.RegisterType((*Post)(nil), "bitchan_pb.Post")
}

func init() { proto.RegisterFile("bitchan.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0xda, 0xb4, 0x28, 0xd3, 0x2c, 0xbb, 0x6b, 0x90, 0x1a, 0x10, 0x88, 0x92, 0x53, 0x2f,
	0xf4, 0x00, 0x17, 0xb4, 0xe2, 0xb4, 0x27, 0x38, 0xc1, 0x81, 0xbb, 0x65, 0xc7, 0x2e, 0xb1, 0x36,
	0xb1, 0x2d, 0x7b, 0x8c, 0x28, 0x97, 0xfd, 0x46, 0xfe, 0x08, 0x65, 0x42, 0x69, 0x50, 0xb5, 0x1c,
	0xc7, 0x7e, 0x6f, 0xe6, 0xcd, 0x7b, 0x03, 0x17, 0xd2, 0x60, 0xd3, 0x0a, 0xbb, 0xf3, 0xc1, 0xa1,
	0x63, 0xf0, 0xa7, 0xe4, 0x5e, 0xd6, 0x3f, 0x61, 0x75, 0xdb, 0xb9, 0xe6, 0xee, 0xa3, 0x16, 0x4a,
	0x07, 0x56, 0xc3, 0x73, 0x1f, 0xf4, 0x77, 0xe3, 0x52, 0xe4, 0x72, 0x78, 0xe7, 0x2d, 0x7d, 0xf0,
	0x56, 0xc4, 0xb6, 0xca, 0x36, 0xd9, 0xb6, 0x64, 0xd7, 0x50, 0x48, 0xa7, 0x0e, 0xe3, 0xd3, 0xec,
	0xf8, 0x84, 0xa6, 0xd7, 0x11, 0x45, 0xef, 0xab, 0xf9, 0x26, 0xdb, 0xe6, 0xec, 0x02, 0x16, 0xd6,
	0xd9, 0x46, 0x57, 0x39, 0x95, 0x0c, 0x40, 0x99, 0xfd, 0xde, 0x34, 0xa9, 0xc3, 0x43, 0xb5, 0xd8,
	0x64, 0xdb, 0x59, 0x7d, 0x03, 0x05, 0xcd, 0xbe, 0x75, 0xea, 0xc0, 0xde, 0x40, 0x89, 0x41, 0xd8,
	0x28, 0x1a, 0x34, 0xce, 0xc6, 0x2a, 0xdb, 0xcc, 0xb7, 0xab, 0xb7, 0xeb, 0xdd, 0x49, 0xeb, 0xee,
	0xeb, 0xe9, 0xbf, 0xfe, 0x95, 0xc1, 0x6a, 0x52, 0xb3, 0xf7, 0xc0, 0x26, 0x74, 0x6e, 0xac, 0x4f,
	0x78, 0x6c, 0xf2, 0xe2, 0x81, 0x26, 0x9f, 0x06, 0x10, 0xbb, 0x81, 0x27, 0x53, 0xa6, 0x4b, 0x48,
	0xd4, 0x19, 0x51, 0x5f, 0x3e, 0x40, 0xfd, 0x4c, 0x28, 0x56, 0xc1, 0x95, 0x74, 0x22, 0x28, 0x6e,
	0x94, 0xb6, 0x68, 0xf6, 0x46, 0x07, 0x5a, 0xbf, 0x18, 0x1c, 0xf1, 0x2e, 0xe2, 0x68, 0x52, 0x4e,
	0x26, 0xbd, 0x82, 0x35, 0xb6, 0x41, 0x0b, 0xc5, 0xa7, 0xf3, 0x08, 0x30, 0xf8, 0x51, 0xd6, 0xf7,
	0x70, 0x75, 0xa6, 0xee, 0x35, 0x3c, 0xfb, 0x1b, 0xc8, 0x19, 0x6d, 0xcc, 0x63, 0x9a, 0xd9, 0xbf,
	0x1e, 0x28, 0xfd, 0x83, 0x02, 0x5a, 0x0c, 0x72, 0xa2, 0xf9, 0x66, 0x05, 0xa6, 0xa0, 0x49, 0x21,
	0x65, 0xe6, 0x93, 0x34, 0x0d, 0xbf, 0xd3, 0x87, 0x51, 0x61, 0xfd, 0x01, 0xae, 0xcf, 0x77, 0x7c,
	0x0c, 0x4b, 0xd1, 0xbb, 0x64, 0x91, 0xc6, 0xcd, 0xd9, 0x1a, 0x2e, 0x7d, 0x92, 0xdd, 0x48, 0x9c,
	0x1c, 0x41, 0x7d, 0x0f, 0xf9, 0x17, 0x17, 0x91, 0x95, 0x90, 0x5b, 0xd1, 0x6b, 0x82, 0x17, 0x43,
	0xd5, 0x0b, 0xd3, 0x11, 0xa6, 0x60, 0x97, 0xf0, 0xa8, 0x71, 0x16, 0xb5, 0xc5, 0x93, 0x4f, 0xa7,
	0xcb, 0x19, 0x4f, 0xe5, 0xbf, 0x2b, 0x93, 0x53, 0xec, 0x29, 0x94, 0x47, 0x2b, 0x0d, 0x76, 0xba,
	0x5a, 0x0e, 0xbd, 0xe4, 0x92, 0xce, 0xfb, 0xdd, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xa4,
	0x70, 0xdf, 0xef, 0x02, 0x00, 0x00,
}
