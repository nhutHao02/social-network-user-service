package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	common "github.com/nhutHao02/social-network-common-service/model"
	"github.com/nhutHao02/social-network-common-service/request"
	"github.com/nhutHao02/social-network-user-service/internal/application"
	"github.com/nhutHao02/social-network-user-service/internal/domain/model"
)

type UserHandler struct {
	userService application.UserSerVice
}

func NewUserHandler(userService application.UserSerVice) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Login(c *gin.Context) {

}

func (h *UserHandler) SignUp(c *gin.Context) {
	var req model.SignUpRequest

	err := request.GetBodyJSON(c, &req)

	if err != nil {
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
