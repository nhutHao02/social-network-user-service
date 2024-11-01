package imp

import (
	"github.com/nhutHao02/social-network-user-service/internal/application"
	"github.com/nhutHao02/social-network-user-service/internal/domain/interface/user"
)

type userService struct {
	userQueryRepository   user.UserQueryRepository
	userCommandRepository user.UserCommandRepository
}

func NewUserService(
	userQueryRepository user.UserQueryRepository,
	userCommandRepository user.UserCommandRepository,
) application.UserSerVice {
	return &userService{userQueryRepository: userQueryRepository, userCommandRepository: userCommandRepository}
}

// Method implements application.UserSerVice.
func (u *userService) Method() {
	panic("unimplemented")
}
