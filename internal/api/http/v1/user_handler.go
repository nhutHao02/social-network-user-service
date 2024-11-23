package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	common "github.com/nhutHao02/social-network-common-service/model"
	"github.com/nhutHao02/social-network-common-service/request"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-common-service/utils/token"
	"github.com/nhutHao02/social-network-user-service/internal/application"
	"github.com/nhutHao02/social-network-user-service/internal/domain/model"
	"github.com/nhutHao02/social-network-user-service/pkg/constants"

	"go.uber.org/zap"
)

type UserHandler struct {
	userService application.UserSerVice
}

func NewUserHandler(userService application.UserSerVice) *UserHandler {
	return &UserHandler{userService: userService}
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login account
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		model.LoginRequest							true	"Login Request"
//	@Success		200		{object}	common.Response{data=model.LoginResponse}	"successfully"
//	@Failure		default	{object}	common.Response{data=nil}					"failure"
//	@Router			/guest/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req model.LoginRequest

	err := request.GetBodyJSON(c, &req)
	if err != nil {
		return
	}

	res, err := h.userService.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), "Login failure"))
		return
	}
	c.JSON(http.StatusOK, common.NewSuccessResponse(res))

}

// SignUp godoc
//
//	@Summary		SignUp
//	@Description	SignUp account
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		model.SignUpRequest			true	"Sign up Request"
//	@Success		200		{object}	common.Response{data=bool}	"successfully"
//	@Failure		default	{object}	common.Response{data=nil}	"failure"
//	@Router			/guest/sign-up [post]
func (h *UserHandler) SignUp(c *gin.Context) {
	var req model.SignUpRequest

	err := request.GetBodyJSON(c, &req)

	if err != nil {
		return
	}

	if err != nil {
		logger.Error("SignUp: hashPassword error: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), "Sign up new user failure"))
		return
	}

	success, err := h.userService.RegisterUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), "Sign up new user failure"))
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessResponse(success))
}

// GetUserInfo godoc
//
//	@Summary		GetUserInfo
//	@Description	Get User Information
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string											true	"Bearer <your_token>"
//	@Param			userID			path		int												true	"User ID"
//	@Success		200				{object}	common.Response{data=model.UserInfoResponse}	"successfully"
//	@Failure		default			{object}	common.Response{data=nil}						"failure"
//	@Router			/user/{userID} [get]
func (h *UserHandler) GetUserInfo(c *gin.Context) {
	var userParam model.UserParam
	err := request.GetParamsFromUrl(c, &userParam)
	if err != nil {
		return
	}

	res, err := h.userService.GetUserInfo(c.Request.Context(), int64(userParam.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), "Get user info failure"))
		return
	}
	c.JSON(http.StatusOK, common.NewSuccessResponse(res))
}

// UpdateUserInfo godoc
//
//	@Summary		UpdateUserInfo
//	@Description	Update User Information
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string							true	"Bearer <your_token>"
//	@Param			body			body		model.UserUpdateRequest			true	"UserUpdateRequest"
//	@Success		200				{object}	common.Response{data=boolean}	"successfully"
//	@Failure		default			{object}	common.Response{data=nil}		"failure"
//	@Router			/user [put]
func (h *UserHandler) UpdateUserInfo(c *gin.Context) {
	var req model.UserUpdateRequest
	err := request.GetBodyJSON(c, &req)
	if err != nil {
		return
	}
	userID, err := token.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), constants.UpdateUserInfoFailure))
		return
	}
	if int(req.ID) != userID {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(constants.InvalidUserID, constants.UpdateUserInfoFailure))
		return
	}

	success, err := h.userService.UpdateUserInfo(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), constants.UpdateUserInfoFailure))
		return
	}
	c.JSON(http.StatusOK, common.NewSuccessResponse(success))

}

