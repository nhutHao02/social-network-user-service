package imp

import (
	"context"

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

// GetUserInfo implements application.UserSerVice.
func (u *userService) GetUserInfo(c context.Context, userID int) (*model.UserInfoResponse, error) {
	res, err := u.userQueryRepository.GetUserInfo(c, userID)
	if err != nil {
		return nil, err
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

	token, err := token.CreateToken(string(user.ID))
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
