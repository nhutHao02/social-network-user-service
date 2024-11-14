package grpc

import (
	"net"

	md "github.com/nhutHao02/social-network-common-service/middleware"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-user-service/config"
	"github.com/nhutHao02/social-network-user-service/internal/application"
	pb "github.com/nhutHao02/social-network-user-service/pkg/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedUserServiceServer
	Cfg         *config.Config
	UserService application.UserSerVice
}

func NewGRPCServer(
	cfg *config.Config,
	userService application.UserSerVice,
) *GRPCServer {
	return &GRPCServer{
		Cfg:         cfg,
		UserService: userService,
	}
}

func (s *GRPCServer) RunGRPCServer() error {
	lis, err := net.Listen("tcp", s.Cfg.GRPC.Port)
	if err != nil {
		logger.Fatal("Failed to listion grpc port", zap.Error(err))
	}
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(md.JWTUnaryInterceptor))

	pb.RegisterUserServiceServer(server, s)

	logger.Info("gRPC server listening at" + s.Cfg.GRPC.Port)

	if err := server.Serve(lis); err != nil {
		logger.Fatal("Failed to serve", zap.Error(err))
		return err
	}
	return nil
}
