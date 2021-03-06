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
	OwnPort   int32  `protobuf:"varint,2,opt,name=own_port" json:"own_port,omitempty"`
	// FIND_NODE or FIND_VALUE targed identifier.
	TargetId []byte `protobuf:"bytes,3,opt,name=target_id,proto3" json:"target_id,omitempty"`
	// Return stored_value instead of nodes. (FIND_VALUE)
	FindValue bool `protobuf:"varint,4,opt,name=find_value" json:"find_value,omitempty"`
	// Response for FIND_NODE or FIND_VALUE
	Nodes []*Node `protobuf:"bytes,5,rep,name=nodes" json:"nodes,omitempty"`
	// STORE
	StoredValue *StoredValue `protobuf:"bytes,6,opt,name=stored_value" json:"stored_value,omitempty"`
	// inv of Bitcoin
	NotifiedHashes []*NotifiedHash `protobuf:"bytes,7,rep,name=notified_hashes" json:"notified_hashes,omitempty"`
	// If set, needs pong.
	IsPing bool `protobuf:"varint,8,opt,name=is_ping" json:"is_ping,omitempty"`
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

func (m *BitchanMessage) GetOwnPort() int32 {
	if m != nil {
		return m.OwnPort
	}
	return 0
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

func (m *BitchanMessage) GetNodes() []*Node {
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

func (m *BitchanMessage) GetIsPing() bool {
	if m != nil {
		return m.IsPing
	}
	return false
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
	// 661 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x4e, 0xdb, 0x3c,
	0x18, 0xfe, 0xd2, 0xa6, 0xb4, 0x7d, 0x13, 0xda, 0x60, 0x90, 0xc8, 0x87, 0x36, 0xd1, 0xe5, 0x84,
	0x68, 0xd2, 0x90, 0xc6, 0x4e, 0xa6, 0x69, 0x27, 0x94, 0xa2, 0x81, 0xb6, 0x51, 0x04, 0xd5, 0xa4,
	0x1d, 0x45, 0x4e, 0xec, 0xb6, 0x16, 0xad, 0x1d, 0xc5, 0x0e, 0xac, 0x3b, 0xe1, 0x46, 0x77, 0x15,
	0xbb, 0x82, 0xc9, 0x36, 0xa5, 0xe9, 0x10, 0x3b, 0xcc, 0xeb, 0xe7, 0xfd, 0x79, 0x7e, 0x14, 0xd8,
	0x4c, 0x99, 0xca, 0xa6, 0x98, 0x1f, 0xe6, 0x85, 0x50, 0x02, 0xc1, 0xc3, 0x67, 0x92, 0xa7, 0xd1,
	0x4f, 0xf0, 0xfa, 0x33, 0x91, 0xdd, 0x9c, 0x51, 0x4c, 0x68, 0x81, 0x22, 0xd8, 0xcb, 0x0b, 0x7a,
	0xcb, 0x44, 0x29, 0x93, 0x54, 0xd7, 0x93, 0xa9, 0x79, 0x48, 0xa6, 0x58, 0x4e, 0x43, 0xa7, 0xe7,
	0xc4, 0x3e, 0xda, 0x82, 0x76, 0x2a, 0xc8, 0xc2, 0x96, 0x6a, 0xcb, 0x92, 0x62, 0x73, 0x2a, 0x15,
	0x9e, 0xe7, 0x61, 0xbd, 0xe7, 0xc4, 0x75, 0xb4, 0x09, 0x0d, 0x2e, 0x78, 0x46, 0x43, 0xb7, 0xe7,
	0xc4, 0x2e, 0x42, 0x00, 0x84, 0x8d, 0xc7, 0x2c, 0x2b, 0x67, 0x6a, 0x11, 0x36, 0x7a, 0x4e, 0x5c,
	0x8b, 0x0e, 0xa0, 0x6d, 0x76, 0xf7, 0x05, 0x59, 0xa0, 0x3d, 0x40, 0xaa, 0xc0, 0x5c, 0xe2, 0x4c,
	0x31, 0xc1, 0xcd, 0x70, 0x2a, 0x43, 0xa7, 0x57, 0x8f, 0xfd, 0xe8, 0x97, 0x03, 0xde, 0x68, 0xf5,
	0x88, 0xde, 0xaf, 0x63, 0x19, 0xcf, 0x4b, 0x65, 0xb1, 0xde, 0xd1, 0x8b, 0xc3, 0x15, 0xbb, 0xc3,
	0x4a, 0xd3, 0xb9, 0x06, 0xa1, 0x0f, 0xb0, 0x5d, 0xed, 0x14, 0xa5, 0x32, 0xad, 0x35, 0xd3, 0xfa,
	0xf2, 0x99, 0xd6, 0xa1, 0x41, 0xa1, 0x00, 0x5a, 0xa9, 0xc0, 0x05, 0x49, 0x18, 0x31, 0x1c, 0xdb,
	0x9a, 0x76, 0x2e, 0xa4, 0xb2, 0x4a, 0xb8, 0x46, 0x89, 0x7d, 0xd8, 0x55, 0xd3, 0x82, 0x62, 0x92,
	0xfc, 0xcd, 0xc6, 0x90, 0x36, 0x52, 0x11, 0x71, 0xc7, 0x6f, 0x85, 0xa2, 0x24, 0xdc, 0xe8, 0x39,
	0x71, 0x2b, 0xba, 0x87, 0xe0, 0xc9, 0xa1, 0xaf, 0xe0, 0xff, 0x47, 0x23, 0x9e, 0x4c, 0xb2, 0x3e,
	0x54, 0xbd, 0x5a, 0x97, 0x83, 0xd0, 0x1f, 0xc6, 0x98, 0x86, 0xde, 0x26, 0xd9, 0x84, 0x63, 0x55,
	0x16, 0xd4, 0x1c, 0x6d, 0x0e, 0xc8, 0xcb, 0x94, 0x65, 0xc9, 0x0d, 0x5d, 0xd8, 0xa3, 0xa3, 0x8f,
	0xb0, 0xf5, 0x94, 0x6e, 0x07, 0x36, 0xf0, 0x5c, 0x94, 0x5c, 0x99, 0x75, 0x75, 0xb4, 0x0b, 0xdd,
	0xbc, 0x4c, 0x67, 0xb6, 0xb1, 0x62, 0x7e, 0x74, 0x0f, 0xee, 0xa5, 0x90, 0x0a, 0xf9, 0xe0, 0x72,
	0x3c, 0xa7, 0x06, 0xde, 0xd6, 0x5f, 0x73, 0xcc, 0x66, 0x06, 0xd3, 0x46, 0x5d, 0x68, 0x66, 0x82,
	0x2b, 0xca, 0xd5, 0x4a, 0xba, 0x55, 0x62, 0x5c, 0xb3, 0xe0, 0x9f, 0x94, 0xad, 0x78, 0x3b, 0xe0,
	0x2f, 0xd5, 0x65, 0x6a, 0x46, 0x8d, 0x7e, 0xed, 0xe8, 0xb7, 0x03, 0x9d, 0xbe, 0x75, 0xee, 0x2b,
	0x95, 0x12, 0x4f, 0x28, 0xda, 0x06, 0x4f, 0xdc, 0xf1, 0x84, 0x0b, 0x42, 0xb5, 0x5d, 0x56, 0xb0,
	0x00, 0x5a, 0xba, 0x98, 0x8b, 0x42, 0xad, 0xe4, 0x51, 0xb8, 0x98, 0x50, 0xb5, 0xf4, 0xd4, 0xd7,
	0x41, 0x1d, 0x33, 0x4e, 0x92, 0x5b, 0x3c, 0x2b, 0x6d, 0x78, 0x5b, 0x68, 0x5f, 0x67, 0x99, 0x50,
	0x19, 0x36, 0x4c, 0x4e, 0x82, 0x6a, 0x4e, 0x2e, 0x04, 0xa1, 0xe8, 0x0d, 0xf8, 0x52, 0x89, 0x82,
	0x2e, 0xdb, 0xf4, 0x5d, 0xde, 0xd1, 0x6e, 0x15, 0x77, 0x6d, 0xde, 0xbf, 0xe9, 0x67, 0xf4, 0x16,
	0xba, 0x5c, 0x28, 0x36, 0x66, 0x94, 0x2c, 0x83, 0xde, 0x34, 0x93, 0xc3, 0xf5, 0xc9, 0x16, 0x72,
	0x86, 0xe5, 0x54, 0x0b, 0xc8, 0x64, 0x92, 0x33, 0x3e, 0x09, 0x5b, 0x26, 0x34, 0xa7, 0xe0, 0xaf,
	0x01, 0x0e, 0xa0, 0x4d, 0xb0, 0xc2, 0x89, 0x5a, 0xe4, 0xd6, 0x82, 0xce, 0xd1, 0x4e, 0x75, 0xda,
	0x00, 0x2b, 0x3c, 0x5a, 0xe4, 0x54, 0x1b, 0x53, 0x31, 0x2f, 0x06, 0xd7, 0x30, 0xe8, 0x42, 0x13,
	0x13, 0x52, 0x50, 0x29, 0x1f, 0xfc, 0xeb, 0x42, 0x73, 0xa9, 0x9e, 0x45, 0x0e, 0xc0, 0xab, 0x72,
	0xf0, 0xc1, 0xd5, 0xfb, 0x1e, 0xa4, 0x5d, 0xdb, 0x5e, 0x7b, 0x7e, 0xfb, 0xeb, 0x4f, 0xd0, 0x7a,
	0xbc, 0x24, 0x00, 0xbf, 0xff, 0x65, 0x78, 0xf2, 0x39, 0x39, 0x3b, 0x3d, 0x1e, 0x9c, 0x5e, 0x05,
	0xff, 0xa1, 0x0e, 0x80, 0xad, 0xf4, 0x87, 0x83, 0xef, 0x81, 0x83, 0xba, 0xe0, 0x8d, 0xae, 0x8e,
	0x2f, 0xae, 0x8f, 0x4f, 0x46, 0xe7, 0xc3, 0x8b, 0xa0, 0x86, 0x5a, 0xe0, 0x5e, 0x0e, 0xaf, 0x47,
	0x41, 0x3d, 0xdd, 0x30, 0xff, 0xb2, 0x77, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xd1, 0x9c,
	0xa9, 0xdc, 0x04, 0x00, 0x00,
}
