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

	"go.uber.org/zap"
)

type UserHandler struct {
	userService application.UserSerVice
}

func NewUserHandler(userService application.UserSerVice) *UserHandler {
	return &UserHandler{userService: userService}
}

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

func (h *UserHandler) GetUserInfo(c *gin.Context) {
	var userParam model.UserParam
	request.GetParamsFromUrl(c, &userParam)

	res, err := h.userService.GetUserInfo(c.Request.Context(), userParam.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), "Get user info failure"))
		return
	}
	c.JSON(http.StatusOK, common.NewSuccessResponse(res))
}

func (h *UserHandler) UpdateUserInfo(c *gin.Context) {
	var req model.UserUpdateRequest
	err := request.GetBodyJSON(c, &req)
	if err != nil {
		return
	}
	userID, err := token.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), "Update user info failure "))
		return
	}
	if req.ID != int(userID) {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse("Invalid user ID", "Update user info failure"))
		return
	}

	success, err := h.userService.UpdateUserInfo(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewErrorResponse(err.Error(), "Update user info failure"))
		return
	}
	c.JSON(http.StatusOK, common.NewSuccessResponse(success))

}
