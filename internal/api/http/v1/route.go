package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nhutHao02/social-network-common-service/middleware"
	"github.com/nhutHao02/social-network-common-service/utils/logger"

	"github.com/gin-contrib/cors"
	_ "github.com/nhutHao02/social-network-user-service/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MapRoutes(
	router *gin.Engine,
	userHandler *UserHandler,
) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
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

	v1Guest := router.Group("api/v1/guest")
	{
		v1Guest.POST("/login", userHandler.Login)
		v1Guest.POST("/sign-up", userHandler.SignUp)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
