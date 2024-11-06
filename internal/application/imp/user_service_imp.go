package imp

import (
	"context"

	resError "github.com/nhutHao02/social-network-common-service/utils/error"
	"github.com/nhutHao02/social-network-user-service/internal/application"
	"github.com/nhutHao02/social-network-user-service/internal/domain/interface/user"
	"github.com/nhutHao02/social-network-user-service/internal/domain/model"
	"github.com/nhutHao02/social-network-user-service/pkg/redis"
)

type userService struct {
	userQueryRepository   user.UserQueryRepository
	userCommandRepository user.UserCommandRepository
	cache                 *redis.RedisClient
}

func NewUserService(
	userQueryRepository user.UserQueryRepository,
	userCommandRepository user.UserCommandRepository,
	cache *redis.RedisClient,
) application.UserSerVice {
	return &userService{userQueryRepository: userQueryRepository, userCommandRepository: userCommandRepository, cache: cache}
}

func (u *userService) RegisterUser(c context.Context, req model.SignUpRequest) (bool, error) {

	existed, err := u.userQueryRepository.CheckUserExisted(c, req.Email)
	if err != nil {
		return false, err
	}

	if existed {
		return false, resError.NewResError(nil, "Email existed")
	}

	success, err := u.userCommandRepository.RegisterUser(c, req)

	if err != nil {
		return false, err
	}

	return success, nil

}
