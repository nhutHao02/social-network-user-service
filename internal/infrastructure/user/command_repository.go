package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/nhutHao02/social-network-user-service/internal/domain/interface/user"
)

type userCommandRepository struct {
	db *sqlx.DB
}

func NewUserCommandRepository(db *sqlx.DB) user.UserCommandRepository {
	return &userCommandRepository{db: db}
}
