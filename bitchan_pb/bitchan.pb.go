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
	BitchanMessage
	NotifiedHash
	Node
	StoredValue
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

type DataType int32

const (
	DataType_BLOCK_HEADER DataType = 0
	DataType_BLOCK_BODY   DataType = 1
	DataType_TRANSACTION  DataType = 2
	DataType_POST         DataType = 3
)

var DataType_name = map[int32]string{
	0: "BLOCK_HEADER",
	1: "BLOCK_BODY",
	2: "TRANSACTION",
	3: "POST",
}
var DataType_value = map[string]int32{
	"BLOCK_HEADER": 0,
	"BLOCK_BODY":   1,
	"TRANSACTION":  2,
	"POST":         3,
}

func (x DataType) String() string {
	return proto.EnumName(DataType_name, int32(x))
}
func (DataType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type BlockHeader struct {
	PreviousBlockHeaderHash []byte  `protobuf:"bytes,1,opt,name=previous_block_header_hash,proto3" json:"previous_block_header_hash,omitempty"`
	BodyHash                []byte  `protobuf:"bytes,2,opt,name=body_hash,proto3" json:"body_hash,omitempty"`
	Timestamp               int64   `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
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

func (m *BlockHeader) GetTimestamp() int64 {
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
	TransactionHashes [][]byte `protobuf:"bytes,1,rep,name=transaction_hashes,proto3" json:"transaction_hashes,omitempty"`
}

func (m *BlockBody) Reset()                    { *m = BlockBody{} }
func (m *BlockBody) String() string            { return proto.CompactTextString(m) }
func (*BlockBody) ProtoMessage()               {}
func (*BlockBody) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BlockBody) GetTransactionHashes() [][]byte {
	if m != nil {
		return m.TransactionHashes
	}
	return nil
}

type Transaction struct {
	TransactionInputs  []*TransactionInput  `protobuf:"bytes,1,rep,name=transaction_inputs" json:"transaction_inputs,omitempty"`
	TransactionOutputs []*TransactionOutput `protobuf:"bytes,2,rep,name=transaction_outputs" json:"transaction_outputs,omitempty"`
	BoardId            string               `protobuf:"bytes,3,opt,name=board_id" json:"board_id,omitempty"`
	PostHash           []byte               `protobuf:"bytes,4,opt,name=post_hash,proto3" json:"post_hash,omitempty"`
	// Set only when the post is successive message in the thread.
	ThreadTransactionHash []byte `protobuf:"bytes,5,opt,name=thread_transaction_hash,proto3" json:"thread_transaction_hash,omitempty"`
	// "sage"
	Downvoted bool `protobuf:"varint,6,opt,name=downvoted" json:"downvoted,omitempty"`
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

func (m *Transaction) GetBoardId() string {
	if m != nil {
		return m.BoardId
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

func (m *Transaction) GetDownvoted() bool {
	if m != nil {
		return m.Downvoted
	}
	return false
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
	Timestamp int64  `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
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

func (m *Post) GetTimestamp() int64 {
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

// Kademlia inspired.
// https://pdos.csail.mit.edu/~petar/papers/maymounkov-kademlia-lncs.pdf
// http://www.bittorrent.org/beps/bep_0005.html
type BitchanMessage struct {
	OwnNodeId []byte `protobuf:"bytes,1,opt,name=own_node_id,proto3" json:"own_node_id,omitempty"`
	// FIND_NODE or FIND_VALUE targed identifier.
	TargetId []byte `protobuf:"bytes,2,opt,name=target_id,proto3" json:"target_id,omitempty"`
	// Return stored_value instead of nodes. (FIND_VALUE)
	FindValue bool `protobuf:"varint,3,opt,name=find_value" json:"find_value,omitempty"`
	// Response for FIND_NODE or FIND_VALUE
	Nodes [][]byte `protobuf:"bytes,4,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// STORE
	StoredValue *StoredValue `protobuf:"bytes,5,opt,name=stored_value" json:"stored_value,omitempty"`
	// inv of Bitcoin
	NotifiedHashes []*NotifiedHash `protobuf:"bytes,6,rep,name=notified_hashes" json:"notified_hashes,omitempty"`
}

func (m *BitchanMessage) Reset()                    { *m = BitchanMessage{} }
func (m *BitchanMessage) String() string            { return proto.CompactTextString(m) }
func (*BitchanMessage) ProtoMessage()               {}
func (*BitchanMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *BitchanMessage) GetOwnNodeId() []byte {
	if m != nil {
		return m.OwnNodeId
	}
	return nil
}

func (m *BitchanMessage) GetTargetId() []byte {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *BitchanMessage) GetFindValue() bool {
	if m != nil {
		return m.FindValue
	}
	return false
}

func (m *BitchanMessage) GetNodes() [][]byte {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *BitchanMessage) GetStoredValue() *StoredValue {
	if m != nil {
		return m.StoredValue
	}
	return nil
}

func (m *BitchanMessage) GetNotifiedHashes() []*NotifiedHash {
	if m != nil {
		return m.NotifiedHashes
	}
	return nil
}

type NotifiedHash struct {
	DataType DataType `protobuf:"varint,1,opt,name=data_type,enum=bitchan_pb.DataType" json:"data_type,omitempty"`
	Hash     []byte   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *NotifiedHash) Reset()                    { *m = NotifiedHash{} }
func (m *NotifiedHash) String() string            { return proto.CompactTextString(m) }
func (*NotifiedHash) ProtoMessage()               {}
func (*NotifiedHash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *NotifiedHash) GetDataType() DataType {
	if m != nil {
		return m.DataType
	}
	return DataType_BLOCK_HEADER
}

func (m *NotifiedHash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Node struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	// 160 bit kademlia ID
	NodeId []byte `protobuf:"bytes,2,opt,name=node_id,proto3" json:"node_id,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Node) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

type StoredValue struct {
	Data     []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	DataType DataType `protobuf:"varint,2,opt,name=data_type,enum=bitchan_pb.DataType" json:"data_type,omitempty"`
}

func (m *StoredValue) Reset()                    { *m = StoredValue{} }
func (m *StoredValue) String() string            { return proto.CompactTextString(m) }
func (*StoredValue) ProtoMessage()               {}
func (*StoredValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *StoredValue) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *StoredValue) GetDataType() DataType {
	if m != nil {
		return m.DataType
	}
	return DataType_BLOCK_HEADER
}

func init() {
	proto.RegisterType((*BlockHeader)(nil), "bitchan_pb.BlockHeader")
	proto.RegisterType((*BlockBody)(nil), "bitchan_pb.BlockBody")
	proto.RegisterType((*Transaction)(nil), "bitchan_pb.Transaction")
	proto.RegisterType((*TransactionInput)(nil), "bitchan_pb.TransactionInput")
	proto.RegisterType((*TransactionOutput)(nil), "bitchan_pb.TransactionOutput")
	proto.RegisterType((*Post)(nil), "bitchan_pb.Post")
	proto.RegisterType((*BitchanMessage)(nil), "bitchan_pb.BitchanMessage")
	proto.RegisterType((*NotifiedHash)(nil), "bitchan_pb.NotifiedHash")
	proto.RegisterType((*Node)(nil), "bitchan_pb.Node")
	proto.RegisterType((*StoredValue)(nil), "bitchan_pb.StoredValue")
	proto.RegisterEnum("bitchan_pb.DataType", DataType_name, DataType_value)
}

func init() { proto.RegisterFile("bitchan.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x51, 0x4f, 0xdb, 0x3c,
	0x14, 0xfd, 0xd2, 0x86, 0xd2, 0xdc, 0x84, 0x36, 0x18, 0x24, 0xf2, 0xa1, 0xef, 0xd3, 0xba, 0xbc,
	0x10, 0x4d, 0x1a, 0xd2, 0xd8, 0xcb, 0x34, 0xed, 0x85, 0x52, 0x34, 0xd0, 0xb6, 0x16, 0x41, 0x35,
	0x69, 0x4f, 0x91, 0x13, 0xbb, 0xd4, 0xa2, 0xb5, 0xa3, 0xd8, 0x81, 0x75, 0x2f, 0xfc, 0xb2, 0xfd,
	0x93, 0xfd, 0x98, 0xc9, 0x76, 0x4b, 0xc3, 0x10, 0x7b, 0xf4, 0xbd, 0xe7, 0xe4, 0xde, 0x7b, 0xce,
	0x51, 0x60, 0x2b, 0x63, 0x2a, 0x9f, 0x62, 0x7e, 0x58, 0x94, 0x42, 0x09, 0x04, 0xcb, 0x67, 0x5a,
	0x64, 0xf1, 0x0f, 0xf0, 0xfb, 0x33, 0x91, 0xdf, 0x9c, 0x51, 0x4c, 0x68, 0x89, 0x62, 0xd8, 0x2f,
	0x4a, 0x7a, 0xcb, 0x44, 0x25, 0xd3, 0x4c, 0xd7, 0xd3, 0xa9, 0x69, 0xa4, 0x53, 0x2c, 0xa7, 0x91,
	0xd3, 0x73, 0x92, 0x00, 0x6d, 0x83, 0x97, 0x09, 0xb2, 0xb0, 0xa5, 0xc6, 0xaa, 0xa4, 0xd8, 0x9c,
	0x4a, 0x85, 0xe7, 0x45, 0xd4, 0xec, 0x39, 0x49, 0x13, 0x6d, 0xc1, 0x06, 0x17, 0x3c, 0xa7, 0x91,
	0xdb, 0x73, 0x12, 0x17, 0x21, 0x00, 0xc2, 0x26, 0x13, 0x96, 0x57, 0x33, 0xb5, 0x88, 0x36, 0x7a,
	0x4e, 0xd2, 0x88, 0x0f, 0xc0, 0x33, 0xb3, 0xfb, 0x82, 0x2c, 0xd0, 0x3e, 0x20, 0x55, 0x62, 0x2e,
	0x71, 0xae, 0x98, 0xe0, 0xe6, 0xe3, 0x54, 0x46, 0x4e, 0xaf, 0x99, 0x04, 0xf1, 0x2f, 0x07, 0xfc,
	0xf1, 0xba, 0x89, 0xde, 0x3d, 0xc6, 0x32, 0x5e, 0x54, 0xca, 0x62, 0xfd, 0xa3, 0xff, 0x0e, 0xd7,
	0xd7, 0x1d, 0xd6, 0x48, 0xe7, 0x1a, 0x84, 0xde, 0xc3, 0x4e, 0x9d, 0x29, 0x2a, 0x65, 0xa8, 0x0d,
	0x43, 0xfd, 0xff, 0x19, 0xea, 0xc8, 0xa0, 0x50, 0x08, 0xed, 0x4c, 0xe0, 0x92, 0xa4, 0x8c, 0x98,
	0x1b, 0x3d, 0x7d, 0x76, 0x21, 0xa4, 0xb2, 0x4a, 0xb8, 0x46, 0x89, 0x17, 0xb0, 0xa7, 0xa6, 0x25,
	0xc5, 0x24, 0xfd, 0xf3, 0x1a, 0x73, 0xb4, 0x91, 0x8a, 0x88, 0x3b, 0x7e, 0x2b, 0x14, 0x25, 0x51,
	0xab, 0xe7, 0x24, 0xed, 0xf8, 0x1e, 0xc2, 0x27, 0x8b, 0xbe, 0x84, 0x7f, 0x1f, 0x8c, 0x78, 0xf2,
	0x25, 0xeb, 0x43, 0xdd, 0xab, 0xc7, 0x72, 0x10, 0xfa, 0xdd, 0x18, 0xb3, 0xa1, 0xa7, 0x49, 0x76,
	0xcd, 0xb1, 0xaa, 0x4a, 0x6a, 0x96, 0x36, 0x0b, 0x14, 0x55, 0xc6, 0xf2, 0xf4, 0x86, 0x2e, 0xec,
	0xd2, 0xf1, 0x07, 0xd8, 0x7e, 0x7a, 0x6e, 0x07, 0x5a, 0x78, 0x2e, 0x2a, 0xae, 0xcc, 0xb8, 0x26,
	0xda, 0x83, 0x6e, 0x51, 0x65, 0x33, 0x4b, 0xac, 0x99, 0x1f, 0xdf, 0x83, 0x7b, 0x21, 0xa4, 0x42,
	0x01, 0xb8, 0x1c, 0xcf, 0xa9, 0x81, 0x7b, 0xfa, 0x35, 0xc7, 0x6c, 0x66, 0x30, 0x1e, 0xea, 0xc2,
	0x66, 0x2e, 0xb8, 0xa2, 0x5c, 0xad, 0xa5, 0x5b, 0x27, 0xc6, 0x35, 0x03, 0xfe, 0x7a, 0xb2, 0x15,
	0x6f, 0x17, 0x82, 0x95, 0xba, 0x4c, 0xcd, 0xa8, 0xd1, 0xcf, 0x8b, 0x7f, 0x3a, 0xd0, 0xe9, 0x5b,
	0xe7, 0xbe, 0x50, 0x29, 0xf1, 0x35, 0x45, 0x3b, 0xe0, 0x8b, 0x3b, 0x9e, 0x72, 0x41, 0xa8, 0xb6,
	0xeb, 0x21, 0xb8, 0x0a, 0x97, 0xd7, 0x54, 0xe9, 0x92, 0x0d, 0x2e, 0x02, 0x98, 0x30, 0x4e, 0xd2,
	0x5b, 0x3c, 0xab, 0xac, 0x40, 0x6d, 0x9b, 0x5c, 0x42, 0x65, 0xe4, 0xea, 0xf0, 0xa1, 0xd7, 0x10,
	0x48, 0x25, 0x4a, 0xba, 0x02, 0xe9, 0x4d, 0xfc, 0xa3, 0xbd, 0x7a, 0x56, 0xae, 0x4c, 0xff, 0xab,
	0x6e, 0xa3, 0x37, 0xd0, 0xe5, 0x42, 0xb1, 0x09, 0xa3, 0x64, 0x15, 0xe2, 0x96, 0x49, 0x57, 0x54,
	0x67, 0x0c, 0x97, 0x90, 0x33, 0x2c, 0xa7, 0xf1, 0x29, 0x04, 0xf5, 0x37, 0x3a, 0x00, 0x8f, 0x60,
	0x85, 0x53, 0xb5, 0x28, 0xac, 0x9a, 0x9d, 0xa3, 0xdd, 0x3a, 0x79, 0x80, 0x15, 0x1e, 0x2f, 0x0a,
	0xaa, 0x35, 0xae, 0xf9, 0x90, 0x80, 0x3b, 0x14, 0x84, 0x6a, 0xad, 0x31, 0x21, 0x25, 0x95, 0x72,
	0x69, 0x45, 0x17, 0x36, 0x57, 0x42, 0x58, 0xe4, 0x00, 0xfc, 0xfa, 0xca, 0x01, 0xb8, 0x7a, 0xde,
	0x52, 0xa5, 0x47, 0xd3, 0x1b, 0xcf, 0x4f, 0x7f, 0xf5, 0x11, 0xda, 0x0f, 0x9b, 0x84, 0x10, 0xf4,
	0x3f, 0x8f, 0x4e, 0x3e, 0xa5, 0x67, 0xa7, 0xc7, 0x83, 0xd3, 0xcb, 0xf0, 0x1f, 0xd4, 0x01, 0xb0,
	0x95, 0xfe, 0x68, 0xf0, 0x2d, 0x74, 0x50, 0x17, 0xfc, 0xf1, 0xe5, 0xf1, 0xf0, 0xea, 0xf8, 0x64,
	0x7c, 0x3e, 0x1a, 0x86, 0x0d, 0xd4, 0x06, 0xf7, 0x62, 0x74, 0x35, 0x0e, 0x9b, 0x59, 0xcb, 0xfc,
	0x96, 0xde, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x14, 0x27, 0x03, 0xa7, 0x04, 0x00, 0x00,
}
