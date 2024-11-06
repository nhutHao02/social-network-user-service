package user

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-user-service/internal/domain/interface/user"
	"go.uber.org/zap"
)

type userQueryRepository struct {
	db *sqlx.DB
}

func (repo *userQueryRepository) CheckUserExisted(ctx context.Context, email string) (bool, error) {
	var count int
	query := "SELECT EXISTS(SELECT * FROM `user` u WHERE email = ?);"

	err := repo.db.GetContext(ctx, &count, query, email)
	if err != nil {
		logger.Error("CheckUserExisted: check exist user error: ", zap.Error(err))
		return false, err
	}
	if count < 1 {
		return false, nil
	}
	return true, nil
}

func NewUserQueryRepository(db *sqlx.DB) user.UserQueryRepository {
	return &userQueryRepository{db: db}
}
