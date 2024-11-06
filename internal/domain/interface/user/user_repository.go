package user

import (
	"context"

	"github.com/nhutHao02/social-network-user-service/internal/domain/model"
)

type UserQueryRepository interface {
	CheckUserExisted(ctx context.Context, email string) (bool, error)
}

type UserCommandRepository interface {
	RegisterUser(ctx context.Context, req model.SignUpRequest) (bool, error)
}
