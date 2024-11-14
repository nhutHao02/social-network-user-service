package grpc

import (
	"context"

	pb "github.com/nhutHao02/social-network-user-service/pkg/grpc"
)

func (s *GRPCServer) GetUserInfo(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{
		Id:       123,
		Email:    "testmail",
		FullName: "fulna",
		UrlAvt:   "url",
	}, nil
}
