package user

import (
	"context"

	"github.com/nhutHao02/social-network-user-service/internal/domain/entity"
	"github.com/nhutHao02/social-network-user-service/internal/domain/model"
)

type UserQueryRepository interface {
	CheckUserExisted(ctx context.Context, email string) (bool, error)
	Login(ctx context.Context, req model.LoginRequest) (*entity.User, error)
	GetUserInfo(ctx context.Context, userID int64) (*model.UserInfoResponse, error)
	CheckUserExistByID(ctx context.Context, ID uint64) (bool, error)
	GetFollow(ctx context.Context, id uint64, isFollower bool) (*model.FollowResponse, error)
}

type UserCommandRepository interface {
	RegisterUser(ctx context.Context, req model.SignUpRequest) (bool, error)
	UpdateUserInfo(ctx context.Context, req model.UserUpdateRequest) (bool, error)
	UpdatePassword(ctx context.Context, req model.UserUpdatePassRequest) (bool, error)
	Follow(ctx context.Context, req model.FollowRequest) (bool, error)
	UnFollow(ctx context.Context, req model.FollowRequest) (bool, error)
}
