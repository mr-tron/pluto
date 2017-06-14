// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

/*
Package imap is a generated protocol buffer package.

It is generated from these files:
	node.proto

It has these top-level messages:
	Context
	Confirmation
	Command
	Reply
*/
package imap

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

type Context struct {
	ClientID string `protobuf:"bytes,1,opt,name=clientID" json:"clientID,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=userName" json:"userName,omitempty"`
}

func (m *Context) Reset()                    { *m = Context{} }
func (m *Context) String() string            { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()               {}
func (*Context) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Context) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *Context) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type Confirmation struct {
	Status uint32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
}

func (m *Confirmation) Reset()                    { *m = Confirmation{} }
func (m *Confirmation) String() string            { return proto.CompactTextString(m) }
func (*Confirmation) ProtoMessage()               {}
func (*Confirmation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Confirmation) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type Command struct {
	Text     string           `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	ClientID string           `protobuf:"bytes,2,opt,name=clientID" json:"clientID,omitempty"`
	HasMsg   *Command_Message `protobuf:"bytes,3,opt,name=hasMsg" json:"hasMsg,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Command) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Command) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *Command) GetHasMsg() *Command_Message {
	if m != nil {
		return m.HasMsg
	}
	return nil
}

type Command_Message struct {
	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *Command_Message) Reset()                    { *m = Command_Message{} }
func (m *Command_Message) String() string            { return proto.CompactTextString(m) }
func (*Command_Message) ProtoMessage()               {}
func (*Command_Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *Command_Message) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type Reply struct {
	Text     string        `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	IsAppend *Reply_APPEND `protobuf:"bytes,2,opt,name=isAppend" json:"isAppend,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Reply) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Reply) GetIsAppend() *Reply_APPEND {
	if m != nil {
		return m.IsAppend
	}
	return nil
}

type Reply_APPEND struct {
	AwaitedNumBytes uint32 `protobuf:"varint,1,opt,name=awaitedNumBytes" json:"awaitedNumBytes,omitempty"`
}

func (m *Reply_APPEND) Reset()                    { *m = Reply_APPEND{} }
func (m *Reply_APPEND) String() string            { return proto.CompactTextString(m) }
func (*Reply_APPEND) ProtoMessage()               {}
func (*Reply_APPEND) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *Reply_APPEND) GetAwaitedNumBytes() uint32 {
	if m != nil {
		return m.AwaitedNumBytes
	}
	return 0
}

