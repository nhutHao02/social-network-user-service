package imp

import (
	"context"
	"strings"
	"time"

	resError "github.com/nhutHao02/social-network-common-service/utils/error"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-common-service/utils/token"
	"github.com/nhutHao02/social-network-user-service/internal/application"
	"github.com/nhutHao02/social-network-user-service/internal/domain/interface/user"
	"github.com/nhutHao02/social-network-user-service/internal/domain/model"
	"github.com/nhutHao02/social-network-user-service/pkg/redis"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userQueryRepository   user.UserQueryRepository
	userCommandRepository user.UserCommandRepository
	cache                 *redis.RedisClient
}

// GetFollower implements application.UserSerVice.
func (u *userService) GetFollow(
	c context.Context, idParam model.FollowIDParam, isFollower bool) (*model.FollowResponse, error) {
	// check user
	_, err := u.userQueryRepository.CheckUserExistByID(c, idParam.ID)
	if err != nil {
		return nil, err
	}

	data, err := u.userQueryRepository.GetFollow(c, idParam.ID, isFollower)
	if err != nil {
		return nil, err
	}
	return data, nil

}

// UnFollow implements application.UserSerVice.
func (u *userService) UnFollow(c context.Context, req model.FollowRequest) (bool, error) {
	// check user
	_, err := u.userQueryRepository.CheckUserExistByID(c, req.FollowerID)
	if err != nil {
		return false, err
	}

	// check user
	_, err = u.userQueryRepository.CheckUserExistByID(c, req.FollowingID)
	if err != nil {
		return false, err
	}

	success, err := u.userCommandRepository.UnFollow(c, req)
	if err != nil {
		return false, err
	}
	return success, nil
}

// Follow implements application.UserSerVice.
func (u *userService) Follow(c context.Context, req model.FollowRequest) (bool, error) {
	// check user
	_, err := u.userQueryRepository.CheckUserExistByID(c, req.FollowerID)
	if err != nil {
		return false, err
	}

	// check user
	_, err = u.userQueryRepository.CheckUserExistByID(c, req.FollowingID)
	if err != nil {
		return false, err
	}

	success, err := u.userCommandRepository.Follow(c, req)
	if err != nil {
		return false, err
	}
	return success, nil
}

// ChangePassword implements application.UserSerVice.
func (u *userService) ChangePassword(c context.Context, req model.UserUpdatePassRequest) (bool, error) {
	hashPass, err := hashPassword(req.Password)
	if err != nil {
		logger.Error("RegisterUser: hash pasword error: ", zap.Error(err))
		return false, err
	}
	req.Password = hashPass
	success, err := u.userCommandRepository.UpdatePassword(c, req)
	if err != nil {
		return false, err
	}

	// clear cache
	err = u.cache.DeleteCache(c, string(req.ID))
	if err != nil {
		logger.Warn("UpdateUserInfo: Clear cache with userID error", zap.Error(err))
	}
	return success, nil
}

// UpdateUserInfo implements application.UserSerVice.
func (u *userService) UpdateUserInfo(c context.Context, req model.UserUpdateRequest) (bool, error) {
	success, err := u.userCommandRepository.UpdateUserInfo(c, req)
	if err != nil {
		return false, err
	}

	// clear cache
	err = u.cache.DeleteCache(c, string(req.ID))
	if err != nil {
		logger.Warn("UpdateUserInfo: Clear cache with userID error", zap.Error(err))
	}

	return success, nil
}

// GetUserInfo implements application.UserSerVice.
func (u *userService) GetUserInfo(c context.Context, userID int64) (*model.UserInfoResponse, error) {
	// check cache
	value, err := u.cache.GetCache(c, string(userID))
	if len(strings.TrimSpace(value)) != 0 && err == nil {
		var res model.UserInfoResponse
		err = u.cache.ConvertDataToStruct(&res, value)
		if err == nil {
			return &res, nil
		}
	}

	res, err := u.userQueryRepository.GetUserInfo(c, userID)
	if err != nil {
		return nil, err
	}

	// save cache
	err = u.cache.SetCacheStructData(c, string(userID), res, 24*time.Hour)
	if err != nil {
		logger.Warn("Save user info to cache error", zap.Error(err))
	}
	return res, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u *userService) RegisterUser(c context.Context, req model.SignUpRequest) (bool, error) {

	existed, err := u.userQueryRepository.CheckUserExisted(c, req.Email)
	if err != nil {
		return false, err
	}

	if existed {
		return false, resError.NewResError(nil, "Email existed")
	}
	hashPass, err := hashPassword(req.Password)
	if err != nil {
		logger.Error("RegisterUser: hash pasword error: ", zap.Error(err))
		return false, err
	}
	req.Password = hashPass

	success, err := u.userCommandRepository.RegisterUser(c, req)

	if err != nil {
		return false, err
	}

	return success, nil
}

// Login implements application.UserSerVice.
func (u *userService) Login(c context.Context, req model.LoginRequest) (model.LoginResponse, error) {
	var res model.LoginResponse

	user, err := u.userQueryRepository.Login(c, req)
	if err != nil {
		return res, err
	}

	veryfi := verifyPassword(req.Password, user.Password)
	if !veryfi {
		logger.Error("Login: invalid password: ")
		return res, resError.NewResError(nil, "Invalid password")
	}

	token, err := token.CreateToken(int(user.ID))
	if err != nil {
		logger.Error("Login: create token error: ", zap.Error(err))
		return res, err
	}

	res.ID = user.ID
	res.Token = token
	return res, nil
}

func NewUserService(
	userQueryRepository user.UserQueryRepository,
	userCommandRepository user.UserCommandRepository,
	cache *redis.RedisClient,
) application.UserSerVice {
	return &userService{
		userQueryRepository:   userQueryRepository,
		userCommandRepository: userCommandRepository,
		cache:                 cache,
	}
}
