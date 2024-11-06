package user

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-user-service/internal/domain/interface/user"
	"github.com/nhutHao02/social-network-user-service/internal/domain/model"
	"go.uber.org/zap"
)

type userCommandRepository struct {
	db *sqlx.DB
}

// RegisterUser implements user.UserCommandRepository.
func (repo *userCommandRepository) RegisterUser(ctx context.Context, req model.SignUpRequest) (bool, error) {
	query := "INSERT INTO `user` (Email, Password) VALUES (:Email, :Password)"

	_, err := repo.db.NamedExecContext(ctx, query, req)
	if err != nil {
		logger.Error("RegisterUser: Insert new usser error: ", zap.Error(err))
		return false, err
	}
	return true, nil
}

func NewUserCommandRepository(db *sqlx.DB) user.UserCommandRepository {
	return &userCommandRepository{db: db}
}
