// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: payloads/payloads.proto

package payloads

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strconv "strconv"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PayloadType int32

const (
	BINARY  PayloadType = 0
	TEXT    PayloadType = 1
	ACK     PayloadType = 2
	SESSION PayloadType = 3
)

var PayloadType_name = map[int32]string{
	0: "BINARY",
	1: "TEXT",
	2: "ACK",
	3: "SESSION",
}
var PayloadType_value = map[string]int32{
	"BINARY":  0,
	"TEXT":    1,
	"ACK":     2,
	"SESSION": 3,
}

func (PayloadType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_payloads_834d6bc1cbaabddf, []int{0}
}

type Message struct {
	Payload      []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Encrypted    bool   `protobuf:"varint,2,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
	Nonce        []byte `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	EncryptedKey []byte `protobuf:"bytes,4,opt,name=encrypted_key,json=encryptedKey,proto3" json:"encrypted_key,omitempty"`
}

func (m *Message) Reset()      { *m = Message{} }
func (*Message) ProtoMessage() {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_payloads_834d6bc1cbaabddf, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetEncrypted() bool {
	if m != nil {
		return m.Encrypted
	}
	return false
}

func (m *Message) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *Message) GetEncryptedKey() []byte {
	if m != nil {
		return m.EncryptedKey
	}
	return nil
}

type Payload struct {
	Type      PayloadType `protobuf:"varint,1,opt,name=type,proto3,enum=payloads.PayloadType" json:"type,omitempty"`
	MessageId []byte      `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Data      []byte      `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	ReplyToId []byte      `protobuf:"bytes,4,opt,name=reply_to_id,json=replyToId,proto3" json:"reply_to_id,omitempty"`
	NoReply   bool        `protobuf:"varint,5,opt,name=no_reply,json=noReply,proto3" json:"no_reply,omitempty"`
}

func (m *Payload) Reset()      { *m = Payload{} }
func (*Payload) ProtoMessage() {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_payloads_834d6bc1cbaabddf, []int{1}
}
func (m *Payload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(dst, src)
}
func (m *Payload) XXX_Size() int {
	return m.Size()
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetType() PayloadType {
	if m != nil {
		return m.Type
	}
	return BINARY
}

func (m *Payload) GetMessageId() []byte {
	if m != nil {
		return m.MessageId
	}
	return nil
}

func (m *Payload) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Payload) GetReplyToId() []byte {
	if m != nil {
		return m.ReplyToId
	}
	return nil
}

func (m *Payload) GetNoReply() bool {
	if m != nil {
		return m.NoReply
	}
	return false
}

type TextData struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *TextData) Reset()      { *m = TextData{} }
func (*TextData) ProtoMessage() {}
func (*TextData) Descriptor() ([]byte, []int) {
	return fileDescriptor_payloads_834d6bc1cbaabddf, []int{2}
}
func (m *TextData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TextData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TextData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TextData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextData.Merge(dst, src)
}
func (m *TextData) XXX_Size() int {
	return m.Size()
}
func (m *TextData) XXX_DiscardUnknown() {
	xxx_messageInfo_TextData.DiscardUnknown(m)
}

var xxx_messageInfo_TextData proto.InternalMessageInfo

