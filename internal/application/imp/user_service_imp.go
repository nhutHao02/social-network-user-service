package imp

import "github.com/nhutHao02/social-network-user-service/internal/application"

type userService struct {
}

func NewUserService() application.UserSerVice {
	return &userService{}
}
