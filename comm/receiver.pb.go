// Code generated by protoc-gen-go. DO NOT EDIT.
// source: receiver.proto

/*
Package comm is a generated protocol buffer package.

It is generated from these files:
	receiver.proto

It has these top-level messages:
	Msg
	BinMsg
	Conf
*/
package comm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Msg struct {
	Replica   string            `protobuf:"bytes,1,opt,name=replica" json:"replica,omitempty"`
	Vclock    map[string]uint32 `protobuf:"bytes,2,rep,name=vclock" json:"vclock,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Operation string            `protobuf:"bytes,3,opt,name=operation" json:"operation,omitempty"`
	Create    *Msg_CREATE       `protobuf:"bytes,4,opt,name=create" json:"create,omitempty"`
	Delete    *Msg_DELETE       `protobuf:"bytes,5,opt,name=delete" json:"delete,omitempty"`
	Append    *Msg_APPEND       `protobuf:"bytes,6,opt,name=append" json:"append,omitempty"`
	Expunge   *Msg_EXPUNGE      `protobuf:"bytes,7,opt,name=expunge" json:"expunge,omitempty"`
	Store     *Msg_STORE        `protobuf:"bytes,8,opt,name=store" json:"store,omitempty"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Msg) GetReplica() string {
	if m != nil {
		return m.Replica
	}
	return ""
}

func (m *Msg) GetVclock() map[string]uint32 {
	if m != nil {
		return m.Vclock
	}
	return nil
}

func (m *Msg) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *Msg) GetCreate() *Msg_CREATE {
	if m != nil {
		return m.Create
	}
	return nil
}

func (m *Msg) GetDelete() *Msg_DELETE {
	if m != nil {
		return m.Delete
	}
	return nil
}

func (m *Msg) GetAppend() *Msg_APPEND {
	if m != nil {
		return m.Append
	}
	return nil
}

func (m *Msg) GetExpunge() *Msg_EXPUNGE {
	if m != nil {
		return m.Expunge
	}
	return nil
}

func (m *Msg) GetStore() *Msg_STORE {
	if m != nil {
		return m.Store
	}
	return nil
}

type Msg_CREATE struct {
	User    string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Mailbox string `protobuf:"bytes,2,opt,name=mailbox" json:"mailbox,omitempty"`
	AddTag  string `protobuf:"bytes,3,opt,name=addTag" json:"addTag,omitempty"`
}

func (m *Msg_CREATE) Reset()                    { *m = Msg_CREATE{} }
func (m *Msg_CREATE) String() string            { return proto.CompactTextString(m) }
func (*Msg_CREATE) ProtoMessage()               {}
func (*Msg_CREATE) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *Msg_CREATE) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Msg_CREATE) GetMailbox() string {
	if m != nil {
		return m.Mailbox
	}
	return ""
}

func (m *Msg_CREATE) GetAddTag() string {
	if m != nil {
		return m.AddTag
	}
	return ""
}

