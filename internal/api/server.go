package api

import (
	grpcServer "github.com/nhutHao02/social-network-user-service/internal/api/grpc"
	"github.com/nhutHao02/social-network-user-service/internal/api/http"
)

type Server struct {
	// httpt server
	HTTPServer *http.HTTPServer
	// grpc server
	GRPCServer *grpcServer.GRPCServer
}

func NewSerVer(httpServer *http.HTTPServer, gRPCServer *grpcServer.GRPCServer) *Server {
	return &Server{
		HTTPServer: httpServer,
		GRPCServer: gRPCServer,
	}
}
