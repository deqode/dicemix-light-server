// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages/messages.proto

package messages

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

// for obtaining Status Code from response messages from server
// to parse response into suitable object
type GenericResponse struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericResponse) Reset()         { *m = GenericResponse{} }
func (m *GenericResponse) String() string { return proto.CompactTextString(m) }
func (*GenericResponse) ProtoMessage()    {}
func (*GenericResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_1a2b715719688b9a, []int{0}
}
func (m *GenericResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericResponse.Unmarshal(m, b)
}
func (m *GenericResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericResponse.Marshal(b, m, deterministic)
}
func (dst *GenericResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericResponse.Merge(dst, src)
}
func (m *GenericResponse) XXX_Size() int {
	return xxx_messageInfo_GenericResponse.Size(m)
}
func (m *GenericResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenericResponse proto.InternalMessageInfo

func (m *GenericResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

// used by server for obtaining Status Code
// from request messages sent from client
// to parse response into suitable object
type GenericRequest struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenericRequest) Reset()         { *m = GenericRequest{} }
func (m *GenericRequest) String() string { return proto.CompactTextString(m) }
func (*GenericRequest) ProtoMessage()    {}
func (*GenericRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_1a2b715719688b9a, []int{1}
}
func (m *GenericRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericRequest.Unmarshal(m, b)
}
func (m *GenericRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericRequest.Marshal(b, m, deterministic)
}
func (dst *GenericRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericRequest.Merge(dst, src)
}
func (m *GenericRequest) XXX_Size() int {
	return xxx_messageInfo_GenericRequest.Size(m)
}
func (m *GenericRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenericRequest proto.InternalMessageInfo

func (m *GenericRequest) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

// Response returned by server when attempt to join dicemix
// S_JOIN_RESPONSE
type RegisterResponse struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Id                   int32    `protobuf:"zigzag32,2,opt,name=Id,proto3" json:"Id,omitempty"`
	Timestamp            string   `protobuf:"bytes,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=Message,proto3" json:"Message,omitempty"`
	Err                  string   `protobuf:"bytes,5,opt,name=Err,proto3" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_1a2b715719688b9a, []int{2}
}
func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(dst, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RegisterResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RegisterResponse) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *RegisterResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RegisterResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

// Response returned by server for -
// StartDiceMix - Code S_START_DICEMIX
// KeyExchangeResponse - Code S_KEY_EXCHANGE
// DCSimpleResponse - Code S_SIMPLE_DC_VECTOR
// ConfirmationRequest - Code S_TX_CONFIRMATION
type DiceMixResponse struct {
	Code                 uint32       `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Peers                []*PeersInfo `protobuf:"bytes,2,rep,name=Peers,proto3" json:"Peers,omitempty"`
	Timestamp            string       `protobuf:"bytes,4,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Message              string       `protobuf:"bytes,5,opt,name=Message,proto3" json:"Message,omitempty"`
	Err                  string       `protobuf:"bytes,6,opt,name=Err,proto3" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DiceMixResponse) Reset()         { *m = DiceMixResponse{} }
func (m *DiceMixResponse) String() string { return proto.CompactTextString(m) }
func (*DiceMixResponse) ProtoMessage()    {}
func (*DiceMixResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_1a2b715719688b9a, []int{3}
}
func (m *DiceMixResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiceMixResponse.Unmarshal(m, b)
}
func (m *DiceMixResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiceMixResponse.Marshal(b, m, deterministic)
}
func (dst *DiceMixResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiceMixResponse.Merge(dst, src)
}
func (m *DiceMixResponse) XXX_Size() int {
	return xxx_messageInfo_DiceMixResponse.Size(m)
}
func (m *DiceMixResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiceMixResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiceMixResponse proto.InternalMessageInfo

func (m *DiceMixResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DiceMixResponse) GetPeers() []*PeersInfo {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *DiceMixResponse) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *DiceMixResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DiceMixResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

// Sub-message for DiceMixResponse
type PeersInfo struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	PublicKey            []byte   `protobuf:"bytes,2,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	PrivateKey           []byte   `protobuf:"bytes,3,opt,name=PrivateKey,proto3" json:"PrivateKey,omitempty"`
	NextPublicKey        []byte   `protobuf:"bytes,4,opt,name=NextPublicKey,proto3" json:"NextPublicKey,omitempty"`
	NumMsgs              uint32   `protobuf:"varint,5,opt,name=NumMsgs,proto3" json:"NumMsgs,omitempty"`
	DCVector             []uint64 `protobuf:"varint,6,rep,packed,name=DCVector,proto3" json:"DCVector,omitempty"`
	DCSimpleVector       [][]byte `protobuf:"bytes,7,rep,name=DCSimpleVector,proto3" json:"DCSimpleVector,omitempty"`
	OK                   bool     `protobuf:"varint,8,opt,name=OK,proto3" json:"OK,omitempty"`
	Messages             [][]byte `protobuf:"bytes,9,rep,name=Messages,proto3" json:"Messages,omitempty"`
	Confirmation         []byte   `protobuf:"bytes,10,opt,name=Confirmation,proto3" json:"Confirmation,omitempty"`
	MessageReceived      bool     `protobuf:"varint,11,opt,name=MessageReceived,proto3" json:"MessageReceived,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeersInfo) Reset()         { *m = PeersInfo{} }
func (m *PeersInfo) String() string { return proto.CompactTextString(m) }
func (*PeersInfo) ProtoMessage()    {}
func (*PeersInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_1a2b715719688b9a, []int{4}
}
func (m *PeersInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeersInfo.Unmarshal(m, b)
}
func (m *PeersInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeersInfo.Marshal(b, m, deterministic)
}
func (dst *PeersInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeersInfo.Merge(dst, src)
}
func (m *PeersInfo) XXX_Size() int {
	return xxx_messageInfo_PeersInfo.Size(m)
}
func (m *PeersInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PeersInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PeersInfo proto.InternalMessageInfo

func (m *PeersInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PeersInfo) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *PeersInfo) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *PeersInfo) GetNextPublicKey() []byte {
	if m != nil {
		return m.NextPublicKey
	}
	return nil
}

func (m *PeersInfo) GetNumMsgs() uint32 {
	if m != nil {
		return m.NumMsgs
	}
	return 0
}

func (m *PeersInfo) GetDCVector() []uint64 {
	if m != nil {
		return m.DCVector
	}
	return nil
}

func (m *PeersInfo) GetDCSimpleVector() [][]byte {
	if m != nil {
		return m.DCSimpleVector
	}
	return nil
}

func (m *PeersInfo) GetOK() bool {
	if m != nil {
		return m.OK
	}
	return false
}

func (m *PeersInfo) GetMessages() [][]byte {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *PeersInfo) GetConfirmation() []byte {
	if m != nil {
		return m.Confirmation
	}
	return nil
}

func (m *PeersInfo) GetMessageReceived() bool {
	if m != nil {
		return m.MessageReceived
	}
	return false
}

// For broadcasting our public key
// to initiate KeyExchange
// Code - C_KEY_EXCHANGE
type KeyExchangeRequest struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Id                   int32    `protobuf:"zigzag32,2,opt,name=Id,proto3" json:"Id,omitempty"`
	PublicKey            []byte   `protobuf:"bytes,3,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	NumMsgs              uint32   `protobuf:"varint,4,opt,name=NumMsgs,proto3" json:"NumMsgs,omitempty"`
	Timestamp            string   `protobuf:"bytes,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyExchangeRequest) Reset()         { *m = KeyExchangeRequest{} }
func (m *KeyExchangeRequest) String() string { return proto.CompactTextString(m) }
func (*KeyExchangeRequest) ProtoMessage()    {}
func (*KeyExchangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_1a2b715719688b9a, []int{5}
}
func (m *KeyExchangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyExchangeRequest.Unmarshal(m, b)
}
func (m *KeyExchangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyExchangeRequest.Marshal(b, m, deterministic)
}
func (dst *KeyExchangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyExchangeRequest.Merge(dst, src)
}
func (m *KeyExchangeRequest) XXX_Size() int {
	return xxx_messageInfo_KeyExchangeRequest.Size(m)
}
func (m *KeyExchangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyExchangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyExchangeRequest proto.InternalMessageInfo

func (m *KeyExchangeRequest) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *KeyExchangeRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *KeyExchangeRequest) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *KeyExchangeRequest) GetNumMsgs() uint32 {
	if m != nil {
		return m.NumMsgs
	}
	return 0
}

