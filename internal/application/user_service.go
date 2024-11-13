package application

import (
	"context"

	"github.com/nhutHao02/social-network-user-service/internal/domain/model"
)

type UserSerVice interface {
	RegisterUser(c context.Context, req model.SignUpRequest) (bool, error)
	Login(c context.Context, req model.LoginRequest) (model.LoginResponse, error)
	GetUserInfo(c context.Context, userID int) (*model.UserInfoResponse, error)
	UpdateUserInfo(c context.Context, req model.UserUpdateRequest) (bool, error)
	ChangePassword(c context.Context, req model.UserUpdatePassRequest) (bool, error)
	Follow(c context.Context, req model.FollowRequest) (bool, error)
	UnFollow(c context.Context, req model.FollowRequest) (bool, error)
	GetFollow(c context.Context, idParam model.FollowIDParam, isFollower bool) (*model.FollowResponse, error)
}
