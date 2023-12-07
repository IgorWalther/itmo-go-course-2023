// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: note.proto

package note

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	NoteManager_CreateNote_FullMethodName     = "/note.NoteManager/CreateNote"
	NoteManager_GetAllNotes_FullMethodName    = "/note.NoteManager/GetAllNotes"
	NoteManager_GetNoteByID_FullMethodName    = "/note.NoteManager/GetNoteByID"
	NoteManager_DeleteNoteByID_FullMethodName = "/note.NoteManager/DeleteNoteByID"
)

// NoteManagerClient is the client API for NoteManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoteManagerClient interface {
	CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error)
	GetAllNotes(ctx context.Context, in *GetAllNotesRequest, opts ...grpc.CallOption) (*GetAllNotesResponse, error)
	GetNoteByID(ctx context.Context, in *GetNoteByIDRequest, opts ...grpc.CallOption) (*GetNoteByIDResponse, error)
	DeleteNoteByID(ctx context.Context, in *DeleteNoteByIDRequest, opts ...grpc.CallOption) (*DeleteNoteByIDResponse, error)
}

type noteManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewNoteManagerClient(cc grpc.ClientConnInterface) NoteManagerClient {
	return &noteManagerClient{cc}
}

func (c *noteManagerClient) CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error) {
	out := new(CreateNoteResponse)
	err := c.cc.Invoke(ctx, NoteManager_CreateNote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteManagerClient) GetAllNotes(ctx context.Context, in *GetAllNotesRequest, opts ...grpc.CallOption) (*GetAllNotesResponse, error) {
	out := new(GetAllNotesResponse)
	err := c.cc.Invoke(ctx, NoteManager_GetAllNotes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteManagerClient) GetNoteByID(ctx context.Context, in *GetNoteByIDRequest, opts ...grpc.CallOption) (*GetNoteByIDResponse, error) {
	out := new(GetNoteByIDResponse)
	err := c.cc.Invoke(ctx, NoteManager_GetNoteByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteManagerClient) DeleteNoteByID(ctx context.Context, in *DeleteNoteByIDRequest, opts ...grpc.CallOption) (*DeleteNoteByIDResponse, error) {
	out := new(DeleteNoteByIDResponse)
	err := c.cc.Invoke(ctx, NoteManager_DeleteNoteByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoteManagerServer is the server API for NoteManager service.
// All implementations must embed UnimplementedNoteManagerServer
// for forward compatibility
type NoteManagerServer interface {
	CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error)
	GetAllNotes(context.Context, *GetAllNotesRequest) (*GetAllNotesResponse, error)
	GetNoteByID(context.Context, *GetNoteByIDRequest) (*GetNoteByIDResponse, error)
	DeleteNoteByID(context.Context, *DeleteNoteByIDRequest) (*DeleteNoteByIDResponse, error)
	mustEmbedUnimplementedNoteManagerServer()
}

// UnimplementedNoteManagerServer must be embedded to have forward compatible implementations.
type UnimplementedNoteManagerServer struct {
}

func (UnimplementedNoteManagerServer) CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (UnimplementedNoteManagerServer) GetAllNotes(context.Context, *GetAllNotesRequest) (*GetAllNotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllNotes not implemented")
}
func (UnimplementedNoteManagerServer) GetNoteByID(context.Context, *GetNoteByIDRequest) (*GetNoteByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNoteByID not implemented")
}
func (UnimplementedNoteManagerServer) DeleteNoteByID(context.Context, *DeleteNoteByIDRequest) (*DeleteNoteByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNoteByID not implemented")
}
func (UnimplementedNoteManagerServer) mustEmbedUnimplementedNoteManagerServer() {}

// UnsafeNoteManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoteManagerServer will
// result in compilation errors.
type UnsafeNoteManagerServer interface {
	mustEmbedUnimplementedNoteManagerServer()
}

func RegisterNoteManagerServer(s grpc.ServiceRegistrar, srv NoteManagerServer) {
	s.RegisterService(&NoteManager_ServiceDesc, srv)
}

func _NoteManager_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteManagerServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteManager_CreateNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteManagerServer).CreateNote(ctx, req.(*CreateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteManager_GetAllNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllNotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteManagerServer).GetAllNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteManager_GetAllNotes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteManagerServer).GetAllNotes(ctx, req.(*GetAllNotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteManager_GetNoteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNoteByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteManagerServer).GetNoteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteManager_GetNoteByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteManagerServer).GetNoteByID(ctx, req.(*GetNoteByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteManager_DeleteNoteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNoteByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteManagerServer).DeleteNoteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteManager_DeleteNoteByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteManagerServer).DeleteNoteByID(ctx, req.(*DeleteNoteByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NoteManager_ServiceDesc is the grpc.ServiceDesc for NoteManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NoteManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "note.NoteManager",
	HandlerType: (*NoteManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNote",
			Handler:    _NoteManager_CreateNote_Handler,
		},
		{
			MethodName: "GetAllNotes",
			Handler:    _NoteManager_GetAllNotes_Handler,
		},
		{
			MethodName: "GetNoteByID",
			Handler:    _NoteManager_GetNoteByID_Handler,
		},
		{
			MethodName: "DeleteNoteByID",
			Handler:    _NoteManager_DeleteNoteByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "note.proto",
}