func (m *KeyExchangeRequest) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

// For broadcasting our DC Exponential Vector
// to initiate DC-EXP
// Code - C_EXP_DC_VECTOR
type DCExpRequest struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Id                   int32    `protobuf:"zigzag32,2,opt,name=Id,proto3" json:"Id,omitempty"`
	DCExpVector          []uint64 `protobuf:"varint,3,rep,packed,name=DCExpVector,proto3" json:"DCExpVector,omitempty"`
	Timestamp            string   `protobuf:"bytes,4,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DCExpRequest) Reset()         { *m = DCExpRequest{} }
func (m *DCExpRequest) String() string { return proto.CompactTextString(m) }
func (*DCExpRequest) ProtoMessage()    {}
func (*DCExpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_1a2b715719688b9a, []int{6}
}
func (m *DCExpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DCExpRequest.Unmarshal(m, b)
}
func (m *DCExpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DCExpRequest.Marshal(b, m, deterministic)
}
func (dst *DCExpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DCExpRequest.Merge(dst, src)
}
func (m *DCExpRequest) XXX_Size() int {
	return xxx_messageInfo_DCExpRequest.Size(m)
}
func (m *DCExpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DCExpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DCExpRequest proto.InternalMessageInfo

func (m *DCExpRequest) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DCExpRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DCExpRequest) GetDCExpVector() []uint64 {
	if m != nil {
		return m.DCExpVector
	}
	return nil
}

func (m *DCExpRequest) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

// Response against DCExpRequest
// conatins ROOTS calculated by server using FLINT
// Code - S_EXP_DC_VECTOR
type DCExpResponse struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Roots                []uint64 `protobuf:"varint,2,rep,packed,name=Roots,proto3" json:"Roots,omitempty"`
	Timestamp            string   `protobuf:"bytes,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=Message,proto3" json:"Message,omitempty"`
	Err                  string   `protobuf:"bytes,5,opt,name=Err,proto3" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DCExpResponse) Reset()         { *m = DCExpResponse{} }