// ChangePassword godoc
//
//	@Summary		ChangePassword
//	@Description	Change Password
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string							true	"Bearer <your_token>"
//	@Param			body			body		model.UserUpdatePassRequest		true	"UserUpdatePassRequest"
//	@Success		200				{object}	common.Response{data=boolean}	"successfully"
//	@Failure		default			{object}	common.Response{data=nil}		"failure"
//	@Router			/user/change-password [put]
func (h *UserHandler) ChangePassword(c *gin.Context) {
	var req model.UserUpdatePassRequest
	err := request.GetBodyJSON(c, &req)
	if err != nil {
		return
	}
	userID, err := token.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), constants.ChangePasswordFailure))
		return
	}
	if int(req.ID) != userID {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(constants.InvalidUserID, constants.ChangePasswordFailure))
		return
	}

	success, err := h.userService.ChangePassword(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), constants.ChangePasswordFailure))
		return
	}
	c.JSON(http.StatusOK, common.NewSuccessResponse(success))
}

// Follow godoc
//
//	@Summary		Follow
//	@Description	Follow
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string							true	"Bearer <your_token>"
//	@Param			body			body		model.FollowRequest				true	"FollowRequest"
//	@Success		200				{object}	common.Response{data=boolean}	"successfully"
//	@Failure		default			{object}	common.Response{data=nil}		"failure"
//	@Router			/user/follow [post]
func (h *UserHandler) Follow(c *gin.Context) {
	var req model.FollowRequest
	err := request.GetBodyJSON(c, &req)
	if err != nil {
		return
	}

	userID, err := token.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), constants.FollowFailure))
		return
	}
	if int(req.FollowerID) != userID {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(constants.InvalidUserID, constants.FollowFailure))
		return
	}

	success, err := h.userService.Follow(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), constants.FollowFailure))
		return
	}
	c.JSON(http.StatusOK, common.NewSuccessResponse(success))

}

// UnFollow godoc
//
//	@Summary		UnFollow
//	@Description	UnFollow
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string							true	"Bearer <your_token>"
//	@Param			body			body		model.FollowRequest				true	"FollowRequest"
//	@Success		200				{object}	common.Response{data=boolean}	"successfully"
//	@Failure		default			{object}	common.Response{data=nil}		"failure"
//	@Router			/user/un-follow [post]
func (h *UserHandler) UnFollow(c *gin.Context) {
	var req model.FollowRequest
	err := request.GetBodyJSON(c, &req)
	if err != nil {
		return
	}

	userID, err := token.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), constants.UnFollowFailure))
		return
	}
	if int(req.FollowerID) != userID {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(constants.InvalidUserID, constants.UnFollowFailure))
		return
	}

	success, err := h.userService.UnFollow(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), constants.UnFollowFailure))
		return
	}
	c.JSON(http.StatusOK, common.NewSuccessResponse(success))
}

// GetFollower godoc
//
//	@Summary		GetFollower
//	@Description	Get Follower
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string										true	"Bearer <your_token>"
//	@Param			userID			path		int											true	"User ID"
//	@Success		200				{object}	common.Response{data=model.FollowResponse}	"successfully"
//	@Failure		default			{object}	common.Response{data=nil}					"failure"
//	@Router			/user/follower/{userID} [get]
func (h *UserHandler) GetFollower(c *gin.Context) {
	getFollow(c, h, true)
}

// GetFollowing godoc
//
//	@Summary		GetFollowing
//	@Description	Get Following
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string										true	"Bearer <your_token>"
//	@Param			userID			path		int											true	"User ID"
//	@Success		200				{object}	common.Response{data=model.FollowResponse}	"successfully"
//	@Failure		default			{object}	common.Response{data=nil}					"failure"
//	@Router			/user/following/{userID} [get]
func (h *UserHandler) GetFollowing(c *gin.Context) {
	getFollow(c, h, false)
}

func getFollow(c *gin.Context, h *UserHandler, isFollower bool) {
	var idParam model.FollowIDParam
	err := request.GetParamsFromUrl(c, &idParam)
	if err != nil {
		return
	}

	data, err := h.userService.GetFollow(c.Request.Context(), idParam, isFollower)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), "Get Follow failure"))
		return
	}
	c.JSON(http.StatusOK, common.NewSuccessResponse(data))
}
