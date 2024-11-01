package v1

import "github.com/nhutHao02/social-network-user-service/internal/application"

type UserHandler struct {
	userService application.UserSerVice
}

func NewUserHandler(userService application.UserSerVice) *UserHandler {
	return &UserHandler{userService: userService}
}
