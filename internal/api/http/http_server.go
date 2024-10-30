package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nhutHao02/social-network-user-service/config"
)

type HTTPServer struct {
	cfg *config.Config
	// handlers
	// example addressHandler at V1
}

func NewHTTPSercer(cfg *config.Config) *HTTPServer {
	return &HTTPServer{cfg: cfg}
}

func RunHTTPServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

}
