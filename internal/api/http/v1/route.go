package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-common-service/utils/token"
)

func MapRoutes(
	router *gin.Engine,
) {

	router.GET("/ping", func(c *gin.Context) {
		logger.Info("this is log test")
		tokenn, _ := token.CreateToken("1q234")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"token":   tokenn,
		})
	})
}