func (m *TextData) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "payloads.Message")
	proto.RegisterType((*Payload)(nil), "payloads.Payload")
	proto.RegisterType((*TextData)(nil), "payloads.TextData")
	proto.RegisterEnum("payloads.PayloadType", PayloadType_name, PayloadType_value)
}
func (x PayloadType) String() string {
	s, ok := PayloadType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *Message) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Message)
	if !ok {
		that2, ok := that.(Message)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Payload, that1.Payload) {
		return false
	}
	if this.Encrypted != that1.Encrypted {
		return false
	}
	if !bytes.Equal(this.Nonce, that1.Nonce) {
		return false
	}
	if !bytes.Equal(this.EncryptedKey, that1.EncryptedKey) {
		return false
	}
	return true
}
func (this *Payload) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Payload)
	if !ok {
		that2, ok := that.(Payload)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if !bytes.Equal(this.MessageId, that1.MessageId) {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	if !bytes.Equal(this.ReplyToId, that1.ReplyToId) {
		return false
	}
	if this.NoReply != that1.NoReply {
		return false
	}
	return true
}
func (this *TextData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TextData)
	if !ok {
		that2, ok := that.(TextData)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Text != that1.Text {
		return false
	}
	return true
}
func (this *Message) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&payloads.Message{")
	s = append(s, "Payload: "+fmt.Sprintf("%#v", this.Payload)+",\n")
	s = append(s, "Encrypted: "+fmt.Sprintf("%#v", this.Encrypted)+",\n")
	s = append(s, "Nonce: "+fmt.Sprintf("%#v", this.Nonce)+",\n")
	s = append(s, "EncryptedKey: "+fmt.Sprintf("%#v", this.EncryptedKey)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Payload) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&payloads.Payload{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "MessageId: "+fmt.Sprintf("%#v", this.MessageId)+",\n")
	s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	s = append(s, "ReplyToId: "+fmt.Sprintf("%#v", this.ReplyToId)+",\n")
	s = append(s, "NoReply: "+fmt.Sprintf("%#v", this.NoReply)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TextData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&payloads.TextData{")
	s = append(s, "Text: "+fmt.Sprintf("%#v", this.Text)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringPayloads(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPayloads(dAtA, i, uint64(len(m.Payload)))
		i += copy(dAtA[i:], m.Payload)
	}
	if m.Encrypted {
		dAtA[i] = 0x10
		i++
		if m.Encrypted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if len(m.Nonce) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPayloads(dAtA, i, uint64(len(m.Nonce)))
		i += copy(dAtA[i:], m.Nonce)
	}
	if len(m.EncryptedKey) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintPayloads(dAtA, i, uint64(len(m.EncryptedKey)))
		i += copy(dAtA[i:], m.EncryptedKey)
	}
	return i, nil
}

func (m *Payload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payload) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPayloads(dAtA, i, uint64(m.Type))
	}
	if len(m.MessageId) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPayloads(dAtA, i, uint64(len(m.MessageId)))
		i += copy(dAtA[i:], m.MessageId)
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPayloads(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if len(m.ReplyToId) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintPayloads(dAtA, i, uint64(len(m.ReplyToId)))
		i += copy(dAtA[i:], m.ReplyToId)
	}
	if m.NoReply {
		dAtA[i] = 0x28
		i++
		if m.NoReply {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *TextData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TextData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Text) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPayloads(dAtA, i, uint64(len(m.Text)))
		i += copy(dAtA[i:], m.Text)
	}
	return i, nil
}

func encodeVarintPayloads(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedMessage(r randyPayloads, easy bool) *Message {
	this := &Message{}
	v1 := r.Intn(100)
	this.Payload = make([]byte, v1)
	for i := 0; i < v1; i++ {
		this.Payload[i] = byte(r.Intn(256))
	}
	this.Encrypted = bool(bool(r.Intn(2) == 0))
	v2 := r.Intn(100)
	this.Nonce = make([]byte, v2)
	for i := 0; i < v2; i++ {
		this.Nonce[i] = byte(r.Intn(256))
	}
	v3 := r.Intn(100)
	this.EncryptedKey = make([]byte, v3)
	for i := 0; i < v3; i++ {
		this.EncryptedKey[i] = byte(r.Intn(256))
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedPayload(r randyPayloads, easy bool) *Payload {
	this := &Payload{}
	this.Type = PayloadType([]int32{0, 1, 2, 3}[r.Intn(4)])
	v4 := r.Intn(100)
	this.MessageId = make([]byte, v4)
	for i := 0; i < v4; i++ {
		this.MessageId[i] = byte(r.Intn(256))
	}
	v5 := r.Intn(100)
	this.Data = make([]byte, v5)
	for i := 0; i < v5; i++ {
		this.Data[i] = byte(r.Intn(256))
	}
	v6 := r.Intn(100)
	this.ReplyToId = make([]byte, v6)
	for i := 0; i < v6; i++ {
		this.ReplyToId[i] = byte(r.Intn(256))
	}
	this.NoReply = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedTextData(r randyPayloads, easy bool) *TextData {
	this := &TextData{}
	this.Text = string(randStringPayloads(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyPayloads interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RunePayloads(r randyPayloads) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringPayloads(r randyPayloads) string {
	v7 := r.Intn(100)
	tmps := make([]rune, v7)
	for i := 0; i < v7; i++ {
		tmps[i] = randUTF8RunePayloads(r)
	}
	return string(tmps)
}
func randUnrecognizedPayloads(r randyPayloads, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldPayloads(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldPayloads(dAtA []byte, r randyPayloads, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulatePayloads(dAtA, uint64(key))
		v8 := r.Int63()
		if r.Intn(2) == 0 {
			v8 *= -1
		}
		dAtA = encodeVarintPopulatePayloads(dAtA, uint64(v8))
	case 1:
		dAtA = encodeVarintPopulatePayloads(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulatePayloads(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulatePayloads(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulatePayloads(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulatePayloads(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovPayloads(uint64(l))
	}
	if m.Encrypted {
		n += 2
	}
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovPayloads(uint64(l))
	}
	l = len(m.EncryptedKey)
	if l > 0 {
		n += 1 + l + sovPayloads(uint64(l))
	}
	return n
}

func (m *Payload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovPayloads(uint64(m.Type))
	}
	l = len(m.MessageId)
	if l > 0 {
		n += 1 + l + sovPayloads(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovPayloads(uint64(l))
	}
	l = len(m.ReplyToId)
	if l > 0 {
		n += 1 + l + sovPayloads(uint64(l))
	}
	if m.NoReply {
		n += 2
	}
	return n
}

func (m *TextData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovPayloads(uint64(l))
	}
	return n
}

func sovPayloads(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPayloads(x uint64) (n int) {
	return sovPayloads(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Message) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Message{`,
		`Payload:` + fmt.Sprintf("%v", this.Payload) + `,`,
		`Encrypted:` + fmt.Sprintf("%v", this.Encrypted) + `,`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`EncryptedKey:` + fmt.Sprintf("%v", this.EncryptedKey) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Payload) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Payload{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`MessageId:` + fmt.Sprintf("%v", this.MessageId) + `,`,
		`Data:` + fmt.Sprintf("%v", this.Data) + `,`,
		`ReplyToId:` + fmt.Sprintf("%v", this.ReplyToId) + `,`,
		`NoReply:` + fmt.Sprintf("%v", this.NoReply) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TextData) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TextData{`,
		`Text:` + fmt.Sprintf("%v", this.Text) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringPayloads(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayloads
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloads
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPayloads
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encrypted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloads
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Encrypted = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloads
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPayloads
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = append(m.Nonce[:0], dAtA[iNdEx:postIndex]...)
			if m.Nonce == nil {
				m.Nonce = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloads
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPayloads
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptedKey = append(m.EncryptedKey[:0], dAtA[iNdEx:postIndex]...)
			if m.EncryptedKey == nil {
				m.EncryptedKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayloads(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPayloads
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
func (m *Payload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayloads
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Payload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Payload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloads
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (PayloadType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloads
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPayloads
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageId = append(m.MessageId[:0], dAtA[iNdEx:postIndex]...)
			if m.MessageId == nil {
				m.MessageId = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloads
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPayloads
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplyToId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloads
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPayloads
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReplyToId = append(m.ReplyToId[:0], dAtA[iNdEx:postIndex]...)
			if m.ReplyToId == nil {
				m.ReplyToId = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoReply", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloads
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NoReply = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPayloads(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPayloads
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
func (m *TextData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayloads
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TextData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TextData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayloads
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPayloads
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayloads(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPayloads
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
func skipPayloads(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPayloads
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
					return 0, ErrIntOverflowPayloads
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPayloads
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPayloads
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPayloads
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPayloads(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPayloads = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPayloads   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("payloads/payloads.proto", fileDescriptor_payloads_834d6bc1cbaabddf) }

var fileDescriptor_payloads_834d6bc1cbaabddf = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x31, 0xcf, 0xd2, 0x50,
	0x14, 0x86, 0x7b, 0xa1, 0xd0, 0xf6, 0x80, 0xa6, 0xb9, 0xd1, 0x58, 0x8d, 0x9e, 0x10, 0x5c, 0xd0,
	0x44, 0x48, 0x74, 0x32, 0x71, 0x01, 0x65, 0x68, 0x88, 0x68, 0x4a, 0x07, 0x9d, 0x9a, 0x42, 0xaf,
	0x95, 0x08, 0xbd, 0x0d, 0x5c, 0x12, 0x9a, 0x38, 0xf8, 0x13, 0xfc, 0x0d, 0x4e, 0xfe, 0x04, 0x7f,
	0x82, 0x23, 0x23, 0xa3, 0xbd, 0x2c, 0x8e, 0x8c, 0xdf, 0xf8, 0x85, 0xdb, 0xc2, 0xf7, 0x6d, 0xef,
	0xf3, 0xf6, 0x3d, 0xf7, 0xbc, 0x27, 0x85, 0x07, 0x69, 0x98, 0x2d, 0x78, 0x18, 0xad, 0x7b, 0x67,
	0xd1, 0x4d, 0x57, 0x5c, 0x70, 0x6a, 0x9e, 0xf9, 0xd1, 0x8b, 0x78, 0x2e, 0xbe, 0x6e, 0xa6, 0xdd,
	0x19, 0x5f, 0xf6, 0x62, 0x1e, 0xf3, 0x9e, 0x0a, 0x4c, 0x37, 0x5f, 0x14, 0x29, 0x50, 0xaa, 0x18,
	0x6c, 0x7f, 0x07, 0xe3, 0x3d, 0x5b, 0xaf, 0xc3, 0x98, 0x51, 0x07, 0x8c, 0xf2, 0x15, 0x87, 0xb4,
	0x48, 0xa7, 0xe9, 0x9d, 0x91, 0x3e, 0x06, 0x8b, 0x25, 0xb3, 0x55, 0x96, 0x0a, 0x16, 0x39, 0x95,
	0x16, 0xe9, 0x98, 0xde, 0x8d, 0x41, 0xef, 0x41, 0x2d, 0xe1, 0xc9, 0x8c, 0x39, 0x55, 0x35, 0x55,
	0x00, 0x7d, 0x0a, 0x77, 0x2e, 0x91, 0xe0, 0x1b, 0xcb, 0x1c, 0x5d, 0x7d, 0x6d, 0x5e, 0xcc, 0x11,
	0xcb, 0xda, 0xbf, 0x08, 0x18, 0x1f, 0xcb, 0x25, 0xcf, 0x40, 0x17, 0x59, 0xca, 0xd4, 0xee, 0xbb,
	0x2f, 0xef, 0x77, 0x2f, 0x17, 0x96, 0x01, 0x3f, 0x4b, 0x99, 0xa7, 0x22, 0xf4, 0x09, 0xc0, 0xb2,
	0x28, 0x1d, 0xcc, 0x8b, 0x42, 0x4d, 0xcf, 0x2a, 0x1d, 0x37, 0xa2, 0x14, 0xf4, 0x28, 0x14, 0x61,
	0xd9, 0x47, 0x69, 0x8a, 0xd0, 0x58, 0xb1, 0x74, 0x91, 0x05, 0x82, 0x9f, 0x66, 0x8a, 0x32, 0x96,
	0xb2, 0x7c, 0xee, 0x46, 0xf4, 0x21, 0x98, 0x09, 0x0f, 0x14, 0x3b, 0x35, 0x75, 0xa1, 0x91, 0x70,
	0xef, 0x84, 0x6d, 0x04, 0xd3, 0x67, 0x5b, 0xf1, 0xee, 0xf4, 0x0c, 0x05, 0x5d, 0xb0, 0xad, 0x50,
	0x25, 0x2d, 0x4f, 0xe9, 0xe7, 0xaf, 0xa1, 0x71, 0xab, 0x22, 0x05, 0xa8, 0x0f, 0xdc, 0x71, 0xdf,
	0xfb, 0x6c, 0x6b, 0xd4, 0x04, 0xdd, 0x1f, 0x7e, 0xf2, 0x6d, 0x42, 0x0d, 0xa8, 0xf6, 0xdf, 0x8e,
	0xec, 0x0a, 0x6d, 0x80, 0x31, 0x19, 0x4e, 0x26, 0xee, 0x87, 0xb1, 0x5d, 0x1d, 0xbc, 0xd9, 0xe5,
	0xa8, 0xed, 0x73, 0xd4, 0x8e, 0x39, 0x92, 0xab, 0x1c, 0xc9, 0x0f, 0x89, 0xe4, 0xb7, 0x44, 0xf2,
	0x47, 0x22, 0xf9, 0x2b, 0x91, 0xec, 0x24, 0x92, 0x7f, 0x12, 0xc9, 0x7f, 0x89, 0xda, 0x51, 0x22,
	0xf9, 0x79, 0x40, 0x6d, 0x77, 0x40, 0x6d, 0x7f, 0x40, 0x6d, 0x5a, 0x57, 0xbf, 0xf0, 0xd5, 0x75,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0x30, 0x1c, 0xf5, 0x16, 0x02, 0x00, 0x00,
}