type Msg_DELETE struct {
	User     string   `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Mailbox  string   `protobuf:"bytes,2,opt,name=mailbox" json:"mailbox,omitempty"`
	RmvTag   []string `protobuf:"bytes,3,rep,name=rmvTag" json:"rmvTag,omitempty"`
	RmvMails []string `protobuf:"bytes,4,rep,name=rmvMails" json:"rmvMails,omitempty"`
}

func (m *Msg_DELETE) Reset()                    { *m = Msg_DELETE{} }
func (m *Msg_DELETE) String() string            { return proto.CompactTextString(m) }
func (*Msg_DELETE) ProtoMessage()               {}
func (*Msg_DELETE) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *Msg_DELETE) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Msg_DELETE) GetMailbox() string {
	if m != nil {
		return m.Mailbox
	}
	return ""
}

func (m *Msg_DELETE) GetRmvTag() []string {
	if m != nil {
		return m.RmvTag
	}
	return nil
}

func (m *Msg_DELETE) GetRmvMails() []string {
	if m != nil {
		return m.RmvMails
	}
	return nil
}

type Msg_APPEND struct {
	User       string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Mailbox    string `protobuf:"bytes,2,opt,name=mailbox" json:"mailbox,omitempty"`
	AddTag     string `protobuf:"bytes,3,opt,name=addTag" json:"addTag,omitempty"`
	AddContent []byte `protobuf:"bytes,4,opt,name=addContent,proto3" json:"addContent,omitempty"`
}

func (m *Msg_APPEND) Reset()                    { *m = Msg_APPEND{} }
func (m *Msg_APPEND) String() string            { return proto.CompactTextString(m) }
func (*Msg_APPEND) ProtoMessage()               {}
func (*Msg_APPEND) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 2} }

func (m *Msg_APPEND) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Msg_APPEND) GetMailbox() string {
	if m != nil {
		return m.Mailbox
	}
	return ""
}

func (m *Msg_APPEND) GetAddTag() string {
	if m != nil {
		return m.AddTag
	}
	return ""
}

func (m *Msg_APPEND) GetAddContent() []byte {
	if m != nil {
		return m.AddContent
	}
	return nil
}

type Msg_EXPUNGE struct {
	User    string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Mailbox string `protobuf:"bytes,2,opt,name=mailbox" json:"mailbox,omitempty"`
	RmvTag  string `protobuf:"bytes,3,opt,name=rmvTag" json:"rmvTag,omitempty"`
	AddTag  string `protobuf:"bytes,4,opt,name=addTag" json:"addTag,omitempty"`
}

func (m *Msg_EXPUNGE) Reset()                    { *m = Msg_EXPUNGE{} }
func (m *Msg_EXPUNGE) String() string            { return proto.CompactTextString(m) }
func (*Msg_EXPUNGE) ProtoMessage()               {}
func (*Msg_EXPUNGE) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 3} }

func (m *Msg_EXPUNGE) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Msg_EXPUNGE) GetMailbox() string {
	if m != nil {
		return m.Mailbox
	}
	return ""
}

func (m *Msg_EXPUNGE) GetRmvTag() string {
	if m != nil {
		return m.RmvTag
	}
	return ""
}

func (m *Msg_EXPUNGE) GetAddTag() string {
	if m != nil {
		return m.AddTag
	}
	return ""
}

type Msg_STORE struct {
	User       string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Mailbox    string `protobuf:"bytes,2,opt,name=mailbox" json:"mailbox,omitempty"`
	RmvTag     string `protobuf:"bytes,3,opt,name=rmvTag" json:"rmvTag,omitempty"`
	AddTag     string `protobuf:"bytes,4,opt,name=addTag" json:"addTag,omitempty"`
	AddContent []byte `protobuf:"bytes,5,opt,name=addContent,proto3" json:"addContent,omitempty"`
}

func (m *Msg_STORE) Reset()                    { *m = Msg_STORE{} }
func (m *Msg_STORE) String() string            { return proto.CompactTextString(m) }
func (*Msg_STORE) ProtoMessage()               {}
func (*Msg_STORE) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 4} }

func (m *Msg_STORE) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Msg_STORE) GetMailbox() string {
	if m != nil {
		return m.Mailbox
	}
	return ""
}

func (m *Msg_STORE) GetRmvTag() string {
	if m != nil {
		return m.RmvTag
	}
	return ""
}

func (m *Msg_STORE) GetAddTag() string {
	if m != nil {
		return m.AddTag
	}
	return ""
}

func (m *Msg_STORE) GetAddContent() []byte {
	if m != nil {
		return m.AddContent
	}
	return nil
}

type BinMsg struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *BinMsg) Reset()                    { *m = BinMsg{} }
func (m *BinMsg) String() string            { return proto.CompactTextString(m) }
func (*BinMsg) ProtoMessage()               {}
func (*BinMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BinMsg) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Conf struct {
	Status uint32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *Conf) Reset()                    { *m = Conf{} }
func (m *Conf) String() string            { return proto.CompactTextString(m) }
func (*Conf) ProtoMessage()               {}
func (*Conf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Conf) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*Msg)(nil), "comm.Msg")
	proto.RegisterType((*Msg_CREATE)(nil), "comm.Msg.CREATE")
	proto.RegisterType((*Msg_DELETE)(nil), "comm.Msg.DELETE")
	proto.RegisterType((*Msg_APPEND)(nil), "comm.Msg.APPEND")
	proto.RegisterType((*Msg_EXPUNGE)(nil), "comm.Msg.EXPUNGE")
	proto.RegisterType((*Msg_STORE)(nil), "comm.Msg.STORE")
	proto.RegisterType((*BinMsg)(nil), "comm.BinMsg")
	proto.RegisterType((*Conf)(nil), "comm.Conf")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Receiver service

type ReceiverClient interface {
	Incoming(ctx context.Context, in *BinMsg, opts ...grpc.CallOption) (*Conf, error)
}

type receiverClient struct {
	cc *grpc.ClientConn
}

func NewReceiverClient(cc *grpc.ClientConn) ReceiverClient {
	return &receiverClient{cc}
}

func (c *receiverClient) Incoming(ctx context.Context, in *BinMsg, opts ...grpc.CallOption) (*Conf, error) {
	out := new(Conf)
	err := grpc.Invoke(ctx, "/comm.Receiver/Incoming", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Receiver service

type ReceiverServer interface {
	Incoming(context.Context, *BinMsg) (*Conf, error)
}

func RegisterReceiverServer(s *grpc.Server, srv ReceiverServer) {
	s.RegisterService(&_Receiver_serviceDesc, srv)
}

func _Receiver_Incoming_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BinMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReceiverServer).Incoming(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comm.Receiver/Incoming",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReceiverServer).Incoming(ctx, req.(*BinMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Receiver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comm.Receiver",
	HandlerType: (*ReceiverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Incoming",
			Handler:    _Receiver_Incoming_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "receiver.proto",
}

func init() { proto.RegisterFile("receiver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x35, 0xcd, 0x57, 0xf7, 0xb6, 0xab, 0xeb, 0xa0, 0x32, 0x84, 0x65, 0x29, 0x05, 0xa5, 0x20,
	0xf6, 0xa1, 0xbe, 0xa8, 0x6f, 0x6b, 0x37, 0x88, 0x60, 0x6b, 0x19, 0xab, 0xf8, 0x3a, 0x9b, 0x8c,
	0x21, 0x6e, 0x32, 0x13, 0x26, 0xd3, 0xb0, 0xfb, 0x03, 0xfc, 0x0b, 0xfe, 0x5e, 0x99, 0x8f, 0x35,
	0xb5, 0x3e, 0x89, 0xfa, 0x36, 0xe7, 0xde, 0x93, 0x93, 0x73, 0xcf, 0xe4, 0x06, 0xee, 0x4a, 0x96,
	0xb1, 0xb2, 0x63, 0x72, 0xde, 0x48, 0xa1, 0x04, 0x0a, 0x32, 0x51, 0xd7, 0xd3, 0xef, 0x31, 0xf8,
	0xab, 0xb6, 0x40, 0x18, 0x62, 0xc9, 0x9a, 0xaa, 0xcc, 0x28, 0xf6, 0x26, 0xde, 0xec, 0x88, 0xdc,
	0x42, 0xf4, 0x0c, 0xa2, 0x2e, 0xab, 0x44, 0x76, 0x85, 0x07, 0x13, 0x7f, 0x36, 0x5a, 0x3c, 0x9c,
	0xeb, 0x07, 0xe7, 0xab, 0xb6, 0x98, 0x7f, 0x32, 0xf5, 0x94, 0x2b, 0x79, 0x43, 0x1c, 0x09, 0x9d,
	0xc2, 0x91, 0x68, 0x98, 0xa4, 0xaa, 0x14, 0x1c, 0xfb, 0x46, 0xaa, 0x2f, 0xa0, 0x19, 0x44, 0x99,
	0x64, 0x54, 0x31, 0x1c, 0x4c, 0xbc, 0xd9, 0x68, 0x71, 0xd2, 0x8b, 0x2d, 0x49, 0x7a, 0xbe, 0x4d,
	0x89, 0xeb, 0x6b, 0x66, 0xce, 0x2a, 0xa6, 0x18, 0x0e, 0x0f, 0x99, 0x17, 0xe9, 0xbb, 0x54, 0x33,
	0x6d, 0x5f, 0x33, 0x69, 0xd3, 0x30, 0x9e, 0xe3, 0xe8, 0x90, 0x79, 0xbe, 0xd9, 0xa4, 0xeb, 0x0b,
	0xe2, 0xfa, 0xe8, 0x29, 0xc4, 0xec, 0xba, 0xd9, 0xf1, 0x82, 0xe1, 0xd8, 0x50, 0xef, 0xf7, 0xd4,
	0xf4, 0xf3, 0xe6, 0xe3, 0xfa, 0x4d, 0x4a, 0x6e, 0x19, 0xe8, 0x31, 0x84, 0xad, 0x12, 0x92, 0xe1,
	0xa1, 0xa1, 0xde, 0xeb, 0xa9, 0x1f, 0xb6, 0xef, 0x49, 0x4a, 0x6c, 0x37, 0x59, 0x43, 0x64, 0x9d,
	0x23, 0x04, 0xc1, 0xae, 0x65, 0xd2, 0xe5, 0x67, 0xce, 0x3a, 0xd6, 0x9a, 0x96, 0xd5, 0xa5, 0xb8,
	0xc6, 0x03, 0x1b, 0xab, 0x83, 0xe8, 0x11, 0x44, 0x34, 0xcf, 0xb7, 0xb4, 0x70, 0x21, 0x39, 0x94,
	0x7c, 0x85, 0xc8, 0xce, 0xf7, 0xe7, 0x7a, 0xb2, 0xee, 0xac, 0x9e, 0xaf, 0xf5, 0x2c, 0x42, 0x09,
	0x0c, 0x65, 0xdd, 0xad, 0x68, 0x59, 0xb5, 0x38, 0x30, 0x9d, 0x9f, 0x38, 0xe1, 0x10, 0xd9, 0x84,
	0xfe, 0x8d, 0x77, 0x74, 0x06, 0x40, 0xf3, 0x7c, 0x29, 0xb8, 0x62, 0x5c, 0x99, 0x1b, 0x1e, 0x93,
	0xbd, 0x4a, 0x52, 0x40, 0xec, 0x62, 0xfe, 0x8b, 0xe1, 0xbc, 0xbd, 0xe1, 0x7a, 0x23, 0xc1, 0x2f,
	0x21, 0x7e, 0xf3, 0x20, 0x34, 0xb7, 0xf4, 0x7f, 0xdf, 0x73, 0x30, 0x70, 0xf8, 0xdb, 0xc0, 0x2f,
	0x61, 0xb4, 0xb7, 0x23, 0xe8, 0x04, 0xfc, 0x2b, 0x76, 0xe3, 0xbc, 0xe8, 0x23, 0x7a, 0x00, 0x61,
	0x47, 0xab, 0x1d, 0x33, 0x46, 0x8e, 0x89, 0x05, 0xaf, 0x06, 0x2f, 0xbc, 0xe9, 0x29, 0x44, 0xaf,
	0x4b, 0xae, 0x57, 0x13, 0x41, 0x90, 0x53, 0x65, 0xf7, 0x72, 0x4c, 0xcc, 0x79, 0x7a, 0x06, 0xc1,
	0x52, 0xf0, 0x2f, 0xda, 0x58, 0xab, 0xa8, 0xda, 0xb5, 0xa6, 0x7b, 0x4c, 0x1c, 0x5a, 0x2c, 0x60,
	0x48, 0xdc, 0xba, 0xa3, 0x27, 0x30, 0x7c, 0xcb, 0x33, 0x51, 0x97, 0xbc, 0x40, 0x63, 0xfb, 0x15,
	0x5b, 0xe5, 0x04, 0x2c, 0xd2, 0x4a, 0xd3, 0x3b, 0x97, 0x91, 0xf9, 0x2f, 0x3c, 0xff, 0x11, 0x00,
	0x00, 0xff, 0xff, 0x90, 0xc2, 0x15, 0xd4, 0x29, 0x04, 0x00, 0x00,
}
