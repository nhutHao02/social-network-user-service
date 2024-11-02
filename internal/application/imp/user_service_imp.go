package imp

import (
	"github.com/nhutHao02/social-network-user-service/internal/application"
	"github.com/nhutHao02/social-network-user-service/internal/domain/interface/user"
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

// Method implements application.UserSerVice.
func (s *userService) Method() {
	panic("unimplemented")
}
