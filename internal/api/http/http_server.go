package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nhutHao02/social-network-user-service/config"
	v1 "github.com/nhutHao02/social-network-user-service/internal/api/http/v1"
)

type HTTPServer struct {
	cfg         *config.Config
	userHandler *v1.UserHandler
}

func NewHTTPServer(cfg *config.Config, userHandler *v1.UserHandler) *HTTPServer {
	return &HTTPServer{cfg: cfg, userHandler: userHandler}
}

func (s *HTTPServer) RunHTTPServer() {
	r := gin.Default()
	v1.MapRoutes(r)
	r.Run(s.cfg.HTTPServer.Address)
}
