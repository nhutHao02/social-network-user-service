package v1

import (
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
		vUser := v1.Group("/user")

		vUser.GET(":id", userHandler.GetUserInfo)
		vUser.GET("follower/:id", userHandler.GetFollower)
		vUser.GET("following/:id", userHandler.GetFollowing)

		vUser.PUT("", userHandler.UpdateUserInfo)
		vUser.PUT("change-password", userHandler.ChangePassword)

		vUser.POST("follow", userHandler.Follow)
		vUser.POST("un-follow", userHandler.UnFollow)
	}

	v1Guest := router.Group("ap1/v1/guest")
	{
		v1Guest.POST("/login", userHandler.Login)
		v1Guest.POST("/sign-up", userHandler.SignUp)
	}

}
