package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/nhutHao02/social-network-user-service/internal/domain/interface/user"
)

type userQueryRepository struct {
	db *sqlx.DB
}

func NewUserQueryRepository(db *sqlx.DB) user.UserQueryRepository {
	return &userQueryRepository{db: db}
}