func init() {
	proto.RegisterType((*Context)(nil), "imap.Context")
	proto.RegisterType((*Confirmation)(nil), "imap.Confirmation")
	proto.RegisterType((*Command)(nil), "imap.Command")
	proto.RegisterType((*Command_Message)(nil), "imap.Command.Message")
	proto.RegisterType((*Reply)(nil), "imap.Reply")
	proto.RegisterType((*Reply_APPEND)(nil), "imap.Reply.APPEND")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Node service

type NodeClient interface {
	Prepare(ctx context.Context, in *Context, opts ...grpc.CallOption) (*Confirmation, error)
	Close(ctx context.Context, in *Context, opts ...grpc.CallOption) (*Confirmation, error)
	Select(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error)
	Create(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error)
	Delete(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error)
	List(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error)
	Append(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error)
	Expunge(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error)
	Store(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Prepare(ctx context.Context, in *Context, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := grpc.Invoke(ctx, "/imap.Node/Prepare", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Close(ctx context.Context, in *Context, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := grpc.Invoke(ctx, "/imap.Node/Close", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Select(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/imap.Node/Select", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Create(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/imap.Node/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Delete(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/imap.Node/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) List(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/imap.Node/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Append(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/imap.Node/Append", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Expunge(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/imap.Node/Expunge", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Store(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/imap.Node/Store", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	Prepare(context.Context, *Context) (*Confirmation, error)
	Close(context.Context, *Context) (*Confirmation, error)
	Select(context.Context, *Command) (*Reply, error)
	Create(context.Context, *Command) (*Reply, error)
	Delete(context.Context, *Command) (*Reply, error)
	List(context.Context, *Command) (*Reply, error)
	Append(context.Context, *Command) (*Reply, error)
	Expunge(context.Context, *Command) (*Reply, error)
	Store(context.Context, *Command) (*Reply, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Context)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/Prepare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Prepare(ctx, req.(*Context))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Context)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Close(ctx, req.(*Context))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Select_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Select(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/Select",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Select(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Create(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Delete(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).List(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Append_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Append(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/Append",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Append(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Expunge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Expunge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/Expunge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Expunge(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imap.Node/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Store(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "imap.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Prepare",
			Handler:    _Node_Prepare_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Node_Close_Handler,
		},
		{
			MethodName: "Select",
			Handler:    _Node_Select_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Node_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Node_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Node_List_Handler,
		},
		{
			MethodName: "Append",
			Handler:    _Node_Append_Handler,
		},
		{
			MethodName: "Expunge",
			Handler:    _Node_Expunge_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _Node_Store_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}

func init() { proto.RegisterFile("node.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x4e, 0xab, 0x40,
	0x14, 0x87, 0x2f, 0xbd, 0x14, 0x7a, 0x4f, 0xdb, 0xdc, 0x64, 0x12, 0x0d, 0x61, 0xd5, 0xa0, 0xd6,
	0x2e, 0x94, 0x05, 0x3e, 0x41, 0xa5, 0x5d, 0x98, 0x58, 0xd2, 0xd0, 0x27, 0x18, 0xcb, 0xb1, 0x92,
	0xc0, 0x0c, 0x99, 0x19, 0x62, 0xbb, 0xf2, 0x09, 0x7c, 0x61, 0x57, 0x86, 0xbf, 0x9a, 0x6a, 0x42,
	0x97, 0xbf, 0x33, 0x1f, 0xf3, 0x9d, 0x73, 0xc8, 0x00, 0x30, 0x1e, 0xa1, 0x9b, 0x09, 0xae, 0x38,
	0xd1, 0xe3, 0x94, 0x66, 0xce, 0x1c, 0x4c, 0x9f, 0x33, 0x85, 0x7b, 0x45, 0x6c, 0x18, 0x6c, 0x93,
	0x18, 0x99, 0x7a, 0x58, 0x58, 0xda, 0x44, 0x9b, 0xfd, 0x0b, 0xdb, 0x5c, 0x9c, 0xe5, 0x12, 0x45,
	0x40, 0x53, 0xb4, 0x7a, 0xd5, 0x59, 0x93, 0x9d, 0x29, 0x8c, 0x7c, 0xce, 0x9e, 0x63, 0x91, 0x52,
	0x15, 0x73, 0x46, 0xce, 0xc1, 0x90, 0x8a, 0xaa, 0x5c, 0x96, 0xb7, 0x8c, 0xc3, 0x3a, 0x39, 0xef,
	0x5a, 0xe1, 0x4a, 0x53, 0xca, 0x22, 0x42, 0x40, 0x2f, 0x9c, 0xb5, 0x47, 0xff, 0xe1, 0xef, 0x1d,
	0xf9, 0x6f, 0xc1, 0x78, 0xa1, 0x72, 0x25, 0x77, 0xd6, 0xdf, 0x89, 0x36, 0x1b, 0x7a, 0x67, 0x6e,
	0xd1, 0xbd, 0x5b, 0x5f, 0xe7, 0xae, 0x50, 0x4a, 0xba, 0xc3, 0xb0, 0x86, 0xec, 0x0b, 0x30, 0xeb,
	0x12, 0xb1, 0xc0, 0xdc, 0x16, 0x03, 0xb2, 0x4a, 0x36, 0x0a, 0x9b, 0xe8, 0xbc, 0x41, 0x3f, 0xc4,
	0x2c, 0x39, 0xfc, 0xda, 0x8c, 0x0b, 0x83, 0x58, 0xce, 0xb3, 0x0c, 0x59, 0x54, 0x36, 0x33, 0xf4,
	0x48, 0xa5, 0x2c, 0x3f, 0x71, 0xe7, 0xeb, 0xf5, 0x32, 0x58, 0x84, 0x2d, 0x63, 0x7b, 0x60, 0x54,
	0x35, 0x32, 0x83, 0xff, 0xf4, 0x95, 0xc6, 0x0a, 0xa3, 0x20, 0x4f, 0xef, 0x0f, 0x0a, 0x9b, 0x3d,
	0x1c, 0x97, 0xbd, 0x8f, 0x1e, 0xe8, 0x01, 0x8f, 0x90, 0xb8, 0x60, 0xae, 0x05, 0x66, 0x54, 0x20,
	0x19, 0x37, 0x83, 0x95, 0xff, 0xc4, 0x26, 0x6d, 0x6c, 0xf7, 0xeb, 0xfc, 0x21, 0x37, 0xd0, 0xf7,
	0x13, 0x2e, 0x4f, 0xa4, 0xa7, 0x60, 0x6c, 0x30, 0xc1, 0xad, 0xfa, 0xc2, 0xcb, 0xad, 0xd9, 0xc3,
	0x6f, 0x13, 0x55, 0x9c, 0x2f, 0x90, 0x2a, 0xec, 0xe6, 0x16, 0x98, 0x60, 0x27, 0x77, 0x09, 0xfa,
	0x63, 0x2c, 0x4f, 0xb0, 0x56, 0x2b, 0xec, 0xe0, 0xae, 0xc1, 0x5c, 0xee, 0xb3, 0x9c, 0xed, 0xba,
	0xb4, 0x57, 0xd0, 0xdf, 0x28, 0x2e, 0x3a, 0xb0, 0x27, 0xa3, 0x7c, 0x05, 0x77, 0x9f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x6e, 0xd8, 0xa1, 0x02, 0x13, 0x03, 0x00, 0x00,
}
