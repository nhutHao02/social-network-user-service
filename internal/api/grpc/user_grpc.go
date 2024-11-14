package grpc

import (
	"context"

	pb "github.com/nhutHao02/social-network-user-service/pkg/grpc"
)

func (s *GRPCServer) GetUserInfo(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.UserService.GetUserInfo(ctx, int(req.UserID))
	if err != nil {
		return nil, err
	}
	fullName := ""
	if user.FullName != nil {
		fullName = *user.FullName
	}

	urlAvt := ""
	if user.UrlAvt != nil {
		urlAvt = *user.UrlAvt
	}

	return &pb.GetUserResponse{
		Id:       int64(user.ID),
		Email:    user.Email,
		FullName: fullName,
		UrlAvt:   urlAvt,
	}, nil
}
