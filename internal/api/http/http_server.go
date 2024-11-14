package http

import (
	"github.com/gin-gonic/gin"
	"github.com/nhutHao02/social-network-common-service/utils/logger"
	"github.com/nhutHao02/social-network-user-service/config"
	v1 "github.com/nhutHao02/social-network-user-service/internal/api/http/v1"
	"go.uber.org/zap"
)

type HTTPServer struct {
	Cfg         *config.Config
	UserHandler *v1.UserHandler
}

func NewHTTPServer(cfg *config.Config, userHandler *v1.UserHandler) *HTTPServer {
	return &HTTPServer{Cfg: cfg, UserHandler: userHandler}
}

func (s *HTTPServer) RunHTTPServer() error {
	r := gin.Default()
	v1.MapRoutes(r, s.UserHandler)
	logger.Info("HTTP Server server listening at" + s.Cfg.HTTPServer.Address)
	err := r.Run(s.Cfg.HTTPServer.Address)
	if err != nil {
		logger.Error("HTTP Server error", zap.Error(err))
		return err
	}
	return nil
}
