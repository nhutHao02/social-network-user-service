package user

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-user-service/internal/domain/entity"
	"github.com/nhutHao02/social-network-user-service/internal/domain/interface/user"
	"github.com/nhutHao02/social-network-user-service/internal/domain/model"
	"go.uber.org/zap"
)

type userQueryRepository struct {
	db *sqlx.DB
}

// GetUserInfo implements user.UserQueryRepository.
func (repo *userQueryRepository) GetUserInfo(ctx context.Context, userID int) (*model.UserInfoResponse, error) {
	var res model.UserInfoResponse
	query := `select 
				u.ID ,
				u.Email,
				u.FullName,
				u.Sex,
				u.Bio,
				u.UrlAvt,
				u.UrlBackground,
				u.CreatedAt,
				u.UpdatedAt,
				l.ID as 'location.ID',
				l.City as 'location.City',
				l.District as 'location.District',
				l.Ward as 'location.Ward',
				l.Description as 'location.Description',
				l.CreatedAt as 'location.CreatedAt',
				l.UpdatedAt as 'location.UpdatedAt',
				l.DeletedAt as 'location.DeletedAt'
				from user u
				left join location l
				on u.LocationID = l.ID
				where u.ID = ? and u.DeletedAt is null and l.DeletedAt is null`

	err := repo.db.GetContext(ctx, &res, query, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("GetUserInfo: repo get user info error", zap.Error(err))
			return nil, nil
		}
		logger.Error("GetUserInfo: repo get user info error", zap.Error(err))
		return nil, err
	}
	return &res, nil

}

// Login implements user.UserQueryRepository.
func (repo *userQueryRepository) Login(ctx context.Context, req model.LoginRequest) (*entity.User, error) {
	var user entity.User
	query := "SELECT * FROM `user` u WHERE Email = :Email AND DeletedAt IS NULL"
	queryString, queryArgs, err := repo.db.BindNamed(query, req)
	if err != nil {
		logger.Error("Login: BindNamed failure", zap.Error(err))
		return nil, err
	}

	err = repo.db.GetContext(ctx, &user, queryString, queryArgs...)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		logger.Error("Login: Login failure ", zap.Error(err))
		return nil, err
	}
	return &user, nil

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