func (m *DCExpResponse) String() string { return proto.CompactTextString(m) }
func (*DCExpResponse) ProtoMessage()    {}
func (*DCExpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_1a2b715719688b9a, []int{7}
}
func (m *DCExpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DCExpResponse.Unmarshal(m, b)
}
func (m *DCExpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DCExpResponse.Marshal(b, m, deterministic)
}
func (dst *DCExpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DCExpResponse.Merge(dst, src)
}
func (m *DCExpResponse) XXX_Size() int {
	return xxx_messageInfo_DCExpResponse.Size(m)
}
func (m *DCExpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DCExpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DCExpResponse proto.InternalMessageInfo

func (m *DCExpResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DCExpResponse) GetRoots() []uint64 {
	if m != nil {
		return m.Roots
	}
	return nil
}

func (m *DCExpResponse) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *DCExpResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *DCExpResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

// For broadcasting our DC Simple Vector
// to initiate DC-SIMPLE
// C_SIMPLE_DC_VECTOR
type DCSimpleRequest struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Id                   int32    `protobuf:"zigzag32,2,opt,name=Id,proto3" json:"Id,omitempty"`
	DCSimpleVector       [][]byte `protobuf:"bytes,3,rep,name=DCSimpleVector,proto3" json:"DCSimpleVector,omitempty"`
	MyOk                 bool     `protobuf:"varint,4,opt,name=MyOk,proto3" json:"MyOk,omitempty"`
	NextPublicKey        []byte   `protobuf:"bytes,5,opt,name=NextPublicKey,proto3" json:"NextPublicKey,omitempty"`
	Timestamp            string   `protobuf:"bytes,6,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DCSimpleRequest) Reset()         { *m = DCSimpleRequest{} }
func (m *DCSimpleRequest) String() string { return proto.CompactTextString(m) }
func (*DCSimpleRequest) ProtoMessage()    {}
func (*DCSimpleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_1a2b715719688b9a, []int{8}
}
func (m *DCSimpleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DCSimpleRequest.Unmarshal(m, b)
}
func (m *DCSimpleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DCSimpleRequest.Marshal(b, m, deterministic)
}
func (dst *DCSimpleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DCSimpleRequest.Merge(dst, src)
}
func (m *DCSimpleRequest) XXX_Size() int {
	return xxx_messageInfo_DCSimpleRequest.Size(m)
}
func (m *DCSimpleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DCSimpleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DCSimpleRequest proto.InternalMessageInfo

func (m *DCSimpleRequest) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *DCSimpleRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DCSimpleRequest) GetDCSimpleVector() [][]byte {
	if m != nil {
		return m.DCSimpleVector
	}
	return nil
}

func (m *DCSimpleRequest) GetMyOk() bool {
	if m != nil {
		return m.MyOk
	}
	return false
}

func (m *DCSimpleRequest) GetNextPublicKey() []byte {
	if m != nil {
		return m.NextPublicKey
	}
	return nil
}

func (m *DCSimpleRequest) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

// For broadcasting our confirmation for messages
// C_TX_CONFIRMATION
type ConfirmationRequest struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Id                   int32    `protobuf:"zigzag32,2,opt,name=Id,proto3" json:"Id,omitempty"`
	Confirmation         []byte   `protobuf:"bytes,3,opt,name=Confirmation,proto3" json:"Confirmation,omitempty"`
	Messages             [][]byte `protobuf:"bytes,4,rep,name=Messages,proto3" json:"Messages,omitempty"`
	Timestamp            string   `protobuf:"bytes,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmationRequest) Reset()         { *m = ConfirmationRequest{} }
func (m *ConfirmationRequest) String() string { return proto.CompactTextString(m) }
func (*ConfirmationRequest) ProtoMessage()    {}
func (*ConfirmationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_1a2b715719688b9a, []int{9}
}
func (m *ConfirmationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmationRequest.Unmarshal(m, b)
}
func (m *ConfirmationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmationRequest.Marshal(b, m, deterministic)
}
func (dst *ConfirmationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmationRequest.Merge(dst, src)
}
func (m *ConfirmationRequest) XXX_Size() int {
	return xxx_messageInfo_ConfirmationRequest.Size(m)
}
func (m *ConfirmationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmationRequest proto.InternalMessageInfo

func (m *ConfirmationRequest) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ConfirmationRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ConfirmationRequest) GetConfirmation() []byte {
	if m != nil {
		return m.Confirmation
	}
	return nil
}

func (m *ConfirmationRequest) GetMessages() [][]byte {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *ConfirmationRequest) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

// Possible response against ConfirmationRequest
// only when all peers send valid confirmations to server
// Code - S_TX_SUCCESSFUL
type TXDoneResponse struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Messages             [][]byte `protobuf:"bytes,2,rep,name=Messages,proto3" json:"Messages,omitempty"`
	Timestamp            string   `protobuf:"bytes,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=Message,proto3" json:"Message,omitempty"`
	Err                  string   `protobuf:"bytes,5,opt,name=Err,proto3" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TXDoneResponse) Reset()         { *m = TXDoneResponse{} }
func (m *TXDoneResponse) String() string { return proto.CompactTextString(m) }
func (*TXDoneResponse) ProtoMessage()    {}
func (*TXDoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_1a2b715719688b9a, []int{10}
}
func (m *TXDoneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TXDoneResponse.Unmarshal(m, b)
}
func (m *TXDoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TXDoneResponse.Marshal(b, m, deterministic)
}
func (dst *TXDoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TXDoneResponse.Merge(dst, src)
}
func (m *TXDoneResponse) XXX_Size() int {
	return xxx_messageInfo_TXDoneResponse.Size(m)
}
func (m *TXDoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TXDoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TXDoneResponse proto.InternalMessageInfo

func (m *TXDoneResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TXDoneResponse) GetMessages() [][]byte {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *TXDoneResponse) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *TXDoneResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *TXDoneResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

// message sent by server
// to initiate KESK
type InitiaiteKESK struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Timestamp            string   `protobuf:"bytes,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=Message,proto3" json:"Message,omitempty"`
	Err                  string   `protobuf:"bytes,4,opt,name=Err,proto3" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitiaiteKESK) Reset()         { *m = InitiaiteKESK{} }
