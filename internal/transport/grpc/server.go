package grpc

import (
	proto "github.com/Imnarka/project-protos/proto/user"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedUserServiceServer
	service    proto.UserServiceServer
	GrpcServer *grpc.Server
}

func NewServer(service proto.UserServiceServer) *Server {
	grpcServer := grpc.NewServer()
	proto.RegisterUserServiceServer(grpcServer, service)
	return &Server{
		service:    service,
		GrpcServer: grpcServer,
	}
}
