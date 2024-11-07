package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nhutHao02/social-network-common-service/middleware"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
)

func MapRoutes(
	router *gin.Engine,
	userHandler *UserHandler,
) {
	v1 := router.Group("/api/v1")
	{
		v1.Use(middleware.JwtAuthMiddleware(logger.GetDefaultLogger()))
		// test api
		v1.GET("/ping", func(c *gin.Context) {
			logger.Info("this is log test")
			c.JSON(http.StatusOK, gin.H{
				"message": "middleware success",
			})
		})
	}

	v1Guest := router.Group("ap1/v1/guest")
	{
		v1Guest.POST("/login", userHandler.Login)
		v1Guest.POST("/sign-up", userHandler.SignUp)
	}

}