func (m *InitiaiteKESK) String() string { return proto.CompactTextString(m) }
func (*InitiaiteKESK) ProtoMessage()    {}
func (*InitiaiteKESK) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_1a2b715719688b9a, []int{11}
}
func (m *InitiaiteKESK) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitiaiteKESK.Unmarshal(m, b)
}
func (m *InitiaiteKESK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitiaiteKESK.Marshal(b, m, deterministic)
}
func (dst *InitiaiteKESK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitiaiteKESK.Merge(dst, src)
}
func (m *InitiaiteKESK) XXX_Size() int {
	return xxx_messageInfo_InitiaiteKESK.Size(m)
}
func (m *InitiaiteKESK) XXX_DiscardUnknown() {
	xxx_messageInfo_InitiaiteKESK.DiscardUnknown(m)
}

var xxx_messageInfo_InitiaiteKESK proto.InternalMessageInfo

func (m *InitiaiteKESK) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *InitiaiteKESK) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *InitiaiteKESK) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *InitiaiteKESK) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

// For broadcasting our KESK
// to initiate BLAME
type InitiaiteKESKResponse struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Id                   int32    `protobuf:"zigzag32,2,opt,name=Id,proto3" json:"Id,omitempty"`
	PrivateKey           []byte   `protobuf:"bytes,3,opt,name=PrivateKey,proto3" json:"PrivateKey,omitempty"`
	Timestamp            string   `protobuf:"bytes,4,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InitiaiteKESKResponse) Reset()         { *m = InitiaiteKESKResponse{} }
func (m *InitiaiteKESKResponse) String() string { return proto.CompactTextString(m) }
func (*InitiaiteKESKResponse) ProtoMessage()    {}
func (*InitiaiteKESKResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_messages_1a2b715719688b9a, []int{12}
}
func (m *InitiaiteKESKResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InitiaiteKESKResponse.Unmarshal(m, b)
}
func (m *InitiaiteKESKResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InitiaiteKESKResponse.Marshal(b, m, deterministic)
}
func (dst *InitiaiteKESKResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitiaiteKESKResponse.Merge(dst, src)
}
func (m *InitiaiteKESKResponse) XXX_Size() int {
	return xxx_messageInfo_InitiaiteKESKResponse.Size(m)
}
func (m *InitiaiteKESKResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InitiaiteKESKResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InitiaiteKESKResponse proto.InternalMessageInfo

func (m *InitiaiteKESKResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *InitiaiteKESKResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InitiaiteKESKResponse) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *InitiaiteKESKResponse) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func init() {
	proto.RegisterType((*GenericResponse)(nil), "messages.GenericResponse")
	proto.RegisterType((*GenericRequest)(nil), "messages.GenericRequest")
	proto.RegisterType((*RegisterResponse)(nil), "messages.RegisterResponse")
	proto.RegisterType((*DiceMixResponse)(nil), "messages.DiceMixResponse")
	proto.RegisterType((*PeersInfo)(nil), "messages.PeersInfo")
	proto.RegisterType((*KeyExchangeRequest)(nil), "messages.KeyExchangeRequest")
	proto.RegisterType((*DCExpRequest)(nil), "messages.DCExpRequest")
	proto.RegisterType((*DCExpResponse)(nil), "messages.DCExpResponse")
	proto.RegisterType((*DCSimpleRequest)(nil), "messages.DCSimpleRequest")
	proto.RegisterType((*ConfirmationRequest)(nil), "messages.ConfirmationRequest")
	proto.RegisterType((*TXDoneResponse)(nil), "messages.TXDoneResponse")
	proto.RegisterType((*InitiaiteKESK)(nil), "messages.InitiaiteKESK")
	proto.RegisterType((*InitiaiteKESKResponse)(nil), "messages.InitiaiteKESKResponse")
}

func init() { proto.RegisterFile("messages/messages.proto", fileDescriptor_messages_1a2b715719688b9a) }

var fileDescriptor_messages_1a2b715719688b9a = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xff, 0xda, 0x64, 0x9a, 0xa4, 0x65, 0x0b, 0xc2, 0x42, 0x08, 0x59, 0x56, 0x41, 0xe6,
	0x52, 0x24, 0x78, 0x84, 0x24, 0x42, 0x91, 0x95, 0x26, 0xda, 0x56, 0x88, 0xab, 0xeb, 0x4c, 0xc3,
	0x8a, 0xda, 0x1b, 0xbc, 0x9b, 0x2a, 0xb9, 0x21, 0x71, 0xe3, 0xc0, 0x15, 0x89, 0xf7, 0xe0, 0x59,
	0x78, 0x1d, 0xe4, 0x8d, 0x13, 0xff, 0xd4, 0x18, 0x45, 0x70, 0xdb, 0xf9, 0x76, 0x3c, 0xf3, 0xed,
	0xb7, 0xdf, 0x78, 0xe1, 0x71, 0x84, 0x42, 0x04, 0x73, 0x14, 0xaf, 0xb6, 0x8b, 0xf3, 0x45, 0xc2,
	0x25, 0x27, 0xad, 0x6d, 0xec, 0x3e, 0x87, 0xe3, 0xb7, 0x18, 0x63, 0xc2, 0x42, 0x8a, 0x62, 0xc1,
	0x63, 0x81, 0x84, 0x80, 0xd9, 0xe7, 0x33, 0xb4, 0x35, 0x47, 0xf3, 0xba, 0x54, 0xad, 0xdd, 0x33,
	0xe8, 0xed, 0xd2, 0x3e, 0x2d, 0x51, 0xc8, 0xda, 0xac, 0xcf, 0x1a, 0x9c, 0x50, 0x9c, 0x33, 0x21,
	0x31, 0x69, 0x2a, 0x47, 0x7a, 0xa0, 0x8f, 0x66, 0xb6, 0xee, 0x68, 0xde, 0x03, 0xaa, 0x8f, 0x66,
	0xe4, 0x29, 0xb4, 0xaf, 0x58, 0x84, 0x42, 0x06, 0xd1, 0xc2, 0x36, 0x1c, 0xcd, 0x6b, 0xd3, 0x1c,
	0x20, 0x36, 0x1c, 0x8e, 0x37, 0x7c, 0x6d, 0x53, 0xed, 0x6d, 0x43, 0x72, 0x02, 0xc6, 0x30, 0x49,
	0x6c, 0x4b, 0xa1, 0xe9, 0xd2, 0xfd, 0xa1, 0xc1, 0xf1, 0x80, 0x85, 0x38, 0x66, 0xab, 0x46, 0x06,
	0x2f, 0xc1, 0x9a, 0x22, 0x26, 0xc2, 0xd6, 0x1d, 0xc3, 0x3b, 0x7a, 0x7d, 0x7a, 0xbe, 0x53, 0x48,
	0xc1, 0xa3, 0xf8, 0x86, 0xd3, 0x4d, 0x46, 0x99, 0x9c, 0xd9, 0x40, 0xce, 0xaa, 0x25, 0x77, 0x90,
	0x93, 0xfb, 0xa5, 0x43, 0x7b, 0x57, 0x3e, 0x13, 0x21, 0x25, 0x65, 0x6d, 0x45, 0x98, 0x2e, 0xaf,
	0x6f, 0x59, 0xe8, 0xe3, 0x5a, 0x69, 0xd3, 0xa1, 0x39, 0x40, 0x9e, 0x01, 0x4c, 0x13, 0x76, 0x17,
	0x48, 0x4c, 0xb7, 0x0d, 0xb5, 0x5d, 0x40, 0xc8, 0x19, 0x74, 0x2f, 0x70, 0x25, 0xf3, 0x0a, 0xa6,
	0x4a, 0x29, 0x83, 0x29, 0xdb, 0x8b, 0x65, 0x34, 0x16, 0x73, 0xa1, 0xd8, 0x76, 0xe9, 0x36, 0x24,
	0x4f, 0xa0, 0x35, 0xe8, 0xbf, 0xc3, 0x50, 0xf2, 0x94, 0xb2, 0xe1, 0x99, 0x74, 0x17, 0x93, 0x17,
	0xd0, 0x1b, 0xf4, 0x2f, 0x59, 0xb4, 0xb8, 0xc5, 0x2c, 0xe3, 0xd0, 0x31, 0xbc, 0x0e, 0xad, 0xa0,
	0xe9, 0x89, 0x26, 0xbe, 0xdd, 0x72, 0x34, 0xaf, 0x45, 0xf5, 0x89, 0x9f, 0xd6, 0xcc, 0xc4, 0x10,
	0x76, 0x5b, 0x7d, 0xb1, 0x8b, 0x89, 0x0b, 0x9d, 0x3e, 0x8f, 0x6f, 0x58, 0x12, 0x05, 0x92, 0xf1,
	0xd8, 0x06, 0x45, 0xb7, 0x84, 0x11, 0x0f, 0x8e, 0xb3, 0x7c, 0x8a, 0x21, 0xb2, 0x3b, 0x9c, 0xd9,
	0x47, 0xaa, 0x78, 0x15, 0x76, 0xbf, 0x69, 0x40, 0x7c, 0x5c, 0x0f, 0x57, 0xe1, 0x87, 0x20, 0x4e,
	0xf1, 0x3f, 0x9a, 0xb4, 0xce, 0x7b, 0xb9, 0x68, 0x46, 0x55, 0xf6, 0x82, 0x60, 0x66, 0x59, 0xb0,
	0x92, 0x2d, 0xac, 0x8a, 0x2d, 0xdc, 0x04, 0x3a, 0x83, 0xfe, 0x70, 0xb5, 0xd8, 0x87, 0x89, 0x03,
	0x47, 0xea, 0x9b, 0x4c, 0x63, 0x43, 0xdd, 0x42, 0x11, 0x6a, 0xb6, 0xa2, 0xfb, 0x45, 0x83, 0x6e,
	0xd6, 0xb4, 0xc1, 0xf9, 0x0f, 0xc1, 0xa2, 0x9c, 0xcb, 0x8d, 0xf3, 0x4d, 0xba, 0x09, 0xfe, 0xe3,
	0x04, 0xfe, 0x4c, 0x27, 0x30, 0xf3, 0xc5, 0x3e, 0xa7, 0xbf, 0x6f, 0x32, 0xa3, 0xd6, 0x64, 0x04,
	0xcc, 0xf1, 0x7a, 0xf2, 0x51, 0x11, 0x69, 0x51, 0xb5, 0xbe, 0x6f, 0x7e, 0xab, 0xce, 0xfc, 0xa5,
	0x33, 0x1e, 0x54, 0xd5, 0xfb, 0xae, 0xc1, 0x69, 0xd1, 0x7d, 0xfb, 0x70, 0xaf, 0x9a, 0xd9, 0xa8,
	0x31, 0x73, 0x71, 0x18, 0xcc, 0xca, 0x30, 0x34, 0x7b, 0xe9, 0xab, 0x06, 0xbd, 0xab, 0xf7, 0x03,
	0x1e, 0x63, 0xe3, 0xc5, 0x16, 0x1b, 0xe8, 0x4d, 0x0d, 0xfe, 0xe9, 0x7a, 0x23, 0xe8, 0x8e, 0x62,
	0x26, 0x59, 0xc0, 0x24, 0xfa, 0xc3, 0x4b, 0xbf, 0x96, 0x4a, 0xa9, 0x9d, 0xde, 0xd0, 0xce, 0xa8,
	0x6d, 0x67, 0xe6, 0xed, 0xd6, 0xf0, 0xa8, 0xd4, 0x6e, 0xaf, 0x67, 0xe5, 0x6f, 0xff, 0xcc, 0xc6,
	0x71, 0xba, 0x3e, 0x50, 0x6f, 0xe5, 0x9b, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xf0, 0x90,
	0xc1, 0x46, 0x07, 0x00, 0x00,
}