// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package stream

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

// FileStreamServiceClient is the client API for FileStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileStreamServiceClient interface {
	StreamFile(ctx context.Context, in *StreamFileRequest, opts ...grpc.CallOption) (FileStreamService_StreamFileClient, error)
}

type fileStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileStreamServiceClient(cc grpc.ClientConnInterface) FileStreamServiceClient {
	return &fileStreamServiceClient{cc}
}

func (c *fileStreamServiceClient) StreamFile(ctx context.Context, in *StreamFileRequest, opts ...grpc.CallOption) (FileStreamService_StreamFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileStreamService_ServiceDesc.Streams[0], "/main.FileStreamService/StreamFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileStreamServiceStreamFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileStreamService_StreamFileClient interface {
	Recv() (*StreamFileResponse, error)
	grpc.ClientStream
}

type fileStreamServiceStreamFileClient struct {
	grpc.ClientStream
}

func (x *fileStreamServiceStreamFileClient) Recv() (*StreamFileResponse, error) {
	m := new(StreamFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileStreamServiceServer is the server API for FileStreamService service.
// All implementations must embed UnimplementedFileStreamServiceServer
// for forward compatibility
type FileStreamServiceServer interface {
	StreamFile(*StreamFileRequest, FileStreamService_StreamFileServer) error
	mustEmbedUnimplementedFileStreamServiceServer()
}

// UnimplementedFileStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileStreamServiceServer struct {
}

func (UnimplementedFileStreamServiceServer) StreamFile(*StreamFileRequest, FileStreamService_StreamFileServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamFile not implemented")
}
func (UnimplementedFileStreamServiceServer) mustEmbedUnimplementedFileStreamServiceServer() {}

// UnsafeFileStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileStreamServiceServer will
// result in compilation errors.
type UnsafeFileStreamServiceServer interface {
	mustEmbedUnimplementedFileStreamServiceServer()
}

func RegisterFileStreamServiceServer(s grpc.ServiceRegistrar, srv FileStreamServiceServer) {
	s.RegisterService(&FileStreamService_ServiceDesc, srv)
}

func _FileStreamService_StreamFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileStreamServiceServer).StreamFile(m, &fileStreamServiceStreamFileServer{stream})
}

type FileStreamService_StreamFileServer interface {
	Send(*StreamFileResponse) error
	grpc.ServerStream
}

type fileStreamServiceStreamFileServer struct {
	grpc.ServerStream
}

func (x *fileStreamServiceStreamFileServer) Send(m *StreamFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

// FileStreamService_ServiceDesc is the grpc.ServiceDesc for FileStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.FileStreamService",
	HandlerType: (*FileStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamFile",
			Handler:       _FileStreamService_StreamFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stream.proto",
}

// JobLogsServiceClient is the client API for JobLogsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobLogsServiceClient interface {
	// GetJobLogs streams the logs of the given Job.
	GetJobLogs(ctx context.Context, in *GetJobLogsRequest, opts ...grpc.CallOption) (JobLogsService_GetJobLogsClient, error)
}

type jobLogsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobLogsServiceClient(cc grpc.ClientConnInterface) JobLogsServiceClient {
	return &jobLogsServiceClient{cc}
}

func (c *jobLogsServiceClient) GetJobLogs(ctx context.Context, in *GetJobLogsRequest, opts ...grpc.CallOption) (JobLogsService_GetJobLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &JobLogsService_ServiceDesc.Streams[0], "/main.JobLogsService/GetJobLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &jobLogsServiceGetJobLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type JobLogsService_GetJobLogsClient interface {
	Recv() (*GetJobLogsResponse, error)
	grpc.ClientStream
}

type jobLogsServiceGetJobLogsClient struct {
	grpc.ClientStream
}

func (x *jobLogsServiceGetJobLogsClient) Recv() (*GetJobLogsResponse, error) {
	m := new(GetJobLogsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// JobLogsServiceServer is the server API for JobLogsService service.
// All implementations must embed UnimplementedJobLogsServiceServer
// for forward compatibility
type JobLogsServiceServer interface {
	// GetJobLogs streams the logs of the given Job.
	GetJobLogs(*GetJobLogsRequest, JobLogsService_GetJobLogsServer) error
	mustEmbedUnimplementedJobLogsServiceServer()
}

// UnimplementedJobLogsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJobLogsServiceServer struct {
}

func (UnimplementedJobLogsServiceServer) GetJobLogs(*GetJobLogsRequest, JobLogsService_GetJobLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetJobLogs not implemented")
}
func (UnimplementedJobLogsServiceServer) mustEmbedUnimplementedJobLogsServiceServer() {}

// UnsafeJobLogsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobLogsServiceServer will
// result in compilation errors.
type UnsafeJobLogsServiceServer interface {
	mustEmbedUnimplementedJobLogsServiceServer()
}

func RegisterJobLogsServiceServer(s grpc.ServiceRegistrar, srv JobLogsServiceServer) {
	s.RegisterService(&JobLogsService_ServiceDesc, srv)
}

func _JobLogsService_GetJobLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetJobLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(JobLogsServiceServer).GetJobLogs(m, &jobLogsServiceGetJobLogsServer{stream})
}

type JobLogsService_GetJobLogsServer interface {
	Send(*GetJobLogsResponse) error
	grpc.ServerStream
}

type jobLogsServiceGetJobLogsServer struct {
	grpc.ServerStream
}

func (x *jobLogsServiceGetJobLogsServer) Send(m *GetJobLogsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// JobLogsService_ServiceDesc is the grpc.ServiceDesc for JobLogsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobLogsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.JobLogsService",
	HandlerType: (*JobLogsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetJobLogs",
			Handler:       _JobLogsService_GetJobLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stream.proto",
}
