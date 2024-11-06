package application

import (
	"context"

	"github.com/nhutHao02/social-network-user-service/internal/domain/model"
)

type UserSerVice interface {
	RegisterUser(c context.Context, req model.SignUpRequest) (bool, error)
}
