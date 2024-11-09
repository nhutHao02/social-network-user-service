package user

import (
	"context"

	"github.com/nhutHao02/social-network-user-service/internal/domain/entity"
	"github.com/nhutHao02/social-network-user-service/internal/domain/model"
)

type UserQueryRepository interface {
	CheckUserExisted(ctx context.Context, email string) (bool, error)
	Login(ctx context.Context, req model.LoginRequest) (*entity.User, error)
	GetUserInfo(ctx context.Context, userID int) (*model.UserInfoResponse, error)
}

type UserCommandRepository interface {
	RegisterUser(ctx context.Context, req model.SignUpRequest) (bool, error)
	UpdateUserInfo(ctx context.Context, req model.UserUpdateRequest) (bool, error)
}
