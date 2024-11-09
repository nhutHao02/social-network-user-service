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

func insertNewLocation(ctx context.Context, tx *sqlx.Tx, req model.UserUpdateRequest) (bool, error) {
	locationInsertQuery := `INSERT INTO location (City, District, Ward, Description) 
							VALUES (:City, :District, :Ward, :Description)
							RETURNING ID`
	var locationID int
	locationParams := map[string]interface{}{
		"City":        req.Location.City,
		"District":    req.Location.District,
		"Ward":        req.Location.Ward,
		"Description": req.Location.Description,
	}
	query, args, err := tx.BindNamed(locationInsertQuery, locationParams)
	if err != nil {
		return false, err
	}

	// err = tx.QueryRowxContext(ctx, query, args).Scan(&locationID)
	err = tx.GetContext(ctx, &locationID, query, args...)
	if err != nil {
		logger.Error("UpdateUserInfo: fail to insert new location ", zap.Error(err))
		return false, err
	}

	// update locationID for user
	updateUserLocationIDQuery := `UPDATE user u SET u.LocationID = :LocationID WHERE ID = :ID`
	_, err = tx.NamedExecContext(ctx, updateUserLocationIDQuery, map[string]interface{}{
		"LocationID": locationID,
		"ID":         req.ID,
	})
	if err != nil {
		logger.Error("UpdateUserInfo: fail to update user LocationID ", zap.Error(err))
		return false, err
	}
	return true, nil
}

func updateUser(ctx context.Context, tx *sqlx.Tx, req model.UserUpdateRequest) error {
	userInfoQuery := `UPDATE user u
						SET 
						u.FullName = COALESCE(:FullName, u.FullName),
						u.Sex = COALESCE(:Sex, u.Sex),
						u.Bio = COALESCE(:Bio, u.Bio),
						u.UrlAvt = COALESCE(:UrlAvt, u.UrlAvt),
						u.UrlBackground = COALESCE(:UrlBackground, u.UrlBackground)
						WHERE u.ID = :ID`
	_, err := tx.NamedExecContext(ctx, userInfoQuery, req)
	if err != nil {
		logger.Error("UpdateUserInfo: Update user infor error: ", zap.Error(err))
		return err
	}
	return nil
}

func updateExistingLocation(ctx context.Context, tx *sqlx.Tx, req model.UserUpdateRequest) (bool, error) {
	locationUpdateQuery := `UPDATE location l
							SET 
							l.City = COALESCE(:City, l.City),
							l.District = COALESCE(:District, l.District),
							l.Ward = COALESCE(:Ward, l.Ward),
							l.Description = COALESCE(:Description, l.Description)
							WHERE ID = (SELECT u.LocationID FROM user u WHERE u.ID = :ID)`
	params := map[string]interface{}{
		"ID":          req.ID,
		"City":        req.Location.City,
		"District":    req.Location.District,
		"Ward":        req.Location.Ward,
		"Description": req.Location.Description,
	}
	_, err := tx.NamedExecContext(ctx, locationUpdateQuery, params)
	if err != nil {
		logger.Error("UpdateUserInfo: Update location of user infor error: ", zap.Error(err))
		return false, err
	}
	return true, nil
}

func handleLocation(ctx context.Context, tx *sqlx.Tx, req model.UserUpdateRequest) (bool, error) {
	var locationID *int
	err := tx.GetContext(ctx, &locationID, `SELECT LocationID FROM user WHERE ID = ?`, req.ID)
	if err != nil {
		logger.Error("UpdateUserInfo: fail to get user LocationID ", zap.Error(err))
		return false, err
	}

	if locationID == nil {
		// insert location
		return insertNewLocation(ctx, tx, req)
	} else {
		// update location
		return updateExistingLocation(ctx, tx, req)
	}
}

// UpdateUserInfo implements user.UserCommandRepository.
func (repo *userCommandRepository) UpdateUserInfo(ctx context.Context, req model.UserUpdateRequest) (bool, error) {
	tx, err := repo.db.BeginTxx(ctx, nil)
	if err != nil {
		logger.Error("UpdateUserInfo: fail to begin transaction ", zap.Error(err))
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			logger.Error("UpdateUserInfo: transaction rollback due to error ", zap.Error(err))
			return
		}
		err = tx.Commit()
		if err != nil {
			logger.Error("UpdateUserInfo: failed to commit transaction ", zap.Error(err))
		}

	}()

	err = updateUser(ctx, tx, req)
	if err != nil {
		return false, err
	}

	if req.Location == nil {
		return true, nil
	}

	success, err := handleLocation(ctx, tx, req)
	if err != nil {
		return false, err
	}
	return success, nil
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
